package container

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"time"

	"code.cloudfoundry.org/winc/hcsclient"
	"code.cloudfoundry.org/winc/sandbox"
	"github.com/Microsoft/hcsshim"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

//go:generate counterfeiter . HCSContainer
type HCSContainer interface {
	Start() error
	Shutdown() error
	Terminate() error
	Wait() error
	WaitTimeout(time.Duration) error
	Pause() error
	Resume() error
	HasPendingUpdates() (bool, error)
	Statistics() (hcsshim.Statistics, error)
	ProcessList() ([]hcsshim.ProcessListItem, error)
	CreateProcess(c *hcsshim.ProcessConfig) (hcsshim.Process, error)
	OpenProcess(pid int) (hcsshim.Process, error)
	Close() error
	Modify(config *hcsshim.ResourceModificationRequestResponse) error
}

type ContainerManager interface {
	Create(rootfsPath string) error
	Delete() error
	State() (*specs.State, error)
}

type containerManager struct {
	hcsClient      hcsclient.Client
	sandboxManager sandbox.SandboxManager
	id             string
}

func NewManager(hcsClient hcsclient.Client, sandboxManager sandbox.SandboxManager, containerId string) ContainerManager {
	return &containerManager{
		hcsClient:      hcsClient,
		sandboxManager: sandboxManager,
		id:             containerId,
	}
}

func (c *containerManager) Create(rootfsPath string) error {
	_, err := c.hcsClient.GetContainerProperties(c.id)
	if err == nil {
		return &hcsclient.AlreadyExistsError{Id: c.id}
	}
	if _, ok := err.(*hcsclient.NotFoundError); !ok {
		return err
	}

	if err := c.sandboxManager.Create(rootfsPath); err != nil {
		return err
	}
	bundlePath := c.sandboxManager.BundlePath()

	if filepath.Base(bundlePath) != c.id {
		return &hcsclient.InvalidIdError{Id: c.id}
	}

	layerChain, err := ioutil.ReadFile(filepath.Join(bundlePath, "layerchain.json"))
	if err != nil {
		return err
	}

	layers := []string{}
	if err := json.Unmarshal(layerChain, &layers); err != nil {
		return err
	}

	layerInfos := []hcsshim.Layer{}
	for _, layerPath := range layers {
		layerId := filepath.Base(layerPath)
		layerGuid, err := c.hcsClient.NameToGuid(layerId)
		if err != nil {
			return err
		}

		layerInfos = append(layerInfos, hcsshim.Layer{
			ID:   layerGuid.ToString(),
			Path: layerPath,
		})
	}

	driverInfo := hcsshim.DriverInfo{
		HomeDir: filepath.Dir(bundlePath),
		Flavour: 1,
	}
	volumePath, err := c.hcsClient.GetLayerMountPath(driverInfo, c.id)
	if err != nil {
		return err
	} else if volumePath == "" {
		return &hcsclient.MissingVolumePathError{Id: c.id}
	}

	containerConfig := &hcsshim.ContainerConfig{
		SystemType:        "Container",
		Name:              bundlePath,
		VolumePath:        volumePath,
		Owner:             "winc",
		LayerFolderPath:   bundlePath,
		Layers:            layerInfos,
		MappedDirectories: []hcsshim.MappedDir{},
		EndpointList:      []string{},
	}

	_, err = c.hcsClient.CreateContainer(c.id, containerConfig)
	if err != nil {
		return err
	}

	return nil
}

func (c *containerManager) Delete() error {
	container, err := c.hcsClient.OpenContainer(c.id)
	if err != nil {
		return err
	}

	err = container.Terminate()
	if c.hcsClient.IsPending(err) {
		err = container.WaitTimeout(time.Minute * 5)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if err := c.sandboxManager.Delete(); err != nil {
		return err
	}

	return nil
}

func (c *containerManager) State() (*specs.State, error) {
	cp, err := c.hcsClient.GetContainerProperties(c.id)
	if err != nil {
		return nil, err
	}

	var status string
	if cp.Stopped {
		status = "stopped"
	} else {
		status = "created"
	}

	return &specs.State{
		Version: specs.Version,
		ID:      c.id,
		Status:  status,
		Bundle:  cp.Name,
	}, nil
}

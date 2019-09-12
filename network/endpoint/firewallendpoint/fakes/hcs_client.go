// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/winc/network/endpoint/firewallendpoint"
	"github.com/Microsoft/hcsshim"
)

type HCSClient struct {
	GetHNSNetworkByNameStub        func(string) (*hcsshim.HNSNetwork, error)
	getHNSNetworkByNameMutex       sync.RWMutex
	getHNSNetworkByNameArgsForCall []struct {
		arg1 string
	}
	getHNSNetworkByNameReturns struct {
		result1 *hcsshim.HNSNetwork
		result2 error
	}
	getHNSNetworkByNameReturnsOnCall map[int]struct {
		result1 *hcsshim.HNSNetwork
		result2 error
	}
	CreateEndpointStub        func(*hcsshim.HNSEndpoint) (*hcsshim.HNSEndpoint, error)
	createEndpointMutex       sync.RWMutex
	createEndpointArgsForCall []struct {
		arg1 *hcsshim.HNSEndpoint
	}
	createEndpointReturns struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}
	createEndpointReturnsOnCall map[int]struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}
	UpdateEndpointStub        func(*hcsshim.HNSEndpoint) (*hcsshim.HNSEndpoint, error)
	updateEndpointMutex       sync.RWMutex
	updateEndpointArgsForCall []struct {
		arg1 *hcsshim.HNSEndpoint
	}
	updateEndpointReturns struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}
	updateEndpointReturnsOnCall map[int]struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}
	GetHNSEndpointByIDStub        func(string) (*hcsshim.HNSEndpoint, error)
	getHNSEndpointByIDMutex       sync.RWMutex
	getHNSEndpointByIDArgsForCall []struct {
		arg1 string
	}
	getHNSEndpointByIDReturns struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}
	getHNSEndpointByIDReturnsOnCall map[int]struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}
	GetHNSEndpointByNameStub        func(string) (*hcsshim.HNSEndpoint, error)
	getHNSEndpointByNameMutex       sync.RWMutex
	getHNSEndpointByNameArgsForCall []struct {
		arg1 string
	}
	getHNSEndpointByNameReturns struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}
	getHNSEndpointByNameReturnsOnCall map[int]struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}
	DeleteEndpointStub        func(*hcsshim.HNSEndpoint) (*hcsshim.HNSEndpoint, error)
	deleteEndpointMutex       sync.RWMutex
	deleteEndpointArgsForCall []struct {
		arg1 *hcsshim.HNSEndpoint
	}
	deleteEndpointReturns struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}
	deleteEndpointReturnsOnCall map[int]struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}
	HotAttachEndpointStub        func(containerID string, endpointID string, endpointReady func() (bool, error)) error
	hotAttachEndpointMutex       sync.RWMutex
	hotAttachEndpointArgsForCall []struct {
		containerID   string
		endpointID    string
		endpointReady func() (bool, error)
	}
	hotAttachEndpointReturns struct {
		result1 error
	}
	hotAttachEndpointReturnsOnCall map[int]struct {
		result1 error
	}
	HotDetachEndpointStub        func(containerID string, endpointID string) error
	hotDetachEndpointMutex       sync.RWMutex
	hotDetachEndpointArgsForCall []struct {
		containerID string
		endpointID  string
	}
	hotDetachEndpointReturns struct {
		result1 error
	}
	hotDetachEndpointReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *HCSClient) GetHNSNetworkByName(arg1 string) (*hcsshim.HNSNetwork, error) {
	fake.getHNSNetworkByNameMutex.Lock()
	ret, specificReturn := fake.getHNSNetworkByNameReturnsOnCall[len(fake.getHNSNetworkByNameArgsForCall)]
	fake.getHNSNetworkByNameArgsForCall = append(fake.getHNSNetworkByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetHNSNetworkByName", []interface{}{arg1})
	fake.getHNSNetworkByNameMutex.Unlock()
	if fake.GetHNSNetworkByNameStub != nil {
		return fake.GetHNSNetworkByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getHNSNetworkByNameReturns.result1, fake.getHNSNetworkByNameReturns.result2
}

func (fake *HCSClient) GetHNSNetworkByNameCallCount() int {
	fake.getHNSNetworkByNameMutex.RLock()
	defer fake.getHNSNetworkByNameMutex.RUnlock()
	return len(fake.getHNSNetworkByNameArgsForCall)
}

func (fake *HCSClient) GetHNSNetworkByNameArgsForCall(i int) string {
	fake.getHNSNetworkByNameMutex.RLock()
	defer fake.getHNSNetworkByNameMutex.RUnlock()
	return fake.getHNSNetworkByNameArgsForCall[i].arg1
}

func (fake *HCSClient) GetHNSNetworkByNameReturns(result1 *hcsshim.HNSNetwork, result2 error) {
	fake.GetHNSNetworkByNameStub = nil
	fake.getHNSNetworkByNameReturns = struct {
		result1 *hcsshim.HNSNetwork
		result2 error
	}{result1, result2}
}

func (fake *HCSClient) GetHNSNetworkByNameReturnsOnCall(i int, result1 *hcsshim.HNSNetwork, result2 error) {
	fake.GetHNSNetworkByNameStub = nil
	if fake.getHNSNetworkByNameReturnsOnCall == nil {
		fake.getHNSNetworkByNameReturnsOnCall = make(map[int]struct {
			result1 *hcsshim.HNSNetwork
			result2 error
		})
	}
	fake.getHNSNetworkByNameReturnsOnCall[i] = struct {
		result1 *hcsshim.HNSNetwork
		result2 error
	}{result1, result2}
}

func (fake *HCSClient) CreateEndpoint(arg1 *hcsshim.HNSEndpoint) (*hcsshim.HNSEndpoint, error) {
	fake.createEndpointMutex.Lock()
	ret, specificReturn := fake.createEndpointReturnsOnCall[len(fake.createEndpointArgsForCall)]
	fake.createEndpointArgsForCall = append(fake.createEndpointArgsForCall, struct {
		arg1 *hcsshim.HNSEndpoint
	}{arg1})
	fake.recordInvocation("CreateEndpoint", []interface{}{arg1})
	fake.createEndpointMutex.Unlock()
	if fake.CreateEndpointStub != nil {
		return fake.CreateEndpointStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createEndpointReturns.result1, fake.createEndpointReturns.result2
}

func (fake *HCSClient) CreateEndpointCallCount() int {
	fake.createEndpointMutex.RLock()
	defer fake.createEndpointMutex.RUnlock()
	return len(fake.createEndpointArgsForCall)
}

func (fake *HCSClient) CreateEndpointArgsForCall(i int) *hcsshim.HNSEndpoint {
	fake.createEndpointMutex.RLock()
	defer fake.createEndpointMutex.RUnlock()
	return fake.createEndpointArgsForCall[i].arg1
}

func (fake *HCSClient) CreateEndpointReturns(result1 *hcsshim.HNSEndpoint, result2 error) {
	fake.CreateEndpointStub = nil
	fake.createEndpointReturns = struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}{result1, result2}
}

func (fake *HCSClient) CreateEndpointReturnsOnCall(i int, result1 *hcsshim.HNSEndpoint, result2 error) {
	fake.CreateEndpointStub = nil
	if fake.createEndpointReturnsOnCall == nil {
		fake.createEndpointReturnsOnCall = make(map[int]struct {
			result1 *hcsshim.HNSEndpoint
			result2 error
		})
	}
	fake.createEndpointReturnsOnCall[i] = struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}{result1, result2}
}

func (fake *HCSClient) UpdateEndpoint(arg1 *hcsshim.HNSEndpoint) (*hcsshim.HNSEndpoint, error) {
	fake.updateEndpointMutex.Lock()
	ret, specificReturn := fake.updateEndpointReturnsOnCall[len(fake.updateEndpointArgsForCall)]
	fake.updateEndpointArgsForCall = append(fake.updateEndpointArgsForCall, struct {
		arg1 *hcsshim.HNSEndpoint
	}{arg1})
	fake.recordInvocation("UpdateEndpoint", []interface{}{arg1})
	fake.updateEndpointMutex.Unlock()
	if fake.UpdateEndpointStub != nil {
		return fake.UpdateEndpointStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateEndpointReturns.result1, fake.updateEndpointReturns.result2
}

func (fake *HCSClient) UpdateEndpointCallCount() int {
	fake.updateEndpointMutex.RLock()
	defer fake.updateEndpointMutex.RUnlock()
	return len(fake.updateEndpointArgsForCall)
}

func (fake *HCSClient) UpdateEndpointArgsForCall(i int) *hcsshim.HNSEndpoint {
	fake.updateEndpointMutex.RLock()
	defer fake.updateEndpointMutex.RUnlock()
	return fake.updateEndpointArgsForCall[i].arg1
}

func (fake *HCSClient) UpdateEndpointReturns(result1 *hcsshim.HNSEndpoint, result2 error) {
	fake.UpdateEndpointStub = nil
	fake.updateEndpointReturns = struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}{result1, result2}
}

func (fake *HCSClient) UpdateEndpointReturnsOnCall(i int, result1 *hcsshim.HNSEndpoint, result2 error) {
	fake.UpdateEndpointStub = nil
	if fake.updateEndpointReturnsOnCall == nil {
		fake.updateEndpointReturnsOnCall = make(map[int]struct {
			result1 *hcsshim.HNSEndpoint
			result2 error
		})
	}
	fake.updateEndpointReturnsOnCall[i] = struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}{result1, result2}
}

func (fake *HCSClient) GetHNSEndpointByID(arg1 string) (*hcsshim.HNSEndpoint, error) {
	fake.getHNSEndpointByIDMutex.Lock()
	ret, specificReturn := fake.getHNSEndpointByIDReturnsOnCall[len(fake.getHNSEndpointByIDArgsForCall)]
	fake.getHNSEndpointByIDArgsForCall = append(fake.getHNSEndpointByIDArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetHNSEndpointByID", []interface{}{arg1})
	fake.getHNSEndpointByIDMutex.Unlock()
	if fake.GetHNSEndpointByIDStub != nil {
		return fake.GetHNSEndpointByIDStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getHNSEndpointByIDReturns.result1, fake.getHNSEndpointByIDReturns.result2
}

func (fake *HCSClient) GetHNSEndpointByIDCallCount() int {
	fake.getHNSEndpointByIDMutex.RLock()
	defer fake.getHNSEndpointByIDMutex.RUnlock()
	return len(fake.getHNSEndpointByIDArgsForCall)
}

func (fake *HCSClient) GetHNSEndpointByIDArgsForCall(i int) string {
	fake.getHNSEndpointByIDMutex.RLock()
	defer fake.getHNSEndpointByIDMutex.RUnlock()
	return fake.getHNSEndpointByIDArgsForCall[i].arg1
}

func (fake *HCSClient) GetHNSEndpointByIDReturns(result1 *hcsshim.HNSEndpoint, result2 error) {
	fake.GetHNSEndpointByIDStub = nil
	fake.getHNSEndpointByIDReturns = struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}{result1, result2}
}

func (fake *HCSClient) GetHNSEndpointByIDReturnsOnCall(i int, result1 *hcsshim.HNSEndpoint, result2 error) {
	fake.GetHNSEndpointByIDStub = nil
	if fake.getHNSEndpointByIDReturnsOnCall == nil {
		fake.getHNSEndpointByIDReturnsOnCall = make(map[int]struct {
			result1 *hcsshim.HNSEndpoint
			result2 error
		})
	}
	fake.getHNSEndpointByIDReturnsOnCall[i] = struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}{result1, result2}
}

func (fake *HCSClient) GetHNSEndpointByName(arg1 string) (*hcsshim.HNSEndpoint, error) {
	fake.getHNSEndpointByNameMutex.Lock()
	ret, specificReturn := fake.getHNSEndpointByNameReturnsOnCall[len(fake.getHNSEndpointByNameArgsForCall)]
	fake.getHNSEndpointByNameArgsForCall = append(fake.getHNSEndpointByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetHNSEndpointByName", []interface{}{arg1})
	fake.getHNSEndpointByNameMutex.Unlock()
	if fake.GetHNSEndpointByNameStub != nil {
		return fake.GetHNSEndpointByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getHNSEndpointByNameReturns.result1, fake.getHNSEndpointByNameReturns.result2
}

func (fake *HCSClient) GetHNSEndpointByNameCallCount() int {
	fake.getHNSEndpointByNameMutex.RLock()
	defer fake.getHNSEndpointByNameMutex.RUnlock()
	return len(fake.getHNSEndpointByNameArgsForCall)
}

func (fake *HCSClient) GetHNSEndpointByNameArgsForCall(i int) string {
	fake.getHNSEndpointByNameMutex.RLock()
	defer fake.getHNSEndpointByNameMutex.RUnlock()
	return fake.getHNSEndpointByNameArgsForCall[i].arg1
}

func (fake *HCSClient) GetHNSEndpointByNameReturns(result1 *hcsshim.HNSEndpoint, result2 error) {
	fake.GetHNSEndpointByNameStub = nil
	fake.getHNSEndpointByNameReturns = struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}{result1, result2}
}

func (fake *HCSClient) GetHNSEndpointByNameReturnsOnCall(i int, result1 *hcsshim.HNSEndpoint, result2 error) {
	fake.GetHNSEndpointByNameStub = nil
	if fake.getHNSEndpointByNameReturnsOnCall == nil {
		fake.getHNSEndpointByNameReturnsOnCall = make(map[int]struct {
			result1 *hcsshim.HNSEndpoint
			result2 error
		})
	}
	fake.getHNSEndpointByNameReturnsOnCall[i] = struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}{result1, result2}
}

func (fake *HCSClient) DeleteEndpoint(arg1 *hcsshim.HNSEndpoint) (*hcsshim.HNSEndpoint, error) {
	fake.deleteEndpointMutex.Lock()
	ret, specificReturn := fake.deleteEndpointReturnsOnCall[len(fake.deleteEndpointArgsForCall)]
	fake.deleteEndpointArgsForCall = append(fake.deleteEndpointArgsForCall, struct {
		arg1 *hcsshim.HNSEndpoint
	}{arg1})
	fake.recordInvocation("DeleteEndpoint", []interface{}{arg1})
	fake.deleteEndpointMutex.Unlock()
	if fake.DeleteEndpointStub != nil {
		return fake.DeleteEndpointStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deleteEndpointReturns.result1, fake.deleteEndpointReturns.result2
}

func (fake *HCSClient) DeleteEndpointCallCount() int {
	fake.deleteEndpointMutex.RLock()
	defer fake.deleteEndpointMutex.RUnlock()
	return len(fake.deleteEndpointArgsForCall)
}

func (fake *HCSClient) DeleteEndpointArgsForCall(i int) *hcsshim.HNSEndpoint {
	fake.deleteEndpointMutex.RLock()
	defer fake.deleteEndpointMutex.RUnlock()
	return fake.deleteEndpointArgsForCall[i].arg1
}

func (fake *HCSClient) DeleteEndpointReturns(result1 *hcsshim.HNSEndpoint, result2 error) {
	fake.DeleteEndpointStub = nil
	fake.deleteEndpointReturns = struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}{result1, result2}
}

func (fake *HCSClient) DeleteEndpointReturnsOnCall(i int, result1 *hcsshim.HNSEndpoint, result2 error) {
	fake.DeleteEndpointStub = nil
	if fake.deleteEndpointReturnsOnCall == nil {
		fake.deleteEndpointReturnsOnCall = make(map[int]struct {
			result1 *hcsshim.HNSEndpoint
			result2 error
		})
	}
	fake.deleteEndpointReturnsOnCall[i] = struct {
		result1 *hcsshim.HNSEndpoint
		result2 error
	}{result1, result2}
}

func (fake *HCSClient) HotAttachEndpoint(containerID string, endpointID string, endpointReady func() (bool, error)) error {
	fake.hotAttachEndpointMutex.Lock()
	ret, specificReturn := fake.hotAttachEndpointReturnsOnCall[len(fake.hotAttachEndpointArgsForCall)]
	fake.hotAttachEndpointArgsForCall = append(fake.hotAttachEndpointArgsForCall, struct {
		containerID   string
		endpointID    string
		endpointReady func() (bool, error)
	}{containerID, endpointID, endpointReady})
	fake.recordInvocation("HotAttachEndpoint", []interface{}{containerID, endpointID, endpointReady})
	fake.hotAttachEndpointMutex.Unlock()
	if fake.HotAttachEndpointStub != nil {
		return fake.HotAttachEndpointStub(containerID, endpointID, endpointReady)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.hotAttachEndpointReturns.result1
}

func (fake *HCSClient) HotAttachEndpointCallCount() int {
	fake.hotAttachEndpointMutex.RLock()
	defer fake.hotAttachEndpointMutex.RUnlock()
	return len(fake.hotAttachEndpointArgsForCall)
}

func (fake *HCSClient) HotAttachEndpointArgsForCall(i int) (string, string, func() (bool, error)) {
	fake.hotAttachEndpointMutex.RLock()
	defer fake.hotAttachEndpointMutex.RUnlock()
	return fake.hotAttachEndpointArgsForCall[i].containerID, fake.hotAttachEndpointArgsForCall[i].endpointID, fake.hotAttachEndpointArgsForCall[i].endpointReady
}

func (fake *HCSClient) HotAttachEndpointReturns(result1 error) {
	fake.HotAttachEndpointStub = nil
	fake.hotAttachEndpointReturns = struct {
		result1 error
	}{result1}
}

func (fake *HCSClient) HotAttachEndpointReturnsOnCall(i int, result1 error) {
	fake.HotAttachEndpointStub = nil
	if fake.hotAttachEndpointReturnsOnCall == nil {
		fake.hotAttachEndpointReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.hotAttachEndpointReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *HCSClient) HotDetachEndpoint(containerID string, endpointID string) error {
	fake.hotDetachEndpointMutex.Lock()
	ret, specificReturn := fake.hotDetachEndpointReturnsOnCall[len(fake.hotDetachEndpointArgsForCall)]
	fake.hotDetachEndpointArgsForCall = append(fake.hotDetachEndpointArgsForCall, struct {
		containerID string
		endpointID  string
	}{containerID, endpointID})
	fake.recordInvocation("HotDetachEndpoint", []interface{}{containerID, endpointID})
	fake.hotDetachEndpointMutex.Unlock()
	if fake.HotDetachEndpointStub != nil {
		return fake.HotDetachEndpointStub(containerID, endpointID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.hotDetachEndpointReturns.result1
}

func (fake *HCSClient) HotDetachEndpointCallCount() int {
	fake.hotDetachEndpointMutex.RLock()
	defer fake.hotDetachEndpointMutex.RUnlock()
	return len(fake.hotDetachEndpointArgsForCall)
}

func (fake *HCSClient) HotDetachEndpointArgsForCall(i int) (string, string) {
	fake.hotDetachEndpointMutex.RLock()
	defer fake.hotDetachEndpointMutex.RUnlock()
	return fake.hotDetachEndpointArgsForCall[i].containerID, fake.hotDetachEndpointArgsForCall[i].endpointID
}

func (fake *HCSClient) HotDetachEndpointReturns(result1 error) {
	fake.HotDetachEndpointStub = nil
	fake.hotDetachEndpointReturns = struct {
		result1 error
	}{result1}
}

func (fake *HCSClient) HotDetachEndpointReturnsOnCall(i int, result1 error) {
	fake.HotDetachEndpointStub = nil
	if fake.hotDetachEndpointReturnsOnCall == nil {
		fake.hotDetachEndpointReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.hotDetachEndpointReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *HCSClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getHNSNetworkByNameMutex.RLock()
	defer fake.getHNSNetworkByNameMutex.RUnlock()
	fake.createEndpointMutex.RLock()
	defer fake.createEndpointMutex.RUnlock()
	fake.updateEndpointMutex.RLock()
	defer fake.updateEndpointMutex.RUnlock()
	fake.getHNSEndpointByIDMutex.RLock()
	defer fake.getHNSEndpointByIDMutex.RUnlock()
	fake.getHNSEndpointByNameMutex.RLock()
	defer fake.getHNSEndpointByNameMutex.RUnlock()
	fake.deleteEndpointMutex.RLock()
	defer fake.deleteEndpointMutex.RUnlock()
	fake.hotAttachEndpointMutex.RLock()
	defer fake.hotAttachEndpointMutex.RUnlock()
	fake.hotDetachEndpointMutex.RLock()
	defer fake.hotDetachEndpointMutex.RUnlock()
	return fake.invocations
}

func (fake *HCSClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ firewallendpoint.HCSClient = new(HCSClient)

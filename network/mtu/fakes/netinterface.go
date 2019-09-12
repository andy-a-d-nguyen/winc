// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/winc/network/mtu"
	"code.cloudfoundry.org/winc/network/netinterface"
)

type NetInterface struct {
	ByNameStub        func(string) (netinterface.AdapterInfo, error)
	byNameMutex       sync.RWMutex
	byNameArgsForCall []struct {
		arg1 string
	}
	byNameReturns struct {
		result1 netinterface.AdapterInfo
		result2 error
	}
	byNameReturnsOnCall map[int]struct {
		result1 netinterface.AdapterInfo
		result2 error
	}
	ByIPStub        func(string) (netinterface.AdapterInfo, error)
	byIPMutex       sync.RWMutex
	byIPArgsForCall []struct {
		arg1 string
	}
	byIPReturns struct {
		result1 netinterface.AdapterInfo
		result2 error
	}
	byIPReturnsOnCall map[int]struct {
		result1 netinterface.AdapterInfo
		result2 error
	}
	SetMTUStub        func(string, uint32, uint32) error
	setMTUMutex       sync.RWMutex
	setMTUArgsForCall []struct {
		arg1 string
		arg2 uint32
		arg3 uint32
	}
	setMTUReturns struct {
		result1 error
	}
	setMTUReturnsOnCall map[int]struct {
		result1 error
	}
	GetMTUStub        func(string, uint32) (uint32, error)
	getMTUMutex       sync.RWMutex
	getMTUArgsForCall []struct {
		arg1 string
		arg2 uint32
	}
	getMTUReturns struct {
		result1 uint32
		result2 error
	}
	getMTUReturnsOnCall map[int]struct {
		result1 uint32
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NetInterface) ByName(arg1 string) (netinterface.AdapterInfo, error) {
	fake.byNameMutex.Lock()
	ret, specificReturn := fake.byNameReturnsOnCall[len(fake.byNameArgsForCall)]
	fake.byNameArgsForCall = append(fake.byNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ByName", []interface{}{arg1})
	fake.byNameMutex.Unlock()
	if fake.ByNameStub != nil {
		return fake.ByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.byNameReturns.result1, fake.byNameReturns.result2
}

func (fake *NetInterface) ByNameCallCount() int {
	fake.byNameMutex.RLock()
	defer fake.byNameMutex.RUnlock()
	return len(fake.byNameArgsForCall)
}

func (fake *NetInterface) ByNameArgsForCall(i int) string {
	fake.byNameMutex.RLock()
	defer fake.byNameMutex.RUnlock()
	return fake.byNameArgsForCall[i].arg1
}

func (fake *NetInterface) ByNameReturns(result1 netinterface.AdapterInfo, result2 error) {
	fake.ByNameStub = nil
	fake.byNameReturns = struct {
		result1 netinterface.AdapterInfo
		result2 error
	}{result1, result2}
}

func (fake *NetInterface) ByNameReturnsOnCall(i int, result1 netinterface.AdapterInfo, result2 error) {
	fake.ByNameStub = nil
	if fake.byNameReturnsOnCall == nil {
		fake.byNameReturnsOnCall = make(map[int]struct {
			result1 netinterface.AdapterInfo
			result2 error
		})
	}
	fake.byNameReturnsOnCall[i] = struct {
		result1 netinterface.AdapterInfo
		result2 error
	}{result1, result2}
}

func (fake *NetInterface) ByIP(arg1 string) (netinterface.AdapterInfo, error) {
	fake.byIPMutex.Lock()
	ret, specificReturn := fake.byIPReturnsOnCall[len(fake.byIPArgsForCall)]
	fake.byIPArgsForCall = append(fake.byIPArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ByIP", []interface{}{arg1})
	fake.byIPMutex.Unlock()
	if fake.ByIPStub != nil {
		return fake.ByIPStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.byIPReturns.result1, fake.byIPReturns.result2
}

func (fake *NetInterface) ByIPCallCount() int {
	fake.byIPMutex.RLock()
	defer fake.byIPMutex.RUnlock()
	return len(fake.byIPArgsForCall)
}

func (fake *NetInterface) ByIPArgsForCall(i int) string {
	fake.byIPMutex.RLock()
	defer fake.byIPMutex.RUnlock()
	return fake.byIPArgsForCall[i].arg1
}

func (fake *NetInterface) ByIPReturns(result1 netinterface.AdapterInfo, result2 error) {
	fake.ByIPStub = nil
	fake.byIPReturns = struct {
		result1 netinterface.AdapterInfo
		result2 error
	}{result1, result2}
}

func (fake *NetInterface) ByIPReturnsOnCall(i int, result1 netinterface.AdapterInfo, result2 error) {
	fake.ByIPStub = nil
	if fake.byIPReturnsOnCall == nil {
		fake.byIPReturnsOnCall = make(map[int]struct {
			result1 netinterface.AdapterInfo
			result2 error
		})
	}
	fake.byIPReturnsOnCall[i] = struct {
		result1 netinterface.AdapterInfo
		result2 error
	}{result1, result2}
}

func (fake *NetInterface) SetMTU(arg1 string, arg2 uint32, arg3 uint32) error {
	fake.setMTUMutex.Lock()
	ret, specificReturn := fake.setMTUReturnsOnCall[len(fake.setMTUArgsForCall)]
	fake.setMTUArgsForCall = append(fake.setMTUArgsForCall, struct {
		arg1 string
		arg2 uint32
		arg3 uint32
	}{arg1, arg2, arg3})
	fake.recordInvocation("SetMTU", []interface{}{arg1, arg2, arg3})
	fake.setMTUMutex.Unlock()
	if fake.SetMTUStub != nil {
		return fake.SetMTUStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setMTUReturns.result1
}

func (fake *NetInterface) SetMTUCallCount() int {
	fake.setMTUMutex.RLock()
	defer fake.setMTUMutex.RUnlock()
	return len(fake.setMTUArgsForCall)
}

func (fake *NetInterface) SetMTUArgsForCall(i int) (string, uint32, uint32) {
	fake.setMTUMutex.RLock()
	defer fake.setMTUMutex.RUnlock()
	return fake.setMTUArgsForCall[i].arg1, fake.setMTUArgsForCall[i].arg2, fake.setMTUArgsForCall[i].arg3
}

func (fake *NetInterface) SetMTUReturns(result1 error) {
	fake.SetMTUStub = nil
	fake.setMTUReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetInterface) SetMTUReturnsOnCall(i int, result1 error) {
	fake.SetMTUStub = nil
	if fake.setMTUReturnsOnCall == nil {
		fake.setMTUReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setMTUReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetInterface) GetMTU(arg1 string, arg2 uint32) (uint32, error) {
	fake.getMTUMutex.Lock()
	ret, specificReturn := fake.getMTUReturnsOnCall[len(fake.getMTUArgsForCall)]
	fake.getMTUArgsForCall = append(fake.getMTUArgsForCall, struct {
		arg1 string
		arg2 uint32
	}{arg1, arg2})
	fake.recordInvocation("GetMTU", []interface{}{arg1, arg2})
	fake.getMTUMutex.Unlock()
	if fake.GetMTUStub != nil {
		return fake.GetMTUStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getMTUReturns.result1, fake.getMTUReturns.result2
}

func (fake *NetInterface) GetMTUCallCount() int {
	fake.getMTUMutex.RLock()
	defer fake.getMTUMutex.RUnlock()
	return len(fake.getMTUArgsForCall)
}

func (fake *NetInterface) GetMTUArgsForCall(i int) (string, uint32) {
	fake.getMTUMutex.RLock()
	defer fake.getMTUMutex.RUnlock()
	return fake.getMTUArgsForCall[i].arg1, fake.getMTUArgsForCall[i].arg2
}

func (fake *NetInterface) GetMTUReturns(result1 uint32, result2 error) {
	fake.GetMTUStub = nil
	fake.getMTUReturns = struct {
		result1 uint32
		result2 error
	}{result1, result2}
}

func (fake *NetInterface) GetMTUReturnsOnCall(i int, result1 uint32, result2 error) {
	fake.GetMTUStub = nil
	if fake.getMTUReturnsOnCall == nil {
		fake.getMTUReturnsOnCall = make(map[int]struct {
			result1 uint32
			result2 error
		})
	}
	fake.getMTUReturnsOnCall[i] = struct {
		result1 uint32
		result2 error
	}{result1, result2}
}

func (fake *NetInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.byNameMutex.RLock()
	defer fake.byNameMutex.RUnlock()
	fake.byIPMutex.RLock()
	defer fake.byIPMutex.RUnlock()
	fake.setMTUMutex.RLock()
	defer fake.setMTUMutex.RUnlock()
	fake.getMTUMutex.RLock()
	defer fake.getMTUMutex.RUnlock()
	return fake.invocations
}

func (fake *NetInterface) recordInvocation(key string, args []interface{}) {
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

var _ mtu.NetInterface = new(NetInterface)

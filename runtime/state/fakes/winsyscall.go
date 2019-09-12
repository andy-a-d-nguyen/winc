// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
	"syscall"

	"code.cloudfoundry.org/winc/runtime/state"
)

type WinSyscall struct {
	OpenProcessStub        func(uint32, bool, uint32) (syscall.Handle, error)
	openProcessMutex       sync.RWMutex
	openProcessArgsForCall []struct {
		arg1 uint32
		arg2 bool
		arg3 uint32
	}
	openProcessReturns struct {
		result1 syscall.Handle
		result2 error
	}
	openProcessReturnsOnCall map[int]struct {
		result1 syscall.Handle
		result2 error
	}
	GetProcessStartTimeStub        func(syscall.Handle) (syscall.Filetime, error)
	getProcessStartTimeMutex       sync.RWMutex
	getProcessStartTimeArgsForCall []struct {
		arg1 syscall.Handle
	}
	getProcessStartTimeReturns struct {
		result1 syscall.Filetime
		result2 error
	}
	getProcessStartTimeReturnsOnCall map[int]struct {
		result1 syscall.Filetime
		result2 error
	}
	CloseHandleStub        func(syscall.Handle) error
	closeHandleMutex       sync.RWMutex
	closeHandleArgsForCall []struct {
		arg1 syscall.Handle
	}
	closeHandleReturns struct {
		result1 error
	}
	closeHandleReturnsOnCall map[int]struct {
		result1 error
	}
	GetExitCodeProcessStub        func(syscall.Handle) (uint32, error)
	getExitCodeProcessMutex       sync.RWMutex
	getExitCodeProcessArgsForCall []struct {
		arg1 syscall.Handle
	}
	getExitCodeProcessReturns struct {
		result1 uint32
		result2 error
	}
	getExitCodeProcessReturnsOnCall map[int]struct {
		result1 uint32
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *WinSyscall) OpenProcess(arg1 uint32, arg2 bool, arg3 uint32) (syscall.Handle, error) {
	fake.openProcessMutex.Lock()
	ret, specificReturn := fake.openProcessReturnsOnCall[len(fake.openProcessArgsForCall)]
	fake.openProcessArgsForCall = append(fake.openProcessArgsForCall, struct {
		arg1 uint32
		arg2 bool
		arg3 uint32
	}{arg1, arg2, arg3})
	fake.recordInvocation("OpenProcess", []interface{}{arg1, arg2, arg3})
	fake.openProcessMutex.Unlock()
	if fake.OpenProcessStub != nil {
		return fake.OpenProcessStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.openProcessReturns.result1, fake.openProcessReturns.result2
}

func (fake *WinSyscall) OpenProcessCallCount() int {
	fake.openProcessMutex.RLock()
	defer fake.openProcessMutex.RUnlock()
	return len(fake.openProcessArgsForCall)
}

func (fake *WinSyscall) OpenProcessArgsForCall(i int) (uint32, bool, uint32) {
	fake.openProcessMutex.RLock()
	defer fake.openProcessMutex.RUnlock()
	return fake.openProcessArgsForCall[i].arg1, fake.openProcessArgsForCall[i].arg2, fake.openProcessArgsForCall[i].arg3
}

func (fake *WinSyscall) OpenProcessReturns(result1 syscall.Handle, result2 error) {
	fake.OpenProcessStub = nil
	fake.openProcessReturns = struct {
		result1 syscall.Handle
		result2 error
	}{result1, result2}
}

func (fake *WinSyscall) OpenProcessReturnsOnCall(i int, result1 syscall.Handle, result2 error) {
	fake.OpenProcessStub = nil
	if fake.openProcessReturnsOnCall == nil {
		fake.openProcessReturnsOnCall = make(map[int]struct {
			result1 syscall.Handle
			result2 error
		})
	}
	fake.openProcessReturnsOnCall[i] = struct {
		result1 syscall.Handle
		result2 error
	}{result1, result2}
}

func (fake *WinSyscall) GetProcessStartTime(arg1 syscall.Handle) (syscall.Filetime, error) {
	fake.getProcessStartTimeMutex.Lock()
	ret, specificReturn := fake.getProcessStartTimeReturnsOnCall[len(fake.getProcessStartTimeArgsForCall)]
	fake.getProcessStartTimeArgsForCall = append(fake.getProcessStartTimeArgsForCall, struct {
		arg1 syscall.Handle
	}{arg1})
	fake.recordInvocation("GetProcessStartTime", []interface{}{arg1})
	fake.getProcessStartTimeMutex.Unlock()
	if fake.GetProcessStartTimeStub != nil {
		return fake.GetProcessStartTimeStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getProcessStartTimeReturns.result1, fake.getProcessStartTimeReturns.result2
}

func (fake *WinSyscall) GetProcessStartTimeCallCount() int {
	fake.getProcessStartTimeMutex.RLock()
	defer fake.getProcessStartTimeMutex.RUnlock()
	return len(fake.getProcessStartTimeArgsForCall)
}

func (fake *WinSyscall) GetProcessStartTimeArgsForCall(i int) syscall.Handle {
	fake.getProcessStartTimeMutex.RLock()
	defer fake.getProcessStartTimeMutex.RUnlock()
	return fake.getProcessStartTimeArgsForCall[i].arg1
}

func (fake *WinSyscall) GetProcessStartTimeReturns(result1 syscall.Filetime, result2 error) {
	fake.GetProcessStartTimeStub = nil
	fake.getProcessStartTimeReturns = struct {
		result1 syscall.Filetime
		result2 error
	}{result1, result2}
}

func (fake *WinSyscall) GetProcessStartTimeReturnsOnCall(i int, result1 syscall.Filetime, result2 error) {
	fake.GetProcessStartTimeStub = nil
	if fake.getProcessStartTimeReturnsOnCall == nil {
		fake.getProcessStartTimeReturnsOnCall = make(map[int]struct {
			result1 syscall.Filetime
			result2 error
		})
	}
	fake.getProcessStartTimeReturnsOnCall[i] = struct {
		result1 syscall.Filetime
		result2 error
	}{result1, result2}
}

func (fake *WinSyscall) CloseHandle(arg1 syscall.Handle) error {
	fake.closeHandleMutex.Lock()
	ret, specificReturn := fake.closeHandleReturnsOnCall[len(fake.closeHandleArgsForCall)]
	fake.closeHandleArgsForCall = append(fake.closeHandleArgsForCall, struct {
		arg1 syscall.Handle
	}{arg1})
	fake.recordInvocation("CloseHandle", []interface{}{arg1})
	fake.closeHandleMutex.Unlock()
	if fake.CloseHandleStub != nil {
		return fake.CloseHandleStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.closeHandleReturns.result1
}

func (fake *WinSyscall) CloseHandleCallCount() int {
	fake.closeHandleMutex.RLock()
	defer fake.closeHandleMutex.RUnlock()
	return len(fake.closeHandleArgsForCall)
}

func (fake *WinSyscall) CloseHandleArgsForCall(i int) syscall.Handle {
	fake.closeHandleMutex.RLock()
	defer fake.closeHandleMutex.RUnlock()
	return fake.closeHandleArgsForCall[i].arg1
}

func (fake *WinSyscall) CloseHandleReturns(result1 error) {
	fake.CloseHandleStub = nil
	fake.closeHandleReturns = struct {
		result1 error
	}{result1}
}

func (fake *WinSyscall) CloseHandleReturnsOnCall(i int, result1 error) {
	fake.CloseHandleStub = nil
	if fake.closeHandleReturnsOnCall == nil {
		fake.closeHandleReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeHandleReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *WinSyscall) GetExitCodeProcess(arg1 syscall.Handle) (uint32, error) {
	fake.getExitCodeProcessMutex.Lock()
	ret, specificReturn := fake.getExitCodeProcessReturnsOnCall[len(fake.getExitCodeProcessArgsForCall)]
	fake.getExitCodeProcessArgsForCall = append(fake.getExitCodeProcessArgsForCall, struct {
		arg1 syscall.Handle
	}{arg1})
	fake.recordInvocation("GetExitCodeProcess", []interface{}{arg1})
	fake.getExitCodeProcessMutex.Unlock()
	if fake.GetExitCodeProcessStub != nil {
		return fake.GetExitCodeProcessStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getExitCodeProcessReturns.result1, fake.getExitCodeProcessReturns.result2
}

func (fake *WinSyscall) GetExitCodeProcessCallCount() int {
	fake.getExitCodeProcessMutex.RLock()
	defer fake.getExitCodeProcessMutex.RUnlock()
	return len(fake.getExitCodeProcessArgsForCall)
}

func (fake *WinSyscall) GetExitCodeProcessArgsForCall(i int) syscall.Handle {
	fake.getExitCodeProcessMutex.RLock()
	defer fake.getExitCodeProcessMutex.RUnlock()
	return fake.getExitCodeProcessArgsForCall[i].arg1
}

func (fake *WinSyscall) GetExitCodeProcessReturns(result1 uint32, result2 error) {
	fake.GetExitCodeProcessStub = nil
	fake.getExitCodeProcessReturns = struct {
		result1 uint32
		result2 error
	}{result1, result2}
}

func (fake *WinSyscall) GetExitCodeProcessReturnsOnCall(i int, result1 uint32, result2 error) {
	fake.GetExitCodeProcessStub = nil
	if fake.getExitCodeProcessReturnsOnCall == nil {
		fake.getExitCodeProcessReturnsOnCall = make(map[int]struct {
			result1 uint32
			result2 error
		})
	}
	fake.getExitCodeProcessReturnsOnCall[i] = struct {
		result1 uint32
		result2 error
	}{result1, result2}
}

func (fake *WinSyscall) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.openProcessMutex.RLock()
	defer fake.openProcessMutex.RUnlock()
	fake.getProcessStartTimeMutex.RLock()
	defer fake.getProcessStartTimeMutex.RUnlock()
	fake.closeHandleMutex.RLock()
	defer fake.closeHandleMutex.RUnlock()
	fake.getExitCodeProcessMutex.RLock()
	defer fake.getExitCodeProcessMutex.RUnlock()
	return fake.invocations
}

func (fake *WinSyscall) recordInvocation(key string, args []interface{}) {
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

var _ state.WinSyscall = new(WinSyscall)

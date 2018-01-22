// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/winc/image"
)

type Statser struct {
	GetCurrentDiskUsageStub        func(string) (uint64, error)
	getCurrentDiskUsageMutex       sync.RWMutex
	getCurrentDiskUsageArgsForCall []struct {
		arg1 string
	}
	getCurrentDiskUsageReturns struct {
		result1 uint64
		result2 error
	}
	getCurrentDiskUsageReturnsOnCall map[int]struct {
		result1 uint64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Statser) GetCurrentDiskUsage(arg1 string) (uint64, error) {
	fake.getCurrentDiskUsageMutex.Lock()
	ret, specificReturn := fake.getCurrentDiskUsageReturnsOnCall[len(fake.getCurrentDiskUsageArgsForCall)]
	fake.getCurrentDiskUsageArgsForCall = append(fake.getCurrentDiskUsageArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetCurrentDiskUsage", []interface{}{arg1})
	fake.getCurrentDiskUsageMutex.Unlock()
	if fake.GetCurrentDiskUsageStub != nil {
		return fake.GetCurrentDiskUsageStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getCurrentDiskUsageReturns.result1, fake.getCurrentDiskUsageReturns.result2
}

func (fake *Statser) GetCurrentDiskUsageCallCount() int {
	fake.getCurrentDiskUsageMutex.RLock()
	defer fake.getCurrentDiskUsageMutex.RUnlock()
	return len(fake.getCurrentDiskUsageArgsForCall)
}

func (fake *Statser) GetCurrentDiskUsageArgsForCall(i int) string {
	fake.getCurrentDiskUsageMutex.RLock()
	defer fake.getCurrentDiskUsageMutex.RUnlock()
	return fake.getCurrentDiskUsageArgsForCall[i].arg1
}

func (fake *Statser) GetCurrentDiskUsageReturns(result1 uint64, result2 error) {
	fake.GetCurrentDiskUsageStub = nil
	fake.getCurrentDiskUsageReturns = struct {
		result1 uint64
		result2 error
	}{result1, result2}
}

func (fake *Statser) GetCurrentDiskUsageReturnsOnCall(i int, result1 uint64, result2 error) {
	fake.GetCurrentDiskUsageStub = nil
	if fake.getCurrentDiskUsageReturnsOnCall == nil {
		fake.getCurrentDiskUsageReturnsOnCall = make(map[int]struct {
			result1 uint64
			result2 error
		})
	}
	fake.getCurrentDiskUsageReturnsOnCall[i] = struct {
		result1 uint64
		result2 error
	}{result1, result2}
}

func (fake *Statser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getCurrentDiskUsageMutex.RLock()
	defer fake.getCurrentDiskUsageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Statser) recordInvocation(key string, args []interface{}) {
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

var _ image.Statser = new(Statser)
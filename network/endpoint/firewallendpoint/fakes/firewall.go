// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/winc/network/endpoint/firewallendpoint"
)

type Firewall struct {
	DeleteRuleStub        func(string) error
	deleteRuleMutex       sync.RWMutex
	deleteRuleArgsForCall []struct {
		arg1 string
	}
	deleteRuleReturns struct {
		result1 error
	}
	deleteRuleReturnsOnCall map[int]struct {
		result1 error
	}
	RuleExistsStub        func(string) (bool, error)
	ruleExistsMutex       sync.RWMutex
	ruleExistsArgsForCall []struct {
		arg1 string
	}
	ruleExistsReturns struct {
		result1 bool
		result2 error
	}
	ruleExistsReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Firewall) DeleteRule(arg1 string) error {
	fake.deleteRuleMutex.Lock()
	ret, specificReturn := fake.deleteRuleReturnsOnCall[len(fake.deleteRuleArgsForCall)]
	fake.deleteRuleArgsForCall = append(fake.deleteRuleArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteRule", []interface{}{arg1})
	fake.deleteRuleMutex.Unlock()
	if fake.DeleteRuleStub != nil {
		return fake.DeleteRuleStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteRuleReturns.result1
}

func (fake *Firewall) DeleteRuleCallCount() int {
	fake.deleteRuleMutex.RLock()
	defer fake.deleteRuleMutex.RUnlock()
	return len(fake.deleteRuleArgsForCall)
}

func (fake *Firewall) DeleteRuleArgsForCall(i int) string {
	fake.deleteRuleMutex.RLock()
	defer fake.deleteRuleMutex.RUnlock()
	return fake.deleteRuleArgsForCall[i].arg1
}

func (fake *Firewall) DeleteRuleReturns(result1 error) {
	fake.DeleteRuleStub = nil
	fake.deleteRuleReturns = struct {
		result1 error
	}{result1}
}

func (fake *Firewall) DeleteRuleReturnsOnCall(i int, result1 error) {
	fake.DeleteRuleStub = nil
	if fake.deleteRuleReturnsOnCall == nil {
		fake.deleteRuleReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteRuleReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Firewall) RuleExists(arg1 string) (bool, error) {
	fake.ruleExistsMutex.Lock()
	ret, specificReturn := fake.ruleExistsReturnsOnCall[len(fake.ruleExistsArgsForCall)]
	fake.ruleExistsArgsForCall = append(fake.ruleExistsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RuleExists", []interface{}{arg1})
	fake.ruleExistsMutex.Unlock()
	if fake.RuleExistsStub != nil {
		return fake.RuleExistsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.ruleExistsReturns.result1, fake.ruleExistsReturns.result2
}

func (fake *Firewall) RuleExistsCallCount() int {
	fake.ruleExistsMutex.RLock()
	defer fake.ruleExistsMutex.RUnlock()
	return len(fake.ruleExistsArgsForCall)
}

func (fake *Firewall) RuleExistsArgsForCall(i int) string {
	fake.ruleExistsMutex.RLock()
	defer fake.ruleExistsMutex.RUnlock()
	return fake.ruleExistsArgsForCall[i].arg1
}

func (fake *Firewall) RuleExistsReturns(result1 bool, result2 error) {
	fake.RuleExistsStub = nil
	fake.ruleExistsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *Firewall) RuleExistsReturnsOnCall(i int, result1 bool, result2 error) {
	fake.RuleExistsStub = nil
	if fake.ruleExistsReturnsOnCall == nil {
		fake.ruleExistsReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.ruleExistsReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *Firewall) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteRuleMutex.RLock()
	defer fake.deleteRuleMutex.RUnlock()
	fake.ruleExistsMutex.RLock()
	defer fake.ruleExistsMutex.RUnlock()
	return fake.invocations
}

func (fake *Firewall) recordInvocation(key string, args []interface{}) {
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

var _ firewallendpoint.Firewall = new(Firewall)

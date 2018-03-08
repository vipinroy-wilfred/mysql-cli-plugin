// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	"sync"

	"github.com/pivotal-cf/mysql-cli-plugin/service"
)

type FakeCfCommandRunner struct {
	CliCommandStub        func(...string) ([]string, error)
	cliCommandMutex       sync.RWMutex
	cliCommandArgsForCall []struct {
		arg1 []string
	}
	cliCommandReturns struct {
		result1 []string
		result2 error
	}
	cliCommandReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCfCommandRunner) CliCommand(arg1 ...string) ([]string, error) {
	fake.cliCommandMutex.Lock()
	ret, specificReturn := fake.cliCommandReturnsOnCall[len(fake.cliCommandArgsForCall)]
	fake.cliCommandArgsForCall = append(fake.cliCommandArgsForCall, struct {
		arg1 []string
	}{arg1})
	fake.recordInvocation("CliCommand", []interface{}{arg1})
	fake.cliCommandMutex.Unlock()
	if fake.CliCommandStub != nil {
		return fake.CliCommandStub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.cliCommandReturns.result1, fake.cliCommandReturns.result2
}

func (fake *FakeCfCommandRunner) CliCommandCallCount() int {
	fake.cliCommandMutex.RLock()
	defer fake.cliCommandMutex.RUnlock()
	return len(fake.cliCommandArgsForCall)
}

func (fake *FakeCfCommandRunner) CliCommandArgsForCall(i int) []string {
	fake.cliCommandMutex.RLock()
	defer fake.cliCommandMutex.RUnlock()
	return fake.cliCommandArgsForCall[i].arg1
}

func (fake *FakeCfCommandRunner) CliCommandReturns(result1 []string, result2 error) {
	fake.CliCommandStub = nil
	fake.cliCommandReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCfCommandRunner) CliCommandReturnsOnCall(i int, result1 []string, result2 error) {
	fake.CliCommandStub = nil
	if fake.cliCommandReturnsOnCall == nil {
		fake.cliCommandReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.cliCommandReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCfCommandRunner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cliCommandMutex.RLock()
	defer fake.cliCommandMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCfCommandRunner) recordInvocation(key string, args []interface{}) {
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

var _ service.CfCommandRunner = new(FakeCfCommandRunner)

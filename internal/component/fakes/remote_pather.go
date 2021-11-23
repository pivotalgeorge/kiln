// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/kiln/internal/component"
	"github.com/pivotal-cf/kiln/pkg/cargo"
)

type RemotePather struct {
	RemotePathStub        func(cargo.ComponentSpec) (string, error)
	remotePathMutex       sync.RWMutex
	remotePathArgsForCall []struct {
		arg1 cargo.ComponentSpec
	}
	remotePathReturns struct {
		result1 string
		result2 error
	}
	remotePathReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *RemotePather) RemotePath(arg1 cargo.ComponentSpec) (string, error) {
	fake.remotePathMutex.Lock()
	ret, specificReturn := fake.remotePathReturnsOnCall[len(fake.remotePathArgsForCall)]
	fake.remotePathArgsForCall = append(fake.remotePathArgsForCall, struct {
		arg1 cargo.ComponentSpec
	}{arg1})
	stub := fake.RemotePathStub
	fakeReturns := fake.remotePathReturns
	fake.recordInvocation("RemotePath", []interface{}{arg1})
	fake.remotePathMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *RemotePather) RemotePathCallCount() int {
	fake.remotePathMutex.RLock()
	defer fake.remotePathMutex.RUnlock()
	return len(fake.remotePathArgsForCall)
}

func (fake *RemotePather) RemotePathCalls(stub func(cargo.ComponentSpec) (string, error)) {
	fake.remotePathMutex.Lock()
	defer fake.remotePathMutex.Unlock()
	fake.RemotePathStub = stub
}

func (fake *RemotePather) RemotePathArgsForCall(i int) cargo.ComponentSpec {
	fake.remotePathMutex.RLock()
	defer fake.remotePathMutex.RUnlock()
	argsForCall := fake.remotePathArgsForCall[i]
	return argsForCall.arg1
}

func (fake *RemotePather) RemotePathReturns(result1 string, result2 error) {
	fake.remotePathMutex.Lock()
	defer fake.remotePathMutex.Unlock()
	fake.RemotePathStub = nil
	fake.remotePathReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *RemotePather) RemotePathReturnsOnCall(i int, result1 string, result2 error) {
	fake.remotePathMutex.Lock()
	defer fake.remotePathMutex.Unlock()
	fake.RemotePathStub = nil
	if fake.remotePathReturnsOnCall == nil {
		fake.remotePathReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.remotePathReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *RemotePather) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.remotePathMutex.RLock()
	defer fake.remotePathMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *RemotePather) recordInvocation(key string, args []interface{}) {
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

var _ component.RemotePather = new(RemotePather)
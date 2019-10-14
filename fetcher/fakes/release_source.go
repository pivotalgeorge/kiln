// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/kiln/fetcher"
	"github.com/pivotal-cf/kiln/internal/cargo"
)

type ReleaseSource struct {
	DownloadReleasesStub        func(string, fetcher.ReleaseSet, int) error
	downloadReleasesMutex       sync.RWMutex
	downloadReleasesArgsForCall []struct {
		arg1 string
		arg2 fetcher.ReleaseSet
		arg3 int
	}
	downloadReleasesReturns struct {
		result1 error
	}
	downloadReleasesReturnsOnCall map[int]struct {
		result1 error
	}
	GetMatchedReleasesStub        func(fetcher.ReleaseSet, cargo.Stemcell) (fetcher.ReleaseSet, error)
	getMatchedReleasesMutex       sync.RWMutex
	getMatchedReleasesArgsForCall []struct {
		arg1 fetcher.ReleaseSet
		arg2 cargo.Stemcell
	}
	getMatchedReleasesReturns struct {
		result1 fetcher.ReleaseSet
		result2 error
	}
	getMatchedReleasesReturnsOnCall map[int]struct {
		result1 fetcher.ReleaseSet
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ReleaseSource) DownloadReleases(arg1 string, arg2 fetcher.ReleaseSet, arg3 int) error {
	fake.downloadReleasesMutex.Lock()
	ret, specificReturn := fake.downloadReleasesReturnsOnCall[len(fake.downloadReleasesArgsForCall)]
	fake.downloadReleasesArgsForCall = append(fake.downloadReleasesArgsForCall, struct {
		arg1 string
		arg2 fetcher.ReleaseSet
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("DownloadReleases", []interface{}{arg1, arg2, arg3})
	fake.downloadReleasesMutex.Unlock()
	if fake.DownloadReleasesStub != nil {
		return fake.DownloadReleasesStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.downloadReleasesReturns
	return fakeReturns.result1
}

func (fake *ReleaseSource) DownloadReleasesCallCount() int {
	fake.downloadReleasesMutex.RLock()
	defer fake.downloadReleasesMutex.RUnlock()
	return len(fake.downloadReleasesArgsForCall)
}

func (fake *ReleaseSource) DownloadReleasesCalls(stub func(string, fetcher.ReleaseSet, int) error) {
	fake.downloadReleasesMutex.Lock()
	defer fake.downloadReleasesMutex.Unlock()
	fake.DownloadReleasesStub = stub
}

func (fake *ReleaseSource) DownloadReleasesArgsForCall(i int) (string, fetcher.ReleaseSet, int) {
	fake.downloadReleasesMutex.RLock()
	defer fake.downloadReleasesMutex.RUnlock()
	argsForCall := fake.downloadReleasesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ReleaseSource) DownloadReleasesReturns(result1 error) {
	fake.downloadReleasesMutex.Lock()
	defer fake.downloadReleasesMutex.Unlock()
	fake.DownloadReleasesStub = nil
	fake.downloadReleasesReturns = struct {
		result1 error
	}{result1}
}

func (fake *ReleaseSource) DownloadReleasesReturnsOnCall(i int, result1 error) {
	fake.downloadReleasesMutex.Lock()
	defer fake.downloadReleasesMutex.Unlock()
	fake.DownloadReleasesStub = nil
	if fake.downloadReleasesReturnsOnCall == nil {
		fake.downloadReleasesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadReleasesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ReleaseSource) GetMatchedReleases(arg1 fetcher.ReleaseSet, arg2 cargo.Stemcell) (fetcher.ReleaseSet, error) {
	fake.getMatchedReleasesMutex.Lock()
	ret, specificReturn := fake.getMatchedReleasesReturnsOnCall[len(fake.getMatchedReleasesArgsForCall)]
	fake.getMatchedReleasesArgsForCall = append(fake.getMatchedReleasesArgsForCall, struct {
		arg1 fetcher.ReleaseSet
		arg2 cargo.Stemcell
	}{arg1, arg2})
	fake.recordInvocation("GetMatchedReleases", []interface{}{arg1, arg2})
	fake.getMatchedReleasesMutex.Unlock()
	if fake.GetMatchedReleasesStub != nil {
		return fake.GetMatchedReleasesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getMatchedReleasesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ReleaseSource) GetMatchedReleasesCallCount() int {
	fake.getMatchedReleasesMutex.RLock()
	defer fake.getMatchedReleasesMutex.RUnlock()
	return len(fake.getMatchedReleasesArgsForCall)
}

func (fake *ReleaseSource) GetMatchedReleasesCalls(stub func(fetcher.ReleaseSet, cargo.Stemcell) (fetcher.ReleaseSet, error)) {
	fake.getMatchedReleasesMutex.Lock()
	defer fake.getMatchedReleasesMutex.Unlock()
	fake.GetMatchedReleasesStub = stub
}

func (fake *ReleaseSource) GetMatchedReleasesArgsForCall(i int) (fetcher.ReleaseSet, cargo.Stemcell) {
	fake.getMatchedReleasesMutex.RLock()
	defer fake.getMatchedReleasesMutex.RUnlock()
	argsForCall := fake.getMatchedReleasesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ReleaseSource) GetMatchedReleasesReturns(result1 fetcher.ReleaseSet, result2 error) {
	fake.getMatchedReleasesMutex.Lock()
	defer fake.getMatchedReleasesMutex.Unlock()
	fake.GetMatchedReleasesStub = nil
	fake.getMatchedReleasesReturns = struct {
		result1 fetcher.ReleaseSet
		result2 error
	}{result1, result2}
}

func (fake *ReleaseSource) GetMatchedReleasesReturnsOnCall(i int, result1 fetcher.ReleaseSet, result2 error) {
	fake.getMatchedReleasesMutex.Lock()
	defer fake.getMatchedReleasesMutex.Unlock()
	fake.GetMatchedReleasesStub = nil
	if fake.getMatchedReleasesReturnsOnCall == nil {
		fake.getMatchedReleasesReturnsOnCall = make(map[int]struct {
			result1 fetcher.ReleaseSet
			result2 error
		})
	}
	fake.getMatchedReleasesReturnsOnCall[i] = struct {
		result1 fetcher.ReleaseSet
		result2 error
	}{result1, result2}
}

func (fake *ReleaseSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadReleasesMutex.RLock()
	defer fake.downloadReleasesMutex.RUnlock()
	fake.getMatchedReleasesMutex.RLock()
	defer fake.getMatchedReleasesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ReleaseSource) recordInvocation(key string, args []interface{}) {
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

var _ fetcher.ReleaseSource = new(ReleaseSource)
// Code generated by counterfeiter. DO NOT EDIT.
package controllers

import (
	"context"
	"sync"

	"github.com/blang/semver"
	"github.com/wonderix/shalm/pkg/shalm"
)

type FakeK8s struct {
	ApplyStub        func(shalm.ObjectStream, *shalm.K8sOptions) error
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		arg1 shalm.ObjectStream
		arg2 *shalm.K8sOptions
	}
	applyReturns struct {
		result1 error
	}
	applyReturnsOnCall map[int]struct {
		result1 error
	}
	ConfigContentStub        func() *string
	configContentMutex       sync.RWMutex
	configContentArgsForCall []struct {
	}
	configContentReturns struct {
		result1 *string
	}
	configContentReturnsOnCall map[int]struct {
		result1 *string
	}
	DeleteStub        func(shalm.ObjectStream, *shalm.K8sOptions) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 shalm.ObjectStream
		arg2 *shalm.K8sOptions
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteObjectStub        func(string, string, *shalm.K8sOptions) error
	deleteObjectMutex       sync.RWMutex
	deleteObjectArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *shalm.K8sOptions
	}
	deleteObjectReturns struct {
		result1 error
	}
	deleteObjectReturnsOnCall map[int]struct {
		result1 error
	}
	ForConfigStub        func(string) (shalm.K8s, error)
	forConfigMutex       sync.RWMutex
	forConfigArgsForCall []struct {
		arg1 string
	}
	forConfigReturns struct {
		result1 shalm.K8s
		result2 error
	}
	forConfigReturnsOnCall map[int]struct {
		result1 shalm.K8s
		result2 error
	}
	ForSubChartStub        func(string, string, semver.Version) shalm.K8s
	forSubChartMutex       sync.RWMutex
	forSubChartArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 semver.Version
	}
	forSubChartReturns struct {
		result1 shalm.K8s
	}
	forSubChartReturnsOnCall map[int]struct {
		result1 shalm.K8s
	}
	GetStub        func(string, string, *shalm.K8sOptions) (*shalm.Object, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *shalm.K8sOptions
	}
	getReturns struct {
		result1 *shalm.Object
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *shalm.Object
		result2 error
	}
	HostStub        func() string
	hostMutex       sync.RWMutex
	hostArgsForCall []struct {
	}
	hostReturns struct {
		result1 string
	}
	hostReturnsOnCall map[int]struct {
		result1 string
	}
	InspectStub        func() string
	inspectMutex       sync.RWMutex
	inspectArgsForCall []struct {
	}
	inspectReturns struct {
		result1 string
	}
	inspectReturnsOnCall map[int]struct {
		result1 string
	}
	IsNotExistStub        func(error) bool
	isNotExistMutex       sync.RWMutex
	isNotExistArgsForCall []struct {
		arg1 error
	}
	isNotExistReturns struct {
		result1 bool
	}
	isNotExistReturnsOnCall map[int]struct {
		result1 bool
	}
	ProgressStub        func(int)
	progressMutex       sync.RWMutex
	progressArgsForCall []struct {
		arg1 int
	}
	RolloutStatusStub        func(string, string, *shalm.K8sOptions) error
	rolloutStatusMutex       sync.RWMutex
	rolloutStatusArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *shalm.K8sOptions
	}
	rolloutStatusReturns struct {
		result1 error
	}
	rolloutStatusReturnsOnCall map[int]struct {
		result1 error
	}
	SetToolStub        func(shalm.Tool)
	setToolMutex       sync.RWMutex
	setToolArgsForCall []struct {
		arg1 shalm.Tool
	}
	ToolStub        func() shalm.Tool
	toolMutex       sync.RWMutex
	toolArgsForCall []struct {
	}
	toolReturns struct {
		result1 shalm.Tool
	}
	toolReturnsOnCall map[int]struct {
		result1 shalm.Tool
	}
	WaitStub        func(string, string, string, *shalm.K8sOptions) error
	waitMutex       sync.RWMutex
	waitArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 *shalm.K8sOptions
	}
	waitReturns struct {
		result1 error
	}
	waitReturnsOnCall map[int]struct {
		result1 error
	}
	WatchStub        func(string, string, *shalm.K8sOptions) shalm.ObjectStream
	watchMutex       sync.RWMutex
	watchArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *shalm.K8sOptions
	}
	watchReturns struct {
		result1 shalm.ObjectStream
	}
	watchReturnsOnCall map[int]struct {
		result1 shalm.ObjectStream
	}
	WithContextStub        func(context.Context) shalm.K8s
	withContextMutex       sync.RWMutex
	withContextArgsForCall []struct {
		arg1 context.Context
	}
	withContextReturns struct {
		result1 shalm.K8s
	}
	withContextReturnsOnCall map[int]struct {
		result1 shalm.K8s
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeK8s) Apply(arg1 shalm.ObjectStream, arg2 *shalm.K8sOptions) error {
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		arg1 shalm.ObjectStream
		arg2 *shalm.K8sOptions
	}{arg1, arg2})
	fake.recordInvocation("Apply", []interface{}{arg1, arg2})
	fake.applyMutex.Unlock()
	if fake.ApplyStub != nil {
		return fake.ApplyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.applyReturns
	return fakeReturns.result1
}

func (fake *FakeK8s) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeK8s) ApplyCalls(stub func(shalm.ObjectStream, *shalm.K8sOptions) error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = stub
}

func (fake *FakeK8s) ApplyArgsForCall(i int) (shalm.ObjectStream, *shalm.K8sOptions) {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	argsForCall := fake.applyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeK8s) ApplyReturns(result1 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8s) ApplyReturnsOnCall(i int, result1 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8s) ConfigContent() *string {
	fake.configContentMutex.Lock()
	ret, specificReturn := fake.configContentReturnsOnCall[len(fake.configContentArgsForCall)]
	fake.configContentArgsForCall = append(fake.configContentArgsForCall, struct {
	}{})
	fake.recordInvocation("ConfigContent", []interface{}{})
	fake.configContentMutex.Unlock()
	if fake.ConfigContentStub != nil {
		return fake.ConfigContentStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.configContentReturns
	return fakeReturns.result1
}

func (fake *FakeK8s) ConfigContentCallCount() int {
	fake.configContentMutex.RLock()
	defer fake.configContentMutex.RUnlock()
	return len(fake.configContentArgsForCall)
}

func (fake *FakeK8s) ConfigContentCalls(stub func() *string) {
	fake.configContentMutex.Lock()
	defer fake.configContentMutex.Unlock()
	fake.ConfigContentStub = stub
}

func (fake *FakeK8s) ConfigContentReturns(result1 *string) {
	fake.configContentMutex.Lock()
	defer fake.configContentMutex.Unlock()
	fake.ConfigContentStub = nil
	fake.configContentReturns = struct {
		result1 *string
	}{result1}
}

func (fake *FakeK8s) ConfigContentReturnsOnCall(i int, result1 *string) {
	fake.configContentMutex.Lock()
	defer fake.configContentMutex.Unlock()
	fake.ConfigContentStub = nil
	if fake.configContentReturnsOnCall == nil {
		fake.configContentReturnsOnCall = make(map[int]struct {
			result1 *string
		})
	}
	fake.configContentReturnsOnCall[i] = struct {
		result1 *string
	}{result1}
}

func (fake *FakeK8s) Delete(arg1 shalm.ObjectStream, arg2 *shalm.K8sOptions) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 shalm.ObjectStream
		arg2 *shalm.K8sOptions
	}{arg1, arg2})
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakeK8s) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeK8s) DeleteCalls(stub func(shalm.ObjectStream, *shalm.K8sOptions) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeK8s) DeleteArgsForCall(i int) (shalm.ObjectStream, *shalm.K8sOptions) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeK8s) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8s) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8s) DeleteObject(arg1 string, arg2 string, arg3 *shalm.K8sOptions) error {
	fake.deleteObjectMutex.Lock()
	ret, specificReturn := fake.deleteObjectReturnsOnCall[len(fake.deleteObjectArgsForCall)]
	fake.deleteObjectArgsForCall = append(fake.deleteObjectArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *shalm.K8sOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("DeleteObject", []interface{}{arg1, arg2, arg3})
	fake.deleteObjectMutex.Unlock()
	if fake.DeleteObjectStub != nil {
		return fake.DeleteObjectStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteObjectReturns
	return fakeReturns.result1
}

func (fake *FakeK8s) DeleteObjectCallCount() int {
	fake.deleteObjectMutex.RLock()
	defer fake.deleteObjectMutex.RUnlock()
	return len(fake.deleteObjectArgsForCall)
}

func (fake *FakeK8s) DeleteObjectCalls(stub func(string, string, *shalm.K8sOptions) error) {
	fake.deleteObjectMutex.Lock()
	defer fake.deleteObjectMutex.Unlock()
	fake.DeleteObjectStub = stub
}

func (fake *FakeK8s) DeleteObjectArgsForCall(i int) (string, string, *shalm.K8sOptions) {
	fake.deleteObjectMutex.RLock()
	defer fake.deleteObjectMutex.RUnlock()
	argsForCall := fake.deleteObjectArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeK8s) DeleteObjectReturns(result1 error) {
	fake.deleteObjectMutex.Lock()
	defer fake.deleteObjectMutex.Unlock()
	fake.DeleteObjectStub = nil
	fake.deleteObjectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8s) DeleteObjectReturnsOnCall(i int, result1 error) {
	fake.deleteObjectMutex.Lock()
	defer fake.deleteObjectMutex.Unlock()
	fake.DeleteObjectStub = nil
	if fake.deleteObjectReturnsOnCall == nil {
		fake.deleteObjectReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteObjectReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8s) ForConfig(arg1 string) (shalm.K8s, error) {
	fake.forConfigMutex.Lock()
	ret, specificReturn := fake.forConfigReturnsOnCall[len(fake.forConfigArgsForCall)]
	fake.forConfigArgsForCall = append(fake.forConfigArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ForConfig", []interface{}{arg1})
	fake.forConfigMutex.Unlock()
	if fake.ForConfigStub != nil {
		return fake.ForConfigStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.forConfigReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeK8s) ForConfigCallCount() int {
	fake.forConfigMutex.RLock()
	defer fake.forConfigMutex.RUnlock()
	return len(fake.forConfigArgsForCall)
}

func (fake *FakeK8s) ForConfigCalls(stub func(string) (shalm.K8s, error)) {
	fake.forConfigMutex.Lock()
	defer fake.forConfigMutex.Unlock()
	fake.ForConfigStub = stub
}

func (fake *FakeK8s) ForConfigArgsForCall(i int) string {
	fake.forConfigMutex.RLock()
	defer fake.forConfigMutex.RUnlock()
	argsForCall := fake.forConfigArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeK8s) ForConfigReturns(result1 shalm.K8s, result2 error) {
	fake.forConfigMutex.Lock()
	defer fake.forConfigMutex.Unlock()
	fake.ForConfigStub = nil
	fake.forConfigReturns = struct {
		result1 shalm.K8s
		result2 error
	}{result1, result2}
}

func (fake *FakeK8s) ForConfigReturnsOnCall(i int, result1 shalm.K8s, result2 error) {
	fake.forConfigMutex.Lock()
	defer fake.forConfigMutex.Unlock()
	fake.ForConfigStub = nil
	if fake.forConfigReturnsOnCall == nil {
		fake.forConfigReturnsOnCall = make(map[int]struct {
			result1 shalm.K8s
			result2 error
		})
	}
	fake.forConfigReturnsOnCall[i] = struct {
		result1 shalm.K8s
		result2 error
	}{result1, result2}
}

func (fake *FakeK8s) ForSubChart(arg1 string, arg2 string, arg3 semver.Version) shalm.K8s {
	fake.forSubChartMutex.Lock()
	ret, specificReturn := fake.forSubChartReturnsOnCall[len(fake.forSubChartArgsForCall)]
	fake.forSubChartArgsForCall = append(fake.forSubChartArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 semver.Version
	}{arg1, arg2, arg3})
	fake.recordInvocation("ForSubChart", []interface{}{arg1, arg2, arg3})
	fake.forSubChartMutex.Unlock()
	if fake.ForSubChartStub != nil {
		return fake.ForSubChartStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.forSubChartReturns
	return fakeReturns.result1
}

func (fake *FakeK8s) ForSubChartCallCount() int {
	fake.forSubChartMutex.RLock()
	defer fake.forSubChartMutex.RUnlock()
	return len(fake.forSubChartArgsForCall)
}

func (fake *FakeK8s) ForSubChartCalls(stub func(string, string, semver.Version) shalm.K8s) {
	fake.forSubChartMutex.Lock()
	defer fake.forSubChartMutex.Unlock()
	fake.ForSubChartStub = stub
}

func (fake *FakeK8s) ForSubChartArgsForCall(i int) (string, string, semver.Version) {
	fake.forSubChartMutex.RLock()
	defer fake.forSubChartMutex.RUnlock()
	argsForCall := fake.forSubChartArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeK8s) ForSubChartReturns(result1 shalm.K8s) {
	fake.forSubChartMutex.Lock()
	defer fake.forSubChartMutex.Unlock()
	fake.ForSubChartStub = nil
	fake.forSubChartReturns = struct {
		result1 shalm.K8s
	}{result1}
}

func (fake *FakeK8s) ForSubChartReturnsOnCall(i int, result1 shalm.K8s) {
	fake.forSubChartMutex.Lock()
	defer fake.forSubChartMutex.Unlock()
	fake.ForSubChartStub = nil
	if fake.forSubChartReturnsOnCall == nil {
		fake.forSubChartReturnsOnCall = make(map[int]struct {
			result1 shalm.K8s
		})
	}
	fake.forSubChartReturnsOnCall[i] = struct {
		result1 shalm.K8s
	}{result1}
}

func (fake *FakeK8s) Get(arg1 string, arg2 string, arg3 *shalm.K8sOptions) (*shalm.Object, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *shalm.K8sOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeK8s) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeK8s) GetCalls(stub func(string, string, *shalm.K8sOptions) (*shalm.Object, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeK8s) GetArgsForCall(i int) (string, string, *shalm.K8sOptions) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeK8s) GetReturns(result1 *shalm.Object, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *shalm.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeK8s) GetReturnsOnCall(i int, result1 *shalm.Object, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *shalm.Object
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *shalm.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeK8s) Host() string {
	fake.hostMutex.Lock()
	ret, specificReturn := fake.hostReturnsOnCall[len(fake.hostArgsForCall)]
	fake.hostArgsForCall = append(fake.hostArgsForCall, struct {
	}{})
	fake.recordInvocation("Host", []interface{}{})
	fake.hostMutex.Unlock()
	if fake.HostStub != nil {
		return fake.HostStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.hostReturns
	return fakeReturns.result1
}

func (fake *FakeK8s) HostCallCount() int {
	fake.hostMutex.RLock()
	defer fake.hostMutex.RUnlock()
	return len(fake.hostArgsForCall)
}

func (fake *FakeK8s) HostCalls(stub func() string) {
	fake.hostMutex.Lock()
	defer fake.hostMutex.Unlock()
	fake.HostStub = stub
}

func (fake *FakeK8s) HostReturns(result1 string) {
	fake.hostMutex.Lock()
	defer fake.hostMutex.Unlock()
	fake.HostStub = nil
	fake.hostReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeK8s) HostReturnsOnCall(i int, result1 string) {
	fake.hostMutex.Lock()
	defer fake.hostMutex.Unlock()
	fake.HostStub = nil
	if fake.hostReturnsOnCall == nil {
		fake.hostReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.hostReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeK8s) Inspect() string {
	fake.inspectMutex.Lock()
	ret, specificReturn := fake.inspectReturnsOnCall[len(fake.inspectArgsForCall)]
	fake.inspectArgsForCall = append(fake.inspectArgsForCall, struct {
	}{})
	fake.recordInvocation("Inspect", []interface{}{})
	fake.inspectMutex.Unlock()
	if fake.InspectStub != nil {
		return fake.InspectStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.inspectReturns
	return fakeReturns.result1
}

func (fake *FakeK8s) InspectCallCount() int {
	fake.inspectMutex.RLock()
	defer fake.inspectMutex.RUnlock()
	return len(fake.inspectArgsForCall)
}

func (fake *FakeK8s) InspectCalls(stub func() string) {
	fake.inspectMutex.Lock()
	defer fake.inspectMutex.Unlock()
	fake.InspectStub = stub
}

func (fake *FakeK8s) InspectReturns(result1 string) {
	fake.inspectMutex.Lock()
	defer fake.inspectMutex.Unlock()
	fake.InspectStub = nil
	fake.inspectReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeK8s) InspectReturnsOnCall(i int, result1 string) {
	fake.inspectMutex.Lock()
	defer fake.inspectMutex.Unlock()
	fake.InspectStub = nil
	if fake.inspectReturnsOnCall == nil {
		fake.inspectReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.inspectReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeK8s) IsNotExist(arg1 error) bool {
	fake.isNotExistMutex.Lock()
	ret, specificReturn := fake.isNotExistReturnsOnCall[len(fake.isNotExistArgsForCall)]
	fake.isNotExistArgsForCall = append(fake.isNotExistArgsForCall, struct {
		arg1 error
	}{arg1})
	fake.recordInvocation("IsNotExist", []interface{}{arg1})
	fake.isNotExistMutex.Unlock()
	if fake.IsNotExistStub != nil {
		return fake.IsNotExistStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.isNotExistReturns
	return fakeReturns.result1
}

func (fake *FakeK8s) IsNotExistCallCount() int {
	fake.isNotExistMutex.RLock()
	defer fake.isNotExistMutex.RUnlock()
	return len(fake.isNotExistArgsForCall)
}

func (fake *FakeK8s) IsNotExistCalls(stub func(error) bool) {
	fake.isNotExistMutex.Lock()
	defer fake.isNotExistMutex.Unlock()
	fake.IsNotExistStub = stub
}

func (fake *FakeK8s) IsNotExistArgsForCall(i int) error {
	fake.isNotExistMutex.RLock()
	defer fake.isNotExistMutex.RUnlock()
	argsForCall := fake.isNotExistArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeK8s) IsNotExistReturns(result1 bool) {
	fake.isNotExistMutex.Lock()
	defer fake.isNotExistMutex.Unlock()
	fake.IsNotExistStub = nil
	fake.isNotExistReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeK8s) IsNotExistReturnsOnCall(i int, result1 bool) {
	fake.isNotExistMutex.Lock()
	defer fake.isNotExistMutex.Unlock()
	fake.IsNotExistStub = nil
	if fake.isNotExistReturnsOnCall == nil {
		fake.isNotExistReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isNotExistReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeK8s) Progress(arg1 int) {
	fake.progressMutex.Lock()
	fake.progressArgsForCall = append(fake.progressArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("Progress", []interface{}{arg1})
	fake.progressMutex.Unlock()
	if fake.ProgressStub != nil {
		fake.ProgressStub(arg1)
	}
}

func (fake *FakeK8s) ProgressCallCount() int {
	fake.progressMutex.RLock()
	defer fake.progressMutex.RUnlock()
	return len(fake.progressArgsForCall)
}

func (fake *FakeK8s) ProgressCalls(stub func(int)) {
	fake.progressMutex.Lock()
	defer fake.progressMutex.Unlock()
	fake.ProgressStub = stub
}

func (fake *FakeK8s) ProgressArgsForCall(i int) int {
	fake.progressMutex.RLock()
	defer fake.progressMutex.RUnlock()
	argsForCall := fake.progressArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeK8s) RolloutStatus(arg1 string, arg2 string, arg3 *shalm.K8sOptions) error {
	fake.rolloutStatusMutex.Lock()
	ret, specificReturn := fake.rolloutStatusReturnsOnCall[len(fake.rolloutStatusArgsForCall)]
	fake.rolloutStatusArgsForCall = append(fake.rolloutStatusArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *shalm.K8sOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("RolloutStatus", []interface{}{arg1, arg2, arg3})
	fake.rolloutStatusMutex.Unlock()
	if fake.RolloutStatusStub != nil {
		return fake.RolloutStatusStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.rolloutStatusReturns
	return fakeReturns.result1
}

func (fake *FakeK8s) RolloutStatusCallCount() int {
	fake.rolloutStatusMutex.RLock()
	defer fake.rolloutStatusMutex.RUnlock()
	return len(fake.rolloutStatusArgsForCall)
}

func (fake *FakeK8s) RolloutStatusCalls(stub func(string, string, *shalm.K8sOptions) error) {
	fake.rolloutStatusMutex.Lock()
	defer fake.rolloutStatusMutex.Unlock()
	fake.RolloutStatusStub = stub
}

func (fake *FakeK8s) RolloutStatusArgsForCall(i int) (string, string, *shalm.K8sOptions) {
	fake.rolloutStatusMutex.RLock()
	defer fake.rolloutStatusMutex.RUnlock()
	argsForCall := fake.rolloutStatusArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeK8s) RolloutStatusReturns(result1 error) {
	fake.rolloutStatusMutex.Lock()
	defer fake.rolloutStatusMutex.Unlock()
	fake.RolloutStatusStub = nil
	fake.rolloutStatusReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8s) RolloutStatusReturnsOnCall(i int, result1 error) {
	fake.rolloutStatusMutex.Lock()
	defer fake.rolloutStatusMutex.Unlock()
	fake.RolloutStatusStub = nil
	if fake.rolloutStatusReturnsOnCall == nil {
		fake.rolloutStatusReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.rolloutStatusReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8s) SetTool(arg1 shalm.Tool) {
	fake.setToolMutex.Lock()
	fake.setToolArgsForCall = append(fake.setToolArgsForCall, struct {
		arg1 shalm.Tool
	}{arg1})
	fake.recordInvocation("SetTool", []interface{}{arg1})
	fake.setToolMutex.Unlock()
	if fake.SetToolStub != nil {
		fake.SetToolStub(arg1)
	}
}

func (fake *FakeK8s) SetToolCallCount() int {
	fake.setToolMutex.RLock()
	defer fake.setToolMutex.RUnlock()
	return len(fake.setToolArgsForCall)
}

func (fake *FakeK8s) SetToolCalls(stub func(shalm.Tool)) {
	fake.setToolMutex.Lock()
	defer fake.setToolMutex.Unlock()
	fake.SetToolStub = stub
}

func (fake *FakeK8s) SetToolArgsForCall(i int) shalm.Tool {
	fake.setToolMutex.RLock()
	defer fake.setToolMutex.RUnlock()
	argsForCall := fake.setToolArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeK8s) Tool() shalm.Tool {
	fake.toolMutex.Lock()
	ret, specificReturn := fake.toolReturnsOnCall[len(fake.toolArgsForCall)]
	fake.toolArgsForCall = append(fake.toolArgsForCall, struct {
	}{})
	fake.recordInvocation("Tool", []interface{}{})
	fake.toolMutex.Unlock()
	if fake.ToolStub != nil {
		return fake.ToolStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.toolReturns
	return fakeReturns.result1
}

func (fake *FakeK8s) ToolCallCount() int {
	fake.toolMutex.RLock()
	defer fake.toolMutex.RUnlock()
	return len(fake.toolArgsForCall)
}

func (fake *FakeK8s) ToolCalls(stub func() shalm.Tool) {
	fake.toolMutex.Lock()
	defer fake.toolMutex.Unlock()
	fake.ToolStub = stub
}

func (fake *FakeK8s) ToolReturns(result1 shalm.Tool) {
	fake.toolMutex.Lock()
	defer fake.toolMutex.Unlock()
	fake.ToolStub = nil
	fake.toolReturns = struct {
		result1 shalm.Tool
	}{result1}
}

func (fake *FakeK8s) ToolReturnsOnCall(i int, result1 shalm.Tool) {
	fake.toolMutex.Lock()
	defer fake.toolMutex.Unlock()
	fake.ToolStub = nil
	if fake.toolReturnsOnCall == nil {
		fake.toolReturnsOnCall = make(map[int]struct {
			result1 shalm.Tool
		})
	}
	fake.toolReturnsOnCall[i] = struct {
		result1 shalm.Tool
	}{result1}
}

func (fake *FakeK8s) Wait(arg1 string, arg2 string, arg3 string, arg4 *shalm.K8sOptions) error {
	fake.waitMutex.Lock()
	ret, specificReturn := fake.waitReturnsOnCall[len(fake.waitArgsForCall)]
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 *shalm.K8sOptions
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Wait", []interface{}{arg1, arg2, arg3, arg4})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.waitReturns
	return fakeReturns.result1
}

func (fake *FakeK8s) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeK8s) WaitCalls(stub func(string, string, string, *shalm.K8sOptions) error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = stub
}

func (fake *FakeK8s) WaitArgsForCall(i int) (string, string, string, *shalm.K8sOptions) {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	argsForCall := fake.waitArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeK8s) WaitReturns(result1 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8s) WaitReturnsOnCall(i int, result1 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	if fake.waitReturnsOnCall == nil {
		fake.waitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.waitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8s) Watch(arg1 string, arg2 string, arg3 *shalm.K8sOptions) shalm.ObjectStream {
	fake.watchMutex.Lock()
	ret, specificReturn := fake.watchReturnsOnCall[len(fake.watchArgsForCall)]
	fake.watchArgsForCall = append(fake.watchArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *shalm.K8sOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("Watch", []interface{}{arg1, arg2, arg3})
	fake.watchMutex.Unlock()
	if fake.WatchStub != nil {
		return fake.WatchStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.watchReturns
	return fakeReturns.result1
}

func (fake *FakeK8s) WatchCallCount() int {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	return len(fake.watchArgsForCall)
}

func (fake *FakeK8s) WatchCalls(stub func(string, string, *shalm.K8sOptions) shalm.ObjectStream) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = stub
}

func (fake *FakeK8s) WatchArgsForCall(i int) (string, string, *shalm.K8sOptions) {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	argsForCall := fake.watchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeK8s) WatchReturns(result1 shalm.ObjectStream) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	fake.watchReturns = struct {
		result1 shalm.ObjectStream
	}{result1}
}

func (fake *FakeK8s) WatchReturnsOnCall(i int, result1 shalm.ObjectStream) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	if fake.watchReturnsOnCall == nil {
		fake.watchReturnsOnCall = make(map[int]struct {
			result1 shalm.ObjectStream
		})
	}
	fake.watchReturnsOnCall[i] = struct {
		result1 shalm.ObjectStream
	}{result1}
}

func (fake *FakeK8s) WithContext(arg1 context.Context) shalm.K8s {
	fake.withContextMutex.Lock()
	ret, specificReturn := fake.withContextReturnsOnCall[len(fake.withContextArgsForCall)]
	fake.withContextArgsForCall = append(fake.withContextArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("WithContext", []interface{}{arg1})
	fake.withContextMutex.Unlock()
	if fake.WithContextStub != nil {
		return fake.WithContextStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.withContextReturns
	return fakeReturns.result1
}

func (fake *FakeK8s) WithContextCallCount() int {
	fake.withContextMutex.RLock()
	defer fake.withContextMutex.RUnlock()
	return len(fake.withContextArgsForCall)
}

func (fake *FakeK8s) WithContextCalls(stub func(context.Context) shalm.K8s) {
	fake.withContextMutex.Lock()
	defer fake.withContextMutex.Unlock()
	fake.WithContextStub = stub
}

func (fake *FakeK8s) WithContextArgsForCall(i int) context.Context {
	fake.withContextMutex.RLock()
	defer fake.withContextMutex.RUnlock()
	argsForCall := fake.withContextArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeK8s) WithContextReturns(result1 shalm.K8s) {
	fake.withContextMutex.Lock()
	defer fake.withContextMutex.Unlock()
	fake.WithContextStub = nil
	fake.withContextReturns = struct {
		result1 shalm.K8s
	}{result1}
}

func (fake *FakeK8s) WithContextReturnsOnCall(i int, result1 shalm.K8s) {
	fake.withContextMutex.Lock()
	defer fake.withContextMutex.Unlock()
	fake.WithContextStub = nil
	if fake.withContextReturnsOnCall == nil {
		fake.withContextReturnsOnCall = make(map[int]struct {
			result1 shalm.K8s
		})
	}
	fake.withContextReturnsOnCall[i] = struct {
		result1 shalm.K8s
	}{result1}
}

func (fake *FakeK8s) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	fake.configContentMutex.RLock()
	defer fake.configContentMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deleteObjectMutex.RLock()
	defer fake.deleteObjectMutex.RUnlock()
	fake.forConfigMutex.RLock()
	defer fake.forConfigMutex.RUnlock()
	fake.forSubChartMutex.RLock()
	defer fake.forSubChartMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.hostMutex.RLock()
	defer fake.hostMutex.RUnlock()
	fake.inspectMutex.RLock()
	defer fake.inspectMutex.RUnlock()
	fake.isNotExistMutex.RLock()
	defer fake.isNotExistMutex.RUnlock()
	fake.progressMutex.RLock()
	defer fake.progressMutex.RUnlock()
	fake.rolloutStatusMutex.RLock()
	defer fake.rolloutStatusMutex.RUnlock()
	fake.setToolMutex.RLock()
	defer fake.setToolMutex.RUnlock()
	fake.toolMutex.RLock()
	defer fake.toolMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	fake.withContextMutex.RLock()
	defer fake.withContextMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeK8s) recordInvocation(key string, args []interface{}) {
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

var _ shalm.K8s = new(FakeK8s)

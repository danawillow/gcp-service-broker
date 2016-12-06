// This file was generated by counterfeiter
package modelsfakes

import (
	"gcp-service-broker/brokerapi/brokers/models"
	"sync"
)

type FakeServiceBrokerHelper struct {
	ProvisionStub        func(instanceId string, details models.ProvisionDetails, plan models.PlanDetails) (models.ServiceInstanceDetails, error)
	provisionMutex       sync.RWMutex
	provisionArgsForCall []struct {
		instanceId string
		details    models.ProvisionDetails
		plan       models.PlanDetails
	}
	provisionReturns struct {
		result1 models.ServiceInstanceDetails
		result2 error
	}
	BindStub        func(instanceID, bindingID string, details models.BindDetails) (models.ServiceBindingCredentials, error)
	bindMutex       sync.RWMutex
	bindArgsForCall []struct {
		instanceID string
		bindingID  string
		details    models.BindDetails
	}
	bindReturns struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}
	BuildInstanceCredentialsStub        func(bindDetails map[string]string, instanceDetails map[string]string) map[string]string
	buildInstanceCredentialsMutex       sync.RWMutex
	buildInstanceCredentialsArgsForCall []struct {
		bindDetails     map[string]string
		instanceDetails map[string]string
	}
	buildInstanceCredentialsReturns struct {
		result1 map[string]string
	}
	UnbindStub        func(details models.ServiceBindingCredentials) error
	unbindMutex       sync.RWMutex
	unbindArgsForCall []struct {
		details models.ServiceBindingCredentials
	}
	unbindReturns struct {
		result1 error
	}
	DeprovisionStub        func(instanceID string, details models.DeprovisionDetails) error
	deprovisionMutex       sync.RWMutex
	deprovisionArgsForCall []struct {
		instanceID string
		details    models.DeprovisionDetails
	}
	deprovisionReturns struct {
		result1 error
	}
	PollInstanceStub        func(instanceID string) (bool, error)
	pollInstanceMutex       sync.RWMutex
	pollInstanceArgsForCall []struct {
		instanceID string
	}
	pollInstanceReturns struct {
		result1 bool
		result2 error
	}
	AsyncStub        func() bool
	asyncMutex       sync.RWMutex
	asyncArgsForCall []struct{}
	asyncReturns     struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceBrokerHelper) Provision(instanceId string, details models.ProvisionDetails, plan models.PlanDetails) (models.ServiceInstanceDetails, error) {
	fake.provisionMutex.Lock()
	fake.provisionArgsForCall = append(fake.provisionArgsForCall, struct {
		instanceId string
		details    models.ProvisionDetails
		plan       models.PlanDetails
	}{instanceId, details, plan})
	fake.recordInvocation("Provision", []interface{}{instanceId, details, plan})
	fake.provisionMutex.Unlock()
	if fake.ProvisionStub != nil {
		return fake.ProvisionStub(instanceId, details, plan)
	} else {
		return fake.provisionReturns.result1, fake.provisionReturns.result2
	}
}

func (fake *FakeServiceBrokerHelper) ProvisionCallCount() int {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return len(fake.provisionArgsForCall)
}

func (fake *FakeServiceBrokerHelper) ProvisionArgsForCall(i int) (string, models.ProvisionDetails, models.PlanDetails) {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return fake.provisionArgsForCall[i].instanceId, fake.provisionArgsForCall[i].details, fake.provisionArgsForCall[i].plan
}

func (fake *FakeServiceBrokerHelper) ProvisionReturns(result1 models.ServiceInstanceDetails, result2 error) {
	fake.ProvisionStub = nil
	fake.provisionReturns = struct {
		result1 models.ServiceInstanceDetails
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) Bind(instanceID string, bindingID string, details models.BindDetails) (models.ServiceBindingCredentials, error) {
	fake.bindMutex.Lock()
	fake.bindArgsForCall = append(fake.bindArgsForCall, struct {
		instanceID string
		bindingID  string
		details    models.BindDetails
	}{instanceID, bindingID, details})
	fake.recordInvocation("Bind", []interface{}{instanceID, bindingID, details})
	fake.bindMutex.Unlock()
	if fake.BindStub != nil {
		return fake.BindStub(instanceID, bindingID, details)
	} else {
		return fake.bindReturns.result1, fake.bindReturns.result2
	}
}

func (fake *FakeServiceBrokerHelper) BindCallCount() int {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return len(fake.bindArgsForCall)
}

func (fake *FakeServiceBrokerHelper) BindArgsForCall(i int) (string, string, models.BindDetails) {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return fake.bindArgsForCall[i].instanceID, fake.bindArgsForCall[i].bindingID, fake.bindArgsForCall[i].details
}

func (fake *FakeServiceBrokerHelper) BindReturns(result1 models.ServiceBindingCredentials, result2 error) {
	fake.BindStub = nil
	fake.bindReturns = struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) BuildInstanceCredentials(bindDetails map[string]string, instanceDetails map[string]string) map[string]string {
	fake.buildInstanceCredentialsMutex.Lock()
	fake.buildInstanceCredentialsArgsForCall = append(fake.buildInstanceCredentialsArgsForCall, struct {
		bindDetails     map[string]string
		instanceDetails map[string]string
	}{bindDetails, instanceDetails})
	fake.recordInvocation("BuildInstanceCredentials", []interface{}{bindDetails, instanceDetails})
	fake.buildInstanceCredentialsMutex.Unlock()
	if fake.BuildInstanceCredentialsStub != nil {
		return fake.BuildInstanceCredentialsStub(bindDetails, instanceDetails)
	} else {
		return fake.buildInstanceCredentialsReturns.result1
	}
}

func (fake *FakeServiceBrokerHelper) BuildInstanceCredentialsCallCount() int {
	fake.buildInstanceCredentialsMutex.RLock()
	defer fake.buildInstanceCredentialsMutex.RUnlock()
	return len(fake.buildInstanceCredentialsArgsForCall)
}

func (fake *FakeServiceBrokerHelper) BuildInstanceCredentialsArgsForCall(i int) (map[string]string, map[string]string) {
	fake.buildInstanceCredentialsMutex.RLock()
	defer fake.buildInstanceCredentialsMutex.RUnlock()
	return fake.buildInstanceCredentialsArgsForCall[i].bindDetails, fake.buildInstanceCredentialsArgsForCall[i].instanceDetails
}

func (fake *FakeServiceBrokerHelper) BuildInstanceCredentialsReturns(result1 map[string]string) {
	fake.BuildInstanceCredentialsStub = nil
	fake.buildInstanceCredentialsReturns = struct {
		result1 map[string]string
	}{result1}
}

func (fake *FakeServiceBrokerHelper) Unbind(details models.ServiceBindingCredentials) error {
	fake.unbindMutex.Lock()
	fake.unbindArgsForCall = append(fake.unbindArgsForCall, struct {
		details models.ServiceBindingCredentials
	}{details})
	fake.recordInvocation("Unbind", []interface{}{details})
	fake.unbindMutex.Unlock()
	if fake.UnbindStub != nil {
		return fake.UnbindStub(details)
	} else {
		return fake.unbindReturns.result1
	}
}

func (fake *FakeServiceBrokerHelper) UnbindCallCount() int {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return len(fake.unbindArgsForCall)
}

func (fake *FakeServiceBrokerHelper) UnbindArgsForCall(i int) models.ServiceBindingCredentials {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return fake.unbindArgsForCall[i].details
}

func (fake *FakeServiceBrokerHelper) UnbindReturns(result1 error) {
	fake.UnbindStub = nil
	fake.unbindReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerHelper) Deprovision(instanceID string, details models.DeprovisionDetails) error {
	fake.deprovisionMutex.Lock()
	fake.deprovisionArgsForCall = append(fake.deprovisionArgsForCall, struct {
		instanceID string
		details    models.DeprovisionDetails
	}{instanceID, details})
	fake.recordInvocation("Deprovision", []interface{}{instanceID, details})
	fake.deprovisionMutex.Unlock()
	if fake.DeprovisionStub != nil {
		return fake.DeprovisionStub(instanceID, details)
	} else {
		return fake.deprovisionReturns.result1
	}
}

func (fake *FakeServiceBrokerHelper) DeprovisionCallCount() int {
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	return len(fake.deprovisionArgsForCall)
}

func (fake *FakeServiceBrokerHelper) DeprovisionArgsForCall(i int) (string, models.DeprovisionDetails) {
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	return fake.deprovisionArgsForCall[i].instanceID, fake.deprovisionArgsForCall[i].details
}

func (fake *FakeServiceBrokerHelper) DeprovisionReturns(result1 error) {
	fake.DeprovisionStub = nil
	fake.deprovisionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerHelper) PollInstance(instanceID string) (bool, error) {
	fake.pollInstanceMutex.Lock()
	fake.pollInstanceArgsForCall = append(fake.pollInstanceArgsForCall, struct {
		instanceID string
	}{instanceID})
	fake.recordInvocation("PollInstance", []interface{}{instanceID})
	fake.pollInstanceMutex.Unlock()
	if fake.PollInstanceStub != nil {
		return fake.PollInstanceStub(instanceID)
	} else {
		return fake.pollInstanceReturns.result1, fake.pollInstanceReturns.result2
	}
}

func (fake *FakeServiceBrokerHelper) PollInstanceCallCount() int {
	fake.pollInstanceMutex.RLock()
	defer fake.pollInstanceMutex.RUnlock()
	return len(fake.pollInstanceArgsForCall)
}

func (fake *FakeServiceBrokerHelper) PollInstanceArgsForCall(i int) string {
	fake.pollInstanceMutex.RLock()
	defer fake.pollInstanceMutex.RUnlock()
	return fake.pollInstanceArgsForCall[i].instanceID
}

func (fake *FakeServiceBrokerHelper) PollInstanceReturns(result1 bool, result2 error) {
	fake.PollInstanceStub = nil
	fake.pollInstanceReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) Async() bool {
	fake.asyncMutex.Lock()
	fake.asyncArgsForCall = append(fake.asyncArgsForCall, struct{}{})
	fake.recordInvocation("Async", []interface{}{})
	fake.asyncMutex.Unlock()
	if fake.AsyncStub != nil {
		return fake.AsyncStub()
	} else {
		return fake.asyncReturns.result1
	}
}

func (fake *FakeServiceBrokerHelper) AsyncCallCount() int {
	fake.asyncMutex.RLock()
	defer fake.asyncMutex.RUnlock()
	return len(fake.asyncArgsForCall)
}

func (fake *FakeServiceBrokerHelper) AsyncReturns(result1 bool) {
	fake.AsyncStub = nil
	fake.asyncReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeServiceBrokerHelper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	fake.buildInstanceCredentialsMutex.RLock()
	defer fake.buildInstanceCredentialsMutex.RUnlock()
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	fake.pollInstanceMutex.RLock()
	defer fake.pollInstanceMutex.RUnlock()
	fake.asyncMutex.RLock()
	defer fake.asyncMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeServiceBrokerHelper) recordInvocation(key string, args []interface{}) {
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

var _ models.ServiceBrokerHelper = new(FakeServiceBrokerHelper)

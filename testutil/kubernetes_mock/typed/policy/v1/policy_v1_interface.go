// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	testing "testing"

	v1 "k8s.io/client-go/kubernetes/typed/policy/v1"
)

// PolicyV1Interface is an autogenerated mock type for the PolicyV1Interface type
type PolicyV1Interface struct {
	mock.Mock
}

// PodDisruptionBudgets provides a mock function with given fields: namespace
func (_m *PolicyV1Interface) PodDisruptionBudgets(namespace string) v1.PodDisruptionBudgetInterface {
	ret := _m.Called(namespace)

	var r0 v1.PodDisruptionBudgetInterface
	if rf, ok := ret.Get(0).(func(string) v1.PodDisruptionBudgetInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.PodDisruptionBudgetInterface)
		}
	}

	return r0
}

// RESTClient provides a mock function with given fields:
func (_m *PolicyV1Interface) RESTClient() rest.Interface {
	ret := _m.Called()

	var r0 rest.Interface
	if rf, ok := ret.Get(0).(func() rest.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.Interface)
		}
	}

	return r0
}

// NewPolicyV1Interface creates a new instance of PolicyV1Interface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewPolicyV1Interface(t testing.TB) *PolicyV1Interface {
	mock := &PolicyV1Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

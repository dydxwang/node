// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// ClusterRoleExpansion is an autogenerated mock type for the ClusterRoleExpansion type
type ClusterRoleExpansion struct {
	mock.Mock
}

// NewClusterRoleExpansion creates a new instance of ClusterRoleExpansion. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewClusterRoleExpansion(t testing.TB) *ClusterRoleExpansion {
	mock := &ClusterRoleExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

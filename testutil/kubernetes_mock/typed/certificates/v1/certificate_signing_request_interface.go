// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	certificatesv1 "k8s.io/api/certificates/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	testing "testing"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/certificates/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// CertificateSigningRequestInterface is an autogenerated mock type for the CertificateSigningRequestInterface type
type CertificateSigningRequestInterface struct {
	mock.Mock
}

// Apply provides a mock function with given fields: ctx, certificateSigningRequest, opts
func (_m *CertificateSigningRequestInterface) Apply(ctx context.Context, certificateSigningRequest *v1.CertificateSigningRequestApplyConfiguration, opts metav1.ApplyOptions) (*certificatesv1.CertificateSigningRequest, error) {
	ret := _m.Called(ctx, certificateSigningRequest, opts)

	var r0 *certificatesv1.CertificateSigningRequest
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CertificateSigningRequestApplyConfiguration, metav1.ApplyOptions) *certificatesv1.CertificateSigningRequest); ok {
		r0 = rf(ctx, certificateSigningRequest, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*certificatesv1.CertificateSigningRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.CertificateSigningRequestApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, certificateSigningRequest, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApplyStatus provides a mock function with given fields: ctx, certificateSigningRequest, opts
func (_m *CertificateSigningRequestInterface) ApplyStatus(ctx context.Context, certificateSigningRequest *v1.CertificateSigningRequestApplyConfiguration, opts metav1.ApplyOptions) (*certificatesv1.CertificateSigningRequest, error) {
	ret := _m.Called(ctx, certificateSigningRequest, opts)

	var r0 *certificatesv1.CertificateSigningRequest
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CertificateSigningRequestApplyConfiguration, metav1.ApplyOptions) *certificatesv1.CertificateSigningRequest); ok {
		r0 = rf(ctx, certificateSigningRequest, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*certificatesv1.CertificateSigningRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.CertificateSigningRequestApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, certificateSigningRequest, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, certificateSigningRequest, opts
func (_m *CertificateSigningRequestInterface) Create(ctx context.Context, certificateSigningRequest *certificatesv1.CertificateSigningRequest, opts metav1.CreateOptions) (*certificatesv1.CertificateSigningRequest, error) {
	ret := _m.Called(ctx, certificateSigningRequest, opts)

	var r0 *certificatesv1.CertificateSigningRequest
	if rf, ok := ret.Get(0).(func(context.Context, *certificatesv1.CertificateSigningRequest, metav1.CreateOptions) *certificatesv1.CertificateSigningRequest); ok {
		r0 = rf(ctx, certificateSigningRequest, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*certificatesv1.CertificateSigningRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *certificatesv1.CertificateSigningRequest, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, certificateSigningRequest, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *CertificateSigningRequestInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *CertificateSigningRequestInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *CertificateSigningRequestInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*certificatesv1.CertificateSigningRequest, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *certificatesv1.CertificateSigningRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *certificatesv1.CertificateSigningRequest); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*certificatesv1.CertificateSigningRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, opts
func (_m *CertificateSigningRequestInterface) List(ctx context.Context, opts metav1.ListOptions) (*certificatesv1.CertificateSigningRequestList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *certificatesv1.CertificateSigningRequestList
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *certificatesv1.CertificateSigningRequestList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*certificatesv1.CertificateSigningRequestList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *CertificateSigningRequestInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*certificatesv1.CertificateSigningRequest, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *certificatesv1.CertificateSigningRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *certificatesv1.CertificateSigningRequest); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*certificatesv1.CertificateSigningRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, certificateSigningRequest, opts
func (_m *CertificateSigningRequestInterface) Update(ctx context.Context, certificateSigningRequest *certificatesv1.CertificateSigningRequest, opts metav1.UpdateOptions) (*certificatesv1.CertificateSigningRequest, error) {
	ret := _m.Called(ctx, certificateSigningRequest, opts)

	var r0 *certificatesv1.CertificateSigningRequest
	if rf, ok := ret.Get(0).(func(context.Context, *certificatesv1.CertificateSigningRequest, metav1.UpdateOptions) *certificatesv1.CertificateSigningRequest); ok {
		r0 = rf(ctx, certificateSigningRequest, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*certificatesv1.CertificateSigningRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *certificatesv1.CertificateSigningRequest, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, certificateSigningRequest, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateApproval provides a mock function with given fields: ctx, certificateSigningRequestName, certificateSigningRequest, opts
func (_m *CertificateSigningRequestInterface) UpdateApproval(ctx context.Context, certificateSigningRequestName string, certificateSigningRequest *certificatesv1.CertificateSigningRequest, opts metav1.UpdateOptions) (*certificatesv1.CertificateSigningRequest, error) {
	ret := _m.Called(ctx, certificateSigningRequestName, certificateSigningRequest, opts)

	var r0 *certificatesv1.CertificateSigningRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, *certificatesv1.CertificateSigningRequest, metav1.UpdateOptions) *certificatesv1.CertificateSigningRequest); ok {
		r0 = rf(ctx, certificateSigningRequestName, certificateSigningRequest, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*certificatesv1.CertificateSigningRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *certificatesv1.CertificateSigningRequest, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, certificateSigningRequestName, certificateSigningRequest, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: ctx, certificateSigningRequest, opts
func (_m *CertificateSigningRequestInterface) UpdateStatus(ctx context.Context, certificateSigningRequest *certificatesv1.CertificateSigningRequest, opts metav1.UpdateOptions) (*certificatesv1.CertificateSigningRequest, error) {
	ret := _m.Called(ctx, certificateSigningRequest, opts)

	var r0 *certificatesv1.CertificateSigningRequest
	if rf, ok := ret.Get(0).(func(context.Context, *certificatesv1.CertificateSigningRequest, metav1.UpdateOptions) *certificatesv1.CertificateSigningRequest); ok {
		r0 = rf(ctx, certificateSigningRequest, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*certificatesv1.CertificateSigningRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *certificatesv1.CertificateSigningRequest, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, certificateSigningRequest, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *CertificateSigningRequestInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCertificateSigningRequestInterface creates a new instance of CertificateSigningRequestInterface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCertificateSigningRequestInterface(t testing.TB) *CertificateSigningRequestInterface {
	mock := &CertificateSigningRequestInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

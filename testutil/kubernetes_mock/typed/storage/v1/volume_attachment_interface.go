// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	storagev1 "k8s.io/api/storage/v1"

	testing "testing"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/storage/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// VolumeAttachmentInterface is an autogenerated mock type for the VolumeAttachmentInterface type
type VolumeAttachmentInterface struct {
	mock.Mock
}

// Apply provides a mock function with given fields: ctx, volumeAttachment, opts
func (_m *VolumeAttachmentInterface) Apply(ctx context.Context, volumeAttachment *v1.VolumeAttachmentApplyConfiguration, opts metav1.ApplyOptions) (*storagev1.VolumeAttachment, error) {
	ret := _m.Called(ctx, volumeAttachment, opts)

	var r0 *storagev1.VolumeAttachment
	if rf, ok := ret.Get(0).(func(context.Context, *v1.VolumeAttachmentApplyConfiguration, metav1.ApplyOptions) *storagev1.VolumeAttachment); ok {
		r0 = rf(ctx, volumeAttachment, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1.VolumeAttachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.VolumeAttachmentApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, volumeAttachment, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApplyStatus provides a mock function with given fields: ctx, volumeAttachment, opts
func (_m *VolumeAttachmentInterface) ApplyStatus(ctx context.Context, volumeAttachment *v1.VolumeAttachmentApplyConfiguration, opts metav1.ApplyOptions) (*storagev1.VolumeAttachment, error) {
	ret := _m.Called(ctx, volumeAttachment, opts)

	var r0 *storagev1.VolumeAttachment
	if rf, ok := ret.Get(0).(func(context.Context, *v1.VolumeAttachmentApplyConfiguration, metav1.ApplyOptions) *storagev1.VolumeAttachment); ok {
		r0 = rf(ctx, volumeAttachment, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1.VolumeAttachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.VolumeAttachmentApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, volumeAttachment, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, volumeAttachment, opts
func (_m *VolumeAttachmentInterface) Create(ctx context.Context, volumeAttachment *storagev1.VolumeAttachment, opts metav1.CreateOptions) (*storagev1.VolumeAttachment, error) {
	ret := _m.Called(ctx, volumeAttachment, opts)

	var r0 *storagev1.VolumeAttachment
	if rf, ok := ret.Get(0).(func(context.Context, *storagev1.VolumeAttachment, metav1.CreateOptions) *storagev1.VolumeAttachment); ok {
		r0 = rf(ctx, volumeAttachment, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1.VolumeAttachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *storagev1.VolumeAttachment, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, volumeAttachment, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *VolumeAttachmentInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
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
func (_m *VolumeAttachmentInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
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
func (_m *VolumeAttachmentInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*storagev1.VolumeAttachment, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *storagev1.VolumeAttachment
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *storagev1.VolumeAttachment); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1.VolumeAttachment)
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
func (_m *VolumeAttachmentInterface) List(ctx context.Context, opts metav1.ListOptions) (*storagev1.VolumeAttachmentList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *storagev1.VolumeAttachmentList
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *storagev1.VolumeAttachmentList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1.VolumeAttachmentList)
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
func (_m *VolumeAttachmentInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*storagev1.VolumeAttachment, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storagev1.VolumeAttachment
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *storagev1.VolumeAttachment); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1.VolumeAttachment)
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

// Update provides a mock function with given fields: ctx, volumeAttachment, opts
func (_m *VolumeAttachmentInterface) Update(ctx context.Context, volumeAttachment *storagev1.VolumeAttachment, opts metav1.UpdateOptions) (*storagev1.VolumeAttachment, error) {
	ret := _m.Called(ctx, volumeAttachment, opts)

	var r0 *storagev1.VolumeAttachment
	if rf, ok := ret.Get(0).(func(context.Context, *storagev1.VolumeAttachment, metav1.UpdateOptions) *storagev1.VolumeAttachment); ok {
		r0 = rf(ctx, volumeAttachment, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1.VolumeAttachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *storagev1.VolumeAttachment, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, volumeAttachment, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: ctx, volumeAttachment, opts
func (_m *VolumeAttachmentInterface) UpdateStatus(ctx context.Context, volumeAttachment *storagev1.VolumeAttachment, opts metav1.UpdateOptions) (*storagev1.VolumeAttachment, error) {
	ret := _m.Called(ctx, volumeAttachment, opts)

	var r0 *storagev1.VolumeAttachment
	if rf, ok := ret.Get(0).(func(context.Context, *storagev1.VolumeAttachment, metav1.UpdateOptions) *storagev1.VolumeAttachment); ok {
		r0 = rf(ctx, volumeAttachment, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1.VolumeAttachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *storagev1.VolumeAttachment, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, volumeAttachment, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *VolumeAttachmentInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
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

// NewVolumeAttachmentInterface creates a new instance of VolumeAttachmentInterface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewVolumeAttachmentInterface(t testing.TB) *VolumeAttachmentInterface {
	mock := &VolumeAttachmentInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

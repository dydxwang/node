// Code generated by mockery v1.0.0. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta1 "k8s.io/api/extensions/v1beta1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// ReplicaSetInterface is an autogenerated mock type for the ReplicaSetInterface type
type ReplicaSetInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, replicaSet, opts
func (_m *ReplicaSetInterface) Create(ctx context.Context, replicaSet *v1beta1.ReplicaSet, opts v1.CreateOptions) (*v1beta1.ReplicaSet, error) {
	ret := _m.Called(ctx, replicaSet, opts)

	var r0 *v1beta1.ReplicaSet
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.ReplicaSet, v1.CreateOptions) *v1beta1.ReplicaSet); ok {
		r0 = rf(ctx, replicaSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.ReplicaSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.ReplicaSet, v1.CreateOptions) error); ok {
		r1 = rf(ctx, replicaSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *ReplicaSetInterface) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *ReplicaSetInterface) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.DeleteOptions, v1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *ReplicaSetInterface) Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ReplicaSet, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *v1beta1.ReplicaSet
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) *v1beta1.ReplicaSet); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.ReplicaSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, v1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScale provides a mock function with given fields: ctx, replicaSetName, options
func (_m *ReplicaSetInterface) GetScale(ctx context.Context, replicaSetName string, options v1.GetOptions) (*v1beta1.Scale, error) {
	ret := _m.Called(ctx, replicaSetName, options)

	var r0 *v1beta1.Scale
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) *v1beta1.Scale); ok {
		r0 = rf(ctx, replicaSetName, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Scale)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, v1.GetOptions) error); ok {
		r1 = rf(ctx, replicaSetName, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, opts
func (_m *ReplicaSetInterface) List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ReplicaSetList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *v1beta1.ReplicaSetList
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) *v1beta1.ReplicaSetList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.ReplicaSetList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *ReplicaSetInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (*v1beta1.ReplicaSet, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1beta1.ReplicaSet
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) *v1beta1.ReplicaSet); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.ReplicaSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, replicaSet, opts
func (_m *ReplicaSetInterface) Update(ctx context.Context, replicaSet *v1beta1.ReplicaSet, opts v1.UpdateOptions) (*v1beta1.ReplicaSet, error) {
	ret := _m.Called(ctx, replicaSet, opts)

	var r0 *v1beta1.ReplicaSet
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.ReplicaSet, v1.UpdateOptions) *v1beta1.ReplicaSet); ok {
		r0 = rf(ctx, replicaSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.ReplicaSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.ReplicaSet, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, replicaSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateScale provides a mock function with given fields: ctx, replicaSetName, scale, opts
func (_m *ReplicaSetInterface) UpdateScale(ctx context.Context, replicaSetName string, scale *v1beta1.Scale, opts v1.UpdateOptions) (*v1beta1.Scale, error) {
	ret := _m.Called(ctx, replicaSetName, scale, opts)

	var r0 *v1beta1.Scale
	if rf, ok := ret.Get(0).(func(context.Context, string, *v1beta1.Scale, v1.UpdateOptions) *v1beta1.Scale); ok {
		r0 = rf(ctx, replicaSetName, scale, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Scale)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *v1beta1.Scale, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, replicaSetName, scale, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: ctx, replicaSet, opts
func (_m *ReplicaSetInterface) UpdateStatus(ctx context.Context, replicaSet *v1beta1.ReplicaSet, opts v1.UpdateOptions) (*v1beta1.ReplicaSet, error) {
	ret := _m.Called(ctx, replicaSet, opts)

	var r0 *v1beta1.ReplicaSet
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.ReplicaSet, v1.UpdateOptions) *v1beta1.ReplicaSet); ok {
		r0 = rf(ctx, replicaSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.ReplicaSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.ReplicaSet, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, replicaSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *ReplicaSetInterface) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
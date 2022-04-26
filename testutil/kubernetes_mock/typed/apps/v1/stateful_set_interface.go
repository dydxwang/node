// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"

	context "context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	testing "testing"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/apps/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// StatefulSetInterface is an autogenerated mock type for the StatefulSetInterface type
type StatefulSetInterface struct {
	mock.Mock
}

// Apply provides a mock function with given fields: ctx, statefulSet, opts
func (_m *StatefulSetInterface) Apply(ctx context.Context, statefulSet *v1.StatefulSetApplyConfiguration, opts metav1.ApplyOptions) (*appsv1.StatefulSet, error) {
	ret := _m.Called(ctx, statefulSet, opts)

	var r0 *appsv1.StatefulSet
	if rf, ok := ret.Get(0).(func(context.Context, *v1.StatefulSetApplyConfiguration, metav1.ApplyOptions) *appsv1.StatefulSet); ok {
		r0 = rf(ctx, statefulSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.StatefulSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.StatefulSetApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, statefulSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApplyStatus provides a mock function with given fields: ctx, statefulSet, opts
func (_m *StatefulSetInterface) ApplyStatus(ctx context.Context, statefulSet *v1.StatefulSetApplyConfiguration, opts metav1.ApplyOptions) (*appsv1.StatefulSet, error) {
	ret := _m.Called(ctx, statefulSet, opts)

	var r0 *appsv1.StatefulSet
	if rf, ok := ret.Get(0).(func(context.Context, *v1.StatefulSetApplyConfiguration, metav1.ApplyOptions) *appsv1.StatefulSet); ok {
		r0 = rf(ctx, statefulSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.StatefulSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.StatefulSetApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, statefulSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, statefulSet, opts
func (_m *StatefulSetInterface) Create(ctx context.Context, statefulSet *appsv1.StatefulSet, opts metav1.CreateOptions) (*appsv1.StatefulSet, error) {
	ret := _m.Called(ctx, statefulSet, opts)

	var r0 *appsv1.StatefulSet
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1.StatefulSet, metav1.CreateOptions) *appsv1.StatefulSet); ok {
		r0 = rf(ctx, statefulSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.StatefulSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *appsv1.StatefulSet, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, statefulSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *StatefulSetInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
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
func (_m *StatefulSetInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
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
func (_m *StatefulSetInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*appsv1.StatefulSet, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *appsv1.StatefulSet
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *appsv1.StatefulSet); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.StatefulSet)
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

// GetScale provides a mock function with given fields: ctx, statefulSetName, options
func (_m *StatefulSetInterface) GetScale(ctx context.Context, statefulSetName string, options metav1.GetOptions) (*autoscalingv1.Scale, error) {
	ret := _m.Called(ctx, statefulSetName, options)

	var r0 *autoscalingv1.Scale
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *autoscalingv1.Scale); ok {
		r0 = rf(ctx, statefulSetName, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingv1.Scale)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, statefulSetName, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, opts
func (_m *StatefulSetInterface) List(ctx context.Context, opts metav1.ListOptions) (*appsv1.StatefulSetList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *appsv1.StatefulSetList
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *appsv1.StatefulSetList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.StatefulSetList)
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
func (_m *StatefulSetInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*appsv1.StatefulSet, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *appsv1.StatefulSet
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *appsv1.StatefulSet); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.StatefulSet)
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

// Update provides a mock function with given fields: ctx, statefulSet, opts
func (_m *StatefulSetInterface) Update(ctx context.Context, statefulSet *appsv1.StatefulSet, opts metav1.UpdateOptions) (*appsv1.StatefulSet, error) {
	ret := _m.Called(ctx, statefulSet, opts)

	var r0 *appsv1.StatefulSet
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1.StatefulSet, metav1.UpdateOptions) *appsv1.StatefulSet); ok {
		r0 = rf(ctx, statefulSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.StatefulSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *appsv1.StatefulSet, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, statefulSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateScale provides a mock function with given fields: ctx, statefulSetName, scale, opts
func (_m *StatefulSetInterface) UpdateScale(ctx context.Context, statefulSetName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (*autoscalingv1.Scale, error) {
	ret := _m.Called(ctx, statefulSetName, scale, opts)

	var r0 *autoscalingv1.Scale
	if rf, ok := ret.Get(0).(func(context.Context, string, *autoscalingv1.Scale, metav1.UpdateOptions) *autoscalingv1.Scale); ok {
		r0 = rf(ctx, statefulSetName, scale, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingv1.Scale)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *autoscalingv1.Scale, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, statefulSetName, scale, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: ctx, statefulSet, opts
func (_m *StatefulSetInterface) UpdateStatus(ctx context.Context, statefulSet *appsv1.StatefulSet, opts metav1.UpdateOptions) (*appsv1.StatefulSet, error) {
	ret := _m.Called(ctx, statefulSet, opts)

	var r0 *appsv1.StatefulSet
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1.StatefulSet, metav1.UpdateOptions) *appsv1.StatefulSet); ok {
		r0 = rf(ctx, statefulSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.StatefulSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *appsv1.StatefulSet, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, statefulSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *StatefulSetInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
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

// NewStatefulSetInterface creates a new instance of StatefulSetInterface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewStatefulSetInterface(t testing.TB) *StatefulSetInterface {
	mock := &StatefulSetInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta1"
)

// PriorityLevelConfigurationsGetter is an autogenerated mock type for the PriorityLevelConfigurationsGetter type
type PriorityLevelConfigurationsGetter struct {
	mock.Mock
}

// PriorityLevelConfigurations provides a mock function with given fields:
func (_m *PriorityLevelConfigurationsGetter) PriorityLevelConfigurations() v1beta1.PriorityLevelConfigurationInterface {
	ret := _m.Called()

	var r0 v1beta1.PriorityLevelConfigurationInterface
	if rf, ok := ret.Get(0).(func() v1beta1.PriorityLevelConfigurationInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.PriorityLevelConfigurationInterface)
		}
	}

	return r0
}

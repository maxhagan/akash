// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1alpha1 "k8s.io/client-go/kubernetes/typed/apiserverinternal/v1alpha1"
)

// StorageVersionsGetter is an autogenerated mock type for the StorageVersionsGetter type
type StorageVersionsGetter struct {
	mock.Mock
}

// StorageVersions provides a mock function with given fields:
func (_m *StorageVersionsGetter) StorageVersions() v1alpha1.StorageVersionInterface {
	ret := _m.Called()

	var r0 v1alpha1.StorageVersionInterface
	if rf, ok := ret.Get(0).(func() v1alpha1.StorageVersionInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.StorageVersionInterface)
		}
	}

	return r0
}

// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	podvolume "github.com/zerospiel/velero/pkg/podvolume"

	volume "github.com/zerospiel/velero/internal/volume"
)

// Restorer is an autogenerated mock type for the Restorer type
type Restorer struct {
	mock.Mock
}

// RestorePodVolumes provides a mock function with given fields: _a0, _a1
func (_m *Restorer) RestorePodVolumes(_a0 podvolume.RestoreData, _a1 *volume.RestoreVolumeInfoTracker) []error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for RestorePodVolumes")
	}

	var r0 []error
	if rf, ok := ret.Get(0).(func(podvolume.RestoreData, *volume.RestoreVolumeInfoTracker) []error); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]error)
		}
	}

	return r0
}

// NewRestorer creates a new instance of Restorer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRestorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Restorer {
	mock := &Restorer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

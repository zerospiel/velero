// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	uploader "github.com/zerospiel/velero/pkg/uploader"
)

// Provider is an autogenerated mock type for the Provider type
type Provider struct {
	mock.Mock
}

// Close provides a mock function with given fields: ctx
func (_m *Provider) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RunBackup provides a mock function with given fields: ctx, path, realSource, tags, forceFull, parentSnapshot, volMode, uploaderCfg, updater
func (_m *Provider) RunBackup(ctx context.Context, path string, realSource string, tags map[string]string, forceFull bool, parentSnapshot string, volMode uploader.PersistentVolumeMode, uploaderCfg map[string]string, updater uploader.ProgressUpdater) (string, bool, error) {
	ret := _m.Called(ctx, path, realSource, tags, forceFull, parentSnapshot, volMode, uploaderCfg, updater)

	var r0 string
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, map[string]string, bool, string, uploader.PersistentVolumeMode, map[string]string, uploader.ProgressUpdater) (string, bool, error)); ok {
		return rf(ctx, path, realSource, tags, forceFull, parentSnapshot, volMode, uploaderCfg, updater)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, map[string]string, bool, string, uploader.PersistentVolumeMode, map[string]string, uploader.ProgressUpdater) string); ok {
		r0 = rf(ctx, path, realSource, tags, forceFull, parentSnapshot, volMode, uploaderCfg, updater)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, map[string]string, bool, string, uploader.PersistentVolumeMode, map[string]string, uploader.ProgressUpdater) bool); ok {
		r1 = rf(ctx, path, realSource, tags, forceFull, parentSnapshot, volMode, uploaderCfg, updater)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, map[string]string, bool, string, uploader.PersistentVolumeMode, map[string]string, uploader.ProgressUpdater) error); ok {
		r2 = rf(ctx, path, realSource, tags, forceFull, parentSnapshot, volMode, uploaderCfg, updater)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RunRestore provides a mock function with given fields: ctx, snapshotID, volumePath, volMode, uploaderConfig, updater
func (_m *Provider) RunRestore(ctx context.Context, snapshotID string, volumePath string, volMode uploader.PersistentVolumeMode, uploaderConfig map[string]string, updater uploader.ProgressUpdater) error {
	ret := _m.Called(ctx, snapshotID, volumePath, volMode, uploaderConfig, updater)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, uploader.PersistentVolumeMode, map[string]string, uploader.ProgressUpdater) error); ok {
		r0 = rf(ctx, snapshotID, volumePath, volMode, uploaderConfig, updater)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewProvider creates a new instance of Provider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProvider(t mockConstructorTestingTNewProvider) *Provider {
	mock := &Provider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

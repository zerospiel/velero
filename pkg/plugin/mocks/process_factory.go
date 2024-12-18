/*
Copyright the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	logrus "github.com/sirupsen/logrus"
	mock "github.com/stretchr/testify/mock"

	process "github.com/zerospiel/velero/pkg/plugin/clientmgmt/process"
)

// ProcessFactory is an autogenerated mock type for the ProcessFactory type
type ProcessFactory struct {
	mock.Mock
}

// newProcess provides a mock function with given fields: command, logger, logLevel
func (_m *ProcessFactory) newProcess(command string, logger logrus.FieldLogger, logLevel logrus.Level) (process.Process, error) {
	ret := _m.Called(command, logger, logLevel)

	var r0 process.Process
	if rf, ok := ret.Get(0).(func(string, logrus.FieldLogger, logrus.Level) process.Process); ok {
		r0 = rf(command, logger, logLevel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(process.Process)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, logrus.FieldLogger, logrus.Level) error); ok {
		r1 = rf(command, logger, logLevel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

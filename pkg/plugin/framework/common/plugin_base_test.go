/*
Copyright 2018, 2019 the Velero contributors.

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
package common

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zerospiel/velero/pkg/test"
)

func TestClientLogger(t *testing.T) {
	base := &PluginBase{}
	logger := test.NewLogger()
	f := ClientLogger(logger)
	f(base)
	assert.Equal(t, logger, base.ClientLogger)
}

func TestServerLogger(t *testing.T) {
	base := &PluginBase{}
	logger := test.NewLogger()
	f := ServerLogger(logger)
	f(base)
	assert.Equal(t, NewServerMux(logger), base.ServerMux)
}

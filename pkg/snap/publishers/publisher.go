// Copyright (c) 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package publishers

import (
	"github.com/intelsdi-x/snap/scheduler/wmap"
	"github.com/intelsdi-x/swan/pkg/conf"
)

// Publisher stores default publisher object
type Publisher struct {
	PluginName string
	Publisher  *wmap.PublishWorkflowMapNode
}

// NewDefaultPublisher construct new snap publisher object
// based on default flag in configuration.
func NewDefaultPublisher() Publisher {

	if conf.DefaultSnapPublisher.Value() == "influxdb" {
		return NewDefaultInfluxDBPublisher()
	}
	// Default is cassandra
	return NewDefaultCassandraPublisher()
}

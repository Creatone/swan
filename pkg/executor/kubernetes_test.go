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

package executor

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/v1"
)

func TestKubernetes(t *testing.T) {

	Convey("Kubernetes config is sane by default", t, func() {
		config := DefaultKubernetesConfig()
		So(config.Privileged, ShouldEqual, false)
		So(config.HostNetwork, ShouldEqual, false)
		So(config.ContainerImage, ShouldEqual, defaultContainerImage)
	})

	Convey("Kubernetes pod executor pod names", t, func() {

		Convey("have desired name", func() {
			podExecutor := &k8s{KubernetesConfig{PodName: "foo"}, nil}
			name := podExecutor.generatePodName()
			So(name, ShouldEqual, "foo")

		})

		Convey("have desired prefix", func() {
			podExecutor := &k8s{KubernetesConfig{PodNamePrefix: "foo"}, nil}
			name := podExecutor.generatePodName()
			So(name, ShouldStartWith, "foo-")

		})

		Convey("with default config", func() {

			podExecutor := &k8s{DefaultKubernetesConfig(), nil}

			Convey("have default prefix", func() {
				name := podExecutor.generatePodName()
				So(name, ShouldStartWith, "swan-")
			})

			Convey("are unique", func() {
				names := make(map[string]struct{})
				N := 1000
				for i := 0; i < N; i++ {
					name := podExecutor.generatePodName()
					names[name] = struct{}{}
				}
				So(names, ShouldHaveLength, N)
			})
		})
	})

}

func v1ToAPI(v1Pod *v1.Pod) *api.Pod {
	apiPod := &api.Pod{}
	scheme := NewRuntimeScheme()
	err := scheme.Convert(v1Pod, apiPod, nil)
	if err != nil {
		panic(err)
	}
	return apiPod
}

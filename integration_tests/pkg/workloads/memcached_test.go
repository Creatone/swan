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

package workloads

import (
	"io/ioutil"
	"testing"

	"github.com/intelsdi-x/swan/pkg/executor"
	"github.com/intelsdi-x/swan/pkg/utils/env"
	"github.com/intelsdi-x/swan/pkg/workloads/memcached"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	netstatCommand = "echo stats | nc -w 1 127.0.0.1 11211"
)

// TestMemcachedWithExecutor is an integration test with local executor.
func TestMemcachedWithExecutor(t *testing.T) {
	log.SetLevel(log.ErrorLevel)

	Convey("While using Local Shell in Memcached launcher", t, func() {
		l := executor.NewLocal()
		config := memcached.DefaultMemcachedConfig()
		// Prefer to run memcached locally using the current user,
		// if it can be determined from the environment.
		config.User = env.GetOrDefault("USER", config.User)
		memcachedLauncher := memcached.New(l, config)

		Convey("When memcached is launched", func() {
			// NOTE: It is needed for memcached to have default port available.
			taskHandle, err := memcachedLauncher.Launch()
			So(err, ShouldBeNil)
			So(taskHandle, ShouldNotBeNil)
			defer taskHandle.Stop()
			defer taskHandle.EraseOutput()

			Convey("There should be no error", func() {
				stopErr := taskHandle.Stop()

				So(err, ShouldBeNil)
				So(stopErr, ShouldBeNil)
			})

			Convey("When we check the memcached endpoint for stats after 1 second", func() {
				netstatTaskHandle, netstatErr := l.Execute(netstatCommand)
				if netstatTaskHandle != nil {
					defer netstatTaskHandle.Stop()
					defer netstatTaskHandle.EraseOutput()
				}
				Convey("There should be no error", func() {
					taskHandle.Stop()
					netstatTaskHandle.Stop()

					So(netstatErr, ShouldBeNil)

				})

				Convey("When we wait for netstat ", func() {
					netstatTaskHandle.Wait(0)

					Convey("The netstat task should be terminated, the task status should be 0"+
						" and output resultes with a STAT information", func() {

						netstatTaskState := netstatTaskHandle.Status()
						So(netstatTaskState, ShouldEqual, executor.TERMINATED)

						exitCode, err := netstatTaskHandle.ExitCode()
						So(err, ShouldBeNil)
						So(exitCode, ShouldEqual, 0)

						stdoutFile, stdoutErr := netstatTaskHandle.StdoutFile()

						So(stdoutErr, ShouldBeNil)
						So(stdoutFile, ShouldNotBeNil)

						data, readErr := ioutil.ReadAll(stdoutFile)
						So(readErr, ShouldBeNil)
						So(string(data[:]), ShouldStartWith, "STAT")
					})
				})
			})

			Convey("When we stop the memcached task", func() {
				err := taskHandle.Stop()

				Convey("There should be no error", func() {
					So(err, ShouldBeNil)
				})

				Convey("The task should be terminated and the task status "+
					"should be -1 or 0", func() {

					taskState := taskHandle.Status()
					So(taskState, ShouldEqual, executor.TERMINATED)

					exitCode, err := taskHandle.ExitCode()

					So(err, ShouldBeNil)
					// Memcached on CentOS returns 0 (successful code) after SIGTERM.
					So(exitCode, ShouldBeIn, -1, 0)
				})
			})
		})
	})
}

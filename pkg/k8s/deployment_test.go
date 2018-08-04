// Copyright 2018 John Deng (hi.devops.io@gmail.com).
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


package k8s


import (
	"testing"
	"github.com/magiconair/properties/assert"
	"istio.io/istio/pkg/log"
)


func TestDeploy(t *testing.T)  {
	deploy := Deployment{
		App: "hello-world",
		Project: "demo",
		Profile: "master",
	}
	d, err := deploy.Deploy()
	assert.Equal(t, nil, err)
	log.Info(d)
}
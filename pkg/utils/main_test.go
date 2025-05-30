/*
Copyright 2022 The KubeVela Authors.

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

package utils

import (
	"testing"

	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/kubevela/pkg/util/singleton"
	"github.com/kubevela/workflow/api/v1alpha1"
)

var (
	cli client.Client
)

func TestMain(m *testing.M) {
	sc := scheme.Scheme
	_ = v1alpha1.AddToScheme(sc)
	cli = fake.NewClientBuilder().WithScheme(sc).WithStatusSubresource(&v1alpha1.WorkflowRun{}).Build()
	singleton.KubeClient.Set(cli)
	m.Run()
}

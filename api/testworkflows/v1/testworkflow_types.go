/*
Copyright 2021.
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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// TestWorkflowSpec defines the desired state of TestWorkflow
type TestWorkflowSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	TestWorkflowSpecBase `json:",inline"`

	// templates to include at a top-level of workflow
	Use []TemplateRef `json:"use,omitempty"`
}

// TemplateRef is the reference for the template inclusion
type TemplateRef struct {
	// name of the template to include
	Name string `json:"name"`
	// trait configuration values if needed
	Config *map[string]intstr.IntOrString `json:"config,omitempty"`
}

// +kubebuilder:object:root=true

// TestWorkflow is the Schema for the workflows API
type TestWorkflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// TestWorkflow readable description
	Description string `json:"description,omitempty"`

	// TestWorkflow specification
	Spec TestWorkflowSpec `json:"spec"`
}

//+kubebuilder:object:root=true

// TestWorkflowList contains a list of TestWorkflow
type TestWorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestWorkflow `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TestWorkflow{}, &TestWorkflowList{})
}

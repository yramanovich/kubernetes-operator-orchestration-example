package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CoinWorkflowSpec defines the desired state of CoinWorkflow
type CoinWorkflowSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of CoinWorkflow. Edit coinworkflow_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// CoinWorkflowStatus defines the observed state of CoinWorkflow
type CoinWorkflowStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CoinWorkflow is the Schema for the coinworkflows API
type CoinWorkflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CoinWorkflowSpec   `json:"spec,omitempty"`
	Status CoinWorkflowStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CoinWorkflowList contains a list of CoinWorkflow
type CoinWorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CoinWorkflow `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CoinWorkflow{}, &CoinWorkflowList{})
}

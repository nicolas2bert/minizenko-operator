package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type CloudSpecServer struct {
	RemoteManagementDisable bool `json:"remoteManagementDisable,omitempty"`
}

// MinizenkoSpec defines the desired state of Minizenko
// +k8s:openapi-gen=true
type MinizenkoSpec struct {
	CloudServer CloudSpecServer `json:"cloudServer,omitempty"`
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// MinizenkoStatus defines the observed state of Minizenko
// +k8s:openapi-gen=true
type MinizenkoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Minizenko is the Schema for the minizenkos API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Minizenko struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MinizenkoSpec   `json:"spec,omitempty"`
	Status MinizenkoStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MinizenkoList contains a list of Minizenko
type MinizenkoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Minizenko `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Minizenko{}, &MinizenkoList{})
}

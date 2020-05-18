package v1aplpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BobRGeoBType is a top-level type
type BobRGeoBType struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Status BobRGeoBTypeStatus `json:"status,omitempty"`
	// This is where you can define
	// your own custom spec
	Spec BobRGeoBSpec `json:"spec,omitempty"`
}

// custom spec
type BobRGeoBSpec struct {
	Image    string `json:"image,omitempty"`
	Replicas int    `json:"replicas,omitempty"`
}

// custom status
type BobRGeoBTypeStatus struct {
	Name string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// no client needed for list as it's been created in above
type BobRGeoBTypeList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []BobRGeoBType `json:"items"`
}

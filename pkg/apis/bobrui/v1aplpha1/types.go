package v1aplpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BobRUIBType is a top-level type
type BobRUIBType struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Status BobRUIBTypeStatus `json:"status,omitempty"`
	// This is where you can define
	// your own custom spec
	Spec BobRUIBSpec `json:"spec,omitempty"`
}

// custom spec
type BobRUIBSpec struct {
	Image    string `json:"image,omitempty"`
	Replicas int    `json:"replicas,omitempty"`
}

// custom status
type BobRUIBTypeStatus struct {
	Name string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// no client needed for list as it's been created in above
type BobRUIBTypeList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []BobRUIBType `json:"items"`
}

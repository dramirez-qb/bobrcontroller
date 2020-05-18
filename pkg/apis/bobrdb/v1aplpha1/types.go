package v1aplpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BobRDBType is a top-level type
type BobRDBType struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Status BobRDBTypeStatus `json:"status,omitempty"`
	// This is where you can define
	// your own custom spec
	Spec BobRDBSpec `json:"spec,omitempty"`
}

// custom spec
type BobRDBSpec struct {
	Image    string `json:"image,omitempty"`
	Replicas int    `json:"replicas,omitempty"`
}

// custom status
type BobRDBTypeStatus struct {
	Name string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// no client needed for list as it's been created in above
type BobRDBTypeList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []BobRDBType `json:"items"`
}

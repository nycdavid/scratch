package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

package v1

import (
metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
"k8s.io/apimachinery/pkg/runtime"
"k8s.io/apimachinery/pkg/runtime/schema"
)

// ------------------------------------------------------------
// Group/Version
// ------------------------------------------------------------

var (
	SchemeGroupVersion = schema.GroupVersion{
		Group:   "vehicles.example.com",
		Version: "v1",
	}
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// ------------------------------------------------------------
// Car Spec & Status (minimal structs)
// ------------------------------------------------------------

type CarSpec struct {
	// Optional: whatever fields match your CRD
}

type CarStatus struct {
	// Optional: whatever fields match your CRD
}

// ------------------------------------------------------------
// Car & CarList types
// ------------------------------------------------------------

type Car struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CarSpec   `json:"spec,omitempty"`
	Status CarStatus `json:"status,omitempty"`
}

type CarList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Car `json:"items"`
}

// ------------------------------------------------------------
// Register the types in the Scheme
// ------------------------------------------------------------

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Car{},
		&CarList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

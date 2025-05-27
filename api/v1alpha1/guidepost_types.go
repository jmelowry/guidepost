package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Guidepost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GuidepostSpec   `json:"spec,omitempty"`
	Status GuidepostStatus `json:"status,omitempty"`
}

type GuidepostSpec struct {
	TemplateRef string            `json:"templateRef"`
	Parameters  map[string]string `json:"parameters,omitempty"`
}

type GuidepostStatus struct {
	ResolvedLinks []string `json:"resolvedLinks,omitempty"`
}

// +kubebuilder:object:root=true
type GuidepostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Guidepost `json:"items"`
}

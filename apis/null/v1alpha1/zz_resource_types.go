/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ResourceInitParameters struct {

	// running any associated provisioners.
	// A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
	// +mapType=granular
	Triggers map[string]*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

type ResourceObservation struct {

	// (String) This is set to a random value at create time.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// running any associated provisioners.
	// A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
	// +mapType=granular
	Triggers map[string]*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

type ResourceParameters struct {

	// running any associated provisioners.
	// A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Triggers map[string]*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

// ResourceSpec defines the desired state of Resource
type ResourceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ResourceInitParameters `json:"initProvider,omitempty"`
}

// ResourceStatus defines the observed state of Resource.
type ResourceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Resource is the Schema for the Resources API. The null_resource resource implements the standard resource lifecycle but takes no further action.hashicorp. The triggers argument allows specifying an arbitrary set of values that, when changed, will cause the resource to be replaced.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,template}
type Resource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceSpec   `json:"spec"`
	Status            ResourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceList contains a list of Resources
type ResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Resource `json:"items"`
}

// Repository type metadata.
var (
	Resource_Kind             = "Resource"
	Resource_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Resource_Kind}.String()
	Resource_KindAPIVersion   = Resource_Kind + "." + CRDGroupVersion.String()
	Resource_GroupVersionKind = CRDGroupVersion.WithKind(Resource_Kind)
)

func init() {
	SchemeBuilder.Register(&Resource{}, &ResourceList{})
}

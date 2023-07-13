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

type MembershipsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MembershipsParameters struct {

	// +crossplane:generate:reference:type=Group
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Required
	RealmID *string `json:"realmId" tf:"realm_id,omitempty"`
}

// MembershipsSpec defines the desired state of Memberships
type MembershipsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MembershipsParameters `json:"forProvider"`
}

// MembershipsStatus defines the observed state of Memberships.
type MembershipsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MembershipsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Memberships is the Schema for the Membershipss API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type Memberships struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MembershipsSpec   `json:"spec"`
	Status            MembershipsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MembershipsList contains a list of Membershipss
type MembershipsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Memberships `json:"items"`
}

// Repository type metadata.
var (
	Memberships_Kind             = "Memberships"
	Memberships_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Memberships_Kind}.String()
	Memberships_KindAPIVersion   = Memberships_Kind + "." + CRDGroupVersion.String()
	Memberships_GroupVersionKind = CRDGroupVersion.WithKind(Memberships_Kind)
)

func init() {
	SchemeBuilder.Register(&Memberships{}, &MembershipsList{})
}
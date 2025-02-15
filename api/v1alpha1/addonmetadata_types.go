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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AddonMetadataSpec defines the desired state of AddonMetadata
// View markers: $ controller-gen -www crd
// TODO add missing fields from schema, only required fields from jsonschema are present
type AddonMetadataSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]{0,30}[A-Za-z0-9]$`
	// Unique ID of the addon
	ID string `json:"id" validate:"required"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[0-9A-Z][A-Za-z0-9-_ ()]+$`
	// Friendly name for the addon, displayed in the UI
	Name string `json:"name" validate:"required"`

	// +kubebuilder:validation:Required
	// Short description for the addon
	Description string `json:"description" validate:"required"`

	// +optional
	// +kubebuilder:validation:Pattern=`^http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+$`
	// Link to the addon documentation
	Link string `json:"link"`

	// +kubebuilder:validation:Required
	// Icon to be shown in UI. Should be around 200px and base64 encoded.
	Icon string `json:"icon" validate:"required"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^api\.openshift\.com/addon-[0-9a-z][0-9a-z-]{0,30}[0-9a-z]$`
	// Kubernetes label for the addon. Needs to match: 'api.openshift.com/<addon-id>'.
	Label string `json:"label" validate:"required"`

	// +kubebuilder:validation:Required
	// Set to true to allow installation of the addon.
	Enabled bool `json:"enabled" validate:"required"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^([A-Za-z -]+ <[0-9A-Za-z_.-]+@redhat\.com>,?)+$`
	// Team or individual responsible for this addon. Needs to match: 'some name <some-email@redhat.com>'.
	AddonOwner string `json:"addonOwner" validate:"required"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^quay\.io/osd-addons/[a-z-]+$`
	// Quay repository for the addon operator. Needs to match: 'quay.io/osd-addons/<my-addon-repo>'.
	QuayRepo string `json:"quayRepo" validate:"required"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^quay\.io/[0-9A-Za-z._-]+/[0-9A-Za-z._-]+(:[A-Za-z0-9._-]+)?$`
	// Quay repository for the testHarness image. Needs to match: 'quay.io/<my-repo>/<my-test-harness>:<my-tag>'.
	TestHarness string `json:"testHarness" validate:"required"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum={AllNamespaces,SingleNamespace,OwnNamespace}
	// OLM InstallMode for the addon operator. One of: AllNamespaces, SingleNamespace or OwnNamespace.
	InstallMode string `json:"installMode" validate:"required"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]{0,60}[A-Za-z0-9]$`
	// Namespace where the addon operator should be installed.
	TargetNamespace string `json:"targetNamespace" validate:"required"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:UniqueItems=true
	// Namespaces managed by the addon-operator. Need to include the TargetNamespace.
	Namespaces []string `json:"namespaces" validate:"required"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-_]{0,35}[A-Za-z0-9]$`
	// Refers to the SKU name for the addon.
	OcmQuotaName string `json:"ocmQuotaName" validate:"required"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=0
	// TODO: what is this?
	OcmQuotaCost int `json:"ocmQuotaCost" validate:"required"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]{0,30}[A-Za-z0-9]$`
	// Name of the addon operator.
	OperatorName string `json:"operatorName" validate:"required"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum={alpha,beta,stable,edge,rc}
	// OLM channel from which to install the addon-operator. One of: alpha, beta, stable, edge or rc.
	DefaultChannel string `json:"defaultChannel" validate:"required"`

	// +kubebuilder:validation:Required
	// List of channels where the addon operator is available.
	Channels []Channel `json:"channels" validate:"required,gt=0"`

	// +kubebuilder:validation:Required
	// +kubebuilder:default:{}
	// Labels to be applied on all listed namespaces.
	NamespaceLabels map[string]string `json:"namespaceLabels" validate:"required"`

	// +kubebuilder:validation:Required
	// +kubebuilder:default:{}
	// Annotations to be applied on all listed namespaces.
	NamespaceAnnotations map[string]string `json:"namespaceAnnotations" validate:"required"`

	// +optional
	// +kubebuilder:validation:Pattern=`^quay\.io/osd-addons/[a-z-]+`
	IndexImage string `json:"indexImage"`
}

// AddonMetadataStatus defines the observed state of AddonMetadata
type AddonMetadataStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AddonMetadata is the Schema for the AddonMetadata API
type AddonMetadata struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AddonMetadataSpec   `json:"spec,omitempty" validate:"required"`
	Status AddonMetadataStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AddonMetadataList contains a list of AddonMetadata
type AddonMetadataList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AddonMetadata `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AddonMetadata{}, &AddonMetadataList{})
}

// *****
// Helper types
// *****

// Channel - list all channels for a given operator
type Channel struct {
	Name       string `json:"name"`
	CurrentCSV string `json:"currentCSV"`
}

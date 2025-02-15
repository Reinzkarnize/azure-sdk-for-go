//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhardwaresecuritymodules

import "time"

// APIEntityReference - The API entity reference.
type APIEntityReference struct {
	// The ARM resource id in the form of /subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/…
	ID *string `json:"id,omitempty"`
}

// DedicatedHsm - Resource information with extended details.
type DedicatedHsm struct {
	// REQUIRED; The supported Azure location where the dedicated HSM should be created.
	Location *string `json:"location,omitempty"`

	// REQUIRED; Properties of the dedicated HSM
	Properties *DedicatedHsmProperties `json:"properties,omitempty"`

	// SKU details
	SKU *SKU `json:"sku,omitempty"`

	// Resource tags
	Tags map[string]*string `json:"tags,omitempty"`

	// The Dedicated Hsm zones.
	Zones []*string `json:"zones,omitempty"`

	// READ-ONLY; The Azure Resource Manager resource ID for the dedicated HSM.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the dedicated HSM.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The resource type of the dedicated HSM.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DedicatedHsmClientBeginCreateOrUpdateOptions contains the optional parameters for the DedicatedHsmClient.BeginCreateOrUpdate
// method.
type DedicatedHsmClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DedicatedHsmClientBeginDeleteOptions contains the optional parameters for the DedicatedHsmClient.BeginDelete method.
type DedicatedHsmClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DedicatedHsmClientBeginUpdateOptions contains the optional parameters for the DedicatedHsmClient.BeginUpdate method.
type DedicatedHsmClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DedicatedHsmClientGetOptions contains the optional parameters for the DedicatedHsmClient.Get method.
type DedicatedHsmClientGetOptions struct {
	// placeholder for future optional parameters
}

// DedicatedHsmClientListByResourceGroupOptions contains the optional parameters for the DedicatedHsmClient.ListByResourceGroup
// method.
type DedicatedHsmClientListByResourceGroupOptions struct {
	// Maximum number of results to return.
	Top *int32
}

// DedicatedHsmClientListBySubscriptionOptions contains the optional parameters for the DedicatedHsmClient.ListBySubscription
// method.
type DedicatedHsmClientListBySubscriptionOptions struct {
	// Maximum number of results to return.
	Top *int32
}

// DedicatedHsmClientListOutboundNetworkDependenciesEndpointsOptions contains the optional parameters for the DedicatedHsmClient.ListOutboundNetworkDependenciesEndpoints
// method.
type DedicatedHsmClientListOutboundNetworkDependenciesEndpointsOptions struct {
	// placeholder for future optional parameters
}

// DedicatedHsmError - The error exception.
type DedicatedHsmError struct {
	// READ-ONLY; The error detail of the operation if any.
	Error *Error `json:"error,omitempty" azure:"ro"`
}

// DedicatedHsmListResult - List of dedicated HSMs
type DedicatedHsmListResult struct {
	// The URL to get the next set of dedicated hsms.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of dedicated HSMs.
	Value []*DedicatedHsm `json:"value,omitempty"`
}

// DedicatedHsmOperation - REST API operation
type DedicatedHsmOperation struct {
	// The display string.
	Display *DedicatedHsmOperationDisplay `json:"display,omitempty"`

	// The name of the Dedicated HSM Resource Provider Operation.
	Name *string `json:"name,omitempty"`

	// READ-ONLY; Gets or sets a value indicating whether it is a data plane action
	IsDataAction *string `json:"isDataAction,omitempty" azure:"ro"`
}

// DedicatedHsmOperationDisplay - The display string.
type DedicatedHsmOperationDisplay struct {
	// The object that represents the operation.
	Description *string `json:"description,omitempty"`

	// Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// The Resource Provider of the operation
	Provider *string `json:"provider,omitempty"`

	// Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
}

// DedicatedHsmOperationListResult - Result of the request to list Dedicated HSM Provider operations. It contains a list of
// operations.
type DedicatedHsmOperationListResult struct {
	// List of Dedicated HSM Resource Provider operations.
	Value []*DedicatedHsmOperation `json:"value,omitempty"`
}

// DedicatedHsmPatchParameters - Patchable properties of the dedicated HSM
type DedicatedHsmPatchParameters struct {
	// Resource tags
	Tags map[string]*string `json:"tags,omitempty"`
}

// DedicatedHsmProperties - Properties of the dedicated hsm
type DedicatedHsmProperties struct {
	// Specifies the management network interfaces of the dedicated hsm.
	ManagementNetworkProfile *NetworkProfile `json:"managementNetworkProfile,omitempty"`

	// Specifies the network interfaces of the dedicated hsm.
	NetworkProfile *NetworkProfile `json:"networkProfile,omitempty"`

	// This field will be used when RP does not support Availability zones.
	StampID *string `json:"stampId,omitempty"`

	// READ-ONLY; Provisioning state.
	ProvisioningState *JSONWebKeyType `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Resource Status Message.
	StatusMessage *string `json:"statusMessage,omitempty" azure:"ro"`
}

// EndpointDependency - A domain name that dedicated hsm services are reaching at.
type EndpointDependency struct {
	// The domain name of the dependency.
	DomainName *string `json:"domainName,omitempty"`

	// The Ports and Protocols used when connecting to domainName.
	EndpointDetails []*EndpointDetail `json:"endpointDetails,omitempty"`
}

// EndpointDetail - Connect information from the dedicated hsm service to a single endpoint.
type EndpointDetail struct {
	// Description of the detail
	Description *string `json:"description,omitempty"`

	// An IP Address that Domain Name currently resolves to.
	IPAddress *string `json:"ipAddress,omitempty"`

	// The port an endpoint is connected to.
	Port *int32 `json:"port,omitempty"`

	// The protocol used for connection
	Protocol *string `json:"protocol,omitempty"`
}

// Error - The key vault server error.
type Error struct {
	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Contains more specific error that narrows down the cause. May be null.
	InnerError *Error `json:"innererror,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// NetworkInterface - The network interface definition.
type NetworkInterface struct {
	// Private Ip address of the interface
	PrivateIPAddress *string `json:"privateIpAddress,omitempty"`

	// READ-ONLY; The ARM resource id in the form of /subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/…
	ID *string `json:"id,omitempty" azure:"ro"`
}

// NetworkProfile - The network profile definition.
type NetworkProfile struct {
	// Specifies the list of resource Ids for the network interfaces associated with the dedicated HSM.
	NetworkInterfaces []*NetworkInterface `json:"networkInterfaces,omitempty"`

	// Specifies the identifier of the subnet.
	Subnet *APIEntityReference `json:"subnet,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// OutboundEnvironmentEndpoint - Egress endpoints which dedicated hsm service connects to for common purpose.
type OutboundEnvironmentEndpoint struct {
	// The category of endpoints accessed by the dedicated hsm service, e.g. azure-resource-management, apiserver, etc.
	Category *string `json:"category,omitempty"`

	// The endpoints that dedicated hsm service connects to
	Endpoints []*EndpointDependency `json:"endpoints,omitempty"`
}

// OutboundEnvironmentEndpointCollection - Collection of OutboundEnvironmentEndpoint
type OutboundEnvironmentEndpointCollection struct {
	// REQUIRED; Collection of resources.
	Value []*OutboundEnvironmentEndpoint `json:"value,omitempty"`

	// READ-ONLY; Link to next page of resources.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// Resource - Dedicated HSM resource
type Resource struct {
	// REQUIRED; The supported Azure location where the dedicated HSM should be created.
	Location *string `json:"location,omitempty"`

	// SKU details
	SKU *SKU `json:"sku,omitempty"`

	// Resource tags
	Tags map[string]*string `json:"tags,omitempty"`

	// The Dedicated Hsm zones.
	Zones []*string `json:"zones,omitempty"`

	// READ-ONLY; The Azure Resource Manager resource ID for the dedicated HSM.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the dedicated HSM.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The resource type of the dedicated HSM.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceListResult - List of dedicated HSM resources.
type ResourceListResult struct {
	// The URL to get the next set of dedicated HSM resources.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of dedicated HSM resources.
	Value []*Resource `json:"value,omitempty"`
}

// SKU of the dedicated HSM
type SKU struct {
	// SKU of the dedicated HSM
	Name *SKUName `json:"name,omitempty"`
}

// SystemData - Metadata pertaining to creation and last modification of dedicated hsm resource.
type SystemData struct {
	// The timestamp of dedicated hsm resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created dedicated hsm resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created dedicated hsm resource.
	CreatedByType *IdentityType `json:"createdByType,omitempty"`

	// The timestamp of dedicated hsm resource last modification (UTC).
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified dedicated hsm resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified dedicated hsm resource.
	LastModifiedByType *IdentityType `json:"lastModifiedByType,omitempty"`
}

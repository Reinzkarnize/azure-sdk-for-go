//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armproviderhub

// ClientCheckinManifestResponse contains the response from method Client.CheckinManifest.
type ClientCheckinManifestResponse struct {
	CheckinManifestInfo
}

// ClientGenerateManifestResponse contains the response from method Client.GenerateManifest.
type ClientGenerateManifestResponse struct {
	ResourceProviderManifest
}

// CustomRolloutsClientCreateOrUpdateResponse contains the response from method CustomRolloutsClient.CreateOrUpdate.
type CustomRolloutsClientCreateOrUpdateResponse struct {
	CustomRollout
}

// CustomRolloutsClientGetResponse contains the response from method CustomRolloutsClient.Get.
type CustomRolloutsClientGetResponse struct {
	CustomRollout
}

// CustomRolloutsClientListByProviderRegistrationResponse contains the response from method CustomRolloutsClient.ListByProviderRegistration.
type CustomRolloutsClientListByProviderRegistrationResponse struct {
	CustomRolloutArrayResponseWithContinuation
}

// DefaultRolloutsClientCreateOrUpdateResponse contains the response from method DefaultRolloutsClient.CreateOrUpdate.
type DefaultRolloutsClientCreateOrUpdateResponse struct {
	DefaultRollout
}

// DefaultRolloutsClientDeleteResponse contains the response from method DefaultRolloutsClient.Delete.
type DefaultRolloutsClientDeleteResponse struct {
	// placeholder for future response values
}

// DefaultRolloutsClientGetResponse contains the response from method DefaultRolloutsClient.Get.
type DefaultRolloutsClientGetResponse struct {
	DefaultRollout
}

// DefaultRolloutsClientListByProviderRegistrationResponse contains the response from method DefaultRolloutsClient.ListByProviderRegistration.
type DefaultRolloutsClientListByProviderRegistrationResponse struct {
	DefaultRolloutArrayResponseWithContinuation
}

// DefaultRolloutsClientStopResponse contains the response from method DefaultRolloutsClient.Stop.
type DefaultRolloutsClientStopResponse struct {
	// placeholder for future response values
}

// NotificationRegistrationsClientCreateOrUpdateResponse contains the response from method NotificationRegistrationsClient.CreateOrUpdate.
type NotificationRegistrationsClientCreateOrUpdateResponse struct {
	NotificationRegistration
}

// NotificationRegistrationsClientDeleteResponse contains the response from method NotificationRegistrationsClient.Delete.
type NotificationRegistrationsClientDeleteResponse struct {
	// placeholder for future response values
}

// NotificationRegistrationsClientGetResponse contains the response from method NotificationRegistrationsClient.Get.
type NotificationRegistrationsClientGetResponse struct {
	NotificationRegistration
}

// NotificationRegistrationsClientListByProviderRegistrationResponse contains the response from method NotificationRegistrationsClient.ListByProviderRegistration.
type NotificationRegistrationsClientListByProviderRegistrationResponse struct {
	NotificationRegistrationArrayResponseWithContinuation
}

// OperationsClientCreateOrUpdateResponse contains the response from method OperationsClient.CreateOrUpdate.
type OperationsClientCreateOrUpdateResponse struct {
	OperationsContent
}

// OperationsClientDeleteResponse contains the response from method OperationsClient.Delete.
type OperationsClientDeleteResponse struct {
	// placeholder for future response values
}

// OperationsClientListByProviderRegistrationResponse contains the response from method OperationsClient.ListByProviderRegistration.
type OperationsClientListByProviderRegistrationResponse struct {
	// Array of OperationsDefinition
	OperationsDefinitionArray []*OperationsDefinition
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsDefinitionArrayResponseWithContinuation
}

// ProviderRegistrationsClientCreateOrUpdateResponse contains the response from method ProviderRegistrationsClient.CreateOrUpdate.
type ProviderRegistrationsClientCreateOrUpdateResponse struct {
	ProviderRegistration
}

// ProviderRegistrationsClientDeleteResponse contains the response from method ProviderRegistrationsClient.Delete.
type ProviderRegistrationsClientDeleteResponse struct {
	// placeholder for future response values
}

// ProviderRegistrationsClientGenerateOperationsResponse contains the response from method ProviderRegistrationsClient.GenerateOperations.
type ProviderRegistrationsClientGenerateOperationsResponse struct {
	// Array of OperationsDefinition
	OperationsDefinitionArray []*OperationsDefinition
}

// ProviderRegistrationsClientGetResponse contains the response from method ProviderRegistrationsClient.Get.
type ProviderRegistrationsClientGetResponse struct {
	ProviderRegistration
}

// ProviderRegistrationsClientListResponse contains the response from method ProviderRegistrationsClient.List.
type ProviderRegistrationsClientListResponse struct {
	ProviderRegistrationArrayResponseWithContinuation
}

// ResourceTypeRegistrationsClientCreateOrUpdateResponse contains the response from method ResourceTypeRegistrationsClient.CreateOrUpdate.
type ResourceTypeRegistrationsClientCreateOrUpdateResponse struct {
	ResourceTypeRegistration
}

// ResourceTypeRegistrationsClientDeleteResponse contains the response from method ResourceTypeRegistrationsClient.Delete.
type ResourceTypeRegistrationsClientDeleteResponse struct {
	// placeholder for future response values
}

// ResourceTypeRegistrationsClientGetResponse contains the response from method ResourceTypeRegistrationsClient.Get.
type ResourceTypeRegistrationsClientGetResponse struct {
	ResourceTypeRegistration
}

// ResourceTypeRegistrationsClientListByProviderRegistrationResponse contains the response from method ResourceTypeRegistrationsClient.ListByProviderRegistration.
type ResourceTypeRegistrationsClientListByProviderRegistrationResponse struct {
	ResourceTypeRegistrationArrayResponseWithContinuation
}

// SKUsClientCreateOrUpdateNestedResourceTypeFirstResponse contains the response from method SKUsClient.CreateOrUpdateNestedResourceTypeFirst.
type SKUsClientCreateOrUpdateNestedResourceTypeFirstResponse struct {
	SKUResource
}

// SKUsClientCreateOrUpdateNestedResourceTypeSecondResponse contains the response from method SKUsClient.CreateOrUpdateNestedResourceTypeSecond.
type SKUsClientCreateOrUpdateNestedResourceTypeSecondResponse struct {
	SKUResource
}

// SKUsClientCreateOrUpdateNestedResourceTypeThirdResponse contains the response from method SKUsClient.CreateOrUpdateNestedResourceTypeThird.
type SKUsClientCreateOrUpdateNestedResourceTypeThirdResponse struct {
	SKUResource
}

// SKUsClientCreateOrUpdateResponse contains the response from method SKUsClient.CreateOrUpdate.
type SKUsClientCreateOrUpdateResponse struct {
	SKUResource
}

// SKUsClientDeleteNestedResourceTypeFirstResponse contains the response from method SKUsClient.DeleteNestedResourceTypeFirst.
type SKUsClientDeleteNestedResourceTypeFirstResponse struct {
	// placeholder for future response values
}

// SKUsClientDeleteNestedResourceTypeSecondResponse contains the response from method SKUsClient.DeleteNestedResourceTypeSecond.
type SKUsClientDeleteNestedResourceTypeSecondResponse struct {
	// placeholder for future response values
}

// SKUsClientDeleteNestedResourceTypeThirdResponse contains the response from method SKUsClient.DeleteNestedResourceTypeThird.
type SKUsClientDeleteNestedResourceTypeThirdResponse struct {
	// placeholder for future response values
}

// SKUsClientDeleteResponse contains the response from method SKUsClient.Delete.
type SKUsClientDeleteResponse struct {
	// placeholder for future response values
}

// SKUsClientGetNestedResourceTypeFirstResponse contains the response from method SKUsClient.GetNestedResourceTypeFirst.
type SKUsClientGetNestedResourceTypeFirstResponse struct {
	SKUResource
}

// SKUsClientGetNestedResourceTypeSecondResponse contains the response from method SKUsClient.GetNestedResourceTypeSecond.
type SKUsClientGetNestedResourceTypeSecondResponse struct {
	SKUResource
}

// SKUsClientGetNestedResourceTypeThirdResponse contains the response from method SKUsClient.GetNestedResourceTypeThird.
type SKUsClientGetNestedResourceTypeThirdResponse struct {
	SKUResource
}

// SKUsClientGetResponse contains the response from method SKUsClient.Get.
type SKUsClientGetResponse struct {
	SKUResource
}

// SKUsClientListByResourceTypeRegistrationsNestedResourceTypeFirstResponse contains the response from method SKUsClient.ListByResourceTypeRegistrationsNestedResourceTypeFirst.
type SKUsClientListByResourceTypeRegistrationsNestedResourceTypeFirstResponse struct {
	SKUResourceArrayResponseWithContinuation
}

// SKUsClientListByResourceTypeRegistrationsNestedResourceTypeSecondResponse contains the response from method SKUsClient.ListByResourceTypeRegistrationsNestedResourceTypeSecond.
type SKUsClientListByResourceTypeRegistrationsNestedResourceTypeSecondResponse struct {
	SKUResourceArrayResponseWithContinuation
}

// SKUsClientListByResourceTypeRegistrationsNestedResourceTypeThirdResponse contains the response from method SKUsClient.ListByResourceTypeRegistrationsNestedResourceTypeThird.
type SKUsClientListByResourceTypeRegistrationsNestedResourceTypeThirdResponse struct {
	SKUResourceArrayResponseWithContinuation
}

// SKUsClientListByResourceTypeRegistrationsResponse contains the response from method SKUsClient.ListByResourceTypeRegistrations.
type SKUsClientListByResourceTypeRegistrationsResponse struct {
	SKUResourceArrayResponseWithContinuation
}

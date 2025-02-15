//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

// EndpointsClientPurgeContentResponse contains the response from method EndpointsClient.PurgeContent.
type EndpointsClientPurgeContentResponse struct {
	// placeholder for future response values
}

// ExperimentsClientCreateOrUpdateResponse contains the response from method ExperimentsClient.CreateOrUpdate.
type ExperimentsClientCreateOrUpdateResponse struct {
	Experiment
}

// ExperimentsClientDeleteResponse contains the response from method ExperimentsClient.Delete.
type ExperimentsClientDeleteResponse struct {
	// placeholder for future response values
}

// ExperimentsClientGetResponse contains the response from method ExperimentsClient.Get.
type ExperimentsClientGetResponse struct {
	Experiment
}

// ExperimentsClientListByProfileResponse contains the response from method ExperimentsClient.ListByProfile.
type ExperimentsClientListByProfileResponse struct {
	ExperimentList
}

// ExperimentsClientUpdateResponse contains the response from method ExperimentsClient.Update.
type ExperimentsClientUpdateResponse struct {
	Experiment
}

// FrontDoorsClientCreateOrUpdateResponse contains the response from method FrontDoorsClient.CreateOrUpdate.
type FrontDoorsClientCreateOrUpdateResponse struct {
	FrontDoor
}

// FrontDoorsClientDeleteResponse contains the response from method FrontDoorsClient.Delete.
type FrontDoorsClientDeleteResponse struct {
	// placeholder for future response values
}

// FrontDoorsClientGetResponse contains the response from method FrontDoorsClient.Get.
type FrontDoorsClientGetResponse struct {
	FrontDoor
}

// FrontDoorsClientListByResourceGroupResponse contains the response from method FrontDoorsClient.ListByResourceGroup.
type FrontDoorsClientListByResourceGroupResponse struct {
	ListResult
}

// FrontDoorsClientListResponse contains the response from method FrontDoorsClient.List.
type FrontDoorsClientListResponse struct {
	ListResult
}

// FrontDoorsClientValidateCustomDomainResponse contains the response from method FrontDoorsClient.ValidateCustomDomain.
type FrontDoorsClientValidateCustomDomainResponse struct {
	ValidateCustomDomainOutput
}

// FrontendEndpointsClientDisableHTTPSResponse contains the response from method FrontendEndpointsClient.DisableHTTPS.
type FrontendEndpointsClientDisableHTTPSResponse struct {
	// placeholder for future response values
}

// FrontendEndpointsClientEnableHTTPSResponse contains the response from method FrontendEndpointsClient.EnableHTTPS.
type FrontendEndpointsClientEnableHTTPSResponse struct {
	// placeholder for future response values
}

// FrontendEndpointsClientGetResponse contains the response from method FrontendEndpointsClient.Get.
type FrontendEndpointsClientGetResponse struct {
	FrontendEndpoint
}

// FrontendEndpointsClientListByFrontDoorResponse contains the response from method FrontendEndpointsClient.ListByFrontDoor.
type FrontendEndpointsClientListByFrontDoorResponse struct {
	FrontendEndpointsListResult
}

// ManagedRuleSetsClientListResponse contains the response from method ManagedRuleSetsClient.List.
type ManagedRuleSetsClientListResponse struct {
	ManagedRuleSetDefinitionList
}

// NameAvailabilityClientCheckResponse contains the response from method NameAvailabilityClient.Check.
type NameAvailabilityClientCheckResponse struct {
	CheckNameAvailabilityOutput
}

// NameAvailabilityWithSubscriptionClientCheckResponse contains the response from method NameAvailabilityWithSubscriptionClient.Check.
type NameAvailabilityWithSubscriptionClientCheckResponse struct {
	CheckNameAvailabilityOutput
}

// NetworkExperimentProfilesClientCreateOrUpdateResponse contains the response from method NetworkExperimentProfilesClient.CreateOrUpdate.
type NetworkExperimentProfilesClientCreateOrUpdateResponse struct {
	Profile
}

// NetworkExperimentProfilesClientDeleteResponse contains the response from method NetworkExperimentProfilesClient.Delete.
type NetworkExperimentProfilesClientDeleteResponse struct {
	// placeholder for future response values
}

// NetworkExperimentProfilesClientGetResponse contains the response from method NetworkExperimentProfilesClient.Get.
type NetworkExperimentProfilesClientGetResponse struct {
	Profile
}

// NetworkExperimentProfilesClientListByResourceGroupResponse contains the response from method NetworkExperimentProfilesClient.ListByResourceGroup.
type NetworkExperimentProfilesClientListByResourceGroupResponse struct {
	ProfileList
}

// NetworkExperimentProfilesClientListResponse contains the response from method NetworkExperimentProfilesClient.List.
type NetworkExperimentProfilesClientListResponse struct {
	ProfileList
}

// NetworkExperimentProfilesClientUpdateResponse contains the response from method NetworkExperimentProfilesClient.Update.
type NetworkExperimentProfilesClientUpdateResponse struct {
	Profile
}

// PoliciesClientCreateOrUpdateResponse contains the response from method PoliciesClient.CreateOrUpdate.
type PoliciesClientCreateOrUpdateResponse struct {
	WebApplicationFirewallPolicy
}

// PoliciesClientDeleteResponse contains the response from method PoliciesClient.Delete.
type PoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// PoliciesClientGetResponse contains the response from method PoliciesClient.Get.
type PoliciesClientGetResponse struct {
	WebApplicationFirewallPolicy
}

// PoliciesClientListResponse contains the response from method PoliciesClient.List.
type PoliciesClientListResponse struct {
	WebApplicationFirewallPolicyList
}

// PreconfiguredEndpointsClientListResponse contains the response from method PreconfiguredEndpointsClient.List.
type PreconfiguredEndpointsClientListResponse struct {
	PreconfiguredEndpointList
}

// ReportsClientGetLatencyScorecardsResponse contains the response from method ReportsClient.GetLatencyScorecards.
type ReportsClientGetLatencyScorecardsResponse struct {
	LatencyScorecard
}

// ReportsClientGetTimeseriesResponse contains the response from method ReportsClient.GetTimeseries.
type ReportsClientGetTimeseriesResponse struct {
	Timeseries
}

// RulesEnginesClientCreateOrUpdateResponse contains the response from method RulesEnginesClient.CreateOrUpdate.
type RulesEnginesClientCreateOrUpdateResponse struct {
	RulesEngine
}

// RulesEnginesClientDeleteResponse contains the response from method RulesEnginesClient.Delete.
type RulesEnginesClientDeleteResponse struct {
	// placeholder for future response values
}

// RulesEnginesClientGetResponse contains the response from method RulesEnginesClient.Get.
type RulesEnginesClientGetResponse struct {
	RulesEngine
}

// RulesEnginesClientListByFrontDoorResponse contains the response from method RulesEnginesClient.ListByFrontDoor.
type RulesEnginesClientListByFrontDoorResponse struct {
	RulesEngineListResult
}

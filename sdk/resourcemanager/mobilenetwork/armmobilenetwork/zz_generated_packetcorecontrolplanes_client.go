//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilenetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PacketCoreControlPlanesClient contains the methods for the PacketCoreControlPlanes group.
// Don't use this type directly, use NewPacketCoreControlPlanesClient() instead.
type PacketCoreControlPlanesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPacketCoreControlPlanesClient creates a new instance of PacketCoreControlPlanesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPacketCoreControlPlanesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PacketCoreControlPlanesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &PacketCoreControlPlanesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a PacketCoreControlPlane.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// packetCoreControlPlaneName - The name of the packet core control plane.
// parameters - Parameters supplied to the create or update packet core control plane operation.
// options - PacketCoreControlPlanesClientBeginCreateOrUpdateOptions contains the optional parameters for the PacketCoreControlPlanesClient.BeginCreateOrUpdate
// method.
func (client *PacketCoreControlPlanesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, parameters PacketCoreControlPlane, options *PacketCoreControlPlanesClientBeginCreateOrUpdateOptions) (PacketCoreControlPlanesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, packetCoreControlPlaneName, parameters, options)
	if err != nil {
		return PacketCoreControlPlanesClientCreateOrUpdatePollerResponse{}, err
	}
	result := PacketCoreControlPlanesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PacketCoreControlPlanesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return PacketCoreControlPlanesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &PacketCoreControlPlanesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a PacketCoreControlPlane.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PacketCoreControlPlanesClient) createOrUpdate(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, parameters PacketCoreControlPlane, options *PacketCoreControlPlanesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PacketCoreControlPlanesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, parameters PacketCoreControlPlane, options *PacketCoreControlPlanesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified packet core control plane.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// packetCoreControlPlaneName - The name of the packet core control plane.
// options - PacketCoreControlPlanesClientBeginDeleteOptions contains the optional parameters for the PacketCoreControlPlanesClient.BeginDelete
// method.
func (client *PacketCoreControlPlanesClient) BeginDelete(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, options *PacketCoreControlPlanesClientBeginDeleteOptions) (PacketCoreControlPlanesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, packetCoreControlPlaneName, options)
	if err != nil {
		return PacketCoreControlPlanesClientDeletePollerResponse{}, err
	}
	result := PacketCoreControlPlanesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PacketCoreControlPlanesClient.Delete", "location", resp, client.pl)
	if err != nil {
		return PacketCoreControlPlanesClientDeletePollerResponse{}, err
	}
	result.Poller = &PacketCoreControlPlanesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified packet core control plane.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PacketCoreControlPlanesClient) deleteOperation(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, options *PacketCoreControlPlanesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PacketCoreControlPlanesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, options *PacketCoreControlPlanesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets information about the specified packet core control plane.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// packetCoreControlPlaneName - The name of the packet core control plane.
// options - PacketCoreControlPlanesClientGetOptions contains the optional parameters for the PacketCoreControlPlanesClient.Get
// method.
func (client *PacketCoreControlPlanesClient) Get(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, options *PacketCoreControlPlanesClientGetOptions) (PacketCoreControlPlanesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, options)
	if err != nil {
		return PacketCoreControlPlanesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PacketCoreControlPlanesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PacketCoreControlPlanesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PacketCoreControlPlanesClient) getCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, options *PacketCoreControlPlanesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PacketCoreControlPlanesClient) getHandleResponse(resp *http.Response) (PacketCoreControlPlanesClientGetResponse, error) {
	result := PacketCoreControlPlanesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PacketCoreControlPlane); err != nil {
		return PacketCoreControlPlanesClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists all the packetCoreControlPlanes in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - PacketCoreControlPlanesClientListByResourceGroupOptions contains the optional parameters for the PacketCoreControlPlanesClient.ListByResourceGroup
// method.
func (client *PacketCoreControlPlanesClient) ListByResourceGroup(resourceGroupName string, options *PacketCoreControlPlanesClientListByResourceGroupOptions) *PacketCoreControlPlanesClientListByResourceGroupPager {
	return &PacketCoreControlPlanesClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp PacketCoreControlPlanesClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PacketCoreControlPlaneListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PacketCoreControlPlanesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PacketCoreControlPlanesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PacketCoreControlPlanesClient) listByResourceGroupHandleResponse(resp *http.Response) (PacketCoreControlPlanesClientListByResourceGroupResponse, error) {
	result := PacketCoreControlPlanesClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PacketCoreControlPlaneListResult); err != nil {
		return PacketCoreControlPlanesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Lists all the packetCoreControlPlanes in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - PacketCoreControlPlanesClientListBySubscriptionOptions contains the optional parameters for the PacketCoreControlPlanesClient.ListBySubscription
// method.
func (client *PacketCoreControlPlanesClient) ListBySubscription(options *PacketCoreControlPlanesClientListBySubscriptionOptions) *PacketCoreControlPlanesClientListBySubscriptionPager {
	return &PacketCoreControlPlanesClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp PacketCoreControlPlanesClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PacketCoreControlPlaneListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *PacketCoreControlPlanesClient) listBySubscriptionCreateRequest(ctx context.Context, options *PacketCoreControlPlanesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *PacketCoreControlPlanesClient) listBySubscriptionHandleResponse(resp *http.Response) (PacketCoreControlPlanesClientListBySubscriptionResponse, error) {
	result := PacketCoreControlPlanesClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PacketCoreControlPlaneListResult); err != nil {
		return PacketCoreControlPlanesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates a PacketCoreControlPlane update tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// packetCoreControlPlaneName - The name of the packet core control plane.
// parameters - Parameters supplied to update PacketCoreControlPlane tags.
// options - PacketCoreControlPlanesClientUpdateTagsOptions contains the optional parameters for the PacketCoreControlPlanesClient.UpdateTags
// method.
func (client *PacketCoreControlPlanesClient) UpdateTags(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, parameters TagsObject, options *PacketCoreControlPlanesClientUpdateTagsOptions) (PacketCoreControlPlanesClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, parameters, options)
	if err != nil {
		return PacketCoreControlPlanesClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PacketCoreControlPlanesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PacketCoreControlPlanesClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *PacketCoreControlPlanesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, parameters TagsObject, options *PacketCoreControlPlanesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *PacketCoreControlPlanesClient) updateTagsHandleResponse(resp *http.Response) (PacketCoreControlPlanesClientUpdateTagsResponse, error) {
	result := PacketCoreControlPlanesClientUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PacketCoreControlPlane); err != nil {
		return PacketCoreControlPlanesClientUpdateTagsResponse{}, err
	}
	return result, nil
}

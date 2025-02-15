//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package armconnectedvmware

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// HybridIdentityMetadataClient contains the methods for the HybridIdentityMetadata group.
// Don't use this type directly, use NewHybridIdentityMetadataClient() instead.
type HybridIdentityMetadataClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewHybridIdentityMetadataClient creates a new instance of HybridIdentityMetadataClient with the specified values.
// subscriptionID - The Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewHybridIdentityMetadataClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HybridIdentityMetadataClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &HybridIdentityMetadataClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Create Or Update HybridIdentityMetadata.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-10-preview
// resourceGroupName - The Resource Group Name.
// virtualMachineName - Name of the vm.
// metadataName - Name of the hybridIdentityMetadata.
// body - Request payload.
// options - HybridIdentityMetadataClientCreateOptions contains the optional parameters for the HybridIdentityMetadataClient.Create
// method.
func (client *HybridIdentityMetadataClient) Create(ctx context.Context, resourceGroupName string, virtualMachineName string, metadataName string, body HybridIdentityMetadata, options *HybridIdentityMetadataClientCreateOptions) (HybridIdentityMetadataClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, virtualMachineName, metadataName, body, options)
	if err != nil {
		return HybridIdentityMetadataClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridIdentityMetadataClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridIdentityMetadataClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *HybridIdentityMetadataClient) createCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, metadataName string, body HybridIdentityMetadata, options *HybridIdentityMetadataClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/{virtualMachineName}/hybridIdentityMetadata/{metadataName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	if metadataName == "" {
		return nil, errors.New("parameter metadataName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metadataName}", url.PathEscape(metadataName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// createHandleResponse handles the Create response.
func (client *HybridIdentityMetadataClient) createHandleResponse(resp *http.Response) (HybridIdentityMetadataClientCreateResponse, error) {
	result := HybridIdentityMetadataClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridIdentityMetadata); err != nil {
		return HybridIdentityMetadataClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Implements HybridIdentityMetadata DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-10-preview
// resourceGroupName - The Resource Group Name.
// virtualMachineName - Name of the vm.
// metadataName - Name of the HybridIdentityMetadata.
// options - HybridIdentityMetadataClientDeleteOptions contains the optional parameters for the HybridIdentityMetadataClient.Delete
// method.
func (client *HybridIdentityMetadataClient) Delete(ctx context.Context, resourceGroupName string, virtualMachineName string, metadataName string, options *HybridIdentityMetadataClientDeleteOptions) (HybridIdentityMetadataClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualMachineName, metadataName, options)
	if err != nil {
		return HybridIdentityMetadataClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridIdentityMetadataClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return HybridIdentityMetadataClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return HybridIdentityMetadataClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HybridIdentityMetadataClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, metadataName string, options *HybridIdentityMetadataClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/{virtualMachineName}/hybridIdentityMetadata/{metadataName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	if metadataName == "" {
		return nil, errors.New("parameter metadataName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metadataName}", url.PathEscape(metadataName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Implements HybridIdentityMetadata GET method.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-10-preview
// resourceGroupName - The Resource Group Name.
// virtualMachineName - Name of the vm.
// metadataName - Name of the HybridIdentityMetadata.
// options - HybridIdentityMetadataClientGetOptions contains the optional parameters for the HybridIdentityMetadataClient.Get
// method.
func (client *HybridIdentityMetadataClient) Get(ctx context.Context, resourceGroupName string, virtualMachineName string, metadataName string, options *HybridIdentityMetadataClientGetOptions) (HybridIdentityMetadataClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualMachineName, metadataName, options)
	if err != nil {
		return HybridIdentityMetadataClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridIdentityMetadataClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridIdentityMetadataClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HybridIdentityMetadataClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, metadataName string, options *HybridIdentityMetadataClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/{virtualMachineName}/hybridIdentityMetadata/{metadataName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	if metadataName == "" {
		return nil, errors.New("parameter metadataName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metadataName}", url.PathEscape(metadataName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HybridIdentityMetadataClient) getHandleResponse(resp *http.Response) (HybridIdentityMetadataClientGetResponse, error) {
	result := HybridIdentityMetadataClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridIdentityMetadata); err != nil {
		return HybridIdentityMetadataClientGetResponse{}, err
	}
	return result, nil
}

// NewListByVMPager - Returns the list of HybridIdentityMetadata of the given vm.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-10-preview
// resourceGroupName - The Resource Group Name.
// virtualMachineName - Name of the vm.
// options - HybridIdentityMetadataClientListByVMOptions contains the optional parameters for the HybridIdentityMetadataClient.ListByVM
// method.
func (client *HybridIdentityMetadataClient) NewListByVMPager(resourceGroupName string, virtualMachineName string, options *HybridIdentityMetadataClientListByVMOptions) *runtime.Pager[HybridIdentityMetadataClientListByVMResponse] {
	return runtime.NewPager(runtime.PagingHandler[HybridIdentityMetadataClientListByVMResponse]{
		More: func(page HybridIdentityMetadataClientListByVMResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HybridIdentityMetadataClientListByVMResponse) (HybridIdentityMetadataClientListByVMResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByVMCreateRequest(ctx, resourceGroupName, virtualMachineName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HybridIdentityMetadataClientListByVMResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return HybridIdentityMetadataClientListByVMResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HybridIdentityMetadataClientListByVMResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByVMHandleResponse(resp)
		},
	})
}

// listByVMCreateRequest creates the ListByVM request.
func (client *HybridIdentityMetadataClient) listByVMCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, options *HybridIdentityMetadataClientListByVMOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/{virtualMachineName}/hybridIdentityMetadata"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByVMHandleResponse handles the ListByVM response.
func (client *HybridIdentityMetadataClient) listByVMHandleResponse(resp *http.Response) (HybridIdentityMetadataClientListByVMResponse, error) {
	result := HybridIdentityMetadataClientListByVMResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridIdentityMetadataList); err != nil {
		return HybridIdentityMetadataClientListByVMResponse{}, err
	}
	return result, nil
}

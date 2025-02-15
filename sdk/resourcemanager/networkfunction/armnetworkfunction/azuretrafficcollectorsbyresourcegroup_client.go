//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetworkfunction

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

// AzureTrafficCollectorsByResourceGroupClient contains the methods for the AzureTrafficCollectorsByResourceGroup group.
// Don't use this type directly, use NewAzureTrafficCollectorsByResourceGroupClient() instead.
type AzureTrafficCollectorsByResourceGroupClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAzureTrafficCollectorsByResourceGroupClient creates a new instance of AzureTrafficCollectorsByResourceGroupClient with the specified values.
// subscriptionID - Azure Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAzureTrafficCollectorsByResourceGroupClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AzureTrafficCollectorsByResourceGroupClient, error) {
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
	client := &AzureTrafficCollectorsByResourceGroupClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Return list of Azure Traffic Collectors in a Resource Group
// Generated from API version 2022-11-01
// resourceGroupName - The name of the resource group.
// options - AzureTrafficCollectorsByResourceGroupClientListOptions contains the optional parameters for the AzureTrafficCollectorsByResourceGroupClient.List
// method.
func (client *AzureTrafficCollectorsByResourceGroupClient) NewListPager(resourceGroupName string, options *AzureTrafficCollectorsByResourceGroupClientListOptions) *runtime.Pager[AzureTrafficCollectorsByResourceGroupClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureTrafficCollectorsByResourceGroupClientListResponse]{
		More: func(page AzureTrafficCollectorsByResourceGroupClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureTrafficCollectorsByResourceGroupClientListResponse) (AzureTrafficCollectorsByResourceGroupClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AzureTrafficCollectorsByResourceGroupClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AzureTrafficCollectorsByResourceGroupClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AzureTrafficCollectorsByResourceGroupClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AzureTrafficCollectorsByResourceGroupClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *AzureTrafficCollectorsByResourceGroupClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkFunction/azureTrafficCollectors"
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
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AzureTrafficCollectorsByResourceGroupClient) listHandleResponse(resp *http.Response) (AzureTrafficCollectorsByResourceGroupClientListResponse, error) {
	result := AzureTrafficCollectorsByResourceGroupClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureTrafficCollectorListResult); err != nil {
		return AzureTrafficCollectorsByResourceGroupClientListResponse{}, err
	}
	return result, nil
}

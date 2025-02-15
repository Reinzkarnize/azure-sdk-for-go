//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridcontainerservice

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// Client contains the methods for the HybridContainerService group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	host string
	pl   runtime.Pipeline
}

// NewClient creates a new instance of Client with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
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
	client := &Client{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// ListOrchestrators - Lists the available orchestrators in a custom location for HybridAKS
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// customLocationResourceURI - The fully qualified Azure Resource manager identifier of the custom location resource.
// options - ClientListOrchestratorsOptions contains the optional parameters for the Client.ListOrchestrators method.
func (client *Client) ListOrchestrators(ctx context.Context, customLocationResourceURI string, options *ClientListOrchestratorsOptions) (ClientListOrchestratorsResponse, error) {
	req, err := client.listOrchestratorsCreateRequest(ctx, customLocationResourceURI, options)
	if err != nil {
		return ClientListOrchestratorsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientListOrchestratorsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientListOrchestratorsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listOrchestratorsHandleResponse(resp)
}

// listOrchestratorsCreateRequest creates the ListOrchestrators request.
func (client *Client) listOrchestratorsCreateRequest(ctx context.Context, customLocationResourceURI string, options *ClientListOrchestratorsOptions) (*policy.Request, error) {
	urlPath := "/{customLocationResourceUri}/providers/Microsoft.HybridContainerService/orchestrators"
	urlPath = strings.ReplaceAll(urlPath, "{customLocationResourceUri}", customLocationResourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listOrchestratorsHandleResponse handles the ListOrchestrators response.
func (client *Client) listOrchestratorsHandleResponse(resp *http.Response) (ClientListOrchestratorsResponse, error) {
	result := ClientListOrchestratorsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrchestratorVersionProfileListResult); err != nil {
		return ClientListOrchestratorsResponse{}, err
	}
	return result, nil
}

// ListVMSKUs - Lists the available VM SKUs in a custom location for HybridAKS
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// customLocationResourceURI - The fully qualified Azure Resource manager identifier of the custom location resource.
// options - ClientListVMSKUsOptions contains the optional parameters for the Client.ListVMSKUs method.
func (client *Client) ListVMSKUs(ctx context.Context, customLocationResourceURI string, options *ClientListVMSKUsOptions) (ClientListVMSKUsResponse, error) {
	req, err := client.listVMSKUsCreateRequest(ctx, customLocationResourceURI, options)
	if err != nil {
		return ClientListVMSKUsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientListVMSKUsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientListVMSKUsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listVMSKUsHandleResponse(resp)
}

// listVMSKUsCreateRequest creates the ListVMSKUs request.
func (client *Client) listVMSKUsCreateRequest(ctx context.Context, customLocationResourceURI string, options *ClientListVMSKUsOptions) (*policy.Request, error) {
	urlPath := "/{customLocationResourceUri}/providers/Microsoft.HybridContainerService/vmSkus"
	urlPath = strings.ReplaceAll(urlPath, "{customLocationResourceUri}", customLocationResourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listVMSKUsHandleResponse handles the ListVMSKUs response.
func (client *Client) listVMSKUsHandleResponse(resp *http.Response) (ClientListVMSKUsResponse, error) {
	result := ClientListVMSKUsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VMSKUListResult); err != nil {
		return ClientListVMSKUsResponse{}, err
	}
	return result, nil
}

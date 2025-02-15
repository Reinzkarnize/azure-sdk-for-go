//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

// FilesClientCreateOrUpdateResponse contains the response from method FilesClient.CreateOrUpdate.
type FilesClientCreateOrUpdateResponse struct {
	ProjectFile
}

// FilesClientDeleteResponse contains the response from method FilesClient.Delete.
type FilesClientDeleteResponse struct {
	// placeholder for future response values
}

// FilesClientGetResponse contains the response from method FilesClient.Get.
type FilesClientGetResponse struct {
	ProjectFile
}

// FilesClientListResponse contains the response from method FilesClient.List.
type FilesClientListResponse struct {
	FileList
}

// FilesClientReadResponse contains the response from method FilesClient.Read.
type FilesClientReadResponse struct {
	FileStorageInfo
}

// FilesClientReadWriteResponse contains the response from method FilesClient.ReadWrite.
type FilesClientReadWriteResponse struct {
	FileStorageInfo
}

// FilesClientUpdateResponse contains the response from method FilesClient.Update.
type FilesClientUpdateResponse struct {
	ProjectFile
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	ServiceOperationList
}

// ProjectsClientCreateOrUpdateResponse contains the response from method ProjectsClient.CreateOrUpdate.
type ProjectsClientCreateOrUpdateResponse struct {
	Project
}

// ProjectsClientDeleteResponse contains the response from method ProjectsClient.Delete.
type ProjectsClientDeleteResponse struct {
	// placeholder for future response values
}

// ProjectsClientGetResponse contains the response from method ProjectsClient.Get.
type ProjectsClientGetResponse struct {
	Project
}

// ProjectsClientListResponse contains the response from method ProjectsClient.List.
type ProjectsClientListResponse struct {
	ProjectList
}

// ProjectsClientUpdateResponse contains the response from method ProjectsClient.Update.
type ProjectsClientUpdateResponse struct {
	Project
}

// ResourceSKUsClientListSKUsResponse contains the response from method ResourceSKUsClient.ListSKUs.
type ResourceSKUsClientListSKUsResponse struct {
	ResourceSKUsResult
}

// ServiceTasksClientCancelResponse contains the response from method ServiceTasksClient.Cancel.
type ServiceTasksClientCancelResponse struct {
	ProjectTask
}

// ServiceTasksClientCreateOrUpdateResponse contains the response from method ServiceTasksClient.CreateOrUpdate.
type ServiceTasksClientCreateOrUpdateResponse struct {
	ProjectTask
}

// ServiceTasksClientDeleteResponse contains the response from method ServiceTasksClient.Delete.
type ServiceTasksClientDeleteResponse struct {
	// placeholder for future response values
}

// ServiceTasksClientGetResponse contains the response from method ServiceTasksClient.Get.
type ServiceTasksClientGetResponse struct {
	ProjectTask
}

// ServiceTasksClientListResponse contains the response from method ServiceTasksClient.List.
type ServiceTasksClientListResponse struct {
	TaskList
}

// ServiceTasksClientUpdateResponse contains the response from method ServiceTasksClient.Update.
type ServiceTasksClientUpdateResponse struct {
	ProjectTask
}

// ServicesClientCheckChildrenNameAvailabilityResponse contains the response from method ServicesClient.CheckChildrenNameAvailability.
type ServicesClientCheckChildrenNameAvailabilityResponse struct {
	NameAvailabilityResponse
}

// ServicesClientCheckNameAvailabilityResponse contains the response from method ServicesClient.CheckNameAvailability.
type ServicesClientCheckNameAvailabilityResponse struct {
	NameAvailabilityResponse
}

// ServicesClientCheckStatusResponse contains the response from method ServicesClient.CheckStatus.
type ServicesClientCheckStatusResponse struct {
	ServiceStatusResponse
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.CreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	Service
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.Delete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	Service
}

// ServicesClientListByResourceGroupResponse contains the response from method ServicesClient.ListByResourceGroup.
type ServicesClientListByResourceGroupResponse struct {
	ServiceList
}

// ServicesClientListResponse contains the response from method ServicesClient.List.
type ServicesClientListResponse struct {
	ServiceList
}

// ServicesClientListSKUsResponse contains the response from method ServicesClient.ListSKUs.
type ServicesClientListSKUsResponse struct {
	ServiceSKUList
}

// ServicesClientStartResponse contains the response from method ServicesClient.Start.
type ServicesClientStartResponse struct {
	// placeholder for future response values
}

// ServicesClientStopResponse contains the response from method ServicesClient.Stop.
type ServicesClientStopResponse struct {
	// placeholder for future response values
}

// ServicesClientUpdateResponse contains the response from method ServicesClient.Update.
type ServicesClientUpdateResponse struct {
	Service
}

// TasksClientCancelResponse contains the response from method TasksClient.Cancel.
type TasksClientCancelResponse struct {
	ProjectTask
}

// TasksClientCommandResponse contains the response from method TasksClient.Command.
type TasksClientCommandResponse struct {
	CommandPropertiesClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TasksClientCommandResponse.
func (t *TasksClientCommandResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalCommandPropertiesClassification(data)
	if err != nil {
		return err
	}
	t.CommandPropertiesClassification = res
	return nil
}

// TasksClientCreateOrUpdateResponse contains the response from method TasksClient.CreateOrUpdate.
type TasksClientCreateOrUpdateResponse struct {
	ProjectTask
}

// TasksClientDeleteResponse contains the response from method TasksClient.Delete.
type TasksClientDeleteResponse struct {
	// placeholder for future response values
}

// TasksClientGetResponse contains the response from method TasksClient.Get.
type TasksClientGetResponse struct {
	ProjectTask
}

// TasksClientListResponse contains the response from method TasksClient.List.
type TasksClientListResponse struct {
	TaskList
}

// TasksClientUpdateResponse contains the response from method TasksClient.Update.
type TasksClientUpdateResponse struct {
	ProjectTask
}

// UsagesClientListResponse contains the response from method UsagesClient.List.
type UsagesClientListResponse struct {
	QuotaList
}

// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package servicebus

import original "github.com/Azure/azure-sdk-for-go/services/servicebus/mgmt/2017-04-01/servicebus"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type DisasterRecoveryConfigsClient = original.DisasterRecoveryConfigsClient
type EventHubsClient = original.EventHubsClient
type MigrationConfigsClient = original.MigrationConfigsClient
type AccessRights = original.AccessRights

const (
	Listen AccessRights = original.Listen
	Manage AccessRights = original.Manage
	Send   AccessRights = original.Send
)

type EncodingCaptureDescription = original.EncodingCaptureDescription

const (
	Avro        EncodingCaptureDescription = original.Avro
	AvroDeflate EncodingCaptureDescription = original.AvroDeflate
)

type EntityStatus = original.EntityStatus

const (
	Active          EntityStatus = original.Active
	Creating        EntityStatus = original.Creating
	Deleting        EntityStatus = original.Deleting
	Disabled        EntityStatus = original.Disabled
	ReceiveDisabled EntityStatus = original.ReceiveDisabled
	Renaming        EntityStatus = original.Renaming
	Restoring       EntityStatus = original.Restoring
	SendDisabled    EntityStatus = original.SendDisabled
	Unknown         EntityStatus = original.Unknown
)

type FilterType = original.FilterType

const (
	FilterTypeCorrelationFilter FilterType = original.FilterTypeCorrelationFilter
	FilterTypeSQLFilter         FilterType = original.FilterTypeSQLFilter
)

type IPAction = original.IPAction

const (
	Accept IPAction = original.Accept
	Reject IPAction = original.Reject
)

type KeyType = original.KeyType

const (
	PrimaryKey   KeyType = original.PrimaryKey
	SecondaryKey KeyType = original.SecondaryKey
)

type NameSpaceType = original.NameSpaceType

const (
	EventHub        NameSpaceType = original.EventHub
	Messaging       NameSpaceType = original.Messaging
	Mixed           NameSpaceType = original.Mixed
	NotificationHub NameSpaceType = original.NotificationHub
	Relay           NameSpaceType = original.Relay
)

type ProvisioningStateDR = original.ProvisioningStateDR

const (
	Accepted  ProvisioningStateDR = original.Accepted
	Failed    ProvisioningStateDR = original.Failed
	Succeeded ProvisioningStateDR = original.Succeeded
)

type RoleDisasterRecovery = original.RoleDisasterRecovery

const (
	Primary               RoleDisasterRecovery = original.Primary
	PrimaryNotReplicating RoleDisasterRecovery = original.PrimaryNotReplicating
	Secondary             RoleDisasterRecovery = original.Secondary
)

type SkuName = original.SkuName

const (
	Basic    SkuName = original.Basic
	Premium  SkuName = original.Premium
	Standard SkuName = original.Standard
)

type SkuTier = original.SkuTier

const (
	SkuTierBasic    SkuTier = original.SkuTierBasic
	SkuTierPremium  SkuTier = original.SkuTierPremium
	SkuTierStandard SkuTier = original.SkuTierStandard
)

type UnavailableReason = original.UnavailableReason

const (
	InvalidName                           UnavailableReason = original.InvalidName
	NameInLockdown                        UnavailableReason = original.NameInLockdown
	NameInUse                             UnavailableReason = original.NameInUse
	None                                  UnavailableReason = original.None
	SubscriptionIsDisabled                UnavailableReason = original.SubscriptionIsDisabled
	TooManyNamespaceInCurrentSubscription UnavailableReason = original.TooManyNamespaceInCurrentSubscription
)

type AccessKeys = original.AccessKeys
type Action = original.Action
type ArmDisasterRecovery = original.ArmDisasterRecovery
type ArmDisasterRecoveryListResult = original.ArmDisasterRecoveryListResult
type ArmDisasterRecoveryListResultIterator = original.ArmDisasterRecoveryListResultIterator
type ArmDisasterRecoveryListResultPage = original.ArmDisasterRecoveryListResultPage
type ArmDisasterRecoveryProperties = original.ArmDisasterRecoveryProperties
type AuthorizationRuleProperties = original.AuthorizationRuleProperties
type CaptureDescription = original.CaptureDescription
type CheckNameAvailability = original.CheckNameAvailability
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type CorrelationFilter = original.CorrelationFilter
type Destination = original.Destination
type DestinationProperties = original.DestinationProperties
type ErrorResponse = original.ErrorResponse
type Eventhub = original.Eventhub
type EventHubListResult = original.EventHubListResult
type EventHubListResultIterator = original.EventHubListResultIterator
type EventHubListResultPage = original.EventHubListResultPage
type EventhubProperties = original.EventhubProperties
type IPFilterRule = original.IPFilterRule
type IPFilterRuleListResult = original.IPFilterRuleListResult
type IPFilterRuleListResultIterator = original.IPFilterRuleListResultIterator
type IPFilterRuleListResultPage = original.IPFilterRuleListResultPage
type IPFilterRuleProperties = original.IPFilterRuleProperties
type MessageCountDetails = original.MessageCountDetails
type MigrationConfigListResult = original.MigrationConfigListResult
type MigrationConfigListResultIterator = original.MigrationConfigListResultIterator
type MigrationConfigListResultPage = original.MigrationConfigListResultPage
type MigrationConfigProperties = original.MigrationConfigProperties
type MigrationConfigPropertiesProperties = original.MigrationConfigPropertiesProperties
type MigrationConfigsCreateAndStartMigrationFuture = original.MigrationConfigsCreateAndStartMigrationFuture
type NamespacesCreateOrUpdateFuture = original.NamespacesCreateOrUpdateFuture
type NamespacesDeleteFuture = original.NamespacesDeleteFuture
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type PremiumMessagingRegions = original.PremiumMessagingRegions
type PremiumMessagingRegionsListResult = original.PremiumMessagingRegionsListResult
type PremiumMessagingRegionsListResultIterator = original.PremiumMessagingRegionsListResultIterator
type PremiumMessagingRegionsListResultPage = original.PremiumMessagingRegionsListResultPage
type PremiumMessagingRegionsProperties = original.PremiumMessagingRegionsProperties
type RegenerateAccessKeyParameters = original.RegenerateAccessKeyParameters
type Resource = original.Resource
type ResourceNamespacePatch = original.ResourceNamespacePatch
type Rule = original.Rule
type RuleListResult = original.RuleListResult
type RuleListResultIterator = original.RuleListResultIterator
type RuleListResultPage = original.RuleListResultPage
type Ruleproperties = original.Ruleproperties
type SBAuthorizationRule = original.SBAuthorizationRule
type SBAuthorizationRuleListResult = original.SBAuthorizationRuleListResult
type SBAuthorizationRuleListResultIterator = original.SBAuthorizationRuleListResultIterator
type SBAuthorizationRuleListResultPage = original.SBAuthorizationRuleListResultPage
type SBAuthorizationRuleProperties = original.SBAuthorizationRuleProperties
type SBNamespace = original.SBNamespace
type SBNamespaceListResult = original.SBNamespaceListResult
type SBNamespaceListResultIterator = original.SBNamespaceListResultIterator
type SBNamespaceListResultPage = original.SBNamespaceListResultPage
type SBNamespaceMigrate = original.SBNamespaceMigrate
type SBNamespaceProperties = original.SBNamespaceProperties
type SBNamespaceUpdateParameters = original.SBNamespaceUpdateParameters
type SBQueue = original.SBQueue
type SBQueueListResult = original.SBQueueListResult
type SBQueueListResultIterator = original.SBQueueListResultIterator
type SBQueueListResultPage = original.SBQueueListResultPage
type SBQueueProperties = original.SBQueueProperties
type SBSku = original.SBSku
type SBSubscription = original.SBSubscription
type SBSubscriptionListResult = original.SBSubscriptionListResult
type SBSubscriptionListResultIterator = original.SBSubscriptionListResultIterator
type SBSubscriptionListResultPage = original.SBSubscriptionListResultPage
type SBSubscriptionProperties = original.SBSubscriptionProperties
type SBTopic = original.SBTopic
type SBTopicListResult = original.SBTopicListResult
type SBTopicListResultIterator = original.SBTopicListResultIterator
type SBTopicListResultPage = original.SBTopicListResultPage
type SBTopicProperties = original.SBTopicProperties
type SQLFilter = original.SQLFilter
type SQLRuleAction = original.SQLRuleAction
type TrackedResource = original.TrackedResource
type VirtualNetworkRule = original.VirtualNetworkRule
type VirtualNetworkRuleListResult = original.VirtualNetworkRuleListResult
type VirtualNetworkRuleListResultIterator = original.VirtualNetworkRuleListResultIterator
type VirtualNetworkRuleListResultPage = original.VirtualNetworkRuleListResultPage
type VirtualNetworkRuleProperties = original.VirtualNetworkRuleProperties
type NamespacesClient = original.NamespacesClient
type OperationsClient = original.OperationsClient
type PremiumMessagingRegionsClient = original.PremiumMessagingRegionsClient
type QueuesClient = original.QueuesClient
type RegionsClient = original.RegionsClient
type RulesClient = original.RulesClient
type SubscriptionsClient = original.SubscriptionsClient
type TopicsClient = original.TopicsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewDisasterRecoveryConfigsClient(subscriptionID string) DisasterRecoveryConfigsClient {
	return original.NewDisasterRecoveryConfigsClient(subscriptionID)
}
func NewDisasterRecoveryConfigsClientWithBaseURI(baseURI string, subscriptionID string) DisasterRecoveryConfigsClient {
	return original.NewDisasterRecoveryConfigsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEventHubsClient(subscriptionID string) EventHubsClient {
	return original.NewEventHubsClient(subscriptionID)
}
func NewEventHubsClientWithBaseURI(baseURI string, subscriptionID string) EventHubsClient {
	return original.NewEventHubsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMigrationConfigsClient(subscriptionID string) MigrationConfigsClient {
	return original.NewMigrationConfigsClient(subscriptionID)
}
func NewMigrationConfigsClientWithBaseURI(baseURI string, subscriptionID string) MigrationConfigsClient {
	return original.NewMigrationConfigsClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessRightsValues() []AccessRights {
	return original.PossibleAccessRightsValues()
}
func PossibleEncodingCaptureDescriptionValues() []EncodingCaptureDescription {
	return original.PossibleEncodingCaptureDescriptionValues()
}
func PossibleEntityStatusValues() []EntityStatus {
	return original.PossibleEntityStatusValues()
}
func PossibleFilterTypeValues() []FilterType {
	return original.PossibleFilterTypeValues()
}
func PossibleIPActionValues() []IPAction {
	return original.PossibleIPActionValues()
}
func PossibleKeyTypeValues() []KeyType {
	return original.PossibleKeyTypeValues()
}
func PossibleNameSpaceTypeValues() []NameSpaceType {
	return original.PossibleNameSpaceTypeValues()
}
func PossibleProvisioningStateDRValues() []ProvisioningStateDR {
	return original.PossibleProvisioningStateDRValues()
}
func PossibleRoleDisasterRecoveryValues() []RoleDisasterRecovery {
	return original.PossibleRoleDisasterRecoveryValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleUnavailableReasonValues() []UnavailableReason {
	return original.PossibleUnavailableReasonValues()
}
func NewNamespacesClient(subscriptionID string) NamespacesClient {
	return original.NewNamespacesClient(subscriptionID)
}
func NewNamespacesClientWithBaseURI(baseURI string, subscriptionID string) NamespacesClient {
	return original.NewNamespacesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPremiumMessagingRegionsClient(subscriptionID string) PremiumMessagingRegionsClient {
	return original.NewPremiumMessagingRegionsClient(subscriptionID)
}
func NewPremiumMessagingRegionsClientWithBaseURI(baseURI string, subscriptionID string) PremiumMessagingRegionsClient {
	return original.NewPremiumMessagingRegionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewQueuesClient(subscriptionID string) QueuesClient {
	return original.NewQueuesClient(subscriptionID)
}
func NewQueuesClientWithBaseURI(baseURI string, subscriptionID string) QueuesClient {
	return original.NewQueuesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRegionsClient(subscriptionID string) RegionsClient {
	return original.NewRegionsClient(subscriptionID)
}
func NewRegionsClientWithBaseURI(baseURI string, subscriptionID string) RegionsClient {
	return original.NewRegionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRulesClient(subscriptionID string) RulesClient {
	return original.NewRulesClient(subscriptionID)
}
func NewRulesClientWithBaseURI(baseURI string, subscriptionID string) RulesClient {
	return original.NewRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSubscriptionsClient(subscriptionID string) SubscriptionsClient {
	return original.NewSubscriptionsClient(subscriptionID)
}
func NewSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionsClient {
	return original.NewSubscriptionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTopicsClient(subscriptionID string) TopicsClient {
	return original.NewTopicsClient(subscriptionID)
}
func NewTopicsClientWithBaseURI(baseURI string, subscriptionID string) TopicsClient {
	return original.NewTopicsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}

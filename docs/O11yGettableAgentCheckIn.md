# O11yGettableAgentCheckIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | Older fields for backward compatibility with existing AWS agents | [optional] 
**CloudAccountId** | Pointer to **string** |  | [optional] 
**CloudIntegrationId** | Pointer to **string** |  | [optional] 
**IntegrationConfigLegacy** | Pointer to [**O11yIntegrationConfig**](O11yIntegrationConfig.md) |  | [optional] 
**IntegrationConfig** | Pointer to [**O11yProviderIntegrationConfig**](O11yProviderIntegrationConfig.md) |  | [optional] 
**ProviderAccountId** | Pointer to **string** |  | [optional] 
**RemovedAtLegacy** | Pointer to **time.Time** |  | [optional] 
**RemovedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewO11yGettableAgentCheckIn

`func NewO11yGettableAgentCheckIn() *O11yGettableAgentCheckIn`

NewO11yGettableAgentCheckIn instantiates a new O11yGettableAgentCheckIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yGettableAgentCheckInWithDefaults

`func NewO11yGettableAgentCheckInWithDefaults() *O11yGettableAgentCheckIn`

NewO11yGettableAgentCheckInWithDefaults instantiates a new O11yGettableAgentCheckIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *O11yGettableAgentCheckIn) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *O11yGettableAgentCheckIn) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *O11yGettableAgentCheckIn) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *O11yGettableAgentCheckIn) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCloudAccountId

`func (o *O11yGettableAgentCheckIn) GetCloudAccountId() string`

GetCloudAccountId returns the CloudAccountId field if non-nil, zero value otherwise.

### GetCloudAccountIdOk

`func (o *O11yGettableAgentCheckIn) GetCloudAccountIdOk() (*string, bool)`

GetCloudAccountIdOk returns a tuple with the CloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountId

`func (o *O11yGettableAgentCheckIn) SetCloudAccountId(v string)`

SetCloudAccountId sets CloudAccountId field to given value.

### HasCloudAccountId

`func (o *O11yGettableAgentCheckIn) HasCloudAccountId() bool`

HasCloudAccountId returns a boolean if a field has been set.

### GetCloudIntegrationId

`func (o *O11yGettableAgentCheckIn) GetCloudIntegrationId() string`

GetCloudIntegrationId returns the CloudIntegrationId field if non-nil, zero value otherwise.

### GetCloudIntegrationIdOk

`func (o *O11yGettableAgentCheckIn) GetCloudIntegrationIdOk() (*string, bool)`

GetCloudIntegrationIdOk returns a tuple with the CloudIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudIntegrationId

`func (o *O11yGettableAgentCheckIn) SetCloudIntegrationId(v string)`

SetCloudIntegrationId sets CloudIntegrationId field to given value.

### HasCloudIntegrationId

`func (o *O11yGettableAgentCheckIn) HasCloudIntegrationId() bool`

HasCloudIntegrationId returns a boolean if a field has been set.

### GetIntegrationConfigLegacy

`func (o *O11yGettableAgentCheckIn) GetIntegrationConfigLegacy() O11yIntegrationConfig`

GetIntegrationConfigLegacy returns the IntegrationConfigLegacy field if non-nil, zero value otherwise.

### GetIntegrationConfigLegacyOk

`func (o *O11yGettableAgentCheckIn) GetIntegrationConfigLegacyOk() (*O11yIntegrationConfig, bool)`

GetIntegrationConfigLegacyOk returns a tuple with the IntegrationConfigLegacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationConfigLegacy

`func (o *O11yGettableAgentCheckIn) SetIntegrationConfigLegacy(v O11yIntegrationConfig)`

SetIntegrationConfigLegacy sets IntegrationConfigLegacy field to given value.

### HasIntegrationConfigLegacy

`func (o *O11yGettableAgentCheckIn) HasIntegrationConfigLegacy() bool`

HasIntegrationConfigLegacy returns a boolean if a field has been set.

### GetIntegrationConfig

`func (o *O11yGettableAgentCheckIn) GetIntegrationConfig() O11yProviderIntegrationConfig`

GetIntegrationConfig returns the IntegrationConfig field if non-nil, zero value otherwise.

### GetIntegrationConfigOk

`func (o *O11yGettableAgentCheckIn) GetIntegrationConfigOk() (*O11yProviderIntegrationConfig, bool)`

GetIntegrationConfigOk returns a tuple with the IntegrationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationConfig

`func (o *O11yGettableAgentCheckIn) SetIntegrationConfig(v O11yProviderIntegrationConfig)`

SetIntegrationConfig sets IntegrationConfig field to given value.

### HasIntegrationConfig

`func (o *O11yGettableAgentCheckIn) HasIntegrationConfig() bool`

HasIntegrationConfig returns a boolean if a field has been set.

### GetProviderAccountId

`func (o *O11yGettableAgentCheckIn) GetProviderAccountId() string`

GetProviderAccountId returns the ProviderAccountId field if non-nil, zero value otherwise.

### GetProviderAccountIdOk

`func (o *O11yGettableAgentCheckIn) GetProviderAccountIdOk() (*string, bool)`

GetProviderAccountIdOk returns a tuple with the ProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountId

`func (o *O11yGettableAgentCheckIn) SetProviderAccountId(v string)`

SetProviderAccountId sets ProviderAccountId field to given value.

### HasProviderAccountId

`func (o *O11yGettableAgentCheckIn) HasProviderAccountId() bool`

HasProviderAccountId returns a boolean if a field has been set.

### GetRemovedAtLegacy

`func (o *O11yGettableAgentCheckIn) GetRemovedAtLegacy() time.Time`

GetRemovedAtLegacy returns the RemovedAtLegacy field if non-nil, zero value otherwise.

### GetRemovedAtLegacyOk

`func (o *O11yGettableAgentCheckIn) GetRemovedAtLegacyOk() (*time.Time, bool)`

GetRemovedAtLegacyOk returns a tuple with the RemovedAtLegacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedAtLegacy

`func (o *O11yGettableAgentCheckIn) SetRemovedAtLegacy(v time.Time)`

SetRemovedAtLegacy sets RemovedAtLegacy field to given value.

### HasRemovedAtLegacy

`func (o *O11yGettableAgentCheckIn) HasRemovedAtLegacy() bool`

HasRemovedAtLegacy returns a boolean if a field has been set.

### GetRemovedAt

`func (o *O11yGettableAgentCheckIn) GetRemovedAt() time.Time`

GetRemovedAt returns the RemovedAt field if non-nil, zero value otherwise.

### GetRemovedAtOk

`func (o *O11yGettableAgentCheckIn) GetRemovedAtOk() (*time.Time, bool)`

GetRemovedAtOk returns a tuple with the RemovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedAt

`func (o *O11yGettableAgentCheckIn) SetRemovedAt(v time.Time)`

SetRemovedAt sets RemovedAt field to given value.

### HasRemovedAt

`func (o *O11yGettableAgentCheckIn) HasRemovedAt() bool`

HasRemovedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



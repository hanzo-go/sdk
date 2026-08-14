# O11yAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentReport** | Pointer to [**O11yAgentReport**](O11yAgentReport.md) |  | [optional] 
**Config** | Pointer to [**O11yAccountConfig**](O11yAccountConfig.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **interface{}** |  | [optional] 
**OrgId** | Pointer to **interface{}** |  | [optional] 
**Provider** | Pointer to **interface{}** |  | [optional] 
**ProviderAccountId** | Pointer to **string** |  | [optional] 
**RemovedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewO11yAccount

`func NewO11yAccount() *O11yAccount`

NewO11yAccount instantiates a new O11yAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAccountWithDefaults

`func NewO11yAccountWithDefaults() *O11yAccount`

NewO11yAccountWithDefaults instantiates a new O11yAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentReport

`func (o *O11yAccount) GetAgentReport() O11yAgentReport`

GetAgentReport returns the AgentReport field if non-nil, zero value otherwise.

### GetAgentReportOk

`func (o *O11yAccount) GetAgentReportOk() (*O11yAgentReport, bool)`

GetAgentReportOk returns a tuple with the AgentReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentReport

`func (o *O11yAccount) SetAgentReport(v O11yAgentReport)`

SetAgentReport sets AgentReport field to given value.

### HasAgentReport

`func (o *O11yAccount) HasAgentReport() bool`

HasAgentReport returns a boolean if a field has been set.

### GetConfig

`func (o *O11yAccount) GetConfig() O11yAccountConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *O11yAccount) GetConfigOk() (*O11yAccountConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *O11yAccount) SetConfig(v O11yAccountConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *O11yAccount) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11yAccount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yAccount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yAccount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *O11yAccount) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yAccount) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yAccount) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *O11yAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *O11yAccount) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *O11yAccount) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetOrgId

`func (o *O11yAccount) GetOrgId() interface{}`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yAccount) GetOrgIdOk() (*interface{}, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yAccount) SetOrgId(v interface{})`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yAccount) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *O11yAccount) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *O11yAccount) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetProvider

`func (o *O11yAccount) GetProvider() interface{}`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *O11yAccount) GetProviderOk() (*interface{}, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *O11yAccount) SetProvider(v interface{})`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *O11yAccount) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *O11yAccount) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *O11yAccount) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetProviderAccountId

`func (o *O11yAccount) GetProviderAccountId() string`

GetProviderAccountId returns the ProviderAccountId field if non-nil, zero value otherwise.

### GetProviderAccountIdOk

`func (o *O11yAccount) GetProviderAccountIdOk() (*string, bool)`

GetProviderAccountIdOk returns a tuple with the ProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountId

`func (o *O11yAccount) SetProviderAccountId(v string)`

SetProviderAccountId sets ProviderAccountId field to given value.

### HasProviderAccountId

`func (o *O11yAccount) HasProviderAccountId() bool`

HasProviderAccountId returns a boolean if a field has been set.

### GetRemovedAt

`func (o *O11yAccount) GetRemovedAt() time.Time`

GetRemovedAt returns the RemovedAt field if non-nil, zero value otherwise.

### GetRemovedAtOk

`func (o *O11yAccount) GetRemovedAtOk() (*time.Time, bool)`

GetRemovedAtOk returns a tuple with the RemovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedAt

`func (o *O11yAccount) SetRemovedAt(v time.Time)`

SetRemovedAt sets RemovedAt field to given value.

### HasRemovedAt

`func (o *O11yAccount) HasRemovedAt() bool`

HasRemovedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yAccount) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yAccount) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yAccount) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yAccount) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# O11yGettableFunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Funnel** | Pointer to [**O11yStorableFunnel**](O11yStorableFunnel.md) |  | [optional] 
**FunnelId** | Pointer to **string** |  | [optional] 
**FunnelName** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Steps** | Pointer to [**[]O11yFunnelStep**](O11yFunnelStep.md) |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**UserEmail** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yGettableFunnel

`func NewO11yGettableFunnel() *O11yGettableFunnel`

NewO11yGettableFunnel instantiates a new O11yGettableFunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yGettableFunnelWithDefaults

`func NewO11yGettableFunnelWithDefaults() *O11yGettableFunnel`

NewO11yGettableFunnelWithDefaults instantiates a new O11yGettableFunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yGettableFunnel) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yGettableFunnel) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yGettableFunnel) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yGettableFunnel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *O11yGettableFunnel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *O11yGettableFunnel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *O11yGettableFunnel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *O11yGettableFunnel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *O11yGettableFunnel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yGettableFunnel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yGettableFunnel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yGettableFunnel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFunnel

`func (o *O11yGettableFunnel) GetFunnel() O11yStorableFunnel`

GetFunnel returns the Funnel field if non-nil, zero value otherwise.

### GetFunnelOk

`func (o *O11yGettableFunnel) GetFunnelOk() (*O11yStorableFunnel, bool)`

GetFunnelOk returns a tuple with the Funnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnel

`func (o *O11yGettableFunnel) SetFunnel(v O11yStorableFunnel)`

SetFunnel sets Funnel field to given value.

### HasFunnel

`func (o *O11yGettableFunnel) HasFunnel() bool`

HasFunnel returns a boolean if a field has been set.

### GetFunnelId

`func (o *O11yGettableFunnel) GetFunnelId() string`

GetFunnelId returns the FunnelId field if non-nil, zero value otherwise.

### GetFunnelIdOk

`func (o *O11yGettableFunnel) GetFunnelIdOk() (*string, bool)`

GetFunnelIdOk returns a tuple with the FunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelId

`func (o *O11yGettableFunnel) SetFunnelId(v string)`

SetFunnelId sets FunnelId field to given value.

### HasFunnelId

`func (o *O11yGettableFunnel) HasFunnelId() bool`

HasFunnelId returns a boolean if a field has been set.

### GetFunnelName

`func (o *O11yGettableFunnel) GetFunnelName() string`

GetFunnelName returns the FunnelName field if non-nil, zero value otherwise.

### GetFunnelNameOk

`func (o *O11yGettableFunnel) GetFunnelNameOk() (*string, bool)`

GetFunnelNameOk returns a tuple with the FunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelName

`func (o *O11yGettableFunnel) SetFunnelName(v string)`

SetFunnelName sets FunnelName field to given value.

### HasFunnelName

`func (o *O11yGettableFunnel) HasFunnelName() bool`

HasFunnelName returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yGettableFunnel) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yGettableFunnel) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yGettableFunnel) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yGettableFunnel) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSteps

`func (o *O11yGettableFunnel) GetSteps() []O11yFunnelStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *O11yGettableFunnel) GetStepsOk() (*[]O11yFunnelStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *O11yGettableFunnel) SetSteps(v []O11yFunnelStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *O11yGettableFunnel) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yGettableFunnel) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yGettableFunnel) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yGettableFunnel) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yGettableFunnel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *O11yGettableFunnel) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *O11yGettableFunnel) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *O11yGettableFunnel) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *O11yGettableFunnel) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUserEmail

`func (o *O11yGettableFunnel) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *O11yGettableFunnel) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *O11yGettableFunnel) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *O11yGettableFunnel) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



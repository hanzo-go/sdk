# O11yStorableFunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FunnelName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **interface{}** |  | [optional] 
**OrgId** | Pointer to **interface{}** |  | [optional] 
**Steps** | Pointer to [**[]O11yFunnelStep**](O11yFunnelStep.md) |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**O11yUser**](O11yUser.md) |  | [optional] 

## Methods

### NewO11yStorableFunnel

`func NewO11yStorableFunnel() *O11yStorableFunnel`

NewO11yStorableFunnel instantiates a new O11yStorableFunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yStorableFunnelWithDefaults

`func NewO11yStorableFunnelWithDefaults() *O11yStorableFunnel`

NewO11yStorableFunnelWithDefaults instantiates a new O11yStorableFunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yStorableFunnel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yStorableFunnel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yStorableFunnel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yStorableFunnel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *O11yStorableFunnel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *O11yStorableFunnel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *O11yStorableFunnel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *O11yStorableFunnel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *O11yStorableFunnel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yStorableFunnel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yStorableFunnel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yStorableFunnel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFunnelName

`func (o *O11yStorableFunnel) GetFunnelName() string`

GetFunnelName returns the FunnelName field if non-nil, zero value otherwise.

### GetFunnelNameOk

`func (o *O11yStorableFunnel) GetFunnelNameOk() (*string, bool)`

GetFunnelNameOk returns a tuple with the FunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelName

`func (o *O11yStorableFunnel) SetFunnelName(v string)`

SetFunnelName sets FunnelName field to given value.

### HasFunnelName

`func (o *O11yStorableFunnel) HasFunnelName() bool`

HasFunnelName returns a boolean if a field has been set.

### GetId

`func (o *O11yStorableFunnel) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yStorableFunnel) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yStorableFunnel) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *O11yStorableFunnel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *O11yStorableFunnel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *O11yStorableFunnel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetOrgId

`func (o *O11yStorableFunnel) GetOrgId() interface{}`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yStorableFunnel) GetOrgIdOk() (*interface{}, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yStorableFunnel) SetOrgId(v interface{})`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yStorableFunnel) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *O11yStorableFunnel) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *O11yStorableFunnel) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetSteps

`func (o *O11yStorableFunnel) GetSteps() []O11yFunnelStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *O11yStorableFunnel) GetStepsOk() (*[]O11yFunnelStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *O11yStorableFunnel) SetSteps(v []O11yFunnelStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *O11yStorableFunnel) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTags

`func (o *O11yStorableFunnel) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yStorableFunnel) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yStorableFunnel) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yStorableFunnel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yStorableFunnel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yStorableFunnel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yStorableFunnel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yStorableFunnel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *O11yStorableFunnel) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *O11yStorableFunnel) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *O11yStorableFunnel) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *O11yStorableFunnel) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUser

`func (o *O11yStorableFunnel) GetUser() O11yUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *O11yStorableFunnel) GetUserOk() (*O11yUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *O11yStorableFunnel) SetUser(v O11yUser)`

SetUser sets User field to given value.

### HasUser

`func (o *O11yStorableFunnel) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



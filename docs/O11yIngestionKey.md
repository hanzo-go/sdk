# O11yIngestionKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Limits** | Pointer to [**[]O11yLimit**](O11yLimit.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yIngestionKey

`func NewO11yIngestionKey() *O11yIngestionKey`

NewO11yIngestionKey instantiates a new O11yIngestionKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yIngestionKeyWithDefaults

`func NewO11yIngestionKeyWithDefaults() *O11yIngestionKey`

NewO11yIngestionKeyWithDefaults instantiates a new O11yIngestionKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yIngestionKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yIngestionKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yIngestionKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yIngestionKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *O11yIngestionKey) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *O11yIngestionKey) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *O11yIngestionKey) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *O11yIngestionKey) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *O11yIngestionKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yIngestionKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yIngestionKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yIngestionKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLimits

`func (o *O11yIngestionKey) GetLimits() []O11yLimit`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *O11yIngestionKey) GetLimitsOk() (*[]O11yLimit, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *O11yIngestionKey) SetLimits(v []O11yLimit)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *O11yIngestionKey) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetName

`func (o *O11yIngestionKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yIngestionKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yIngestionKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yIngestionKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *O11yIngestionKey) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yIngestionKey) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yIngestionKey) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yIngestionKey) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yIngestionKey) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yIngestionKey) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yIngestionKey) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yIngestionKey) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValue

`func (o *O11yIngestionKey) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yIngestionKey) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yIngestionKey) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yIngestionKey) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *O11yIngestionKey) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *O11yIngestionKey) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *O11yIngestionKey) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *O11yIngestionKey) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



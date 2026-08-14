# O11ySpanMapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**O11ySpanMapperConfig**](O11ySpanMapperConfig.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**FieldContext** | Pointer to **interface{}** |  | [optional] 
**GroupId** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewO11ySpanMapper

`func NewO11ySpanMapper() *O11ySpanMapper`

NewO11ySpanMapper instantiates a new O11ySpanMapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11ySpanMapperWithDefaults

`func NewO11ySpanMapperWithDefaults() *O11ySpanMapper`

NewO11ySpanMapperWithDefaults instantiates a new O11ySpanMapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *O11ySpanMapper) GetConfig() O11ySpanMapperConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *O11ySpanMapper) GetConfigOk() (*O11ySpanMapperConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *O11ySpanMapper) SetConfig(v O11ySpanMapperConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *O11ySpanMapper) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11ySpanMapper) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11ySpanMapper) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11ySpanMapper) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11ySpanMapper) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *O11ySpanMapper) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *O11ySpanMapper) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *O11ySpanMapper) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *O11ySpanMapper) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEnabled

`func (o *O11ySpanMapper) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *O11ySpanMapper) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *O11ySpanMapper) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *O11ySpanMapper) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFieldContext

`func (o *O11ySpanMapper) GetFieldContext() interface{}`

GetFieldContext returns the FieldContext field if non-nil, zero value otherwise.

### GetFieldContextOk

`func (o *O11ySpanMapper) GetFieldContextOk() (*interface{}, bool)`

GetFieldContextOk returns a tuple with the FieldContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldContext

`func (o *O11ySpanMapper) SetFieldContext(v interface{})`

SetFieldContext sets FieldContext field to given value.

### HasFieldContext

`func (o *O11ySpanMapper) HasFieldContext() bool`

HasFieldContext returns a boolean if a field has been set.

### SetFieldContextNil

`func (o *O11ySpanMapper) SetFieldContextNil(b bool)`

 SetFieldContextNil sets the value for FieldContext to be an explicit nil

### UnsetFieldContext
`func (o *O11ySpanMapper) UnsetFieldContext()`

UnsetFieldContext ensures that no value is present for FieldContext, not even an explicit nil
### GetGroupId

`func (o *O11ySpanMapper) GetGroupId() interface{}`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *O11ySpanMapper) GetGroupIdOk() (*interface{}, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *O11ySpanMapper) SetGroupId(v interface{})`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *O11ySpanMapper) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *O11ySpanMapper) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *O11ySpanMapper) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetId

`func (o *O11ySpanMapper) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11ySpanMapper) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11ySpanMapper) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *O11ySpanMapper) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *O11ySpanMapper) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *O11ySpanMapper) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *O11ySpanMapper) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11ySpanMapper) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11ySpanMapper) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11ySpanMapper) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11ySpanMapper) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11ySpanMapper) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11ySpanMapper) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11ySpanMapper) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *O11ySpanMapper) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *O11ySpanMapper) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *O11ySpanMapper) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *O11ySpanMapper) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



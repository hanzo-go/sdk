# O11ySpanMapperGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to [**O11ySpanMapperGroupCondition**](O11ySpanMapperGroupCondition.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewO11ySpanMapperGroup

`func NewO11ySpanMapperGroup() *O11ySpanMapperGroup`

NewO11ySpanMapperGroup instantiates a new O11ySpanMapperGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11ySpanMapperGroupWithDefaults

`func NewO11ySpanMapperGroupWithDefaults() *O11ySpanMapperGroup`

NewO11ySpanMapperGroupWithDefaults instantiates a new O11ySpanMapperGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *O11ySpanMapperGroup) GetCondition() O11ySpanMapperGroupCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *O11ySpanMapperGroup) GetConditionOk() (*O11ySpanMapperGroupCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *O11ySpanMapperGroup) SetCondition(v O11ySpanMapperGroupCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *O11ySpanMapperGroup) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11ySpanMapperGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11ySpanMapperGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11ySpanMapperGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11ySpanMapperGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *O11ySpanMapperGroup) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *O11ySpanMapperGroup) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *O11ySpanMapperGroup) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *O11ySpanMapperGroup) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEnabled

`func (o *O11ySpanMapperGroup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *O11ySpanMapperGroup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *O11ySpanMapperGroup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *O11ySpanMapperGroup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *O11ySpanMapperGroup) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11ySpanMapperGroup) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11ySpanMapperGroup) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *O11ySpanMapperGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *O11ySpanMapperGroup) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *O11ySpanMapperGroup) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *O11ySpanMapperGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11ySpanMapperGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11ySpanMapperGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11ySpanMapperGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *O11ySpanMapperGroup) GetOrgId() interface{}`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11ySpanMapperGroup) GetOrgIdOk() (*interface{}, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11ySpanMapperGroup) SetOrgId(v interface{})`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11ySpanMapperGroup) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *O11ySpanMapperGroup) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *O11ySpanMapperGroup) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetUpdatedAt

`func (o *O11ySpanMapperGroup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11ySpanMapperGroup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11ySpanMapperGroup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11ySpanMapperGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *O11ySpanMapperGroup) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *O11ySpanMapperGroup) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *O11ySpanMapperGroup) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *O11ySpanMapperGroup) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



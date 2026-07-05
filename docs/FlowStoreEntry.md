# FlowStoreEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFlowStoreEntry

`func NewFlowStoreEntry() *FlowStoreEntry`

NewFlowStoreEntry instantiates a new FlowStoreEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowStoreEntryWithDefaults

`func NewFlowStoreEntryWithDefaults() *FlowStoreEntry`

NewFlowStoreEntryWithDefaults instantiates a new FlowStoreEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlowStoreEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowStoreEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowStoreEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlowStoreEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *FlowStoreEntry) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlowStoreEntry) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlowStoreEntry) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FlowStoreEntry) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *FlowStoreEntry) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FlowStoreEntry) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FlowStoreEntry) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *FlowStoreEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *FlowStoreEntry) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FlowStoreEntry) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetProjectId

`func (o *FlowStoreEntry) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *FlowStoreEntry) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *FlowStoreEntry) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *FlowStoreEntry) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetCreated

`func (o *FlowStoreEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FlowStoreEntry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FlowStoreEntry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FlowStoreEntry) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *FlowStoreEntry) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FlowStoreEntry) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FlowStoreEntry) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FlowStoreEntry) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



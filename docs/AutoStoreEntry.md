# AutoStoreEntry

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

### NewAutoStoreEntry

`func NewAutoStoreEntry() *AutoStoreEntry`

NewAutoStoreEntry instantiates a new AutoStoreEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoStoreEntryWithDefaults

`func NewAutoStoreEntryWithDefaults() *AutoStoreEntry`

NewAutoStoreEntryWithDefaults instantiates a new AutoStoreEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoStoreEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoStoreEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoStoreEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoStoreEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *AutoStoreEntry) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AutoStoreEntry) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AutoStoreEntry) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AutoStoreEntry) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *AutoStoreEntry) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AutoStoreEntry) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AutoStoreEntry) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *AutoStoreEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *AutoStoreEntry) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *AutoStoreEntry) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetProjectId

`func (o *AutoStoreEntry) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AutoStoreEntry) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AutoStoreEntry) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AutoStoreEntry) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetCreated

`func (o *AutoStoreEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AutoStoreEntry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AutoStoreEntry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AutoStoreEntry) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *AutoStoreEntry) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AutoStoreEntry) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AutoStoreEntry) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AutoStoreEntry) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



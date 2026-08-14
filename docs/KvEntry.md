# KvEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | Created is when this revision was written, RFC3339. | [optional] 
**Key** | Pointer to **string** | Key is the entry&#39;s key. | [optional] 
**Operation** | Pointer to **string** | Operation is what wrote the revision: put, del or purge. | [optional] 
**Revision** | Pointer to **int32** | Revision is the entry&#39;s revision in the bucket. | [optional] 
**Value** | Pointer to **string** | Value is the value as UTF-8 text; empty for delete and purge markers. | [optional] 

## Methods

### NewKvEntry

`func NewKvEntry() *KvEntry`

NewKvEntry instantiates a new KvEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvEntryWithDefaults

`func NewKvEntryWithDefaults() *KvEntry`

NewKvEntryWithDefaults instantiates a new KvEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *KvEntry) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *KvEntry) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *KvEntry) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *KvEntry) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetKey

`func (o *KvEntry) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KvEntry) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KvEntry) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *KvEntry) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOperation

`func (o *KvEntry) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *KvEntry) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *KvEntry) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *KvEntry) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetRevision

`func (o *KvEntry) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *KvEntry) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *KvEntry) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *KvEntry) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetValue

`func (o *KvEntry) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KvEntry) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KvEntry) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KvEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



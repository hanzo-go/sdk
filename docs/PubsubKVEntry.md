# PubsubKVEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** | Base64-encoded value | [optional] 
**Revision** | Pointer to **int32** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Operation** | Pointer to **string** |  | [optional] 

## Methods

### NewPubsubKVEntry

`func NewPubsubKVEntry() *PubsubKVEntry`

NewPubsubKVEntry instantiates a new PubsubKVEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubKVEntryWithDefaults

`func NewPubsubKVEntryWithDefaults() *PubsubKVEntry`

NewPubsubKVEntryWithDefaults instantiates a new PubsubKVEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PubsubKVEntry) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PubsubKVEntry) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PubsubKVEntry) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PubsubKVEntry) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *PubsubKVEntry) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PubsubKVEntry) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PubsubKVEntry) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PubsubKVEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRevision

`func (o *PubsubKVEntry) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PubsubKVEntry) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PubsubKVEntry) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *PubsubKVEntry) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetCreated

`func (o *PubsubKVEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PubsubKVEntry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PubsubKVEntry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PubsubKVEntry) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetOperation

`func (o *PubsubKVEntry) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *PubsubKVEntry) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *PubsubKVEntry) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *PubsubKVEntry) HasOperation() bool`

HasOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



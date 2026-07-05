# MqKVEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key name. | [optional] 
**Value** | Pointer to **string** | Value (base64-encoded for binary data). | [optional] 
**Revision** | Pointer to **int32** | Revision number. | [optional] 
**Created** | Pointer to **time.Time** | Timestamp of this revision. | [optional] 
**Operation** | Pointer to **string** | Operation that produced this revision. &#x60;delete&#x60; is a tombstone. &#x60;purge&#x60; removes all prior revisions.  | [optional] 
**Bucket** | Pointer to **string** | Bucket name. | [optional] 
**Delta** | Pointer to **int32** | Number of revisions since this entry (for watch operations).  | [optional] 

## Methods

### NewMqKVEntry

`func NewMqKVEntry() *MqKVEntry`

NewMqKVEntry instantiates a new MqKVEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqKVEntryWithDefaults

`func NewMqKVEntryWithDefaults() *MqKVEntry`

NewMqKVEntryWithDefaults instantiates a new MqKVEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MqKVEntry) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MqKVEntry) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MqKVEntry) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MqKVEntry) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *MqKVEntry) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MqKVEntry) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MqKVEntry) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MqKVEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRevision

`func (o *MqKVEntry) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *MqKVEntry) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *MqKVEntry) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *MqKVEntry) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetCreated

`func (o *MqKVEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MqKVEntry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MqKVEntry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *MqKVEntry) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetOperation

`func (o *MqKVEntry) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *MqKVEntry) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *MqKVEntry) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *MqKVEntry) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetBucket

`func (o *MqKVEntry) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *MqKVEntry) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *MqKVEntry) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *MqKVEntry) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetDelta

`func (o *MqKVEntry) GetDelta() int32`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *MqKVEntry) GetDeltaOk() (*int32, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *MqKVEntry) SetDelta(v int32)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *MqKVEntry) HasDelta() bool`

HasDelta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StreamWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discard** | Pointer to **string** |  | [optional] 
**MaxAge** | Pointer to **int32** |  | [optional] 
**MaxBytes** | Pointer to **int32** |  | [optional] 
**MaxMsgs** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** | Name is the stream&#39;s name within the org: 1–64 of [A-Za-z0-9_], no dash. | [optional] 
**Retention** | Pointer to **string** |  | [optional] 
**Storage** | Pointer to **string** |  | [optional] 
**Subjects** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStreamWrite

`func NewStreamWrite() *StreamWrite`

NewStreamWrite instantiates a new StreamWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamWriteWithDefaults

`func NewStreamWriteWithDefaults() *StreamWrite`

NewStreamWriteWithDefaults instantiates a new StreamWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscard

`func (o *StreamWrite) GetDiscard() string`

GetDiscard returns the Discard field if non-nil, zero value otherwise.

### GetDiscardOk

`func (o *StreamWrite) GetDiscardOk() (*string, bool)`

GetDiscardOk returns a tuple with the Discard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscard

`func (o *StreamWrite) SetDiscard(v string)`

SetDiscard sets Discard field to given value.

### HasDiscard

`func (o *StreamWrite) HasDiscard() bool`

HasDiscard returns a boolean if a field has been set.

### GetMaxAge

`func (o *StreamWrite) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *StreamWrite) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *StreamWrite) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *StreamWrite) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetMaxBytes

`func (o *StreamWrite) GetMaxBytes() int32`

GetMaxBytes returns the MaxBytes field if non-nil, zero value otherwise.

### GetMaxBytesOk

`func (o *StreamWrite) GetMaxBytesOk() (*int32, bool)`

GetMaxBytesOk returns a tuple with the MaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytes

`func (o *StreamWrite) SetMaxBytes(v int32)`

SetMaxBytes sets MaxBytes field to given value.

### HasMaxBytes

`func (o *StreamWrite) HasMaxBytes() bool`

HasMaxBytes returns a boolean if a field has been set.

### GetMaxMsgs

`func (o *StreamWrite) GetMaxMsgs() int32`

GetMaxMsgs returns the MaxMsgs field if non-nil, zero value otherwise.

### GetMaxMsgsOk

`func (o *StreamWrite) GetMaxMsgsOk() (*int32, bool)`

GetMaxMsgsOk returns a tuple with the MaxMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgs

`func (o *StreamWrite) SetMaxMsgs(v int32)`

SetMaxMsgs sets MaxMsgs field to given value.

### HasMaxMsgs

`func (o *StreamWrite) HasMaxMsgs() bool`

HasMaxMsgs returns a boolean if a field has been set.

### GetName

`func (o *StreamWrite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamWrite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamWrite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StreamWrite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetention

`func (o *StreamWrite) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *StreamWrite) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *StreamWrite) SetRetention(v string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *StreamWrite) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetStorage

`func (o *StreamWrite) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *StreamWrite) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *StreamWrite) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *StreamWrite) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetSubjects

`func (o *StreamWrite) GetSubjects() []string`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *StreamWrite) GetSubjectsOk() (*[]string, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *StreamWrite) SetSubjects(v []string)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *StreamWrite) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



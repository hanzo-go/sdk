# MqPublishResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stream** | Pointer to **string** | Stream name if the subject is bound to a JetStream stream. Absent for core NATS subjects.  | [optional] 
**Sequence** | Pointer to **int32** | Stream sequence number. Present only for JetStream-bound subjects.  | [optional] 
**Duplicate** | Pointer to **bool** | True if the message was a duplicate (dedup window). | [optional] 

## Methods

### NewMqPublishResponse

`func NewMqPublishResponse() *MqPublishResponse`

NewMqPublishResponse instantiates a new MqPublishResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqPublishResponseWithDefaults

`func NewMqPublishResponseWithDefaults() *MqPublishResponse`

NewMqPublishResponseWithDefaults instantiates a new MqPublishResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStream

`func (o *MqPublishResponse) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *MqPublishResponse) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *MqPublishResponse) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *MqPublishResponse) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetSequence

`func (o *MqPublishResponse) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *MqPublishResponse) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *MqPublishResponse) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *MqPublishResponse) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetDuplicate

`func (o *MqPublishResponse) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *MqPublishResponse) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *MqPublishResponse) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.

### HasDuplicate

`func (o *MqPublishResponse) HasDuplicate() bool`

HasDuplicate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



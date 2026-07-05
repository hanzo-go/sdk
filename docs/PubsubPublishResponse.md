# PubsubPublishResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | Pointer to **bool** |  | [optional] 
**Stream** | Pointer to **string** | JetStream stream name (if subject is captured) | [optional] 
**Seq** | Pointer to **int32** | JetStream sequence number | [optional] 
**Duplicate** | Pointer to **bool** | Whether message was a duplicate (dedup by Msg-Id header) | [optional] 

## Methods

### NewPubsubPublishResponse

`func NewPubsubPublishResponse() *PubsubPublishResponse`

NewPubsubPublishResponse instantiates a new PubsubPublishResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubPublishResponseWithDefaults

`func NewPubsubPublishResponseWithDefaults() *PubsubPublishResponse`

NewPubsubPublishResponseWithDefaults instantiates a new PubsubPublishResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *PubsubPublishResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *PubsubPublishResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *PubsubPublishResponse) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *PubsubPublishResponse) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetStream

`func (o *PubsubPublishResponse) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *PubsubPublishResponse) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *PubsubPublishResponse) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *PubsubPublishResponse) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetSeq

`func (o *PubsubPublishResponse) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *PubsubPublishResponse) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *PubsubPublishResponse) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *PubsubPublishResponse) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetDuplicate

`func (o *PubsubPublishResponse) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *PubsubPublishResponse) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *PubsubPublishResponse) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.

### HasDuplicate

`func (o *PubsubPublishResponse) HasDuplicate() bool`

HasDuplicate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



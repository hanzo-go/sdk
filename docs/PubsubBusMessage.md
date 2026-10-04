# PubsubBusMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | Data is the payload as UTF-8 text. | [optional] 
**Headers** | Pointer to **map[string][]string** | Headers are the message&#39;s headers, when it carries any. | [optional] 
**Seq** | Pointer to **int32** | Seq is the message&#39;s stream sequence — fetched messages only. | [optional] 
**Subject** | Pointer to **string** | Subject is the message&#39;s subject in the org&#39;s own namespace. | [optional] 
**Time** | Pointer to **string** | Time is when the stream stored the message, RFC3339 — fetched messages only. | [optional] 

## Methods

### NewPubsubBusMessage

`func NewPubsubBusMessage() *PubsubBusMessage`

NewPubsubBusMessage instantiates a new PubsubBusMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubBusMessageWithDefaults

`func NewPubsubBusMessageWithDefaults() *PubsubBusMessage`

NewPubsubBusMessageWithDefaults instantiates a new PubsubBusMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PubsubBusMessage) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PubsubBusMessage) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PubsubBusMessage) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *PubsubBusMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeaders

`func (o *PubsubBusMessage) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PubsubBusMessage) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PubsubBusMessage) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PubsubBusMessage) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetSeq

`func (o *PubsubBusMessage) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *PubsubBusMessage) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *PubsubBusMessage) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *PubsubBusMessage) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetSubject

`func (o *PubsubBusMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PubsubBusMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PubsubBusMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PubsubBusMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTime

`func (o *PubsubBusMessage) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PubsubBusMessage) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PubsubBusMessage) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *PubsubBusMessage) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



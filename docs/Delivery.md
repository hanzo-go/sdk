# Delivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | Data is the payload, base64-encoded. | [optional] 
**Headers** | Pointer to **map[string][]string** | Headers are the message headers, when any were published. | [optional] 
**NumDelivered** | Pointer to **int32** | Delivered is how many times a consumer has been handed this message (pulls only). | [optional] 
**NumPending** | Pointer to **int32** | Remaining is how many messages follow this one for the consumer (pulls only). | [optional] 
**Sequence** | Pointer to **int32** | Sequence is the message&#39;s stream sequence. | [optional] 
**Subject** | Pointer to **string** | Subject is the org-relative subject the message was stored under. | [optional] 
**Timestamp** | Pointer to **time.Time** | Timestamp is when the broker stored the message. | [optional] 

## Methods

### NewDelivery

`func NewDelivery() *Delivery`

NewDelivery instantiates a new Delivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryWithDefaults

`func NewDeliveryWithDefaults() *Delivery`

NewDeliveryWithDefaults instantiates a new Delivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Delivery) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Delivery) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Delivery) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *Delivery) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeaders

`func (o *Delivery) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Delivery) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Delivery) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *Delivery) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetNumDelivered

`func (o *Delivery) GetNumDelivered() int32`

GetNumDelivered returns the NumDelivered field if non-nil, zero value otherwise.

### GetNumDeliveredOk

`func (o *Delivery) GetNumDeliveredOk() (*int32, bool)`

GetNumDeliveredOk returns a tuple with the NumDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDelivered

`func (o *Delivery) SetNumDelivered(v int32)`

SetNumDelivered sets NumDelivered field to given value.

### HasNumDelivered

`func (o *Delivery) HasNumDelivered() bool`

HasNumDelivered returns a boolean if a field has been set.

### GetNumPending

`func (o *Delivery) GetNumPending() int32`

GetNumPending returns the NumPending field if non-nil, zero value otherwise.

### GetNumPendingOk

`func (o *Delivery) GetNumPendingOk() (*int32, bool)`

GetNumPendingOk returns a tuple with the NumPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPending

`func (o *Delivery) SetNumPending(v int32)`

SetNumPending sets NumPending field to given value.

### HasNumPending

`func (o *Delivery) HasNumPending() bool`

HasNumPending returns a boolean if a field has been set.

### GetSequence

`func (o *Delivery) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *Delivery) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *Delivery) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *Delivery) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSubject

`func (o *Delivery) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Delivery) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Delivery) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Delivery) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTimestamp

`func (o *Delivery) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Delivery) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Delivery) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Delivery) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



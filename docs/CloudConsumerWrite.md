# CloudConsumerWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ack** | Pointer to **string** |  | [optional] 
**AckWait** | Pointer to **int32** |  | [optional] 
**Deliver** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to **string** |  | [optional] 
**MaxDeliver** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** | Name is the durable consumer name: 1–64 of [A-Za-z0-9_-]. | [optional] 
**Stream** | Pointer to **string** | Stream is the stream to consume, from the path. | [optional] 

## Methods

### NewCloudConsumerWrite

`func NewCloudConsumerWrite() *CloudConsumerWrite`

NewCloudConsumerWrite instantiates a new CloudConsumerWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudConsumerWriteWithDefaults

`func NewCloudConsumerWriteWithDefaults() *CloudConsumerWrite`

NewCloudConsumerWriteWithDefaults instantiates a new CloudConsumerWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAck

`func (o *CloudConsumerWrite) GetAck() string`

GetAck returns the Ack field if non-nil, zero value otherwise.

### GetAckOk

`func (o *CloudConsumerWrite) GetAckOk() (*string, bool)`

GetAckOk returns a tuple with the Ack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAck

`func (o *CloudConsumerWrite) SetAck(v string)`

SetAck sets Ack field to given value.

### HasAck

`func (o *CloudConsumerWrite) HasAck() bool`

HasAck returns a boolean if a field has been set.

### GetAckWait

`func (o *CloudConsumerWrite) GetAckWait() int32`

GetAckWait returns the AckWait field if non-nil, zero value otherwise.

### GetAckWaitOk

`func (o *CloudConsumerWrite) GetAckWaitOk() (*int32, bool)`

GetAckWaitOk returns a tuple with the AckWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckWait

`func (o *CloudConsumerWrite) SetAckWait(v int32)`

SetAckWait sets AckWait field to given value.

### HasAckWait

`func (o *CloudConsumerWrite) HasAckWait() bool`

HasAckWait returns a boolean if a field has been set.

### GetDeliver

`func (o *CloudConsumerWrite) GetDeliver() string`

GetDeliver returns the Deliver field if non-nil, zero value otherwise.

### GetDeliverOk

`func (o *CloudConsumerWrite) GetDeliverOk() (*string, bool)`

GetDeliverOk returns a tuple with the Deliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliver

`func (o *CloudConsumerWrite) SetDeliver(v string)`

SetDeliver sets Deliver field to given value.

### HasDeliver

`func (o *CloudConsumerWrite) HasDeliver() bool`

HasDeliver returns a boolean if a field has been set.

### GetFilter

`func (o *CloudConsumerWrite) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CloudConsumerWrite) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CloudConsumerWrite) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CloudConsumerWrite) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetMaxDeliver

`func (o *CloudConsumerWrite) GetMaxDeliver() int32`

GetMaxDeliver returns the MaxDeliver field if non-nil, zero value otherwise.

### GetMaxDeliverOk

`func (o *CloudConsumerWrite) GetMaxDeliverOk() (*int32, bool)`

GetMaxDeliverOk returns a tuple with the MaxDeliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliver

`func (o *CloudConsumerWrite) SetMaxDeliver(v int32)`

SetMaxDeliver sets MaxDeliver field to given value.

### HasMaxDeliver

`func (o *CloudConsumerWrite) HasMaxDeliver() bool`

HasMaxDeliver returns a boolean if a field has been set.

### GetName

`func (o *CloudConsumerWrite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudConsumerWrite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudConsumerWrite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudConsumerWrite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStream

`func (o *CloudConsumerWrite) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *CloudConsumerWrite) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *CloudConsumerWrite) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *CloudConsumerWrite) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



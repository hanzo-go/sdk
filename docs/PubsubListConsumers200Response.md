# PubsubListConsumers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consumers** | Pointer to [**[]PubsubConsumerInfo**](PubsubConsumerInfo.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewPubsubListConsumers200Response

`func NewPubsubListConsumers200Response() *PubsubListConsumers200Response`

NewPubsubListConsumers200Response instantiates a new PubsubListConsumers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubListConsumers200ResponseWithDefaults

`func NewPubsubListConsumers200ResponseWithDefaults() *PubsubListConsumers200Response`

NewPubsubListConsumers200ResponseWithDefaults instantiates a new PubsubListConsumers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumers

`func (o *PubsubListConsumers200Response) GetConsumers() []PubsubConsumerInfo`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *PubsubListConsumers200Response) GetConsumersOk() (*[]PubsubConsumerInfo, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *PubsubListConsumers200Response) SetConsumers(v []PubsubConsumerInfo)`

SetConsumers sets Consumers field to given value.

### HasConsumers

`func (o *PubsubListConsumers200Response) HasConsumers() bool`

HasConsumers returns a boolean if a field has been set.

### GetTotal

`func (o *PubsubListConsumers200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PubsubListConsumers200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PubsubListConsumers200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PubsubListConsumers200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



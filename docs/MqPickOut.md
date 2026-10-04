# MqPickOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consumers** | Pointer to [**[]MqConsumer**](MqConsumer.md) | Consumers is the page, ordered by name. | [optional] 
**Total** | Pointer to **int64** | Total is the stream&#39;s consumer count before paging. | [optional] 

## Methods

### NewMqPickOut

`func NewMqPickOut() *MqPickOut`

NewMqPickOut instantiates a new MqPickOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqPickOutWithDefaults

`func NewMqPickOutWithDefaults() *MqPickOut`

NewMqPickOutWithDefaults instantiates a new MqPickOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumers

`func (o *MqPickOut) GetConsumers() []MqConsumer`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *MqPickOut) GetConsumersOk() (*[]MqConsumer, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *MqPickOut) SetConsumers(v []MqConsumer)`

SetConsumers sets Consumers field to given value.

### HasConsumers

`func (o *MqPickOut) HasConsumers() bool`

HasConsumers returns a boolean if a field has been set.

### GetTotal

`func (o *MqPickOut) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MqPickOut) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MqPickOut) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MqPickOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



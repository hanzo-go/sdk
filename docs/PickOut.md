# PickOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consumers** | Pointer to [**[]Consumer**](Consumer.md) | Consumers is the page, ordered by name. | [optional] 
**Total** | Pointer to **int32** | Total is the stream&#39;s consumer count before paging. | [optional] 

## Methods

### NewPickOut

`func NewPickOut() *PickOut`

NewPickOut instantiates a new PickOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPickOutWithDefaults

`func NewPickOutWithDefaults() *PickOut`

NewPickOutWithDefaults instantiates a new PickOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumers

`func (o *PickOut) GetConsumers() []Consumer`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *PickOut) GetConsumersOk() (*[]Consumer, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *PickOut) SetConsumers(v []Consumer)`

SetConsumers sets Consumers field to given value.

### HasConsumers

`func (o *PickOut) HasConsumers() bool`

HasConsumers returns a boolean if a field has been set.

### GetTotal

`func (o *PickOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PickOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PickOut) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PickOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



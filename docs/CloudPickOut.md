# CloudPickOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consumers** | Pointer to [**[]CloudConsumer**](CloudConsumer.md) | Consumers is the page, ordered by name. | [optional] 
**Total** | Pointer to **int32** | Total is the stream&#39;s consumer count before paging. | [optional] 

## Methods

### NewCloudPickOut

`func NewCloudPickOut() *CloudPickOut`

NewCloudPickOut instantiates a new CloudPickOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPickOutWithDefaults

`func NewCloudPickOutWithDefaults() *CloudPickOut`

NewCloudPickOutWithDefaults instantiates a new CloudPickOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumers

`func (o *CloudPickOut) GetConsumers() []CloudConsumer`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *CloudPickOut) GetConsumersOk() (*[]CloudConsumer, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *CloudPickOut) SetConsumers(v []CloudConsumer)`

SetConsumers sets Consumers field to given value.

### HasConsumers

`func (o *CloudPickOut) HasConsumers() bool`

HasConsumers returns a boolean if a field has been set.

### GetTotal

`func (o *CloudPickOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CloudPickOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CloudPickOut) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CloudPickOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



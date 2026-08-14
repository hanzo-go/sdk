# ConsumerPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ConsumerRecord**](ConsumerRecord.md) | Data are the stream&#39;s consumers. | [optional] 

## Methods

### NewConsumerPage

`func NewConsumerPage() *ConsumerPage`

NewConsumerPage instantiates a new ConsumerPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerPageWithDefaults

`func NewConsumerPageWithDefaults() *ConsumerPage`

NewConsumerPageWithDefaults instantiates a new ConsumerPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ConsumerPage) GetData() []ConsumerRecord`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConsumerPage) GetDataOk() (*[]ConsumerRecord, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConsumerPage) SetData(v []ConsumerRecord)`

SetData sets Data field to given value.

### HasData

`func (o *ConsumerPage) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



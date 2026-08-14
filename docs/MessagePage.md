# MessagePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]BusMessage**](BusMessage.md) | Data are the messages in this batch — empty when nothing was pending within the wait. | [optional] 

## Methods

### NewMessagePage

`func NewMessagePage() *MessagePage`

NewMessagePage instantiates a new MessagePage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagePageWithDefaults

`func NewMessagePageWithDefaults() *MessagePage`

NewMessagePageWithDefaults instantiates a new MessagePage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MessagePage) GetData() []BusMessage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MessagePage) GetDataOk() (*[]BusMessage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MessagePage) SetData(v []BusMessage)`

SetData sets Data field to given value.

### HasData

`func (o *MessagePage) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DeliveryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]DeliveryRow**](DeliveryRow.md) | Data is the matching attempts, newest first. | [optional] 

## Methods

### NewDeliveryList

`func NewDeliveryList() *DeliveryList`

NewDeliveryList instantiates a new DeliveryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryListWithDefaults

`func NewDeliveryListWithDefaults() *DeliveryList`

NewDeliveryListWithDefaults instantiates a new DeliveryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DeliveryList) GetData() []DeliveryRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeliveryList) GetDataOk() (*[]DeliveryRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeliveryList) SetData(v []DeliveryRow)`

SetData sets Data field to given value.

### HasData

`func (o *DeliveryList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



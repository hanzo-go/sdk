# CloudDeliveryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudDeliveryRow**](CloudDeliveryRow.md) | Data is the matching attempts, newest first. | [optional] 

## Methods

### NewCloudDeliveryList

`func NewCloudDeliveryList() *CloudDeliveryList`

NewCloudDeliveryList instantiates a new CloudDeliveryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDeliveryListWithDefaults

`func NewCloudDeliveryListWithDefaults() *CloudDeliveryList`

NewCloudDeliveryListWithDefaults instantiates a new CloudDeliveryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudDeliveryList) GetData() []CloudDeliveryRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudDeliveryList) GetDataOk() (*[]CloudDeliveryRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudDeliveryList) SetData(v []CloudDeliveryRow)`

SetData sets Data field to given value.

### HasData

`func (o *CloudDeliveryList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



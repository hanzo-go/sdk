# WebhookDeliveryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]WebhookDeliveryRow**](WebhookDeliveryRow.md) | Data is the matching attempts, newest first. | [optional] 

## Methods

### NewWebhookDeliveryList

`func NewWebhookDeliveryList() *WebhookDeliveryList`

NewWebhookDeliveryList instantiates a new WebhookDeliveryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeliveryListWithDefaults

`func NewWebhookDeliveryListWithDefaults() *WebhookDeliveryList`

NewWebhookDeliveryListWithDefaults instantiates a new WebhookDeliveryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WebhookDeliveryList) GetData() []WebhookDeliveryRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookDeliveryList) GetDataOk() (*[]WebhookDeliveryRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookDeliveryList) SetData(v []WebhookDeliveryRow)`

SetData sets Data field to given value.

### HasData

`func (o *WebhookDeliveryList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



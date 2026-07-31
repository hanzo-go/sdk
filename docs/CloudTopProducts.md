# CloudTopProducts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when the product-event table could not be read. | [optional] 
**Items** | Pointer to [**[]CloudProductRow**](CloudProductRow.md) | Items is the ranked products, highest revenue first. Empty rather than absent. | [optional] 
**Reason** | Pointer to **string** | Reason says why the lens is unavailable. Omitted when it is available. | [optional] 
**Source** | Pointer to **string** | Source is the warehouse table the lens read. | [optional] 

## Methods

### NewCloudTopProducts

`func NewCloudTopProducts() *CloudTopProducts`

NewCloudTopProducts instantiates a new CloudTopProducts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTopProductsWithDefaults

`func NewCloudTopProductsWithDefaults() *CloudTopProducts`

NewCloudTopProductsWithDefaults instantiates a new CloudTopProducts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *CloudTopProducts) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CloudTopProducts) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CloudTopProducts) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CloudTopProducts) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetItems

`func (o *CloudTopProducts) GetItems() []CloudProductRow`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CloudTopProducts) GetItemsOk() (*[]CloudProductRow, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CloudTopProducts) SetItems(v []CloudProductRow)`

SetItems sets Items field to given value.

### HasItems

`func (o *CloudTopProducts) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetReason

`func (o *CloudTopProducts) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CloudTopProducts) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CloudTopProducts) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CloudTopProducts) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSource

`func (o *CloudTopProducts) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudTopProducts) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudTopProducts) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudTopProducts) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



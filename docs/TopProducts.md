# TopProducts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when the product-event table could not be read. | [optional] 
**Items** | Pointer to [**[]ProductRow**](ProductRow.md) | Items is the ranked products, highest revenue first. Empty rather than absent. | [optional] 
**Reason** | Pointer to **string** | Reason says why the lens is unavailable. Omitted when it is available. | [optional] 
**Source** | Pointer to **string** | Source is the warehouse table the lens read. | [optional] 

## Methods

### NewTopProducts

`func NewTopProducts() *TopProducts`

NewTopProducts instantiates a new TopProducts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopProductsWithDefaults

`func NewTopProductsWithDefaults() *TopProducts`

NewTopProductsWithDefaults instantiates a new TopProducts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *TopProducts) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *TopProducts) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *TopProducts) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *TopProducts) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetItems

`func (o *TopProducts) GetItems() []ProductRow`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TopProducts) GetItemsOk() (*[]ProductRow, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TopProducts) SetItems(v []ProductRow)`

SetItems sets Items field to given value.

### HasItems

`func (o *TopProducts) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetReason

`func (o *TopProducts) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TopProducts) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TopProducts) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TopProducts) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSource

`func (o *TopProducts) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TopProducts) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TopProducts) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TopProducts) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



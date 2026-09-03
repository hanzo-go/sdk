# ProductRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to **int64** | Orders is how many order_completed events carried it. | [optional] 
**ProductId** | Pointer to **string** | ProductID is the product the order events named. | [optional] 
**Revenue** | Pointer to **float64** | Revenue is the total they carried, in the events&#39; own currency unit. | [optional] 
**Units** | Pointer to **int64** | Units is the summed quantity sold. | [optional] 

## Methods

### NewProductRow

`func NewProductRow() *ProductRow`

NewProductRow instantiates a new ProductRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductRowWithDefaults

`func NewProductRowWithDefaults() *ProductRow`

NewProductRowWithDefaults instantiates a new ProductRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *ProductRow) GetOrders() int64`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *ProductRow) GetOrdersOk() (*int64, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *ProductRow) SetOrders(v int64)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *ProductRow) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetProductId

`func (o *ProductRow) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductRow) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductRow) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductRow) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetRevenue

`func (o *ProductRow) GetRevenue() float64`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *ProductRow) GetRevenueOk() (*float64, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *ProductRow) SetRevenue(v float64)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *ProductRow) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetUnits

`func (o *ProductRow) GetUnits() int64`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ProductRow) GetUnitsOk() (*int64, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ProductRow) SetUnits(v int64)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *ProductRow) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



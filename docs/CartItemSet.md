# CartItemSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the cart to amend, from the path. | [optional] 
**Product** | Pointer to **string** | Product names the catalog product to set, by its id or its URL slug. Give this or Variant, never both; a request naming neither is refused. | [optional] 
**Quantity** | Pointer to **int64** | Quantity is how many of that item the cart should hold AFTER this call — it is the resulting count, not a delta, so sending 3 twice leaves 3 and not 6. ZERO REMOVES the line, which is the only way to take an item out. | [optional] 
**Variant** | Pointer to **string** | Variant names the specific sellable variant to set, by its id or its SKU. Prefer it over Product for anything sold in sizes, colours or tiers — the price and the stock are the variant&#39;s, not the product&#39;s. | [optional] 

## Methods

### NewCartItemSet

`func NewCartItemSet() *CartItemSet`

NewCartItemSet instantiates a new CartItemSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartItemSetWithDefaults

`func NewCartItemSetWithDefaults() *CartItemSet`

NewCartItemSetWithDefaults instantiates a new CartItemSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CartItemSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CartItemSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CartItemSet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CartItemSet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProduct

`func (o *CartItemSet) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CartItemSet) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CartItemSet) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CartItemSet) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetQuantity

`func (o *CartItemSet) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CartItemSet) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CartItemSet) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CartItemSet) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetVariant

`func (o *CartItemSet) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *CartItemSet) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *CartItemSet) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *CartItemSet) HasVariant() bool`

HasVariant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



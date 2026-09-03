# CartItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Free** | Pointer to **bool** | Free reports a line that costs nothing because a coupon or a promotion made it so, rather than because its price is zero. | [optional] 
**Id** | Pointer to **string** | ID is the line&#39;s identity — the variant id when the line is a variant, otherwise the product id. It is what a subsequent set call addresses. | [optional] 
**Kind** | Pointer to **string** | Kind is \&quot;variant\&quot; when this line is a specific sellable variant and \&quot;product\&quot; when it is the product itself. | [optional] 
**Name** | Pointer to **string** | Name is the item&#39;s display name, cached onto the line when it was added so a cart renders without a second read. | [optional] 
**PriceCents** | Pointer to **int64** | PriceCents is the unit price in whole cents, cached at the moment the line was added. The line&#39;s contribution to the cart is this times Quantity. | [optional] 
**Quantity** | Pointer to **int64** | Quantity is how many units of this item the cart holds. | [optional] 
**Sku** | Pointer to **string** | SKU is the line&#39;s stock-keeping unit — the variant&#39;s when it has one, otherwise the product&#39;s. Empty when neither carries one. | [optional] 

## Methods

### NewCartItem

`func NewCartItem() *CartItem`

NewCartItem instantiates a new CartItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartItemWithDefaults

`func NewCartItemWithDefaults() *CartItem`

NewCartItemWithDefaults instantiates a new CartItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFree

`func (o *CartItem) GetFree() bool`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *CartItem) GetFreeOk() (*bool, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *CartItem) SetFree(v bool)`

SetFree sets Free field to given value.

### HasFree

`func (o *CartItem) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetId

`func (o *CartItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CartItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CartItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CartItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CartItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CartItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CartItem) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CartItem) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *CartItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CartItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CartItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CartItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriceCents

`func (o *CartItem) GetPriceCents() int64`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *CartItem) GetPriceCentsOk() (*int64, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *CartItem) SetPriceCents(v int64)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *CartItem) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.

### GetQuantity

`func (o *CartItem) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CartItem) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CartItem) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CartItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSku

`func (o *CartItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CartItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CartItem) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CartItem) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



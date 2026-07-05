# CommerceLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **string** |  | [optional] 
**ProductName** | Pointer to **string** |  | [optional] 
**ProductSlug** | Pointer to **string** |  | [optional] 
**ProductSKU** | Pointer to **string** |  | [optional] 
**VariantId** | Pointer to **string** |  | [optional] 
**VariantName** | Pointer to **string** |  | [optional] 
**VariantSKU** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Price** | Pointer to **int32** | Price in cents | [optional] 
**Taxable** | Pointer to **bool** |  | [optional] 
**Free** | Pointer to **bool** |  | [optional] 

## Methods

### NewCommerceLineItem

`func NewCommerceLineItem() *CommerceLineItem`

NewCommerceLineItem instantiates a new CommerceLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceLineItemWithDefaults

`func NewCommerceLineItemWithDefaults() *CommerceLineItem`

NewCommerceLineItemWithDefaults instantiates a new CommerceLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *CommerceLineItem) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CommerceLineItem) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CommerceLineItem) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *CommerceLineItem) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductName

`func (o *CommerceLineItem) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *CommerceLineItem) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *CommerceLineItem) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *CommerceLineItem) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductSlug

`func (o *CommerceLineItem) GetProductSlug() string`

GetProductSlug returns the ProductSlug field if non-nil, zero value otherwise.

### GetProductSlugOk

`func (o *CommerceLineItem) GetProductSlugOk() (*string, bool)`

GetProductSlugOk returns a tuple with the ProductSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSlug

`func (o *CommerceLineItem) SetProductSlug(v string)`

SetProductSlug sets ProductSlug field to given value.

### HasProductSlug

`func (o *CommerceLineItem) HasProductSlug() bool`

HasProductSlug returns a boolean if a field has been set.

### GetProductSKU

`func (o *CommerceLineItem) GetProductSKU() string`

GetProductSKU returns the ProductSKU field if non-nil, zero value otherwise.

### GetProductSKUOk

`func (o *CommerceLineItem) GetProductSKUOk() (*string, bool)`

GetProductSKUOk returns a tuple with the ProductSKU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSKU

`func (o *CommerceLineItem) SetProductSKU(v string)`

SetProductSKU sets ProductSKU field to given value.

### HasProductSKU

`func (o *CommerceLineItem) HasProductSKU() bool`

HasProductSKU returns a boolean if a field has been set.

### GetVariantId

`func (o *CommerceLineItem) GetVariantId() string`

GetVariantId returns the VariantId field if non-nil, zero value otherwise.

### GetVariantIdOk

`func (o *CommerceLineItem) GetVariantIdOk() (*string, bool)`

GetVariantIdOk returns a tuple with the VariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantId

`func (o *CommerceLineItem) SetVariantId(v string)`

SetVariantId sets VariantId field to given value.

### HasVariantId

`func (o *CommerceLineItem) HasVariantId() bool`

HasVariantId returns a boolean if a field has been set.

### GetVariantName

`func (o *CommerceLineItem) GetVariantName() string`

GetVariantName returns the VariantName field if non-nil, zero value otherwise.

### GetVariantNameOk

`func (o *CommerceLineItem) GetVariantNameOk() (*string, bool)`

GetVariantNameOk returns a tuple with the VariantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantName

`func (o *CommerceLineItem) SetVariantName(v string)`

SetVariantName sets VariantName field to given value.

### HasVariantName

`func (o *CommerceLineItem) HasVariantName() bool`

HasVariantName returns a boolean if a field has been set.

### GetVariantSKU

`func (o *CommerceLineItem) GetVariantSKU() string`

GetVariantSKU returns the VariantSKU field if non-nil, zero value otherwise.

### GetVariantSKUOk

`func (o *CommerceLineItem) GetVariantSKUOk() (*string, bool)`

GetVariantSKUOk returns a tuple with the VariantSKU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantSKU

`func (o *CommerceLineItem) SetVariantSKU(v string)`

SetVariantSKU sets VariantSKU field to given value.

### HasVariantSKU

`func (o *CommerceLineItem) HasVariantSKU() bool`

HasVariantSKU returns a boolean if a field has been set.

### GetQuantity

`func (o *CommerceLineItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CommerceLineItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CommerceLineItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CommerceLineItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPrice

`func (o *CommerceLineItem) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CommerceLineItem) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CommerceLineItem) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CommerceLineItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTaxable

`func (o *CommerceLineItem) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *CommerceLineItem) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *CommerceLineItem) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *CommerceLineItem) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetFree

`func (o *CommerceLineItem) GetFree() bool`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *CommerceLineItem) GetFreeOk() (*bool, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *CommerceLineItem) SetFree(v bool)`

SetFree sets Free field to given value.

### HasFree

`func (o *CommerceLineItem) HasFree() bool`

HasFree returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



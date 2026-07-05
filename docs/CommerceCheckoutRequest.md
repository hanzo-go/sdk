# CommerceCheckoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CartId** | Pointer to **string** |  | [optional] 
**Buyer** | Pointer to [**CommerceBuyer**](CommerceBuyer.md) |  | [optional] 
**Items** | Pointer to [**[]CommerceLineItem**](CommerceLineItem.md) |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ShippingMethod** | Pointer to **string** |  | [optional] 
**CouponCodes** | Pointer to **[]string** |  | [optional] 
**Payment** | Pointer to [**CommercePaymentRequest**](CommercePaymentRequest.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCommerceCheckoutRequest

`func NewCommerceCheckoutRequest() *CommerceCheckoutRequest`

NewCommerceCheckoutRequest instantiates a new CommerceCheckoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceCheckoutRequestWithDefaults

`func NewCommerceCheckoutRequestWithDefaults() *CommerceCheckoutRequest`

NewCommerceCheckoutRequestWithDefaults instantiates a new CommerceCheckoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCartId

`func (o *CommerceCheckoutRequest) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *CommerceCheckoutRequest) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *CommerceCheckoutRequest) SetCartId(v string)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *CommerceCheckoutRequest) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### GetBuyer

`func (o *CommerceCheckoutRequest) GetBuyer() CommerceBuyer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *CommerceCheckoutRequest) GetBuyerOk() (*CommerceBuyer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *CommerceCheckoutRequest) SetBuyer(v CommerceBuyer)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *CommerceCheckoutRequest) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetItems

`func (o *CommerceCheckoutRequest) GetItems() []CommerceLineItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CommerceCheckoutRequest) GetItemsOk() (*[]CommerceLineItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CommerceCheckoutRequest) SetItems(v []CommerceLineItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *CommerceCheckoutRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetCurrency

`func (o *CommerceCheckoutRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CommerceCheckoutRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CommerceCheckoutRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CommerceCheckoutRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetShippingMethod

`func (o *CommerceCheckoutRequest) GetShippingMethod() string`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *CommerceCheckoutRequest) GetShippingMethodOk() (*string, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *CommerceCheckoutRequest) SetShippingMethod(v string)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *CommerceCheckoutRequest) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### GetCouponCodes

`func (o *CommerceCheckoutRequest) GetCouponCodes() []string`

GetCouponCodes returns the CouponCodes field if non-nil, zero value otherwise.

### GetCouponCodesOk

`func (o *CommerceCheckoutRequest) GetCouponCodesOk() (*[]string, bool)`

GetCouponCodesOk returns a tuple with the CouponCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCodes

`func (o *CommerceCheckoutRequest) SetCouponCodes(v []string)`

SetCouponCodes sets CouponCodes field to given value.

### HasCouponCodes

`func (o *CommerceCheckoutRequest) HasCouponCodes() bool`

HasCouponCodes returns a boolean if a field has been set.

### GetPayment

`func (o *CommerceCheckoutRequest) GetPayment() CommercePaymentRequest`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *CommerceCheckoutRequest) GetPaymentOk() (*CommercePaymentRequest, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *CommerceCheckoutRequest) SetPayment(v CommercePaymentRequest)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *CommerceCheckoutRequest) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetMetadata

`func (o *CommerceCheckoutRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CommerceCheckoutRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CommerceCheckoutRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CommerceCheckoutRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



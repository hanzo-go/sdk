# CommerceCart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**StoreId** | Pointer to **string** |  | [optional] 
**CampaignId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**LineTotal** | Pointer to **int32** |  | [optional] 
**Discount** | Pointer to **int32** |  | [optional] 
**Subtotal** | Pointer to **int32** |  | [optional] 
**Shipping** | Pointer to **int32** |  | [optional] 
**Tax** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**BillingAddress** | Pointer to [**CommerceAddress**](CommerceAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**CommerceAddress**](CommerceAddress.md) |  | [optional] 
**Items** | Pointer to [**[]CommerceLineItem**](CommerceLineItem.md) |  | [optional] 
**Coupons** | Pointer to [**[]CommerceCoupon**](CommerceCoupon.md) |  | [optional] 
**CouponCodes** | Pointer to **[]string** |  | [optional] 
**ReferrerId** | Pointer to **string** |  | [optional] 
**Gift** | Pointer to **bool** |  | [optional] 
**GiftMessage** | Pointer to **string** |  | [optional] 
**GiftEmail** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewCommerceCart

`func NewCommerceCart() *CommerceCart`

NewCommerceCart instantiates a new CommerceCart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceCartWithDefaults

`func NewCommerceCartWithDefaults() *CommerceCart`

NewCommerceCartWithDefaults instantiates a new CommerceCart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommerceCart) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommerceCart) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommerceCart) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommerceCart) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStoreId

`func (o *CommerceCart) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *CommerceCart) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *CommerceCart) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *CommerceCart) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetCampaignId

`func (o *CommerceCart) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CommerceCart) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CommerceCart) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *CommerceCart) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetUserId

`func (o *CommerceCart) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CommerceCart) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CommerceCart) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CommerceCart) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEmail

`func (o *CommerceCart) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CommerceCart) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CommerceCart) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CommerceCart) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOrderId

`func (o *CommerceCart) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CommerceCart) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CommerceCart) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CommerceCart) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetStatus

`func (o *CommerceCart) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommerceCart) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommerceCart) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommerceCart) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrency

`func (o *CommerceCart) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CommerceCart) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CommerceCart) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CommerceCart) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetLineTotal

`func (o *CommerceCart) GetLineTotal() int32`

GetLineTotal returns the LineTotal field if non-nil, zero value otherwise.

### GetLineTotalOk

`func (o *CommerceCart) GetLineTotalOk() (*int32, bool)`

GetLineTotalOk returns a tuple with the LineTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineTotal

`func (o *CommerceCart) SetLineTotal(v int32)`

SetLineTotal sets LineTotal field to given value.

### HasLineTotal

`func (o *CommerceCart) HasLineTotal() bool`

HasLineTotal returns a boolean if a field has been set.

### GetDiscount

`func (o *CommerceCart) GetDiscount() int32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *CommerceCart) GetDiscountOk() (*int32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *CommerceCart) SetDiscount(v int32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *CommerceCart) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetSubtotal

`func (o *CommerceCart) GetSubtotal() int32`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *CommerceCart) GetSubtotalOk() (*int32, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *CommerceCart) SetSubtotal(v int32)`

SetSubtotal sets Subtotal field to given value.

### HasSubtotal

`func (o *CommerceCart) HasSubtotal() bool`

HasSubtotal returns a boolean if a field has been set.

### GetShipping

`func (o *CommerceCart) GetShipping() int32`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *CommerceCart) GetShippingOk() (*int32, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *CommerceCart) SetShipping(v int32)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *CommerceCart) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetTax

`func (o *CommerceCart) GetTax() int32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *CommerceCart) GetTaxOk() (*int32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *CommerceCart) SetTax(v int32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *CommerceCart) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTotal

`func (o *CommerceCart) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CommerceCart) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CommerceCart) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CommerceCart) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetBillingAddress

`func (o *CommerceCart) GetBillingAddress() CommerceAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *CommerceCart) GetBillingAddressOk() (*CommerceAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *CommerceCart) SetBillingAddress(v CommerceAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *CommerceCart) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *CommerceCart) GetShippingAddress() CommerceAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *CommerceCart) GetShippingAddressOk() (*CommerceAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *CommerceCart) SetShippingAddress(v CommerceAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *CommerceCart) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetItems

`func (o *CommerceCart) GetItems() []CommerceLineItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CommerceCart) GetItemsOk() (*[]CommerceLineItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CommerceCart) SetItems(v []CommerceLineItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *CommerceCart) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetCoupons

`func (o *CommerceCart) GetCoupons() []CommerceCoupon`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *CommerceCart) GetCouponsOk() (*[]CommerceCoupon, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupons

`func (o *CommerceCart) SetCoupons(v []CommerceCoupon)`

SetCoupons sets Coupons field to given value.

### HasCoupons

`func (o *CommerceCart) HasCoupons() bool`

HasCoupons returns a boolean if a field has been set.

### GetCouponCodes

`func (o *CommerceCart) GetCouponCodes() []string`

GetCouponCodes returns the CouponCodes field if non-nil, zero value otherwise.

### GetCouponCodesOk

`func (o *CommerceCart) GetCouponCodesOk() (*[]string, bool)`

GetCouponCodesOk returns a tuple with the CouponCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCodes

`func (o *CommerceCart) SetCouponCodes(v []string)`

SetCouponCodes sets CouponCodes field to given value.

### HasCouponCodes

`func (o *CommerceCart) HasCouponCodes() bool`

HasCouponCodes returns a boolean if a field has been set.

### GetReferrerId

`func (o *CommerceCart) GetReferrerId() string`

GetReferrerId returns the ReferrerId field if non-nil, zero value otherwise.

### GetReferrerIdOk

`func (o *CommerceCart) GetReferrerIdOk() (*string, bool)`

GetReferrerIdOk returns a tuple with the ReferrerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrerId

`func (o *CommerceCart) SetReferrerId(v string)`

SetReferrerId sets ReferrerId field to given value.

### HasReferrerId

`func (o *CommerceCart) HasReferrerId() bool`

HasReferrerId returns a boolean if a field has been set.

### GetGift

`func (o *CommerceCart) GetGift() bool`

GetGift returns the Gift field if non-nil, zero value otherwise.

### GetGiftOk

`func (o *CommerceCart) GetGiftOk() (*bool, bool)`

GetGiftOk returns a tuple with the Gift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGift

`func (o *CommerceCart) SetGift(v bool)`

SetGift sets Gift field to given value.

### HasGift

`func (o *CommerceCart) HasGift() bool`

HasGift returns a boolean if a field has been set.

### GetGiftMessage

`func (o *CommerceCart) GetGiftMessage() string`

GetGiftMessage returns the GiftMessage field if non-nil, zero value otherwise.

### GetGiftMessageOk

`func (o *CommerceCart) GetGiftMessageOk() (*string, bool)`

GetGiftMessageOk returns a tuple with the GiftMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftMessage

`func (o *CommerceCart) SetGiftMessage(v string)`

SetGiftMessage sets GiftMessage field to given value.

### HasGiftMessage

`func (o *CommerceCart) HasGiftMessage() bool`

HasGiftMessage returns a boolean if a field has been set.

### GetGiftEmail

`func (o *CommerceCart) GetGiftEmail() string`

GetGiftEmail returns the GiftEmail field if non-nil, zero value otherwise.

### GetGiftEmailOk

`func (o *CommerceCart) GetGiftEmailOk() (*string, bool)`

GetGiftEmailOk returns a tuple with the GiftEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftEmail

`func (o *CommerceCart) SetGiftEmail(v string)`

SetGiftEmail sets GiftEmail field to given value.

### HasGiftEmail

`func (o *CommerceCart) HasGiftEmail() bool`

HasGiftEmail returns a boolean if a field has been set.

### GetMetadata

`func (o *CommerceCart) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CommerceCart) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CommerceCart) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CommerceCart) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommerceCart) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommerceCart) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommerceCart) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommerceCart) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CommerceCart) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommerceCart) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommerceCart) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CommerceCart) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



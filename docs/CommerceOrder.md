# CommerceOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Number** | Pointer to **int32** |  | [optional] [readonly] 
**StoreId** | Pointer to **string** |  | [optional] 
**CampaignId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**CartId** | Pointer to **string** |  | [optional] 
**ReferrerId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**CommerceOrderStatus**](CommerceOrderStatus.md) |  | [optional] 
**PaymentStatus** | Pointer to [**CommercePaymentStatus**](CommercePaymentStatus.md) |  | [optional] 
**Preorder** | Pointer to **bool** |  | [optional] 
**Currency** | Pointer to **string** | 3-letter ISO currency code | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**ShippingMethod** | Pointer to **string** |  | [optional] 
**LineTotal** | Pointer to **int32** | Sum of line items in cents | [optional] 
**Discount** | Pointer to **int32** | Discount in cents | [optional] 
**Subtotal** | Pointer to **int32** | Subtotal in cents | [optional] 
**Shipping** | Pointer to **int32** | Shipping cost in cents | [optional] 
**Tax** | Pointer to **int32** | Tax in cents | [optional] 
**Total** | Pointer to **int32** | Total in cents | [optional] 
**Balance** | Pointer to **int32** | Balance owed in cents | [optional] 
**Paid** | Pointer to **int32** | Amount paid in cents | [optional] 
**Refunded** | Pointer to **int32** | Amount refunded in cents | [optional] 
**BillingAddress** | Pointer to [**CommerceAddress**](CommerceAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**CommerceAddress**](CommerceAddress.md) |  | [optional] 
**Items** | Pointer to [**[]CommerceLineItem**](CommerceLineItem.md) |  | [optional] 
**Coupons** | Pointer to [**[]CommerceCoupon**](CommerceCoupon.md) |  | [optional] 
**CouponCodes** | Pointer to **[]string** |  | [optional] 
**Fulfillment** | Pointer to [**CommerceFulfillment**](CommerceFulfillment.md) |  | [optional] 
**Gift** | Pointer to **bool** |  | [optional] 
**GiftMessage** | Pointer to **string** |  | [optional] 
**GiftEmail** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Test** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewCommerceOrder

`func NewCommerceOrder() *CommerceOrder`

NewCommerceOrder instantiates a new CommerceOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceOrderWithDefaults

`func NewCommerceOrderWithDefaults() *CommerceOrder`

NewCommerceOrderWithDefaults instantiates a new CommerceOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommerceOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommerceOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommerceOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommerceOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumber

`func (o *CommerceOrder) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CommerceOrder) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CommerceOrder) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CommerceOrder) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetStoreId

`func (o *CommerceOrder) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *CommerceOrder) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *CommerceOrder) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *CommerceOrder) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetCampaignId

`func (o *CommerceOrder) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CommerceOrder) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CommerceOrder) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *CommerceOrder) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetUserId

`func (o *CommerceOrder) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CommerceOrder) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CommerceOrder) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CommerceOrder) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEmail

`func (o *CommerceOrder) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CommerceOrder) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CommerceOrder) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CommerceOrder) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCartId

`func (o *CommerceOrder) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *CommerceOrder) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *CommerceOrder) SetCartId(v string)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *CommerceOrder) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### GetReferrerId

`func (o *CommerceOrder) GetReferrerId() string`

GetReferrerId returns the ReferrerId field if non-nil, zero value otherwise.

### GetReferrerIdOk

`func (o *CommerceOrder) GetReferrerIdOk() (*string, bool)`

GetReferrerIdOk returns a tuple with the ReferrerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrerId

`func (o *CommerceOrder) SetReferrerId(v string)`

SetReferrerId sets ReferrerId field to given value.

### HasReferrerId

`func (o *CommerceOrder) HasReferrerId() bool`

HasReferrerId returns a boolean if a field has been set.

### GetStatus

`func (o *CommerceOrder) GetStatus() CommerceOrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommerceOrder) GetStatusOk() (*CommerceOrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommerceOrder) SetStatus(v CommerceOrderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommerceOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *CommerceOrder) GetPaymentStatus() CommercePaymentStatus`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *CommerceOrder) GetPaymentStatusOk() (*CommercePaymentStatus, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *CommerceOrder) SetPaymentStatus(v CommercePaymentStatus)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *CommerceOrder) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetPreorder

`func (o *CommerceOrder) GetPreorder() bool`

GetPreorder returns the Preorder field if non-nil, zero value otherwise.

### GetPreorderOk

`func (o *CommerceOrder) GetPreorderOk() (*bool, bool)`

GetPreorderOk returns a tuple with the Preorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreorder

`func (o *CommerceOrder) SetPreorder(v bool)`

SetPreorder sets Preorder field to given value.

### HasPreorder

`func (o *CommerceOrder) HasPreorder() bool`

HasPreorder returns a boolean if a field has been set.

### GetCurrency

`func (o *CommerceOrder) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CommerceOrder) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CommerceOrder) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CommerceOrder) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMode

`func (o *CommerceOrder) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CommerceOrder) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CommerceOrder) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CommerceOrder) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetShippingMethod

`func (o *CommerceOrder) GetShippingMethod() string`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *CommerceOrder) GetShippingMethodOk() (*string, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *CommerceOrder) SetShippingMethod(v string)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *CommerceOrder) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### GetLineTotal

`func (o *CommerceOrder) GetLineTotal() int32`

GetLineTotal returns the LineTotal field if non-nil, zero value otherwise.

### GetLineTotalOk

`func (o *CommerceOrder) GetLineTotalOk() (*int32, bool)`

GetLineTotalOk returns a tuple with the LineTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineTotal

`func (o *CommerceOrder) SetLineTotal(v int32)`

SetLineTotal sets LineTotal field to given value.

### HasLineTotal

`func (o *CommerceOrder) HasLineTotal() bool`

HasLineTotal returns a boolean if a field has been set.

### GetDiscount

`func (o *CommerceOrder) GetDiscount() int32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *CommerceOrder) GetDiscountOk() (*int32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *CommerceOrder) SetDiscount(v int32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *CommerceOrder) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetSubtotal

`func (o *CommerceOrder) GetSubtotal() int32`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *CommerceOrder) GetSubtotalOk() (*int32, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *CommerceOrder) SetSubtotal(v int32)`

SetSubtotal sets Subtotal field to given value.

### HasSubtotal

`func (o *CommerceOrder) HasSubtotal() bool`

HasSubtotal returns a boolean if a field has been set.

### GetShipping

`func (o *CommerceOrder) GetShipping() int32`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *CommerceOrder) GetShippingOk() (*int32, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *CommerceOrder) SetShipping(v int32)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *CommerceOrder) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetTax

`func (o *CommerceOrder) GetTax() int32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *CommerceOrder) GetTaxOk() (*int32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *CommerceOrder) SetTax(v int32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *CommerceOrder) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTotal

`func (o *CommerceOrder) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CommerceOrder) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CommerceOrder) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CommerceOrder) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetBalance

`func (o *CommerceOrder) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CommerceOrder) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CommerceOrder) SetBalance(v int32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *CommerceOrder) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetPaid

`func (o *CommerceOrder) GetPaid() int32`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *CommerceOrder) GetPaidOk() (*int32, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *CommerceOrder) SetPaid(v int32)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *CommerceOrder) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetRefunded

`func (o *CommerceOrder) GetRefunded() int32`

GetRefunded returns the Refunded field if non-nil, zero value otherwise.

### GetRefundedOk

`func (o *CommerceOrder) GetRefundedOk() (*int32, bool)`

GetRefundedOk returns a tuple with the Refunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunded

`func (o *CommerceOrder) SetRefunded(v int32)`

SetRefunded sets Refunded field to given value.

### HasRefunded

`func (o *CommerceOrder) HasRefunded() bool`

HasRefunded returns a boolean if a field has been set.

### GetBillingAddress

`func (o *CommerceOrder) GetBillingAddress() CommerceAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *CommerceOrder) GetBillingAddressOk() (*CommerceAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *CommerceOrder) SetBillingAddress(v CommerceAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *CommerceOrder) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *CommerceOrder) GetShippingAddress() CommerceAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *CommerceOrder) GetShippingAddressOk() (*CommerceAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *CommerceOrder) SetShippingAddress(v CommerceAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *CommerceOrder) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetItems

`func (o *CommerceOrder) GetItems() []CommerceLineItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CommerceOrder) GetItemsOk() (*[]CommerceLineItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CommerceOrder) SetItems(v []CommerceLineItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *CommerceOrder) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetCoupons

`func (o *CommerceOrder) GetCoupons() []CommerceCoupon`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *CommerceOrder) GetCouponsOk() (*[]CommerceCoupon, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupons

`func (o *CommerceOrder) SetCoupons(v []CommerceCoupon)`

SetCoupons sets Coupons field to given value.

### HasCoupons

`func (o *CommerceOrder) HasCoupons() bool`

HasCoupons returns a boolean if a field has been set.

### GetCouponCodes

`func (o *CommerceOrder) GetCouponCodes() []string`

GetCouponCodes returns the CouponCodes field if non-nil, zero value otherwise.

### GetCouponCodesOk

`func (o *CommerceOrder) GetCouponCodesOk() (*[]string, bool)`

GetCouponCodesOk returns a tuple with the CouponCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCodes

`func (o *CommerceOrder) SetCouponCodes(v []string)`

SetCouponCodes sets CouponCodes field to given value.

### HasCouponCodes

`func (o *CommerceOrder) HasCouponCodes() bool`

HasCouponCodes returns a boolean if a field has been set.

### GetFulfillment

`func (o *CommerceOrder) GetFulfillment() CommerceFulfillment`

GetFulfillment returns the Fulfillment field if non-nil, zero value otherwise.

### GetFulfillmentOk

`func (o *CommerceOrder) GetFulfillmentOk() (*CommerceFulfillment, bool)`

GetFulfillmentOk returns a tuple with the Fulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillment

`func (o *CommerceOrder) SetFulfillment(v CommerceFulfillment)`

SetFulfillment sets Fulfillment field to given value.

### HasFulfillment

`func (o *CommerceOrder) HasFulfillment() bool`

HasFulfillment returns a boolean if a field has been set.

### GetGift

`func (o *CommerceOrder) GetGift() bool`

GetGift returns the Gift field if non-nil, zero value otherwise.

### GetGiftOk

`func (o *CommerceOrder) GetGiftOk() (*bool, bool)`

GetGiftOk returns a tuple with the Gift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGift

`func (o *CommerceOrder) SetGift(v bool)`

SetGift sets Gift field to given value.

### HasGift

`func (o *CommerceOrder) HasGift() bool`

HasGift returns a boolean if a field has been set.

### GetGiftMessage

`func (o *CommerceOrder) GetGiftMessage() string`

GetGiftMessage returns the GiftMessage field if non-nil, zero value otherwise.

### GetGiftMessageOk

`func (o *CommerceOrder) GetGiftMessageOk() (*string, bool)`

GetGiftMessageOk returns a tuple with the GiftMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftMessage

`func (o *CommerceOrder) SetGiftMessage(v string)`

SetGiftMessage sets GiftMessage field to given value.

### HasGiftMessage

`func (o *CommerceOrder) HasGiftMessage() bool`

HasGiftMessage returns a boolean if a field has been set.

### GetGiftEmail

`func (o *CommerceOrder) GetGiftEmail() string`

GetGiftEmail returns the GiftEmail field if non-nil, zero value otherwise.

### GetGiftEmailOk

`func (o *CommerceOrder) GetGiftEmailOk() (*string, bool)`

GetGiftEmailOk returns a tuple with the GiftEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftEmail

`func (o *CommerceOrder) SetGiftEmail(v string)`

SetGiftEmail sets GiftEmail field to given value.

### HasGiftEmail

`func (o *CommerceOrder) HasGiftEmail() bool`

HasGiftEmail returns a boolean if a field has been set.

### GetMetadata

`func (o *CommerceOrder) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CommerceOrder) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CommerceOrder) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CommerceOrder) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTest

`func (o *CommerceOrder) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *CommerceOrder) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *CommerceOrder) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *CommerceOrder) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommerceOrder) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommerceOrder) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommerceOrder) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommerceOrder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CommerceOrder) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommerceOrder) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommerceOrder) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CommerceOrder) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



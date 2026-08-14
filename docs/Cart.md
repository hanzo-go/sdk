# Cart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is when the cart was opened, RFC3339. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO 4217 code every amount below is denominated in. | [optional] 
**DiscountCents** | Pointer to **int32** | DiscountCents is what coupons and promotions took off, in whole cents. | [optional] 
**Email** | Pointer to **string** | Email is the shopper&#39;s address, when the cart carries one. | [optional] 
**Id** | Pointer to **string** | ID is the cart&#39;s id — what every other cart op addresses it by, and what a storefront persists against the browser session. | [optional] 
**Items** | Pointer to [**[]CartItem**](CartItem.md) | Items are the cart&#39;s lines, in the order they were added. | [optional] 
**LineTotalCents** | Pointer to **int32** | LineTotalCents is the sum of the lines before any discount, in whole cents. | [optional] 
**Order** | Pointer to **string** | Order is the order this cart became, once checkout completed it. Empty until then, and its presence is what makes a cart final. | [optional] 
**ShippingCents** | Pointer to **int32** | ShippingCents is the shipping charge, in whole cents. It stays zero until a shipping option is priced at checkout. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;active\&quot; for a cart still being filled, \&quot;ordered\&quot; once checkout turned it into an order, and \&quot;discarded\&quot; when the shopper abandoned it. | [optional] 
**Store** | Pointer to **string** | Store is the storefront the cart is being filled on. | [optional] 
**SubtotalCents** | Pointer to **int32** | SubtotalCents is LineTotalCents less DiscountCents, in whole cents. | [optional] 
**TaxCents** | Pointer to **int32** | TaxCents is the sales tax, in whole cents. It stays zero until checkout resolves the shopper&#39;s tax region. | [optional] 
**TotalCents** | Pointer to **int32** | TotalCents is what the shopper pays: subtotal plus shipping plus tax, in whole cents. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when the cart was last amended, RFC3339. | [optional] 
**User** | Pointer to **string** | User is the signed-in shopper this cart belongs to, empty for a guest cart. | [optional] 

## Methods

### NewCart

`func NewCart() *Cart`

NewCart instantiates a new Cart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartWithDefaults

`func NewCartWithDefaults() *Cart`

NewCartWithDefaults instantiates a new Cart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Cart) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Cart) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Cart) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Cart) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *Cart) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Cart) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Cart) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Cart) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDiscountCents

`func (o *Cart) GetDiscountCents() int32`

GetDiscountCents returns the DiscountCents field if non-nil, zero value otherwise.

### GetDiscountCentsOk

`func (o *Cart) GetDiscountCentsOk() (*int32, bool)`

GetDiscountCentsOk returns a tuple with the DiscountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCents

`func (o *Cart) SetDiscountCents(v int32)`

SetDiscountCents sets DiscountCents field to given value.

### HasDiscountCents

`func (o *Cart) HasDiscountCents() bool`

HasDiscountCents returns a boolean if a field has been set.

### GetEmail

`func (o *Cart) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Cart) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Cart) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Cart) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *Cart) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cart) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cart) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Cart) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *Cart) GetItems() []CartItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Cart) GetItemsOk() (*[]CartItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Cart) SetItems(v []CartItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Cart) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLineTotalCents

`func (o *Cart) GetLineTotalCents() int32`

GetLineTotalCents returns the LineTotalCents field if non-nil, zero value otherwise.

### GetLineTotalCentsOk

`func (o *Cart) GetLineTotalCentsOk() (*int32, bool)`

GetLineTotalCentsOk returns a tuple with the LineTotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineTotalCents

`func (o *Cart) SetLineTotalCents(v int32)`

SetLineTotalCents sets LineTotalCents field to given value.

### HasLineTotalCents

`func (o *Cart) HasLineTotalCents() bool`

HasLineTotalCents returns a boolean if a field has been set.

### GetOrder

`func (o *Cart) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Cart) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Cart) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Cart) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetShippingCents

`func (o *Cart) GetShippingCents() int32`

GetShippingCents returns the ShippingCents field if non-nil, zero value otherwise.

### GetShippingCentsOk

`func (o *Cart) GetShippingCentsOk() (*int32, bool)`

GetShippingCentsOk returns a tuple with the ShippingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCents

`func (o *Cart) SetShippingCents(v int32)`

SetShippingCents sets ShippingCents field to given value.

### HasShippingCents

`func (o *Cart) HasShippingCents() bool`

HasShippingCents returns a boolean if a field has been set.

### GetStatus

`func (o *Cart) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cart) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cart) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Cart) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStore

`func (o *Cart) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *Cart) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *Cart) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *Cart) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetSubtotalCents

`func (o *Cart) GetSubtotalCents() int32`

GetSubtotalCents returns the SubtotalCents field if non-nil, zero value otherwise.

### GetSubtotalCentsOk

`func (o *Cart) GetSubtotalCentsOk() (*int32, bool)`

GetSubtotalCentsOk returns a tuple with the SubtotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalCents

`func (o *Cart) SetSubtotalCents(v int32)`

SetSubtotalCents sets SubtotalCents field to given value.

### HasSubtotalCents

`func (o *Cart) HasSubtotalCents() bool`

HasSubtotalCents returns a boolean if a field has been set.

### GetTaxCents

`func (o *Cart) GetTaxCents() int32`

GetTaxCents returns the TaxCents field if non-nil, zero value otherwise.

### GetTaxCentsOk

`func (o *Cart) GetTaxCentsOk() (*int32, bool)`

GetTaxCentsOk returns a tuple with the TaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCents

`func (o *Cart) SetTaxCents(v int32)`

SetTaxCents sets TaxCents field to given value.

### HasTaxCents

`func (o *Cart) HasTaxCents() bool`

HasTaxCents returns a boolean if a field has been set.

### GetTotalCents

`func (o *Cart) GetTotalCents() int32`

GetTotalCents returns the TotalCents field if non-nil, zero value otherwise.

### GetTotalCentsOk

`func (o *Cart) GetTotalCentsOk() (*int32, bool)`

GetTotalCentsOk returns a tuple with the TotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCents

`func (o *Cart) SetTotalCents(v int32)`

SetTotalCents sets TotalCents field to given value.

### HasTotalCents

`func (o *Cart) HasTotalCents() bool`

HasTotalCents returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Cart) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Cart) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Cart) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Cart) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *Cart) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Cart) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Cart) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Cart) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



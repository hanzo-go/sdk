# CartOpen

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Currency is the ISO 4217 code the cart is priced in, lower-cased. Empty means usd. | [optional] 
**Email** | Pointer to **string** | Email is the shopper&#39;s address, for a cart that belongs to someone who has not signed in. It is what a guest checkout and an abandoned-cart follow-up key on. Empty is fine. | [optional] 
**Store** | Pointer to **string** | Store is the storefront this cart is being filled on. Empty uses the org&#39;s default store, which is what a single-storefront merchant always wants. | [optional] 
**User** | Pointer to **string** | User is the id of the signed-in shopper this cart belongs to, when there is one. Empty means a guest cart identified only by its own id. | [optional] 

## Methods

### NewCartOpen

`func NewCartOpen() *CartOpen`

NewCartOpen instantiates a new CartOpen object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartOpenWithDefaults

`func NewCartOpenWithDefaults() *CartOpen`

NewCartOpenWithDefaults instantiates a new CartOpen object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CartOpen) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CartOpen) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CartOpen) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CartOpen) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEmail

`func (o *CartOpen) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CartOpen) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CartOpen) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CartOpen) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetStore

`func (o *CartOpen) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *CartOpen) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *CartOpen) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *CartOpen) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetUser

`func (o *CartOpen) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CartOpen) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CartOpen) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CartOpen) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TopupIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** | AmountCents is how much to charge, in cents of Currency. Required. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO-4217 code to charge in. Empty takes the deployment&#39;s own default. | [optional] 
**PaymentMethodId** | Pointer to **string** | MethodID names a card the subject already saved, for the saved-card door. | [optional] 
**SourceId** | Pointer to **string** | SourceID is a single-use card token from the payment form, for the token door. It is vaulted as part of the charge, so a caller never holds card numbers and this service never sees one. | [optional] 

## Methods

### NewTopupIn

`func NewTopupIn() *TopupIn`

NewTopupIn instantiates a new TopupIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopupInWithDefaults

`func NewTopupInWithDefaults() *TopupIn`

NewTopupInWithDefaults instantiates a new TopupIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *TopupIn) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *TopupIn) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *TopupIn) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *TopupIn) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetCurrency

`func (o *TopupIn) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TopupIn) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TopupIn) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TopupIn) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *TopupIn) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *TopupIn) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *TopupIn) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *TopupIn) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetSourceId

`func (o *TopupIn) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *TopupIn) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *TopupIn) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *TopupIn) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PaymentIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** | AmountCents is the amount to charge, in whole cents (5000 is $50.00). Server-side bounds apply and are authoritative — the default floor is $1 and the ceiling $5,000, so a fat-fingered or hostile amount is refused before any money moves. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO 4217 code, lower-cased. Empty means usd. | [optional] 
**IdempotencyKey** | Pointer to **string** | IdempotencyKey makes a retry safe: the same key never charges twice, it replays the first result. Sending one is strongly recommended for an agent, which retries by construction. Empty falls back to a windowed key derived from the amount and currency, so a double-submit inside 15 minutes still collapses onto one charge. | [optional] 
**SourceId** | Pointer to **string** | SourceID is the single-use payment token that stands in for the card: a Square Web Payments SDK nonce minted in the browser, or a Square sandbox test nonce when the org&#39;s credentials are sandbox ones. The card number itself never reaches this process, which is what keeps it out of PCI scope. | [optional] 

## Methods

### NewPaymentIn

`func NewPaymentIn() *PaymentIn`

NewPaymentIn instantiates a new PaymentIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInWithDefaults

`func NewPaymentInWithDefaults() *PaymentIn`

NewPaymentInWithDefaults instantiates a new PaymentIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *PaymentIn) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *PaymentIn) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *PaymentIn) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *PaymentIn) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentIn) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentIn) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentIn) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentIn) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *PaymentIn) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *PaymentIn) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *PaymentIn) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *PaymentIn) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetSourceId

`func (o *PaymentIn) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *PaymentIn) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *PaymentIn) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *PaymentIn) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



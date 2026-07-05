# AdminCreditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | **int64** | Positive credit to add (capped at 10,000,000 cents per grant) | 
**Currency** | Pointer to **string** |  | [optional] [default to "usd"]
**Reason** | Pointer to **string** | Operator justification (recorded in the audit trail) | [optional] 

## Methods

### NewAdminCreditRequest

`func NewAdminCreditRequest(amountCents int64, ) *AdminCreditRequest`

NewAdminCreditRequest instantiates a new AdminCreditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminCreditRequestWithDefaults

`func NewAdminCreditRequestWithDefaults() *AdminCreditRequest`

NewAdminCreditRequestWithDefaults instantiates a new AdminCreditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *AdminCreditRequest) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *AdminCreditRequest) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *AdminCreditRequest) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.


### GetCurrency

`func (o *AdminCreditRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AdminCreditRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AdminCreditRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AdminCreditRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetReason

`func (o *AdminCreditRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AdminCreditRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AdminCreditRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AdminCreditRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AffiliatesPayoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | **int64** | Payout amount in USD cents (must be positive; cannot exceed pending commission). | 
**Method** | **string** | Payout method. &#x60;credits&#x60; issues a commerce grant; any other value (wire/paypal/check/…) is record-only. | 
**Reference** | Pointer to **string** | Optional external reference for a cash payout. | [optional] 

## Methods

### NewAffiliatesPayoutRequest

`func NewAffiliatesPayoutRequest(amountCents int64, method string, ) *AffiliatesPayoutRequest`

NewAffiliatesPayoutRequest instantiates a new AffiliatesPayoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliatesPayoutRequestWithDefaults

`func NewAffiliatesPayoutRequestWithDefaults() *AffiliatesPayoutRequest`

NewAffiliatesPayoutRequestWithDefaults instantiates a new AffiliatesPayoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *AffiliatesPayoutRequest) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *AffiliatesPayoutRequest) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *AffiliatesPayoutRequest) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.


### GetMethod

`func (o *AffiliatesPayoutRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AffiliatesPayoutRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AffiliatesPayoutRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetReference

`func (o *AffiliatesPayoutRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AffiliatesPayoutRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AffiliatesPayoutRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AffiliatesPayoutRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



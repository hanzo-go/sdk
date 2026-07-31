# AuthorsPayoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | **int64** | Payout amount in USD minor units (cents). Must be positive and not exceed pending. | 
**Method** | **string** | Payout method. &#39;credits&#39; issues a commerce grant; any other value (wire, paypal, check, …) is record-only. | 
**Reference** | Pointer to **string** | Optional external reference for a cash disbursement. | [optional] 

## Methods

### NewAuthorsPayoutRequest

`func NewAuthorsPayoutRequest(amountCents int64, method string, ) *AuthorsPayoutRequest`

NewAuthorsPayoutRequest instantiates a new AuthorsPayoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsPayoutRequestWithDefaults

`func NewAuthorsPayoutRequestWithDefaults() *AuthorsPayoutRequest`

NewAuthorsPayoutRequestWithDefaults instantiates a new AuthorsPayoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *AuthorsPayoutRequest) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *AuthorsPayoutRequest) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *AuthorsPayoutRequest) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.


### GetMethod

`func (o *AuthorsPayoutRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AuthorsPayoutRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AuthorsPayoutRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetReference

`func (o *AuthorsPayoutRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AuthorsPayoutRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AuthorsPayoutRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AuthorsPayoutRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



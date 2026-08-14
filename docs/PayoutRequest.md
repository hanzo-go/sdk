# PayoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** | AmountCents is how much to pay, in cents. Must be positive and can never exceed the author&#39;s pending royalty (accrued minus paid). | [optional] 
**Id** | Pointer to **string** | ID is the author to pay, from the path. | [optional] 
**Method** | Pointer to **string** | Method is how it settles: \&quot;credits\&quot; issues a grant into the author&#39;s wallet; wire, paypal and the like are record-only. Required. | [optional] 
**Reference** | Pointer to **string** | Reference is the operator&#39;s external reference for a cash settlement — a wire confirmation, a PayPal transaction id. | [optional] 

## Methods

### NewPayoutRequest

`func NewPayoutRequest() *PayoutRequest`

NewPayoutRequest instantiates a new PayoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutRequestWithDefaults

`func NewPayoutRequestWithDefaults() *PayoutRequest`

NewPayoutRequestWithDefaults instantiates a new PayoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *PayoutRequest) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *PayoutRequest) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *PayoutRequest) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *PayoutRequest) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetId

`func (o *PayoutRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PayoutRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PayoutRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PayoutRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *PayoutRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PayoutRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PayoutRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PayoutRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetReference

`func (o *PayoutRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PayoutRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PayoutRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PayoutRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



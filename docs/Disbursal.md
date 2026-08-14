# Disbursal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** | AmountCents is the payout, integer cents; it must be positive and can never exceed the affiliate&#39;s pending commission. Body-only (&#x60;url:\&quot;-\&quot;&#x60;, like every money field here): a payout must never ride the URL into access logs, and the raw handler read only the body. | [optional] 
**Id** | Pointer to **string** | ID is the affiliate to pay, from the path. | [optional] 
**Method** | Pointer to **string** | Method decides whether money moves: &#x60;credits&#x60; issues a commerce grant, every other method (wire, paypal, …) is record-only. | [optional] 
**Reference** | Pointer to **string** | Reference is the operator&#39;s settlement note (a bank id, a ledger ref). | [optional] 

## Methods

### NewDisbursal

`func NewDisbursal() *Disbursal`

NewDisbursal instantiates a new Disbursal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisbursalWithDefaults

`func NewDisbursalWithDefaults() *Disbursal`

NewDisbursalWithDefaults instantiates a new Disbursal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *Disbursal) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *Disbursal) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *Disbursal) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *Disbursal) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetId

`func (o *Disbursal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Disbursal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Disbursal) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Disbursal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *Disbursal) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Disbursal) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Disbursal) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Disbursal) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetReference

`func (o *Disbursal) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Disbursal) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Disbursal) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *Disbursal) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



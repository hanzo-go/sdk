# GrantIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** | AmountCents is the credit, in whole cents. Must be positive and within the per-grant cap. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO code, lower-cased. Empty means usd. | [optional] 
**Org** | Pointer to **string** | Org is the tenant to credit. Required. | [optional] 
**Reason** | Pointer to **string** | Reason is the operator&#39;s justification, recorded on the audit row. | [optional] 
**Source** | Pointer to **string** | Source is the money bucket: \&quot;trial\&quot; (default) for a non-cash comp that is never refundable, or \&quot;prepaid\&quot; for real money. Anything unknown falls back to trial. | [optional] 
**User** | Pointer to **string** | User optionally names a MEMBER to credit, by bare IAM username. Empty credits the org. Which of the two the money actually lands on is decided by account.Payer, not here: a pooled org keeps one balance whatever is named. | [optional] 

## Methods

### NewGrantIn

`func NewGrantIn() *GrantIn`

NewGrantIn instantiates a new GrantIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantInWithDefaults

`func NewGrantInWithDefaults() *GrantIn`

NewGrantInWithDefaults instantiates a new GrantIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *GrantIn) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *GrantIn) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *GrantIn) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *GrantIn) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetCurrency

`func (o *GrantIn) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GrantIn) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GrantIn) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GrantIn) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetOrg

`func (o *GrantIn) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *GrantIn) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *GrantIn) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *GrantIn) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetReason

`func (o *GrantIn) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GrantIn) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GrantIn) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GrantIn) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSource

`func (o *GrantIn) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GrantIn) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GrantIn) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GrantIn) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUser

`func (o *GrantIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GrantIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GrantIn) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *GrantIn) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



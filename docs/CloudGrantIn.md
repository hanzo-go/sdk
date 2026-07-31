# CloudGrantIn

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

### NewCloudGrantIn

`func NewCloudGrantIn() *CloudGrantIn`

NewCloudGrantIn instantiates a new CloudGrantIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGrantInWithDefaults

`func NewCloudGrantInWithDefaults() *CloudGrantIn`

NewCloudGrantInWithDefaults instantiates a new CloudGrantIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *CloudGrantIn) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CloudGrantIn) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CloudGrantIn) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *CloudGrantIn) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetCurrency

`func (o *CloudGrantIn) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudGrantIn) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudGrantIn) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudGrantIn) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetOrg

`func (o *CloudGrantIn) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudGrantIn) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudGrantIn) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudGrantIn) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetReason

`func (o *CloudGrantIn) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CloudGrantIn) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CloudGrantIn) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CloudGrantIn) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSource

`func (o *CloudGrantIn) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudGrantIn) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudGrantIn) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudGrantIn) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUser

`func (o *CloudGrantIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CloudGrantIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CloudGrantIn) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CloudGrantIn) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



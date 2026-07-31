# CloudGrantResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceCents** | Pointer to **int32** | BalanceCents is the account balance AFTER the grant, in whole cents. | [optional] 
**BalanceExact** | Pointer to **string** | BalanceExact is that same balance at full 18-decimal precision, so a sub-cent debit is visible rather than rounded away. | [optional] 
**Currency** | Pointer to **string** | Currency is the lower-cased ISO code the grant was denominated in. | [optional] 
**GrantedCents** | Pointer to **int32** | GrantedCents is the amount actually credited. | [optional] 
**Org** | Pointer to **string** | Org is the tenant whose ledger was credited. | [optional] 
**Source** | Pointer to **string** | Source is the money bucket: \&quot;trial\&quot; (non-cash comp) or \&quot;prepaid\&quot; (real money). | [optional] 
**Subject** | Pointer to **string** | Subject is the ACCOUNT the credit landed on inside that ledger: the org slug for a pooled org, \&quot;&lt;org&gt;/&lt;name&gt;\&quot; for a member of a per-member one. It is echoed because the operator does not choose it — account.Payer does — so naming a member of a pooled org credits the pool and the receipt has to say so. | [optional] 
**TransactionId** | Pointer to **string** | TransactionID is the ledger entry id, for reconciliation against commerce. | [optional] 

## Methods

### NewCloudGrantResult

`func NewCloudGrantResult() *CloudGrantResult`

NewCloudGrantResult instantiates a new CloudGrantResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGrantResultWithDefaults

`func NewCloudGrantResultWithDefaults() *CloudGrantResult`

NewCloudGrantResultWithDefaults instantiates a new CloudGrantResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceCents

`func (o *CloudGrantResult) GetBalanceCents() int32`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *CloudGrantResult) GetBalanceCentsOk() (*int32, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *CloudGrantResult) SetBalanceCents(v int32)`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *CloudGrantResult) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### GetBalanceExact

`func (o *CloudGrantResult) GetBalanceExact() string`

GetBalanceExact returns the BalanceExact field if non-nil, zero value otherwise.

### GetBalanceExactOk

`func (o *CloudGrantResult) GetBalanceExactOk() (*string, bool)`

GetBalanceExactOk returns a tuple with the BalanceExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceExact

`func (o *CloudGrantResult) SetBalanceExact(v string)`

SetBalanceExact sets BalanceExact field to given value.

### HasBalanceExact

`func (o *CloudGrantResult) HasBalanceExact() bool`

HasBalanceExact returns a boolean if a field has been set.

### GetCurrency

`func (o *CloudGrantResult) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudGrantResult) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudGrantResult) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudGrantResult) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGrantedCents

`func (o *CloudGrantResult) GetGrantedCents() int32`

GetGrantedCents returns the GrantedCents field if non-nil, zero value otherwise.

### GetGrantedCentsOk

`func (o *CloudGrantResult) GetGrantedCentsOk() (*int32, bool)`

GetGrantedCentsOk returns a tuple with the GrantedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedCents

`func (o *CloudGrantResult) SetGrantedCents(v int32)`

SetGrantedCents sets GrantedCents field to given value.

### HasGrantedCents

`func (o *CloudGrantResult) HasGrantedCents() bool`

HasGrantedCents returns a boolean if a field has been set.

### GetOrg

`func (o *CloudGrantResult) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudGrantResult) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudGrantResult) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudGrantResult) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSource

`func (o *CloudGrantResult) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudGrantResult) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudGrantResult) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudGrantResult) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSubject

`func (o *CloudGrantResult) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CloudGrantResult) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CloudGrantResult) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CloudGrantResult) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTransactionId

`func (o *CloudGrantResult) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CloudGrantResult) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CloudGrantResult) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *CloudGrantResult) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



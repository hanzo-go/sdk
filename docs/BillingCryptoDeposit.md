# BillingCryptoDeposit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressTag** | Pointer to **string** | AddressTag is the destination tag the payment must carry, on a chain whose address is shared. Absent where the address alone names the payer. | [optional] 
**AmountCents** | Pointer to **int64** | AmountCents is what the payer said they would send, in USD cents, when the address was asked for: a record of intent, not a quote. What is credited is what arrives. Absent when nothing was named. | [optional] 
**Chain** | Pointer to **string** | Chain is the network to send on. | [optional] 
**Confirmations** | Pointer to **int64** | Confirmations is how deep the transfer being confirmed is now; zero until one is seen. | [optional] 
**CreditedCents** | Pointer to **int64** | CreditedCents is what the transfer this intent confirmed credited, in USD cents, once it succeeded. A later transfer to the same address is credited too, and shows in the balance rather than here. | [optional] 
**DepositAddress** | Pointer to **string** | DepositAddress is where to send it. | [optional] 
**ExpiresAt** | Pointer to **string** | ExpiresAt is when this address stops being handed out again (RFC3339). A transfer that arrives after it is still credited. | [optional] 
**FeeCents** | Pointer to **int64** | FeeCents is deducted from every deposit before it is credited: the cost of sweeping it out of the custody address. Absent deducts nothing. | [optional] 
**Id** | Pointer to **string** | ID names the intent; read it back at GET /v1/billing/crypto/deposit/{id}. | [optional] 
**Instructions** | Pointer to **string** | Instructions says, in one line, what to send where and when it is credited. | [optional] 
**MinimumCents** | Pointer to **int64** | MinimumCents is the smallest deposit, in USD cents, that credits anything; less credits nothing. Absent when nothing is watching this asset now, and then nothing sent to it is credited. | [optional] 
**RequiredConfirmations** | Pointer to **int64** | RequiredConfirmations is how many blocks deep the transfer must be before it is credited. | [optional] 
**Status** | Pointer to **string** | Status is pending until a transfer to the address is seen, confirming while the chain buries it, and succeeded once it is credited as cash in the payer&#39;s balance (GET /v1/billing/balance, cash). | [optional] 
**Token** | Pointer to **string** | Token is the asset to send; this intent credits that token alone. | [optional] 

## Methods

### NewBillingCryptoDeposit

`func NewBillingCryptoDeposit() *BillingCryptoDeposit`

NewBillingCryptoDeposit instantiates a new BillingCryptoDeposit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCryptoDepositWithDefaults

`func NewBillingCryptoDepositWithDefaults() *BillingCryptoDeposit`

NewBillingCryptoDepositWithDefaults instantiates a new BillingCryptoDeposit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressTag

`func (o *BillingCryptoDeposit) GetAddressTag() string`

GetAddressTag returns the AddressTag field if non-nil, zero value otherwise.

### GetAddressTagOk

`func (o *BillingCryptoDeposit) GetAddressTagOk() (*string, bool)`

GetAddressTagOk returns a tuple with the AddressTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressTag

`func (o *BillingCryptoDeposit) SetAddressTag(v string)`

SetAddressTag sets AddressTag field to given value.

### HasAddressTag

`func (o *BillingCryptoDeposit) HasAddressTag() bool`

HasAddressTag returns a boolean if a field has been set.

### GetAmountCents

`func (o *BillingCryptoDeposit) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *BillingCryptoDeposit) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *BillingCryptoDeposit) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *BillingCryptoDeposit) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetChain

`func (o *BillingCryptoDeposit) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *BillingCryptoDeposit) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *BillingCryptoDeposit) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *BillingCryptoDeposit) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetConfirmations

`func (o *BillingCryptoDeposit) GetConfirmations() int64`

GetConfirmations returns the Confirmations field if non-nil, zero value otherwise.

### GetConfirmationsOk

`func (o *BillingCryptoDeposit) GetConfirmationsOk() (*int64, bool)`

GetConfirmationsOk returns a tuple with the Confirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmations

`func (o *BillingCryptoDeposit) SetConfirmations(v int64)`

SetConfirmations sets Confirmations field to given value.

### HasConfirmations

`func (o *BillingCryptoDeposit) HasConfirmations() bool`

HasConfirmations returns a boolean if a field has been set.

### GetCreditedCents

`func (o *BillingCryptoDeposit) GetCreditedCents() int64`

GetCreditedCents returns the CreditedCents field if non-nil, zero value otherwise.

### GetCreditedCentsOk

`func (o *BillingCryptoDeposit) GetCreditedCentsOk() (*int64, bool)`

GetCreditedCentsOk returns a tuple with the CreditedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditedCents

`func (o *BillingCryptoDeposit) SetCreditedCents(v int64)`

SetCreditedCents sets CreditedCents field to given value.

### HasCreditedCents

`func (o *BillingCryptoDeposit) HasCreditedCents() bool`

HasCreditedCents returns a boolean if a field has been set.

### GetDepositAddress

`func (o *BillingCryptoDeposit) GetDepositAddress() string`

GetDepositAddress returns the DepositAddress field if non-nil, zero value otherwise.

### GetDepositAddressOk

`func (o *BillingCryptoDeposit) GetDepositAddressOk() (*string, bool)`

GetDepositAddressOk returns a tuple with the DepositAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositAddress

`func (o *BillingCryptoDeposit) SetDepositAddress(v string)`

SetDepositAddress sets DepositAddress field to given value.

### HasDepositAddress

`func (o *BillingCryptoDeposit) HasDepositAddress() bool`

HasDepositAddress returns a boolean if a field has been set.

### GetExpiresAt

`func (o *BillingCryptoDeposit) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *BillingCryptoDeposit) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *BillingCryptoDeposit) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *BillingCryptoDeposit) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFeeCents

`func (o *BillingCryptoDeposit) GetFeeCents() int64`

GetFeeCents returns the FeeCents field if non-nil, zero value otherwise.

### GetFeeCentsOk

`func (o *BillingCryptoDeposit) GetFeeCentsOk() (*int64, bool)`

GetFeeCentsOk returns a tuple with the FeeCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCents

`func (o *BillingCryptoDeposit) SetFeeCents(v int64)`

SetFeeCents sets FeeCents field to given value.

### HasFeeCents

`func (o *BillingCryptoDeposit) HasFeeCents() bool`

HasFeeCents returns a boolean if a field has been set.

### GetId

`func (o *BillingCryptoDeposit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingCryptoDeposit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingCryptoDeposit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingCryptoDeposit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstructions

`func (o *BillingCryptoDeposit) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *BillingCryptoDeposit) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *BillingCryptoDeposit) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *BillingCryptoDeposit) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetMinimumCents

`func (o *BillingCryptoDeposit) GetMinimumCents() int64`

GetMinimumCents returns the MinimumCents field if non-nil, zero value otherwise.

### GetMinimumCentsOk

`func (o *BillingCryptoDeposit) GetMinimumCentsOk() (*int64, bool)`

GetMinimumCentsOk returns a tuple with the MinimumCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCents

`func (o *BillingCryptoDeposit) SetMinimumCents(v int64)`

SetMinimumCents sets MinimumCents field to given value.

### HasMinimumCents

`func (o *BillingCryptoDeposit) HasMinimumCents() bool`

HasMinimumCents returns a boolean if a field has been set.

### GetRequiredConfirmations

`func (o *BillingCryptoDeposit) GetRequiredConfirmations() int64`

GetRequiredConfirmations returns the RequiredConfirmations field if non-nil, zero value otherwise.

### GetRequiredConfirmationsOk

`func (o *BillingCryptoDeposit) GetRequiredConfirmationsOk() (*int64, bool)`

GetRequiredConfirmationsOk returns a tuple with the RequiredConfirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredConfirmations

`func (o *BillingCryptoDeposit) SetRequiredConfirmations(v int64)`

SetRequiredConfirmations sets RequiredConfirmations field to given value.

### HasRequiredConfirmations

`func (o *BillingCryptoDeposit) HasRequiredConfirmations() bool`

HasRequiredConfirmations returns a boolean if a field has been set.

### GetStatus

`func (o *BillingCryptoDeposit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingCryptoDeposit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingCryptoDeposit) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingCryptoDeposit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *BillingCryptoDeposit) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BillingCryptoDeposit) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BillingCryptoDeposit) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *BillingCryptoDeposit) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



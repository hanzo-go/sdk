# BillingUsageReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the wallet the debit drew from, the same key GET /v1/billing/balance reports for this caller. | [optional] 
**Amount** | Pointer to [**BillingMoney**](BillingMoney.md) | Amount is the debit, exactly as recorded. | [optional] 
**Id** | Pointer to **string** | ID is the act&#39;s name, echoed. | [optional] 
**Org** | Pointer to **string** | Org is the organization whose ledger holds the debit. | [optional] 

## Methods

### NewBillingUsageReceipt

`func NewBillingUsageReceipt() *BillingUsageReceipt`

NewBillingUsageReceipt instantiates a new BillingUsageReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingUsageReceiptWithDefaults

`func NewBillingUsageReceiptWithDefaults() *BillingUsageReceipt`

NewBillingUsageReceiptWithDefaults instantiates a new BillingUsageReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *BillingUsageReceipt) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *BillingUsageReceipt) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *BillingUsageReceipt) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *BillingUsageReceipt) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAmount

`func (o *BillingUsageReceipt) GetAmount() BillingMoney`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingUsageReceipt) GetAmountOk() (*BillingMoney, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingUsageReceipt) SetAmount(v BillingMoney)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingUsageReceipt) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetId

`func (o *BillingUsageReceipt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingUsageReceipt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingUsageReceipt) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingUsageReceipt) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrg

`func (o *BillingUsageReceipt) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *BillingUsageReceipt) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *BillingUsageReceipt) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *BillingUsageReceipt) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



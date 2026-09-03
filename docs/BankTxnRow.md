# BankTxnRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int64** | AmountCents is the size of the movement in whole cents, always POSITIVE — direction carries the sign, so a caller must read both to know which way money went. | [optional] 
**Connector** | Pointer to **string** | Connector names the feed this row arrived on — which bank or processor connection it was synced from. With externalId it is the row&#39;s identity, so re-syncing the same statement never books a second copy. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO code the bank reported the line in. | [optional] 
**Description** | Pointer to **string** | Description is the statement memo as the bank wrote it. | [optional] 
**Direction** | Pointer to **string** | Direction is which way the money moved: an inflow into the account or an outflow from it, from the org&#39;s point of view. | [optional] 
**ExternalId** | Pointer to **string** | ExternalID is the bank&#39;s OWN id for the line, carried verbatim. It is unique only within its connector. | [optional] 
**MatchedVoucher** | Pointer to **string** | MatchedVoucher names the ledger voucher this line was reconciled against — the bill it paid, or the settlement it cleared. Absent when nothing matched, which for an inflow is what raises a question. | [optional] 
**Merchant** | Pointer to **string** | Merchant is the counterparty the feed identified, where it did. | [optional] 
**PostedAt** | Pointer to **string** | PostedAt is the bank&#39;s posting date for the line, not when we synced it. | [optional] 
**Status** | Pointer to **string** | Status is where the line got to: posted (an outflow booked straight to an expense), settled (an outflow that paid down a scanned bill), reconciled (an inflow that cleared a pending settlement), transfer (a move between the org&#39;s own accounts, recorded but with no effect on the books), or unmatched (an inflow nobody could place, which is waiting on a human answer). | [optional] 

## Methods

### NewBankTxnRow

`func NewBankTxnRow() *BankTxnRow`

NewBankTxnRow instantiates a new BankTxnRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTxnRowWithDefaults

`func NewBankTxnRowWithDefaults() *BankTxnRow`

NewBankTxnRowWithDefaults instantiates a new BankTxnRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *BankTxnRow) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *BankTxnRow) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *BankTxnRow) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *BankTxnRow) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetConnector

`func (o *BankTxnRow) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *BankTxnRow) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *BankTxnRow) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *BankTxnRow) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetCurrency

`func (o *BankTxnRow) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BankTxnRow) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BankTxnRow) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BankTxnRow) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *BankTxnRow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BankTxnRow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BankTxnRow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BankTxnRow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirection

`func (o *BankTxnRow) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *BankTxnRow) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *BankTxnRow) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *BankTxnRow) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetExternalId

`func (o *BankTxnRow) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BankTxnRow) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BankTxnRow) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *BankTxnRow) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMatchedVoucher

`func (o *BankTxnRow) GetMatchedVoucher() string`

GetMatchedVoucher returns the MatchedVoucher field if non-nil, zero value otherwise.

### GetMatchedVoucherOk

`func (o *BankTxnRow) GetMatchedVoucherOk() (*string, bool)`

GetMatchedVoucherOk returns a tuple with the MatchedVoucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedVoucher

`func (o *BankTxnRow) SetMatchedVoucher(v string)`

SetMatchedVoucher sets MatchedVoucher field to given value.

### HasMatchedVoucher

`func (o *BankTxnRow) HasMatchedVoucher() bool`

HasMatchedVoucher returns a boolean if a field has been set.

### GetMerchant

`func (o *BankTxnRow) GetMerchant() string`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *BankTxnRow) GetMerchantOk() (*string, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *BankTxnRow) SetMerchant(v string)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *BankTxnRow) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetPostedAt

`func (o *BankTxnRow) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *BankTxnRow) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *BankTxnRow) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.

### HasPostedAt

`func (o *BankTxnRow) HasPostedAt() bool`

HasPostedAt returns a boolean if a field has been set.

### GetStatus

`func (o *BankTxnRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BankTxnRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BankTxnRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BankTxnRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



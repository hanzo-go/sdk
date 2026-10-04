# BooksBankTxnRow

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

### NewBooksBankTxnRow

`func NewBooksBankTxnRow() *BooksBankTxnRow`

NewBooksBankTxnRow instantiates a new BooksBankTxnRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksBankTxnRowWithDefaults

`func NewBooksBankTxnRowWithDefaults() *BooksBankTxnRow`

NewBooksBankTxnRowWithDefaults instantiates a new BooksBankTxnRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *BooksBankTxnRow) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *BooksBankTxnRow) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *BooksBankTxnRow) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *BooksBankTxnRow) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetConnector

`func (o *BooksBankTxnRow) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *BooksBankTxnRow) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *BooksBankTxnRow) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *BooksBankTxnRow) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetCurrency

`func (o *BooksBankTxnRow) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BooksBankTxnRow) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BooksBankTxnRow) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BooksBankTxnRow) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *BooksBankTxnRow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BooksBankTxnRow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BooksBankTxnRow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BooksBankTxnRow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirection

`func (o *BooksBankTxnRow) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *BooksBankTxnRow) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *BooksBankTxnRow) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *BooksBankTxnRow) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetExternalId

`func (o *BooksBankTxnRow) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BooksBankTxnRow) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BooksBankTxnRow) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *BooksBankTxnRow) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMatchedVoucher

`func (o *BooksBankTxnRow) GetMatchedVoucher() string`

GetMatchedVoucher returns the MatchedVoucher field if non-nil, zero value otherwise.

### GetMatchedVoucherOk

`func (o *BooksBankTxnRow) GetMatchedVoucherOk() (*string, bool)`

GetMatchedVoucherOk returns a tuple with the MatchedVoucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedVoucher

`func (o *BooksBankTxnRow) SetMatchedVoucher(v string)`

SetMatchedVoucher sets MatchedVoucher field to given value.

### HasMatchedVoucher

`func (o *BooksBankTxnRow) HasMatchedVoucher() bool`

HasMatchedVoucher returns a boolean if a field has been set.

### GetMerchant

`func (o *BooksBankTxnRow) GetMerchant() string`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *BooksBankTxnRow) GetMerchantOk() (*string, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *BooksBankTxnRow) SetMerchant(v string)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *BooksBankTxnRow) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetPostedAt

`func (o *BooksBankTxnRow) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *BooksBankTxnRow) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *BooksBankTxnRow) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.

### HasPostedAt

`func (o *BooksBankTxnRow) HasPostedAt() bool`

HasPostedAt returns a boolean if a field has been set.

### GetStatus

`func (o *BooksBankTxnRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BooksBankTxnRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BooksBankTxnRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BooksBankTxnRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



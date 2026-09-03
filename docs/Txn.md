# Txn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int64** | AmountCents is the voucher&#39;s total, in whole cents — its total debit, which equals its total credit because every voucher balances. It is the size of the entry and carries no direction; the category says which way it went. | [optional] 
**Category** | Pointer to **string** | Category is the chart-of-accounts NUMBER of the income or expense account this voucher touched — where it lands on the P&amp;L, not a free-text label. | [optional] 
**CategoryName** | Pointer to **string** | CategoryName is that account&#39;s human name, so a caller need not carry the chart to render the row. | [optional] 
**Date** | Pointer to **string** | Date is when the voucher POSTED — the accounting date the reports window on, which for an imported bank row is the bank&#39;s date and not the day it landed here. | [optional] 
**Description** | Pointer to **string** | Description is the line a person reads: the memo carried in from the source. | [optional] 
**Source** | Pointer to **string** | Source is where the entry came from: bank_txn for an imported statement line, scan for a receipt or bill read by the scanner, commerce_txn for a sale booked by the store. | [optional] 
**Vendor** | Pointer to **string** | Vendor is the counterparty, resolved from whatever the source knew — a bank row&#39;s merchant, a scanned bill&#39;s supplier. Absent when the source named none. | [optional] 
**VoucherId** | Pointer to **int64** | VoucherID identifies the underlying double-entry voucher, so a caller can open the full set of legs behind this single register line. | [optional] 

## Methods

### NewTxn

`func NewTxn() *Txn`

NewTxn instantiates a new Txn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTxnWithDefaults

`func NewTxnWithDefaults() *Txn`

NewTxnWithDefaults instantiates a new Txn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *Txn) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *Txn) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *Txn) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *Txn) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetCategory

`func (o *Txn) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Txn) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Txn) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Txn) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCategoryName

`func (o *Txn) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *Txn) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *Txn) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *Txn) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.

### GetDate

`func (o *Txn) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Txn) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Txn) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Txn) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDescription

`func (o *Txn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Txn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Txn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Txn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSource

`func (o *Txn) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Txn) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Txn) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Txn) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVendor

`func (o *Txn) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Txn) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Txn) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *Txn) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVoucherId

`func (o *Txn) GetVoucherId() int64`

GetVoucherId returns the VoucherId field if non-nil, zero value otherwise.

### GetVoucherIdOk

`func (o *Txn) GetVoucherIdOk() (*int64, bool)`

GetVoucherIdOk returns a tuple with the VoucherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherId

`func (o *Txn) SetVoucherId(v int64)`

SetVoucherId sets VoucherId field to given value.

### HasVoucherId

`func (o *Txn) HasVoucherId() bool`

HasVoucherId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



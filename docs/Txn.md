# Txn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** |  | [optional] 
**Category** | Pointer to **string** | COA account number of the P&amp;L line | [optional] 
**CategoryName** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** | source_kind: bank_txn | scan | commerce_txn | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**VoucherId** | Pointer to **int32** |  | [optional] 

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

`func (o *Txn) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *Txn) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *Txn) SetAmountCents(v int32)`

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

`func (o *Txn) GetVoucherId() int32`

GetVoucherId returns the VoucherId field if non-nil, zero value otherwise.

### GetVoucherIdOk

`func (o *Txn) GetVoucherIdOk() (*int32, bool)`

GetVoucherIdOk returns a tuple with the VoucherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherId

`func (o *Txn) SetVoucherId(v int32)`

SetVoucherId sets VoucherId field to given value.

### HasVoucherId

`func (o *Txn) HasVoucherId() bool`

HasVoucherId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TaxFiling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is when it was exported, unix seconds. | [optional] 
**Due** | Pointer to **string** | Due is when the return is due to the IRS. | [optional] 
**Forms** | Pointer to **[]string** | Forms are the form ids the filing carries. | [optional] 
**Id** | Pointer to **string** | ID is the filing&#39;s id, \&quot;iris_\&quot;-prefixed. | [optional] 
**Kind** | Pointer to **string** | Kind is 1099-NEC or 1099-MISC — one per filing, as IRIS takes them. | [optional] 
**Layout** | Pointer to **string** | Layout says what the files are. | [optional] 
**MissingTin** | Pointer to **[]string** | MissingTIN are forms filed with no recipient TIN. | [optional] 
**Parts** | Pointer to [**[]TaxPart**](TaxPart.md) | Parts are its files. | [optional] 
**ReceiptId** | Pointer to **string** | ReceiptID is the Receipt ID IRIS returned, once the payer recorded one. | [optional] 
**Status** | Pointer to **string** | Status is exported, submitted, accepted, accepted_with_errors or rejected. Only the payer&#39;s own record moves it past exported: nothing here submits. | [optional] 
**Steps** | Pointer to **[]string** | Steps are what remains for the payer to do, in order. | [optional] 
**Tcc** | Pointer to **string** | TCC is the payer&#39;s IRIS Transmitter Control Code, once it recorded a submission. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when its status last moved, unix seconds. | [optional] 
**Year** | Pointer to **int64** | Year is the tax year. | [optional] 

## Methods

### NewTaxFiling

`func NewTaxFiling() *TaxFiling`

NewTaxFiling instantiates a new TaxFiling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxFilingWithDefaults

`func NewTaxFilingWithDefaults() *TaxFiling`

NewTaxFilingWithDefaults instantiates a new TaxFiling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TaxFiling) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaxFiling) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaxFiling) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaxFiling) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDue

`func (o *TaxFiling) GetDue() string`

GetDue returns the Due field if non-nil, zero value otherwise.

### GetDueOk

`func (o *TaxFiling) GetDueOk() (*string, bool)`

GetDueOk returns a tuple with the Due field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDue

`func (o *TaxFiling) SetDue(v string)`

SetDue sets Due field to given value.

### HasDue

`func (o *TaxFiling) HasDue() bool`

HasDue returns a boolean if a field has been set.

### GetForms

`func (o *TaxFiling) GetForms() []string`

GetForms returns the Forms field if non-nil, zero value otherwise.

### GetFormsOk

`func (o *TaxFiling) GetFormsOk() (*[]string, bool)`

GetFormsOk returns a tuple with the Forms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForms

`func (o *TaxFiling) SetForms(v []string)`

SetForms sets Forms field to given value.

### HasForms

`func (o *TaxFiling) HasForms() bool`

HasForms returns a boolean if a field has been set.

### GetId

`func (o *TaxFiling) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxFiling) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxFiling) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxFiling) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *TaxFiling) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TaxFiling) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TaxFiling) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TaxFiling) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLayout

`func (o *TaxFiling) GetLayout() string`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *TaxFiling) GetLayoutOk() (*string, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *TaxFiling) SetLayout(v string)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *TaxFiling) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetMissingTin

`func (o *TaxFiling) GetMissingTin() []string`

GetMissingTin returns the MissingTin field if non-nil, zero value otherwise.

### GetMissingTinOk

`func (o *TaxFiling) GetMissingTinOk() (*[]string, bool)`

GetMissingTinOk returns a tuple with the MissingTin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingTin

`func (o *TaxFiling) SetMissingTin(v []string)`

SetMissingTin sets MissingTin field to given value.

### HasMissingTin

`func (o *TaxFiling) HasMissingTin() bool`

HasMissingTin returns a boolean if a field has been set.

### GetParts

`func (o *TaxFiling) GetParts() []TaxPart`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *TaxFiling) GetPartsOk() (*[]TaxPart, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *TaxFiling) SetParts(v []TaxPart)`

SetParts sets Parts field to given value.

### HasParts

`func (o *TaxFiling) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetReceiptId

`func (o *TaxFiling) GetReceiptId() string`

GetReceiptId returns the ReceiptId field if non-nil, zero value otherwise.

### GetReceiptIdOk

`func (o *TaxFiling) GetReceiptIdOk() (*string, bool)`

GetReceiptIdOk returns a tuple with the ReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptId

`func (o *TaxFiling) SetReceiptId(v string)`

SetReceiptId sets ReceiptId field to given value.

### HasReceiptId

`func (o *TaxFiling) HasReceiptId() bool`

HasReceiptId returns a boolean if a field has been set.

### GetStatus

`func (o *TaxFiling) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaxFiling) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaxFiling) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaxFiling) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSteps

`func (o *TaxFiling) GetSteps() []string`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *TaxFiling) GetStepsOk() (*[]string, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *TaxFiling) SetSteps(v []string)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *TaxFiling) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTcc

`func (o *TaxFiling) GetTcc() string`

GetTcc returns the Tcc field if non-nil, zero value otherwise.

### GetTccOk

`func (o *TaxFiling) GetTccOk() (*string, bool)`

GetTccOk returns a tuple with the Tcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcc

`func (o *TaxFiling) SetTcc(v string)`

SetTcc sets Tcc field to given value.

### HasTcc

`func (o *TaxFiling) HasTcc() bool`

HasTcc returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaxFiling) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaxFiling) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaxFiling) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TaxFiling) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetYear

`func (o *TaxFiling) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *TaxFiling) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *TaxFiling) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *TaxFiling) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



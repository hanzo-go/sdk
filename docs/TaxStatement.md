# TaxStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the account number the form prints. | [optional] 
**Backup** | Pointer to [**TaxWithholding**](TaxWithholding.md) | Backup is the backup withholding the payer applied. | [optional] 
**Boxes** | Pointer to [**[]TaxAmount**](TaxAmount.md) | Boxes are the figures. | [optional] 
**Corrected** | Pointer to **bool** | Corrected is true for a statement that corrects an earlier one. | [optional] 
**FurnishedAt** | Pointer to **int64** | FurnishedAt is when it was delivered, unix seconds. | [optional] 
**Id** | Pointer to **string** | ID is the form id the payer issued it under. | [optional] 
**Kind** | Pointer to **string** | Kind is 1099-NEC or 1099-MISC. | [optional] 
**Payer** | Pointer to [**TaxParty**](TaxParty.md) | Payer is the org that furnished it. | [optional] 
**Recipient** | Pointer to [**TaxParty**](TaxParty.md) | Recipient is this org, as the form names it, TIN truncated. | [optional] 
**SupersededBy** | Pointer to **string** | SupersededBy is the later statement that corrects this one. | [optional] 
**Supersedes** | Pointer to **string** | Supersedes is the earlier statement this corrects. | [optional] 
**Year** | Pointer to **int64** | Year is the calendar year of the payments. | [optional] 

## Methods

### NewTaxStatement

`func NewTaxStatement() *TaxStatement`

NewTaxStatement instantiates a new TaxStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxStatementWithDefaults

`func NewTaxStatementWithDefaults() *TaxStatement`

NewTaxStatementWithDefaults instantiates a new TaxStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *TaxStatement) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TaxStatement) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TaxStatement) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TaxStatement) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBackup

`func (o *TaxStatement) GetBackup() TaxWithholding`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *TaxStatement) GetBackupOk() (*TaxWithholding, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *TaxStatement) SetBackup(v TaxWithholding)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *TaxStatement) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetBoxes

`func (o *TaxStatement) GetBoxes() []TaxAmount`

GetBoxes returns the Boxes field if non-nil, zero value otherwise.

### GetBoxesOk

`func (o *TaxStatement) GetBoxesOk() (*[]TaxAmount, bool)`

GetBoxesOk returns a tuple with the Boxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxes

`func (o *TaxStatement) SetBoxes(v []TaxAmount)`

SetBoxes sets Boxes field to given value.

### HasBoxes

`func (o *TaxStatement) HasBoxes() bool`

HasBoxes returns a boolean if a field has been set.

### GetCorrected

`func (o *TaxStatement) GetCorrected() bool`

GetCorrected returns the Corrected field if non-nil, zero value otherwise.

### GetCorrectedOk

`func (o *TaxStatement) GetCorrectedOk() (*bool, bool)`

GetCorrectedOk returns a tuple with the Corrected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrected

`func (o *TaxStatement) SetCorrected(v bool)`

SetCorrected sets Corrected field to given value.

### HasCorrected

`func (o *TaxStatement) HasCorrected() bool`

HasCorrected returns a boolean if a field has been set.

### GetFurnishedAt

`func (o *TaxStatement) GetFurnishedAt() int64`

GetFurnishedAt returns the FurnishedAt field if non-nil, zero value otherwise.

### GetFurnishedAtOk

`func (o *TaxStatement) GetFurnishedAtOk() (*int64, bool)`

GetFurnishedAtOk returns a tuple with the FurnishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFurnishedAt

`func (o *TaxStatement) SetFurnishedAt(v int64)`

SetFurnishedAt sets FurnishedAt field to given value.

### HasFurnishedAt

`func (o *TaxStatement) HasFurnishedAt() bool`

HasFurnishedAt returns a boolean if a field has been set.

### GetId

`func (o *TaxStatement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxStatement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxStatement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxStatement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *TaxStatement) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TaxStatement) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TaxStatement) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TaxStatement) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPayer

`func (o *TaxStatement) GetPayer() TaxParty`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *TaxStatement) GetPayerOk() (*TaxParty, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *TaxStatement) SetPayer(v TaxParty)`

SetPayer sets Payer field to given value.

### HasPayer

`func (o *TaxStatement) HasPayer() bool`

HasPayer returns a boolean if a field has been set.

### GetRecipient

`func (o *TaxStatement) GetRecipient() TaxParty`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *TaxStatement) GetRecipientOk() (*TaxParty, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *TaxStatement) SetRecipient(v TaxParty)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *TaxStatement) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetSupersededBy

`func (o *TaxStatement) GetSupersededBy() string`

GetSupersededBy returns the SupersededBy field if non-nil, zero value otherwise.

### GetSupersededByOk

`func (o *TaxStatement) GetSupersededByOk() (*string, bool)`

GetSupersededByOk returns a tuple with the SupersededBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersededBy

`func (o *TaxStatement) SetSupersededBy(v string)`

SetSupersededBy sets SupersededBy field to given value.

### HasSupersededBy

`func (o *TaxStatement) HasSupersededBy() bool`

HasSupersededBy returns a boolean if a field has been set.

### GetSupersedes

`func (o *TaxStatement) GetSupersedes() string`

GetSupersedes returns the Supersedes field if non-nil, zero value otherwise.

### GetSupersedesOk

`func (o *TaxStatement) GetSupersedesOk() (*string, bool)`

GetSupersedesOk returns a tuple with the Supersedes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersedes

`func (o *TaxStatement) SetSupersedes(v string)`

SetSupersedes sets Supersedes field to given value.

### HasSupersedes

`func (o *TaxStatement) HasSupersedes() bool`

HasSupersedes returns a boolean if a field has been set.

### GetYear

`func (o *TaxStatement) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *TaxStatement) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *TaxStatement) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *TaxStatement) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TaxForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the account number the form prints — the original form&#39;s id, which a correction keeps, so the IRS can tie the two. | [optional] 
**Backup** | Pointer to [**TaxWithholding**](TaxWithholding.md) | Backup is the backup withholding the form carries. | [optional] 
**Boxes** | Pointer to [**[]TaxAmount**](TaxAmount.md) | Boxes are the figures, box 4 (federal income tax withheld) included. | [optional] 
**Certified** | Pointer to **bool** | Certified is whether the payee&#39;s W-9 was certified when the form was prepared. | [optional] 
**Corrected** | Pointer to **bool** | Corrected is true for a form that corrects one already furnished. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the form was prepared, unix seconds. | [optional] 
**Delivery** | Pointer to **string** | Delivery is electronic (Copy B in the payee&#39;s inbox) or paper, once furnished. | [optional] 
**Due** | Pointer to [**TaxDeadline**](TaxDeadline.md) | Due is when Copy B is due to the payee and the return to the IRS. | [optional] 
**FurnishedAt** | Pointer to **int64** | FurnishedAt is when Copy B was delivered in the inbox, or marked owed on paper. | [optional] 
**Id** | Pointer to **string** | ID is the form&#39;s id, \&quot;f1099_\&quot;-prefixed. | [optional] 
**Kind** | Pointer to **string** | Kind is 1099-NEC or 1099-MISC. | [optional] 
**MailedAt** | Pointer to **int64** | MailedAt is when the payer recorded mailing a paper Copy B. | [optional] 
**Notes** | Pointer to **[]string** | Notes are what still needs a person, in words. | [optional] 
**Payer** | Pointer to [**TaxParty**](TaxParty.md) | Payer is the filer. | [optional] 
**Payments** | Pointer to **[]string** | Payments are the rail payment ids summed into the boxes. | [optional] 
**Reason** | Pointer to **string** | Reason is why a corrected form was made. | [optional] 
**Recipient** | Pointer to [**TaxParty**](TaxParty.md) | Recipient is the payee. | [optional] 
**Status** | Pointer to **string** | Status is draft, reviewed, furnished, owed or void. | [optional] 
**SupersededBy** | Pointer to **string** | SupersededBy is the form that corrects this one, once there is one. | [optional] 
**Supersedes** | Pointer to **string** | Supersedes is the form this one corrects. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when its status last moved, unix seconds. | [optional] 
**W9** | Pointer to **string** | W9 is where the payer&#39;s W-9 request stood when the form was prepared. | [optional] 
**Year** | Pointer to **int64** | Year is the calendar year the payments were made. | [optional] 

## Methods

### NewTaxForm

`func NewTaxForm() *TaxForm`

NewTaxForm instantiates a new TaxForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxFormWithDefaults

`func NewTaxFormWithDefaults() *TaxForm`

NewTaxFormWithDefaults instantiates a new TaxForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *TaxForm) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TaxForm) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TaxForm) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TaxForm) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBackup

`func (o *TaxForm) GetBackup() TaxWithholding`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *TaxForm) GetBackupOk() (*TaxWithholding, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *TaxForm) SetBackup(v TaxWithholding)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *TaxForm) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetBoxes

`func (o *TaxForm) GetBoxes() []TaxAmount`

GetBoxes returns the Boxes field if non-nil, zero value otherwise.

### GetBoxesOk

`func (o *TaxForm) GetBoxesOk() (*[]TaxAmount, bool)`

GetBoxesOk returns a tuple with the Boxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxes

`func (o *TaxForm) SetBoxes(v []TaxAmount)`

SetBoxes sets Boxes field to given value.

### HasBoxes

`func (o *TaxForm) HasBoxes() bool`

HasBoxes returns a boolean if a field has been set.

### GetCertified

`func (o *TaxForm) GetCertified() bool`

GetCertified returns the Certified field if non-nil, zero value otherwise.

### GetCertifiedOk

`func (o *TaxForm) GetCertifiedOk() (*bool, bool)`

GetCertifiedOk returns a tuple with the Certified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertified

`func (o *TaxForm) SetCertified(v bool)`

SetCertified sets Certified field to given value.

### HasCertified

`func (o *TaxForm) HasCertified() bool`

HasCertified returns a boolean if a field has been set.

### GetCorrected

`func (o *TaxForm) GetCorrected() bool`

GetCorrected returns the Corrected field if non-nil, zero value otherwise.

### GetCorrectedOk

`func (o *TaxForm) GetCorrectedOk() (*bool, bool)`

GetCorrectedOk returns a tuple with the Corrected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrected

`func (o *TaxForm) SetCorrected(v bool)`

SetCorrected sets Corrected field to given value.

### HasCorrected

`func (o *TaxForm) HasCorrected() bool`

HasCorrected returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaxForm) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaxForm) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaxForm) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaxForm) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDelivery

`func (o *TaxForm) GetDelivery() string`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *TaxForm) GetDeliveryOk() (*string, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *TaxForm) SetDelivery(v string)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *TaxForm) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.

### GetDue

`func (o *TaxForm) GetDue() TaxDeadline`

GetDue returns the Due field if non-nil, zero value otherwise.

### GetDueOk

`func (o *TaxForm) GetDueOk() (*TaxDeadline, bool)`

GetDueOk returns a tuple with the Due field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDue

`func (o *TaxForm) SetDue(v TaxDeadline)`

SetDue sets Due field to given value.

### HasDue

`func (o *TaxForm) HasDue() bool`

HasDue returns a boolean if a field has been set.

### GetFurnishedAt

`func (o *TaxForm) GetFurnishedAt() int64`

GetFurnishedAt returns the FurnishedAt field if non-nil, zero value otherwise.

### GetFurnishedAtOk

`func (o *TaxForm) GetFurnishedAtOk() (*int64, bool)`

GetFurnishedAtOk returns a tuple with the FurnishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFurnishedAt

`func (o *TaxForm) SetFurnishedAt(v int64)`

SetFurnishedAt sets FurnishedAt field to given value.

### HasFurnishedAt

`func (o *TaxForm) HasFurnishedAt() bool`

HasFurnishedAt returns a boolean if a field has been set.

### GetId

`func (o *TaxForm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxForm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxForm) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxForm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *TaxForm) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TaxForm) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TaxForm) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TaxForm) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMailedAt

`func (o *TaxForm) GetMailedAt() int64`

GetMailedAt returns the MailedAt field if non-nil, zero value otherwise.

### GetMailedAtOk

`func (o *TaxForm) GetMailedAtOk() (*int64, bool)`

GetMailedAtOk returns a tuple with the MailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailedAt

`func (o *TaxForm) SetMailedAt(v int64)`

SetMailedAt sets MailedAt field to given value.

### HasMailedAt

`func (o *TaxForm) HasMailedAt() bool`

HasMailedAt returns a boolean if a field has been set.

### GetNotes

`func (o *TaxForm) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TaxForm) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TaxForm) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TaxForm) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPayer

`func (o *TaxForm) GetPayer() TaxParty`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *TaxForm) GetPayerOk() (*TaxParty, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *TaxForm) SetPayer(v TaxParty)`

SetPayer sets Payer field to given value.

### HasPayer

`func (o *TaxForm) HasPayer() bool`

HasPayer returns a boolean if a field has been set.

### GetPayments

`func (o *TaxForm) GetPayments() []string`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *TaxForm) GetPaymentsOk() (*[]string, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *TaxForm) SetPayments(v []string)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *TaxForm) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetReason

`func (o *TaxForm) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TaxForm) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TaxForm) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TaxForm) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRecipient

`func (o *TaxForm) GetRecipient() TaxParty`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *TaxForm) GetRecipientOk() (*TaxParty, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *TaxForm) SetRecipient(v TaxParty)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *TaxForm) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetStatus

`func (o *TaxForm) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaxForm) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaxForm) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaxForm) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupersededBy

`func (o *TaxForm) GetSupersededBy() string`

GetSupersededBy returns the SupersededBy field if non-nil, zero value otherwise.

### GetSupersededByOk

`func (o *TaxForm) GetSupersededByOk() (*string, bool)`

GetSupersededByOk returns a tuple with the SupersededBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersededBy

`func (o *TaxForm) SetSupersededBy(v string)`

SetSupersededBy sets SupersededBy field to given value.

### HasSupersededBy

`func (o *TaxForm) HasSupersededBy() bool`

HasSupersededBy returns a boolean if a field has been set.

### GetSupersedes

`func (o *TaxForm) GetSupersedes() string`

GetSupersedes returns the Supersedes field if non-nil, zero value otherwise.

### GetSupersedesOk

`func (o *TaxForm) GetSupersedesOk() (*string, bool)`

GetSupersedesOk returns a tuple with the Supersedes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersedes

`func (o *TaxForm) SetSupersedes(v string)`

SetSupersedes sets Supersedes field to given value.

### HasSupersedes

`func (o *TaxForm) HasSupersedes() bool`

HasSupersedes returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaxForm) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaxForm) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaxForm) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TaxForm) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetW9

`func (o *TaxForm) GetW9() string`

GetW9 returns the W9 field if non-nil, zero value otherwise.

### GetW9Ok

`func (o *TaxForm) GetW9Ok() (*string, bool)`

GetW9Ok returns a tuple with the W9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW9

`func (o *TaxForm) SetW9(v string)`

SetW9 sets W9 field to given value.

### HasW9

`func (o *TaxForm) HasW9() bool`

HasW9 returns a boolean if a field has been set.

### GetYear

`func (o *TaxForm) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *TaxForm) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *TaxForm) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *TaxForm) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



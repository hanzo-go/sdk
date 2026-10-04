# TaxW9

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category is what the payer states it pays this payee for. Only the payer sees it. | [optional] 
**DecidedAt** | Pointer to **int64** | DecidedAt is when the payee last granted, declined or revoked. | [optional] 
**Id** | Pointer to **string** | ID addresses the relationship on both sides. | [optional] 
**Match** | Pointer to [**TaxMatch**](TaxMatch.md) | Match is the IRS TIN Matching result the payer recorded, when it recorded one. | [optional] 
**Payee** | Pointer to **string** | Payee is the org whose W-9 it is. | [optional] 
**Payer** | Pointer to **string** | Payer is the org that asked. | [optional] 
**Profile** | Pointer to [**TaxProfile**](TaxProfile.md) | Profile is the payee&#39;s W-9, TIN masked — present only for the payer, only while the grant is live. | [optional] 
**RequestedAt** | Pointer to **int64** | RequestedAt is when the payer asked, unix seconds. | [optional] 
**Role** | Pointer to **string** | Role is the caller&#39;s side of it: payer or payee. | [optional] 
**Status** | Pointer to **string** | Status is requested, granted, declined or revoked. | [optional] 

## Methods

### NewTaxW9

`func NewTaxW9() *TaxW9`

NewTaxW9 instantiates a new TaxW9 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxW9WithDefaults

`func NewTaxW9WithDefaults() *TaxW9`

NewTaxW9WithDefaults instantiates a new TaxW9 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *TaxW9) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TaxW9) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TaxW9) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TaxW9) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDecidedAt

`func (o *TaxW9) GetDecidedAt() int64`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *TaxW9) GetDecidedAtOk() (*int64, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *TaxW9) SetDecidedAt(v int64)`

SetDecidedAt sets DecidedAt field to given value.

### HasDecidedAt

`func (o *TaxW9) HasDecidedAt() bool`

HasDecidedAt returns a boolean if a field has been set.

### GetId

`func (o *TaxW9) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxW9) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxW9) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxW9) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMatch

`func (o *TaxW9) GetMatch() TaxMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *TaxW9) GetMatchOk() (*TaxMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *TaxW9) SetMatch(v TaxMatch)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *TaxW9) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetPayee

`func (o *TaxW9) GetPayee() string`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *TaxW9) GetPayeeOk() (*string, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *TaxW9) SetPayee(v string)`

SetPayee sets Payee field to given value.

### HasPayee

`func (o *TaxW9) HasPayee() bool`

HasPayee returns a boolean if a field has been set.

### GetPayer

`func (o *TaxW9) GetPayer() string`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *TaxW9) GetPayerOk() (*string, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *TaxW9) SetPayer(v string)`

SetPayer sets Payer field to given value.

### HasPayer

`func (o *TaxW9) HasPayer() bool`

HasPayer returns a boolean if a field has been set.

### GetProfile

`func (o *TaxW9) GetProfile() TaxProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *TaxW9) GetProfileOk() (*TaxProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *TaxW9) SetProfile(v TaxProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *TaxW9) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetRequestedAt

`func (o *TaxW9) GetRequestedAt() int64`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *TaxW9) GetRequestedAtOk() (*int64, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *TaxW9) SetRequestedAt(v int64)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *TaxW9) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetRole

`func (o *TaxW9) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TaxW9) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TaxW9) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *TaxW9) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *TaxW9) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaxW9) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaxW9) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaxW9) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



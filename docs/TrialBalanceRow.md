# TrialBalanceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**ClosingCredit** | Pointer to **int32** |  | [optional] 
**ClosingDebit** | Pointer to **int32** |  | [optional] 
**Credit** | Pointer to **int32** | period movement | [optional] 
**Debit** | Pointer to **int32** | period movement | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OpeningCredit** | Pointer to **int32** |  | [optional] 
**OpeningDebit** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewTrialBalanceRow

`func NewTrialBalanceRow() *TrialBalanceRow`

NewTrialBalanceRow instantiates a new TrialBalanceRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialBalanceRowWithDefaults

`func NewTrialBalanceRowWithDefaults() *TrialBalanceRow`

NewTrialBalanceRowWithDefaults instantiates a new TrialBalanceRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *TrialBalanceRow) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TrialBalanceRow) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TrialBalanceRow) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TrialBalanceRow) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetClosingCredit

`func (o *TrialBalanceRow) GetClosingCredit() int32`

GetClosingCredit returns the ClosingCredit field if non-nil, zero value otherwise.

### GetClosingCreditOk

`func (o *TrialBalanceRow) GetClosingCreditOk() (*int32, bool)`

GetClosingCreditOk returns a tuple with the ClosingCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingCredit

`func (o *TrialBalanceRow) SetClosingCredit(v int32)`

SetClosingCredit sets ClosingCredit field to given value.

### HasClosingCredit

`func (o *TrialBalanceRow) HasClosingCredit() bool`

HasClosingCredit returns a boolean if a field has been set.

### GetClosingDebit

`func (o *TrialBalanceRow) GetClosingDebit() int32`

GetClosingDebit returns the ClosingDebit field if non-nil, zero value otherwise.

### GetClosingDebitOk

`func (o *TrialBalanceRow) GetClosingDebitOk() (*int32, bool)`

GetClosingDebitOk returns a tuple with the ClosingDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingDebit

`func (o *TrialBalanceRow) SetClosingDebit(v int32)`

SetClosingDebit sets ClosingDebit field to given value.

### HasClosingDebit

`func (o *TrialBalanceRow) HasClosingDebit() bool`

HasClosingDebit returns a boolean if a field has been set.

### GetCredit

`func (o *TrialBalanceRow) GetCredit() int32`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *TrialBalanceRow) GetCreditOk() (*int32, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *TrialBalanceRow) SetCredit(v int32)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *TrialBalanceRow) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetDebit

`func (o *TrialBalanceRow) GetDebit() int32`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *TrialBalanceRow) GetDebitOk() (*int32, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *TrialBalanceRow) SetDebit(v int32)`

SetDebit sets Debit field to given value.

### HasDebit

`func (o *TrialBalanceRow) HasDebit() bool`

HasDebit returns a boolean if a field has been set.

### GetName

`func (o *TrialBalanceRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrialBalanceRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrialBalanceRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrialBalanceRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpeningCredit

`func (o *TrialBalanceRow) GetOpeningCredit() int32`

GetOpeningCredit returns the OpeningCredit field if non-nil, zero value otherwise.

### GetOpeningCreditOk

`func (o *TrialBalanceRow) GetOpeningCreditOk() (*int32, bool)`

GetOpeningCreditOk returns a tuple with the OpeningCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningCredit

`func (o *TrialBalanceRow) SetOpeningCredit(v int32)`

SetOpeningCredit sets OpeningCredit field to given value.

### HasOpeningCredit

`func (o *TrialBalanceRow) HasOpeningCredit() bool`

HasOpeningCredit returns a boolean if a field has been set.

### GetOpeningDebit

`func (o *TrialBalanceRow) GetOpeningDebit() int32`

GetOpeningDebit returns the OpeningDebit field if non-nil, zero value otherwise.

### GetOpeningDebitOk

`func (o *TrialBalanceRow) GetOpeningDebitOk() (*int32, bool)`

GetOpeningDebitOk returns a tuple with the OpeningDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningDebit

`func (o *TrialBalanceRow) SetOpeningDebit(v int32)`

SetOpeningDebit sets OpeningDebit field to given value.

### HasOpeningDebit

`func (o *TrialBalanceRow) HasOpeningDebit() bool`

HasOpeningDebit returns a boolean if a field has been set.

### GetType

`func (o *TrialBalanceRow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TrialBalanceRow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TrialBalanceRow) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TrialBalanceRow) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



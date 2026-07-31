# CloudTrialBalanceRow

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

### NewCloudTrialBalanceRow

`func NewCloudTrialBalanceRow() *CloudTrialBalanceRow`

NewCloudTrialBalanceRow instantiates a new CloudTrialBalanceRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTrialBalanceRowWithDefaults

`func NewCloudTrialBalanceRowWithDefaults() *CloudTrialBalanceRow`

NewCloudTrialBalanceRowWithDefaults instantiates a new CloudTrialBalanceRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudTrialBalanceRow) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudTrialBalanceRow) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudTrialBalanceRow) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudTrialBalanceRow) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetClosingCredit

`func (o *CloudTrialBalanceRow) GetClosingCredit() int32`

GetClosingCredit returns the ClosingCredit field if non-nil, zero value otherwise.

### GetClosingCreditOk

`func (o *CloudTrialBalanceRow) GetClosingCreditOk() (*int32, bool)`

GetClosingCreditOk returns a tuple with the ClosingCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingCredit

`func (o *CloudTrialBalanceRow) SetClosingCredit(v int32)`

SetClosingCredit sets ClosingCredit field to given value.

### HasClosingCredit

`func (o *CloudTrialBalanceRow) HasClosingCredit() bool`

HasClosingCredit returns a boolean if a field has been set.

### GetClosingDebit

`func (o *CloudTrialBalanceRow) GetClosingDebit() int32`

GetClosingDebit returns the ClosingDebit field if non-nil, zero value otherwise.

### GetClosingDebitOk

`func (o *CloudTrialBalanceRow) GetClosingDebitOk() (*int32, bool)`

GetClosingDebitOk returns a tuple with the ClosingDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingDebit

`func (o *CloudTrialBalanceRow) SetClosingDebit(v int32)`

SetClosingDebit sets ClosingDebit field to given value.

### HasClosingDebit

`func (o *CloudTrialBalanceRow) HasClosingDebit() bool`

HasClosingDebit returns a boolean if a field has been set.

### GetCredit

`func (o *CloudTrialBalanceRow) GetCredit() int32`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *CloudTrialBalanceRow) GetCreditOk() (*int32, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *CloudTrialBalanceRow) SetCredit(v int32)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *CloudTrialBalanceRow) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetDebit

`func (o *CloudTrialBalanceRow) GetDebit() int32`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *CloudTrialBalanceRow) GetDebitOk() (*int32, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *CloudTrialBalanceRow) SetDebit(v int32)`

SetDebit sets Debit field to given value.

### HasDebit

`func (o *CloudTrialBalanceRow) HasDebit() bool`

HasDebit returns a boolean if a field has been set.

### GetName

`func (o *CloudTrialBalanceRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudTrialBalanceRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudTrialBalanceRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudTrialBalanceRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpeningCredit

`func (o *CloudTrialBalanceRow) GetOpeningCredit() int32`

GetOpeningCredit returns the OpeningCredit field if non-nil, zero value otherwise.

### GetOpeningCreditOk

`func (o *CloudTrialBalanceRow) GetOpeningCreditOk() (*int32, bool)`

GetOpeningCreditOk returns a tuple with the OpeningCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningCredit

`func (o *CloudTrialBalanceRow) SetOpeningCredit(v int32)`

SetOpeningCredit sets OpeningCredit field to given value.

### HasOpeningCredit

`func (o *CloudTrialBalanceRow) HasOpeningCredit() bool`

HasOpeningCredit returns a boolean if a field has been set.

### GetOpeningDebit

`func (o *CloudTrialBalanceRow) GetOpeningDebit() int32`

GetOpeningDebit returns the OpeningDebit field if non-nil, zero value otherwise.

### GetOpeningDebitOk

`func (o *CloudTrialBalanceRow) GetOpeningDebitOk() (*int32, bool)`

GetOpeningDebitOk returns a tuple with the OpeningDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningDebit

`func (o *CloudTrialBalanceRow) SetOpeningDebit(v int32)`

SetOpeningDebit sets OpeningDebit field to given value.

### HasOpeningDebit

`func (o *CloudTrialBalanceRow) HasOpeningDebit() bool`

HasOpeningDebit returns a boolean if a field has been set.

### GetType

`func (o *CloudTrialBalanceRow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudTrialBalanceRow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudTrialBalanceRow) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudTrialBalanceRow) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



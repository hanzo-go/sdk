# BooksTrialBalanceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the chart-of-accounts NUMBER this line reports on (\&quot;1000\&quot;, \&quot;4000\&quot;) — the stable posting key, not a display label. | [optional] 
**ClosingCredit** | Pointer to **int64** | ClosingCredit is that closing balance in cents when it is a credit balance. | [optional] 
**ClosingDebit** | Pointer to **int64** | ClosingDebit is the balance at the end of the window, in cents, when it is a debit balance. This is the column the report&#39;s totals are summed from. | [optional] 
**Credit** | Pointer to **int64** | Credit is the same window movement in cents when it was net credit. | [optional] 
**Debit** | Pointer to **int64** | Debit is the account&#39;s MOVEMENT within the window — closing minus opening, not the closing balance — in cents, when that movement was net debit. Zero when the account moved net credit. | [optional] 
**Name** | Pointer to **string** | Name is that account&#39;s human name from the fixed chart. | [optional] 
**OpeningCredit** | Pointer to **int64** | OpeningCredit is the same opening balance in cents when it fell on the credit side. Zero when the balance was a debit one. | [optional] 
**OpeningDebit** | Pointer to **int64** | OpeningDebit is the account&#39;s balance before the window began, in whole cents, when that balance was on the debit side. Zero when the balance was a credit one — the pair is exclusive, never two halves of one number. | [optional] 
**Type** | Pointer to **string** | Type is the account&#39;s fundamental class — asset, liability, income, expense or equity — which is also its normal balance side. It is carried for presentation and does NOT decide which column an amount lands in: placement follows the sign of the real net, so a contra balance shows up as one. | [optional] 

## Methods

### NewBooksTrialBalanceRow

`func NewBooksTrialBalanceRow() *BooksTrialBalanceRow`

NewBooksTrialBalanceRow instantiates a new BooksTrialBalanceRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksTrialBalanceRowWithDefaults

`func NewBooksTrialBalanceRowWithDefaults() *BooksTrialBalanceRow`

NewBooksTrialBalanceRowWithDefaults instantiates a new BooksTrialBalanceRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *BooksTrialBalanceRow) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *BooksTrialBalanceRow) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *BooksTrialBalanceRow) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *BooksTrialBalanceRow) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetClosingCredit

`func (o *BooksTrialBalanceRow) GetClosingCredit() int64`

GetClosingCredit returns the ClosingCredit field if non-nil, zero value otherwise.

### GetClosingCreditOk

`func (o *BooksTrialBalanceRow) GetClosingCreditOk() (*int64, bool)`

GetClosingCreditOk returns a tuple with the ClosingCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingCredit

`func (o *BooksTrialBalanceRow) SetClosingCredit(v int64)`

SetClosingCredit sets ClosingCredit field to given value.

### HasClosingCredit

`func (o *BooksTrialBalanceRow) HasClosingCredit() bool`

HasClosingCredit returns a boolean if a field has been set.

### GetClosingDebit

`func (o *BooksTrialBalanceRow) GetClosingDebit() int64`

GetClosingDebit returns the ClosingDebit field if non-nil, zero value otherwise.

### GetClosingDebitOk

`func (o *BooksTrialBalanceRow) GetClosingDebitOk() (*int64, bool)`

GetClosingDebitOk returns a tuple with the ClosingDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingDebit

`func (o *BooksTrialBalanceRow) SetClosingDebit(v int64)`

SetClosingDebit sets ClosingDebit field to given value.

### HasClosingDebit

`func (o *BooksTrialBalanceRow) HasClosingDebit() bool`

HasClosingDebit returns a boolean if a field has been set.

### GetCredit

`func (o *BooksTrialBalanceRow) GetCredit() int64`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *BooksTrialBalanceRow) GetCreditOk() (*int64, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *BooksTrialBalanceRow) SetCredit(v int64)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *BooksTrialBalanceRow) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetDebit

`func (o *BooksTrialBalanceRow) GetDebit() int64`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *BooksTrialBalanceRow) GetDebitOk() (*int64, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *BooksTrialBalanceRow) SetDebit(v int64)`

SetDebit sets Debit field to given value.

### HasDebit

`func (o *BooksTrialBalanceRow) HasDebit() bool`

HasDebit returns a boolean if a field has been set.

### GetName

`func (o *BooksTrialBalanceRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BooksTrialBalanceRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BooksTrialBalanceRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BooksTrialBalanceRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpeningCredit

`func (o *BooksTrialBalanceRow) GetOpeningCredit() int64`

GetOpeningCredit returns the OpeningCredit field if non-nil, zero value otherwise.

### GetOpeningCreditOk

`func (o *BooksTrialBalanceRow) GetOpeningCreditOk() (*int64, bool)`

GetOpeningCreditOk returns a tuple with the OpeningCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningCredit

`func (o *BooksTrialBalanceRow) SetOpeningCredit(v int64)`

SetOpeningCredit sets OpeningCredit field to given value.

### HasOpeningCredit

`func (o *BooksTrialBalanceRow) HasOpeningCredit() bool`

HasOpeningCredit returns a boolean if a field has been set.

### GetOpeningDebit

`func (o *BooksTrialBalanceRow) GetOpeningDebit() int64`

GetOpeningDebit returns the OpeningDebit field if non-nil, zero value otherwise.

### GetOpeningDebitOk

`func (o *BooksTrialBalanceRow) GetOpeningDebitOk() (*int64, bool)`

GetOpeningDebitOk returns a tuple with the OpeningDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningDebit

`func (o *BooksTrialBalanceRow) SetOpeningDebit(v int64)`

SetOpeningDebit sets OpeningDebit field to given value.

### HasOpeningDebit

`func (o *BooksTrialBalanceRow) HasOpeningDebit() bool`

HasOpeningDebit returns a boolean if a field has been set.

### GetType

`func (o *BooksTrialBalanceRow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BooksTrialBalanceRow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BooksTrialBalanceRow) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BooksTrialBalanceRow) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



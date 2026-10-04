# BooksFinancialPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceSheet** | Pointer to [**BooksBalanceSheet**](BooksBalanceSheet.md) | BalanceSheet is struck as of the period END, not the start. | [optional] 
**From** | Pointer to **string** | From opens the reporting period. Absent means from the beginning of the ledger. | [optional] 
**GeneratedAt** | Pointer to **string** | GeneratedAt is when the bundle was assembled — the moment the statements were struck, which is what makes two exports of the same period comparable. | [optional] 
**Gl** | Pointer to [**[]BooksGLRow**](BooksGLRow.md) | GL is the newest slice of ledger detail, as the audit trail behind the statements. It is CAPPED, so on a busy ledger it is a sample rather than the full support for the figures above. | [optional] 
**Org** | Pointer to **string** | Org is the organisation whose books these are — the validated caller&#39;s own, stamped so a downloaded bundle still says whose it is. | [optional] 
**Pnl** | Pointer to [**BooksPnL**](BooksPnL.md) | PnL is the income statement for the period, on an accrual basis. | [optional] 
**To** | Pointer to **string** | To closes it. Absent means up to now. | [optional] 
**TrialBalance** | Pointer to [**BooksTrialBalance**](BooksTrialBalance.md) | TrialBalance is the proof the ledger balances over the period. | [optional] 

## Methods

### NewBooksFinancialPackage

`func NewBooksFinancialPackage() *BooksFinancialPackage`

NewBooksFinancialPackage instantiates a new BooksFinancialPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksFinancialPackageWithDefaults

`func NewBooksFinancialPackageWithDefaults() *BooksFinancialPackage`

NewBooksFinancialPackageWithDefaults instantiates a new BooksFinancialPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceSheet

`func (o *BooksFinancialPackage) GetBalanceSheet() BooksBalanceSheet`

GetBalanceSheet returns the BalanceSheet field if non-nil, zero value otherwise.

### GetBalanceSheetOk

`func (o *BooksFinancialPackage) GetBalanceSheetOk() (*BooksBalanceSheet, bool)`

GetBalanceSheetOk returns a tuple with the BalanceSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceSheet

`func (o *BooksFinancialPackage) SetBalanceSheet(v BooksBalanceSheet)`

SetBalanceSheet sets BalanceSheet field to given value.

### HasBalanceSheet

`func (o *BooksFinancialPackage) HasBalanceSheet() bool`

HasBalanceSheet returns a boolean if a field has been set.

### GetFrom

`func (o *BooksFinancialPackage) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BooksFinancialPackage) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BooksFinancialPackage) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *BooksFinancialPackage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *BooksFinancialPackage) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *BooksFinancialPackage) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *BooksFinancialPackage) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *BooksFinancialPackage) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetGl

`func (o *BooksFinancialPackage) GetGl() []BooksGLRow`

GetGl returns the Gl field if non-nil, zero value otherwise.

### GetGlOk

`func (o *BooksFinancialPackage) GetGlOk() (*[]BooksGLRow, bool)`

GetGlOk returns a tuple with the Gl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGl

`func (o *BooksFinancialPackage) SetGl(v []BooksGLRow)`

SetGl sets Gl field to given value.

### HasGl

`func (o *BooksFinancialPackage) HasGl() bool`

HasGl returns a boolean if a field has been set.

### GetOrg

`func (o *BooksFinancialPackage) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *BooksFinancialPackage) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *BooksFinancialPackage) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *BooksFinancialPackage) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPnl

`func (o *BooksFinancialPackage) GetPnl() BooksPnL`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *BooksFinancialPackage) GetPnlOk() (*BooksPnL, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *BooksFinancialPackage) SetPnl(v BooksPnL)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *BooksFinancialPackage) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetTo

`func (o *BooksFinancialPackage) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BooksFinancialPackage) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BooksFinancialPackage) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *BooksFinancialPackage) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTrialBalance

`func (o *BooksFinancialPackage) GetTrialBalance() BooksTrialBalance`

GetTrialBalance returns the TrialBalance field if non-nil, zero value otherwise.

### GetTrialBalanceOk

`func (o *BooksFinancialPackage) GetTrialBalanceOk() (*BooksTrialBalance, bool)`

GetTrialBalanceOk returns a tuple with the TrialBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialBalance

`func (o *BooksFinancialPackage) SetTrialBalance(v BooksTrialBalance)`

SetTrialBalance sets TrialBalance field to given value.

### HasTrialBalance

`func (o *BooksFinancialPackage) HasTrialBalance() bool`

HasTrialBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



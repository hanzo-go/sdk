# FinancialPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceSheet** | Pointer to [**BalanceSheet**](BalanceSheet.md) | BalanceSheet is struck as of the period END, not the start. | [optional] 
**From** | Pointer to **string** | From opens the reporting period. Absent means from the beginning of the ledger. | [optional] 
**GeneratedAt** | Pointer to **string** | GeneratedAt is when the bundle was assembled — the moment the statements were struck, which is what makes two exports of the same period comparable. | [optional] 
**Gl** | Pointer to [**[]GLRow**](GLRow.md) | GL is the newest slice of ledger detail, as the audit trail behind the statements. It is CAPPED, so on a busy ledger it is a sample rather than the full support for the figures above. | [optional] 
**Org** | Pointer to **string** | Org is the organisation whose books these are — the validated caller&#39;s own, stamped so a downloaded bundle still says whose it is. | [optional] 
**Pnl** | Pointer to [**PnL**](PnL.md) | PnL is the income statement for the period, on an accrual basis. | [optional] 
**To** | Pointer to **string** | To closes it. Absent means up to now. | [optional] 
**TrialBalance** | Pointer to [**TrialBalance**](TrialBalance.md) | TrialBalance is the proof the ledger balances over the period. | [optional] 

## Methods

### NewFinancialPackage

`func NewFinancialPackage() *FinancialPackage`

NewFinancialPackage instantiates a new FinancialPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialPackageWithDefaults

`func NewFinancialPackageWithDefaults() *FinancialPackage`

NewFinancialPackageWithDefaults instantiates a new FinancialPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceSheet

`func (o *FinancialPackage) GetBalanceSheet() BalanceSheet`

GetBalanceSheet returns the BalanceSheet field if non-nil, zero value otherwise.

### GetBalanceSheetOk

`func (o *FinancialPackage) GetBalanceSheetOk() (*BalanceSheet, bool)`

GetBalanceSheetOk returns a tuple with the BalanceSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceSheet

`func (o *FinancialPackage) SetBalanceSheet(v BalanceSheet)`

SetBalanceSheet sets BalanceSheet field to given value.

### HasBalanceSheet

`func (o *FinancialPackage) HasBalanceSheet() bool`

HasBalanceSheet returns a boolean if a field has been set.

### GetFrom

`func (o *FinancialPackage) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *FinancialPackage) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *FinancialPackage) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *FinancialPackage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *FinancialPackage) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *FinancialPackage) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *FinancialPackage) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *FinancialPackage) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetGl

`func (o *FinancialPackage) GetGl() []GLRow`

GetGl returns the Gl field if non-nil, zero value otherwise.

### GetGlOk

`func (o *FinancialPackage) GetGlOk() (*[]GLRow, bool)`

GetGlOk returns a tuple with the Gl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGl

`func (o *FinancialPackage) SetGl(v []GLRow)`

SetGl sets Gl field to given value.

### HasGl

`func (o *FinancialPackage) HasGl() bool`

HasGl returns a boolean if a field has been set.

### GetOrg

`func (o *FinancialPackage) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *FinancialPackage) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *FinancialPackage) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *FinancialPackage) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPnl

`func (o *FinancialPackage) GetPnl() PnL`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *FinancialPackage) GetPnlOk() (*PnL, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *FinancialPackage) SetPnl(v PnL)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *FinancialPackage) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetTo

`func (o *FinancialPackage) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *FinancialPackage) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *FinancialPackage) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *FinancialPackage) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTrialBalance

`func (o *FinancialPackage) GetTrialBalance() TrialBalance`

GetTrialBalance returns the TrialBalance field if non-nil, zero value otherwise.

### GetTrialBalanceOk

`func (o *FinancialPackage) GetTrialBalanceOk() (*TrialBalance, bool)`

GetTrialBalanceOk returns a tuple with the TrialBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialBalance

`func (o *FinancialPackage) SetTrialBalance(v TrialBalance)`

SetTrialBalance sets TrialBalance field to given value.

### HasTrialBalance

`func (o *FinancialPackage) HasTrialBalance() bool`

HasTrialBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



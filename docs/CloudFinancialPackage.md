# CloudFinancialPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceSheet** | Pointer to [**CloudBalanceSheet**](CloudBalanceSheet.md) |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**GeneratedAt** | Pointer to **string** |  | [optional] 
**Gl** | Pointer to [**[]CloudGLRow**](CloudGLRow.md) |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Pnl** | Pointer to [**CloudPnL**](CloudPnL.md) |  | [optional] 
**To** | Pointer to **string** |  | [optional] 
**TrialBalance** | Pointer to [**CloudTrialBalance**](CloudTrialBalance.md) |  | [optional] 

## Methods

### NewCloudFinancialPackage

`func NewCloudFinancialPackage() *CloudFinancialPackage`

NewCloudFinancialPackage instantiates a new CloudFinancialPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFinancialPackageWithDefaults

`func NewCloudFinancialPackageWithDefaults() *CloudFinancialPackage`

NewCloudFinancialPackageWithDefaults instantiates a new CloudFinancialPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceSheet

`func (o *CloudFinancialPackage) GetBalanceSheet() CloudBalanceSheet`

GetBalanceSheet returns the BalanceSheet field if non-nil, zero value otherwise.

### GetBalanceSheetOk

`func (o *CloudFinancialPackage) GetBalanceSheetOk() (*CloudBalanceSheet, bool)`

GetBalanceSheetOk returns a tuple with the BalanceSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceSheet

`func (o *CloudFinancialPackage) SetBalanceSheet(v CloudBalanceSheet)`

SetBalanceSheet sets BalanceSheet field to given value.

### HasBalanceSheet

`func (o *CloudFinancialPackage) HasBalanceSheet() bool`

HasBalanceSheet returns a boolean if a field has been set.

### GetFrom

`func (o *CloudFinancialPackage) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CloudFinancialPackage) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CloudFinancialPackage) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CloudFinancialPackage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *CloudFinancialPackage) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *CloudFinancialPackage) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *CloudFinancialPackage) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *CloudFinancialPackage) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetGl

`func (o *CloudFinancialPackage) GetGl() []CloudGLRow`

GetGl returns the Gl field if non-nil, zero value otherwise.

### GetGlOk

`func (o *CloudFinancialPackage) GetGlOk() (*[]CloudGLRow, bool)`

GetGlOk returns a tuple with the Gl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGl

`func (o *CloudFinancialPackage) SetGl(v []CloudGLRow)`

SetGl sets Gl field to given value.

### HasGl

`func (o *CloudFinancialPackage) HasGl() bool`

HasGl returns a boolean if a field has been set.

### GetOrg

`func (o *CloudFinancialPackage) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudFinancialPackage) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudFinancialPackage) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudFinancialPackage) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPnl

`func (o *CloudFinancialPackage) GetPnl() CloudPnL`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *CloudFinancialPackage) GetPnlOk() (*CloudPnL, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *CloudFinancialPackage) SetPnl(v CloudPnL)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *CloudFinancialPackage) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetTo

`func (o *CloudFinancialPackage) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CloudFinancialPackage) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CloudFinancialPackage) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CloudFinancialPackage) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTrialBalance

`func (o *CloudFinancialPackage) GetTrialBalance() CloudTrialBalance`

GetTrialBalance returns the TrialBalance field if non-nil, zero value otherwise.

### GetTrialBalanceOk

`func (o *CloudFinancialPackage) GetTrialBalanceOk() (*CloudTrialBalance, bool)`

GetTrialBalanceOk returns a tuple with the TrialBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialBalance

`func (o *CloudFinancialPackage) SetTrialBalance(v CloudTrialBalance)`

SetTrialBalance sets TrialBalance field to given value.

### HasTrialBalance

`func (o *CloudFinancialPackage) HasTrialBalance() bool`

HasTrialBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



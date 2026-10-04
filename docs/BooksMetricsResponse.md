# BooksMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arr** | Pointer to **int64** | ARR is annualized recurring revenue in cents (MRR × 12). | [optional] 
**Burn** | Pointer to **int64** | Burn is total expense in cents over the period. | [optional] 
**Cash** | Pointer to **int64** | Cash is the bank + processor-clearing balance in cents as of To. | [optional] 
**Cogs** | Pointer to **int64** | COGS is cost of goods sold in cents over the period. | [optional] 
**DeferredRevenue** | Pointer to **int64** | DeferredRevenue is the customer-wallet liability in cents as of To. | [optional] 
**Figures** | Pointer to [**[]BooksFigure**](BooksFigure.md) | Figures is the same snapshot rendered through books&#39; one money formatter. | [optional] 
**From** | Pointer to **string** | From is the RFC3339 start of the reporting window, exclusive; absent for all time. | [optional] 
**GrossMarginBps** | Pointer to **int64** | GrossMarginBps is GrossProfit / Revenue in basis points (7000 &#x3D; 70%). | [optional] 
**GrossProfit** | Pointer to **int64** | GrossProfit is Revenue − COGS, in cents. | [optional] 
**MonthlyBurn** | Pointer to **int64** | MonthlyBurn is net cash burned per month in cents; 0 when not losing cash. | [optional] 
**Months** | Pointer to **int64** | Months is the window length in whole months used to normalize MRR and burn. | [optional] 
**Mrr** | Pointer to **int64** | MRR is monthly recurring revenue in cents. | [optional] 
**NetIncome** | Pointer to **int64** | NetIncome is Revenue − Burn, in cents. | [optional] 
**Period** | Pointer to **string** | Period is the human window label, e.g. \&quot;2026-07\&quot; or \&quot;all-time\&quot;. | [optional] 
**Revenue** | Pointer to **int64** | Revenue is recognized revenue in cents over the period. | [optional] 
**RunwayMonths** | Pointer to **int64** | RunwayMonths is Cash / MonthlyBurn; -1 means infinite (the org is not burning). | [optional] 
**To** | Pointer to **string** | To is the RFC3339 end of the reporting window, inclusive; absent for up to now. | [optional] 

## Methods

### NewBooksMetricsResponse

`func NewBooksMetricsResponse() *BooksMetricsResponse`

NewBooksMetricsResponse instantiates a new BooksMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksMetricsResponseWithDefaults

`func NewBooksMetricsResponseWithDefaults() *BooksMetricsResponse`

NewBooksMetricsResponseWithDefaults instantiates a new BooksMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArr

`func (o *BooksMetricsResponse) GetArr() int64`

GetArr returns the Arr field if non-nil, zero value otherwise.

### GetArrOk

`func (o *BooksMetricsResponse) GetArrOk() (*int64, bool)`

GetArrOk returns a tuple with the Arr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArr

`func (o *BooksMetricsResponse) SetArr(v int64)`

SetArr sets Arr field to given value.

### HasArr

`func (o *BooksMetricsResponse) HasArr() bool`

HasArr returns a boolean if a field has been set.

### GetBurn

`func (o *BooksMetricsResponse) GetBurn() int64`

GetBurn returns the Burn field if non-nil, zero value otherwise.

### GetBurnOk

`func (o *BooksMetricsResponse) GetBurnOk() (*int64, bool)`

GetBurnOk returns a tuple with the Burn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurn

`func (o *BooksMetricsResponse) SetBurn(v int64)`

SetBurn sets Burn field to given value.

### HasBurn

`func (o *BooksMetricsResponse) HasBurn() bool`

HasBurn returns a boolean if a field has been set.

### GetCash

`func (o *BooksMetricsResponse) GetCash() int64`

GetCash returns the Cash field if non-nil, zero value otherwise.

### GetCashOk

`func (o *BooksMetricsResponse) GetCashOk() (*int64, bool)`

GetCashOk returns a tuple with the Cash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCash

`func (o *BooksMetricsResponse) SetCash(v int64)`

SetCash sets Cash field to given value.

### HasCash

`func (o *BooksMetricsResponse) HasCash() bool`

HasCash returns a boolean if a field has been set.

### GetCogs

`func (o *BooksMetricsResponse) GetCogs() int64`

GetCogs returns the Cogs field if non-nil, zero value otherwise.

### GetCogsOk

`func (o *BooksMetricsResponse) GetCogsOk() (*int64, bool)`

GetCogsOk returns a tuple with the Cogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs

`func (o *BooksMetricsResponse) SetCogs(v int64)`

SetCogs sets Cogs field to given value.

### HasCogs

`func (o *BooksMetricsResponse) HasCogs() bool`

HasCogs returns a boolean if a field has been set.

### GetDeferredRevenue

`func (o *BooksMetricsResponse) GetDeferredRevenue() int64`

GetDeferredRevenue returns the DeferredRevenue field if non-nil, zero value otherwise.

### GetDeferredRevenueOk

`func (o *BooksMetricsResponse) GetDeferredRevenueOk() (*int64, bool)`

GetDeferredRevenueOk returns a tuple with the DeferredRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredRevenue

`func (o *BooksMetricsResponse) SetDeferredRevenue(v int64)`

SetDeferredRevenue sets DeferredRevenue field to given value.

### HasDeferredRevenue

`func (o *BooksMetricsResponse) HasDeferredRevenue() bool`

HasDeferredRevenue returns a boolean if a field has been set.

### GetFigures

`func (o *BooksMetricsResponse) GetFigures() []BooksFigure`

GetFigures returns the Figures field if non-nil, zero value otherwise.

### GetFiguresOk

`func (o *BooksMetricsResponse) GetFiguresOk() (*[]BooksFigure, bool)`

GetFiguresOk returns a tuple with the Figures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFigures

`func (o *BooksMetricsResponse) SetFigures(v []BooksFigure)`

SetFigures sets Figures field to given value.

### HasFigures

`func (o *BooksMetricsResponse) HasFigures() bool`

HasFigures returns a boolean if a field has been set.

### GetFrom

`func (o *BooksMetricsResponse) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BooksMetricsResponse) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BooksMetricsResponse) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *BooksMetricsResponse) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGrossMarginBps

`func (o *BooksMetricsResponse) GetGrossMarginBps() int64`

GetGrossMarginBps returns the GrossMarginBps field if non-nil, zero value otherwise.

### GetGrossMarginBpsOk

`func (o *BooksMetricsResponse) GetGrossMarginBpsOk() (*int64, bool)`

GetGrossMarginBpsOk returns a tuple with the GrossMarginBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossMarginBps

`func (o *BooksMetricsResponse) SetGrossMarginBps(v int64)`

SetGrossMarginBps sets GrossMarginBps field to given value.

### HasGrossMarginBps

`func (o *BooksMetricsResponse) HasGrossMarginBps() bool`

HasGrossMarginBps returns a boolean if a field has been set.

### GetGrossProfit

`func (o *BooksMetricsResponse) GetGrossProfit() int64`

GetGrossProfit returns the GrossProfit field if non-nil, zero value otherwise.

### GetGrossProfitOk

`func (o *BooksMetricsResponse) GetGrossProfitOk() (*int64, bool)`

GetGrossProfitOk returns a tuple with the GrossProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossProfit

`func (o *BooksMetricsResponse) SetGrossProfit(v int64)`

SetGrossProfit sets GrossProfit field to given value.

### HasGrossProfit

`func (o *BooksMetricsResponse) HasGrossProfit() bool`

HasGrossProfit returns a boolean if a field has been set.

### GetMonthlyBurn

`func (o *BooksMetricsResponse) GetMonthlyBurn() int64`

GetMonthlyBurn returns the MonthlyBurn field if non-nil, zero value otherwise.

### GetMonthlyBurnOk

`func (o *BooksMetricsResponse) GetMonthlyBurnOk() (*int64, bool)`

GetMonthlyBurnOk returns a tuple with the MonthlyBurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyBurn

`func (o *BooksMetricsResponse) SetMonthlyBurn(v int64)`

SetMonthlyBurn sets MonthlyBurn field to given value.

### HasMonthlyBurn

`func (o *BooksMetricsResponse) HasMonthlyBurn() bool`

HasMonthlyBurn returns a boolean if a field has been set.

### GetMonths

`func (o *BooksMetricsResponse) GetMonths() int64`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *BooksMetricsResponse) GetMonthsOk() (*int64, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *BooksMetricsResponse) SetMonths(v int64)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *BooksMetricsResponse) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### GetMrr

`func (o *BooksMetricsResponse) GetMrr() int64`

GetMrr returns the Mrr field if non-nil, zero value otherwise.

### GetMrrOk

`func (o *BooksMetricsResponse) GetMrrOk() (*int64, bool)`

GetMrrOk returns a tuple with the Mrr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrr

`func (o *BooksMetricsResponse) SetMrr(v int64)`

SetMrr sets Mrr field to given value.

### HasMrr

`func (o *BooksMetricsResponse) HasMrr() bool`

HasMrr returns a boolean if a field has been set.

### GetNetIncome

`func (o *BooksMetricsResponse) GetNetIncome() int64`

GetNetIncome returns the NetIncome field if non-nil, zero value otherwise.

### GetNetIncomeOk

`func (o *BooksMetricsResponse) GetNetIncomeOk() (*int64, bool)`

GetNetIncomeOk returns a tuple with the NetIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIncome

`func (o *BooksMetricsResponse) SetNetIncome(v int64)`

SetNetIncome sets NetIncome field to given value.

### HasNetIncome

`func (o *BooksMetricsResponse) HasNetIncome() bool`

HasNetIncome returns a boolean if a field has been set.

### GetPeriod

`func (o *BooksMetricsResponse) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *BooksMetricsResponse) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *BooksMetricsResponse) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *BooksMetricsResponse) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRevenue

`func (o *BooksMetricsResponse) GetRevenue() int64`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *BooksMetricsResponse) GetRevenueOk() (*int64, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *BooksMetricsResponse) SetRevenue(v int64)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *BooksMetricsResponse) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetRunwayMonths

`func (o *BooksMetricsResponse) GetRunwayMonths() int64`

GetRunwayMonths returns the RunwayMonths field if non-nil, zero value otherwise.

### GetRunwayMonthsOk

`func (o *BooksMetricsResponse) GetRunwayMonthsOk() (*int64, bool)`

GetRunwayMonthsOk returns a tuple with the RunwayMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunwayMonths

`func (o *BooksMetricsResponse) SetRunwayMonths(v int64)`

SetRunwayMonths sets RunwayMonths field to given value.

### HasRunwayMonths

`func (o *BooksMetricsResponse) HasRunwayMonths() bool`

HasRunwayMonths returns a boolean if a field has been set.

### GetTo

`func (o *BooksMetricsResponse) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BooksMetricsResponse) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BooksMetricsResponse) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *BooksMetricsResponse) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



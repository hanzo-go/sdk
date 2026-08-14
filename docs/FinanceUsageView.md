# FinanceUsageView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] 
**End** | Pointer to **string** |  | [optional] 
**Lines** | Pointer to [**[]UsageLine**](UsageLine.md) |  | [optional] 
**Series** | Pointer to [**[]Sample**](Sample.md) |  | [optional] 
**Start** | Pointer to **string** |  | [optional] 
**TotalCents** | Pointer to **int32** |  | [optional] 

## Methods

### NewFinanceUsageView

`func NewFinanceUsageView() *FinanceUsageView`

NewFinanceUsageView instantiates a new FinanceUsageView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinanceUsageViewWithDefaults

`func NewFinanceUsageViewWithDefaults() *FinanceUsageView`

NewFinanceUsageViewWithDefaults instantiates a new FinanceUsageView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *FinanceUsageView) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FinanceUsageView) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FinanceUsageView) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FinanceUsageView) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEnd

`func (o *FinanceUsageView) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *FinanceUsageView) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *FinanceUsageView) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *FinanceUsageView) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetLines

`func (o *FinanceUsageView) GetLines() []UsageLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *FinanceUsageView) GetLinesOk() (*[]UsageLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *FinanceUsageView) SetLines(v []UsageLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *FinanceUsageView) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetSeries

`func (o *FinanceUsageView) GetSeries() []Sample`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *FinanceUsageView) GetSeriesOk() (*[]Sample, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *FinanceUsageView) SetSeries(v []Sample)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *FinanceUsageView) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetStart

`func (o *FinanceUsageView) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *FinanceUsageView) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *FinanceUsageView) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *FinanceUsageView) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTotalCents

`func (o *FinanceUsageView) GetTotalCents() int32`

GetTotalCents returns the TotalCents field if non-nil, zero value otherwise.

### GetTotalCentsOk

`func (o *FinanceUsageView) GetTotalCentsOk() (*int32, bool)`

GetTotalCentsOk returns a tuple with the TotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCents

`func (o *FinanceUsageView) SetTotalCents(v int32)`

SetTotalCents sets TotalCents field to given value.

### HasTotalCents

`func (o *FinanceUsageView) HasTotalCents() bool`

HasTotalCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



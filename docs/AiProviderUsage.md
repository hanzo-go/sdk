# AiProviderUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** |  | [optional] 
**ByModel** | Pointer to [**[]AiProviderUsageModelSpend**](AiProviderUsageModelSpend.md) |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**End** | Pointer to **string** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Series** | Pointer to [**[]AiProviderUsageSeriesPoint**](AiProviderUsageSeriesPoint.md) |  | [optional] 
**Start** | Pointer to **string** |  | [optional] 
**Totals** | Pointer to [**AiProviderUsageTotals**](AiProviderUsageTotals.md) |  | [optional] 

## Methods

### NewAiProviderUsage

`func NewAiProviderUsage() *AiProviderUsage`

NewAiProviderUsage instantiates a new AiProviderUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiProviderUsageWithDefaults

`func NewAiProviderUsageWithDefaults() *AiProviderUsage`

NewAiProviderUsageWithDefaults instantiates a new AiProviderUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *AiProviderUsage) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *AiProviderUsage) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *AiProviderUsage) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *AiProviderUsage) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetByModel

`func (o *AiProviderUsage) GetByModel() []AiProviderUsageModelSpend`

GetByModel returns the ByModel field if non-nil, zero value otherwise.

### GetByModelOk

`func (o *AiProviderUsage) GetByModelOk() (*[]AiProviderUsageModelSpend, bool)`

GetByModelOk returns a tuple with the ByModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByModel

`func (o *AiProviderUsage) SetByModel(v []AiProviderUsageModelSpend)`

SetByModel sets ByModel field to given value.

### HasByModel

`func (o *AiProviderUsage) HasByModel() bool`

HasByModel returns a boolean if a field has been set.

### GetConnected

`func (o *AiProviderUsage) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *AiProviderUsage) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *AiProviderUsage) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *AiProviderUsage) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetCurrency

`func (o *AiProviderUsage) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AiProviderUsage) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AiProviderUsage) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AiProviderUsage) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEnd

`func (o *AiProviderUsage) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AiProviderUsage) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AiProviderUsage) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *AiProviderUsage) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetInterval

`func (o *AiProviderUsage) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *AiProviderUsage) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *AiProviderUsage) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *AiProviderUsage) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetNote

`func (o *AiProviderUsage) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AiProviderUsage) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AiProviderUsage) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *AiProviderUsage) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetProvider

`func (o *AiProviderUsage) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AiProviderUsage) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AiProviderUsage) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AiProviderUsage) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSeries

`func (o *AiProviderUsage) GetSeries() []AiProviderUsageSeriesPoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *AiProviderUsage) GetSeriesOk() (*[]AiProviderUsageSeriesPoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *AiProviderUsage) SetSeries(v []AiProviderUsageSeriesPoint)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *AiProviderUsage) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetStart

`func (o *AiProviderUsage) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AiProviderUsage) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AiProviderUsage) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *AiProviderUsage) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTotals

`func (o *AiProviderUsage) GetTotals() AiProviderUsageTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *AiProviderUsage) GetTotalsOk() (*AiProviderUsageTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *AiProviderUsage) SetTotals(v AiProviderUsageTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *AiProviderUsage) HasTotals() bool`

HasTotals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



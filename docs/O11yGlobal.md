# O11yGlobal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** |  | [optional] 
**Llm** | Pointer to [**O11yLLM**](O11yLLM.md) |  | [optional] 
**LogSeries** | Pointer to [**[]O11yLogPoint**](O11yLogPoint.md) |  | [optional] 
**Range** | Pointer to **string** |  | [optional] 
**Series** | Pointer to [**[]O11ySeries**](O11ySeries.md) |  | [optional] 
**Start** | Pointer to **string** |  | [optional] 
**TopModels** | Pointer to [**[]O11yModelStat**](O11yModelStat.md) |  | [optional] 
**TopOrgs** | Pointer to [**[]O11yOrgStat**](O11yOrgStat.md) |  | [optional] 
**TopServices** | Pointer to [**[]O11ySvcStat**](O11ySvcStat.md) |  | [optional] 
**Totals** | Pointer to [**O11yTotals**](O11yTotals.md) |  | [optional] 

## Methods

### NewO11yGlobal

`func NewO11yGlobal() *O11yGlobal`

NewO11yGlobal instantiates a new O11yGlobal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yGlobalWithDefaults

`func NewO11yGlobalWithDefaults() *O11yGlobal`

NewO11yGlobalWithDefaults instantiates a new O11yGlobal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *O11yGlobal) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yGlobal) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yGlobal) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yGlobal) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetLlm

`func (o *O11yGlobal) GetLlm() O11yLLM`

GetLlm returns the Llm field if non-nil, zero value otherwise.

### GetLlmOk

`func (o *O11yGlobal) GetLlmOk() (*O11yLLM, bool)`

GetLlmOk returns a tuple with the Llm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlm

`func (o *O11yGlobal) SetLlm(v O11yLLM)`

SetLlm sets Llm field to given value.

### HasLlm

`func (o *O11yGlobal) HasLlm() bool`

HasLlm returns a boolean if a field has been set.

### GetLogSeries

`func (o *O11yGlobal) GetLogSeries() []O11yLogPoint`

GetLogSeries returns the LogSeries field if non-nil, zero value otherwise.

### GetLogSeriesOk

`func (o *O11yGlobal) GetLogSeriesOk() (*[]O11yLogPoint, bool)`

GetLogSeriesOk returns a tuple with the LogSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSeries

`func (o *O11yGlobal) SetLogSeries(v []O11yLogPoint)`

SetLogSeries sets LogSeries field to given value.

### HasLogSeries

`func (o *O11yGlobal) HasLogSeries() bool`

HasLogSeries returns a boolean if a field has been set.

### GetRange

`func (o *O11yGlobal) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *O11yGlobal) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *O11yGlobal) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *O11yGlobal) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSeries

`func (o *O11yGlobal) GetSeries() []O11ySeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *O11yGlobal) GetSeriesOk() (*[]O11ySeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *O11yGlobal) SetSeries(v []O11ySeries)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *O11yGlobal) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetStart

`func (o *O11yGlobal) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yGlobal) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yGlobal) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yGlobal) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTopModels

`func (o *O11yGlobal) GetTopModels() []O11yModelStat`

GetTopModels returns the TopModels field if non-nil, zero value otherwise.

### GetTopModelsOk

`func (o *O11yGlobal) GetTopModelsOk() (*[]O11yModelStat, bool)`

GetTopModelsOk returns a tuple with the TopModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopModels

`func (o *O11yGlobal) SetTopModels(v []O11yModelStat)`

SetTopModels sets TopModels field to given value.

### HasTopModels

`func (o *O11yGlobal) HasTopModels() bool`

HasTopModels returns a boolean if a field has been set.

### GetTopOrgs

`func (o *O11yGlobal) GetTopOrgs() []O11yOrgStat`

GetTopOrgs returns the TopOrgs field if non-nil, zero value otherwise.

### GetTopOrgsOk

`func (o *O11yGlobal) GetTopOrgsOk() (*[]O11yOrgStat, bool)`

GetTopOrgsOk returns a tuple with the TopOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopOrgs

`func (o *O11yGlobal) SetTopOrgs(v []O11yOrgStat)`

SetTopOrgs sets TopOrgs field to given value.

### HasTopOrgs

`func (o *O11yGlobal) HasTopOrgs() bool`

HasTopOrgs returns a boolean if a field has been set.

### GetTopServices

`func (o *O11yGlobal) GetTopServices() []O11ySvcStat`

GetTopServices returns the TopServices field if non-nil, zero value otherwise.

### GetTopServicesOk

`func (o *O11yGlobal) GetTopServicesOk() (*[]O11ySvcStat, bool)`

GetTopServicesOk returns a tuple with the TopServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopServices

`func (o *O11yGlobal) SetTopServices(v []O11ySvcStat)`

SetTopServices sets TopServices field to given value.

### HasTopServices

`func (o *O11yGlobal) HasTopServices() bool`

HasTopServices returns a boolean if a field has been set.

### GetTotals

`func (o *O11yGlobal) GetTotals() O11yTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *O11yGlobal) GetTotalsOk() (*O11yTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *O11yGlobal) SetTotals(v O11yTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *O11yGlobal) HasTotals() bool`

HasTotals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



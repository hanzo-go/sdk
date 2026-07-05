# AdminO11yGlobal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **string** |  | [optional] 
**End** | Pointer to **string** |  | [optional] 
**Totals** | Pointer to [**AdminO11yTotals**](AdminO11yTotals.md) |  | [optional] 
**Series** | Pointer to [**[]AdminO11ySeries**](AdminO11ySeries.md) |  | [optional] 
**LogSeries** | Pointer to [**[]AdminO11yLogPoint**](AdminO11yLogPoint.md) |  | [optional] 
**TopOrgs** | Pointer to [**[]AdminO11yOrgStat**](AdminO11yOrgStat.md) |  | [optional] 
**TopModels** | Pointer to [**[]AdminO11yModelStat**](AdminO11yModelStat.md) |  | [optional] 
**TopServices** | Pointer to [**[]AdminO11ySvcStat**](AdminO11ySvcStat.md) |  | [optional] 
**Llm** | Pointer to [**AdminO11yLLM**](AdminO11yLLM.md) |  | [optional] 

## Methods

### NewAdminO11yGlobal

`func NewAdminO11yGlobal() *AdminO11yGlobal`

NewAdminO11yGlobal instantiates a new AdminO11yGlobal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminO11yGlobalWithDefaults

`func NewAdminO11yGlobalWithDefaults() *AdminO11yGlobal`

NewAdminO11yGlobalWithDefaults instantiates a new AdminO11yGlobal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *AdminO11yGlobal) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *AdminO11yGlobal) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *AdminO11yGlobal) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *AdminO11yGlobal) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetStart

`func (o *AdminO11yGlobal) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AdminO11yGlobal) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AdminO11yGlobal) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *AdminO11yGlobal) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *AdminO11yGlobal) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AdminO11yGlobal) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AdminO11yGlobal) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *AdminO11yGlobal) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetTotals

`func (o *AdminO11yGlobal) GetTotals() AdminO11yTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *AdminO11yGlobal) GetTotalsOk() (*AdminO11yTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *AdminO11yGlobal) SetTotals(v AdminO11yTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *AdminO11yGlobal) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetSeries

`func (o *AdminO11yGlobal) GetSeries() []AdminO11ySeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *AdminO11yGlobal) GetSeriesOk() (*[]AdminO11ySeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *AdminO11yGlobal) SetSeries(v []AdminO11ySeries)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *AdminO11yGlobal) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetLogSeries

`func (o *AdminO11yGlobal) GetLogSeries() []AdminO11yLogPoint`

GetLogSeries returns the LogSeries field if non-nil, zero value otherwise.

### GetLogSeriesOk

`func (o *AdminO11yGlobal) GetLogSeriesOk() (*[]AdminO11yLogPoint, bool)`

GetLogSeriesOk returns a tuple with the LogSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSeries

`func (o *AdminO11yGlobal) SetLogSeries(v []AdminO11yLogPoint)`

SetLogSeries sets LogSeries field to given value.

### HasLogSeries

`func (o *AdminO11yGlobal) HasLogSeries() bool`

HasLogSeries returns a boolean if a field has been set.

### GetTopOrgs

`func (o *AdminO11yGlobal) GetTopOrgs() []AdminO11yOrgStat`

GetTopOrgs returns the TopOrgs field if non-nil, zero value otherwise.

### GetTopOrgsOk

`func (o *AdminO11yGlobal) GetTopOrgsOk() (*[]AdminO11yOrgStat, bool)`

GetTopOrgsOk returns a tuple with the TopOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopOrgs

`func (o *AdminO11yGlobal) SetTopOrgs(v []AdminO11yOrgStat)`

SetTopOrgs sets TopOrgs field to given value.

### HasTopOrgs

`func (o *AdminO11yGlobal) HasTopOrgs() bool`

HasTopOrgs returns a boolean if a field has been set.

### GetTopModels

`func (o *AdminO11yGlobal) GetTopModels() []AdminO11yModelStat`

GetTopModels returns the TopModels field if non-nil, zero value otherwise.

### GetTopModelsOk

`func (o *AdminO11yGlobal) GetTopModelsOk() (*[]AdminO11yModelStat, bool)`

GetTopModelsOk returns a tuple with the TopModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopModels

`func (o *AdminO11yGlobal) SetTopModels(v []AdminO11yModelStat)`

SetTopModels sets TopModels field to given value.

### HasTopModels

`func (o *AdminO11yGlobal) HasTopModels() bool`

HasTopModels returns a boolean if a field has been set.

### GetTopServices

`func (o *AdminO11yGlobal) GetTopServices() []AdminO11ySvcStat`

GetTopServices returns the TopServices field if non-nil, zero value otherwise.

### GetTopServicesOk

`func (o *AdminO11yGlobal) GetTopServicesOk() (*[]AdminO11ySvcStat, bool)`

GetTopServicesOk returns a tuple with the TopServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopServices

`func (o *AdminO11yGlobal) SetTopServices(v []AdminO11ySvcStat)`

SetTopServices sets TopServices field to given value.

### HasTopServices

`func (o *AdminO11yGlobal) HasTopServices() bool`

HasTopServices returns a boolean if a field has been set.

### GetLlm

`func (o *AdminO11yGlobal) GetLlm() AdminO11yLLM`

GetLlm returns the Llm field if non-nil, zero value otherwise.

### GetLlmOk

`func (o *AdminO11yGlobal) GetLlmOk() (*AdminO11yLLM, bool)`

GetLlmOk returns a tuple with the Llm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlm

`func (o *AdminO11yGlobal) SetLlm(v AdminO11yLLM)`

SetLlm sets Llm field to given value.

### HasLlm

`func (o *AdminO11yGlobal) HasLlm() bool`

HasLlm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



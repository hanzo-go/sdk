# CloudO11yGlobal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** |  | [optional] 
**Llm** | Pointer to [**CloudO11yLLM**](CloudO11yLLM.md) |  | [optional] 
**LogSeries** | Pointer to [**[]CloudO11yLogPoint**](CloudO11yLogPoint.md) |  | [optional] 
**Range** | Pointer to **string** |  | [optional] 
**Series** | Pointer to [**[]CloudO11ySeries**](CloudO11ySeries.md) |  | [optional] 
**Start** | Pointer to **string** |  | [optional] 
**TopModels** | Pointer to [**[]CloudO11yModelStat**](CloudO11yModelStat.md) |  | [optional] 
**TopOrgs** | Pointer to [**[]CloudO11yOrgStat**](CloudO11yOrgStat.md) |  | [optional] 
**TopServices** | Pointer to [**[]CloudO11ySvcStat**](CloudO11ySvcStat.md) |  | [optional] 
**Totals** | Pointer to [**CloudO11yTotals**](CloudO11yTotals.md) |  | [optional] 

## Methods

### NewCloudO11yGlobal

`func NewCloudO11yGlobal() *CloudO11yGlobal`

NewCloudO11yGlobal instantiates a new CloudO11yGlobal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudO11yGlobalWithDefaults

`func NewCloudO11yGlobalWithDefaults() *CloudO11yGlobal`

NewCloudO11yGlobalWithDefaults instantiates a new CloudO11yGlobal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *CloudO11yGlobal) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CloudO11yGlobal) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CloudO11yGlobal) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CloudO11yGlobal) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetLlm

`func (o *CloudO11yGlobal) GetLlm() CloudO11yLLM`

GetLlm returns the Llm field if non-nil, zero value otherwise.

### GetLlmOk

`func (o *CloudO11yGlobal) GetLlmOk() (*CloudO11yLLM, bool)`

GetLlmOk returns a tuple with the Llm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlm

`func (o *CloudO11yGlobal) SetLlm(v CloudO11yLLM)`

SetLlm sets Llm field to given value.

### HasLlm

`func (o *CloudO11yGlobal) HasLlm() bool`

HasLlm returns a boolean if a field has been set.

### GetLogSeries

`func (o *CloudO11yGlobal) GetLogSeries() []CloudO11yLogPoint`

GetLogSeries returns the LogSeries field if non-nil, zero value otherwise.

### GetLogSeriesOk

`func (o *CloudO11yGlobal) GetLogSeriesOk() (*[]CloudO11yLogPoint, bool)`

GetLogSeriesOk returns a tuple with the LogSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSeries

`func (o *CloudO11yGlobal) SetLogSeries(v []CloudO11yLogPoint)`

SetLogSeries sets LogSeries field to given value.

### HasLogSeries

`func (o *CloudO11yGlobal) HasLogSeries() bool`

HasLogSeries returns a boolean if a field has been set.

### GetRange

`func (o *CloudO11yGlobal) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CloudO11yGlobal) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CloudO11yGlobal) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *CloudO11yGlobal) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSeries

`func (o *CloudO11yGlobal) GetSeries() []CloudO11ySeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *CloudO11yGlobal) GetSeriesOk() (*[]CloudO11ySeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *CloudO11yGlobal) SetSeries(v []CloudO11ySeries)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *CloudO11yGlobal) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetStart

`func (o *CloudO11yGlobal) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CloudO11yGlobal) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CloudO11yGlobal) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *CloudO11yGlobal) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTopModels

`func (o *CloudO11yGlobal) GetTopModels() []CloudO11yModelStat`

GetTopModels returns the TopModels field if non-nil, zero value otherwise.

### GetTopModelsOk

`func (o *CloudO11yGlobal) GetTopModelsOk() (*[]CloudO11yModelStat, bool)`

GetTopModelsOk returns a tuple with the TopModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopModels

`func (o *CloudO11yGlobal) SetTopModels(v []CloudO11yModelStat)`

SetTopModels sets TopModels field to given value.

### HasTopModels

`func (o *CloudO11yGlobal) HasTopModels() bool`

HasTopModels returns a boolean if a field has been set.

### GetTopOrgs

`func (o *CloudO11yGlobal) GetTopOrgs() []CloudO11yOrgStat`

GetTopOrgs returns the TopOrgs field if non-nil, zero value otherwise.

### GetTopOrgsOk

`func (o *CloudO11yGlobal) GetTopOrgsOk() (*[]CloudO11yOrgStat, bool)`

GetTopOrgsOk returns a tuple with the TopOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopOrgs

`func (o *CloudO11yGlobal) SetTopOrgs(v []CloudO11yOrgStat)`

SetTopOrgs sets TopOrgs field to given value.

### HasTopOrgs

`func (o *CloudO11yGlobal) HasTopOrgs() bool`

HasTopOrgs returns a boolean if a field has been set.

### GetTopServices

`func (o *CloudO11yGlobal) GetTopServices() []CloudO11ySvcStat`

GetTopServices returns the TopServices field if non-nil, zero value otherwise.

### GetTopServicesOk

`func (o *CloudO11yGlobal) GetTopServicesOk() (*[]CloudO11ySvcStat, bool)`

GetTopServicesOk returns a tuple with the TopServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopServices

`func (o *CloudO11yGlobal) SetTopServices(v []CloudO11ySvcStat)`

SetTopServices sets TopServices field to given value.

### HasTopServices

`func (o *CloudO11yGlobal) HasTopServices() bool`

HasTopServices returns a boolean if a field has been set.

### GetTotals

`func (o *CloudO11yGlobal) GetTotals() CloudO11yTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *CloudO11yGlobal) GetTotalsOk() (*CloudO11yTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *CloudO11yGlobal) SetTotals(v CloudO11yTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *CloudO11yGlobal) HasTotals() bool`

HasTotals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



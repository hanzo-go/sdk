# Analysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alpha** | Pointer to **float64** | the two-tailed threshold significance was judged at | [optional] 
**Experiment** | Pointer to **string** | the experiment that was analysed | [optional] 
**ExposedTotal** | Pointer to **int64** | subjects enrolled across every arm | [optional] 
**Metric** | Pointer to **string** | the event a conversion is counted from | [optional] 
**Results** | Pointer to [**[]Outcome**](Outcome.md) | one row per declared arm, control first | [optional] 
**Winner** | Pointer to **string** | ADVISORY: the significant, control-beating arm with the highest rate, else empty | [optional] 

## Methods

### NewAnalysis

`func NewAnalysis() *Analysis`

NewAnalysis instantiates a new Analysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisWithDefaults

`func NewAnalysisWithDefaults() *Analysis`

NewAnalysisWithDefaults instantiates a new Analysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlpha

`func (o *Analysis) GetAlpha() float64`

GetAlpha returns the Alpha field if non-nil, zero value otherwise.

### GetAlphaOk

`func (o *Analysis) GetAlphaOk() (*float64, bool)`

GetAlphaOk returns a tuple with the Alpha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpha

`func (o *Analysis) SetAlpha(v float64)`

SetAlpha sets Alpha field to given value.

### HasAlpha

`func (o *Analysis) HasAlpha() bool`

HasAlpha returns a boolean if a field has been set.

### GetExperiment

`func (o *Analysis) GetExperiment() string`

GetExperiment returns the Experiment field if non-nil, zero value otherwise.

### GetExperimentOk

`func (o *Analysis) GetExperimentOk() (*string, bool)`

GetExperimentOk returns a tuple with the Experiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiment

`func (o *Analysis) SetExperiment(v string)`

SetExperiment sets Experiment field to given value.

### HasExperiment

`func (o *Analysis) HasExperiment() bool`

HasExperiment returns a boolean if a field has been set.

### GetExposedTotal

`func (o *Analysis) GetExposedTotal() int64`

GetExposedTotal returns the ExposedTotal field if non-nil, zero value otherwise.

### GetExposedTotalOk

`func (o *Analysis) GetExposedTotalOk() (*int64, bool)`

GetExposedTotalOk returns a tuple with the ExposedTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedTotal

`func (o *Analysis) SetExposedTotal(v int64)`

SetExposedTotal sets ExposedTotal field to given value.

### HasExposedTotal

`func (o *Analysis) HasExposedTotal() bool`

HasExposedTotal returns a boolean if a field has been set.

### GetMetric

`func (o *Analysis) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *Analysis) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *Analysis) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *Analysis) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetResults

`func (o *Analysis) GetResults() []Outcome`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Analysis) GetResultsOk() (*[]Outcome, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Analysis) SetResults(v []Outcome)`

SetResults sets Results field to given value.

### HasResults

`func (o *Analysis) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetWinner

`func (o *Analysis) GetWinner() string`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *Analysis) GetWinnerOk() (*string, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *Analysis) SetWinner(v string)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *Analysis) HasWinner() bool`

HasWinner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



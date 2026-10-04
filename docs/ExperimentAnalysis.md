# ExperimentAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alpha** | Pointer to **float64** | the two-tailed threshold significance was judged at | [optional] 
**Experiment** | Pointer to **string** | the experiment that was analysed | [optional] 
**ExposedTotal** | Pointer to **int64** | subjects enrolled across every arm | [optional] 
**Metric** | Pointer to **string** | the event a conversion is counted from | [optional] 
**Results** | Pointer to [**[]ExperimentOutcome**](ExperimentOutcome.md) | one row per declared arm, control first | [optional] 
**Winner** | Pointer to **string** | ADVISORY: the significant, control-beating arm with the highest rate, else empty | [optional] 

## Methods

### NewExperimentAnalysis

`func NewExperimentAnalysis() *ExperimentAnalysis`

NewExperimentAnalysis instantiates a new ExperimentAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentAnalysisWithDefaults

`func NewExperimentAnalysisWithDefaults() *ExperimentAnalysis`

NewExperimentAnalysisWithDefaults instantiates a new ExperimentAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlpha

`func (o *ExperimentAnalysis) GetAlpha() float64`

GetAlpha returns the Alpha field if non-nil, zero value otherwise.

### GetAlphaOk

`func (o *ExperimentAnalysis) GetAlphaOk() (*float64, bool)`

GetAlphaOk returns a tuple with the Alpha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpha

`func (o *ExperimentAnalysis) SetAlpha(v float64)`

SetAlpha sets Alpha field to given value.

### HasAlpha

`func (o *ExperimentAnalysis) HasAlpha() bool`

HasAlpha returns a boolean if a field has been set.

### GetExperiment

`func (o *ExperimentAnalysis) GetExperiment() string`

GetExperiment returns the Experiment field if non-nil, zero value otherwise.

### GetExperimentOk

`func (o *ExperimentAnalysis) GetExperimentOk() (*string, bool)`

GetExperimentOk returns a tuple with the Experiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiment

`func (o *ExperimentAnalysis) SetExperiment(v string)`

SetExperiment sets Experiment field to given value.

### HasExperiment

`func (o *ExperimentAnalysis) HasExperiment() bool`

HasExperiment returns a boolean if a field has been set.

### GetExposedTotal

`func (o *ExperimentAnalysis) GetExposedTotal() int64`

GetExposedTotal returns the ExposedTotal field if non-nil, zero value otherwise.

### GetExposedTotalOk

`func (o *ExperimentAnalysis) GetExposedTotalOk() (*int64, bool)`

GetExposedTotalOk returns a tuple with the ExposedTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedTotal

`func (o *ExperimentAnalysis) SetExposedTotal(v int64)`

SetExposedTotal sets ExposedTotal field to given value.

### HasExposedTotal

`func (o *ExperimentAnalysis) HasExposedTotal() bool`

HasExposedTotal returns a boolean if a field has been set.

### GetMetric

`func (o *ExperimentAnalysis) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ExperimentAnalysis) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ExperimentAnalysis) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *ExperimentAnalysis) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetResults

`func (o *ExperimentAnalysis) GetResults() []ExperimentOutcome`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ExperimentAnalysis) GetResultsOk() (*[]ExperimentOutcome, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ExperimentAnalysis) SetResults(v []ExperimentOutcome)`

SetResults sets Results field to given value.

### HasResults

`func (o *ExperimentAnalysis) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetWinner

`func (o *ExperimentAnalysis) GetWinner() string`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *ExperimentAnalysis) GetWinnerOk() (*string, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *ExperimentAnalysis) SetWinner(v string)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *ExperimentAnalysis) HasWinner() bool`

HasWinner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



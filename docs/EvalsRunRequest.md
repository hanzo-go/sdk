# EvalsRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | **string** |  | 
**Model** | **string** | Model-under-test (catalog / fine-tuned / BYOM / router) | 
**RunName** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] [default to 20]
**Judge** | Pointer to [**EvalsJudgeSpec**](EvalsJudgeSpec.md) |  | [optional] 
**Metrics** | Pointer to [**[]EvalsMetric**](EvalsMetric.md) |  | [optional] 
**Backend** | Pointer to **string** | Evaluation backend — native Hanzo engine or DigitalOcean GenAI | [optional] [default to "hanzo"]
**Preset** | Pointer to **string** | Named preset to reuse a saved configuration | [optional] 

## Methods

### NewEvalsRunRequest

`func NewEvalsRunRequest(dataset string, model string, ) *EvalsRunRequest`

NewEvalsRunRequest instantiates a new EvalsRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalsRunRequestWithDefaults

`func NewEvalsRunRequestWithDefaults() *EvalsRunRequest`

NewEvalsRunRequestWithDefaults instantiates a new EvalsRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *EvalsRunRequest) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *EvalsRunRequest) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *EvalsRunRequest) SetDataset(v string)`

SetDataset sets Dataset field to given value.


### GetModel

`func (o *EvalsRunRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EvalsRunRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EvalsRunRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetRunName

`func (o *EvalsRunRequest) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *EvalsRunRequest) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *EvalsRunRequest) SetRunName(v string)`

SetRunName sets RunName field to given value.

### HasRunName

`func (o *EvalsRunRequest) HasRunName() bool`

HasRunName returns a boolean if a field has been set.

### GetLimit

`func (o *EvalsRunRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *EvalsRunRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *EvalsRunRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *EvalsRunRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetJudge

`func (o *EvalsRunRequest) GetJudge() EvalsJudgeSpec`

GetJudge returns the Judge field if non-nil, zero value otherwise.

### GetJudgeOk

`func (o *EvalsRunRequest) GetJudgeOk() (*EvalsJudgeSpec, bool)`

GetJudgeOk returns a tuple with the Judge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudge

`func (o *EvalsRunRequest) SetJudge(v EvalsJudgeSpec)`

SetJudge sets Judge field to given value.

### HasJudge

`func (o *EvalsRunRequest) HasJudge() bool`

HasJudge returns a boolean if a field has been set.

### GetMetrics

`func (o *EvalsRunRequest) GetMetrics() []EvalsMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *EvalsRunRequest) GetMetricsOk() (*[]EvalsMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *EvalsRunRequest) SetMetrics(v []EvalsMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *EvalsRunRequest) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetBackend

`func (o *EvalsRunRequest) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *EvalsRunRequest) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *EvalsRunRequest) SetBackend(v string)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *EvalsRunRequest) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetPreset

`func (o *EvalsRunRequest) GetPreset() string`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *EvalsRunRequest) GetPresetOk() (*string, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *EvalsRunRequest) SetPreset(v string)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *EvalsRunRequest) HasPreset() bool`

HasPreset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



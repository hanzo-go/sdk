# AiJudgePanelState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** |  | [optional] 
**Benchmark** | Pointer to [**AiJudgeBenchmark**](AiJudgeBenchmark.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Judges** | Pointer to [**[]AiPanelJudge**](AiPanelJudge.md) |  | [optional] 
**Models** | Pointer to **[]string** |  | [optional] 
**SampleRate** | Pointer to **float32** |  | [optional] 

## Methods

### NewAiJudgePanelState

`func NewAiJudgePanelState() *AiJudgePanelState`

NewAiJudgePanelState instantiates a new AiJudgePanelState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiJudgePanelStateWithDefaults

`func NewAiJudgePanelStateWithDefaults() *AiJudgePanelState`

NewAiJudgePanelStateWithDefaults instantiates a new AiJudgePanelState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *AiJudgePanelState) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *AiJudgePanelState) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *AiJudgePanelState) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *AiJudgePanelState) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBenchmark

`func (o *AiJudgePanelState) GetBenchmark() AiJudgeBenchmark`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *AiJudgePanelState) GetBenchmarkOk() (*AiJudgeBenchmark, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *AiJudgePanelState) SetBenchmark(v AiJudgeBenchmark)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *AiJudgePanelState) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetEnabled

`func (o *AiJudgePanelState) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AiJudgePanelState) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AiJudgePanelState) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AiJudgePanelState) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetJudges

`func (o *AiJudgePanelState) GetJudges() []AiPanelJudge`

GetJudges returns the Judges field if non-nil, zero value otherwise.

### GetJudgesOk

`func (o *AiJudgePanelState) GetJudgesOk() (*[]AiPanelJudge, bool)`

GetJudgesOk returns a tuple with the Judges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudges

`func (o *AiJudgePanelState) SetJudges(v []AiPanelJudge)`

SetJudges sets Judges field to given value.

### HasJudges

`func (o *AiJudgePanelState) HasJudges() bool`

HasJudges returns a boolean if a field has been set.

### GetModels

`func (o *AiJudgePanelState) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *AiJudgePanelState) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *AiJudgePanelState) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *AiJudgePanelState) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetSampleRate

`func (o *AiJudgePanelState) GetSampleRate() float32`

GetSampleRate returns the SampleRate field if non-nil, zero value otherwise.

### GetSampleRateOk

`func (o *AiJudgePanelState) GetSampleRateOk() (*float32, bool)`

GetSampleRateOk returns a tuple with the SampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleRate

`func (o *AiJudgePanelState) SetSampleRate(v float32)`

SetSampleRate sets SampleRate field to given value.

### HasSampleRate

`func (o *AiJudgePanelState) HasSampleRate() bool`

HasSampleRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



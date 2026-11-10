# TrainFamily

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adaptations** | Pointer to **[]string** | Adaptations are the methods it trains by: a create&#39;s adaptation.mode. | [optional] 
**Base** | Pointer to **string** | Base is the base model: &#x60;kai&#x60;, or &#x60;engine&#x60; for the models the engine loads. | [optional] 
**Datasets** | Pointer to **[]string** | Datasets are the dataset URI schemes it reads. | [optional] 
**Devices** | Pointer to **[]string** | Devices are the accelerator kinds it runs on, wherever it runs. | [optional] 
**Executor** | Pointer to **string** | Executor is what runs it. | [optional] 
**Nouns** | Pointer to **[]string** | Nouns are the resources that run it: jobs, clients. | [optional] 
**Outputs** | Pointer to **[]string** | Outputs are the output kinds it produces. | [optional] 
**Protect** | Pointer to **[]string** | Protect are the protections it runs: projection, distillation, capabilities. | [optional] 
**Runs** | Pointer to [**[]TrainRun**](TrainRun.md) | Runs are where it runs, each with what a device-second costs there. | [optional] 
**Terms** | Pointer to **[]string** | Terms are the objective terms it runs. | [optional] 

## Methods

### NewTrainFamily

`func NewTrainFamily() *TrainFamily`

NewTrainFamily instantiates a new TrainFamily object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainFamilyWithDefaults

`func NewTrainFamilyWithDefaults() *TrainFamily`

NewTrainFamilyWithDefaults instantiates a new TrainFamily object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdaptations

`func (o *TrainFamily) GetAdaptations() []string`

GetAdaptations returns the Adaptations field if non-nil, zero value otherwise.

### GetAdaptationsOk

`func (o *TrainFamily) GetAdaptationsOk() (*[]string, bool)`

GetAdaptationsOk returns a tuple with the Adaptations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptations

`func (o *TrainFamily) SetAdaptations(v []string)`

SetAdaptations sets Adaptations field to given value.

### HasAdaptations

`func (o *TrainFamily) HasAdaptations() bool`

HasAdaptations returns a boolean if a field has been set.

### GetBase

`func (o *TrainFamily) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *TrainFamily) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *TrainFamily) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *TrainFamily) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetDatasets

`func (o *TrainFamily) GetDatasets() []string`

GetDatasets returns the Datasets field if non-nil, zero value otherwise.

### GetDatasetsOk

`func (o *TrainFamily) GetDatasetsOk() (*[]string, bool)`

GetDatasetsOk returns a tuple with the Datasets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasets

`func (o *TrainFamily) SetDatasets(v []string)`

SetDatasets sets Datasets field to given value.

### HasDatasets

`func (o *TrainFamily) HasDatasets() bool`

HasDatasets returns a boolean if a field has been set.

### GetDevices

`func (o *TrainFamily) GetDevices() []string`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *TrainFamily) GetDevicesOk() (*[]string, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *TrainFamily) SetDevices(v []string)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *TrainFamily) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetExecutor

`func (o *TrainFamily) GetExecutor() string`

GetExecutor returns the Executor field if non-nil, zero value otherwise.

### GetExecutorOk

`func (o *TrainFamily) GetExecutorOk() (*string, bool)`

GetExecutorOk returns a tuple with the Executor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutor

`func (o *TrainFamily) SetExecutor(v string)`

SetExecutor sets Executor field to given value.

### HasExecutor

`func (o *TrainFamily) HasExecutor() bool`

HasExecutor returns a boolean if a field has been set.

### GetNouns

`func (o *TrainFamily) GetNouns() []string`

GetNouns returns the Nouns field if non-nil, zero value otherwise.

### GetNounsOk

`func (o *TrainFamily) GetNounsOk() (*[]string, bool)`

GetNounsOk returns a tuple with the Nouns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNouns

`func (o *TrainFamily) SetNouns(v []string)`

SetNouns sets Nouns field to given value.

### HasNouns

`func (o *TrainFamily) HasNouns() bool`

HasNouns returns a boolean if a field has been set.

### GetOutputs

`func (o *TrainFamily) GetOutputs() []string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *TrainFamily) GetOutputsOk() (*[]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *TrainFamily) SetOutputs(v []string)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *TrainFamily) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetProtect

`func (o *TrainFamily) GetProtect() []string`

GetProtect returns the Protect field if non-nil, zero value otherwise.

### GetProtectOk

`func (o *TrainFamily) GetProtectOk() (*[]string, bool)`

GetProtectOk returns a tuple with the Protect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtect

`func (o *TrainFamily) SetProtect(v []string)`

SetProtect sets Protect field to given value.

### HasProtect

`func (o *TrainFamily) HasProtect() bool`

HasProtect returns a boolean if a field has been set.

### GetRuns

`func (o *TrainFamily) GetRuns() []TrainRun`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *TrainFamily) GetRunsOk() (*[]TrainRun, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *TrainFamily) SetRuns(v []TrainRun)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *TrainFamily) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetTerms

`func (o *TrainFamily) GetTerms() []string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *TrainFamily) GetTermsOk() (*[]string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *TrainFamily) SetTerms(v []string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *TrainFamily) HasTerms() bool`

HasTerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



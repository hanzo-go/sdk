# EngineListPipelineRuns200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Runs** | Pointer to [**[]EnginePipelineRun**](EnginePipelineRun.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewEngineListPipelineRuns200Response

`func NewEngineListPipelineRuns200Response() *EngineListPipelineRuns200Response`

NewEngineListPipelineRuns200Response instantiates a new EngineListPipelineRuns200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineListPipelineRuns200ResponseWithDefaults

`func NewEngineListPipelineRuns200ResponseWithDefaults() *EngineListPipelineRuns200Response`

NewEngineListPipelineRuns200ResponseWithDefaults instantiates a new EngineListPipelineRuns200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuns

`func (o *EngineListPipelineRuns200Response) GetRuns() []EnginePipelineRun`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *EngineListPipelineRuns200Response) GetRunsOk() (*[]EnginePipelineRun, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *EngineListPipelineRuns200Response) SetRuns(v []EnginePipelineRun)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *EngineListPipelineRuns200Response) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetTotal

`func (o *EngineListPipelineRuns200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EngineListPipelineRuns200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EngineListPipelineRuns200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *EngineListPipelineRuns200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



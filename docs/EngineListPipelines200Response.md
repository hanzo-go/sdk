# EngineListPipelines200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pipelines** | Pointer to [**[]EnginePipeline**](EnginePipeline.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewEngineListPipelines200Response

`func NewEngineListPipelines200Response() *EngineListPipelines200Response`

NewEngineListPipelines200Response instantiates a new EngineListPipelines200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineListPipelines200ResponseWithDefaults

`func NewEngineListPipelines200ResponseWithDefaults() *EngineListPipelines200Response`

NewEngineListPipelines200ResponseWithDefaults instantiates a new EngineListPipelines200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipelines

`func (o *EngineListPipelines200Response) GetPipelines() []EnginePipeline`

GetPipelines returns the Pipelines field if non-nil, zero value otherwise.

### GetPipelinesOk

`func (o *EngineListPipelines200Response) GetPipelinesOk() (*[]EnginePipeline, bool)`

GetPipelinesOk returns a tuple with the Pipelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelines

`func (o *EngineListPipelines200Response) SetPipelines(v []EnginePipeline)`

SetPipelines sets Pipelines field to given value.

### HasPipelines

`func (o *EngineListPipelines200Response) HasPipelines() bool`

HasPipelines returns a boolean if a field has been set.

### GetTotal

`func (o *EngineListPipelines200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EngineListPipelines200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EngineListPipelines200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *EngineListPipelines200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



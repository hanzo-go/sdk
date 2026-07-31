# MlListPipelines200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pipelines** | Pointer to [**[]MlPipeline**](MlPipeline.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewMlListPipelines200Response

`func NewMlListPipelines200Response() *MlListPipelines200Response`

NewMlListPipelines200Response instantiates a new MlListPipelines200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlListPipelines200ResponseWithDefaults

`func NewMlListPipelines200ResponseWithDefaults() *MlListPipelines200Response`

NewMlListPipelines200ResponseWithDefaults instantiates a new MlListPipelines200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipelines

`func (o *MlListPipelines200Response) GetPipelines() []MlPipeline`

GetPipelines returns the Pipelines field if non-nil, zero value otherwise.

### GetPipelinesOk

`func (o *MlListPipelines200Response) GetPipelinesOk() (*[]MlPipeline, bool)`

GetPipelinesOk returns a tuple with the Pipelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelines

`func (o *MlListPipelines200Response) SetPipelines(v []MlPipeline)`

SetPipelines sets Pipelines field to given value.

### HasPipelines

`func (o *MlListPipelines200Response) HasPipelines() bool`

HasPipelines returns a boolean if a field has been set.

### GetTotal

`func (o *MlListPipelines200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MlListPipelines200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MlListPipelines200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MlListPipelines200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



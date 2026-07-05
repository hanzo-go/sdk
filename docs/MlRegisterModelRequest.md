# MlRegisterModelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Version** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**SourceRunId** | Pointer to **string** |  | [optional] 
**Artifacts** | Pointer to [**MlRegisterModelRequestArtifacts**](MlRegisterModelRequestArtifacts.md) |  | [optional] 
**Metrics** | Pointer to **map[string]float32** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMlRegisterModelRequest

`func NewMlRegisterModelRequest(name string, version string, ) *MlRegisterModelRequest`

NewMlRegisterModelRequest instantiates a new MlRegisterModelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlRegisterModelRequestWithDefaults

`func NewMlRegisterModelRequestWithDefaults() *MlRegisterModelRequest`

NewMlRegisterModelRequestWithDefaults instantiates a new MlRegisterModelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MlRegisterModelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MlRegisterModelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MlRegisterModelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *MlRegisterModelRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MlRegisterModelRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MlRegisterModelRequest) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDescription

`func (o *MlRegisterModelRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MlRegisterModelRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MlRegisterModelRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MlRegisterModelRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSourceRunId

`func (o *MlRegisterModelRequest) GetSourceRunId() string`

GetSourceRunId returns the SourceRunId field if non-nil, zero value otherwise.

### GetSourceRunIdOk

`func (o *MlRegisterModelRequest) GetSourceRunIdOk() (*string, bool)`

GetSourceRunIdOk returns a tuple with the SourceRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRunId

`func (o *MlRegisterModelRequest) SetSourceRunId(v string)`

SetSourceRunId sets SourceRunId field to given value.

### HasSourceRunId

`func (o *MlRegisterModelRequest) HasSourceRunId() bool`

HasSourceRunId returns a boolean if a field has been set.

### GetArtifacts

`func (o *MlRegisterModelRequest) GetArtifacts() MlRegisterModelRequestArtifacts`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *MlRegisterModelRequest) GetArtifactsOk() (*MlRegisterModelRequestArtifacts, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *MlRegisterModelRequest) SetArtifacts(v MlRegisterModelRequestArtifacts)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *MlRegisterModelRequest) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetMetrics

`func (o *MlRegisterModelRequest) GetMetrics() map[string]float32`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MlRegisterModelRequest) GetMetricsOk() (*map[string]float32, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MlRegisterModelRequest) SetMetrics(v map[string]float32)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *MlRegisterModelRequest) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetTags

`func (o *MlRegisterModelRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MlRegisterModelRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MlRegisterModelRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MlRegisterModelRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



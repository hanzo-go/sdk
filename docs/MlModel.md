# MlModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Stage** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**SourceRunId** | Pointer to **string** | Experiment run that produced this model | [optional] 
**Artifacts** | Pointer to [**MlRegisterModelRequestArtifacts**](MlRegisterModelRequestArtifacts.md) |  | [optional] 
**Metrics** | Pointer to **map[string]float32** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMlModel

`func NewMlModel() *MlModel`

NewMlModel instantiates a new MlModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlModelWithDefaults

`func NewMlModelWithDefaults() *MlModel`

NewMlModelWithDefaults instantiates a new MlModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MlModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MlModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MlModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MlModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MlModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MlModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MlModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MlModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *MlModel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MlModel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MlModel) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MlModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStage

`func (o *MlModel) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *MlModel) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *MlModel) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *MlModel) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetDescription

`func (o *MlModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MlModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MlModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MlModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSourceRunId

`func (o *MlModel) GetSourceRunId() string`

GetSourceRunId returns the SourceRunId field if non-nil, zero value otherwise.

### GetSourceRunIdOk

`func (o *MlModel) GetSourceRunIdOk() (*string, bool)`

GetSourceRunIdOk returns a tuple with the SourceRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRunId

`func (o *MlModel) SetSourceRunId(v string)`

SetSourceRunId sets SourceRunId field to given value.

### HasSourceRunId

`func (o *MlModel) HasSourceRunId() bool`

HasSourceRunId returns a boolean if a field has been set.

### GetArtifacts

`func (o *MlModel) GetArtifacts() MlRegisterModelRequestArtifacts`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *MlModel) GetArtifactsOk() (*MlRegisterModelRequestArtifacts, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *MlModel) SetArtifacts(v MlRegisterModelRequestArtifacts)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *MlModel) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetMetrics

`func (o *MlModel) GetMetrics() map[string]float32`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MlModel) GetMetricsOk() (*map[string]float32, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MlModel) SetMetrics(v map[string]float32)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *MlModel) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetTags

`func (o *MlModel) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MlModel) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MlModel) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MlModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MlModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MlModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MlModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MlModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



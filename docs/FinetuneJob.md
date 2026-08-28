# FinetuneJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseModel** | Pointer to **string** |  | [optional] 
**CostCents** | Pointer to **int32** |  | [optional] 
**CrName** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Dataset** | Pointer to **string** |  | [optional] 
**DeployUrl** | Pointer to **string** |  | [optional] 
**DeployedModel** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**FinishedTime** | Pointer to **string** |  | [optional] 
**GpuCount** | Pointer to **int32** |  | [optional] 
**GpuSeconds** | Pointer to **int32** |  | [optional] 
**GpuType** | Pointer to **string** |  | [optional] 
**Hyperparams** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Metered** | Pointer to **bool** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**NumNodes** | Pointer to **int32** |  | [optional] 
**OutputUri** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Preset** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**Runtime** | Pointer to **string** |  | [optional] 
**StartedTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Task** | Pointer to **string** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 

## Methods

### NewFinetuneJob

`func NewFinetuneJob() *FinetuneJob`

NewFinetuneJob instantiates a new FinetuneJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinetuneJobWithDefaults

`func NewFinetuneJobWithDefaults() *FinetuneJob`

NewFinetuneJobWithDefaults instantiates a new FinetuneJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseModel

`func (o *FinetuneJob) GetBaseModel() string`

GetBaseModel returns the BaseModel field if non-nil, zero value otherwise.

### GetBaseModelOk

`func (o *FinetuneJob) GetBaseModelOk() (*string, bool)`

GetBaseModelOk returns a tuple with the BaseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseModel

`func (o *FinetuneJob) SetBaseModel(v string)`

SetBaseModel sets BaseModel field to given value.

### HasBaseModel

`func (o *FinetuneJob) HasBaseModel() bool`

HasBaseModel returns a boolean if a field has been set.

### GetCostCents

`func (o *FinetuneJob) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *FinetuneJob) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *FinetuneJob) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *FinetuneJob) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetCrName

`func (o *FinetuneJob) GetCrName() string`

GetCrName returns the CrName field if non-nil, zero value otherwise.

### GetCrNameOk

`func (o *FinetuneJob) GetCrNameOk() (*string, bool)`

GetCrNameOk returns a tuple with the CrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrName

`func (o *FinetuneJob) SetCrName(v string)`

SetCrName sets CrName field to given value.

### HasCrName

`func (o *FinetuneJob) HasCrName() bool`

HasCrName returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FinetuneJob) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FinetuneJob) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FinetuneJob) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FinetuneJob) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedTime

`func (o *FinetuneJob) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *FinetuneJob) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *FinetuneJob) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *FinetuneJob) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDataset

`func (o *FinetuneJob) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *FinetuneJob) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *FinetuneJob) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *FinetuneJob) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetDeployUrl

`func (o *FinetuneJob) GetDeployUrl() string`

GetDeployUrl returns the DeployUrl field if non-nil, zero value otherwise.

### GetDeployUrlOk

`func (o *FinetuneJob) GetDeployUrlOk() (*string, bool)`

GetDeployUrlOk returns a tuple with the DeployUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployUrl

`func (o *FinetuneJob) SetDeployUrl(v string)`

SetDeployUrl sets DeployUrl field to given value.

### HasDeployUrl

`func (o *FinetuneJob) HasDeployUrl() bool`

HasDeployUrl returns a boolean if a field has been set.

### GetDeployedModel

`func (o *FinetuneJob) GetDeployedModel() string`

GetDeployedModel returns the DeployedModel field if non-nil, zero value otherwise.

### GetDeployedModelOk

`func (o *FinetuneJob) GetDeployedModelOk() (*string, bool)`

GetDeployedModelOk returns a tuple with the DeployedModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedModel

`func (o *FinetuneJob) SetDeployedModel(v string)`

SetDeployedModel sets DeployedModel field to given value.

### HasDeployedModel

`func (o *FinetuneJob) HasDeployedModel() bool`

HasDeployedModel returns a boolean if a field has been set.

### GetDisplayName

`func (o *FinetuneJob) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FinetuneJob) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FinetuneJob) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FinetuneJob) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetError

`func (o *FinetuneJob) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FinetuneJob) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FinetuneJob) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *FinetuneJob) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFinishedTime

`func (o *FinetuneJob) GetFinishedTime() string`

GetFinishedTime returns the FinishedTime field if non-nil, zero value otherwise.

### GetFinishedTimeOk

`func (o *FinetuneJob) GetFinishedTimeOk() (*string, bool)`

GetFinishedTimeOk returns a tuple with the FinishedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedTime

`func (o *FinetuneJob) SetFinishedTime(v string)`

SetFinishedTime sets FinishedTime field to given value.

### HasFinishedTime

`func (o *FinetuneJob) HasFinishedTime() bool`

HasFinishedTime returns a boolean if a field has been set.

### GetGpuCount

`func (o *FinetuneJob) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *FinetuneJob) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *FinetuneJob) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *FinetuneJob) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetGpuSeconds

`func (o *FinetuneJob) GetGpuSeconds() int32`

GetGpuSeconds returns the GpuSeconds field if non-nil, zero value otherwise.

### GetGpuSecondsOk

`func (o *FinetuneJob) GetGpuSecondsOk() (*int32, bool)`

GetGpuSecondsOk returns a tuple with the GpuSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuSeconds

`func (o *FinetuneJob) SetGpuSeconds(v int32)`

SetGpuSeconds sets GpuSeconds field to given value.

### HasGpuSeconds

`func (o *FinetuneJob) HasGpuSeconds() bool`

HasGpuSeconds returns a boolean if a field has been set.

### GetGpuType

`func (o *FinetuneJob) GetGpuType() string`

GetGpuType returns the GpuType field if non-nil, zero value otherwise.

### GetGpuTypeOk

`func (o *FinetuneJob) GetGpuTypeOk() (*string, bool)`

GetGpuTypeOk returns a tuple with the GpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuType

`func (o *FinetuneJob) SetGpuType(v string)`

SetGpuType sets GpuType field to given value.

### HasGpuType

`func (o *FinetuneJob) HasGpuType() bool`

HasGpuType returns a boolean if a field has been set.

### GetHyperparams

`func (o *FinetuneJob) GetHyperparams() string`

GetHyperparams returns the Hyperparams field if non-nil, zero value otherwise.

### GetHyperparamsOk

`func (o *FinetuneJob) GetHyperparamsOk() (*string, bool)`

GetHyperparamsOk returns a tuple with the Hyperparams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperparams

`func (o *FinetuneJob) SetHyperparams(v string)`

SetHyperparams sets Hyperparams field to given value.

### HasHyperparams

`func (o *FinetuneJob) HasHyperparams() bool`

HasHyperparams returns a boolean if a field has been set.

### GetMessage

`func (o *FinetuneJob) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FinetuneJob) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FinetuneJob) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FinetuneJob) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetered

`func (o *FinetuneJob) GetMetered() bool`

GetMetered returns the Metered field if non-nil, zero value otherwise.

### GetMeteredOk

`func (o *FinetuneJob) GetMeteredOk() (*bool, bool)`

GetMeteredOk returns a tuple with the Metered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetered

`func (o *FinetuneJob) SetMetered(v bool)`

SetMetered sets Metered field to given value.

### HasMetered

`func (o *FinetuneJob) HasMetered() bool`

HasMetered returns a boolean if a field has been set.

### GetMethod

`func (o *FinetuneJob) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *FinetuneJob) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *FinetuneJob) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *FinetuneJob) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetName

`func (o *FinetuneJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FinetuneJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FinetuneJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FinetuneJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *FinetuneJob) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *FinetuneJob) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *FinetuneJob) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *FinetuneJob) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNumNodes

`func (o *FinetuneJob) GetNumNodes() int32`

GetNumNodes returns the NumNodes field if non-nil, zero value otherwise.

### GetNumNodesOk

`func (o *FinetuneJob) GetNumNodesOk() (*int32, bool)`

GetNumNodesOk returns a tuple with the NumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNodes

`func (o *FinetuneJob) SetNumNodes(v int32)`

SetNumNodes sets NumNodes field to given value.

### HasNumNodes

`func (o *FinetuneJob) HasNumNodes() bool`

HasNumNodes returns a boolean if a field has been set.

### GetOutputUri

`func (o *FinetuneJob) GetOutputUri() string`

GetOutputUri returns the OutputUri field if non-nil, zero value otherwise.

### GetOutputUriOk

`func (o *FinetuneJob) GetOutputUriOk() (*string, bool)`

GetOutputUriOk returns a tuple with the OutputUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputUri

`func (o *FinetuneJob) SetOutputUri(v string)`

SetOutputUri sets OutputUri field to given value.

### HasOutputUri

`func (o *FinetuneJob) HasOutputUri() bool`

HasOutputUri returns a boolean if a field has been set.

### GetOwner

`func (o *FinetuneJob) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *FinetuneJob) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *FinetuneJob) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *FinetuneJob) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPreset

`func (o *FinetuneJob) GetPreset() string`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *FinetuneJob) GetPresetOk() (*string, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *FinetuneJob) SetPreset(v string)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *FinetuneJob) HasPreset() bool`

HasPreset returns a boolean if a field has been set.

### GetProgress

`func (o *FinetuneJob) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *FinetuneJob) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *FinetuneJob) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *FinetuneJob) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetRuntime

`func (o *FinetuneJob) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *FinetuneJob) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *FinetuneJob) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *FinetuneJob) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetStartedTime

`func (o *FinetuneJob) GetStartedTime() string`

GetStartedTime returns the StartedTime field if non-nil, zero value otherwise.

### GetStartedTimeOk

`func (o *FinetuneJob) GetStartedTimeOk() (*string, bool)`

GetStartedTimeOk returns a tuple with the StartedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedTime

`func (o *FinetuneJob) SetStartedTime(v string)`

SetStartedTime sets StartedTime field to given value.

### HasStartedTime

`func (o *FinetuneJob) HasStartedTime() bool`

HasStartedTime returns a boolean if a field has been set.

### GetStatus

`func (o *FinetuneJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FinetuneJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FinetuneJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FinetuneJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTask

`func (o *FinetuneJob) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *FinetuneJob) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *FinetuneJob) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *FinetuneJob) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *FinetuneJob) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *FinetuneJob) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *FinetuneJob) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *FinetuneJob) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



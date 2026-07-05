# EngineJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** | Container image | [optional] 
**Command** | Pointer to **[]string** |  | [optional] 
**Resources** | Pointer to [**EngineJobResources**](EngineJobResources.md) |  | [optional] 
**Env** | Pointer to **map[string]string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] [default to 0]
**MaxRetries** | Pointer to **int32** |  | [optional] [default to 0]
**TimeoutSeconds** | Pointer to **int32** |  | [optional] 
**Metrics** | Pointer to [**EngineJobMetrics**](EngineJobMetrics.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEngineJob

`func NewEngineJob() *EngineJob`

NewEngineJob instantiates a new EngineJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineJobWithDefaults

`func NewEngineJobWithDefaults() *EngineJob`

NewEngineJobWithDefaults instantiates a new EngineJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EngineJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EngineJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EngineJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EngineJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EngineJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngineJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngineJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EngineJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *EngineJob) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EngineJob) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EngineJob) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EngineJob) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *EngineJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EngineJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EngineJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EngineJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetClusterId

`func (o *EngineJob) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EngineJob) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EngineJob) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *EngineJob) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetImage

`func (o *EngineJob) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *EngineJob) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *EngineJob) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *EngineJob) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetCommand

`func (o *EngineJob) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *EngineJob) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *EngineJob) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *EngineJob) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetResources

`func (o *EngineJob) GetResources() EngineJobResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *EngineJob) GetResourcesOk() (*EngineJobResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *EngineJob) SetResources(v EngineJobResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *EngineJob) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetEnv

`func (o *EngineJob) GetEnv() map[string]string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *EngineJob) GetEnvOk() (*map[string]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *EngineJob) SetEnv(v map[string]string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *EngineJob) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetPriority

`func (o *EngineJob) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EngineJob) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EngineJob) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *EngineJob) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetMaxRetries

`func (o *EngineJob) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *EngineJob) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *EngineJob) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *EngineJob) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *EngineJob) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *EngineJob) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *EngineJob) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *EngineJob) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetMetrics

`func (o *EngineJob) GetMetrics() EngineJobMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *EngineJob) GetMetricsOk() (*EngineJobMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *EngineJob) SetMetrics(v EngineJobMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *EngineJob) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetStartedAt

`func (o *EngineJob) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *EngineJob) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *EngineJob) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *EngineJob) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *EngineJob) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *EngineJob) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *EngineJob) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *EngineJob) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EngineJob) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EngineJob) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EngineJob) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EngineJob) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



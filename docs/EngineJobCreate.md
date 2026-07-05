# EngineJobCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**ClusterId** | Pointer to **string** | Target cluster (auto-selected if omitted) | [optional] 
**Image** | **string** |  | 
**Command** | Pointer to **[]string** |  | [optional] 
**Resources** | [**EngineJobResources**](EngineJobResources.md) |  | 
**Env** | Pointer to **map[string]string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] [default to 0]
**MaxRetries** | Pointer to **int32** |  | [optional] [default to 0]
**TimeoutSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewEngineJobCreate

`func NewEngineJobCreate(name string, type_ string, image string, resources EngineJobResources, ) *EngineJobCreate`

NewEngineJobCreate instantiates a new EngineJobCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineJobCreateWithDefaults

`func NewEngineJobCreateWithDefaults() *EngineJobCreate`

NewEngineJobCreateWithDefaults instantiates a new EngineJobCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EngineJobCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngineJobCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngineJobCreate) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *EngineJobCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EngineJobCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EngineJobCreate) SetType(v string)`

SetType sets Type field to given value.


### GetClusterId

`func (o *EngineJobCreate) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EngineJobCreate) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EngineJobCreate) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *EngineJobCreate) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetImage

`func (o *EngineJobCreate) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *EngineJobCreate) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *EngineJobCreate) SetImage(v string)`

SetImage sets Image field to given value.


### GetCommand

`func (o *EngineJobCreate) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *EngineJobCreate) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *EngineJobCreate) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *EngineJobCreate) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetResources

`func (o *EngineJobCreate) GetResources() EngineJobResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *EngineJobCreate) GetResourcesOk() (*EngineJobResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *EngineJobCreate) SetResources(v EngineJobResources)`

SetResources sets Resources field to given value.


### GetEnv

`func (o *EngineJobCreate) GetEnv() map[string]string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *EngineJobCreate) GetEnvOk() (*map[string]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *EngineJobCreate) SetEnv(v map[string]string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *EngineJobCreate) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetPriority

`func (o *EngineJobCreate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EngineJobCreate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EngineJobCreate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *EngineJobCreate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetMaxRetries

`func (o *EngineJobCreate) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *EngineJobCreate) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *EngineJobCreate) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *EngineJobCreate) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *EngineJobCreate) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *EngineJobCreate) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *EngineJobCreate) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *EngineJobCreate) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



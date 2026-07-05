# KmsCreateSecretSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Destination** | **string** |  | 
**ProjectId** | **string** |  | 
**SourceEnvironment** | **string** |  | 
**SourcePath** | **string** |  | 
**ConnectionId** | **string** |  | 
**IsAutoSyncEnabled** | Pointer to **bool** |  | [optional] [default to true]
**DestinationConfig** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewKmsCreateSecretSyncRequest

`func NewKmsCreateSecretSyncRequest(name string, destination string, projectId string, sourceEnvironment string, sourcePath string, connectionId string, ) *KmsCreateSecretSyncRequest`

NewKmsCreateSecretSyncRequest instantiates a new KmsCreateSecretSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCreateSecretSyncRequestWithDefaults

`func NewKmsCreateSecretSyncRequestWithDefaults() *KmsCreateSecretSyncRequest`

NewKmsCreateSecretSyncRequestWithDefaults instantiates a new KmsCreateSecretSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KmsCreateSecretSyncRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsCreateSecretSyncRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsCreateSecretSyncRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDestination

`func (o *KmsCreateSecretSyncRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *KmsCreateSecretSyncRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *KmsCreateSecretSyncRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetProjectId

`func (o *KmsCreateSecretSyncRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *KmsCreateSecretSyncRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *KmsCreateSecretSyncRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSourceEnvironment

`func (o *KmsCreateSecretSyncRequest) GetSourceEnvironment() string`

GetSourceEnvironment returns the SourceEnvironment field if non-nil, zero value otherwise.

### GetSourceEnvironmentOk

`func (o *KmsCreateSecretSyncRequest) GetSourceEnvironmentOk() (*string, bool)`

GetSourceEnvironmentOk returns a tuple with the SourceEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnvironment

`func (o *KmsCreateSecretSyncRequest) SetSourceEnvironment(v string)`

SetSourceEnvironment sets SourceEnvironment field to given value.


### GetSourcePath

`func (o *KmsCreateSecretSyncRequest) GetSourcePath() string`

GetSourcePath returns the SourcePath field if non-nil, zero value otherwise.

### GetSourcePathOk

`func (o *KmsCreateSecretSyncRequest) GetSourcePathOk() (*string, bool)`

GetSourcePathOk returns a tuple with the SourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePath

`func (o *KmsCreateSecretSyncRequest) SetSourcePath(v string)`

SetSourcePath sets SourcePath field to given value.


### GetConnectionId

`func (o *KmsCreateSecretSyncRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *KmsCreateSecretSyncRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *KmsCreateSecretSyncRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetIsAutoSyncEnabled

`func (o *KmsCreateSecretSyncRequest) GetIsAutoSyncEnabled() bool`

GetIsAutoSyncEnabled returns the IsAutoSyncEnabled field if non-nil, zero value otherwise.

### GetIsAutoSyncEnabledOk

`func (o *KmsCreateSecretSyncRequest) GetIsAutoSyncEnabledOk() (*bool, bool)`

GetIsAutoSyncEnabledOk returns a tuple with the IsAutoSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoSyncEnabled

`func (o *KmsCreateSecretSyncRequest) SetIsAutoSyncEnabled(v bool)`

SetIsAutoSyncEnabled sets IsAutoSyncEnabled field to given value.

### HasIsAutoSyncEnabled

`func (o *KmsCreateSecretSyncRequest) HasIsAutoSyncEnabled() bool`

HasIsAutoSyncEnabled returns a boolean if a field has been set.

### GetDestinationConfig

`func (o *KmsCreateSecretSyncRequest) GetDestinationConfig() map[string]interface{}`

GetDestinationConfig returns the DestinationConfig field if non-nil, zero value otherwise.

### GetDestinationConfigOk

`func (o *KmsCreateSecretSyncRequest) GetDestinationConfigOk() (*map[string]interface{}, bool)`

GetDestinationConfigOk returns a tuple with the DestinationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationConfig

`func (o *KmsCreateSecretSyncRequest) SetDestinationConfig(v map[string]interface{})`

SetDestinationConfig sets DestinationConfig field to given value.

### HasDestinationConfig

`func (o *KmsCreateSecretSyncRequest) HasDestinationConfig() bool`

HasDestinationConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



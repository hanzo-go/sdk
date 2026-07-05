# KmsDeleteSecretFolderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**Environment** | **string** |  | 
**Directory** | Pointer to **string** |  | [optional] 

## Methods

### NewKmsDeleteSecretFolderRequest

`func NewKmsDeleteSecretFolderRequest(workspaceId string, environment string, ) *KmsDeleteSecretFolderRequest`

NewKmsDeleteSecretFolderRequest instantiates a new KmsDeleteSecretFolderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsDeleteSecretFolderRequestWithDefaults

`func NewKmsDeleteSecretFolderRequestWithDefaults() *KmsDeleteSecretFolderRequest`

NewKmsDeleteSecretFolderRequestWithDefaults instantiates a new KmsDeleteSecretFolderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *KmsDeleteSecretFolderRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *KmsDeleteSecretFolderRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *KmsDeleteSecretFolderRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetEnvironment

`func (o *KmsDeleteSecretFolderRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KmsDeleteSecretFolderRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KmsDeleteSecretFolderRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetDirectory

`func (o *KmsDeleteSecretFolderRequest) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *KmsDeleteSecretFolderRequest) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *KmsDeleteSecretFolderRequest) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *KmsDeleteSecretFolderRequest) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



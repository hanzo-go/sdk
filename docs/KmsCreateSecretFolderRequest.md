# KmsCreateSecretFolderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**Environment** | **string** |  | 
**Name** | **string** |  | 
**Directory** | Pointer to **string** |  | [optional] [default to "/"]

## Methods

### NewKmsCreateSecretFolderRequest

`func NewKmsCreateSecretFolderRequest(workspaceId string, environment string, name string, ) *KmsCreateSecretFolderRequest`

NewKmsCreateSecretFolderRequest instantiates a new KmsCreateSecretFolderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCreateSecretFolderRequestWithDefaults

`func NewKmsCreateSecretFolderRequestWithDefaults() *KmsCreateSecretFolderRequest`

NewKmsCreateSecretFolderRequestWithDefaults instantiates a new KmsCreateSecretFolderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *KmsCreateSecretFolderRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *KmsCreateSecretFolderRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *KmsCreateSecretFolderRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetEnvironment

`func (o *KmsCreateSecretFolderRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KmsCreateSecretFolderRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KmsCreateSecretFolderRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetName

`func (o *KmsCreateSecretFolderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsCreateSecretFolderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsCreateSecretFolderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDirectory

`func (o *KmsCreateSecretFolderRequest) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *KmsCreateSecretFolderRequest) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *KmsCreateSecretFolderRequest) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *KmsCreateSecretFolderRequest) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



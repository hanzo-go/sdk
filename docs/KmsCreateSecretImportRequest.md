# KmsCreateSecretImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**Environment** | **string** |  | 
**Directory** | **string** |  | [default to "/"]
**Import** | [**KmsCreateSecretImportRequestImport**](KmsCreateSecretImportRequestImport.md) |  | 

## Methods

### NewKmsCreateSecretImportRequest

`func NewKmsCreateSecretImportRequest(workspaceId string, environment string, directory string, import_ KmsCreateSecretImportRequestImport, ) *KmsCreateSecretImportRequest`

NewKmsCreateSecretImportRequest instantiates a new KmsCreateSecretImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCreateSecretImportRequestWithDefaults

`func NewKmsCreateSecretImportRequestWithDefaults() *KmsCreateSecretImportRequest`

NewKmsCreateSecretImportRequestWithDefaults instantiates a new KmsCreateSecretImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *KmsCreateSecretImportRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *KmsCreateSecretImportRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *KmsCreateSecretImportRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetEnvironment

`func (o *KmsCreateSecretImportRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KmsCreateSecretImportRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KmsCreateSecretImportRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetDirectory

`func (o *KmsCreateSecretImportRequest) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *KmsCreateSecretImportRequest) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *KmsCreateSecretImportRequest) SetDirectory(v string)`

SetDirectory sets Directory field to given value.


### GetImport

`func (o *KmsCreateSecretImportRequest) GetImport() KmsCreateSecretImportRequestImport`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *KmsCreateSecretImportRequest) GetImportOk() (*KmsCreateSecretImportRequestImport, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *KmsCreateSecretImportRequest) SetImport(v KmsCreateSecretImportRequestImport)`

SetImport sets Import field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



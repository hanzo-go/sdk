# ProjectsCompleteDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Final status. | 
**Commit** | Pointer to **string** |  | [optional] 
**LiveUrl** | Pointer to **string** | Defaults to the project&#39;s canonical live URL when omitted on a live completion. | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Files** | Pointer to **int32** |  | [optional] 
**Bytes** | Pointer to **int64** |  | [optional] 

## Methods

### NewProjectsCompleteDeploymentRequest

`func NewProjectsCompleteDeploymentRequest(status string, ) *ProjectsCompleteDeploymentRequest`

NewProjectsCompleteDeploymentRequest instantiates a new ProjectsCompleteDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsCompleteDeploymentRequestWithDefaults

`func NewProjectsCompleteDeploymentRequestWithDefaults() *ProjectsCompleteDeploymentRequest`

NewProjectsCompleteDeploymentRequestWithDefaults instantiates a new ProjectsCompleteDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ProjectsCompleteDeploymentRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectsCompleteDeploymentRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectsCompleteDeploymentRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCommit

`func (o *ProjectsCompleteDeploymentRequest) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ProjectsCompleteDeploymentRequest) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ProjectsCompleteDeploymentRequest) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *ProjectsCompleteDeploymentRequest) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetLiveUrl

`func (o *ProjectsCompleteDeploymentRequest) GetLiveUrl() string`

GetLiveUrl returns the LiveUrl field if non-nil, zero value otherwise.

### GetLiveUrlOk

`func (o *ProjectsCompleteDeploymentRequest) GetLiveUrlOk() (*string, bool)`

GetLiveUrlOk returns a tuple with the LiveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveUrl

`func (o *ProjectsCompleteDeploymentRequest) SetLiveUrl(v string)`

SetLiveUrl sets LiveUrl field to given value.

### HasLiveUrl

`func (o *ProjectsCompleteDeploymentRequest) HasLiveUrl() bool`

HasLiveUrl returns a boolean if a field has been set.

### GetMessage

`func (o *ProjectsCompleteDeploymentRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProjectsCompleteDeploymentRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProjectsCompleteDeploymentRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProjectsCompleteDeploymentRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFiles

`func (o *ProjectsCompleteDeploymentRequest) GetFiles() int32`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ProjectsCompleteDeploymentRequest) GetFilesOk() (*int32, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ProjectsCompleteDeploymentRequest) SetFiles(v int32)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ProjectsCompleteDeploymentRequest) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetBytes

`func (o *ProjectsCompleteDeploymentRequest) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ProjectsCompleteDeploymentRequest) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ProjectsCompleteDeploymentRequest) SetBytes(v int64)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *ProjectsCompleteDeploymentRequest) HasBytes() bool`

HasBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



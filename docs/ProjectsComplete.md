# ProjectsComplete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **int32** | Bytes is their total size in bytes. | [optional] 
**Commit** | Pointer to **string** | Commit is the revision that was built, recorded on the deployment. | [optional] 
**Files** | Pointer to **int32** | Files is how many objects CI published. | [optional] 
**Id** | Pointer to **string** | ID is the queued deployment to complete, from the path. | [optional] 
**Keys** | Pointer to **[]string** | Keys is the manifest CI just uploaded, RELATIVE to the deployment prefix. It is what replaces &#x60;aws s3 sync --delete&#x60;: an upload grant authorizes writes only, so CI cannot remove a file, and cloud reconciles the prefix against this list instead (grant.go). Omit it and nothing is deleted — the prefix only grows, which is the old pre-grant behaviour and a safe default. | [optional] 
**LiveUrl** | Pointer to **string** | LiveURL is a HINT at the address the site should serve at. The public host is claimed by cloud first, so this can refine the URL a deployment reports but can never assert a subdomain another tenant holds. | [optional] 
**Message** | Pointer to **string** | Message is what happened, in words — on an error completion, why it failed. | [optional] 
**Slug** | Pointer to **string** | Slug is the project the deployment belongs to, from the path. | [optional] 
**Status** | Pointer to **string** | Status is how the build ended: &#x60;live&#x60; if it succeeded, &#x60;error&#x60; if it did not. Nothing else is accepted. | [optional] 

## Methods

### NewProjectsComplete

`func NewProjectsComplete() *ProjectsComplete`

NewProjectsComplete instantiates a new ProjectsComplete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsCompleteWithDefaults

`func NewProjectsCompleteWithDefaults() *ProjectsComplete`

NewProjectsCompleteWithDefaults instantiates a new ProjectsComplete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *ProjectsComplete) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ProjectsComplete) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ProjectsComplete) SetBytes(v int32)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *ProjectsComplete) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetCommit

`func (o *ProjectsComplete) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ProjectsComplete) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ProjectsComplete) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *ProjectsComplete) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetFiles

`func (o *ProjectsComplete) GetFiles() int32`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ProjectsComplete) GetFilesOk() (*int32, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ProjectsComplete) SetFiles(v int32)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ProjectsComplete) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetId

`func (o *ProjectsComplete) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectsComplete) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectsComplete) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectsComplete) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeys

`func (o *ProjectsComplete) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ProjectsComplete) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ProjectsComplete) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ProjectsComplete) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetLiveUrl

`func (o *ProjectsComplete) GetLiveUrl() string`

GetLiveUrl returns the LiveUrl field if non-nil, zero value otherwise.

### GetLiveUrlOk

`func (o *ProjectsComplete) GetLiveUrlOk() (*string, bool)`

GetLiveUrlOk returns a tuple with the LiveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveUrl

`func (o *ProjectsComplete) SetLiveUrl(v string)`

SetLiveUrl sets LiveUrl field to given value.

### HasLiveUrl

`func (o *ProjectsComplete) HasLiveUrl() bool`

HasLiveUrl returns a boolean if a field has been set.

### GetMessage

`func (o *ProjectsComplete) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProjectsComplete) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProjectsComplete) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProjectsComplete) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectsComplete) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsComplete) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsComplete) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectsComplete) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectsComplete) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectsComplete) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectsComplete) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectsComplete) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



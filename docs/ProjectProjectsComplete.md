# ProjectProjectsComplete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **int64** | Bytes is their total size in bytes. | [optional] 
**Commit** | Pointer to **string** | Commit is the revision that was built, recorded on the deployment. | [optional] 
**Copy** | Pointer to **[]string** | Copy lists the keys to take from CopyFrom&#39;s prefix, relative to it. They are also named in Keys, which stays the whole release. | [optional] 
**CopyFrom** | Pointer to **string** | CopyFrom names another site of the same org whose prefix already holds the bytes of Copy. A host that shares a build with another differs from it by a page or two, so CI uploads those and asks for the rest to be copied inside the object store instead of sent again over the runner&#39;s link. | [optional] 
**Files** | Pointer to **int64** | Files is how many objects CI published. | [optional] 
**Id** | Pointer to **string** | ID is the queued deployment to complete, from the path. | [optional] 
**Keys** | Pointer to **[]string** | Keys is the manifest CI just uploaded, RELATIVE to the deployment prefix. It is what replaces &#x60;aws s3 sync --delete&#x60;: an upload grant authorizes writes only, so CI cannot remove a file, and cloud reconciles the prefix against this list instead (grant.go). Omit it and nothing is deleted — the prefix only grows, which is the old pre-grant behaviour and a safe default. | [optional] 
**LiveUrl** | Pointer to **string** | LiveURL is a HINT at the address the site should serve at. The public host is claimed by cloud first, so this can refine the URL a deployment reports but can never assert a subdomain another tenant holds. | [optional] 
**Message** | Pointer to **string** | Message is what happened, in words — on an error completion, why it failed. | [optional] 
**Slug** | Pointer to **string** | Slug is the project the deployment belongs to, from the path. | [optional] 
**Status** | Pointer to **string** | Status is how the build ended: &#x60;live&#x60; if it succeeded, &#x60;error&#x60; if it did not. Nothing else is accepted. | [optional] 

## Methods

### NewProjectProjectsComplete

`func NewProjectProjectsComplete() *ProjectProjectsComplete`

NewProjectProjectsComplete instantiates a new ProjectProjectsComplete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectsCompleteWithDefaults

`func NewProjectProjectsCompleteWithDefaults() *ProjectProjectsComplete`

NewProjectProjectsCompleteWithDefaults instantiates a new ProjectProjectsComplete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *ProjectProjectsComplete) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ProjectProjectsComplete) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ProjectProjectsComplete) SetBytes(v int64)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *ProjectProjectsComplete) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetCommit

`func (o *ProjectProjectsComplete) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ProjectProjectsComplete) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ProjectProjectsComplete) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *ProjectProjectsComplete) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetCopy

`func (o *ProjectProjectsComplete) GetCopy() []string`

GetCopy returns the Copy field if non-nil, zero value otherwise.

### GetCopyOk

`func (o *ProjectProjectsComplete) GetCopyOk() (*[]string, bool)`

GetCopyOk returns a tuple with the Copy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopy

`func (o *ProjectProjectsComplete) SetCopy(v []string)`

SetCopy sets Copy field to given value.

### HasCopy

`func (o *ProjectProjectsComplete) HasCopy() bool`

HasCopy returns a boolean if a field has been set.

### GetCopyFrom

`func (o *ProjectProjectsComplete) GetCopyFrom() string`

GetCopyFrom returns the CopyFrom field if non-nil, zero value otherwise.

### GetCopyFromOk

`func (o *ProjectProjectsComplete) GetCopyFromOk() (*string, bool)`

GetCopyFromOk returns a tuple with the CopyFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyFrom

`func (o *ProjectProjectsComplete) SetCopyFrom(v string)`

SetCopyFrom sets CopyFrom field to given value.

### HasCopyFrom

`func (o *ProjectProjectsComplete) HasCopyFrom() bool`

HasCopyFrom returns a boolean if a field has been set.

### GetFiles

`func (o *ProjectProjectsComplete) GetFiles() int64`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ProjectProjectsComplete) GetFilesOk() (*int64, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ProjectProjectsComplete) SetFiles(v int64)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ProjectProjectsComplete) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetId

`func (o *ProjectProjectsComplete) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectProjectsComplete) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectProjectsComplete) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectProjectsComplete) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeys

`func (o *ProjectProjectsComplete) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ProjectProjectsComplete) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ProjectProjectsComplete) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ProjectProjectsComplete) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetLiveUrl

`func (o *ProjectProjectsComplete) GetLiveUrl() string`

GetLiveUrl returns the LiveUrl field if non-nil, zero value otherwise.

### GetLiveUrlOk

`func (o *ProjectProjectsComplete) GetLiveUrlOk() (*string, bool)`

GetLiveUrlOk returns a tuple with the LiveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveUrl

`func (o *ProjectProjectsComplete) SetLiveUrl(v string)`

SetLiveUrl sets LiveUrl field to given value.

### HasLiveUrl

`func (o *ProjectProjectsComplete) HasLiveUrl() bool`

HasLiveUrl returns a boolean if a field has been set.

### GetMessage

`func (o *ProjectProjectsComplete) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProjectProjectsComplete) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProjectProjectsComplete) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProjectProjectsComplete) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectProjectsComplete) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectProjectsComplete) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectProjectsComplete) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectProjectsComplete) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectProjectsComplete) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectProjectsComplete) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectProjectsComplete) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectProjectsComplete) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



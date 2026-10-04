# AutoFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **int64** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**ProjectId** | Pointer to **string** | projectId &#x3D;&#x3D; org (server-derived) | [optional] 
**PublishedVersionId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **int64** |  | [optional] 

## Methods

### NewAutoFlow

`func NewAutoFlow() *AutoFlow`

NewAutoFlow instantiates a new AutoFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoFlowWithDefaults

`func NewAutoFlowWithDefaults() *AutoFlow`

NewAutoFlowWithDefaults instantiates a new AutoFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AutoFlow) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AutoFlow) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AutoFlow) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AutoFlow) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExternalId

`func (o *AutoFlow) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AutoFlow) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AutoFlow) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AutoFlow) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFolderId

`func (o *AutoFlow) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AutoFlow) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AutoFlow) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *AutoFlow) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetId

`func (o *AutoFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoFlow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoFlow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *AutoFlow) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AutoFlow) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AutoFlow) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AutoFlow) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *AutoFlow) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *AutoFlow) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetProjectId

`func (o *AutoFlow) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AutoFlow) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AutoFlow) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AutoFlow) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetPublishedVersionId

`func (o *AutoFlow) GetPublishedVersionId() string`

GetPublishedVersionId returns the PublishedVersionId field if non-nil, zero value otherwise.

### GetPublishedVersionIdOk

`func (o *AutoFlow) GetPublishedVersionIdOk() (*string, bool)`

GetPublishedVersionIdOk returns a tuple with the PublishedVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersionId

`func (o *AutoFlow) SetPublishedVersionId(v string)`

SetPublishedVersionId sets PublishedVersionId field to given value.

### HasPublishedVersionId

`func (o *AutoFlow) HasPublishedVersionId() bool`

HasPublishedVersionId returns a boolean if a field has been set.

### GetStatus

`func (o *AutoFlow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutoFlow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutoFlow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutoFlow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdated

`func (o *AutoFlow) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AutoFlow) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AutoFlow) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AutoFlow) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



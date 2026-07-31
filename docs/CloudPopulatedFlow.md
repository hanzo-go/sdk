# CloudPopulatedFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **int32** | Created and Updated are unix milliseconds. | [optional] 
**ExternalId** | Pointer to **string** | ExternalID is the caller&#39;s own id for this flow, if it set one. | [optional] 
**FolderId** | Pointer to **string** | FolderID groups the flow in the builder&#39;s tree. | [optional] 
**Id** | Pointer to **string** | ID is the flow&#39;s id. | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**ProjectId** | Pointer to **string** | Org is the owning org, which this surface names projectId. Server-derived from the validated principal — never read from a request. | [optional] 
**PublishedVersionId** | Pointer to **string** | PublishedVersionID is the version a run executes when set; empty means the latest version runs. | [optional] 
**Status** | Pointer to **string** | Status is ENABLED or DISABLED — whether the flow&#39;s trigger is armed. | [optional] 
**Updated** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to [**CloudFlowVersion**](CloudFlowVersion.md) | Version is the flow&#39;s latest version — its display name and step tree. | [optional] 

## Methods

### NewCloudPopulatedFlow

`func NewCloudPopulatedFlow() *CloudPopulatedFlow`

NewCloudPopulatedFlow instantiates a new CloudPopulatedFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPopulatedFlowWithDefaults

`func NewCloudPopulatedFlowWithDefaults() *CloudPopulatedFlow`

NewCloudPopulatedFlowWithDefaults instantiates a new CloudPopulatedFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CloudPopulatedFlow) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CloudPopulatedFlow) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CloudPopulatedFlow) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CloudPopulatedFlow) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExternalId

`func (o *CloudPopulatedFlow) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CloudPopulatedFlow) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CloudPopulatedFlow) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CloudPopulatedFlow) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFolderId

`func (o *CloudPopulatedFlow) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *CloudPopulatedFlow) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *CloudPopulatedFlow) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *CloudPopulatedFlow) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetId

`func (o *CloudPopulatedFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudPopulatedFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudPopulatedFlow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudPopulatedFlow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CloudPopulatedFlow) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CloudPopulatedFlow) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CloudPopulatedFlow) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CloudPopulatedFlow) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CloudPopulatedFlow) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CloudPopulatedFlow) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetProjectId

`func (o *CloudPopulatedFlow) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CloudPopulatedFlow) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CloudPopulatedFlow) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CloudPopulatedFlow) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetPublishedVersionId

`func (o *CloudPopulatedFlow) GetPublishedVersionId() string`

GetPublishedVersionId returns the PublishedVersionId field if non-nil, zero value otherwise.

### GetPublishedVersionIdOk

`func (o *CloudPopulatedFlow) GetPublishedVersionIdOk() (*string, bool)`

GetPublishedVersionIdOk returns a tuple with the PublishedVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersionId

`func (o *CloudPopulatedFlow) SetPublishedVersionId(v string)`

SetPublishedVersionId sets PublishedVersionId field to given value.

### HasPublishedVersionId

`func (o *CloudPopulatedFlow) HasPublishedVersionId() bool`

HasPublishedVersionId returns a boolean if a field has been set.

### GetStatus

`func (o *CloudPopulatedFlow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudPopulatedFlow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudPopulatedFlow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudPopulatedFlow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdated

`func (o *CloudPopulatedFlow) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CloudPopulatedFlow) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CloudPopulatedFlow) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CloudPopulatedFlow) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetVersion

`func (o *CloudPopulatedFlow) GetVersion() CloudFlowVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudPopulatedFlow) GetVersionOk() (*CloudFlowVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudPopulatedFlow) SetVersion(v CloudFlowVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudPopulatedFlow) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



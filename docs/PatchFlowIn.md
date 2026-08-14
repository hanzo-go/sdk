# PatchFlowIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | ExternalID sets the caller&#39;s own id for this flow. | [optional] 
**FolderId** | Pointer to **string** | FolderID moves the flow in the builder&#39;s tree. | [optional] 
**Id** | Pointer to **string** | ID is the flow to update, from the path. | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**PublishedVersionId** | Pointer to **string** | PublishedVersionID pins the version runs execute. It must name a version OF THIS FLOW; empty clears the pin, so runs take the latest version again. | [optional] 

## Methods

### NewPatchFlowIn

`func NewPatchFlowIn() *PatchFlowIn`

NewPatchFlowIn instantiates a new PatchFlowIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchFlowInWithDefaults

`func NewPatchFlowInWithDefaults() *PatchFlowIn`

NewPatchFlowInWithDefaults instantiates a new PatchFlowIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *PatchFlowIn) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PatchFlowIn) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PatchFlowIn) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *PatchFlowIn) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFolderId

`func (o *PatchFlowIn) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *PatchFlowIn) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *PatchFlowIn) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *PatchFlowIn) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetId

`func (o *PatchFlowIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchFlowIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchFlowIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchFlowIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *PatchFlowIn) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchFlowIn) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchFlowIn) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchFlowIn) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PatchFlowIn) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PatchFlowIn) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPublishedVersionId

`func (o *PatchFlowIn) GetPublishedVersionId() string`

GetPublishedVersionId returns the PublishedVersionId field if non-nil, zero value otherwise.

### GetPublishedVersionIdOk

`func (o *PatchFlowIn) GetPublishedVersionIdOk() (*string, bool)`

GetPublishedVersionIdOk returns a tuple with the PublishedVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersionId

`func (o *PatchFlowIn) SetPublishedVersionId(v string)`

SetPublishedVersionId sets PublishedVersionId field to given value.

### HasPublishedVersionId

`func (o *PatchFlowIn) HasPublishedVersionId() bool`

HasPublishedVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



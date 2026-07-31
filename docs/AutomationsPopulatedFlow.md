# AutomationsPopulatedFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** | The owner org (server-derived) | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**PublishedVersionId** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**Created** | Pointer to **int64** | Unix ms | [optional] 
**Updated** | Pointer to **int64** | Unix ms | [optional] 
**Version** | Pointer to [**AutomationsFlowVersion**](AutomationsFlowVersion.md) |  | [optional] 

## Methods

### NewAutomationsPopulatedFlow

`func NewAutomationsPopulatedFlow() *AutomationsPopulatedFlow`

NewAutomationsPopulatedFlow instantiates a new AutomationsPopulatedFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationsPopulatedFlowWithDefaults

`func NewAutomationsPopulatedFlowWithDefaults() *AutomationsPopulatedFlow`

NewAutomationsPopulatedFlowWithDefaults instantiates a new AutomationsPopulatedFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutomationsPopulatedFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationsPopulatedFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationsPopulatedFlow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutomationsPopulatedFlow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *AutomationsPopulatedFlow) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AutomationsPopulatedFlow) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AutomationsPopulatedFlow) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AutomationsPopulatedFlow) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetExternalId

`func (o *AutomationsPopulatedFlow) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AutomationsPopulatedFlow) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AutomationsPopulatedFlow) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AutomationsPopulatedFlow) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFolderId

`func (o *AutomationsPopulatedFlow) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AutomationsPopulatedFlow) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AutomationsPopulatedFlow) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *AutomationsPopulatedFlow) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetStatus

`func (o *AutomationsPopulatedFlow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutomationsPopulatedFlow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutomationsPopulatedFlow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutomationsPopulatedFlow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPublishedVersionId

`func (o *AutomationsPopulatedFlow) GetPublishedVersionId() string`

GetPublishedVersionId returns the PublishedVersionId field if non-nil, zero value otherwise.

### GetPublishedVersionIdOk

`func (o *AutomationsPopulatedFlow) GetPublishedVersionIdOk() (*string, bool)`

GetPublishedVersionIdOk returns a tuple with the PublishedVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersionId

`func (o *AutomationsPopulatedFlow) SetPublishedVersionId(v string)`

SetPublishedVersionId sets PublishedVersionId field to given value.

### HasPublishedVersionId

`func (o *AutomationsPopulatedFlow) HasPublishedVersionId() bool`

HasPublishedVersionId returns a boolean if a field has been set.

### GetMetadata

`func (o *AutomationsPopulatedFlow) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AutomationsPopulatedFlow) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AutomationsPopulatedFlow) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AutomationsPopulatedFlow) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *AutomationsPopulatedFlow) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *AutomationsPopulatedFlow) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCreated

`func (o *AutomationsPopulatedFlow) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AutomationsPopulatedFlow) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AutomationsPopulatedFlow) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AutomationsPopulatedFlow) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *AutomationsPopulatedFlow) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AutomationsPopulatedFlow) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AutomationsPopulatedFlow) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AutomationsPopulatedFlow) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetVersion

`func (o *AutomationsPopulatedFlow) GetVersion() AutomationsFlowVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AutomationsPopulatedFlow) GetVersionOk() (*AutomationsFlowVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AutomationsPopulatedFlow) SetVersion(v AutomationsFlowVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AutomationsPopulatedFlow) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



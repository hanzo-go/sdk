# CreateFlowReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | DisplayName names the flow&#39;s initial draft version. | [optional] 
**ExternalId** | Pointer to **string** | ExternalID is the caller&#39;s own id for this flow. Optional. | [optional] 
**FolderId** | Pointer to **string** | FolderID groups the flow in the builder&#39;s tree. Optional. | [optional] 
**Trigger** | Pointer to [**FlowTrigger**](FlowTrigger.md) | Trigger is the root of the step tree — how the flow starts, and the action chain that follows. Optional: a flow may be created empty and edited later. | [optional] 

## Methods

### NewCreateFlowReq

`func NewCreateFlowReq() *CreateFlowReq`

NewCreateFlowReq instantiates a new CreateFlowReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFlowReqWithDefaults

`func NewCreateFlowReqWithDefaults() *CreateFlowReq`

NewCreateFlowReqWithDefaults instantiates a new CreateFlowReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CreateFlowReq) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateFlowReq) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateFlowReq) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateFlowReq) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateFlowReq) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateFlowReq) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateFlowReq) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateFlowReq) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFolderId

`func (o *CreateFlowReq) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *CreateFlowReq) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *CreateFlowReq) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *CreateFlowReq) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetTrigger

`func (o *CreateFlowReq) GetTrigger() FlowTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *CreateFlowReq) GetTriggerOk() (*FlowTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *CreateFlowReq) SetTrigger(v FlowTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *CreateFlowReq) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



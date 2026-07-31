# CloudCreateFlowReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | DisplayName names the flow&#39;s initial draft version. | [optional] 
**ExternalId** | Pointer to **string** | ExternalID is the caller&#39;s own id for this flow. Optional. | [optional] 
**FolderId** | Pointer to **string** | FolderID groups the flow in the builder&#39;s tree. Optional. | [optional] 
**Trigger** | Pointer to [**CloudFlowTrigger**](CloudFlowTrigger.md) | Trigger is the root of the step tree — how the flow starts, and the action chain that follows. Optional: a flow may be created empty and edited later. | [optional] 

## Methods

### NewCloudCreateFlowReq

`func NewCloudCreateFlowReq() *CloudCreateFlowReq`

NewCloudCreateFlowReq instantiates a new CloudCreateFlowReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCreateFlowReqWithDefaults

`func NewCloudCreateFlowReqWithDefaults() *CloudCreateFlowReq`

NewCloudCreateFlowReqWithDefaults instantiates a new CloudCreateFlowReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CloudCreateFlowReq) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudCreateFlowReq) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudCreateFlowReq) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudCreateFlowReq) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalId

`func (o *CloudCreateFlowReq) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CloudCreateFlowReq) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CloudCreateFlowReq) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CloudCreateFlowReq) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFolderId

`func (o *CloudCreateFlowReq) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *CloudCreateFlowReq) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *CloudCreateFlowReq) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *CloudCreateFlowReq) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetTrigger

`func (o *CloudCreateFlowReq) GetTrigger() CloudFlowTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *CloudCreateFlowReq) GetTriggerOk() (*CloudFlowTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *CloudCreateFlowReq) SetTrigger(v CloudFlowTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *CloudCreateFlowReq) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



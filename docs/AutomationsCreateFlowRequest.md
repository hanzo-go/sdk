# AutomationsCreateFlowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**Trigger** | Pointer to [**AutomationsFlowTrigger**](AutomationsFlowTrigger.md) |  | [optional] 

## Methods

### NewAutomationsCreateFlowRequest

`func NewAutomationsCreateFlowRequest() *AutomationsCreateFlowRequest`

NewAutomationsCreateFlowRequest instantiates a new AutomationsCreateFlowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationsCreateFlowRequestWithDefaults

`func NewAutomationsCreateFlowRequestWithDefaults() *AutomationsCreateFlowRequest`

NewAutomationsCreateFlowRequestWithDefaults instantiates a new AutomationsCreateFlowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *AutomationsCreateFlowRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AutomationsCreateFlowRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AutomationsCreateFlowRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AutomationsCreateFlowRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalId

`func (o *AutomationsCreateFlowRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AutomationsCreateFlowRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AutomationsCreateFlowRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AutomationsCreateFlowRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFolderId

`func (o *AutomationsCreateFlowRequest) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AutomationsCreateFlowRequest) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AutomationsCreateFlowRequest) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *AutomationsCreateFlowRequest) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetTrigger

`func (o *AutomationsCreateFlowRequest) GetTrigger() AutomationsFlowTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *AutomationsCreateFlowRequest) GetTriggerOk() (*AutomationsFlowTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *AutomationsCreateFlowRequest) SetTrigger(v AutomationsFlowTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *AutomationsCreateFlowRequest) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



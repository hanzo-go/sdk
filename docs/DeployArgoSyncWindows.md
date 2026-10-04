# DeployArgoSyncWindows

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveWindows** | Pointer to **[]interface{}** | ActiveWindows are the sync windows in force right now. Always null: this platform declares none, so nothing is ever in force. | [optional] 
**AssignedWindows** | Pointer to **[]interface{}** | AssignedWindows are the windows configured for this application at all, whether or not currently in force. Always null, for the same reason. | [optional] 
**CanSync** | Pointer to **bool** | CanSync is whether a sync would be permitted at this moment. Always true — with no windows there is nothing to deny it. A caller must not read this as \&quot;a sync will succeed\&quot;; it only means no window is blocking one. | [optional] 

## Methods

### NewDeployArgoSyncWindows

`func NewDeployArgoSyncWindows() *DeployArgoSyncWindows`

NewDeployArgoSyncWindows instantiates a new DeployArgoSyncWindows object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployArgoSyncWindowsWithDefaults

`func NewDeployArgoSyncWindowsWithDefaults() *DeployArgoSyncWindows`

NewDeployArgoSyncWindowsWithDefaults instantiates a new DeployArgoSyncWindows object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveWindows

`func (o *DeployArgoSyncWindows) GetActiveWindows() []interface{}`

GetActiveWindows returns the ActiveWindows field if non-nil, zero value otherwise.

### GetActiveWindowsOk

`func (o *DeployArgoSyncWindows) GetActiveWindowsOk() (*[]interface{}, bool)`

GetActiveWindowsOk returns a tuple with the ActiveWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveWindows

`func (o *DeployArgoSyncWindows) SetActiveWindows(v []interface{})`

SetActiveWindows sets ActiveWindows field to given value.

### HasActiveWindows

`func (o *DeployArgoSyncWindows) HasActiveWindows() bool`

HasActiveWindows returns a boolean if a field has been set.

### GetAssignedWindows

`func (o *DeployArgoSyncWindows) GetAssignedWindows() []interface{}`

GetAssignedWindows returns the AssignedWindows field if non-nil, zero value otherwise.

### GetAssignedWindowsOk

`func (o *DeployArgoSyncWindows) GetAssignedWindowsOk() (*[]interface{}, bool)`

GetAssignedWindowsOk returns a tuple with the AssignedWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedWindows

`func (o *DeployArgoSyncWindows) SetAssignedWindows(v []interface{})`

SetAssignedWindows sets AssignedWindows field to given value.

### HasAssignedWindows

`func (o *DeployArgoSyncWindows) HasAssignedWindows() bool`

HasAssignedWindows returns a boolean if a field has been set.

### GetCanSync

`func (o *DeployArgoSyncWindows) GetCanSync() bool`

GetCanSync returns the CanSync field if non-nil, zero value otherwise.

### GetCanSyncOk

`func (o *DeployArgoSyncWindows) GetCanSyncOk() (*bool, bool)`

GetCanSyncOk returns a tuple with the CanSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSync

`func (o *DeployArgoSyncWindows) SetCanSync(v bool)`

SetCanSync sets CanSync field to given value.

### HasCanSync

`func (o *DeployArgoSyncWindows) HasCanSync() bool`

HasCanSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



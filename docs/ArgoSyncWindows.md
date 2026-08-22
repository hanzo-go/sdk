# ArgoSyncWindows

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveWindows** | Pointer to **[]map[string]interface{}** | ActiveWindows are the sync windows in force right now. Always null: this platform declares none, so nothing is ever in force. | [optional] 
**AssignedWindows** | Pointer to **[]map[string]interface{}** | AssignedWindows are the windows configured for this application at all, whether or not currently in force. Always null, for the same reason. | [optional] 
**CanSync** | Pointer to **bool** | CanSync is whether a sync would be permitted at this moment. Always true — with no windows there is nothing to deny it. A caller must not read this as \&quot;a sync will succeed\&quot;; it only means no window is blocking one. | [optional] 

## Methods

### NewArgoSyncWindows

`func NewArgoSyncWindows() *ArgoSyncWindows`

NewArgoSyncWindows instantiates a new ArgoSyncWindows object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoSyncWindowsWithDefaults

`func NewArgoSyncWindowsWithDefaults() *ArgoSyncWindows`

NewArgoSyncWindowsWithDefaults instantiates a new ArgoSyncWindows object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveWindows

`func (o *ArgoSyncWindows) GetActiveWindows() []map[string]interface{}`

GetActiveWindows returns the ActiveWindows field if non-nil, zero value otherwise.

### GetActiveWindowsOk

`func (o *ArgoSyncWindows) GetActiveWindowsOk() (*[]map[string]interface{}, bool)`

GetActiveWindowsOk returns a tuple with the ActiveWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveWindows

`func (o *ArgoSyncWindows) SetActiveWindows(v []map[string]interface{})`

SetActiveWindows sets ActiveWindows field to given value.

### HasActiveWindows

`func (o *ArgoSyncWindows) HasActiveWindows() bool`

HasActiveWindows returns a boolean if a field has been set.

### GetAssignedWindows

`func (o *ArgoSyncWindows) GetAssignedWindows() []map[string]interface{}`

GetAssignedWindows returns the AssignedWindows field if non-nil, zero value otherwise.

### GetAssignedWindowsOk

`func (o *ArgoSyncWindows) GetAssignedWindowsOk() (*[]map[string]interface{}, bool)`

GetAssignedWindowsOk returns a tuple with the AssignedWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedWindows

`func (o *ArgoSyncWindows) SetAssignedWindows(v []map[string]interface{})`

SetAssignedWindows sets AssignedWindows field to given value.

### HasAssignedWindows

`func (o *ArgoSyncWindows) HasAssignedWindows() bool`

HasAssignedWindows returns a boolean if a field has been set.

### GetCanSync

`func (o *ArgoSyncWindows) GetCanSync() bool`

GetCanSync returns the CanSync field if non-nil, zero value otherwise.

### GetCanSyncOk

`func (o *ArgoSyncWindows) GetCanSyncOk() (*bool, bool)`

GetCanSyncOk returns a tuple with the CanSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSync

`func (o *ArgoSyncWindows) SetCanSync(v bool)`

SetCanSync sets CanSync field to given value.

### HasCanSync

`func (o *ArgoSyncWindows) HasCanSync() bool`

HasCanSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CloudArgoSyncWindows

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveWindows** | Pointer to **[]map[string]interface{}** |  | [optional] 
**AssignedWindows** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CanSync** | Pointer to **bool** |  | [optional] 

## Methods

### NewCloudArgoSyncWindows

`func NewCloudArgoSyncWindows() *CloudArgoSyncWindows`

NewCloudArgoSyncWindows instantiates a new CloudArgoSyncWindows object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudArgoSyncWindowsWithDefaults

`func NewCloudArgoSyncWindowsWithDefaults() *CloudArgoSyncWindows`

NewCloudArgoSyncWindowsWithDefaults instantiates a new CloudArgoSyncWindows object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveWindows

`func (o *CloudArgoSyncWindows) GetActiveWindows() []map[string]interface{}`

GetActiveWindows returns the ActiveWindows field if non-nil, zero value otherwise.

### GetActiveWindowsOk

`func (o *CloudArgoSyncWindows) GetActiveWindowsOk() (*[]map[string]interface{}, bool)`

GetActiveWindowsOk returns a tuple with the ActiveWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveWindows

`func (o *CloudArgoSyncWindows) SetActiveWindows(v []map[string]interface{})`

SetActiveWindows sets ActiveWindows field to given value.

### HasActiveWindows

`func (o *CloudArgoSyncWindows) HasActiveWindows() bool`

HasActiveWindows returns a boolean if a field has been set.

### GetAssignedWindows

`func (o *CloudArgoSyncWindows) GetAssignedWindows() []map[string]interface{}`

GetAssignedWindows returns the AssignedWindows field if non-nil, zero value otherwise.

### GetAssignedWindowsOk

`func (o *CloudArgoSyncWindows) GetAssignedWindowsOk() (*[]map[string]interface{}, bool)`

GetAssignedWindowsOk returns a tuple with the AssignedWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedWindows

`func (o *CloudArgoSyncWindows) SetAssignedWindows(v []map[string]interface{})`

SetAssignedWindows sets AssignedWindows field to given value.

### HasAssignedWindows

`func (o *CloudArgoSyncWindows) HasAssignedWindows() bool`

HasAssignedWindows returns a boolean if a field has been set.

### GetCanSync

`func (o *CloudArgoSyncWindows) GetCanSync() bool`

GetCanSync returns the CanSync field if non-nil, zero value otherwise.

### GetCanSyncOk

`func (o *CloudArgoSyncWindows) GetCanSyncOk() (*bool, bool)`

GetCanSyncOk returns a tuple with the CanSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSync

`func (o *CloudArgoSyncWindows) SetCanSync(v bool)`

SetCanSync sets CanSync field to given value.

### HasCanSync

`func (o *CloudArgoSyncWindows) HasCanSync() bool`

HasCanSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



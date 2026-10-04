# PlatformProjectBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CdUnavailable** | Pointer to [**PlatformUnreadable**](PlatformUnreadable.md) | CD says why the reconciliation columns are missing, when they are. | [optional] 
**ClusterUnavailable** | Pointer to [**PlatformUnreadable**](PlatformUnreadable.md) | Cluster says why the running columns are missing, when they are. | [optional] 
**Projects** | Pointer to [**[]PlatformProject**](PlatformProject.md) | Projects are sorted by owner, then name. | [optional] 
**Unassigned** | Pointer to [**[]PlatformDeclRef**](PlatformDeclRef.md) | Unassigned are declarations that name no project. Every declaration names one, so a non-empty list is drift in git itself. | [optional] 
**Unreadable** | Pointer to [**[]PlatformUnreadableFile**](PlatformUnreadableFile.md) | Unreadable are the values files in the caller&#39;s scope that do not parse, so their projects are unknown. Absent when every file parsed. | [optional] 

## Methods

### NewPlatformProjectBoard

`func NewPlatformProjectBoard() *PlatformProjectBoard`

NewPlatformProjectBoard instantiates a new PlatformProjectBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformProjectBoardWithDefaults

`func NewPlatformProjectBoardWithDefaults() *PlatformProjectBoard`

NewPlatformProjectBoardWithDefaults instantiates a new PlatformProjectBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdUnavailable

`func (o *PlatformProjectBoard) GetCdUnavailable() PlatformUnreadable`

GetCdUnavailable returns the CdUnavailable field if non-nil, zero value otherwise.

### GetCdUnavailableOk

`func (o *PlatformProjectBoard) GetCdUnavailableOk() (*PlatformUnreadable, bool)`

GetCdUnavailableOk returns a tuple with the CdUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdUnavailable

`func (o *PlatformProjectBoard) SetCdUnavailable(v PlatformUnreadable)`

SetCdUnavailable sets CdUnavailable field to given value.

### HasCdUnavailable

`func (o *PlatformProjectBoard) HasCdUnavailable() bool`

HasCdUnavailable returns a boolean if a field has been set.

### GetClusterUnavailable

`func (o *PlatformProjectBoard) GetClusterUnavailable() PlatformUnreadable`

GetClusterUnavailable returns the ClusterUnavailable field if non-nil, zero value otherwise.

### GetClusterUnavailableOk

`func (o *PlatformProjectBoard) GetClusterUnavailableOk() (*PlatformUnreadable, bool)`

GetClusterUnavailableOk returns a tuple with the ClusterUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUnavailable

`func (o *PlatformProjectBoard) SetClusterUnavailable(v PlatformUnreadable)`

SetClusterUnavailable sets ClusterUnavailable field to given value.

### HasClusterUnavailable

`func (o *PlatformProjectBoard) HasClusterUnavailable() bool`

HasClusterUnavailable returns a boolean if a field has been set.

### GetProjects

`func (o *PlatformProjectBoard) GetProjects() []PlatformProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *PlatformProjectBoard) GetProjectsOk() (*[]PlatformProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *PlatformProjectBoard) SetProjects(v []PlatformProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *PlatformProjectBoard) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetUnassigned

`func (o *PlatformProjectBoard) GetUnassigned() []PlatformDeclRef`

GetUnassigned returns the Unassigned field if non-nil, zero value otherwise.

### GetUnassignedOk

`func (o *PlatformProjectBoard) GetUnassignedOk() (*[]PlatformDeclRef, bool)`

GetUnassignedOk returns a tuple with the Unassigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassigned

`func (o *PlatformProjectBoard) SetUnassigned(v []PlatformDeclRef)`

SetUnassigned sets Unassigned field to given value.

### HasUnassigned

`func (o *PlatformProjectBoard) HasUnassigned() bool`

HasUnassigned returns a boolean if a field has been set.

### GetUnreadable

`func (o *PlatformProjectBoard) GetUnreadable() []PlatformUnreadableFile`

GetUnreadable returns the Unreadable field if non-nil, zero value otherwise.

### GetUnreadableOk

`func (o *PlatformProjectBoard) GetUnreadableOk() (*[]PlatformUnreadableFile, bool)`

GetUnreadableOk returns a tuple with the Unreadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadable

`func (o *PlatformProjectBoard) SetUnreadable(v []PlatformUnreadableFile)`

SetUnreadable sets Unreadable field to given value.

### HasUnreadable

`func (o *PlatformProjectBoard) HasUnreadable() bool`

HasUnreadable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



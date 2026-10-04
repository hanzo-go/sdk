# PlatformProjectView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to **int64** |  | [optional] 
**CdUnavailable** | Pointer to [**PlatformUnreadable**](PlatformUnreadable.md) | CD says why the reconciliation columns are missing, when they are. | [optional] 
**ClusterUnavailable** | Pointer to [**PlatformUnreadable**](PlatformUnreadable.md) | Cluster says why the running columns are missing, when they are. | [optional] 
**Drift** | Pointer to [**PlatformDriftTally**](PlatformDriftTally.md) |  | [optional] 
**Health** | Pointer to [**PlatformHealthTally**](PlatformHealthTally.md) |  | [optional] 
**Members** | Pointer to [**[]PlatformProjectApp**](PlatformProjectApp.md) | Members are the project&#39;s apps, sorted by namespace then name. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespaces** | Pointer to **[]string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Sync** | Pointer to [**PlatformSyncTally**](PlatformSyncTally.md) |  | [optional] 
**Unreadable** | Pointer to [**[]PlatformUnreadableFile**](PlatformUnreadableFile.md) | Unreadable are the values files in the owner&#39;s scope that do not parse, any of which may name this project. Absent when every file parsed. | [optional] 

## Methods

### NewPlatformProjectView

`func NewPlatformProjectView() *PlatformProjectView`

NewPlatformProjectView instantiates a new PlatformProjectView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformProjectViewWithDefaults

`func NewPlatformProjectViewWithDefaults() *PlatformProjectView`

NewPlatformProjectViewWithDefaults instantiates a new PlatformProjectView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *PlatformProjectView) GetApps() int64`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *PlatformProjectView) GetAppsOk() (*int64, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *PlatformProjectView) SetApps(v int64)`

SetApps sets Apps field to given value.

### HasApps

`func (o *PlatformProjectView) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetCdUnavailable

`func (o *PlatformProjectView) GetCdUnavailable() PlatformUnreadable`

GetCdUnavailable returns the CdUnavailable field if non-nil, zero value otherwise.

### GetCdUnavailableOk

`func (o *PlatformProjectView) GetCdUnavailableOk() (*PlatformUnreadable, bool)`

GetCdUnavailableOk returns a tuple with the CdUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdUnavailable

`func (o *PlatformProjectView) SetCdUnavailable(v PlatformUnreadable)`

SetCdUnavailable sets CdUnavailable field to given value.

### HasCdUnavailable

`func (o *PlatformProjectView) HasCdUnavailable() bool`

HasCdUnavailable returns a boolean if a field has been set.

### GetClusterUnavailable

`func (o *PlatformProjectView) GetClusterUnavailable() PlatformUnreadable`

GetClusterUnavailable returns the ClusterUnavailable field if non-nil, zero value otherwise.

### GetClusterUnavailableOk

`func (o *PlatformProjectView) GetClusterUnavailableOk() (*PlatformUnreadable, bool)`

GetClusterUnavailableOk returns a tuple with the ClusterUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUnavailable

`func (o *PlatformProjectView) SetClusterUnavailable(v PlatformUnreadable)`

SetClusterUnavailable sets ClusterUnavailable field to given value.

### HasClusterUnavailable

`func (o *PlatformProjectView) HasClusterUnavailable() bool`

HasClusterUnavailable returns a boolean if a field has been set.

### GetDrift

`func (o *PlatformProjectView) GetDrift() PlatformDriftTally`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *PlatformProjectView) GetDriftOk() (*PlatformDriftTally, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *PlatformProjectView) SetDrift(v PlatformDriftTally)`

SetDrift sets Drift field to given value.

### HasDrift

`func (o *PlatformProjectView) HasDrift() bool`

HasDrift returns a boolean if a field has been set.

### GetHealth

`func (o *PlatformProjectView) GetHealth() PlatformHealthTally`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *PlatformProjectView) GetHealthOk() (*PlatformHealthTally, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *PlatformProjectView) SetHealth(v PlatformHealthTally)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *PlatformProjectView) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetMembers

`func (o *PlatformProjectView) GetMembers() []PlatformProjectApp`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *PlatformProjectView) GetMembersOk() (*[]PlatformProjectApp, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *PlatformProjectView) SetMembers(v []PlatformProjectApp)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *PlatformProjectView) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *PlatformProjectView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformProjectView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformProjectView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformProjectView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespaces

`func (o *PlatformProjectView) GetNamespaces() []string`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *PlatformProjectView) GetNamespacesOk() (*[]string, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *PlatformProjectView) SetNamespaces(v []string)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *PlatformProjectView) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetOrg

`func (o *PlatformProjectView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PlatformProjectView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PlatformProjectView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PlatformProjectView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSync

`func (o *PlatformProjectView) GetSync() PlatformSyncTally`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *PlatformProjectView) GetSyncOk() (*PlatformSyncTally, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *PlatformProjectView) SetSync(v PlatformSyncTally)`

SetSync sets Sync field to given value.

### HasSync

`func (o *PlatformProjectView) HasSync() bool`

HasSync returns a boolean if a field has been set.

### GetUnreadable

`func (o *PlatformProjectView) GetUnreadable() []PlatformUnreadableFile`

GetUnreadable returns the Unreadable field if non-nil, zero value otherwise.

### GetUnreadableOk

`func (o *PlatformProjectView) GetUnreadableOk() (*[]PlatformUnreadableFile, bool)`

GetUnreadableOk returns a tuple with the Unreadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadable

`func (o *PlatformProjectView) SetUnreadable(v []PlatformUnreadableFile)`

SetUnreadable sets Unreadable field to given value.

### HasUnreadable

`func (o *PlatformProjectView) HasUnreadable() bool`

HasUnreadable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



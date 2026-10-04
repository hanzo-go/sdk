# PlatformProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to **int64** | Apps is how many declarations name the project. | [optional] 
**Drift** | Pointer to [**PlatformDriftTally**](PlatformDriftTally.md) | Drift counts its apps by drift severity. | [optional] 
**Health** | Pointer to [**PlatformHealthTally**](PlatformHealthTally.md) | Health counts its apps by CD&#39;s health verdict. | [optional] 
**Name** | Pointer to **string** | Name is the project, as the values files spell it in &#x60;partOf&#x60;. | [optional] 
**Namespaces** | Pointer to **[]string** | Namespaces are the namespaces its apps deploy into, sorted. | [optional] 
**Org** | Pointer to **string** | Org is the organization that owns the project. Absent for the platform&#39;s own projects, whose apps live in the platform&#39;s directories. | [optional] 
**Sync** | Pointer to [**PlatformSyncTally**](PlatformSyncTally.md) | Sync counts its apps by CD&#39;s sync verdict. | [optional] 

## Methods

### NewPlatformProject

`func NewPlatformProject() *PlatformProject`

NewPlatformProject instantiates a new PlatformProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformProjectWithDefaults

`func NewPlatformProjectWithDefaults() *PlatformProject`

NewPlatformProjectWithDefaults instantiates a new PlatformProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *PlatformProject) GetApps() int64`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *PlatformProject) GetAppsOk() (*int64, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *PlatformProject) SetApps(v int64)`

SetApps sets Apps field to given value.

### HasApps

`func (o *PlatformProject) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetDrift

`func (o *PlatformProject) GetDrift() PlatformDriftTally`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *PlatformProject) GetDriftOk() (*PlatformDriftTally, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *PlatformProject) SetDrift(v PlatformDriftTally)`

SetDrift sets Drift field to given value.

### HasDrift

`func (o *PlatformProject) HasDrift() bool`

HasDrift returns a boolean if a field has been set.

### GetHealth

`func (o *PlatformProject) GetHealth() PlatformHealthTally`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *PlatformProject) GetHealthOk() (*PlatformHealthTally, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *PlatformProject) SetHealth(v PlatformHealthTally)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *PlatformProject) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetName

`func (o *PlatformProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespaces

`func (o *PlatformProject) GetNamespaces() []string`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *PlatformProject) GetNamespacesOk() (*[]string, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *PlatformProject) SetNamespaces(v []string)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *PlatformProject) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetOrg

`func (o *PlatformProject) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PlatformProject) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PlatformProject) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PlatformProject) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSync

`func (o *PlatformProject) GetSync() PlatformSyncTally`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *PlatformProject) GetSyncOk() (*PlatformSyncTally, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *PlatformProject) SetSync(v PlatformSyncTally)`

SetSync sets Sync field to given value.

### HasSync

`func (o *PlatformProject) HasSync() bool`

HasSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



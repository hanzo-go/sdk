# PlatformEnvironmentRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the environment&#39;s name, which is also its identity — an environment is derived from the apps that target it, so it has no id of its own. | [optional] 
**Name** | Pointer to **string** | Name is the environment&#39;s name as an app declared it. | [optional] 
**Services** | Pointer to **[]string** | Services are the apps that target this environment, by name. | [optional] 
**Status** | Pointer to **string** | Status rolls up the real states of this environment&#39;s apps: degraded, active, idle or empty. | [optional] 
**Type** | Pointer to **string** | Type buckets the name for display: production, staging, development or custom. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when any of them last changed, RFC3339 UTC; empty when unset. | [optional] 

## Methods

### NewPlatformEnvironmentRow

`func NewPlatformEnvironmentRow() *PlatformEnvironmentRow`

NewPlatformEnvironmentRow instantiates a new PlatformEnvironmentRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformEnvironmentRowWithDefaults

`func NewPlatformEnvironmentRowWithDefaults() *PlatformEnvironmentRow`

NewPlatformEnvironmentRowWithDefaults instantiates a new PlatformEnvironmentRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlatformEnvironmentRow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlatformEnvironmentRow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlatformEnvironmentRow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlatformEnvironmentRow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PlatformEnvironmentRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformEnvironmentRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformEnvironmentRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformEnvironmentRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServices

`func (o *PlatformEnvironmentRow) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *PlatformEnvironmentRow) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *PlatformEnvironmentRow) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *PlatformEnvironmentRow) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetStatus

`func (o *PlatformEnvironmentRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlatformEnvironmentRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlatformEnvironmentRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlatformEnvironmentRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *PlatformEnvironmentRow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlatformEnvironmentRow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlatformEnvironmentRow) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlatformEnvironmentRow) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PlatformEnvironmentRow) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlatformEnvironmentRow) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlatformEnvironmentRow) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PlatformEnvironmentRow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DbEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**BranchId** | Pointer to **string** |  | [optional] 
**RegionId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**CurrentState** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**DbEndpointSettings**](DbEndpointSettings.md) |  | [optional] 
**PoolerEnabled** | Pointer to **bool** |  | [optional] 
**PoolerMode** | Pointer to **string** |  | [optional] 
**LastActive** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDbEndpoint

`func NewDbEndpoint() *DbEndpoint`

NewDbEndpoint instantiates a new DbEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbEndpointWithDefaults

`func NewDbEndpointWithDefaults() *DbEndpoint`

NewDbEndpointWithDefaults instantiates a new DbEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbEndpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DbEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHost

`func (o *DbEndpoint) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DbEndpoint) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DbEndpoint) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DbEndpoint) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetProjectId

`func (o *DbEndpoint) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DbEndpoint) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DbEndpoint) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *DbEndpoint) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetBranchId

`func (o *DbEndpoint) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *DbEndpoint) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *DbEndpoint) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.

### HasBranchId

`func (o *DbEndpoint) HasBranchId() bool`

HasBranchId returns a boolean if a field has been set.

### GetRegionId

`func (o *DbEndpoint) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *DbEndpoint) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *DbEndpoint) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *DbEndpoint) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetType

`func (o *DbEndpoint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DbEndpoint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DbEndpoint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DbEndpoint) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCurrentState

`func (o *DbEndpoint) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *DbEndpoint) GetCurrentStateOk() (*string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *DbEndpoint) SetCurrentState(v string)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *DbEndpoint) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetSettings

`func (o *DbEndpoint) GetSettings() DbEndpointSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *DbEndpoint) GetSettingsOk() (*DbEndpointSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *DbEndpoint) SetSettings(v DbEndpointSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *DbEndpoint) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetPoolerEnabled

`func (o *DbEndpoint) GetPoolerEnabled() bool`

GetPoolerEnabled returns the PoolerEnabled field if non-nil, zero value otherwise.

### GetPoolerEnabledOk

`func (o *DbEndpoint) GetPoolerEnabledOk() (*bool, bool)`

GetPoolerEnabledOk returns a tuple with the PoolerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolerEnabled

`func (o *DbEndpoint) SetPoolerEnabled(v bool)`

SetPoolerEnabled sets PoolerEnabled field to given value.

### HasPoolerEnabled

`func (o *DbEndpoint) HasPoolerEnabled() bool`

HasPoolerEnabled returns a boolean if a field has been set.

### GetPoolerMode

`func (o *DbEndpoint) GetPoolerMode() string`

GetPoolerMode returns the PoolerMode field if non-nil, zero value otherwise.

### GetPoolerModeOk

`func (o *DbEndpoint) GetPoolerModeOk() (*string, bool)`

GetPoolerModeOk returns a tuple with the PoolerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolerMode

`func (o *DbEndpoint) SetPoolerMode(v string)`

SetPoolerMode sets PoolerMode field to given value.

### HasPoolerMode

`func (o *DbEndpoint) HasPoolerMode() bool`

HasPoolerMode returns a boolean if a field has been set.

### GetLastActive

`func (o *DbEndpoint) GetLastActive() time.Time`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *DbEndpoint) GetLastActiveOk() (*time.Time, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *DbEndpoint) SetLastActive(v time.Time)`

SetLastActive sets LastActive field to given value.

### HasLastActive

`func (o *DbEndpoint) HasLastActive() bool`

HasLastActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DbEndpoint) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbEndpoint) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbEndpoint) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbEndpoint) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbEndpoint) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbEndpoint) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbEndpoint) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbEndpoint) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



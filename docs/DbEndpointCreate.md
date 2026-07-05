# DbEndpointCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchId** | **string** |  | 
**Type** | **string** |  | 
**Settings** | Pointer to [**DbEndpointSettings**](DbEndpointSettings.md) |  | [optional] 
**PoolerEnabled** | Pointer to **bool** |  | [optional] [default to true]
**PoolerMode** | Pointer to **string** |  | [optional] [default to "transaction"]

## Methods

### NewDbEndpointCreate

`func NewDbEndpointCreate(branchId string, type_ string, ) *DbEndpointCreate`

NewDbEndpointCreate instantiates a new DbEndpointCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbEndpointCreateWithDefaults

`func NewDbEndpointCreateWithDefaults() *DbEndpointCreate`

NewDbEndpointCreateWithDefaults instantiates a new DbEndpointCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchId

`func (o *DbEndpointCreate) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *DbEndpointCreate) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *DbEndpointCreate) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.


### GetType

`func (o *DbEndpointCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DbEndpointCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DbEndpointCreate) SetType(v string)`

SetType sets Type field to given value.


### GetSettings

`func (o *DbEndpointCreate) GetSettings() DbEndpointSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *DbEndpointCreate) GetSettingsOk() (*DbEndpointSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *DbEndpointCreate) SetSettings(v DbEndpointSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *DbEndpointCreate) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetPoolerEnabled

`func (o *DbEndpointCreate) GetPoolerEnabled() bool`

GetPoolerEnabled returns the PoolerEnabled field if non-nil, zero value otherwise.

### GetPoolerEnabledOk

`func (o *DbEndpointCreate) GetPoolerEnabledOk() (*bool, bool)`

GetPoolerEnabledOk returns a tuple with the PoolerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolerEnabled

`func (o *DbEndpointCreate) SetPoolerEnabled(v bool)`

SetPoolerEnabled sets PoolerEnabled field to given value.

### HasPoolerEnabled

`func (o *DbEndpointCreate) HasPoolerEnabled() bool`

HasPoolerEnabled returns a boolean if a field has been set.

### GetPoolerMode

`func (o *DbEndpointCreate) GetPoolerMode() string`

GetPoolerMode returns the PoolerMode field if non-nil, zero value otherwise.

### GetPoolerModeOk

`func (o *DbEndpointCreate) GetPoolerModeOk() (*string, bool)`

GetPoolerModeOk returns a tuple with the PoolerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolerMode

`func (o *DbEndpointCreate) SetPoolerMode(v string)`

SetPoolerMode sets PoolerMode field to given value.

### HasPoolerMode

`func (o *DbEndpointCreate) HasPoolerMode() bool`

HasPoolerMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



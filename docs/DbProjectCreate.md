# DbProjectCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**RegionId** | Pointer to **string** |  | [optional] [default to "us-east-1"]
**PgVersion** | Pointer to **int32** |  | [optional] [default to 16]
**StorePasswords** | Pointer to **bool** |  | [optional] [default to true]
**DefaultEndpointSettings** | Pointer to [**DbEndpointSettings**](DbEndpointSettings.md) |  | [optional] 

## Methods

### NewDbProjectCreate

`func NewDbProjectCreate(name string, ) *DbProjectCreate`

NewDbProjectCreate instantiates a new DbProjectCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbProjectCreateWithDefaults

`func NewDbProjectCreateWithDefaults() *DbProjectCreate`

NewDbProjectCreateWithDefaults instantiates a new DbProjectCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DbProjectCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DbProjectCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DbProjectCreate) SetName(v string)`

SetName sets Name field to given value.


### GetRegionId

`func (o *DbProjectCreate) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *DbProjectCreate) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *DbProjectCreate) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *DbProjectCreate) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetPgVersion

`func (o *DbProjectCreate) GetPgVersion() int32`

GetPgVersion returns the PgVersion field if non-nil, zero value otherwise.

### GetPgVersionOk

`func (o *DbProjectCreate) GetPgVersionOk() (*int32, bool)`

GetPgVersionOk returns a tuple with the PgVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgVersion

`func (o *DbProjectCreate) SetPgVersion(v int32)`

SetPgVersion sets PgVersion field to given value.

### HasPgVersion

`func (o *DbProjectCreate) HasPgVersion() bool`

HasPgVersion returns a boolean if a field has been set.

### GetStorePasswords

`func (o *DbProjectCreate) GetStorePasswords() bool`

GetStorePasswords returns the StorePasswords field if non-nil, zero value otherwise.

### GetStorePasswordsOk

`func (o *DbProjectCreate) GetStorePasswordsOk() (*bool, bool)`

GetStorePasswordsOk returns a tuple with the StorePasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePasswords

`func (o *DbProjectCreate) SetStorePasswords(v bool)`

SetStorePasswords sets StorePasswords field to given value.

### HasStorePasswords

`func (o *DbProjectCreate) HasStorePasswords() bool`

HasStorePasswords returns a boolean if a field has been set.

### GetDefaultEndpointSettings

`func (o *DbProjectCreate) GetDefaultEndpointSettings() DbEndpointSettings`

GetDefaultEndpointSettings returns the DefaultEndpointSettings field if non-nil, zero value otherwise.

### GetDefaultEndpointSettingsOk

`func (o *DbProjectCreate) GetDefaultEndpointSettingsOk() (*DbEndpointSettings, bool)`

GetDefaultEndpointSettingsOk returns a tuple with the DefaultEndpointSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEndpointSettings

`func (o *DbProjectCreate) SetDefaultEndpointSettings(v DbEndpointSettings)`

SetDefaultEndpointSettings sets DefaultEndpointSettings field to given value.

### HasDefaultEndpointSettings

`func (o *DbProjectCreate) HasDefaultEndpointSettings() bool`

HasDefaultEndpointSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DbProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PlatformId** | Pointer to **string** |  | [optional] 
**RegionId** | Pointer to **string** |  | [optional] 
**PgVersion** | Pointer to **int32** | PostgreSQL major version | [optional] 
**StorePasswords** | Pointer to **bool** |  | [optional] 
**ActiveTimeSeconds** | Pointer to **int64** |  | [optional] 
**ComputeTimeSeconds** | Pointer to **int64** |  | [optional] 
**DataStorageBytesHour** | Pointer to **int64** |  | [optional] 
**DataTransferBytes** | Pointer to **int64** |  | [optional] 
**BranchLogicalSizeLimit** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDbProject

`func NewDbProject() *DbProject`

NewDbProject instantiates a new DbProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbProjectWithDefaults

`func NewDbProjectWithDefaults() *DbProject`

NewDbProjectWithDefaults instantiates a new DbProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbProject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DbProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DbProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DbProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DbProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DbProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatformId

`func (o *DbProject) GetPlatformId() string`

GetPlatformId returns the PlatformId field if non-nil, zero value otherwise.

### GetPlatformIdOk

`func (o *DbProject) GetPlatformIdOk() (*string, bool)`

GetPlatformIdOk returns a tuple with the PlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformId

`func (o *DbProject) SetPlatformId(v string)`

SetPlatformId sets PlatformId field to given value.

### HasPlatformId

`func (o *DbProject) HasPlatformId() bool`

HasPlatformId returns a boolean if a field has been set.

### GetRegionId

`func (o *DbProject) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *DbProject) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *DbProject) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *DbProject) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetPgVersion

`func (o *DbProject) GetPgVersion() int32`

GetPgVersion returns the PgVersion field if non-nil, zero value otherwise.

### GetPgVersionOk

`func (o *DbProject) GetPgVersionOk() (*int32, bool)`

GetPgVersionOk returns a tuple with the PgVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgVersion

`func (o *DbProject) SetPgVersion(v int32)`

SetPgVersion sets PgVersion field to given value.

### HasPgVersion

`func (o *DbProject) HasPgVersion() bool`

HasPgVersion returns a boolean if a field has been set.

### GetStorePasswords

`func (o *DbProject) GetStorePasswords() bool`

GetStorePasswords returns the StorePasswords field if non-nil, zero value otherwise.

### GetStorePasswordsOk

`func (o *DbProject) GetStorePasswordsOk() (*bool, bool)`

GetStorePasswordsOk returns a tuple with the StorePasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePasswords

`func (o *DbProject) SetStorePasswords(v bool)`

SetStorePasswords sets StorePasswords field to given value.

### HasStorePasswords

`func (o *DbProject) HasStorePasswords() bool`

HasStorePasswords returns a boolean if a field has been set.

### GetActiveTimeSeconds

`func (o *DbProject) GetActiveTimeSeconds() int64`

GetActiveTimeSeconds returns the ActiveTimeSeconds field if non-nil, zero value otherwise.

### GetActiveTimeSecondsOk

`func (o *DbProject) GetActiveTimeSecondsOk() (*int64, bool)`

GetActiveTimeSecondsOk returns a tuple with the ActiveTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTimeSeconds

`func (o *DbProject) SetActiveTimeSeconds(v int64)`

SetActiveTimeSeconds sets ActiveTimeSeconds field to given value.

### HasActiveTimeSeconds

`func (o *DbProject) HasActiveTimeSeconds() bool`

HasActiveTimeSeconds returns a boolean if a field has been set.

### GetComputeTimeSeconds

`func (o *DbProject) GetComputeTimeSeconds() int64`

GetComputeTimeSeconds returns the ComputeTimeSeconds field if non-nil, zero value otherwise.

### GetComputeTimeSecondsOk

`func (o *DbProject) GetComputeTimeSecondsOk() (*int64, bool)`

GetComputeTimeSecondsOk returns a tuple with the ComputeTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeTimeSeconds

`func (o *DbProject) SetComputeTimeSeconds(v int64)`

SetComputeTimeSeconds sets ComputeTimeSeconds field to given value.

### HasComputeTimeSeconds

`func (o *DbProject) HasComputeTimeSeconds() bool`

HasComputeTimeSeconds returns a boolean if a field has been set.

### GetDataStorageBytesHour

`func (o *DbProject) GetDataStorageBytesHour() int64`

GetDataStorageBytesHour returns the DataStorageBytesHour field if non-nil, zero value otherwise.

### GetDataStorageBytesHourOk

`func (o *DbProject) GetDataStorageBytesHourOk() (*int64, bool)`

GetDataStorageBytesHourOk returns a tuple with the DataStorageBytesHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStorageBytesHour

`func (o *DbProject) SetDataStorageBytesHour(v int64)`

SetDataStorageBytesHour sets DataStorageBytesHour field to given value.

### HasDataStorageBytesHour

`func (o *DbProject) HasDataStorageBytesHour() bool`

HasDataStorageBytesHour returns a boolean if a field has been set.

### GetDataTransferBytes

`func (o *DbProject) GetDataTransferBytes() int64`

GetDataTransferBytes returns the DataTransferBytes field if non-nil, zero value otherwise.

### GetDataTransferBytesOk

`func (o *DbProject) GetDataTransferBytesOk() (*int64, bool)`

GetDataTransferBytesOk returns a tuple with the DataTransferBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferBytes

`func (o *DbProject) SetDataTransferBytes(v int64)`

SetDataTransferBytes sets DataTransferBytes field to given value.

### HasDataTransferBytes

`func (o *DbProject) HasDataTransferBytes() bool`

HasDataTransferBytes returns a boolean if a field has been set.

### GetBranchLogicalSizeLimit

`func (o *DbProject) GetBranchLogicalSizeLimit() int64`

GetBranchLogicalSizeLimit returns the BranchLogicalSizeLimit field if non-nil, zero value otherwise.

### GetBranchLogicalSizeLimitOk

`func (o *DbProject) GetBranchLogicalSizeLimitOk() (*int64, bool)`

GetBranchLogicalSizeLimitOk returns a tuple with the BranchLogicalSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchLogicalSizeLimit

`func (o *DbProject) SetBranchLogicalSizeLimit(v int64)`

SetBranchLogicalSizeLimit sets BranchLogicalSizeLimit field to given value.

### HasBranchLogicalSizeLimit

`func (o *DbProject) HasBranchLogicalSizeLimit() bool`

HasBranchLogicalSizeLimit returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DbProject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbProject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbProject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbProject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbProject) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbProject) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbProject) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbProject) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# KmsSecretSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **string** | Sync destination (e.g., aws-parameter-store, github) | [optional] 
**SourceEnvironment** | Pointer to **string** |  | [optional] 
**SourcePath** | Pointer to **string** |  | [optional] 
**ConnectionId** | Pointer to **string** |  | [optional] 
**IsAutoSyncEnabled** | Pointer to **bool** |  | [optional] 
**LastSyncedAt** | Pointer to **time.Time** |  | [optional] 
**SyncStatus** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewKmsSecretSync

`func NewKmsSecretSync() *KmsSecretSync`

NewKmsSecretSync instantiates a new KmsSecretSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsSecretSyncWithDefaults

`func NewKmsSecretSyncWithDefaults() *KmsSecretSync`

NewKmsSecretSyncWithDefaults instantiates a new KmsSecretSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KmsSecretSync) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KmsSecretSync) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KmsSecretSync) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KmsSecretSync) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KmsSecretSync) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsSecretSync) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsSecretSync) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KmsSecretSync) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDestination

`func (o *KmsSecretSync) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *KmsSecretSync) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *KmsSecretSync) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *KmsSecretSync) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetSourceEnvironment

`func (o *KmsSecretSync) GetSourceEnvironment() string`

GetSourceEnvironment returns the SourceEnvironment field if non-nil, zero value otherwise.

### GetSourceEnvironmentOk

`func (o *KmsSecretSync) GetSourceEnvironmentOk() (*string, bool)`

GetSourceEnvironmentOk returns a tuple with the SourceEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnvironment

`func (o *KmsSecretSync) SetSourceEnvironment(v string)`

SetSourceEnvironment sets SourceEnvironment field to given value.

### HasSourceEnvironment

`func (o *KmsSecretSync) HasSourceEnvironment() bool`

HasSourceEnvironment returns a boolean if a field has been set.

### GetSourcePath

`func (o *KmsSecretSync) GetSourcePath() string`

GetSourcePath returns the SourcePath field if non-nil, zero value otherwise.

### GetSourcePathOk

`func (o *KmsSecretSync) GetSourcePathOk() (*string, bool)`

GetSourcePathOk returns a tuple with the SourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePath

`func (o *KmsSecretSync) SetSourcePath(v string)`

SetSourcePath sets SourcePath field to given value.

### HasSourcePath

`func (o *KmsSecretSync) HasSourcePath() bool`

HasSourcePath returns a boolean if a field has been set.

### GetConnectionId

`func (o *KmsSecretSync) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *KmsSecretSync) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *KmsSecretSync) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *KmsSecretSync) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetIsAutoSyncEnabled

`func (o *KmsSecretSync) GetIsAutoSyncEnabled() bool`

GetIsAutoSyncEnabled returns the IsAutoSyncEnabled field if non-nil, zero value otherwise.

### GetIsAutoSyncEnabledOk

`func (o *KmsSecretSync) GetIsAutoSyncEnabledOk() (*bool, bool)`

GetIsAutoSyncEnabledOk returns a tuple with the IsAutoSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoSyncEnabled

`func (o *KmsSecretSync) SetIsAutoSyncEnabled(v bool)`

SetIsAutoSyncEnabled sets IsAutoSyncEnabled field to given value.

### HasIsAutoSyncEnabled

`func (o *KmsSecretSync) HasIsAutoSyncEnabled() bool`

HasIsAutoSyncEnabled returns a boolean if a field has been set.

### GetLastSyncedAt

`func (o *KmsSecretSync) GetLastSyncedAt() time.Time`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *KmsSecretSync) GetLastSyncedAtOk() (*time.Time, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *KmsSecretSync) SetLastSyncedAt(v time.Time)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *KmsSecretSync) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### GetSyncStatus

`func (o *KmsSecretSync) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *KmsSecretSync) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *KmsSecretSync) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *KmsSecretSync) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KmsSecretSync) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KmsSecretSync) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KmsSecretSync) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KmsSecretSync) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KmsSecretSync) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KmsSecretSync) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KmsSecretSync) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KmsSecretSync) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PlatformSyncTally

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutOfSync** | Pointer to **int64** | OutOfSync is how many differ from git until someone syncs them. | [optional] 
**Synced** | Pointer to **int64** | Synced is how many run what git declares. | [optional] 
**Unknown** | Pointer to **int64** | Unknown is how many CD holds no verdict for, or could not be read. | [optional] 

## Methods

### NewPlatformSyncTally

`func NewPlatformSyncTally() *PlatformSyncTally`

NewPlatformSyncTally instantiates a new PlatformSyncTally object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformSyncTallyWithDefaults

`func NewPlatformSyncTallyWithDefaults() *PlatformSyncTally`

NewPlatformSyncTallyWithDefaults instantiates a new PlatformSyncTally object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutOfSync

`func (o *PlatformSyncTally) GetOutOfSync() int64`

GetOutOfSync returns the OutOfSync field if non-nil, zero value otherwise.

### GetOutOfSyncOk

`func (o *PlatformSyncTally) GetOutOfSyncOk() (*int64, bool)`

GetOutOfSyncOk returns a tuple with the OutOfSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfSync

`func (o *PlatformSyncTally) SetOutOfSync(v int64)`

SetOutOfSync sets OutOfSync field to given value.

### HasOutOfSync

`func (o *PlatformSyncTally) HasOutOfSync() bool`

HasOutOfSync returns a boolean if a field has been set.

### GetSynced

`func (o *PlatformSyncTally) GetSynced() int64`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *PlatformSyncTally) GetSyncedOk() (*int64, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *PlatformSyncTally) SetSynced(v int64)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *PlatformSyncTally) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### GetUnknown

`func (o *PlatformSyncTally) GetUnknown() int64`

GetUnknown returns the Unknown field if non-nil, zero value otherwise.

### GetUnknownOk

`func (o *PlatformSyncTally) GetUnknownOk() (*int64, bool)`

GetUnknownOk returns a tuple with the Unknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknown

`func (o *PlatformSyncTally) SetUnknown(v int64)`

SetUnknown sets Unknown field to given value.

### HasUnknown

`func (o *PlatformSyncTally) HasUnknown() bool`

HasUnknown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



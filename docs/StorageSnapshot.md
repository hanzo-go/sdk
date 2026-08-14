# StorageSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**[]StorageAlert**](StorageAlert.md) |  | [optional] 
**Datastore** | Pointer to [**DatastoreVolume**](DatastoreVolume.md) |  | [optional] 
**Fleet** | Pointer to [**StorageFleet**](StorageFleet.md) |  | [optional] 
**Volumes** | Pointer to [**[]StorageVolume**](StorageVolume.md) |  | [optional] 

## Methods

### NewStorageSnapshot

`func NewStorageSnapshot() *StorageSnapshot`

NewStorageSnapshot instantiates a new StorageSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSnapshotWithDefaults

`func NewStorageSnapshotWithDefaults() *StorageSnapshot`

NewStorageSnapshotWithDefaults instantiates a new StorageSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *StorageSnapshot) GetAlerts() []StorageAlert`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *StorageSnapshot) GetAlertsOk() (*[]StorageAlert, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *StorageSnapshot) SetAlerts(v []StorageAlert)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *StorageSnapshot) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetDatastore

`func (o *StorageSnapshot) GetDatastore() DatastoreVolume`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *StorageSnapshot) GetDatastoreOk() (*DatastoreVolume, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *StorageSnapshot) SetDatastore(v DatastoreVolume)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *StorageSnapshot) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetFleet

`func (o *StorageSnapshot) GetFleet() StorageFleet`

GetFleet returns the Fleet field if non-nil, zero value otherwise.

### GetFleetOk

`func (o *StorageSnapshot) GetFleetOk() (*StorageFleet, bool)`

GetFleetOk returns a tuple with the Fleet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleet

`func (o *StorageSnapshot) SetFleet(v StorageFleet)`

SetFleet sets Fleet field to given value.

### HasFleet

`func (o *StorageSnapshot) HasFleet() bool`

HasFleet returns a boolean if a field has been set.

### GetVolumes

`func (o *StorageSnapshot) GetVolumes() []StorageVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *StorageSnapshot) GetVolumesOk() (*[]StorageVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *StorageSnapshot) SetVolumes(v []StorageVolume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *StorageSnapshot) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



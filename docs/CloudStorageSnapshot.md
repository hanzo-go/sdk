# CloudStorageSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**[]CloudStorageAlert**](CloudStorageAlert.md) |  | [optional] 
**Datastore** | Pointer to [**CloudDatastoreVolume**](CloudDatastoreVolume.md) |  | [optional] 
**Fleet** | Pointer to [**CloudStorageFleet**](CloudStorageFleet.md) |  | [optional] 
**Volumes** | Pointer to [**[]CloudStorageVolume**](CloudStorageVolume.md) |  | [optional] 

## Methods

### NewCloudStorageSnapshot

`func NewCloudStorageSnapshot() *CloudStorageSnapshot`

NewCloudStorageSnapshot instantiates a new CloudStorageSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudStorageSnapshotWithDefaults

`func NewCloudStorageSnapshotWithDefaults() *CloudStorageSnapshot`

NewCloudStorageSnapshotWithDefaults instantiates a new CloudStorageSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *CloudStorageSnapshot) GetAlerts() []CloudStorageAlert`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *CloudStorageSnapshot) GetAlertsOk() (*[]CloudStorageAlert, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *CloudStorageSnapshot) SetAlerts(v []CloudStorageAlert)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *CloudStorageSnapshot) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetDatastore

`func (o *CloudStorageSnapshot) GetDatastore() CloudDatastoreVolume`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *CloudStorageSnapshot) GetDatastoreOk() (*CloudDatastoreVolume, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *CloudStorageSnapshot) SetDatastore(v CloudDatastoreVolume)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *CloudStorageSnapshot) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetFleet

`func (o *CloudStorageSnapshot) GetFleet() CloudStorageFleet`

GetFleet returns the Fleet field if non-nil, zero value otherwise.

### GetFleetOk

`func (o *CloudStorageSnapshot) GetFleetOk() (*CloudStorageFleet, bool)`

GetFleetOk returns a tuple with the Fleet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleet

`func (o *CloudStorageSnapshot) SetFleet(v CloudStorageFleet)`

SetFleet sets Fleet field to given value.

### HasFleet

`func (o *CloudStorageSnapshot) HasFleet() bool`

HasFleet returns a boolean if a field has been set.

### GetVolumes

`func (o *CloudStorageSnapshot) GetVolumes() []CloudStorageVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *CloudStorageSnapshot) GetVolumesOk() (*[]CloudStorageVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *CloudStorageSnapshot) SetVolumes(v []CloudStorageVolume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *CloudStorageSnapshot) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



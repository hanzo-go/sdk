# CloudVolumeIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the DO volume id, from the path. | [optional] 
**Name** | Pointer to **string** | Name is the snapshot name on the snapshot action. Blank gets a deterministic \&quot;&lt;volume&gt;-predelete-&lt;unix&gt;\&quot; so the undo is findable in the DO console. | [optional] 
**SizeGiB** | Pointer to **int32** | SizeGiB is the target size on the resize action. A volume only ever grows — ExpandTo is the verdict that refuses a shrink, so this is not validated here. | [optional] 
**Snapshot** | Pointer to **string** | Snapshot is the snapshot-first switch on DELETE. Anything other than the literal \&quot;false\&quot; snapshots before destroying — the snapshot IS the undo, so waiving it is deliberate and explicit. | [optional] 

## Methods

### NewCloudVolumeIn

`func NewCloudVolumeIn() *CloudVolumeIn`

NewCloudVolumeIn instantiates a new CloudVolumeIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVolumeInWithDefaults

`func NewCloudVolumeInWithDefaults() *CloudVolumeIn`

NewCloudVolumeInWithDefaults instantiates a new CloudVolumeIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudVolumeIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudVolumeIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudVolumeIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudVolumeIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudVolumeIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudVolumeIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudVolumeIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudVolumeIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSizeGiB

`func (o *CloudVolumeIn) GetSizeGiB() int32`

GetSizeGiB returns the SizeGiB field if non-nil, zero value otherwise.

### GetSizeGiBOk

`func (o *CloudVolumeIn) GetSizeGiBOk() (*int32, bool)`

GetSizeGiBOk returns a tuple with the SizeGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGiB

`func (o *CloudVolumeIn) SetSizeGiB(v int32)`

SetSizeGiB sets SizeGiB field to given value.

### HasSizeGiB

`func (o *CloudVolumeIn) HasSizeGiB() bool`

HasSizeGiB returns a boolean if a field has been set.

### GetSnapshot

`func (o *CloudVolumeIn) GetSnapshot() string`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *CloudVolumeIn) GetSnapshotOk() (*string, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *CloudVolumeIn) SetSnapshot(v string)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *CloudVolumeIn) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# VolumeSnapshotOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**DigitaloceanSnapshot**](DigitaloceanSnapshot.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewVolumeSnapshotOut

`func NewVolumeSnapshotOut() *VolumeSnapshotOut`

NewVolumeSnapshotOut instantiates a new VolumeSnapshotOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeSnapshotOutWithDefaults

`func NewVolumeSnapshotOutWithDefaults() *VolumeSnapshotOut`

NewVolumeSnapshotOutWithDefaults instantiates a new VolumeSnapshotOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VolumeSnapshotOut) GetData() DigitaloceanSnapshot`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VolumeSnapshotOut) GetDataOk() (*DigitaloceanSnapshot, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VolumeSnapshotOut) SetData(v DigitaloceanSnapshot)`

SetData sets Data field to given value.

### HasData

`func (o *VolumeSnapshotOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *VolumeSnapshotOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *VolumeSnapshotOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *VolumeSnapshotOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *VolumeSnapshotOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *VolumeSnapshotOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VolumeSnapshotOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VolumeSnapshotOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VolumeSnapshotOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



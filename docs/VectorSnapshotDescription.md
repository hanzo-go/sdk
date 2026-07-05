# VectorSnapshotDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 

## Methods

### NewVectorSnapshotDescription

`func NewVectorSnapshotDescription() *VectorSnapshotDescription`

NewVectorSnapshotDescription instantiates a new VectorSnapshotDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorSnapshotDescriptionWithDefaults

`func NewVectorSnapshotDescriptionWithDefaults() *VectorSnapshotDescription`

NewVectorSnapshotDescriptionWithDefaults instantiates a new VectorSnapshotDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VectorSnapshotDescription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VectorSnapshotDescription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VectorSnapshotDescription) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VectorSnapshotDescription) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreationTime

`func (o *VectorSnapshotDescription) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *VectorSnapshotDescription) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *VectorSnapshotDescription) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *VectorSnapshotDescription) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetSize

`func (o *VectorSnapshotDescription) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VectorSnapshotDescription) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VectorSnapshotDescription) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *VectorSnapshotDescription) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



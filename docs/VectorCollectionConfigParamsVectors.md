# VectorCollectionConfigParamsVectors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **int32** | Vector dimension | 
**Distance** | **string** |  | 
**OnDisk** | Pointer to **bool** |  | [optional] 

## Methods

### NewVectorCollectionConfigParamsVectors

`func NewVectorCollectionConfigParamsVectors(size int32, distance string, ) *VectorCollectionConfigParamsVectors`

NewVectorCollectionConfigParamsVectors instantiates a new VectorCollectionConfigParamsVectors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorCollectionConfigParamsVectorsWithDefaults

`func NewVectorCollectionConfigParamsVectorsWithDefaults() *VectorCollectionConfigParamsVectors`

NewVectorCollectionConfigParamsVectorsWithDefaults instantiates a new VectorCollectionConfigParamsVectors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *VectorCollectionConfigParamsVectors) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VectorCollectionConfigParamsVectors) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VectorCollectionConfigParamsVectors) SetSize(v int32)`

SetSize sets Size field to given value.


### GetDistance

`func (o *VectorCollectionConfigParamsVectors) GetDistance() string`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *VectorCollectionConfigParamsVectors) GetDistanceOk() (*string, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *VectorCollectionConfigParamsVectors) SetDistance(v string)`

SetDistance sets Distance field to given value.


### GetOnDisk

`func (o *VectorCollectionConfigParamsVectors) GetOnDisk() bool`

GetOnDisk returns the OnDisk field if non-nil, zero value otherwise.

### GetOnDiskOk

`func (o *VectorCollectionConfigParamsVectors) GetOnDiskOk() (*bool, bool)`

GetOnDiskOk returns a tuple with the OnDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDisk

`func (o *VectorCollectionConfigParamsVectors) SetOnDisk(v bool)`

SetOnDisk sets OnDisk field to given value.

### HasOnDisk

`func (o *VectorCollectionConfigParamsVectors) HasOnDisk() bool`

HasOnDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



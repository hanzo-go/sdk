# VectorVectorParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **int32** | Vector dimension | 
**Distance** | **string** |  | 
**OnDisk** | Pointer to **bool** |  | [optional] 

## Methods

### NewVectorVectorParams

`func NewVectorVectorParams(size int32, distance string, ) *VectorVectorParams`

NewVectorVectorParams instantiates a new VectorVectorParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorVectorParamsWithDefaults

`func NewVectorVectorParamsWithDefaults() *VectorVectorParams`

NewVectorVectorParamsWithDefaults instantiates a new VectorVectorParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *VectorVectorParams) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VectorVectorParams) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VectorVectorParams) SetSize(v int32)`

SetSize sets Size field to given value.


### GetDistance

`func (o *VectorVectorParams) GetDistance() string`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *VectorVectorParams) GetDistanceOk() (*string, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *VectorVectorParams) SetDistance(v string)`

SetDistance sets Distance field to given value.


### GetOnDisk

`func (o *VectorVectorParams) GetOnDisk() bool`

GetOnDisk returns the OnDisk field if non-nil, zero value otherwise.

### GetOnDiskOk

`func (o *VectorVectorParams) GetOnDiskOk() (*bool, bool)`

GetOnDiskOk returns a tuple with the OnDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDisk

`func (o *VectorVectorParams) SetOnDisk(v bool)`

SetOnDisk sets OnDisk field to given value.

### HasOnDisk

`func (o *VectorVectorParams) HasOnDisk() bool`

HasOnDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



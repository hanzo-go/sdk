# ProductVectorCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Collection name. | 
**VectorCount** | **int64** | Number of points/vectors in the collection. | 
**Dimension** | **int64** | Vector dimension. | 
**DistanceMetric** | **string** | Distance metric (defaults to \&quot;cosine\&quot; when the upstream omits it). | 
**StorageBytes** | Pointer to **int64** | Storage size in bytes (omitted when zero). | [optional] 
**CreatedAt** | **string** | Created timestamp (RFC 3339). | 

## Methods

### NewProductVectorCollection

`func NewProductVectorCollection(name string, vectorCount int64, dimension int64, distanceMetric string, createdAt string, ) *ProductVectorCollection`

NewProductVectorCollection instantiates a new ProductVectorCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductVectorCollectionWithDefaults

`func NewProductVectorCollectionWithDefaults() *ProductVectorCollection`

NewProductVectorCollectionWithDefaults instantiates a new ProductVectorCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductVectorCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductVectorCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductVectorCollection) SetName(v string)`

SetName sets Name field to given value.


### GetVectorCount

`func (o *ProductVectorCollection) GetVectorCount() int64`

GetVectorCount returns the VectorCount field if non-nil, zero value otherwise.

### GetVectorCountOk

`func (o *ProductVectorCollection) GetVectorCountOk() (*int64, bool)`

GetVectorCountOk returns a tuple with the VectorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorCount

`func (o *ProductVectorCollection) SetVectorCount(v int64)`

SetVectorCount sets VectorCount field to given value.


### GetDimension

`func (o *ProductVectorCollection) GetDimension() int64`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *ProductVectorCollection) GetDimensionOk() (*int64, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *ProductVectorCollection) SetDimension(v int64)`

SetDimension sets Dimension field to given value.


### GetDistanceMetric

`func (o *ProductVectorCollection) GetDistanceMetric() string`

GetDistanceMetric returns the DistanceMetric field if non-nil, zero value otherwise.

### GetDistanceMetricOk

`func (o *ProductVectorCollection) GetDistanceMetricOk() (*string, bool)`

GetDistanceMetricOk returns a tuple with the DistanceMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceMetric

`func (o *ProductVectorCollection) SetDistanceMetric(v string)`

SetDistanceMetric sets DistanceMetric field to given value.


### GetStorageBytes

`func (o *ProductVectorCollection) GetStorageBytes() int64`

GetStorageBytes returns the StorageBytes field if non-nil, zero value otherwise.

### GetStorageBytesOk

`func (o *ProductVectorCollection) GetStorageBytesOk() (*int64, bool)`

GetStorageBytesOk returns a tuple with the StorageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageBytes

`func (o *ProductVectorCollection) SetStorageBytes(v int64)`

SetStorageBytes sets StorageBytes field to given value.

### HasStorageBytes

`func (o *ProductVectorCollection) HasStorageBytes() bool`

HasStorageBytes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProductVectorCollection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProductVectorCollection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProductVectorCollection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



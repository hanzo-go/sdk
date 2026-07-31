# CloudVectorCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is the collection&#39;s creation time (RFC 3339); Qdrant does not report one, so it is empty today. | [optional] 
**Dimension** | Pointer to **int32** | Dimension is the size of one vector in the collection. | [optional] 
**DistanceMetric** | Pointer to **string** | DistanceMetric is the collection&#39;s distance function; \&quot;cosine\&quot; when the collection&#39;s detail could not be read. | [optional] 
**Name** | Pointer to **string** | Name is the collection name. | [optional] 
**StorageBytes** | Pointer to **int32** | StorageBytes is the collection&#39;s on-disk size, omitted when unknown. | [optional] 
**VectorCount** | Pointer to **int32** | VectorCount is the collection&#39;s point count. | [optional] 

## Methods

### NewCloudVectorCollection

`func NewCloudVectorCollection() *CloudVectorCollection`

NewCloudVectorCollection instantiates a new CloudVectorCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVectorCollectionWithDefaults

`func NewCloudVectorCollectionWithDefaults() *CloudVectorCollection`

NewCloudVectorCollectionWithDefaults instantiates a new CloudVectorCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloudVectorCollection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudVectorCollection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudVectorCollection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudVectorCollection) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDimension

`func (o *CloudVectorCollection) GetDimension() int32`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *CloudVectorCollection) GetDimensionOk() (*int32, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *CloudVectorCollection) SetDimension(v int32)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *CloudVectorCollection) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetDistanceMetric

`func (o *CloudVectorCollection) GetDistanceMetric() string`

GetDistanceMetric returns the DistanceMetric field if non-nil, zero value otherwise.

### GetDistanceMetricOk

`func (o *CloudVectorCollection) GetDistanceMetricOk() (*string, bool)`

GetDistanceMetricOk returns a tuple with the DistanceMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceMetric

`func (o *CloudVectorCollection) SetDistanceMetric(v string)`

SetDistanceMetric sets DistanceMetric field to given value.

### HasDistanceMetric

`func (o *CloudVectorCollection) HasDistanceMetric() bool`

HasDistanceMetric returns a boolean if a field has been set.

### GetName

`func (o *CloudVectorCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudVectorCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudVectorCollection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudVectorCollection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorageBytes

`func (o *CloudVectorCollection) GetStorageBytes() int32`

GetStorageBytes returns the StorageBytes field if non-nil, zero value otherwise.

### GetStorageBytesOk

`func (o *CloudVectorCollection) GetStorageBytesOk() (*int32, bool)`

GetStorageBytesOk returns a tuple with the StorageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageBytes

`func (o *CloudVectorCollection) SetStorageBytes(v int32)`

SetStorageBytes sets StorageBytes field to given value.

### HasStorageBytes

`func (o *CloudVectorCollection) HasStorageBytes() bool`

HasStorageBytes returns a boolean if a field has been set.

### GetVectorCount

`func (o *CloudVectorCollection) GetVectorCount() int32`

GetVectorCount returns the VectorCount field if non-nil, zero value otherwise.

### GetVectorCountOk

`func (o *CloudVectorCollection) GetVectorCountOk() (*int32, bool)`

GetVectorCountOk returns a tuple with the VectorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorCount

`func (o *CloudVectorCollection) SetVectorCount(v int32)`

SetVectorCount sets VectorCount field to given value.

### HasVectorCount

`func (o *CloudVectorCollection) HasVectorCount() bool`

HasVectorCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



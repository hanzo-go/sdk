# VectorSearchRequestVector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Vector** | Pointer to **[]float32** |  | [optional] 

## Methods

### NewVectorSearchRequestVector

`func NewVectorSearchRequestVector() *VectorSearchRequestVector`

NewVectorSearchRequestVector instantiates a new VectorSearchRequestVector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorSearchRequestVectorWithDefaults

`func NewVectorSearchRequestVectorWithDefaults() *VectorSearchRequestVector`

NewVectorSearchRequestVectorWithDefaults instantiates a new VectorSearchRequestVector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VectorSearchRequestVector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VectorSearchRequestVector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VectorSearchRequestVector) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VectorSearchRequestVector) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVector

`func (o *VectorSearchRequestVector) GetVector() []float32`

GetVector returns the Vector field if non-nil, zero value otherwise.

### GetVectorOk

`func (o *VectorSearchRequestVector) GetVectorOk() (*[]float32, bool)`

GetVectorOk returns a tuple with the Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVector

`func (o *VectorSearchRequestVector) SetVector(v []float32)`

SetVector sets Vector field to given value.

### HasVector

`func (o *VectorSearchRequestVector) HasVector() bool`

HasVector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



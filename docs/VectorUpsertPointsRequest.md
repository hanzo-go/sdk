# VectorUpsertPointsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Points** | [**[]VectorPointStruct**](VectorPointStruct.md) |  | 

## Methods

### NewVectorUpsertPointsRequest

`func NewVectorUpsertPointsRequest(points []VectorPointStruct, ) *VectorUpsertPointsRequest`

NewVectorUpsertPointsRequest instantiates a new VectorUpsertPointsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorUpsertPointsRequestWithDefaults

`func NewVectorUpsertPointsRequestWithDefaults() *VectorUpsertPointsRequest`

NewVectorUpsertPointsRequestWithDefaults instantiates a new VectorUpsertPointsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoints

`func (o *VectorUpsertPointsRequest) GetPoints() []VectorPointStruct`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *VectorUpsertPointsRequest) GetPointsOk() (*[]VectorPointStruct, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *VectorUpsertPointsRequest) SetPoints(v []VectorPointStruct)`

SetPoints sets Points field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



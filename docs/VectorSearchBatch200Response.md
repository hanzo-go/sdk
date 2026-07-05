# VectorSearchBatch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**[][]VectorScoredPoint**]([]VectorScoredPoint.md) |  | [optional] 

## Methods

### NewVectorSearchBatch200Response

`func NewVectorSearchBatch200Response() *VectorSearchBatch200Response`

NewVectorSearchBatch200Response instantiates a new VectorSearchBatch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorSearchBatch200ResponseWithDefaults

`func NewVectorSearchBatch200ResponseWithDefaults() *VectorSearchBatch200Response`

NewVectorSearchBatch200ResponseWithDefaults instantiates a new VectorSearchBatch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *VectorSearchBatch200Response) GetResult() [][]VectorScoredPoint`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *VectorSearchBatch200Response) GetResultOk() (*[][]VectorScoredPoint, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *VectorSearchBatch200Response) SetResult(v [][]VectorScoredPoint)`

SetResult sets Result field to given value.

### HasResult

`func (o *VectorSearchBatch200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



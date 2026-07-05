# VectorSearchBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Searches** | [**[]VectorSearchRequest**](VectorSearchRequest.md) |  | 

## Methods

### NewVectorSearchBatchRequest

`func NewVectorSearchBatchRequest(searches []VectorSearchRequest, ) *VectorSearchBatchRequest`

NewVectorSearchBatchRequest instantiates a new VectorSearchBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorSearchBatchRequestWithDefaults

`func NewVectorSearchBatchRequestWithDefaults() *VectorSearchBatchRequest`

NewVectorSearchBatchRequestWithDefaults instantiates a new VectorSearchBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearches

`func (o *VectorSearchBatchRequest) GetSearches() []VectorSearchRequest`

GetSearches returns the Searches field if non-nil, zero value otherwise.

### GetSearchesOk

`func (o *VectorSearchBatchRequest) GetSearchesOk() (*[]VectorSearchRequest, bool)`

GetSearchesOk returns a tuple with the Searches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearches

`func (o *VectorSearchBatchRequest) SetSearches(v []VectorSearchRequest)`

SetSearches sets Searches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



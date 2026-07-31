# VectorRecommendPointsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Positive** | [**[]VectorPointId**](VectorPointId.md) |  | 
**Negative** | Pointer to [**[]VectorPointId**](VectorPointId.md) |  | [optional] 
**Filter** | Pointer to [**VectorFilter**](VectorFilter.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] [default to 10]
**WithPayload** | Pointer to **bool** |  | [optional] 
**WithVector** | Pointer to **bool** |  | [optional] 

## Methods

### NewVectorRecommendPointsRequest

`func NewVectorRecommendPointsRequest(positive []VectorPointId, ) *VectorRecommendPointsRequest`

NewVectorRecommendPointsRequest instantiates a new VectorRecommendPointsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorRecommendPointsRequestWithDefaults

`func NewVectorRecommendPointsRequestWithDefaults() *VectorRecommendPointsRequest`

NewVectorRecommendPointsRequestWithDefaults instantiates a new VectorRecommendPointsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositive

`func (o *VectorRecommendPointsRequest) GetPositive() []VectorPointId`

GetPositive returns the Positive field if non-nil, zero value otherwise.

### GetPositiveOk

`func (o *VectorRecommendPointsRequest) GetPositiveOk() (*[]VectorPointId, bool)`

GetPositiveOk returns a tuple with the Positive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositive

`func (o *VectorRecommendPointsRequest) SetPositive(v []VectorPointId)`

SetPositive sets Positive field to given value.


### GetNegative

`func (o *VectorRecommendPointsRequest) GetNegative() []VectorPointId`

GetNegative returns the Negative field if non-nil, zero value otherwise.

### GetNegativeOk

`func (o *VectorRecommendPointsRequest) GetNegativeOk() (*[]VectorPointId, bool)`

GetNegativeOk returns a tuple with the Negative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegative

`func (o *VectorRecommendPointsRequest) SetNegative(v []VectorPointId)`

SetNegative sets Negative field to given value.

### HasNegative

`func (o *VectorRecommendPointsRequest) HasNegative() bool`

HasNegative returns a boolean if a field has been set.

### GetFilter

`func (o *VectorRecommendPointsRequest) GetFilter() VectorFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *VectorRecommendPointsRequest) GetFilterOk() (*VectorFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *VectorRecommendPointsRequest) SetFilter(v VectorFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *VectorRecommendPointsRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetLimit

`func (o *VectorRecommendPointsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *VectorRecommendPointsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *VectorRecommendPointsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *VectorRecommendPointsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetWithPayload

`func (o *VectorRecommendPointsRequest) GetWithPayload() bool`

GetWithPayload returns the WithPayload field if non-nil, zero value otherwise.

### GetWithPayloadOk

`func (o *VectorRecommendPointsRequest) GetWithPayloadOk() (*bool, bool)`

GetWithPayloadOk returns a tuple with the WithPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithPayload

`func (o *VectorRecommendPointsRequest) SetWithPayload(v bool)`

SetWithPayload sets WithPayload field to given value.

### HasWithPayload

`func (o *VectorRecommendPointsRequest) HasWithPayload() bool`

HasWithPayload returns a boolean if a field has been set.

### GetWithVector

`func (o *VectorRecommendPointsRequest) GetWithVector() bool`

GetWithVector returns the WithVector field if non-nil, zero value otherwise.

### GetWithVectorOk

`func (o *VectorRecommendPointsRequest) GetWithVectorOk() (*bool, bool)`

GetWithVectorOk returns a tuple with the WithVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithVector

`func (o *VectorRecommendPointsRequest) SetWithVector(v bool)`

SetWithVector sets WithVector field to given value.

### HasWithVector

`func (o *VectorRecommendPointsRequest) HasWithVector() bool`

HasWithVector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



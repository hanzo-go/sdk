# VectorSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vector** | [**VectorVectorQuery**](VectorVectorQuery.md) |  | 
**Filter** | Pointer to [**VectorFilter**](VectorFilter.md) |  | [optional] 
**Params** | Pointer to [**VectorSearchParams**](VectorSearchParams.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] [default to 10]
**Offset** | Pointer to **int32** |  | [optional] [default to 0]
**WithPayload** | Pointer to [**VectorFieldSelector**](VectorFieldSelector.md) |  | [optional] 
**WithVector** | Pointer to [**VectorFieldSelector**](VectorFieldSelector.md) |  | [optional] 
**ScoreThreshold** | Pointer to **float32** |  | [optional] 

## Methods

### NewVectorSearchRequest

`func NewVectorSearchRequest(vector VectorVectorQuery, ) *VectorSearchRequest`

NewVectorSearchRequest instantiates a new VectorSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorSearchRequestWithDefaults

`func NewVectorSearchRequestWithDefaults() *VectorSearchRequest`

NewVectorSearchRequestWithDefaults instantiates a new VectorSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVector

`func (o *VectorSearchRequest) GetVector() VectorVectorQuery`

GetVector returns the Vector field if non-nil, zero value otherwise.

### GetVectorOk

`func (o *VectorSearchRequest) GetVectorOk() (*VectorVectorQuery, bool)`

GetVectorOk returns a tuple with the Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVector

`func (o *VectorSearchRequest) SetVector(v VectorVectorQuery)`

SetVector sets Vector field to given value.


### GetFilter

`func (o *VectorSearchRequest) GetFilter() VectorFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *VectorSearchRequest) GetFilterOk() (*VectorFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *VectorSearchRequest) SetFilter(v VectorFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *VectorSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetParams

`func (o *VectorSearchRequest) GetParams() VectorSearchParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *VectorSearchRequest) GetParamsOk() (*VectorSearchParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *VectorSearchRequest) SetParams(v VectorSearchParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *VectorSearchRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetLimit

`func (o *VectorSearchRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *VectorSearchRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *VectorSearchRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *VectorSearchRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *VectorSearchRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *VectorSearchRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *VectorSearchRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *VectorSearchRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetWithPayload

`func (o *VectorSearchRequest) GetWithPayload() VectorFieldSelector`

GetWithPayload returns the WithPayload field if non-nil, zero value otherwise.

### GetWithPayloadOk

`func (o *VectorSearchRequest) GetWithPayloadOk() (*VectorFieldSelector, bool)`

GetWithPayloadOk returns a tuple with the WithPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithPayload

`func (o *VectorSearchRequest) SetWithPayload(v VectorFieldSelector)`

SetWithPayload sets WithPayload field to given value.

### HasWithPayload

`func (o *VectorSearchRequest) HasWithPayload() bool`

HasWithPayload returns a boolean if a field has been set.

### GetWithVector

`func (o *VectorSearchRequest) GetWithVector() VectorFieldSelector`

GetWithVector returns the WithVector field if non-nil, zero value otherwise.

### GetWithVectorOk

`func (o *VectorSearchRequest) GetWithVectorOk() (*VectorFieldSelector, bool)`

GetWithVectorOk returns a tuple with the WithVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithVector

`func (o *VectorSearchRequest) SetWithVector(v VectorFieldSelector)`

SetWithVector sets WithVector field to given value.

### HasWithVector

`func (o *VectorSearchRequest) HasWithVector() bool`

HasWithVector returns a boolean if a field has been set.

### GetScoreThreshold

`func (o *VectorSearchRequest) GetScoreThreshold() float32`

GetScoreThreshold returns the ScoreThreshold field if non-nil, zero value otherwise.

### GetScoreThresholdOk

`func (o *VectorSearchRequest) GetScoreThresholdOk() (*float32, bool)`

GetScoreThresholdOk returns a tuple with the ScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreThreshold

`func (o *VectorSearchRequest) SetScoreThreshold(v float32)`

SetScoreThreshold sets ScoreThreshold field to given value.

### HasScoreThreshold

`func (o *VectorSearchRequest) HasScoreThreshold() bool`

HasScoreThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



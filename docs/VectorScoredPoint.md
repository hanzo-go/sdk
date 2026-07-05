# VectorScoredPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**VectorDeletePointsRequestPointsInner**](VectorDeletePointsRequestPointsInner.md) |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Score** | Pointer to **float32** |  | [optional] 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 
**Vector** | Pointer to [**VectorScoredPointVector**](VectorScoredPointVector.md) |  | [optional] 

## Methods

### NewVectorScoredPoint

`func NewVectorScoredPoint() *VectorScoredPoint`

NewVectorScoredPoint instantiates a new VectorScoredPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorScoredPointWithDefaults

`func NewVectorScoredPointWithDefaults() *VectorScoredPoint`

NewVectorScoredPointWithDefaults instantiates a new VectorScoredPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VectorScoredPoint) GetId() VectorDeletePointsRequestPointsInner`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VectorScoredPoint) GetIdOk() (*VectorDeletePointsRequestPointsInner, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VectorScoredPoint) SetId(v VectorDeletePointsRequestPointsInner)`

SetId sets Id field to given value.

### HasId

`func (o *VectorScoredPoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *VectorScoredPoint) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VectorScoredPoint) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VectorScoredPoint) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VectorScoredPoint) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetScore

`func (o *VectorScoredPoint) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *VectorScoredPoint) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *VectorScoredPoint) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *VectorScoredPoint) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetPayload

`func (o *VectorScoredPoint) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *VectorScoredPoint) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *VectorScoredPoint) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *VectorScoredPoint) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetVector

`func (o *VectorScoredPoint) GetVector() VectorScoredPointVector`

GetVector returns the Vector field if non-nil, zero value otherwise.

### GetVectorOk

`func (o *VectorScoredPoint) GetVectorOk() (*VectorScoredPointVector, bool)`

GetVectorOk returns a tuple with the Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVector

`func (o *VectorScoredPoint) SetVector(v VectorScoredPointVector)`

SetVector sets Vector field to given value.

### HasVector

`func (o *VectorScoredPoint) HasVector() bool`

HasVector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



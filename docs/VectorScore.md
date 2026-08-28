# VectorScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **float32** |  | [optional] 
**Vector** | Pointer to **string** |  | [optional] 

## Methods

### NewVectorScore

`func NewVectorScore() *VectorScore`

NewVectorScore instantiates a new VectorScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorScoreWithDefaults

`func NewVectorScoreWithDefaults() *VectorScore`

NewVectorScoreWithDefaults instantiates a new VectorScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *VectorScore) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *VectorScore) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *VectorScore) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *VectorScore) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetVector

`func (o *VectorScore) GetVector() string`

GetVector returns the Vector field if non-nil, zero value otherwise.

### GetVectorOk

`func (o *VectorScore) GetVectorOk() (*string, bool)`

GetVectorOk returns a tuple with the Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVector

`func (o *VectorScore) SetVector(v string)`

SetVector sets Vector field to given value.

### HasVector

`func (o *VectorScore) HasVector() bool`

HasVector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



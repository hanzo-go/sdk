# O11yPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**T** | Pointer to **string** | T is the bucket start, RFC3339 in UTC. | [optional] 
**V** | Pointer to **float32** | V is the bucket&#39;s value. | [optional] 

## Methods

### NewO11yPoint

`func NewO11yPoint() *O11yPoint`

NewO11yPoint instantiates a new O11yPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yPointWithDefaults

`func NewO11yPointWithDefaults() *O11yPoint`

NewO11yPointWithDefaults instantiates a new O11yPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetT

`func (o *O11yPoint) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *O11yPoint) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *O11yPoint) SetT(v string)`

SetT sets T field to given value.

### HasT

`func (o *O11yPoint) HasT() bool`

HasT returns a boolean if a field has been set.

### GetV

`func (o *O11yPoint) GetV() float32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *O11yPoint) GetVOk() (*float32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *O11yPoint) SetV(v float32)`

SetV sets V field to given value.

### HasV

`func (o *O11yPoint) HasV() bool`

HasV returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



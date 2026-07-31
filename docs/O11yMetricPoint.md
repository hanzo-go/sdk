# O11yMetricPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**T** | Pointer to **time.Time** | RFC3339 UTC bucket start. | [optional] 
**V** | Pointer to **float32** |  | [optional] 

## Methods

### NewO11yMetricPoint

`func NewO11yMetricPoint() *O11yMetricPoint`

NewO11yMetricPoint instantiates a new O11yMetricPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yMetricPointWithDefaults

`func NewO11yMetricPointWithDefaults() *O11yMetricPoint`

NewO11yMetricPointWithDefaults instantiates a new O11yMetricPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetT

`func (o *O11yMetricPoint) GetT() time.Time`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *O11yMetricPoint) GetTOk() (*time.Time, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *O11yMetricPoint) SetT(v time.Time)`

SetT sets T field to given value.

### HasT

`func (o *O11yMetricPoint) HasT() bool`

HasT returns a boolean if a field has been set.

### GetV

`func (o *O11yMetricPoint) GetV() float32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *O11yMetricPoint) GetVOk() (*float32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *O11yMetricPoint) SetV(v float32)`

SetV sets V field to given value.

### HasV

`func (o *O11yMetricPoint) HasV() bool`

HasV returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



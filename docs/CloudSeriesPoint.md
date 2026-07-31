# CloudSeriesPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**T** | Pointer to **string** | bucket start, RFC3339 UTC | [optional] 
**V** | Pointer to **int32** | real invocation count in the bucket | [optional] 

## Methods

### NewCloudSeriesPoint

`func NewCloudSeriesPoint() *CloudSeriesPoint`

NewCloudSeriesPoint instantiates a new CloudSeriesPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSeriesPointWithDefaults

`func NewCloudSeriesPointWithDefaults() *CloudSeriesPoint`

NewCloudSeriesPointWithDefaults instantiates a new CloudSeriesPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetT

`func (o *CloudSeriesPoint) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *CloudSeriesPoint) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *CloudSeriesPoint) SetT(v string)`

SetT sets T field to given value.

### HasT

`func (o *CloudSeriesPoint) HasT() bool`

HasT returns a boolean if a field has been set.

### GetV

`func (o *CloudSeriesPoint) GetV() int32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *CloudSeriesPoint) GetVOk() (*int32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *CloudSeriesPoint) SetV(v int32)`

SetV sets V field to given value.

### HasV

`func (o *CloudSeriesPoint) HasV() bool`

HasV returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



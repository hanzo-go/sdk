# VectorDeletePointsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Points** | Pointer to [**[]VectorDeletePointsRequestPointsInner**](VectorDeletePointsRequestPointsInner.md) |  | [optional] 
**Filter** | Pointer to [**VectorFilter**](VectorFilter.md) |  | [optional] 

## Methods

### NewVectorDeletePointsRequest

`func NewVectorDeletePointsRequest() *VectorDeletePointsRequest`

NewVectorDeletePointsRequest instantiates a new VectorDeletePointsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorDeletePointsRequestWithDefaults

`func NewVectorDeletePointsRequestWithDefaults() *VectorDeletePointsRequest`

NewVectorDeletePointsRequestWithDefaults instantiates a new VectorDeletePointsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoints

`func (o *VectorDeletePointsRequest) GetPoints() []VectorDeletePointsRequestPointsInner`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *VectorDeletePointsRequest) GetPointsOk() (*[]VectorDeletePointsRequestPointsInner, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *VectorDeletePointsRequest) SetPoints(v []VectorDeletePointsRequestPointsInner)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *VectorDeletePointsRequest) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetFilter

`func (o *VectorDeletePointsRequest) GetFilter() VectorFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *VectorDeletePointsRequest) GetFilterOk() (*VectorFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *VectorDeletePointsRequest) SetFilter(v VectorFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *VectorDeletePointsRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



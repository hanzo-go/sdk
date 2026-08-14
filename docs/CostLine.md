# CostLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | the function the line is about | [optional] 
**Points** | Pointer to [**[]PointView**](PointView.md) | one point per bucket, oldest first | [optional] 

## Methods

### NewCostLine

`func NewCostLine() *CostLine`

NewCostLine instantiates a new CostLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostLineWithDefaults

`func NewCostLineWithDefaults() *CostLine`

NewCostLineWithDefaults instantiates a new CostLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CostLine) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CostLine) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CostLine) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CostLine) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPoints

`func (o *CostLine) GetPoints() []PointView`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *CostLine) GetPointsOk() (*[]PointView, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *CostLine) SetPoints(v []PointView)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *CostLine) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CloudSeriesLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | agent name | [optional] 
**Points** | Pointer to [**[]CloudSeriesPoint**](CloudSeriesPoint.md) |  | [optional] 

## Methods

### NewCloudSeriesLine

`func NewCloudSeriesLine() *CloudSeriesLine`

NewCloudSeriesLine instantiates a new CloudSeriesLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSeriesLineWithDefaults

`func NewCloudSeriesLineWithDefaults() *CloudSeriesLine`

NewCloudSeriesLineWithDefaults instantiates a new CloudSeriesLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CloudSeriesLine) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CloudSeriesLine) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CloudSeriesLine) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CloudSeriesLine) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPoints

`func (o *CloudSeriesLine) GetPoints() []CloudSeriesPoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *CloudSeriesLine) GetPointsOk() (*[]CloudSeriesPoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *CloudSeriesLine) SetPoints(v []CloudSeriesPoint)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *CloudSeriesLine) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



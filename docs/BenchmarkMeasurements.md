# BenchmarkMeasurements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | Pointer to **string** | By is the verified user who asked. | [optional] 
**Data** | Pointer to [**[]BenchmarkMeasurement**](BenchmarkMeasurement.md) | Data is one answer per benchmark, in the order the request named them. | [optional] 
**Model** | Pointer to **string** | Model is the model asked about. | [optional] 
**Org** | Pointer to **string** | Org is the caller&#39;s verified org. | [optional] 

## Methods

### NewBenchmarkMeasurements

`func NewBenchmarkMeasurements() *BenchmarkMeasurements`

NewBenchmarkMeasurements instantiates a new BenchmarkMeasurements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkMeasurementsWithDefaults

`func NewBenchmarkMeasurementsWithDefaults() *BenchmarkMeasurements`

NewBenchmarkMeasurementsWithDefaults instantiates a new BenchmarkMeasurements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *BenchmarkMeasurements) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *BenchmarkMeasurements) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *BenchmarkMeasurements) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *BenchmarkMeasurements) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetData

`func (o *BenchmarkMeasurements) GetData() []BenchmarkMeasurement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BenchmarkMeasurements) GetDataOk() (*[]BenchmarkMeasurement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BenchmarkMeasurements) SetData(v []BenchmarkMeasurement)`

SetData sets Data field to given value.

### HasData

`func (o *BenchmarkMeasurements) HasData() bool`

HasData returns a boolean if a field has been set.

### GetModel

`func (o *BenchmarkMeasurements) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BenchmarkMeasurements) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BenchmarkMeasurements) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *BenchmarkMeasurements) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOrg

`func (o *BenchmarkMeasurements) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *BenchmarkMeasurements) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *BenchmarkMeasurements) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *BenchmarkMeasurements) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



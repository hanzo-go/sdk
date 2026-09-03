# RiskOrgFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blind** | Pointer to **bool** | Blind is true when the dimension is present in no bucket at all: this organisation&#39;s surface does not carry it, and saying so is the difference between no risk and no data. | [optional] 
**Buckets** | Pointer to **int64** | Buckets is how many five-minute buckets of this organisation&#39;s surface were measured. | [optional] 
**Max** | Pointer to **float64** | Max is the largest value it reached in the window. | [optional] 
**Mean** | Pointer to **float64** | Mean is the dimension&#39;s average where it was present. | [optional] 
**Name** | Pointer to **string** | Name is the dimension as this API publishes it. | [optional] 
**Present** | Pointer to **int64** | Present is in how many of them the dimension carried a value at all. | [optional] 
**Source** | Pointer to **string** | Source names the plane it is rolled up from, so a dimension that reads zero everywhere traces to a plane the organisation does not use rather than to a defect. | [optional] 
**Unit** | Pointer to **string** | Unit is how to read the numbers below. | [optional] 

## Methods

### NewRiskOrgFeature

`func NewRiskOrgFeature() *RiskOrgFeature`

NewRiskOrgFeature instantiates a new RiskOrgFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskOrgFeatureWithDefaults

`func NewRiskOrgFeatureWithDefaults() *RiskOrgFeature`

NewRiskOrgFeatureWithDefaults instantiates a new RiskOrgFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlind

`func (o *RiskOrgFeature) GetBlind() bool`

GetBlind returns the Blind field if non-nil, zero value otherwise.

### GetBlindOk

`func (o *RiskOrgFeature) GetBlindOk() (*bool, bool)`

GetBlindOk returns a tuple with the Blind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlind

`func (o *RiskOrgFeature) SetBlind(v bool)`

SetBlind sets Blind field to given value.

### HasBlind

`func (o *RiskOrgFeature) HasBlind() bool`

HasBlind returns a boolean if a field has been set.

### GetBuckets

`func (o *RiskOrgFeature) GetBuckets() int64`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *RiskOrgFeature) GetBucketsOk() (*int64, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *RiskOrgFeature) SetBuckets(v int64)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *RiskOrgFeature) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetMax

`func (o *RiskOrgFeature) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *RiskOrgFeature) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *RiskOrgFeature) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *RiskOrgFeature) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMean

`func (o *RiskOrgFeature) GetMean() float64`

GetMean returns the Mean field if non-nil, zero value otherwise.

### GetMeanOk

`func (o *RiskOrgFeature) GetMeanOk() (*float64, bool)`

GetMeanOk returns a tuple with the Mean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMean

`func (o *RiskOrgFeature) SetMean(v float64)`

SetMean sets Mean field to given value.

### HasMean

`func (o *RiskOrgFeature) HasMean() bool`

HasMean returns a boolean if a field has been set.

### GetName

`func (o *RiskOrgFeature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskOrgFeature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskOrgFeature) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RiskOrgFeature) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPresent

`func (o *RiskOrgFeature) GetPresent() int64`

GetPresent returns the Present field if non-nil, zero value otherwise.

### GetPresentOk

`func (o *RiskOrgFeature) GetPresentOk() (*int64, bool)`

GetPresentOk returns a tuple with the Present field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresent

`func (o *RiskOrgFeature) SetPresent(v int64)`

SetPresent sets Present field to given value.

### HasPresent

`func (o *RiskOrgFeature) HasPresent() bool`

HasPresent returns a boolean if a field has been set.

### GetSource

`func (o *RiskOrgFeature) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RiskOrgFeature) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RiskOrgFeature) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RiskOrgFeature) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUnit

`func (o *RiskOrgFeature) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *RiskOrgFeature) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *RiskOrgFeature) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *RiskOrgFeature) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



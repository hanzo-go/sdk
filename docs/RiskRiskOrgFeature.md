# RiskRiskOrgFeature

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

### NewRiskRiskOrgFeature

`func NewRiskRiskOrgFeature() *RiskRiskOrgFeature`

NewRiskRiskOrgFeature instantiates a new RiskRiskOrgFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRiskOrgFeatureWithDefaults

`func NewRiskRiskOrgFeatureWithDefaults() *RiskRiskOrgFeature`

NewRiskRiskOrgFeatureWithDefaults instantiates a new RiskRiskOrgFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlind

`func (o *RiskRiskOrgFeature) GetBlind() bool`

GetBlind returns the Blind field if non-nil, zero value otherwise.

### GetBlindOk

`func (o *RiskRiskOrgFeature) GetBlindOk() (*bool, bool)`

GetBlindOk returns a tuple with the Blind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlind

`func (o *RiskRiskOrgFeature) SetBlind(v bool)`

SetBlind sets Blind field to given value.

### HasBlind

`func (o *RiskRiskOrgFeature) HasBlind() bool`

HasBlind returns a boolean if a field has been set.

### GetBuckets

`func (o *RiskRiskOrgFeature) GetBuckets() int64`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *RiskRiskOrgFeature) GetBucketsOk() (*int64, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *RiskRiskOrgFeature) SetBuckets(v int64)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *RiskRiskOrgFeature) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetMax

`func (o *RiskRiskOrgFeature) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *RiskRiskOrgFeature) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *RiskRiskOrgFeature) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *RiskRiskOrgFeature) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMean

`func (o *RiskRiskOrgFeature) GetMean() float64`

GetMean returns the Mean field if non-nil, zero value otherwise.

### GetMeanOk

`func (o *RiskRiskOrgFeature) GetMeanOk() (*float64, bool)`

GetMeanOk returns a tuple with the Mean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMean

`func (o *RiskRiskOrgFeature) SetMean(v float64)`

SetMean sets Mean field to given value.

### HasMean

`func (o *RiskRiskOrgFeature) HasMean() bool`

HasMean returns a boolean if a field has been set.

### GetName

`func (o *RiskRiskOrgFeature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskRiskOrgFeature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskRiskOrgFeature) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RiskRiskOrgFeature) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPresent

`func (o *RiskRiskOrgFeature) GetPresent() int64`

GetPresent returns the Present field if non-nil, zero value otherwise.

### GetPresentOk

`func (o *RiskRiskOrgFeature) GetPresentOk() (*int64, bool)`

GetPresentOk returns a tuple with the Present field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresent

`func (o *RiskRiskOrgFeature) SetPresent(v int64)`

SetPresent sets Present field to given value.

### HasPresent

`func (o *RiskRiskOrgFeature) HasPresent() bool`

HasPresent returns a boolean if a field has been set.

### GetSource

`func (o *RiskRiskOrgFeature) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RiskRiskOrgFeature) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RiskRiskOrgFeature) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RiskRiskOrgFeature) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUnit

`func (o *RiskRiskOrgFeature) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *RiskRiskOrgFeature) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *RiskRiskOrgFeature) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *RiskRiskOrgFeature) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



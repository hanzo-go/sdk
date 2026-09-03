# RiskValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Baseline** | Pointer to **float64** | Baseline is what Observed was measured against: this organisation&#39;s own history for this subject. | [optional] 
**Blind** | Pointer to **bool** | Blind marks a coordinate that could not be computed and took its neutral value. A model silently reading neutral for a dimension it never has data for is indistinguishable from one reading a genuine absence of risk. | [optional] 
**Feature** | Pointer to **string** | Feature is the dimension. | [optional] 
**Observed** | Pointer to **float64** | Observed is the raw number X was computed from, quoted so the coordinate reads back as a sentence rather than a bare ratio. | [optional] 
**Unit** | Pointer to **string** | Unit is how to read Observed. | [optional] 
**X** | Pointer to **float64** | X is the coordinate in the model space, always dimensionless. | [optional] 

## Methods

### NewRiskValue

`func NewRiskValue() *RiskValue`

NewRiskValue instantiates a new RiskValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskValueWithDefaults

`func NewRiskValueWithDefaults() *RiskValue`

NewRiskValueWithDefaults instantiates a new RiskValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseline

`func (o *RiskValue) GetBaseline() float64`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *RiskValue) GetBaselineOk() (*float64, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *RiskValue) SetBaseline(v float64)`

SetBaseline sets Baseline field to given value.

### HasBaseline

`func (o *RiskValue) HasBaseline() bool`

HasBaseline returns a boolean if a field has been set.

### GetBlind

`func (o *RiskValue) GetBlind() bool`

GetBlind returns the Blind field if non-nil, zero value otherwise.

### GetBlindOk

`func (o *RiskValue) GetBlindOk() (*bool, bool)`

GetBlindOk returns a tuple with the Blind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlind

`func (o *RiskValue) SetBlind(v bool)`

SetBlind sets Blind field to given value.

### HasBlind

`func (o *RiskValue) HasBlind() bool`

HasBlind returns a boolean if a field has been set.

### GetFeature

`func (o *RiskValue) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *RiskValue) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *RiskValue) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *RiskValue) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetObserved

`func (o *RiskValue) GetObserved() float64`

GetObserved returns the Observed field if non-nil, zero value otherwise.

### GetObservedOk

`func (o *RiskValue) GetObservedOk() (*float64, bool)`

GetObservedOk returns a tuple with the Observed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObserved

`func (o *RiskValue) SetObserved(v float64)`

SetObserved sets Observed field to given value.

### HasObserved

`func (o *RiskValue) HasObserved() bool`

HasObserved returns a boolean if a field has been set.

### GetUnit

`func (o *RiskValue) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *RiskValue) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *RiskValue) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *RiskValue) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetX

`func (o *RiskValue) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *RiskValue) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *RiskValue) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *RiskValue) HasX() bool`

HasX returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



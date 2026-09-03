# RiskModelFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blind** | Pointer to **int64** | Blind is how often this dimension took that neutral value for THIS organisation. | [optional] 
**Citation** | Pointer to **string** | Citation is where those words come from, so the claim is checkable rather than asserted. | [optional] 
**Indicator** | Pointer to **string** | Indicator is the supervisor&#39;s own words for the thing being looked for. | [optional] 
**Name** | Pointer to **string** | Name is the dimension. | [optional] 
**Neutral** | Pointer to **float64** | Neutral is the value the coordinate takes when the data cannot support it. | [optional] 
**Severity** | Pointer to **string** | Severity is how much weight an alert on it carries. | [optional] 
**Typology** | Pointer to **string** | Typology is the pattern this dimension detects. | [optional] 
**Unit** | Pointer to **string** | Unit is how to read the raw number, which is what turns a coordinate into a sentence an investigator can put in a file. | [optional] 
**Window** | Pointer to **string** | Window is the sliding aggregate it reads. | [optional] 

## Methods

### NewRiskModelFeature

`func NewRiskModelFeature() *RiskModelFeature`

NewRiskModelFeature instantiates a new RiskModelFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskModelFeatureWithDefaults

`func NewRiskModelFeatureWithDefaults() *RiskModelFeature`

NewRiskModelFeatureWithDefaults instantiates a new RiskModelFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlind

`func (o *RiskModelFeature) GetBlind() int64`

GetBlind returns the Blind field if non-nil, zero value otherwise.

### GetBlindOk

`func (o *RiskModelFeature) GetBlindOk() (*int64, bool)`

GetBlindOk returns a tuple with the Blind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlind

`func (o *RiskModelFeature) SetBlind(v int64)`

SetBlind sets Blind field to given value.

### HasBlind

`func (o *RiskModelFeature) HasBlind() bool`

HasBlind returns a boolean if a field has been set.

### GetCitation

`func (o *RiskModelFeature) GetCitation() string`

GetCitation returns the Citation field if non-nil, zero value otherwise.

### GetCitationOk

`func (o *RiskModelFeature) GetCitationOk() (*string, bool)`

GetCitationOk returns a tuple with the Citation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitation

`func (o *RiskModelFeature) SetCitation(v string)`

SetCitation sets Citation field to given value.

### HasCitation

`func (o *RiskModelFeature) HasCitation() bool`

HasCitation returns a boolean if a field has been set.

### GetIndicator

`func (o *RiskModelFeature) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *RiskModelFeature) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *RiskModelFeature) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.

### HasIndicator

`func (o *RiskModelFeature) HasIndicator() bool`

HasIndicator returns a boolean if a field has been set.

### GetName

`func (o *RiskModelFeature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskModelFeature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskModelFeature) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RiskModelFeature) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeutral

`func (o *RiskModelFeature) GetNeutral() float64`

GetNeutral returns the Neutral field if non-nil, zero value otherwise.

### GetNeutralOk

`func (o *RiskModelFeature) GetNeutralOk() (*float64, bool)`

GetNeutralOk returns a tuple with the Neutral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeutral

`func (o *RiskModelFeature) SetNeutral(v float64)`

SetNeutral sets Neutral field to given value.

### HasNeutral

`func (o *RiskModelFeature) HasNeutral() bool`

HasNeutral returns a boolean if a field has been set.

### GetSeverity

`func (o *RiskModelFeature) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RiskModelFeature) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RiskModelFeature) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *RiskModelFeature) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTypology

`func (o *RiskModelFeature) GetTypology() string`

GetTypology returns the Typology field if non-nil, zero value otherwise.

### GetTypologyOk

`func (o *RiskModelFeature) GetTypologyOk() (*string, bool)`

GetTypologyOk returns a tuple with the Typology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypology

`func (o *RiskModelFeature) SetTypology(v string)`

SetTypology sets Typology field to given value.

### HasTypology

`func (o *RiskModelFeature) HasTypology() bool`

HasTypology returns a boolean if a field has been set.

### GetUnit

`func (o *RiskModelFeature) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *RiskModelFeature) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *RiskModelFeature) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *RiskModelFeature) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetWindow

`func (o *RiskModelFeature) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *RiskModelFeature) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *RiskModelFeature) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *RiskModelFeature) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



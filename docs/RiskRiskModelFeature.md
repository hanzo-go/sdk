# RiskRiskModelFeature

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

### NewRiskRiskModelFeature

`func NewRiskRiskModelFeature() *RiskRiskModelFeature`

NewRiskRiskModelFeature instantiates a new RiskRiskModelFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRiskModelFeatureWithDefaults

`func NewRiskRiskModelFeatureWithDefaults() *RiskRiskModelFeature`

NewRiskRiskModelFeatureWithDefaults instantiates a new RiskRiskModelFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlind

`func (o *RiskRiskModelFeature) GetBlind() int64`

GetBlind returns the Blind field if non-nil, zero value otherwise.

### GetBlindOk

`func (o *RiskRiskModelFeature) GetBlindOk() (*int64, bool)`

GetBlindOk returns a tuple with the Blind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlind

`func (o *RiskRiskModelFeature) SetBlind(v int64)`

SetBlind sets Blind field to given value.

### HasBlind

`func (o *RiskRiskModelFeature) HasBlind() bool`

HasBlind returns a boolean if a field has been set.

### GetCitation

`func (o *RiskRiskModelFeature) GetCitation() string`

GetCitation returns the Citation field if non-nil, zero value otherwise.

### GetCitationOk

`func (o *RiskRiskModelFeature) GetCitationOk() (*string, bool)`

GetCitationOk returns a tuple with the Citation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitation

`func (o *RiskRiskModelFeature) SetCitation(v string)`

SetCitation sets Citation field to given value.

### HasCitation

`func (o *RiskRiskModelFeature) HasCitation() bool`

HasCitation returns a boolean if a field has been set.

### GetIndicator

`func (o *RiskRiskModelFeature) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *RiskRiskModelFeature) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *RiskRiskModelFeature) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.

### HasIndicator

`func (o *RiskRiskModelFeature) HasIndicator() bool`

HasIndicator returns a boolean if a field has been set.

### GetName

`func (o *RiskRiskModelFeature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskRiskModelFeature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskRiskModelFeature) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RiskRiskModelFeature) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeutral

`func (o *RiskRiskModelFeature) GetNeutral() float64`

GetNeutral returns the Neutral field if non-nil, zero value otherwise.

### GetNeutralOk

`func (o *RiskRiskModelFeature) GetNeutralOk() (*float64, bool)`

GetNeutralOk returns a tuple with the Neutral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeutral

`func (o *RiskRiskModelFeature) SetNeutral(v float64)`

SetNeutral sets Neutral field to given value.

### HasNeutral

`func (o *RiskRiskModelFeature) HasNeutral() bool`

HasNeutral returns a boolean if a field has been set.

### GetSeverity

`func (o *RiskRiskModelFeature) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RiskRiskModelFeature) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RiskRiskModelFeature) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *RiskRiskModelFeature) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTypology

`func (o *RiskRiskModelFeature) GetTypology() string`

GetTypology returns the Typology field if non-nil, zero value otherwise.

### GetTypologyOk

`func (o *RiskRiskModelFeature) GetTypologyOk() (*string, bool)`

GetTypologyOk returns a tuple with the Typology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypology

`func (o *RiskRiskModelFeature) SetTypology(v string)`

SetTypology sets Typology field to given value.

### HasTypology

`func (o *RiskRiskModelFeature) HasTypology() bool`

HasTypology returns a boolean if a field has been set.

### GetUnit

`func (o *RiskRiskModelFeature) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *RiskRiskModelFeature) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *RiskRiskModelFeature) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *RiskRiskModelFeature) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetWindow

`func (o *RiskRiskModelFeature) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *RiskRiskModelFeature) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *RiskRiskModelFeature) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *RiskRiskModelFeature) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RiskRiskCause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Baseline** | Pointer to **float64** | Baseline is the number it was measured against — always this organisation&#39;s own history, never a fixed limit and never another organisation&#39;s. | [optional] 
**Citation** | Pointer to **string** | Citation is where those words come from, so the claim is checkable rather than asserted — which is what a chargeback network or a regulator asks for. | [optional] 
**Feature** | Pointer to **string** | Feature is the dimension that contributed. | [optional] 
**Indicator** | Pointer to **string** | Indicator is the supervisor&#39;s own words for the thing being looked for. | [optional] 
**Observed** | Pointer to **float64** | Observed is the raw number the coordinate was computed from. | [optional] 
**Severity** | Pointer to **string** | Severity is how much weight this dimension carries. | [optional] 
**Share** | Pointer to **float64** | Share is this feature&#39;s part of the score, in [0,1]. Zero across every cause means no single feature accounts for the alert and the combination does; the causes are then ordered by how far each sits from unremarkable. | [optional] 
**Typology** | Pointer to **string** | Typology is the laundering or abuse pattern this dimension detects. | [optional] 
**Unit** | Pointer to **string** | Unit is how to read Observed, which is what turns a coordinate into a sentence. | [optional] 
**Without** | Pointer to **float64** | Without is the score the same event would have received with this coordinate at its neutral value — the counterfactual itself. | [optional] 

## Methods

### NewRiskRiskCause

`func NewRiskRiskCause() *RiskRiskCause`

NewRiskRiskCause instantiates a new RiskRiskCause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRiskCauseWithDefaults

`func NewRiskRiskCauseWithDefaults() *RiskRiskCause`

NewRiskRiskCauseWithDefaults instantiates a new RiskRiskCause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseline

`func (o *RiskRiskCause) GetBaseline() float64`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *RiskRiskCause) GetBaselineOk() (*float64, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *RiskRiskCause) SetBaseline(v float64)`

SetBaseline sets Baseline field to given value.

### HasBaseline

`func (o *RiskRiskCause) HasBaseline() bool`

HasBaseline returns a boolean if a field has been set.

### GetCitation

`func (o *RiskRiskCause) GetCitation() string`

GetCitation returns the Citation field if non-nil, zero value otherwise.

### GetCitationOk

`func (o *RiskRiskCause) GetCitationOk() (*string, bool)`

GetCitationOk returns a tuple with the Citation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitation

`func (o *RiskRiskCause) SetCitation(v string)`

SetCitation sets Citation field to given value.

### HasCitation

`func (o *RiskRiskCause) HasCitation() bool`

HasCitation returns a boolean if a field has been set.

### GetFeature

`func (o *RiskRiskCause) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *RiskRiskCause) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *RiskRiskCause) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *RiskRiskCause) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetIndicator

`func (o *RiskRiskCause) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *RiskRiskCause) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *RiskRiskCause) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.

### HasIndicator

`func (o *RiskRiskCause) HasIndicator() bool`

HasIndicator returns a boolean if a field has been set.

### GetObserved

`func (o *RiskRiskCause) GetObserved() float64`

GetObserved returns the Observed field if non-nil, zero value otherwise.

### GetObservedOk

`func (o *RiskRiskCause) GetObservedOk() (*float64, bool)`

GetObservedOk returns a tuple with the Observed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObserved

`func (o *RiskRiskCause) SetObserved(v float64)`

SetObserved sets Observed field to given value.

### HasObserved

`func (o *RiskRiskCause) HasObserved() bool`

HasObserved returns a boolean if a field has been set.

### GetSeverity

`func (o *RiskRiskCause) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RiskRiskCause) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RiskRiskCause) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *RiskRiskCause) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetShare

`func (o *RiskRiskCause) GetShare() float64`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *RiskRiskCause) GetShareOk() (*float64, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *RiskRiskCause) SetShare(v float64)`

SetShare sets Share field to given value.

### HasShare

`func (o *RiskRiskCause) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetTypology

`func (o *RiskRiskCause) GetTypology() string`

GetTypology returns the Typology field if non-nil, zero value otherwise.

### GetTypologyOk

`func (o *RiskRiskCause) GetTypologyOk() (*string, bool)`

GetTypologyOk returns a tuple with the Typology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypology

`func (o *RiskRiskCause) SetTypology(v string)`

SetTypology sets Typology field to given value.

### HasTypology

`func (o *RiskRiskCause) HasTypology() bool`

HasTypology returns a boolean if a field has been set.

### GetUnit

`func (o *RiskRiskCause) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *RiskRiskCause) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *RiskRiskCause) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *RiskRiskCause) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetWithout

`func (o *RiskRiskCause) GetWithout() float64`

GetWithout returns the Without field if non-nil, zero value otherwise.

### GetWithoutOk

`func (o *RiskRiskCause) GetWithoutOk() (*float64, bool)`

GetWithoutOk returns a tuple with the Without field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithout

`func (o *RiskRiskCause) SetWithout(v float64)`

SetWithout sets Without field to given value.

### HasWithout

`func (o *RiskRiskCause) HasWithout() bool`

HasWithout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



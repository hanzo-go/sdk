# RiskCause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Baseline** | Pointer to **float32** | Baseline is the number it was measured against — always this organisation&#39;s own history, never a fixed limit and never another organisation&#39;s. | [optional] 
**Citation** | Pointer to **string** | Citation is where those words come from, so the claim is checkable rather than asserted — which is what a chargeback network or a regulator asks for. | [optional] 
**Feature** | Pointer to **string** | Feature is the dimension that contributed. | [optional] 
**Indicator** | Pointer to **string** | Indicator is the supervisor&#39;s own words for the thing being looked for. | [optional] 
**Observed** | Pointer to **float32** | Observed is the raw number the coordinate was computed from. | [optional] 
**Severity** | Pointer to **string** | Severity is how much weight this dimension carries. | [optional] 
**Share** | Pointer to **float32** | Share is this feature&#39;s part of the score, in [0,1]. Zero across every cause means no single feature accounts for the alert and the combination does; the causes are then ordered by how far each sits from unremarkable. | [optional] 
**Typology** | Pointer to **string** | Typology is the laundering or abuse pattern this dimension detects. | [optional] 
**Unit** | Pointer to **string** | Unit is how to read Observed, which is what turns a coordinate into a sentence. | [optional] 
**Without** | Pointer to **float32** | Without is the score the same event would have received with this coordinate at its neutral value — the counterfactual itself. | [optional] 

## Methods

### NewRiskCause

`func NewRiskCause() *RiskCause`

NewRiskCause instantiates a new RiskCause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskCauseWithDefaults

`func NewRiskCauseWithDefaults() *RiskCause`

NewRiskCauseWithDefaults instantiates a new RiskCause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseline

`func (o *RiskCause) GetBaseline() float32`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *RiskCause) GetBaselineOk() (*float32, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *RiskCause) SetBaseline(v float32)`

SetBaseline sets Baseline field to given value.

### HasBaseline

`func (o *RiskCause) HasBaseline() bool`

HasBaseline returns a boolean if a field has been set.

### GetCitation

`func (o *RiskCause) GetCitation() string`

GetCitation returns the Citation field if non-nil, zero value otherwise.

### GetCitationOk

`func (o *RiskCause) GetCitationOk() (*string, bool)`

GetCitationOk returns a tuple with the Citation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitation

`func (o *RiskCause) SetCitation(v string)`

SetCitation sets Citation field to given value.

### HasCitation

`func (o *RiskCause) HasCitation() bool`

HasCitation returns a boolean if a field has been set.

### GetFeature

`func (o *RiskCause) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *RiskCause) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *RiskCause) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *RiskCause) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetIndicator

`func (o *RiskCause) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *RiskCause) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *RiskCause) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.

### HasIndicator

`func (o *RiskCause) HasIndicator() bool`

HasIndicator returns a boolean if a field has been set.

### GetObserved

`func (o *RiskCause) GetObserved() float32`

GetObserved returns the Observed field if non-nil, zero value otherwise.

### GetObservedOk

`func (o *RiskCause) GetObservedOk() (*float32, bool)`

GetObservedOk returns a tuple with the Observed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObserved

`func (o *RiskCause) SetObserved(v float32)`

SetObserved sets Observed field to given value.

### HasObserved

`func (o *RiskCause) HasObserved() bool`

HasObserved returns a boolean if a field has been set.

### GetSeverity

`func (o *RiskCause) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RiskCause) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RiskCause) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *RiskCause) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetShare

`func (o *RiskCause) GetShare() float32`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *RiskCause) GetShareOk() (*float32, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *RiskCause) SetShare(v float32)`

SetShare sets Share field to given value.

### HasShare

`func (o *RiskCause) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetTypology

`func (o *RiskCause) GetTypology() string`

GetTypology returns the Typology field if non-nil, zero value otherwise.

### GetTypologyOk

`func (o *RiskCause) GetTypologyOk() (*string, bool)`

GetTypologyOk returns a tuple with the Typology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypology

`func (o *RiskCause) SetTypology(v string)`

SetTypology sets Typology field to given value.

### HasTypology

`func (o *RiskCause) HasTypology() bool`

HasTypology returns a boolean if a field has been set.

### GetUnit

`func (o *RiskCause) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *RiskCause) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *RiskCause) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *RiskCause) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetWithout

`func (o *RiskCause) GetWithout() float32`

GetWithout returns the Without field if non-nil, zero value otherwise.

### GetWithoutOk

`func (o *RiskCause) GetWithoutOk() (*float32, bool)`

GetWithoutOk returns a tuple with the Without field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithout

`func (o *RiskCause) SetWithout(v float32)`

SetWithout sets Without field to given value.

### HasWithout

`func (o *RiskCause) HasWithout() bool`

HasWithout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



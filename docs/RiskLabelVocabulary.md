# RiskLabelVocabulary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dispositions** | Pointer to **[]string** | Dispositions is the closed set a write&#39;s &#x60;disposition&#x60; must be drawn from, published in full so a caller can validate a batch before filing it instead of discovering a refusal per member: \&quot;productive\&quot;, \&quot;unproductive\&quot;, and \&quot;\&quot; — the EMPTY STRING is a member and means an explicit unjudged, so a client that filters empties out of this list drops a third of the vocabulary and can never file \&quot;we looked and could not say\&quot;. They are the AML engine&#39;s own spelling, verbatim, which is what lets a replay there report against these values. | [optional] 
**Kinds** | Pointer to **[]string** | Kinds, Dispositions and Sources are the closed vocabularies. A value outside them is refused at the door. | [optional] 
**Precedence** | Pointer to **[]string** | Precedence is the sources in the order that resolves a conflict, strongest first. It is DERIVED from the same declaration the resolver reads, so the published order is the enforced order and cannot drift from it. | [optional] 
**Retention** | Pointer to **int32** | Retention is the platform floor in days: no tenant may dispose of a label younger than this, because a label can be the input to an adverse action. | [optional] 
**Rule** | Pointer to **[]string** | Rule states the tie-breaks below rank, in order, so a caller reading a contested resolution can reproduce it. | [optional] 

## Methods

### NewRiskLabelVocabulary

`func NewRiskLabelVocabulary() *RiskLabelVocabulary`

NewRiskLabelVocabulary instantiates a new RiskLabelVocabulary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskLabelVocabularyWithDefaults

`func NewRiskLabelVocabularyWithDefaults() *RiskLabelVocabulary`

NewRiskLabelVocabularyWithDefaults instantiates a new RiskLabelVocabulary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDispositions

`func (o *RiskLabelVocabulary) GetDispositions() []string`

GetDispositions returns the Dispositions field if non-nil, zero value otherwise.

### GetDispositionsOk

`func (o *RiskLabelVocabulary) GetDispositionsOk() (*[]string, bool)`

GetDispositionsOk returns a tuple with the Dispositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositions

`func (o *RiskLabelVocabulary) SetDispositions(v []string)`

SetDispositions sets Dispositions field to given value.

### HasDispositions

`func (o *RiskLabelVocabulary) HasDispositions() bool`

HasDispositions returns a boolean if a field has been set.

### GetKinds

`func (o *RiskLabelVocabulary) GetKinds() []string`

GetKinds returns the Kinds field if non-nil, zero value otherwise.

### GetKindsOk

`func (o *RiskLabelVocabulary) GetKindsOk() (*[]string, bool)`

GetKindsOk returns a tuple with the Kinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKinds

`func (o *RiskLabelVocabulary) SetKinds(v []string)`

SetKinds sets Kinds field to given value.

### HasKinds

`func (o *RiskLabelVocabulary) HasKinds() bool`

HasKinds returns a boolean if a field has been set.

### GetPrecedence

`func (o *RiskLabelVocabulary) GetPrecedence() []string`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *RiskLabelVocabulary) GetPrecedenceOk() (*[]string, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *RiskLabelVocabulary) SetPrecedence(v []string)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *RiskLabelVocabulary) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetRetention

`func (o *RiskLabelVocabulary) GetRetention() int32`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *RiskLabelVocabulary) GetRetentionOk() (*int32, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *RiskLabelVocabulary) SetRetention(v int32)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *RiskLabelVocabulary) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetRule

`func (o *RiskLabelVocabulary) GetRule() []string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *RiskLabelVocabulary) GetRuleOk() (*[]string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *RiskLabelVocabulary) SetRule(v []string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *RiskLabelVocabulary) HasRule() bool`

HasRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



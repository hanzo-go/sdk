# RiskSourceCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facts** | Pointer to **int64** | Facts is how many assertions this source filed; Won is how many judged events it was the assertion in force for. A source with many facts and few wins is one that is being outranked, which is worth knowing before concluding it is wired correctly. | [optional] 
**Source** | Pointer to **string** | Source is the asserter these two counts are for — chargeoff, dispute, case, refund, review or sample. There is one entry per source that either filed in the window or won in it, in precedence order, strongest first. A source no longer in the vocabulary still has rows and is reported after the known ones rather than dropped out of a total that is supposed to add up. | [optional] 
**Won** | Pointer to **int64** | Won is how many JUDGED events this source&#39;s assertion was the one IN FORCE for, at that event&#39;s own as-of — it beat every other visible claim under the precedence rule. Summed over the sources it is Judged. Read against Facts it is the ratio that matters: many filed and few won is a source being outranked, not a source that is broken, and one source winning nearly everything is a plane that looks labelled because one noisy filer dominates it. | [optional] 

## Methods

### NewRiskSourceCoverage

`func NewRiskSourceCoverage() *RiskSourceCoverage`

NewRiskSourceCoverage instantiates a new RiskSourceCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskSourceCoverageWithDefaults

`func NewRiskSourceCoverageWithDefaults() *RiskSourceCoverage`

NewRiskSourceCoverageWithDefaults instantiates a new RiskSourceCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacts

`func (o *RiskSourceCoverage) GetFacts() int64`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *RiskSourceCoverage) GetFactsOk() (*int64, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *RiskSourceCoverage) SetFacts(v int64)`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *RiskSourceCoverage) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetSource

`func (o *RiskSourceCoverage) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RiskSourceCoverage) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RiskSourceCoverage) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RiskSourceCoverage) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetWon

`func (o *RiskSourceCoverage) GetWon() int64`

GetWon returns the Won field if non-nil, zero value otherwise.

### GetWonOk

`func (o *RiskSourceCoverage) GetWonOk() (*int64, bool)`

GetWonOk returns a tuple with the Won field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWon

`func (o *RiskSourceCoverage) SetWon(v int64)`

SetWon sets Won field to given value.

### HasWon

`func (o *RiskSourceCoverage) HasWon() bool`

HasWon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



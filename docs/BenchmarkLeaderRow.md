# BenchmarkLeaderRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CiHigh** | Pointer to **float64** | CIHigh is the upper bound of that interval. Wilson rather than the normal approximation because the normal one produces bounds past 100 exactly where benchmark scores live — at 194/198 that is the top of the board, not a corner case. | [optional] 
**CiLow** | Pointer to **float64** | CILow and CIHigh are the 95% Wilson interval on Measured, in percent. They are what makes the score comparable: at n&#x3D;198 a 98% carries roughly ±2 points, so most differences at the top of a board are not distinguishable and a bare number implies a precision it does not have. Absent when there is no measurement. | [optional] 
**Claims** | Pointer to **int64** | Claims is how many independent claims the platform holds for this model on this benchmark. More than one means several sources reported it. | [optional] 
**Evidence** | Pointer to **[]string** | Evidence are the evidence uris the Published claim rests on (GET /v1/research/evidence?uri&#x3D;). | [optional] 
**From** | Pointer to **string** | From is the org whose claim Published was promoted from, when a SuperAdmin published another org&#39;s claim rather than writing one under org admin: \&quot;hanzo\&quot; on a row Hanzo measured and the platform published. | [optional] 
**Gap** | Pointer to **float64** | published − measured (the arena signal) | [optional] 
**Mean** | Pointer to **float64** | Mean is the unweighted average of every claim, which answers a different question from Published: what the field says on average, rather than what the vendor says about itself. With one claim the two are equal. | [optional] 
**Measured** | Pointer to **float64** | hanzo-measured accuracy % (nil if unrun) | [optional] 
**MeasuredAt** | Pointer to **time.Time** | MeasuredAt is when the run behind Measured was recorded. | [optional] 
**Model** | Pointer to **string** | the model this row scores | [optional] 
**N** | Pointer to **int64** | coverage — NEVER compare across different n | [optional] 
**Protocol** | Pointer to **string** | how the vendor scored their claim: single-attempt, pass@k or agentic | [optional] 
**Published** | Pointer to **float64** | the platform&#39;s selected claim % (nil if none) | [optional] 
**Run** | Pointer to **string** | Run names the measurement Measured came from, and MeasuredAt is when it ran. A score with no date is not a fact about a model, it is a fact about a model on a day — and models change, so the date is what makes the number checkable rather than merely quoted. | [optional] 
**Runs** | Pointer to **[]string** | Runs are the research run ids the Published claim rests on (GET /v1/research/runs?id&#x3D;). | [optional] 
**Spread** | Pointer to **float64** | Spread is the distance between the highest and lowest of them, nil when there is only one. It is the disagreement AMONG sources, which a single Published number cannot show — signal in the same way the published-minus-measured gap is. | [optional] 

## Methods

### NewBenchmarkLeaderRow

`func NewBenchmarkLeaderRow() *BenchmarkLeaderRow`

NewBenchmarkLeaderRow instantiates a new BenchmarkLeaderRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkLeaderRowWithDefaults

`func NewBenchmarkLeaderRowWithDefaults() *BenchmarkLeaderRow`

NewBenchmarkLeaderRowWithDefaults instantiates a new BenchmarkLeaderRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiHigh

`func (o *BenchmarkLeaderRow) GetCiHigh() float64`

GetCiHigh returns the CiHigh field if non-nil, zero value otherwise.

### GetCiHighOk

`func (o *BenchmarkLeaderRow) GetCiHighOk() (*float64, bool)`

GetCiHighOk returns a tuple with the CiHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiHigh

`func (o *BenchmarkLeaderRow) SetCiHigh(v float64)`

SetCiHigh sets CiHigh field to given value.

### HasCiHigh

`func (o *BenchmarkLeaderRow) HasCiHigh() bool`

HasCiHigh returns a boolean if a field has been set.

### GetCiLow

`func (o *BenchmarkLeaderRow) GetCiLow() float64`

GetCiLow returns the CiLow field if non-nil, zero value otherwise.

### GetCiLowOk

`func (o *BenchmarkLeaderRow) GetCiLowOk() (*float64, bool)`

GetCiLowOk returns a tuple with the CiLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiLow

`func (o *BenchmarkLeaderRow) SetCiLow(v float64)`

SetCiLow sets CiLow field to given value.

### HasCiLow

`func (o *BenchmarkLeaderRow) HasCiLow() bool`

HasCiLow returns a boolean if a field has been set.

### GetClaims

`func (o *BenchmarkLeaderRow) GetClaims() int64`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *BenchmarkLeaderRow) GetClaimsOk() (*int64, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *BenchmarkLeaderRow) SetClaims(v int64)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *BenchmarkLeaderRow) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetEvidence

`func (o *BenchmarkLeaderRow) GetEvidence() []string`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *BenchmarkLeaderRow) GetEvidenceOk() (*[]string, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *BenchmarkLeaderRow) SetEvidence(v []string)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *BenchmarkLeaderRow) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetFrom

`func (o *BenchmarkLeaderRow) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BenchmarkLeaderRow) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BenchmarkLeaderRow) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *BenchmarkLeaderRow) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGap

`func (o *BenchmarkLeaderRow) GetGap() float64`

GetGap returns the Gap field if non-nil, zero value otherwise.

### GetGapOk

`func (o *BenchmarkLeaderRow) GetGapOk() (*float64, bool)`

GetGapOk returns a tuple with the Gap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGap

`func (o *BenchmarkLeaderRow) SetGap(v float64)`

SetGap sets Gap field to given value.

### HasGap

`func (o *BenchmarkLeaderRow) HasGap() bool`

HasGap returns a boolean if a field has been set.

### GetMean

`func (o *BenchmarkLeaderRow) GetMean() float64`

GetMean returns the Mean field if non-nil, zero value otherwise.

### GetMeanOk

`func (o *BenchmarkLeaderRow) GetMeanOk() (*float64, bool)`

GetMeanOk returns a tuple with the Mean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMean

`func (o *BenchmarkLeaderRow) SetMean(v float64)`

SetMean sets Mean field to given value.

### HasMean

`func (o *BenchmarkLeaderRow) HasMean() bool`

HasMean returns a boolean if a field has been set.

### GetMeasured

`func (o *BenchmarkLeaderRow) GetMeasured() float64`

GetMeasured returns the Measured field if non-nil, zero value otherwise.

### GetMeasuredOk

`func (o *BenchmarkLeaderRow) GetMeasuredOk() (*float64, bool)`

GetMeasuredOk returns a tuple with the Measured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasured

`func (o *BenchmarkLeaderRow) SetMeasured(v float64)`

SetMeasured sets Measured field to given value.

### HasMeasured

`func (o *BenchmarkLeaderRow) HasMeasured() bool`

HasMeasured returns a boolean if a field has been set.

### GetMeasuredAt

`func (o *BenchmarkLeaderRow) GetMeasuredAt() time.Time`

GetMeasuredAt returns the MeasuredAt field if non-nil, zero value otherwise.

### GetMeasuredAtOk

`func (o *BenchmarkLeaderRow) GetMeasuredAtOk() (*time.Time, bool)`

GetMeasuredAtOk returns a tuple with the MeasuredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasuredAt

`func (o *BenchmarkLeaderRow) SetMeasuredAt(v time.Time)`

SetMeasuredAt sets MeasuredAt field to given value.

### HasMeasuredAt

`func (o *BenchmarkLeaderRow) HasMeasuredAt() bool`

HasMeasuredAt returns a boolean if a field has been set.

### GetModel

`func (o *BenchmarkLeaderRow) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BenchmarkLeaderRow) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BenchmarkLeaderRow) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *BenchmarkLeaderRow) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetN

`func (o *BenchmarkLeaderRow) GetN() int64`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *BenchmarkLeaderRow) GetNOk() (*int64, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *BenchmarkLeaderRow) SetN(v int64)`

SetN sets N field to given value.

### HasN

`func (o *BenchmarkLeaderRow) HasN() bool`

HasN returns a boolean if a field has been set.

### GetProtocol

`func (o *BenchmarkLeaderRow) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BenchmarkLeaderRow) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BenchmarkLeaderRow) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *BenchmarkLeaderRow) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPublished

`func (o *BenchmarkLeaderRow) GetPublished() float64`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *BenchmarkLeaderRow) GetPublishedOk() (*float64, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *BenchmarkLeaderRow) SetPublished(v float64)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *BenchmarkLeaderRow) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRun

`func (o *BenchmarkLeaderRow) GetRun() string`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *BenchmarkLeaderRow) GetRunOk() (*string, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *BenchmarkLeaderRow) SetRun(v string)`

SetRun sets Run field to given value.

### HasRun

`func (o *BenchmarkLeaderRow) HasRun() bool`

HasRun returns a boolean if a field has been set.

### GetRuns

`func (o *BenchmarkLeaderRow) GetRuns() []string`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *BenchmarkLeaderRow) GetRunsOk() (*[]string, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *BenchmarkLeaderRow) SetRuns(v []string)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *BenchmarkLeaderRow) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetSpread

`func (o *BenchmarkLeaderRow) GetSpread() float64`

GetSpread returns the Spread field if non-nil, zero value otherwise.

### GetSpreadOk

`func (o *BenchmarkLeaderRow) GetSpreadOk() (*float64, bool)`

GetSpreadOk returns a tuple with the Spread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpread

`func (o *BenchmarkLeaderRow) SetSpread(v float64)`

SetSpread sets Spread field to given value.

### HasSpread

`func (o *BenchmarkLeaderRow) HasSpread() bool`

HasSpread returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



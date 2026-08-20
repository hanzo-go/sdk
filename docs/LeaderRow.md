# LeaderRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CiHigh** | Pointer to **float32** | CIHigh is the upper bound of that interval. Wilson rather than the normal approximation because the normal one produces bounds past 100 exactly where benchmark scores live — at 194/198 that is the top of the board, not a corner case. | [optional] 
**CiLow** | Pointer to **float32** | CILow and CIHigh are the 95% Wilson interval on Measured, in percent. They are what makes the score comparable: at n&#x3D;198 a 98% carries roughly ±2 points, so most differences at the top of a board are not distinguishable and a bare number implies a precision it does not have. Absent when there is no measurement. | [optional] 
**Claims** | Pointer to **int32** | Claims is how many independent claims exist for this model on this benchmark. More than one means several sources reported it. | [optional] 
**Gap** | Pointer to **float32** | published − measured (the arena signal) | [optional] 
**Mean** | Pointer to **float32** | Mean is the unweighted average of every claim, which answers a different question from Published: what the field says on average, rather than what the vendor says about itself. With one claim the two are equal. | [optional] 
**Measured** | Pointer to **float32** | hanzo-measured accuracy % (nil if unrun) | [optional] 
**MeasuredAt** | Pointer to **time.Time** | MeasuredAt is when the run behind Measured was recorded. | [optional] 
**Model** | Pointer to **string** | the model this row scores | [optional] 
**N** | Pointer to **int32** | coverage — NEVER compare across different n | [optional] 
**Protocol** | Pointer to **string** | how the vendor scored their claim: single-attempt, pass@k or agentic | [optional] 
**Published** | Pointer to **float32** | provider-claimed % (nil if none) | [optional] 
**Run** | Pointer to **string** | Run names the measurement Measured came from, and MeasuredAt is when it ran. A score with no date is not a fact about a model, it is a fact about a model on a day — and models change, so the date is what makes the number checkable rather than merely quoted. | [optional] 
**Spread** | Pointer to **float32** | Spread is the distance between the highest and lowest of them, nil when there is only one. It is the disagreement AMONG sources, which a single Published number cannot show — signal in the same way the published-minus-measured gap is. | [optional] 

## Methods

### NewLeaderRow

`func NewLeaderRow() *LeaderRow`

NewLeaderRow instantiates a new LeaderRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderRowWithDefaults

`func NewLeaderRowWithDefaults() *LeaderRow`

NewLeaderRowWithDefaults instantiates a new LeaderRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiHigh

`func (o *LeaderRow) GetCiHigh() float32`

GetCiHigh returns the CiHigh field if non-nil, zero value otherwise.

### GetCiHighOk

`func (o *LeaderRow) GetCiHighOk() (*float32, bool)`

GetCiHighOk returns a tuple with the CiHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiHigh

`func (o *LeaderRow) SetCiHigh(v float32)`

SetCiHigh sets CiHigh field to given value.

### HasCiHigh

`func (o *LeaderRow) HasCiHigh() bool`

HasCiHigh returns a boolean if a field has been set.

### GetCiLow

`func (o *LeaderRow) GetCiLow() float32`

GetCiLow returns the CiLow field if non-nil, zero value otherwise.

### GetCiLowOk

`func (o *LeaderRow) GetCiLowOk() (*float32, bool)`

GetCiLowOk returns a tuple with the CiLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiLow

`func (o *LeaderRow) SetCiLow(v float32)`

SetCiLow sets CiLow field to given value.

### HasCiLow

`func (o *LeaderRow) HasCiLow() bool`

HasCiLow returns a boolean if a field has been set.

### GetClaims

`func (o *LeaderRow) GetClaims() int32`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *LeaderRow) GetClaimsOk() (*int32, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *LeaderRow) SetClaims(v int32)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *LeaderRow) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetGap

`func (o *LeaderRow) GetGap() float32`

GetGap returns the Gap field if non-nil, zero value otherwise.

### GetGapOk

`func (o *LeaderRow) GetGapOk() (*float32, bool)`

GetGapOk returns a tuple with the Gap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGap

`func (o *LeaderRow) SetGap(v float32)`

SetGap sets Gap field to given value.

### HasGap

`func (o *LeaderRow) HasGap() bool`

HasGap returns a boolean if a field has been set.

### GetMean

`func (o *LeaderRow) GetMean() float32`

GetMean returns the Mean field if non-nil, zero value otherwise.

### GetMeanOk

`func (o *LeaderRow) GetMeanOk() (*float32, bool)`

GetMeanOk returns a tuple with the Mean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMean

`func (o *LeaderRow) SetMean(v float32)`

SetMean sets Mean field to given value.

### HasMean

`func (o *LeaderRow) HasMean() bool`

HasMean returns a boolean if a field has been set.

### GetMeasured

`func (o *LeaderRow) GetMeasured() float32`

GetMeasured returns the Measured field if non-nil, zero value otherwise.

### GetMeasuredOk

`func (o *LeaderRow) GetMeasuredOk() (*float32, bool)`

GetMeasuredOk returns a tuple with the Measured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasured

`func (o *LeaderRow) SetMeasured(v float32)`

SetMeasured sets Measured field to given value.

### HasMeasured

`func (o *LeaderRow) HasMeasured() bool`

HasMeasured returns a boolean if a field has been set.

### GetMeasuredAt

`func (o *LeaderRow) GetMeasuredAt() time.Time`

GetMeasuredAt returns the MeasuredAt field if non-nil, zero value otherwise.

### GetMeasuredAtOk

`func (o *LeaderRow) GetMeasuredAtOk() (*time.Time, bool)`

GetMeasuredAtOk returns a tuple with the MeasuredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasuredAt

`func (o *LeaderRow) SetMeasuredAt(v time.Time)`

SetMeasuredAt sets MeasuredAt field to given value.

### HasMeasuredAt

`func (o *LeaderRow) HasMeasuredAt() bool`

HasMeasuredAt returns a boolean if a field has been set.

### GetModel

`func (o *LeaderRow) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *LeaderRow) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *LeaderRow) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *LeaderRow) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetN

`func (o *LeaderRow) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *LeaderRow) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *LeaderRow) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *LeaderRow) HasN() bool`

HasN returns a boolean if a field has been set.

### GetProtocol

`func (o *LeaderRow) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LeaderRow) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LeaderRow) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *LeaderRow) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPublished

`func (o *LeaderRow) GetPublished() float32`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *LeaderRow) GetPublishedOk() (*float32, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *LeaderRow) SetPublished(v float32)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *LeaderRow) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRun

`func (o *LeaderRow) GetRun() string`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *LeaderRow) GetRunOk() (*string, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *LeaderRow) SetRun(v string)`

SetRun sets Run field to given value.

### HasRun

`func (o *LeaderRow) HasRun() bool`

HasRun returns a boolean if a field has been set.

### GetSpread

`func (o *LeaderRow) GetSpread() float32`

GetSpread returns the Spread field if non-nil, zero value otherwise.

### GetSpreadOk

`func (o *LeaderRow) GetSpreadOk() (*float32, bool)`

GetSpreadOk returns a tuple with the Spread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpread

`func (o *LeaderRow) SetSpread(v float32)`

SetSpread sets Spread field to given value.

### HasSpread

`func (o *LeaderRow) HasSpread() bool`

HasSpread returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



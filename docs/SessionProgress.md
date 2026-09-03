# SessionProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to **string** | Activity is the one line saying what the run is doing right now (\&quot;running the reaper&#39;s tests\&quot;), up to 120 characters, in the model&#39;s words. Empty when nothing has estimated it yet. | [optional] 
**At** | Pointer to **string** | At is when this was determined, RFC 3339 in UTC to the second. Read it as the estimate&#39;s AGE: an estimate is not refreshed while nothing has happened, and a stale one beside a run that is still moving is itself worth seeing. Empty when nothing has estimated it yet. | [optional] 
**Estimated** | Pointer to **bool** | Estimated says a MODEL produced this, from the run&#39;s transcript, and it may be wrong. False means the session&#39;s own row said it: a finished run is 100% because it finished, not because anything guessed. Never treat a true here as a measurement — it is the reason to look, not the answer. | [optional] 
**Pct** | Pointer to **int64** | Pct is how much of the run is done, 0 to 100. THE KEY IS ABSENT when progress is indeterminate — a run nobody can estimate is not a run that has done nothing, and rendering the second for the first is the mistake this omission exists to make impossible. Read &#x60;phase&#x60; before reaching for it. | [optional] 
**Phase** | Pointer to **string** | Phase is what shape the run is in: running, blocked, done, error, or unknown when nothing has estimated it yet. blocked means the transcript shows the run waiting on something — an approval, a credential, an answer — which is the one state the running surface cannot report about itself. error only ever comes from the session&#39;s own terminal status. | [optional] 

## Methods

### NewSessionProgress

`func NewSessionProgress() *SessionProgress`

NewSessionProgress instantiates a new SessionProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionProgressWithDefaults

`func NewSessionProgressWithDefaults() *SessionProgress`

NewSessionProgressWithDefaults instantiates a new SessionProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *SessionProgress) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *SessionProgress) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *SessionProgress) SetActivity(v string)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *SessionProgress) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetAt

`func (o *SessionProgress) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *SessionProgress) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *SessionProgress) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *SessionProgress) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetEstimated

`func (o *SessionProgress) GetEstimated() bool`

GetEstimated returns the Estimated field if non-nil, zero value otherwise.

### GetEstimatedOk

`func (o *SessionProgress) GetEstimatedOk() (*bool, bool)`

GetEstimatedOk returns a tuple with the Estimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimated

`func (o *SessionProgress) SetEstimated(v bool)`

SetEstimated sets Estimated field to given value.

### HasEstimated

`func (o *SessionProgress) HasEstimated() bool`

HasEstimated returns a boolean if a field has been set.

### GetPct

`func (o *SessionProgress) GetPct() int64`

GetPct returns the Pct field if non-nil, zero value otherwise.

### GetPctOk

`func (o *SessionProgress) GetPctOk() (*int64, bool)`

GetPctOk returns a tuple with the Pct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPct

`func (o *SessionProgress) SetPct(v int64)`

SetPct sets Pct field to given value.

### HasPct

`func (o *SessionProgress) HasPct() bool`

HasPct returns a boolean if a field has been set.

### GetPhase

`func (o *SessionProgress) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *SessionProgress) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *SessionProgress) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *SessionProgress) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



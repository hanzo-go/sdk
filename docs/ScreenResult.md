# ScreenResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DraftReply** | Pointer to **string** | DraftReply is a suggested email reply for staff to edit and send. | [optional] 
**Error** | Pointer to **string** | Error says why a failed screen failed — no AI gateway configured, a gateway error, or a reply that carried no parseable JSON. Absent on success. | [optional] 
**Model** | Pointer to **string** | Model is the LLM the screen ran on. | [optional] 
**Score** | Pointer to **int32** | Score is the model&#39;s 0..100 fit score, clamped to that range. | [optional] 
**ScreenedAt** | Pointer to **int32** | ScreenedAt is the unix second the screen finished (0 while pending). | [optional] 
**Status** | Pointer to **string** | Status is the screen&#39;s state: pending | done | failed. | [optional] 
**SuggestedCredits** | Pointer to **int32** | SuggestedCredits is the recommended credit grant in USD, snapped to the nearest allowed rung: 0 | 5000 | 25000 | 50000 | 150000. | [optional] 
**Summary** | Pointer to **string** | Summary is the model&#39;s short assessment of the application. | [optional] 
**Tier1Backed** | Pointer to **string** | Tier1Backed is the model&#39;s read on tier-1 backing, normalized to \&quot;yes\&quot;, \&quot;no\&quot; or \&quot;unclear\&quot; (anything it cannot resolve reads \&quot;unclear\&quot;). | [optional] 

## Methods

### NewScreenResult

`func NewScreenResult() *ScreenResult`

NewScreenResult instantiates a new ScreenResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScreenResultWithDefaults

`func NewScreenResultWithDefaults() *ScreenResult`

NewScreenResultWithDefaults instantiates a new ScreenResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDraftReply

`func (o *ScreenResult) GetDraftReply() string`

GetDraftReply returns the DraftReply field if non-nil, zero value otherwise.

### GetDraftReplyOk

`func (o *ScreenResult) GetDraftReplyOk() (*string, bool)`

GetDraftReplyOk returns a tuple with the DraftReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftReply

`func (o *ScreenResult) SetDraftReply(v string)`

SetDraftReply sets DraftReply field to given value.

### HasDraftReply

`func (o *ScreenResult) HasDraftReply() bool`

HasDraftReply returns a boolean if a field has been set.

### GetError

`func (o *ScreenResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ScreenResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ScreenResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ScreenResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetModel

`func (o *ScreenResult) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ScreenResult) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ScreenResult) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ScreenResult) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetScore

`func (o *ScreenResult) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ScreenResult) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ScreenResult) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ScreenResult) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetScreenedAt

`func (o *ScreenResult) GetScreenedAt() int32`

GetScreenedAt returns the ScreenedAt field if non-nil, zero value otherwise.

### GetScreenedAtOk

`func (o *ScreenResult) GetScreenedAtOk() (*int32, bool)`

GetScreenedAtOk returns a tuple with the ScreenedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenedAt

`func (o *ScreenResult) SetScreenedAt(v int32)`

SetScreenedAt sets ScreenedAt field to given value.

### HasScreenedAt

`func (o *ScreenResult) HasScreenedAt() bool`

HasScreenedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ScreenResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScreenResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScreenResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScreenResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuggestedCredits

`func (o *ScreenResult) GetSuggestedCredits() int32`

GetSuggestedCredits returns the SuggestedCredits field if non-nil, zero value otherwise.

### GetSuggestedCreditsOk

`func (o *ScreenResult) GetSuggestedCreditsOk() (*int32, bool)`

GetSuggestedCreditsOk returns a tuple with the SuggestedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedCredits

`func (o *ScreenResult) SetSuggestedCredits(v int32)`

SetSuggestedCredits sets SuggestedCredits field to given value.

### HasSuggestedCredits

`func (o *ScreenResult) HasSuggestedCredits() bool`

HasSuggestedCredits returns a boolean if a field has been set.

### GetSummary

`func (o *ScreenResult) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ScreenResult) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ScreenResult) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ScreenResult) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTier1Backed

`func (o *ScreenResult) GetTier1Backed() string`

GetTier1Backed returns the Tier1Backed field if non-nil, zero value otherwise.

### GetTier1BackedOk

`func (o *ScreenResult) GetTier1BackedOk() (*string, bool)`

GetTier1BackedOk returns a tuple with the Tier1Backed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier1Backed

`func (o *ScreenResult) SetTier1Backed(v string)`

SetTier1Backed sets Tier1Backed field to given value.

### HasTier1Backed

`func (o *ScreenResult) HasTier1Backed() bool`

HasTier1Backed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



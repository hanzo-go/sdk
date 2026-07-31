# CloudScreenResult

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

### NewCloudScreenResult

`func NewCloudScreenResult() *CloudScreenResult`

NewCloudScreenResult instantiates a new CloudScreenResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudScreenResultWithDefaults

`func NewCloudScreenResultWithDefaults() *CloudScreenResult`

NewCloudScreenResultWithDefaults instantiates a new CloudScreenResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDraftReply

`func (o *CloudScreenResult) GetDraftReply() string`

GetDraftReply returns the DraftReply field if non-nil, zero value otherwise.

### GetDraftReplyOk

`func (o *CloudScreenResult) GetDraftReplyOk() (*string, bool)`

GetDraftReplyOk returns a tuple with the DraftReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftReply

`func (o *CloudScreenResult) SetDraftReply(v string)`

SetDraftReply sets DraftReply field to given value.

### HasDraftReply

`func (o *CloudScreenResult) HasDraftReply() bool`

HasDraftReply returns a boolean if a field has been set.

### GetError

`func (o *CloudScreenResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CloudScreenResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CloudScreenResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CloudScreenResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetModel

`func (o *CloudScreenResult) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CloudScreenResult) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CloudScreenResult) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CloudScreenResult) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetScore

`func (o *CloudScreenResult) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CloudScreenResult) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CloudScreenResult) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *CloudScreenResult) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetScreenedAt

`func (o *CloudScreenResult) GetScreenedAt() int32`

GetScreenedAt returns the ScreenedAt field if non-nil, zero value otherwise.

### GetScreenedAtOk

`func (o *CloudScreenResult) GetScreenedAtOk() (*int32, bool)`

GetScreenedAtOk returns a tuple with the ScreenedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenedAt

`func (o *CloudScreenResult) SetScreenedAt(v int32)`

SetScreenedAt sets ScreenedAt field to given value.

### HasScreenedAt

`func (o *CloudScreenResult) HasScreenedAt() bool`

HasScreenedAt returns a boolean if a field has been set.

### GetStatus

`func (o *CloudScreenResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudScreenResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudScreenResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudScreenResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuggestedCredits

`func (o *CloudScreenResult) GetSuggestedCredits() int32`

GetSuggestedCredits returns the SuggestedCredits field if non-nil, zero value otherwise.

### GetSuggestedCreditsOk

`func (o *CloudScreenResult) GetSuggestedCreditsOk() (*int32, bool)`

GetSuggestedCreditsOk returns a tuple with the SuggestedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedCredits

`func (o *CloudScreenResult) SetSuggestedCredits(v int32)`

SetSuggestedCredits sets SuggestedCredits field to given value.

### HasSuggestedCredits

`func (o *CloudScreenResult) HasSuggestedCredits() bool`

HasSuggestedCredits returns a boolean if a field has been set.

### GetSummary

`func (o *CloudScreenResult) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CloudScreenResult) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CloudScreenResult) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CloudScreenResult) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTier1Backed

`func (o *CloudScreenResult) GetTier1Backed() string`

GetTier1Backed returns the Tier1Backed field if non-nil, zero value otherwise.

### GetTier1BackedOk

`func (o *CloudScreenResult) GetTier1BackedOk() (*string, bool)`

GetTier1BackedOk returns a tuple with the Tier1Backed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier1Backed

`func (o *CloudScreenResult) SetTier1Backed(v string)`

SetTier1Backed sets Tier1Backed field to given value.

### HasTier1Backed

`func (o *CloudScreenResult) HasTier1Backed() bool`

HasTier1Backed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



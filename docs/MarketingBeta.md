# MarketingBeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]MarketingBetaAction**](MarketingBetaAction.md) | Actions are the actions that earned points, oldest first. | [optional] 
**Answers** | Pointer to [**[]MarketingBetaAnswer**](MarketingBetaAnswer.md) | Answers are the application&#39;s answers. | [optional] 
**Applied** | Pointer to **bool** | Applied is whether the account has applied: its record names a program. | [optional] 
**Betas** | Pointer to **[]string** | Betas are the ids of the unreleased products asked for. | [optional] 
**Github** | Pointer to **string** | GitHub is the GitHub login the account&#39;s stars are counted for. | [optional] 
**Lead** | Pointer to **string** | Lead is the reference the application was filed under as a lead. | [optional] 
**Of** | Pointer to **int64** | Of is how many people are on the waitlist, as the board counts them. | [optional] 
**Place** | Pointer to **int64** | Place is the record&#39;s place on the waitlist, 1 first; 0 with no record, or when the board did not answer. | [optional] 
**Points** | Pointer to **int64** | Points is the sum of what Actions earned. | [optional] 
**Program** | Pointer to **string** | Program is the program applied to: beta, research or partner. | [optional] 
**Requested** | Pointer to **map[string]string** | Requested is when each of Betas was asked for, RFC 3339 in UTC, by id. | [optional] 
**Score** | Pointer to **int64** | Score is what the waitlist ranks the record by: Points, the account&#39;s referrals and its token usage (GET /v1/leaderboard/beta says how). 0 with no record, or when the board did not answer. | [optional] 
**Since** | Pointer to **string** | Since is when the record began, as the CRM stamped it; empty with no record. | [optional] 

## Methods

### NewMarketingBeta

`func NewMarketingBeta() *MarketingBeta`

NewMarketingBeta instantiates a new MarketingBeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingBetaWithDefaults

`func NewMarketingBetaWithDefaults() *MarketingBeta`

NewMarketingBetaWithDefaults instantiates a new MarketingBeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *MarketingBeta) GetActions() []MarketingBetaAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *MarketingBeta) GetActionsOk() (*[]MarketingBetaAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *MarketingBeta) SetActions(v []MarketingBetaAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *MarketingBeta) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAnswers

`func (o *MarketingBeta) GetAnswers() []MarketingBetaAnswer`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *MarketingBeta) GetAnswersOk() (*[]MarketingBetaAnswer, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *MarketingBeta) SetAnswers(v []MarketingBetaAnswer)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *MarketingBeta) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetApplied

`func (o *MarketingBeta) GetApplied() bool`

GetApplied returns the Applied field if non-nil, zero value otherwise.

### GetAppliedOk

`func (o *MarketingBeta) GetAppliedOk() (*bool, bool)`

GetAppliedOk returns a tuple with the Applied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplied

`func (o *MarketingBeta) SetApplied(v bool)`

SetApplied sets Applied field to given value.

### HasApplied

`func (o *MarketingBeta) HasApplied() bool`

HasApplied returns a boolean if a field has been set.

### GetBetas

`func (o *MarketingBeta) GetBetas() []string`

GetBetas returns the Betas field if non-nil, zero value otherwise.

### GetBetasOk

`func (o *MarketingBeta) GetBetasOk() (*[]string, bool)`

GetBetasOk returns a tuple with the Betas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetas

`func (o *MarketingBeta) SetBetas(v []string)`

SetBetas sets Betas field to given value.

### HasBetas

`func (o *MarketingBeta) HasBetas() bool`

HasBetas returns a boolean if a field has been set.

### GetGithub

`func (o *MarketingBeta) GetGithub() string`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *MarketingBeta) GetGithubOk() (*string, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *MarketingBeta) SetGithub(v string)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *MarketingBeta) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetLead

`func (o *MarketingBeta) GetLead() string`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *MarketingBeta) GetLeadOk() (*string, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *MarketingBeta) SetLead(v string)`

SetLead sets Lead field to given value.

### HasLead

`func (o *MarketingBeta) HasLead() bool`

HasLead returns a boolean if a field has been set.

### GetOf

`func (o *MarketingBeta) GetOf() int64`

GetOf returns the Of field if non-nil, zero value otherwise.

### GetOfOk

`func (o *MarketingBeta) GetOfOk() (*int64, bool)`

GetOfOk returns a tuple with the Of field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOf

`func (o *MarketingBeta) SetOf(v int64)`

SetOf sets Of field to given value.

### HasOf

`func (o *MarketingBeta) HasOf() bool`

HasOf returns a boolean if a field has been set.

### GetPlace

`func (o *MarketingBeta) GetPlace() int64`

GetPlace returns the Place field if non-nil, zero value otherwise.

### GetPlaceOk

`func (o *MarketingBeta) GetPlaceOk() (*int64, bool)`

GetPlaceOk returns a tuple with the Place field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlace

`func (o *MarketingBeta) SetPlace(v int64)`

SetPlace sets Place field to given value.

### HasPlace

`func (o *MarketingBeta) HasPlace() bool`

HasPlace returns a boolean if a field has been set.

### GetPoints

`func (o *MarketingBeta) GetPoints() int64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *MarketingBeta) GetPointsOk() (*int64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *MarketingBeta) SetPoints(v int64)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *MarketingBeta) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetProgram

`func (o *MarketingBeta) GetProgram() string`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *MarketingBeta) GetProgramOk() (*string, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *MarketingBeta) SetProgram(v string)`

SetProgram sets Program field to given value.

### HasProgram

`func (o *MarketingBeta) HasProgram() bool`

HasProgram returns a boolean if a field has been set.

### GetRequested

`func (o *MarketingBeta) GetRequested() map[string]string`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *MarketingBeta) GetRequestedOk() (*map[string]string, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *MarketingBeta) SetRequested(v map[string]string)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *MarketingBeta) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### GetScore

`func (o *MarketingBeta) GetScore() int64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *MarketingBeta) GetScoreOk() (*int64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *MarketingBeta) SetScore(v int64)`

SetScore sets Score field to given value.

### HasScore

`func (o *MarketingBeta) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSince

`func (o *MarketingBeta) GetSince() string`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *MarketingBeta) GetSinceOk() (*string, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *MarketingBeta) SetSince(v string)`

SetSince sets Since field to given value.

### HasSince

`func (o *MarketingBeta) HasSince() bool`

HasSince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



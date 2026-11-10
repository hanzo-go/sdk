# LeaderboardBetaBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Of** | Pointer to **int64** | Of is how many people are on the waitlist: every account that applied, and every lead that applied before it had an account. | [optional] 
**Referrals** | Pointer to **bool** | Referrals is whether referrals were counted; false when the referral service did not answer, and no record&#39;s score has them then. | [optional] 
**Rows** | Pointer to [**[]LeaderboardBetaRow**](LeaderboardBetaRow.md) | Rows are the top of the waitlist, best first. | [optional] 
**Tokens** | Pointer to **bool** | Tokens is whether token usage was counted; false when the warehouse did not answer, and no record&#39;s score has it then. | [optional] 
**You** | Pointer to [**LeaderboardBetaRow**](LeaderboardBetaRow.md) | You is the caller&#39;s own row wherever it ranks; null when the caller holds no record. | [optional] 

## Methods

### NewLeaderboardBetaBoard

`func NewLeaderboardBetaBoard() *LeaderboardBetaBoard`

NewLeaderboardBetaBoard instantiates a new LeaderboardBetaBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardBetaBoardWithDefaults

`func NewLeaderboardBetaBoardWithDefaults() *LeaderboardBetaBoard`

NewLeaderboardBetaBoardWithDefaults instantiates a new LeaderboardBetaBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOf

`func (o *LeaderboardBetaBoard) GetOf() int64`

GetOf returns the Of field if non-nil, zero value otherwise.

### GetOfOk

`func (o *LeaderboardBetaBoard) GetOfOk() (*int64, bool)`

GetOfOk returns a tuple with the Of field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOf

`func (o *LeaderboardBetaBoard) SetOf(v int64)`

SetOf sets Of field to given value.

### HasOf

`func (o *LeaderboardBetaBoard) HasOf() bool`

HasOf returns a boolean if a field has been set.

### GetReferrals

`func (o *LeaderboardBetaBoard) GetReferrals() bool`

GetReferrals returns the Referrals field if non-nil, zero value otherwise.

### GetReferralsOk

`func (o *LeaderboardBetaBoard) GetReferralsOk() (*bool, bool)`

GetReferralsOk returns a tuple with the Referrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrals

`func (o *LeaderboardBetaBoard) SetReferrals(v bool)`

SetReferrals sets Referrals field to given value.

### HasReferrals

`func (o *LeaderboardBetaBoard) HasReferrals() bool`

HasReferrals returns a boolean if a field has been set.

### GetRows

`func (o *LeaderboardBetaBoard) GetRows() []LeaderboardBetaRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *LeaderboardBetaBoard) GetRowsOk() (*[]LeaderboardBetaRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *LeaderboardBetaBoard) SetRows(v []LeaderboardBetaRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *LeaderboardBetaBoard) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTokens

`func (o *LeaderboardBetaBoard) GetTokens() bool`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *LeaderboardBetaBoard) GetTokensOk() (*bool, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *LeaderboardBetaBoard) SetTokens(v bool)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *LeaderboardBetaBoard) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetYou

`func (o *LeaderboardBetaBoard) GetYou() LeaderboardBetaRow`

GetYou returns the You field if non-nil, zero value otherwise.

### GetYouOk

`func (o *LeaderboardBetaBoard) GetYouOk() (*LeaderboardBetaRow, bool)`

GetYouOk returns a tuple with the You field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYou

`func (o *LeaderboardBetaBoard) SetYou(v LeaderboardBetaRow)`

SetYou sets You field to given value.

### HasYou

`func (o *LeaderboardBetaBoard) HasYou() bool`

HasYou returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# LeaderboardPulseView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when no series could be read at all. | [optional] 
**Days** | Pointer to **[]string** | Days are the thirty days counted, oldest first, \&quot;2006-01-02\&quot; in UTC, ending today. | [optional] 
**Hours** | Pointer to **[]string** | Hours are the twenty-four hours counted, oldest first, RFC 3339 in UTC, ending in the current hour. | [optional] 
**Missing** | Pointer to **[]string** | Missing names every series not drawn: the ones nothing records (products, shares) and the ones whose source could not be read just now. | [optional] 
**Models** | Pointer to [**LeaderboardPulseSeries**](LeaderboardPulseSeries.md) | Models is the orgs that used a model for the first time, by day. | [optional] 
**Orgs** | Pointer to [**LeaderboardPulseSeries**](LeaderboardPulseSeries.md) | Orgs is the orgs that used anything. | [optional] 
**Referred** | Pointer to [**LeaderboardPulseSeries**](LeaderboardPulseSeries.md) | Referred is the orgs that signed up through a referral link. | [optional] 
**Signups** | Pointer to [**LeaderboardPulseSeries**](LeaderboardPulseSeries.md) | Signups is the orgs that signed up. | [optional] 
**Tokens** | Pointer to [**LeaderboardPulseSeries**](LeaderboardPulseSeries.md) | Tokens is the tokens used. | [optional] 
**Tries** | Pointer to [**[]LeaderboardPulseTally**](LeaderboardPulseTally.md) | Tries is the models orgs used for the first time in the window, most first. | [optional] 

## Methods

### NewLeaderboardPulseView

`func NewLeaderboardPulseView() *LeaderboardPulseView`

NewLeaderboardPulseView instantiates a new LeaderboardPulseView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardPulseViewWithDefaults

`func NewLeaderboardPulseViewWithDefaults() *LeaderboardPulseView`

NewLeaderboardPulseViewWithDefaults instantiates a new LeaderboardPulseView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *LeaderboardPulseView) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *LeaderboardPulseView) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *LeaderboardPulseView) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *LeaderboardPulseView) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetDays

`func (o *LeaderboardPulseView) GetDays() []string`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *LeaderboardPulseView) GetDaysOk() (*[]string, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *LeaderboardPulseView) SetDays(v []string)`

SetDays sets Days field to given value.

### HasDays

`func (o *LeaderboardPulseView) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetHours

`func (o *LeaderboardPulseView) GetHours() []string`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *LeaderboardPulseView) GetHoursOk() (*[]string, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *LeaderboardPulseView) SetHours(v []string)`

SetHours sets Hours field to given value.

### HasHours

`func (o *LeaderboardPulseView) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetMissing

`func (o *LeaderboardPulseView) GetMissing() []string`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *LeaderboardPulseView) GetMissingOk() (*[]string, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *LeaderboardPulseView) SetMissing(v []string)`

SetMissing sets Missing field to given value.

### HasMissing

`func (o *LeaderboardPulseView) HasMissing() bool`

HasMissing returns a boolean if a field has been set.

### GetModels

`func (o *LeaderboardPulseView) GetModels() LeaderboardPulseSeries`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *LeaderboardPulseView) GetModelsOk() (*LeaderboardPulseSeries, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *LeaderboardPulseView) SetModels(v LeaderboardPulseSeries)`

SetModels sets Models field to given value.

### HasModels

`func (o *LeaderboardPulseView) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetOrgs

`func (o *LeaderboardPulseView) GetOrgs() LeaderboardPulseSeries`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *LeaderboardPulseView) GetOrgsOk() (*LeaderboardPulseSeries, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *LeaderboardPulseView) SetOrgs(v LeaderboardPulseSeries)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *LeaderboardPulseView) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetReferred

`func (o *LeaderboardPulseView) GetReferred() LeaderboardPulseSeries`

GetReferred returns the Referred field if non-nil, zero value otherwise.

### GetReferredOk

`func (o *LeaderboardPulseView) GetReferredOk() (*LeaderboardPulseSeries, bool)`

GetReferredOk returns a tuple with the Referred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferred

`func (o *LeaderboardPulseView) SetReferred(v LeaderboardPulseSeries)`

SetReferred sets Referred field to given value.

### HasReferred

`func (o *LeaderboardPulseView) HasReferred() bool`

HasReferred returns a boolean if a field has been set.

### GetSignups

`func (o *LeaderboardPulseView) GetSignups() LeaderboardPulseSeries`

GetSignups returns the Signups field if non-nil, zero value otherwise.

### GetSignupsOk

`func (o *LeaderboardPulseView) GetSignupsOk() (*LeaderboardPulseSeries, bool)`

GetSignupsOk returns a tuple with the Signups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignups

`func (o *LeaderboardPulseView) SetSignups(v LeaderboardPulseSeries)`

SetSignups sets Signups field to given value.

### HasSignups

`func (o *LeaderboardPulseView) HasSignups() bool`

HasSignups returns a boolean if a field has been set.

### GetTokens

`func (o *LeaderboardPulseView) GetTokens() LeaderboardPulseSeries`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *LeaderboardPulseView) GetTokensOk() (*LeaderboardPulseSeries, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *LeaderboardPulseView) SetTokens(v LeaderboardPulseSeries)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *LeaderboardPulseView) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetTries

`func (o *LeaderboardPulseView) GetTries() []LeaderboardPulseTally`

GetTries returns the Tries field if non-nil, zero value otherwise.

### GetTriesOk

`func (o *LeaderboardPulseView) GetTriesOk() (*[]LeaderboardPulseTally, bool)`

GetTriesOk returns a tuple with the Tries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTries

`func (o *LeaderboardPulseView) SetTries(v []LeaderboardPulseTally)`

SetTries sets Tries field to given value.

### HasTries

`func (o *LeaderboardPulseView) HasTries() bool`

HasTries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



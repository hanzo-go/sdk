# TrafficView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blind** | Pointer to **int64** | Blind is how many requests in the window carried no identity to attribute them to — no validated credential and no client address. Non-zero on a public plane means the client address is not reaching this process (a TCP load balancer with no PROXY protocol in front of it, typically), so this scope&#39;s callers cannot be told apart and nothing can be held against them. | [optional] 
**Callers** | Pointer to [**[]TrafficCaller**](TrafficCaller.md) | Callers is the scope&#39;s busiest callers this window. A credentialed caller appears as a FINGERPRINT — a per-process one-way digest: enough to recognise the same caller across requests, never enough to reconstruct the credential. | [optional] 
**Ceiling** | Pointer to **int64** | Ceiling is the most callers this scope may hold at once. | [optional] 
**Denied** | Pointer to **int64** | Denied is how many of them the gate refused. | [optional] 
**Lanes** | Pointer to **map[string]int64** | Lanes is the request count per lane — agent, human, bot, unknown. This is the split that separates a customer&#39;s automation from a scraper. | [optional] 
**Mode** | Pointer to **string** | Mode is the abuse gate&#39;s posture for this scope: \&quot;shadow\&quot; records the scorer&#39;s action without enforcing it, \&quot;live\&quot; enforces it. | [optional] 
**Org** | Pointer to **string** | Org is the scope this view was taken for — the validated principal&#39;s own, never a value the caller supplied. Empty names the anonymous lane, the one scope that has no tenant. | [optional] 
**Refused** | Pointer to **int64** | Refused is how many callers this scope&#39;s ceilings turned away in the window. | [optional] 
**Requests** | Pointer to **int64** | Requests is how many requests this scope made in the window. | [optional] 
**Screens** | Pointer to **int64** | Screens is how many of them were put to the scorer — the billable unit of the risk product. Counted from the first request, whatever the SKU costs. | [optional] 
**Strain** | Pointer to **string** | Strain is what this scope&#39;s ceilings are doing: \&quot;clear\&quot; below them, \&quot;full\&quot; at them, \&quot;refuse\&quot; once a caller has been turned away inside this window — which means that caller is UNMEASURED and the numbers here are a sample rather than a census. It is reported rather than logged because the alternative — a bound that degrades a scope silently — is the failure this design exists to rule out. No other scope can move it. | [optional] 
**Tracked** | Pointer to **int64** | Tracked is how many callers this scope holds state for right now, and Ceiling is the most it may hold. Tracked &#x3D;&#x3D; Ceiling is the fact a bound that binds cannot hide. | [optional] 
**Unscored** | Pointer to **int64** | Unscored is how many of those screens got NO answer — the scorer was absent, stuck, slow, erroring or silent. An unanswered screen allows ordinary traffic, so this is the number that separates \&quot;a quiet day\&quot; from \&quot;the judge stopped answering and nothing said so\&quot;. | [optional] 
**WindowSec** | Pointer to **int64** | WindowSec is the span the counts cover, in seconds. | [optional] 

## Methods

### NewTrafficView

`func NewTrafficView() *TrafficView`

NewTrafficView instantiates a new TrafficView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficViewWithDefaults

`func NewTrafficViewWithDefaults() *TrafficView`

NewTrafficViewWithDefaults instantiates a new TrafficView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlind

`func (o *TrafficView) GetBlind() int64`

GetBlind returns the Blind field if non-nil, zero value otherwise.

### GetBlindOk

`func (o *TrafficView) GetBlindOk() (*int64, bool)`

GetBlindOk returns a tuple with the Blind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlind

`func (o *TrafficView) SetBlind(v int64)`

SetBlind sets Blind field to given value.

### HasBlind

`func (o *TrafficView) HasBlind() bool`

HasBlind returns a boolean if a field has been set.

### GetCallers

`func (o *TrafficView) GetCallers() []TrafficCaller`

GetCallers returns the Callers field if non-nil, zero value otherwise.

### GetCallersOk

`func (o *TrafficView) GetCallersOk() (*[]TrafficCaller, bool)`

GetCallersOk returns a tuple with the Callers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallers

`func (o *TrafficView) SetCallers(v []TrafficCaller)`

SetCallers sets Callers field to given value.

### HasCallers

`func (o *TrafficView) HasCallers() bool`

HasCallers returns a boolean if a field has been set.

### GetCeiling

`func (o *TrafficView) GetCeiling() int64`

GetCeiling returns the Ceiling field if non-nil, zero value otherwise.

### GetCeilingOk

`func (o *TrafficView) GetCeilingOk() (*int64, bool)`

GetCeilingOk returns a tuple with the Ceiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeiling

`func (o *TrafficView) SetCeiling(v int64)`

SetCeiling sets Ceiling field to given value.

### HasCeiling

`func (o *TrafficView) HasCeiling() bool`

HasCeiling returns a boolean if a field has been set.

### GetDenied

`func (o *TrafficView) GetDenied() int64`

GetDenied returns the Denied field if non-nil, zero value otherwise.

### GetDeniedOk

`func (o *TrafficView) GetDeniedOk() (*int64, bool)`

GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenied

`func (o *TrafficView) SetDenied(v int64)`

SetDenied sets Denied field to given value.

### HasDenied

`func (o *TrafficView) HasDenied() bool`

HasDenied returns a boolean if a field has been set.

### GetLanes

`func (o *TrafficView) GetLanes() map[string]int64`

GetLanes returns the Lanes field if non-nil, zero value otherwise.

### GetLanesOk

`func (o *TrafficView) GetLanesOk() (*map[string]int64, bool)`

GetLanesOk returns a tuple with the Lanes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanes

`func (o *TrafficView) SetLanes(v map[string]int64)`

SetLanes sets Lanes field to given value.

### HasLanes

`func (o *TrafficView) HasLanes() bool`

HasLanes returns a boolean if a field has been set.

### GetMode

`func (o *TrafficView) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TrafficView) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TrafficView) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *TrafficView) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOrg

`func (o *TrafficView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *TrafficView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *TrafficView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *TrafficView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRefused

`func (o *TrafficView) GetRefused() int64`

GetRefused returns the Refused field if non-nil, zero value otherwise.

### GetRefusedOk

`func (o *TrafficView) GetRefusedOk() (*int64, bool)`

GetRefusedOk returns a tuple with the Refused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefused

`func (o *TrafficView) SetRefused(v int64)`

SetRefused sets Refused field to given value.

### HasRefused

`func (o *TrafficView) HasRefused() bool`

HasRefused returns a boolean if a field has been set.

### GetRequests

`func (o *TrafficView) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *TrafficView) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *TrafficView) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *TrafficView) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetScreens

`func (o *TrafficView) GetScreens() int64`

GetScreens returns the Screens field if non-nil, zero value otherwise.

### GetScreensOk

`func (o *TrafficView) GetScreensOk() (*int64, bool)`

GetScreensOk returns a tuple with the Screens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreens

`func (o *TrafficView) SetScreens(v int64)`

SetScreens sets Screens field to given value.

### HasScreens

`func (o *TrafficView) HasScreens() bool`

HasScreens returns a boolean if a field has been set.

### GetStrain

`func (o *TrafficView) GetStrain() string`

GetStrain returns the Strain field if non-nil, zero value otherwise.

### GetStrainOk

`func (o *TrafficView) GetStrainOk() (*string, bool)`

GetStrainOk returns a tuple with the Strain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrain

`func (o *TrafficView) SetStrain(v string)`

SetStrain sets Strain field to given value.

### HasStrain

`func (o *TrafficView) HasStrain() bool`

HasStrain returns a boolean if a field has been set.

### GetTracked

`func (o *TrafficView) GetTracked() int64`

GetTracked returns the Tracked field if non-nil, zero value otherwise.

### GetTrackedOk

`func (o *TrafficView) GetTrackedOk() (*int64, bool)`

GetTrackedOk returns a tuple with the Tracked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracked

`func (o *TrafficView) SetTracked(v int64)`

SetTracked sets Tracked field to given value.

### HasTracked

`func (o *TrafficView) HasTracked() bool`

HasTracked returns a boolean if a field has been set.

### GetUnscored

`func (o *TrafficView) GetUnscored() int64`

GetUnscored returns the Unscored field if non-nil, zero value otherwise.

### GetUnscoredOk

`func (o *TrafficView) GetUnscoredOk() (*int64, bool)`

GetUnscoredOk returns a tuple with the Unscored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnscored

`func (o *TrafficView) SetUnscored(v int64)`

SetUnscored sets Unscored field to given value.

### HasUnscored

`func (o *TrafficView) HasUnscored() bool`

HasUnscored returns a boolean if a field has been set.

### GetWindowSec

`func (o *TrafficView) GetWindowSec() int64`

GetWindowSec returns the WindowSec field if non-nil, zero value otherwise.

### GetWindowSecOk

`func (o *TrafficView) GetWindowSecOk() (*int64, bool)`

GetWindowSecOk returns a tuple with the WindowSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSec

`func (o *TrafficView) SetWindowSec(v int64)`

SetWindowSec sets WindowSec field to given value.

### HasWindowSec

`func (o *TrafficView) HasWindowSec() bool`

HasWindowSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



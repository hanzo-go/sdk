# GatewayTrafficView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blind** | Pointer to **int64** | Blind is how many requests in the window carried no identity to attribute them to — no validated credential and no client address. Non-zero on a public plane means the client address is not reaching this process (a TCP load balancer with no PROXY protocol in front of it, typically), so this scope&#39;s callers cannot be told apart and nothing can be held against them. | [optional] 
**Callers** | Pointer to [**[]GatewayTrafficCaller**](GatewayTrafficCaller.md) | Callers is the scope&#39;s busiest callers this window. A credentialed caller appears as a FINGERPRINT — a per-process one-way digest: enough to recognise the same caller across requests, never enough to reconstruct the credential. | [optional] 
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

### NewGatewayTrafficView

`func NewGatewayTrafficView() *GatewayTrafficView`

NewGatewayTrafficView instantiates a new GatewayTrafficView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTrafficViewWithDefaults

`func NewGatewayTrafficViewWithDefaults() *GatewayTrafficView`

NewGatewayTrafficViewWithDefaults instantiates a new GatewayTrafficView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlind

`func (o *GatewayTrafficView) GetBlind() int64`

GetBlind returns the Blind field if non-nil, zero value otherwise.

### GetBlindOk

`func (o *GatewayTrafficView) GetBlindOk() (*int64, bool)`

GetBlindOk returns a tuple with the Blind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlind

`func (o *GatewayTrafficView) SetBlind(v int64)`

SetBlind sets Blind field to given value.

### HasBlind

`func (o *GatewayTrafficView) HasBlind() bool`

HasBlind returns a boolean if a field has been set.

### GetCallers

`func (o *GatewayTrafficView) GetCallers() []GatewayTrafficCaller`

GetCallers returns the Callers field if non-nil, zero value otherwise.

### GetCallersOk

`func (o *GatewayTrafficView) GetCallersOk() (*[]GatewayTrafficCaller, bool)`

GetCallersOk returns a tuple with the Callers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallers

`func (o *GatewayTrafficView) SetCallers(v []GatewayTrafficCaller)`

SetCallers sets Callers field to given value.

### HasCallers

`func (o *GatewayTrafficView) HasCallers() bool`

HasCallers returns a boolean if a field has been set.

### GetCeiling

`func (o *GatewayTrafficView) GetCeiling() int64`

GetCeiling returns the Ceiling field if non-nil, zero value otherwise.

### GetCeilingOk

`func (o *GatewayTrafficView) GetCeilingOk() (*int64, bool)`

GetCeilingOk returns a tuple with the Ceiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeiling

`func (o *GatewayTrafficView) SetCeiling(v int64)`

SetCeiling sets Ceiling field to given value.

### HasCeiling

`func (o *GatewayTrafficView) HasCeiling() bool`

HasCeiling returns a boolean if a field has been set.

### GetDenied

`func (o *GatewayTrafficView) GetDenied() int64`

GetDenied returns the Denied field if non-nil, zero value otherwise.

### GetDeniedOk

`func (o *GatewayTrafficView) GetDeniedOk() (*int64, bool)`

GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenied

`func (o *GatewayTrafficView) SetDenied(v int64)`

SetDenied sets Denied field to given value.

### HasDenied

`func (o *GatewayTrafficView) HasDenied() bool`

HasDenied returns a boolean if a field has been set.

### GetLanes

`func (o *GatewayTrafficView) GetLanes() map[string]int64`

GetLanes returns the Lanes field if non-nil, zero value otherwise.

### GetLanesOk

`func (o *GatewayTrafficView) GetLanesOk() (*map[string]int64, bool)`

GetLanesOk returns a tuple with the Lanes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanes

`func (o *GatewayTrafficView) SetLanes(v map[string]int64)`

SetLanes sets Lanes field to given value.

### HasLanes

`func (o *GatewayTrafficView) HasLanes() bool`

HasLanes returns a boolean if a field has been set.

### GetMode

`func (o *GatewayTrafficView) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GatewayTrafficView) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GatewayTrafficView) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GatewayTrafficView) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOrg

`func (o *GatewayTrafficView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *GatewayTrafficView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *GatewayTrafficView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *GatewayTrafficView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRefused

`func (o *GatewayTrafficView) GetRefused() int64`

GetRefused returns the Refused field if non-nil, zero value otherwise.

### GetRefusedOk

`func (o *GatewayTrafficView) GetRefusedOk() (*int64, bool)`

GetRefusedOk returns a tuple with the Refused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefused

`func (o *GatewayTrafficView) SetRefused(v int64)`

SetRefused sets Refused field to given value.

### HasRefused

`func (o *GatewayTrafficView) HasRefused() bool`

HasRefused returns a boolean if a field has been set.

### GetRequests

`func (o *GatewayTrafficView) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *GatewayTrafficView) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *GatewayTrafficView) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *GatewayTrafficView) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetScreens

`func (o *GatewayTrafficView) GetScreens() int64`

GetScreens returns the Screens field if non-nil, zero value otherwise.

### GetScreensOk

`func (o *GatewayTrafficView) GetScreensOk() (*int64, bool)`

GetScreensOk returns a tuple with the Screens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreens

`func (o *GatewayTrafficView) SetScreens(v int64)`

SetScreens sets Screens field to given value.

### HasScreens

`func (o *GatewayTrafficView) HasScreens() bool`

HasScreens returns a boolean if a field has been set.

### GetStrain

`func (o *GatewayTrafficView) GetStrain() string`

GetStrain returns the Strain field if non-nil, zero value otherwise.

### GetStrainOk

`func (o *GatewayTrafficView) GetStrainOk() (*string, bool)`

GetStrainOk returns a tuple with the Strain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrain

`func (o *GatewayTrafficView) SetStrain(v string)`

SetStrain sets Strain field to given value.

### HasStrain

`func (o *GatewayTrafficView) HasStrain() bool`

HasStrain returns a boolean if a field has been set.

### GetTracked

`func (o *GatewayTrafficView) GetTracked() int64`

GetTracked returns the Tracked field if non-nil, zero value otherwise.

### GetTrackedOk

`func (o *GatewayTrafficView) GetTrackedOk() (*int64, bool)`

GetTrackedOk returns a tuple with the Tracked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracked

`func (o *GatewayTrafficView) SetTracked(v int64)`

SetTracked sets Tracked field to given value.

### HasTracked

`func (o *GatewayTrafficView) HasTracked() bool`

HasTracked returns a boolean if a field has been set.

### GetUnscored

`func (o *GatewayTrafficView) GetUnscored() int64`

GetUnscored returns the Unscored field if non-nil, zero value otherwise.

### GetUnscoredOk

`func (o *GatewayTrafficView) GetUnscoredOk() (*int64, bool)`

GetUnscoredOk returns a tuple with the Unscored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnscored

`func (o *GatewayTrafficView) SetUnscored(v int64)`

SetUnscored sets Unscored field to given value.

### HasUnscored

`func (o *GatewayTrafficView) HasUnscored() bool`

HasUnscored returns a boolean if a field has been set.

### GetWindowSec

`func (o *GatewayTrafficView) GetWindowSec() int64`

GetWindowSec returns the WindowSec field if non-nil, zero value otherwise.

### GetWindowSecOk

`func (o *GatewayTrafficView) GetWindowSecOk() (*int64, bool)`

GetWindowSecOk returns a tuple with the WindowSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSec

`func (o *GatewayTrafficView) SetWindowSec(v int64)`

SetWindowSec sets WindowSec field to given value.

### HasWindowSec

`func (o *GatewayTrafficView) HasWindowSec() bool`

HasWindowSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



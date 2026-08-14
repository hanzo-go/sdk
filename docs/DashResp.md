# DashResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the linked account that was asked about, when one was named. | [optional] 
**Available** | Pointer to **bool** | Available is false when the warehouse could not be read. That means \&quot;no answer\&quot;, NOT \&quot;no usage\&quot; — the two lists below are then empty for a reason. | [optional] 
**Current** | Pointer to [**[]UsageWindowView**](UsageWindowView.md) | Current is the newest window instance of each lane — the dash headline. | [optional] 
**From** | Pointer to **string** | From is the inclusive start of that window, RFC3339 UTC. | [optional] 
**Provider** | Pointer to **string** | Provider is the upstream that was asked about, echoed back. | [optional] 
**Range** | Pointer to **string** | Range is the window that was served: 1h, 24h, 7d or 30d. | [optional] 
**Scope** | Pointer to **string** | Scope says whose rows these are: the caller&#39;s own linked accounts. | [optional] 
**Source** | Pointer to **string** | Source names the meter of record — the provider&#39;s own login, not Hanzo. | [optional] 
**To** | Pointer to **string** | To is the exclusive end of that window, RFC3339 UTC. | [optional] 
**Windows** | Pointer to [**[]UsageWindowView**](UsageWindowView.md) | Windows is every instance in range, newest first — the history behind it. | [optional] 

## Methods

### NewDashResp

`func NewDashResp() *DashResp`

NewDashResp instantiates a new DashResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashRespWithDefaults

`func NewDashRespWithDefaults() *DashResp`

NewDashRespWithDefaults instantiates a new DashResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *DashResp) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *DashResp) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *DashResp) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *DashResp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAvailable

`func (o *DashResp) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *DashResp) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *DashResp) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *DashResp) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCurrent

`func (o *DashResp) GetCurrent() []UsageWindowView`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *DashResp) GetCurrentOk() (*[]UsageWindowView, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *DashResp) SetCurrent(v []UsageWindowView)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *DashResp) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetFrom

`func (o *DashResp) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *DashResp) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *DashResp) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *DashResp) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetProvider

`func (o *DashResp) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DashResp) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DashResp) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *DashResp) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRange

`func (o *DashResp) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *DashResp) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *DashResp) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *DashResp) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScope

`func (o *DashResp) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DashResp) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DashResp) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *DashResp) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSource

`func (o *DashResp) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DashResp) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DashResp) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DashResp) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTo

`func (o *DashResp) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *DashResp) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *DashResp) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *DashResp) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetWindows

`func (o *DashResp) GetWindows() []UsageWindowView`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *DashResp) GetWindowsOk() (*[]UsageWindowView, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *DashResp) SetWindows(v []UsageWindowView)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *DashResp) HasWindows() bool`

HasWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



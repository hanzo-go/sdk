# BoardResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the account the series narrows to, when one was named. | [optional] 
**Available** | Pointer to **bool** | Available reports whether the warehouse answered; false is an honest \&quot;we have no data\&quot;, NOT zero usage. | [optional] 
**Current** | Pointer to [**[]ReadingView**](ReadingView.md) | Current is the live state of each lane — the dash headline. | [optional] 
**From** | Pointer to **string** | From is when the resolved window opens, RFC 3339 UTC. | [optional] 
**Provider** | Pointer to **string** | Provider is the provider whose meter answered. | [optional] 
**Range** | Pointer to **string** | Range is the resolved period label. | [optional] 
**Scope** | Pointer to **string** | Scope is always \&quot;user\&quot;: the caller&#39;s own linked accounts. | [optional] 
**Source** | Pointer to **string** | Source is always \&quot;account\&quot;: the provider&#39;s own meter, not a Hanzo charge. | [optional] 
**To** | Pointer to **string** | To is where it closes, EXCLUSIVE, RFC 3339 UTC — the instant the read was served, so the window walks forward with the clock and two reads a minute apart do not cover the same period. | [optional] 
**Windows** | Pointer to [**[]ReadingView**](ReadingView.md) | Windows is every window instance in range, newest first. | [optional] 

## Methods

### NewBoardResp

`func NewBoardResp() *BoardResp`

NewBoardResp instantiates a new BoardResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardRespWithDefaults

`func NewBoardRespWithDefaults() *BoardResp`

NewBoardRespWithDefaults instantiates a new BoardResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *BoardResp) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *BoardResp) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *BoardResp) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *BoardResp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAvailable

`func (o *BoardResp) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *BoardResp) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *BoardResp) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *BoardResp) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCurrent

`func (o *BoardResp) GetCurrent() []ReadingView`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *BoardResp) GetCurrentOk() (*[]ReadingView, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *BoardResp) SetCurrent(v []ReadingView)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *BoardResp) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetFrom

`func (o *BoardResp) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BoardResp) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BoardResp) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *BoardResp) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetProvider

`func (o *BoardResp) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BoardResp) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BoardResp) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *BoardResp) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRange

`func (o *BoardResp) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *BoardResp) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *BoardResp) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *BoardResp) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScope

`func (o *BoardResp) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *BoardResp) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *BoardResp) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *BoardResp) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSource

`func (o *BoardResp) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BoardResp) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BoardResp) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BoardResp) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTo

`func (o *BoardResp) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BoardResp) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BoardResp) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *BoardResp) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetWindows

`func (o *BoardResp) GetWindows() []ReadingView`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *BoardResp) GetWindowsOk() (*[]ReadingView, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *BoardResp) SetWindows(v []ReadingView)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *BoardResp) HasWindows() bool`

HasWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



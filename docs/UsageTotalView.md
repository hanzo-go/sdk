# UsageTotalView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confidence** | Pointer to **string** | Confidence says how much the counters mean; a percentage-only meter leaves them at zero. | [optional] 
**CostCents** | Pointer to **int64** | CostCents is the row&#39;s cost in US cents. For an \&quot;account\&quot; row this is the PROVIDER&#39;s own charge, not a Hanzo one. | [optional] 
**Provider** | Pointer to **string** | Provider is the upstream the usage was measured against. | [optional] 
**Requests** | Pointer to **int64** | Requests is how many requests the row covers. | [optional] 
**Scope** | Pointer to **string** | Scope is whose row it is: \&quot;user\&quot; for the caller&#39;s own linked accounts, \&quot;org\&quot; for the whole tenant&#39;s Hanzo-routed usage. | [optional] 
**Source** | Pointer to **string** | Source is where the row came from: \&quot;account\&quot; is the provider&#39;s own meter on the caller&#39;s linked account, \&quot;hanzo\&quot; is Hanzo-routed inference. The two are never summed. | [optional] 
**Tokens** | Pointer to **int64** | Tokens is the total tokens the row covers. | [optional] 
**UsedPct** | Pointer to **float64** | UsedPct is how much of a plan window the row consumed, 0–100. It is a share, never money. | [optional] 
**Window** | Pointer to **string** | Window is the meter window class the row rolls up, when it has one. | [optional] 
**Windows** | Pointer to **int64** | Windows is how many window instances rolled up into the row. | [optional] 

## Methods

### NewUsageTotalView

`func NewUsageTotalView() *UsageTotalView`

NewUsageTotalView instantiates a new UsageTotalView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageTotalViewWithDefaults

`func NewUsageTotalViewWithDefaults() *UsageTotalView`

NewUsageTotalViewWithDefaults instantiates a new UsageTotalView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidence

`func (o *UsageTotalView) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *UsageTotalView) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *UsageTotalView) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *UsageTotalView) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetCostCents

`func (o *UsageTotalView) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *UsageTotalView) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *UsageTotalView) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *UsageTotalView) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetProvider

`func (o *UsageTotalView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UsageTotalView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UsageTotalView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UsageTotalView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequests

`func (o *UsageTotalView) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *UsageTotalView) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *UsageTotalView) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *UsageTotalView) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetScope

`func (o *UsageTotalView) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UsageTotalView) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UsageTotalView) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UsageTotalView) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSource

`func (o *UsageTotalView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UsageTotalView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UsageTotalView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *UsageTotalView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTokens

`func (o *UsageTotalView) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *UsageTotalView) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *UsageTotalView) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *UsageTotalView) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetUsedPct

`func (o *UsageTotalView) GetUsedPct() float64`

GetUsedPct returns the UsedPct field if non-nil, zero value otherwise.

### GetUsedPctOk

`func (o *UsageTotalView) GetUsedPctOk() (*float64, bool)`

GetUsedPctOk returns a tuple with the UsedPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPct

`func (o *UsageTotalView) SetUsedPct(v float64)`

SetUsedPct sets UsedPct field to given value.

### HasUsedPct

`func (o *UsageTotalView) HasUsedPct() bool`

HasUsedPct returns a boolean if a field has been set.

### GetWindow

`func (o *UsageTotalView) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *UsageTotalView) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *UsageTotalView) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *UsageTotalView) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetWindows

`func (o *UsageTotalView) GetWindows() int64`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *UsageTotalView) GetWindowsOk() (*int64, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *UsageTotalView) SetWindows(v int64)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *UsageTotalView) HasWindows() bool`

HasWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



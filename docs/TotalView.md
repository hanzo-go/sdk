# TotalView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confidence** | Pointer to **string** | Confidence says how real the row&#39;s numbers are. | [optional] 
**CostCents** | Pointer to **int32** | CostCents is the period&#39;s spend in cents, in the row&#39;s own ledger. | [optional] 
**Provider** | Pointer to **string** | Provider is the provider the row totals. | [optional] 
**Requests** | Pointer to **int32** | Requests is the period&#39;s request count. | [optional] 
**Scope** | Pointer to **string** | Scope is whose usage the row measures: user or org. | [optional] 
**Source** | Pointer to **string** | Source is whose meter the row came from: account or hanzo. | [optional] 
**Tokens** | Pointer to **int32** | Tokens is the period&#39;s total token count. | [optional] 
**UsedPct** | Pointer to **float32** | UsedPct is the plan consumption percentage, on the account side. | [optional] 
**Window** | Pointer to **string** | Window is the window class the row totals, on the account side. | [optional] 
**Windows** | Pointer to **int32** | Windows is how many window instances the row folds. | [optional] 

## Methods

### NewTotalView

`func NewTotalView() *TotalView`

NewTotalView instantiates a new TotalView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTotalViewWithDefaults

`func NewTotalViewWithDefaults() *TotalView`

NewTotalViewWithDefaults instantiates a new TotalView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidence

`func (o *TotalView) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *TotalView) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *TotalView) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *TotalView) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetCostCents

`func (o *TotalView) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *TotalView) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *TotalView) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *TotalView) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetProvider

`func (o *TotalView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *TotalView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *TotalView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *TotalView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequests

`func (o *TotalView) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *TotalView) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *TotalView) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *TotalView) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetScope

`func (o *TotalView) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *TotalView) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *TotalView) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *TotalView) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSource

`func (o *TotalView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TotalView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TotalView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TotalView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTokens

`func (o *TotalView) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *TotalView) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *TotalView) SetTokens(v int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *TotalView) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetUsedPct

`func (o *TotalView) GetUsedPct() float32`

GetUsedPct returns the UsedPct field if non-nil, zero value otherwise.

### GetUsedPctOk

`func (o *TotalView) GetUsedPctOk() (*float32, bool)`

GetUsedPctOk returns a tuple with the UsedPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPct

`func (o *TotalView) SetUsedPct(v float32)`

SetUsedPct sets UsedPct field to given value.

### HasUsedPct

`func (o *TotalView) HasUsedPct() bool`

HasUsedPct returns a boolean if a field has been set.

### GetWindow

`func (o *TotalView) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *TotalView) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *TotalView) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *TotalView) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetWindows

`func (o *TotalView) GetWindows() int32`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *TotalView) GetWindowsOk() (*int32, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *TotalView) SetWindows(v int32)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *TotalView) HasWindows() bool`

HasWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



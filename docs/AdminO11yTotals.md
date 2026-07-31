# AdminO11yTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to **int64** |  | [optional] 
**Tokens** | Pointer to **int64** |  | [optional] 
**PromptTokens** | Pointer to **int64** |  | [optional] 
**CompletionTokens** | Pointer to **int64** |  | [optional] 
**CostCents** | Pointer to **int64** |  | [optional] 
**Errors** | Pointer to **int64** |  | [optional] 
**Orgs** | Pointer to **int64** |  | [optional] 
**Models** | Pointer to **int64** |  | [optional] 
**TraceCount** | Pointer to **int64** |  | [optional] 
**LatencyP50Ms** | Pointer to **float64** |  | [optional] 
**LatencyP95Ms** | Pointer to **float64** |  | [optional] 
**LatencyP99Ms** | Pointer to **float64** |  | [optional] 
**TraceErrorRate** | Pointer to **float64** |  | [optional] 
**Services** | Pointer to **int64** |  | [optional] 
**LogVolume** | Pointer to **int64** |  | [optional] 

## Methods

### NewAdminO11yTotals

`func NewAdminO11yTotals() *AdminO11yTotals`

NewAdminO11yTotals instantiates a new AdminO11yTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminO11yTotalsWithDefaults

`func NewAdminO11yTotalsWithDefaults() *AdminO11yTotals`

NewAdminO11yTotalsWithDefaults instantiates a new AdminO11yTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *AdminO11yTotals) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *AdminO11yTotals) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *AdminO11yTotals) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *AdminO11yTotals) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTokens

`func (o *AdminO11yTotals) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *AdminO11yTotals) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *AdminO11yTotals) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *AdminO11yTotals) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetPromptTokens

`func (o *AdminO11yTotals) GetPromptTokens() int64`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *AdminO11yTotals) GetPromptTokensOk() (*int64, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *AdminO11yTotals) SetPromptTokens(v int64)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *AdminO11yTotals) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetCompletionTokens

`func (o *AdminO11yTotals) GetCompletionTokens() int64`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *AdminO11yTotals) GetCompletionTokensOk() (*int64, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *AdminO11yTotals) SetCompletionTokens(v int64)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *AdminO11yTotals) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetCostCents

`func (o *AdminO11yTotals) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *AdminO11yTotals) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *AdminO11yTotals) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *AdminO11yTotals) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetErrors

`func (o *AdminO11yTotals) GetErrors() int64`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AdminO11yTotals) GetErrorsOk() (*int64, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AdminO11yTotals) SetErrors(v int64)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AdminO11yTotals) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetOrgs

`func (o *AdminO11yTotals) GetOrgs() int64`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *AdminO11yTotals) GetOrgsOk() (*int64, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *AdminO11yTotals) SetOrgs(v int64)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *AdminO11yTotals) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetModels

`func (o *AdminO11yTotals) GetModels() int64`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *AdminO11yTotals) GetModelsOk() (*int64, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *AdminO11yTotals) SetModels(v int64)`

SetModels sets Models field to given value.

### HasModels

`func (o *AdminO11yTotals) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetTraceCount

`func (o *AdminO11yTotals) GetTraceCount() int64`

GetTraceCount returns the TraceCount field if non-nil, zero value otherwise.

### GetTraceCountOk

`func (o *AdminO11yTotals) GetTraceCountOk() (*int64, bool)`

GetTraceCountOk returns a tuple with the TraceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceCount

`func (o *AdminO11yTotals) SetTraceCount(v int64)`

SetTraceCount sets TraceCount field to given value.

### HasTraceCount

`func (o *AdminO11yTotals) HasTraceCount() bool`

HasTraceCount returns a boolean if a field has been set.

### GetLatencyP50Ms

`func (o *AdminO11yTotals) GetLatencyP50Ms() float64`

GetLatencyP50Ms returns the LatencyP50Ms field if non-nil, zero value otherwise.

### GetLatencyP50MsOk

`func (o *AdminO11yTotals) GetLatencyP50MsOk() (*float64, bool)`

GetLatencyP50MsOk returns a tuple with the LatencyP50Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP50Ms

`func (o *AdminO11yTotals) SetLatencyP50Ms(v float64)`

SetLatencyP50Ms sets LatencyP50Ms field to given value.

### HasLatencyP50Ms

`func (o *AdminO11yTotals) HasLatencyP50Ms() bool`

HasLatencyP50Ms returns a boolean if a field has been set.

### GetLatencyP95Ms

`func (o *AdminO11yTotals) GetLatencyP95Ms() float64`

GetLatencyP95Ms returns the LatencyP95Ms field if non-nil, zero value otherwise.

### GetLatencyP95MsOk

`func (o *AdminO11yTotals) GetLatencyP95MsOk() (*float64, bool)`

GetLatencyP95MsOk returns a tuple with the LatencyP95Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP95Ms

`func (o *AdminO11yTotals) SetLatencyP95Ms(v float64)`

SetLatencyP95Ms sets LatencyP95Ms field to given value.

### HasLatencyP95Ms

`func (o *AdminO11yTotals) HasLatencyP95Ms() bool`

HasLatencyP95Ms returns a boolean if a field has been set.

### GetLatencyP99Ms

`func (o *AdminO11yTotals) GetLatencyP99Ms() float64`

GetLatencyP99Ms returns the LatencyP99Ms field if non-nil, zero value otherwise.

### GetLatencyP99MsOk

`func (o *AdminO11yTotals) GetLatencyP99MsOk() (*float64, bool)`

GetLatencyP99MsOk returns a tuple with the LatencyP99Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP99Ms

`func (o *AdminO11yTotals) SetLatencyP99Ms(v float64)`

SetLatencyP99Ms sets LatencyP99Ms field to given value.

### HasLatencyP99Ms

`func (o *AdminO11yTotals) HasLatencyP99Ms() bool`

HasLatencyP99Ms returns a boolean if a field has been set.

### GetTraceErrorRate

`func (o *AdminO11yTotals) GetTraceErrorRate() float64`

GetTraceErrorRate returns the TraceErrorRate field if non-nil, zero value otherwise.

### GetTraceErrorRateOk

`func (o *AdminO11yTotals) GetTraceErrorRateOk() (*float64, bool)`

GetTraceErrorRateOk returns a tuple with the TraceErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceErrorRate

`func (o *AdminO11yTotals) SetTraceErrorRate(v float64)`

SetTraceErrorRate sets TraceErrorRate field to given value.

### HasTraceErrorRate

`func (o *AdminO11yTotals) HasTraceErrorRate() bool`

HasTraceErrorRate returns a boolean if a field has been set.

### GetServices

`func (o *AdminO11yTotals) GetServices() int64`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AdminO11yTotals) GetServicesOk() (*int64, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AdminO11yTotals) SetServices(v int64)`

SetServices sets Services field to given value.

### HasServices

`func (o *AdminO11yTotals) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetLogVolume

`func (o *AdminO11yTotals) GetLogVolume() int64`

GetLogVolume returns the LogVolume field if non-nil, zero value otherwise.

### GetLogVolumeOk

`func (o *AdminO11yTotals) GetLogVolumeOk() (*int64, bool)`

GetLogVolumeOk returns a tuple with the LogVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogVolume

`func (o *AdminO11yTotals) SetLogVolume(v int64)`

SetLogVolume sets LogVolume field to given value.

### HasLogVolume

`func (o *AdminO11yTotals) HasLogVolume() bool`

HasLogVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



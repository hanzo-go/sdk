# CloudO11yTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionTokens** | Pointer to **int32** |  | [optional] 
**CostCents** | Pointer to **int32** |  | [optional] 
**Errors** | Pointer to **int32** |  | [optional] 
**LatencyP50Ms** | Pointer to **float32** |  | [optional] 
**LatencyP95Ms** | Pointer to **float32** |  | [optional] 
**LatencyP99Ms** | Pointer to **float32** |  | [optional] 
**LogVolume** | Pointer to **int32** | Logs (distributed_logs_v2), fleet volume over the window. | [optional] 
**Models** | Pointer to **int32** |  | [optional] 
**Orgs** | Pointer to **int32** |  | [optional] 
**PromptTokens** | Pointer to **int32** |  | [optional] 
**Requests** | Pointer to **int32** | LLM usage (hanzo.cloud_usage), all orgs. | [optional] 
**Services** | Pointer to **int32** |  | [optional] 
**Tokens** | Pointer to **int32** |  | [optional] 
**TraceCount** | Pointer to **int32** | Traces (o11y_index_v3), all services. | [optional] 
**TraceErrorRate** | Pointer to **float32** | percent (0..100) | [optional] 

## Methods

### NewCloudO11yTotals

`func NewCloudO11yTotals() *CloudO11yTotals`

NewCloudO11yTotals instantiates a new CloudO11yTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudO11yTotalsWithDefaults

`func NewCloudO11yTotalsWithDefaults() *CloudO11yTotals`

NewCloudO11yTotalsWithDefaults instantiates a new CloudO11yTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionTokens

`func (o *CloudO11yTotals) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *CloudO11yTotals) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *CloudO11yTotals) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *CloudO11yTotals) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetCostCents

`func (o *CloudO11yTotals) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *CloudO11yTotals) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *CloudO11yTotals) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *CloudO11yTotals) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetErrors

`func (o *CloudO11yTotals) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CloudO11yTotals) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CloudO11yTotals) SetErrors(v int32)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CloudO11yTotals) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetLatencyP50Ms

`func (o *CloudO11yTotals) GetLatencyP50Ms() float32`

GetLatencyP50Ms returns the LatencyP50Ms field if non-nil, zero value otherwise.

### GetLatencyP50MsOk

`func (o *CloudO11yTotals) GetLatencyP50MsOk() (*float32, bool)`

GetLatencyP50MsOk returns a tuple with the LatencyP50Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP50Ms

`func (o *CloudO11yTotals) SetLatencyP50Ms(v float32)`

SetLatencyP50Ms sets LatencyP50Ms field to given value.

### HasLatencyP50Ms

`func (o *CloudO11yTotals) HasLatencyP50Ms() bool`

HasLatencyP50Ms returns a boolean if a field has been set.

### GetLatencyP95Ms

`func (o *CloudO11yTotals) GetLatencyP95Ms() float32`

GetLatencyP95Ms returns the LatencyP95Ms field if non-nil, zero value otherwise.

### GetLatencyP95MsOk

`func (o *CloudO11yTotals) GetLatencyP95MsOk() (*float32, bool)`

GetLatencyP95MsOk returns a tuple with the LatencyP95Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP95Ms

`func (o *CloudO11yTotals) SetLatencyP95Ms(v float32)`

SetLatencyP95Ms sets LatencyP95Ms field to given value.

### HasLatencyP95Ms

`func (o *CloudO11yTotals) HasLatencyP95Ms() bool`

HasLatencyP95Ms returns a boolean if a field has been set.

### GetLatencyP99Ms

`func (o *CloudO11yTotals) GetLatencyP99Ms() float32`

GetLatencyP99Ms returns the LatencyP99Ms field if non-nil, zero value otherwise.

### GetLatencyP99MsOk

`func (o *CloudO11yTotals) GetLatencyP99MsOk() (*float32, bool)`

GetLatencyP99MsOk returns a tuple with the LatencyP99Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP99Ms

`func (o *CloudO11yTotals) SetLatencyP99Ms(v float32)`

SetLatencyP99Ms sets LatencyP99Ms field to given value.

### HasLatencyP99Ms

`func (o *CloudO11yTotals) HasLatencyP99Ms() bool`

HasLatencyP99Ms returns a boolean if a field has been set.

### GetLogVolume

`func (o *CloudO11yTotals) GetLogVolume() int32`

GetLogVolume returns the LogVolume field if non-nil, zero value otherwise.

### GetLogVolumeOk

`func (o *CloudO11yTotals) GetLogVolumeOk() (*int32, bool)`

GetLogVolumeOk returns a tuple with the LogVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogVolume

`func (o *CloudO11yTotals) SetLogVolume(v int32)`

SetLogVolume sets LogVolume field to given value.

### HasLogVolume

`func (o *CloudO11yTotals) HasLogVolume() bool`

HasLogVolume returns a boolean if a field has been set.

### GetModels

`func (o *CloudO11yTotals) GetModels() int32`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *CloudO11yTotals) GetModelsOk() (*int32, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *CloudO11yTotals) SetModels(v int32)`

SetModels sets Models field to given value.

### HasModels

`func (o *CloudO11yTotals) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetOrgs

`func (o *CloudO11yTotals) GetOrgs() int32`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *CloudO11yTotals) GetOrgsOk() (*int32, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *CloudO11yTotals) SetOrgs(v int32)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *CloudO11yTotals) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetPromptTokens

`func (o *CloudO11yTotals) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *CloudO11yTotals) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *CloudO11yTotals) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *CloudO11yTotals) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetRequests

`func (o *CloudO11yTotals) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *CloudO11yTotals) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *CloudO11yTotals) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *CloudO11yTotals) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetServices

`func (o *CloudO11yTotals) GetServices() int32`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *CloudO11yTotals) GetServicesOk() (*int32, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *CloudO11yTotals) SetServices(v int32)`

SetServices sets Services field to given value.

### HasServices

`func (o *CloudO11yTotals) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetTokens

`func (o *CloudO11yTotals) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *CloudO11yTotals) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *CloudO11yTotals) SetTokens(v int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *CloudO11yTotals) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetTraceCount

`func (o *CloudO11yTotals) GetTraceCount() int32`

GetTraceCount returns the TraceCount field if non-nil, zero value otherwise.

### GetTraceCountOk

`func (o *CloudO11yTotals) GetTraceCountOk() (*int32, bool)`

GetTraceCountOk returns a tuple with the TraceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceCount

`func (o *CloudO11yTotals) SetTraceCount(v int32)`

SetTraceCount sets TraceCount field to given value.

### HasTraceCount

`func (o *CloudO11yTotals) HasTraceCount() bool`

HasTraceCount returns a boolean if a field has been set.

### GetTraceErrorRate

`func (o *CloudO11yTotals) GetTraceErrorRate() float32`

GetTraceErrorRate returns the TraceErrorRate field if non-nil, zero value otherwise.

### GetTraceErrorRateOk

`func (o *CloudO11yTotals) GetTraceErrorRateOk() (*float32, bool)`

GetTraceErrorRateOk returns a tuple with the TraceErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceErrorRate

`func (o *CloudO11yTotals) SetTraceErrorRate(v float32)`

SetTraceErrorRate sets TraceErrorRate field to given value.

### HasTraceErrorRate

`func (o *CloudO11yTotals) HasTraceErrorRate() bool`

HasTraceErrorRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



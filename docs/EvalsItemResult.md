# EvalsItemResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | Pointer to **string** |  | [optional] 
**TraceId** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **float32** |  | [optional] 
**Rationale** | Pointer to **string** |  | [optional] 
**Output** | Pointer to **string** |  | [optional] 
**LatencyMs** | Pointer to **int32** |  | [optional] 
**Tokens** | Pointer to **int32** |  | [optional] 
**CostCents** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewEvalsItemResult

`func NewEvalsItemResult() *EvalsItemResult`

NewEvalsItemResult instantiates a new EvalsItemResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalsItemResultWithDefaults

`func NewEvalsItemResultWithDefaults() *EvalsItemResult`

NewEvalsItemResultWithDefaults instantiates a new EvalsItemResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *EvalsItemResult) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *EvalsItemResult) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *EvalsItemResult) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *EvalsItemResult) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetTraceId

`func (o *EvalsItemResult) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *EvalsItemResult) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *EvalsItemResult) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *EvalsItemResult) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetScore

`func (o *EvalsItemResult) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *EvalsItemResult) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *EvalsItemResult) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *EvalsItemResult) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetRationale

`func (o *EvalsItemResult) GetRationale() string`

GetRationale returns the Rationale field if non-nil, zero value otherwise.

### GetRationaleOk

`func (o *EvalsItemResult) GetRationaleOk() (*string, bool)`

GetRationaleOk returns a tuple with the Rationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRationale

`func (o *EvalsItemResult) SetRationale(v string)`

SetRationale sets Rationale field to given value.

### HasRationale

`func (o *EvalsItemResult) HasRationale() bool`

HasRationale returns a boolean if a field has been set.

### GetOutput

`func (o *EvalsItemResult) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *EvalsItemResult) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *EvalsItemResult) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *EvalsItemResult) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetLatencyMs

`func (o *EvalsItemResult) GetLatencyMs() int32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *EvalsItemResult) GetLatencyMsOk() (*int32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *EvalsItemResult) SetLatencyMs(v int32)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *EvalsItemResult) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### GetTokens

`func (o *EvalsItemResult) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *EvalsItemResult) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *EvalsItemResult) SetTokens(v int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *EvalsItemResult) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetCostCents

`func (o *EvalsItemResult) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *EvalsItemResult) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *EvalsItemResult) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *EvalsItemResult) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetError

`func (o *EvalsItemResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EvalsItemResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EvalsItemResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *EvalsItemResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



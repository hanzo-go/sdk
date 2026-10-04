# EventUsagePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to **int64** | Requests is how many LLM calls fell in this bucket. | [optional] 
**SpendCents** | Pointer to **int64** | SpendCents is what they cost, in cents. | [optional] 
**T** | Pointer to **string** | T is the bucket&#39;s start, RFC3339 UTC, aligned to the interval. | [optional] 
**Tokens** | Pointer to **int64** | Tokens is prompt plus completion tokens over those calls. | [optional] 

## Methods

### NewEventUsagePoint

`func NewEventUsagePoint() *EventUsagePoint`

NewEventUsagePoint instantiates a new EventUsagePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventUsagePointWithDefaults

`func NewEventUsagePointWithDefaults() *EventUsagePoint`

NewEventUsagePointWithDefaults instantiates a new EventUsagePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *EventUsagePoint) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *EventUsagePoint) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *EventUsagePoint) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *EventUsagePoint) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSpendCents

`func (o *EventUsagePoint) GetSpendCents() int64`

GetSpendCents returns the SpendCents field if non-nil, zero value otherwise.

### GetSpendCentsOk

`func (o *EventUsagePoint) GetSpendCentsOk() (*int64, bool)`

GetSpendCentsOk returns a tuple with the SpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents

`func (o *EventUsagePoint) SetSpendCents(v int64)`

SetSpendCents sets SpendCents field to given value.

### HasSpendCents

`func (o *EventUsagePoint) HasSpendCents() bool`

HasSpendCents returns a boolean if a field has been set.

### GetT

`func (o *EventUsagePoint) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *EventUsagePoint) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *EventUsagePoint) SetT(v string)`

SetT sets T field to given value.

### HasT

`func (o *EventUsagePoint) HasT() bool`

HasT returns a boolean if a field has been set.

### GetTokens

`func (o *EventUsagePoint) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *EventUsagePoint) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *EventUsagePoint) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *EventUsagePoint) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



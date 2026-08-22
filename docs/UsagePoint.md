# UsagePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to **int32** | Requests is how many LLM calls fell in this bucket. | [optional] 
**SpendCents** | Pointer to **int32** | SpendCents is what they cost, in cents. | [optional] 
**T** | Pointer to **string** | T is the bucket&#39;s start, RFC3339 UTC, aligned to the interval. | [optional] 
**Tokens** | Pointer to **int32** | Tokens is prompt plus completion tokens over those calls. | [optional] 

## Methods

### NewUsagePoint

`func NewUsagePoint() *UsagePoint`

NewUsagePoint instantiates a new UsagePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsagePointWithDefaults

`func NewUsagePointWithDefaults() *UsagePoint`

NewUsagePointWithDefaults instantiates a new UsagePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *UsagePoint) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *UsagePoint) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *UsagePoint) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *UsagePoint) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSpendCents

`func (o *UsagePoint) GetSpendCents() int32`

GetSpendCents returns the SpendCents field if non-nil, zero value otherwise.

### GetSpendCentsOk

`func (o *UsagePoint) GetSpendCentsOk() (*int32, bool)`

GetSpendCentsOk returns a tuple with the SpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents

`func (o *UsagePoint) SetSpendCents(v int32)`

SetSpendCents sets SpendCents field to given value.

### HasSpendCents

`func (o *UsagePoint) HasSpendCents() bool`

HasSpendCents returns a boolean if a field has been set.

### GetT

`func (o *UsagePoint) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UsagePoint) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UsagePoint) SetT(v string)`

SetT sets T field to given value.

### HasT

`func (o *UsagePoint) HasT() bool`

HasT returns a boolean if a field has been set.

### GetTokens

`func (o *UsagePoint) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *UsagePoint) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *UsagePoint) SetTokens(v int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *UsagePoint) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



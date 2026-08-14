# O11yUsageBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calls** | Pointer to **int32** | Calls is how many LLM calls landed in the bucket. | [optional] 
**CostCents** | Pointer to **int32** | CostCents is what they cost, in cents. | [optional] 
**T** | Pointer to **string** | T is the bucket start, RFC3339 in UTC. | [optional] 
**Tokens** | Pointer to **int32** | Tokens is how many tokens they consumed. | [optional] 

## Methods

### NewO11yUsageBucket

`func NewO11yUsageBucket() *O11yUsageBucket`

NewO11yUsageBucket instantiates a new O11yUsageBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yUsageBucketWithDefaults

`func NewO11yUsageBucketWithDefaults() *O11yUsageBucket`

NewO11yUsageBucketWithDefaults instantiates a new O11yUsageBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalls

`func (o *O11yUsageBucket) GetCalls() int32`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *O11yUsageBucket) GetCallsOk() (*int32, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *O11yUsageBucket) SetCalls(v int32)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *O11yUsageBucket) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetCostCents

`func (o *O11yUsageBucket) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *O11yUsageBucket) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *O11yUsageBucket) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *O11yUsageBucket) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetT

`func (o *O11yUsageBucket) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *O11yUsageBucket) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *O11yUsageBucket) SetT(v string)`

SetT sets T field to given value.

### HasT

`func (o *O11yUsageBucket) HasT() bool`

HasT returns a boolean if a field has been set.

### GetTokens

`func (o *O11yUsageBucket) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *O11yUsageBucket) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *O11yUsageBucket) SetTokens(v int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *O11yUsageBucket) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



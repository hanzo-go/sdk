# ObserveUsagePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**T** | Pointer to **string** | RFC3339 bucket start (UTC). | [optional] 
**Calls** | Pointer to **int64** |  | [optional] 
**Tokens** | Pointer to **int64** |  | [optional] 
**CostCents** | Pointer to **int64** |  | [optional] 

## Methods

### NewObserveUsagePoint

`func NewObserveUsagePoint() *ObserveUsagePoint`

NewObserveUsagePoint instantiates a new ObserveUsagePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObserveUsagePointWithDefaults

`func NewObserveUsagePointWithDefaults() *ObserveUsagePoint`

NewObserveUsagePointWithDefaults instantiates a new ObserveUsagePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetT

`func (o *ObserveUsagePoint) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *ObserveUsagePoint) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *ObserveUsagePoint) SetT(v string)`

SetT sets T field to given value.

### HasT

`func (o *ObserveUsagePoint) HasT() bool`

HasT returns a boolean if a field has been set.

### GetCalls

`func (o *ObserveUsagePoint) GetCalls() int64`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *ObserveUsagePoint) GetCallsOk() (*int64, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *ObserveUsagePoint) SetCalls(v int64)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *ObserveUsagePoint) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetTokens

`func (o *ObserveUsagePoint) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ObserveUsagePoint) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ObserveUsagePoint) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *ObserveUsagePoint) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetCostCents

`func (o *ObserveUsagePoint) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *ObserveUsagePoint) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *ObserveUsagePoint) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *ObserveUsagePoint) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



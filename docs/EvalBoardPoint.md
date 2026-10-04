# EvalBoardPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostCents** | Pointer to **int64** | what this bucket cost, in cents | [optional] 
**Errors** | Pointer to **int64** | calls in this bucket that did not succeed | [optional] 
**Generations** | Pointer to **int64** | model calls in this bucket | [optional] 
**T** | Pointer to **string** | RFC3339 (UTC) bucket start | [optional] 
**TotalTokens** | Pointer to **int64** | tokens in this bucket | [optional] 

## Methods

### NewEvalBoardPoint

`func NewEvalBoardPoint() *EvalBoardPoint`

NewEvalBoardPoint instantiates a new EvalBoardPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalBoardPointWithDefaults

`func NewEvalBoardPointWithDefaults() *EvalBoardPoint`

NewEvalBoardPointWithDefaults instantiates a new EvalBoardPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostCents

`func (o *EvalBoardPoint) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *EvalBoardPoint) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *EvalBoardPoint) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *EvalBoardPoint) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetErrors

`func (o *EvalBoardPoint) GetErrors() int64`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *EvalBoardPoint) GetErrorsOk() (*int64, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *EvalBoardPoint) SetErrors(v int64)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *EvalBoardPoint) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetGenerations

`func (o *EvalBoardPoint) GetGenerations() int64`

GetGenerations returns the Generations field if non-nil, zero value otherwise.

### GetGenerationsOk

`func (o *EvalBoardPoint) GetGenerationsOk() (*int64, bool)`

GetGenerationsOk returns a tuple with the Generations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerations

`func (o *EvalBoardPoint) SetGenerations(v int64)`

SetGenerations sets Generations field to given value.

### HasGenerations

`func (o *EvalBoardPoint) HasGenerations() bool`

HasGenerations returns a boolean if a field has been set.

### GetT

`func (o *EvalBoardPoint) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *EvalBoardPoint) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *EvalBoardPoint) SetT(v string)`

SetT sets T field to given value.

### HasT

`func (o *EvalBoardPoint) HasT() bool`

HasT returns a boolean if a field has been set.

### GetTotalTokens

`func (o *EvalBoardPoint) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *EvalBoardPoint) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *EvalBoardPoint) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *EvalBoardPoint) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



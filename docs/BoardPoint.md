# BoardPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostCents** | Pointer to **int32** | what this bucket cost, in cents | [optional] 
**Errors** | Pointer to **int32** | calls in this bucket that did not succeed | [optional] 
**Generations** | Pointer to **int32** | model calls in this bucket | [optional] 
**T** | Pointer to **string** | RFC3339 (UTC) bucket start | [optional] 
**TotalTokens** | Pointer to **int32** | tokens in this bucket | [optional] 

## Methods

### NewBoardPoint

`func NewBoardPoint() *BoardPoint`

NewBoardPoint instantiates a new BoardPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardPointWithDefaults

`func NewBoardPointWithDefaults() *BoardPoint`

NewBoardPointWithDefaults instantiates a new BoardPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostCents

`func (o *BoardPoint) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *BoardPoint) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *BoardPoint) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *BoardPoint) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetErrors

`func (o *BoardPoint) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BoardPoint) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BoardPoint) SetErrors(v int32)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BoardPoint) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetGenerations

`func (o *BoardPoint) GetGenerations() int32`

GetGenerations returns the Generations field if non-nil, zero value otherwise.

### GetGenerationsOk

`func (o *BoardPoint) GetGenerationsOk() (*int32, bool)`

GetGenerationsOk returns a tuple with the Generations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerations

`func (o *BoardPoint) SetGenerations(v int32)`

SetGenerations sets Generations field to given value.

### HasGenerations

`func (o *BoardPoint) HasGenerations() bool`

HasGenerations returns a boolean if a field has been set.

### GetT

`func (o *BoardPoint) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *BoardPoint) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *BoardPoint) SetT(v string)`

SetT sets T field to given value.

### HasT

`func (o *BoardPoint) HasT() bool`

HasT returns a boolean if a field has been set.

### GetTotalTokens

`func (o *BoardPoint) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *BoardPoint) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *BoardPoint) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *BoardPoint) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



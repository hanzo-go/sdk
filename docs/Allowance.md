# Allowance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | calls the plan allows per period; 0 &#x3D; unbounded | [optional] 
**Plan** | Pointer to **string** | the tier the limit came from | [optional] 
**Resets** | Pointer to **int32** | unix seconds; when the count starts again | [optional] 
**Spent** | Pointer to **bool** | the subject is at the limit | [optional] 
**Used** | Pointer to **int32** |  | [optional] 

## Methods

### NewAllowance

`func NewAllowance() *Allowance`

NewAllowance instantiates a new Allowance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowanceWithDefaults

`func NewAllowanceWithDefaults() *Allowance`

NewAllowanceWithDefaults instantiates a new Allowance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *Allowance) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Allowance) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Allowance) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Allowance) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetPlan

`func (o *Allowance) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Allowance) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Allowance) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Allowance) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetResets

`func (o *Allowance) GetResets() int32`

GetResets returns the Resets field if non-nil, zero value otherwise.

### GetResetsOk

`func (o *Allowance) GetResetsOk() (*int32, bool)`

GetResetsOk returns a tuple with the Resets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResets

`func (o *Allowance) SetResets(v int32)`

SetResets sets Resets field to given value.

### HasResets

`func (o *Allowance) HasResets() bool`

HasResets returns a boolean if a field has been set.

### GetSpent

`func (o *Allowance) GetSpent() bool`

GetSpent returns the Spent field if non-nil, zero value otherwise.

### GetSpentOk

`func (o *Allowance) GetSpentOk() (*bool, bool)`

GetSpentOk returns a tuple with the Spent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpent

`func (o *Allowance) SetSpent(v bool)`

SetSpent sets Spent field to given value.

### HasSpent

`func (o *Allowance) HasSpent() bool`

HasSpent returns a boolean if a field has been set.

### GetUsed

`func (o *Allowance) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *Allowance) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *Allowance) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *Allowance) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



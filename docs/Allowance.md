# Allowance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int64** | calls the plan allows per period; 0 &#x3D; unbounded | [optional] 
**Plan** | Pointer to **string** | the tier the limit came from | [optional] 
**Resets** | Pointer to **int64** | unix seconds; when THAT window starts again | [optional] 
**Spent** | Pointer to **bool** | the subject is at the limit | [optional] 
**Used** | Pointer to **int64** | Used is how many zero-priced calls this subject has been SERVED in the period ending at Resets — the UTC calendar day. Only a served call counts, so an admission check, a refusal, or a vendor that never answered leaves it where it stood. It stops AT Limit rather than climbing past it, so Limit-Used is what remains and never goes negative. | [optional] 
**Window** | Pointer to **string** | Window is which ceiling these numbers describe — \&quot;hour\&quot; or \&quot;day\&quot; — because a caller is held to both and only one of them is the answer. It is the window that REFUSED where one did, and otherwise the one with least left, so Limit-Used is always the number that will actually stop them next. Empty where no window bounds the subject at all. | [optional] 

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

`func (o *Allowance) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Allowance) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Allowance) SetLimit(v int64)`

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

`func (o *Allowance) GetResets() int64`

GetResets returns the Resets field if non-nil, zero value otherwise.

### GetResetsOk

`func (o *Allowance) GetResetsOk() (*int64, bool)`

GetResetsOk returns a tuple with the Resets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResets

`func (o *Allowance) SetResets(v int64)`

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

`func (o *Allowance) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *Allowance) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *Allowance) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *Allowance) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetWindow

`func (o *Allowance) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *Allowance) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *Allowance) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *Allowance) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



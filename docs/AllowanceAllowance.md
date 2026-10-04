# AllowanceAllowance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int64** | calls the plan allows per period; 0 &#x3D; unbounded | [optional] 
**Plan** | Pointer to **string** | the tier the limit came from | [optional] 
**Pool** | Pointer to [**AllowancePool**](AllowancePool.md) | Pool is that shared pool&#39;s standing, where it could be read. Absent when the subject is not pooled, and when the process that holds it did not answer — never a guess. | [optional] 
**Pooled** | Pointer to **bool** | Pooled says the calls counted here are served from the pool every free user shares — the platform&#39;s vendor accounts for free models — rather than from the subject&#39;s own plan. True exactly where a window bounds the subject. | [optional] 
**Resets** | Pointer to **int64** | unix seconds; when THAT window starts again | [optional] 
**Spent** | Pointer to **bool** | Spent says the ceiling refuses: on a take, that THIS call was refused and nothing was counted; on a read, that the next call would be. | [optional] 
**Used** | Pointer to **int64** | Used is how many free calls this subject has been admitted in the period ending at Resets — the UTC calendar day. A call counts when it is admitted, whatever its upstream then does; a refused call counts nothing. It stops AT Limit rather than climbing past it, so Limit-Used is what remains and never goes negative. | [optional] 
**Window** | Pointer to **string** | Window is which ceiling these numbers describe — \&quot;hour\&quot; or \&quot;day\&quot; — because a caller is held to both and only one of them is the answer. It is the window that REFUSED where one did, and otherwise the one with least left, so Limit-Used is always the number that will actually stop them next. Empty where no window bounds the subject at all. | [optional] 

## Methods

### NewAllowanceAllowance

`func NewAllowanceAllowance() *AllowanceAllowance`

NewAllowanceAllowance instantiates a new AllowanceAllowance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowanceAllowanceWithDefaults

`func NewAllowanceAllowanceWithDefaults() *AllowanceAllowance`

NewAllowanceAllowanceWithDefaults instantiates a new AllowanceAllowance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *AllowanceAllowance) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AllowanceAllowance) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AllowanceAllowance) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AllowanceAllowance) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetPlan

`func (o *AllowanceAllowance) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AllowanceAllowance) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AllowanceAllowance) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AllowanceAllowance) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPool

`func (o *AllowanceAllowance) GetPool() AllowancePool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *AllowanceAllowance) GetPoolOk() (*AllowancePool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *AllowanceAllowance) SetPool(v AllowancePool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *AllowanceAllowance) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPooled

`func (o *AllowanceAllowance) GetPooled() bool`

GetPooled returns the Pooled field if non-nil, zero value otherwise.

### GetPooledOk

`func (o *AllowanceAllowance) GetPooledOk() (*bool, bool)`

GetPooledOk returns a tuple with the Pooled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPooled

`func (o *AllowanceAllowance) SetPooled(v bool)`

SetPooled sets Pooled field to given value.

### HasPooled

`func (o *AllowanceAllowance) HasPooled() bool`

HasPooled returns a boolean if a field has been set.

### GetResets

`func (o *AllowanceAllowance) GetResets() int64`

GetResets returns the Resets field if non-nil, zero value otherwise.

### GetResetsOk

`func (o *AllowanceAllowance) GetResetsOk() (*int64, bool)`

GetResetsOk returns a tuple with the Resets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResets

`func (o *AllowanceAllowance) SetResets(v int64)`

SetResets sets Resets field to given value.

### HasResets

`func (o *AllowanceAllowance) HasResets() bool`

HasResets returns a boolean if a field has been set.

### GetSpent

`func (o *AllowanceAllowance) GetSpent() bool`

GetSpent returns the Spent field if non-nil, zero value otherwise.

### GetSpentOk

`func (o *AllowanceAllowance) GetSpentOk() (*bool, bool)`

GetSpentOk returns a tuple with the Spent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpent

`func (o *AllowanceAllowance) SetSpent(v bool)`

SetSpent sets Spent field to given value.

### HasSpent

`func (o *AllowanceAllowance) HasSpent() bool`

HasSpent returns a boolean if a field has been set.

### GetUsed

`func (o *AllowanceAllowance) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *AllowanceAllowance) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *AllowanceAllowance) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *AllowanceAllowance) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetWindow

`func (o *AllowanceAllowance) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *AllowanceAllowance) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *AllowanceAllowance) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *AllowanceAllowance) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



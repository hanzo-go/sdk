# AllowancePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resets** | Pointer to **int64** | Resets is when the pool refills, unix seconds: set when it is exhausted, and when it is busy for less than a minute. Absent otherwise. | [optional] 
**State** | Pointer to **string** | State is \&quot;available\&quot; (every account takes free requests), \&quot;busy\&quot; (some account is at its vendor limit, or all are and one is back within a minute) or \&quot;exhausted\&quot; (none is until Resets). | [optional] 

## Methods

### NewAllowancePool

`func NewAllowancePool() *AllowancePool`

NewAllowancePool instantiates a new AllowancePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowancePoolWithDefaults

`func NewAllowancePoolWithDefaults() *AllowancePool`

NewAllowancePoolWithDefaults instantiates a new AllowancePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResets

`func (o *AllowancePool) GetResets() int64`

GetResets returns the Resets field if non-nil, zero value otherwise.

### GetResetsOk

`func (o *AllowancePool) GetResetsOk() (*int64, bool)`

GetResetsOk returns a tuple with the Resets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResets

`func (o *AllowancePool) SetResets(v int64)`

SetResets sets Resets field to given value.

### HasResets

`func (o *AllowancePool) HasResets() bool`

HasResets returns a boolean if a field has been set.

### GetState

`func (o *AllowancePool) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AllowancePool) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AllowancePool) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AllowancePool) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



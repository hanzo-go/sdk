# Loss

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exhausted** | Pointer to **int32** | Exhausted counts facts the bus abandoned after maxDeliver failed inserts. | [optional] 
**Undecodable** | Pointer to **int32** | Undecodable counts messages acked without landing because they did not parse. | [optional] 

## Methods

### NewLoss

`func NewLoss() *Loss`

NewLoss instantiates a new Loss object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLossWithDefaults

`func NewLossWithDefaults() *Loss`

NewLossWithDefaults instantiates a new Loss object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExhausted

`func (o *Loss) GetExhausted() int32`

GetExhausted returns the Exhausted field if non-nil, zero value otherwise.

### GetExhaustedOk

`func (o *Loss) GetExhaustedOk() (*int32, bool)`

GetExhaustedOk returns a tuple with the Exhausted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExhausted

`func (o *Loss) SetExhausted(v int32)`

SetExhausted sets Exhausted field to given value.

### HasExhausted

`func (o *Loss) HasExhausted() bool`

HasExhausted returns a boolean if a field has been set.

### GetUndecodable

`func (o *Loss) GetUndecodable() int32`

GetUndecodable returns the Undecodable field if non-nil, zero value otherwise.

### GetUndecodableOk

`func (o *Loss) GetUndecodableOk() (*int32, bool)`

GetUndecodableOk returns a tuple with the Undecodable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndecodable

`func (o *Loss) SetUndecodable(v int32)`

SetUndecodable sets Undecodable field to given value.

### HasUndecodable

`func (o *Loss) HasUndecodable() bool`

HasUndecodable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



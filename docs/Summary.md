# Summary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calls** | Pointer to **int64** | Calls is how many calls this org has placed or received, over its whole history — a running total, not a window. | [optional] 
**Messages** | Pointer to **int64** | Messages is the same running total for messages. | [optional] 
**Numbers** | Pointer to **int64** | Numbers is how many numbers this org holds right now. | [optional] 

## Methods

### NewSummary

`func NewSummary() *Summary`

NewSummary instantiates a new Summary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryWithDefaults

`func NewSummaryWithDefaults() *Summary`

NewSummaryWithDefaults instantiates a new Summary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalls

`func (o *Summary) GetCalls() int64`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *Summary) GetCallsOk() (*int64, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *Summary) SetCalls(v int64)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *Summary) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetMessages

`func (o *Summary) GetMessages() int64`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *Summary) GetMessagesOk() (*int64, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *Summary) SetMessages(v int64)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *Summary) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetNumbers

`func (o *Summary) GetNumbers() int64`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *Summary) GetNumbersOk() (*int64, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *Summary) SetNumbers(v int64)`

SetNumbers sets Numbers field to given value.

### HasNumbers

`func (o *Summary) HasNumbers() bool`

HasNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# History

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is the token this is the history of, lowercased. | [optional] 
**Chain** | Pointer to **string** |  | [optional] 
**Days** | Pointer to [**[]Day**](Day.md) | Days is oldest first, which is the order a chart draws. &#x60;[]&#x60; for a token the indexer holds no day for, &#x60;null&#x60; where the read failed. | [optional] 
**Reach** | Pointer to [**Reach**](Reach.md) |  | [optional] 

## Methods

### NewHistory

`func NewHistory() *History`

NewHistory instantiates a new History object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryWithDefaults

`func NewHistoryWithDefaults() *History`

NewHistoryWithDefaults instantiates a new History object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *History) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *History) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *History) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *History) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetChain

`func (o *History) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *History) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *History) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *History) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetDays

`func (o *History) GetDays() []Day`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *History) GetDaysOk() (*[]Day, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *History) SetDays(v []Day)`

SetDays sets Days field to given value.

### HasDays

`func (o *History) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetReach

`func (o *History) GetReach() Reach`

GetReach returns the Reach field if non-nil, zero value otherwise.

### GetReachOk

`func (o *History) GetReachOk() (*Reach, bool)`

GetReachOk returns a tuple with the Reach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReach

`func (o *History) SetReach(v Reach)`

SetReach sets Reach field to given value.

### HasReach

`func (o *History) HasReach() bool`

HasReach returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



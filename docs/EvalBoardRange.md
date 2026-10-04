# EvalBoardRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** | RFC3339 (UTC) | [optional] 
**Interval** | Pointer to **string** | hour | day | [optional] 
**Range** | Pointer to **string** | echoed label (24h | 7d | 30d | custom) | [optional] 
**Start** | Pointer to **string** | RFC3339 (UTC) | [optional] 

## Methods

### NewEvalBoardRange

`func NewEvalBoardRange() *EvalBoardRange`

NewEvalBoardRange instantiates a new EvalBoardRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalBoardRangeWithDefaults

`func NewEvalBoardRangeWithDefaults() *EvalBoardRange`

NewEvalBoardRangeWithDefaults instantiates a new EvalBoardRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *EvalBoardRange) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *EvalBoardRange) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *EvalBoardRange) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *EvalBoardRange) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetInterval

`func (o *EvalBoardRange) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *EvalBoardRange) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *EvalBoardRange) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *EvalBoardRange) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetRange

`func (o *EvalBoardRange) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *EvalBoardRange) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *EvalBoardRange) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *EvalBoardRange) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetStart

`func (o *EvalBoardRange) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *EvalBoardRange) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *EvalBoardRange) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *EvalBoardRange) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



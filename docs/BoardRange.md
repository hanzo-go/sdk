# BoardRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** | RFC3339 (UTC) | [optional] 
**Interval** | Pointer to **string** | hour | day | [optional] 
**Range** | Pointer to **string** | echoed label (24h | 7d | 30d | custom) | [optional] 
**Start** | Pointer to **string** | RFC3339 (UTC) | [optional] 

## Methods

### NewBoardRange

`func NewBoardRange() *BoardRange`

NewBoardRange instantiates a new BoardRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardRangeWithDefaults

`func NewBoardRangeWithDefaults() *BoardRange`

NewBoardRangeWithDefaults instantiates a new BoardRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *BoardRange) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *BoardRange) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *BoardRange) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *BoardRange) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetInterval

`func (o *BoardRange) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *BoardRange) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *BoardRange) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *BoardRange) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetRange

`func (o *BoardRange) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *BoardRange) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *BoardRange) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *BoardRange) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetStart

`func (o *BoardRange) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *BoardRange) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *BoardRange) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *BoardRange) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



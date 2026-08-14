# O11yO11yStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** | Time is the start of the bucket. | [optional] 
**Value** | Pointer to **int32** | Value is how many events fell in it. | [optional] 

## Methods

### NewO11yO11yStat

`func NewO11yO11yStat() *O11yO11yStat`

NewO11yO11yStat instantiates a new O11yO11yStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yStatWithDefaults

`func NewO11yO11yStatWithDefaults() *O11yO11yStat`

NewO11yO11yStatWithDefaults instantiates a new O11yO11yStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *O11yO11yStat) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *O11yO11yStat) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *O11yO11yStat) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *O11yO11yStat) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetValue

`func (o *O11yO11yStat) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yO11yStat) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yO11yStat) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yO11yStat) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



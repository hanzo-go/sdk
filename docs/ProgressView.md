# ProgressView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Done** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **string** |  | [optional] 
**Percent** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewProgressView

`func NewProgressView() *ProgressView`

NewProgressView instantiates a new ProgressView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgressViewWithDefaults

`func NewProgressViewWithDefaults() *ProgressView`

NewProgressViewWithDefaults instantiates a new ProgressView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDone

`func (o *ProgressView) GetDone() int32`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *ProgressView) GetDoneOk() (*int32, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *ProgressView) SetDone(v int32)`

SetDone sets Done field to given value.

### HasDone

`func (o *ProgressView) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetNext

`func (o *ProgressView) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ProgressView) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ProgressView) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ProgressView) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPercent

`func (o *ProgressView) GetPercent() int32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *ProgressView) GetPercentOk() (*int32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *ProgressView) SetPercent(v int32)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *ProgressView) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetTotal

`func (o *ProgressView) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProgressView) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProgressView) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProgressView) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



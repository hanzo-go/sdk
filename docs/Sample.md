# Sample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cents** | Pointer to **int32** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 

## Methods

### NewSample

`func NewSample() *Sample`

NewSample instantiates a new Sample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSampleWithDefaults

`func NewSampleWithDefaults() *Sample`

NewSampleWithDefaults instantiates a new Sample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCents

`func (o *Sample) GetCents() int32`

GetCents returns the Cents field if non-nil, zero value otherwise.

### GetCentsOk

`func (o *Sample) GetCentsOk() (*int32, bool)`

GetCentsOk returns a tuple with the Cents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCents

`func (o *Sample) SetCents(v int32)`

SetCents sets Cents field to given value.

### HasCents

`func (o *Sample) HasCents() bool`

HasCents returns a boolean if a field has been set.

### GetDate

`func (o *Sample) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Sample) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Sample) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Sample) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



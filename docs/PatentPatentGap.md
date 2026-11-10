# PatentPatentGap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | Field is the wire name, e.g. prior_art. | [optional] 
**Missing** | Pointer to **int64** | Missing is how many covered records leave it empty. | [optional] 
**Why** | Pointer to **string** | Why names what the absence costs, so a gap is a decision to fix rather than a lint warning. | [optional] 

## Methods

### NewPatentPatentGap

`func NewPatentPatentGap() *PatentPatentGap`

NewPatentPatentGap instantiates a new PatentPatentGap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentPatentGapWithDefaults

`func NewPatentPatentGapWithDefaults() *PatentPatentGap`

NewPatentPatentGapWithDefaults instantiates a new PatentPatentGap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *PatentPatentGap) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *PatentPatentGap) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *PatentPatentGap) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *PatentPatentGap) HasField() bool`

HasField returns a boolean if a field has been set.

### GetMissing

`func (o *PatentPatentGap) GetMissing() int64`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *PatentPatentGap) GetMissingOk() (*int64, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *PatentPatentGap) SetMissing(v int64)`

SetMissing sets Missing field to given value.

### HasMissing

`func (o *PatentPatentGap) HasMissing() bool`

HasMissing returns a boolean if a field has been set.

### GetWhy

`func (o *PatentPatentGap) GetWhy() string`

GetWhy returns the Why field if non-nil, zero value otherwise.

### GetWhyOk

`func (o *PatentPatentGap) GetWhyOk() (*string, bool)`

GetWhyOk returns a tuple with the Why field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhy

`func (o *PatentPatentGap) SetWhy(v string)`

SetWhy sets Why field to given value.

### HasWhy

`func (o *PatentPatentGap) HasWhy() bool`

HasWhy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



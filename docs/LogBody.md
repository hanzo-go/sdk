# LogBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 

## Methods

### NewLogBody

`func NewLogBody() *LogBody`

NewLogBody instantiates a new LogBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogBodyWithDefaults

`func NewLogBodyWithDefaults() *LogBody`

NewLogBodyWithDefaults instantiates a new LogBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *LogBody) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *LogBody) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *LogBody) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *LogBody) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetNumber

`func (o *LogBody) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *LogBody) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *LogBody) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *LogBody) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetSeverity

`func (o *LogBody) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *LogBody) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *LogBody) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *LogBody) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



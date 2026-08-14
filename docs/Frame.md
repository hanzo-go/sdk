# Frame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | Pointer to **int32** |  | [optional] 
**File** | Pointer to **string** |  | [optional] 
**Function** | Pointer to **string** |  | [optional] 
**Line** | Pointer to **int32** |  | [optional] 

## Methods

### NewFrame

`func NewFrame() *Frame`

NewFrame instantiates a new Frame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameWithDefaults

`func NewFrameWithDefaults() *Frame`

NewFrameWithDefaults instantiates a new Frame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *Frame) GetColumn() int32`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *Frame) GetColumnOk() (*int32, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *Frame) SetColumn(v int32)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *Frame) HasColumn() bool`

HasColumn returns a boolean if a field has been set.

### GetFile

`func (o *Frame) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *Frame) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *Frame) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *Frame) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFunction

`func (o *Frame) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *Frame) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *Frame) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *Frame) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetLine

`func (o *Frame) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *Frame) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *Frame) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *Frame) HasLine() bool`

HasLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



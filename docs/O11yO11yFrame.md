# O11yO11yFrame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | Pointer to **int32** | Column is the column number. | [optional] 
**File** | Pointer to **string** | File is the file it is in. | [optional] 
**Function** | Pointer to **string** | Function is the function the frame is in. | [optional] 
**Line** | Pointer to **int32** | Line is the line number. | [optional] 
**Own** | Pointer to **bool** | Own marks a frame in the reporting application&#39;s own code rather than in a dependency or the runtime. | [optional] 

## Methods

### NewO11yO11yFrame

`func NewO11yO11yFrame() *O11yO11yFrame`

NewO11yO11yFrame instantiates a new O11yO11yFrame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yFrameWithDefaults

`func NewO11yO11yFrameWithDefaults() *O11yO11yFrame`

NewO11yO11yFrameWithDefaults instantiates a new O11yO11yFrame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *O11yO11yFrame) GetColumn() int32`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *O11yO11yFrame) GetColumnOk() (*int32, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *O11yO11yFrame) SetColumn(v int32)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *O11yO11yFrame) HasColumn() bool`

HasColumn returns a boolean if a field has been set.

### GetFile

`func (o *O11yO11yFrame) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *O11yO11yFrame) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *O11yO11yFrame) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *O11yO11yFrame) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFunction

`func (o *O11yO11yFrame) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *O11yO11yFrame) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *O11yO11yFrame) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *O11yO11yFrame) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetLine

`func (o *O11yO11yFrame) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *O11yO11yFrame) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *O11yO11yFrame) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *O11yO11yFrame) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetOwn

`func (o *O11yO11yFrame) GetOwn() bool`

GetOwn returns the Own field if non-nil, zero value otherwise.

### GetOwnOk

`func (o *O11yO11yFrame) GetOwnOk() (*bool, bool)`

GetOwnOk returns a tuple with the Own field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwn

`func (o *O11yO11yFrame) SetOwn(v bool)`

SetOwn sets Own field to given value.

### HasOwn

`func (o *O11yO11yFrame) HasOwn() bool`

HasOwn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



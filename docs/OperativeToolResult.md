# OperativeToolResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Output** | Pointer to **string** | Standard output text | [optional] 
**Error** | Pointer to **string** | Error message if the tool failed | [optional] 
**Base64Image** | Pointer to **string** | Base64-encoded PNG screenshot | [optional] 
**System** | Pointer to **string** | System message (e.g. tool restart notice) | [optional] 

## Methods

### NewOperativeToolResult

`func NewOperativeToolResult() *OperativeToolResult`

NewOperativeToolResult instantiates a new OperativeToolResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperativeToolResultWithDefaults

`func NewOperativeToolResultWithDefaults() *OperativeToolResult`

NewOperativeToolResultWithDefaults instantiates a new OperativeToolResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutput

`func (o *OperativeToolResult) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *OperativeToolResult) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *OperativeToolResult) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *OperativeToolResult) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetError

`func (o *OperativeToolResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OperativeToolResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OperativeToolResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *OperativeToolResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetBase64Image

`func (o *OperativeToolResult) GetBase64Image() string`

GetBase64Image returns the Base64Image field if non-nil, zero value otherwise.

### GetBase64ImageOk

`func (o *OperativeToolResult) GetBase64ImageOk() (*string, bool)`

GetBase64ImageOk returns a tuple with the Base64Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64Image

`func (o *OperativeToolResult) SetBase64Image(v string)`

SetBase64Image sets Base64Image field to given value.

### HasBase64Image

`func (o *OperativeToolResult) HasBase64Image() bool`

HasBase64Image returns a boolean if a field has been set.

### GetSystem

`func (o *OperativeToolResult) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *OperativeToolResult) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *OperativeToolResult) SetSystem(v string)`

SetSystem sets System field to given value.

### HasSystem

`func (o *OperativeToolResult) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



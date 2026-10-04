# FunctionLogLines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to **string** | Logs is that run&#39;s error text when it failed, else its output. It is empty when the function has never run. | [optional] 

## Methods

### NewFunctionLogLines

`func NewFunctionLogLines() *FunctionLogLines`

NewFunctionLogLines instantiates a new FunctionLogLines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionLogLinesWithDefaults

`func NewFunctionLogLinesWithDefaults() *FunctionLogLines`

NewFunctionLogLinesWithDefaults instantiates a new FunctionLogLines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *FunctionLogLines) GetLogs() string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *FunctionLogLines) GetLogsOk() (*string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *FunctionLogLines) SetLogs(v string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *FunctionLogLines) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



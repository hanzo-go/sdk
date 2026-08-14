# RunResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error is the connector-level failure message when not ok. | [optional] 
**Ok** | Pointer to **bool** | Ok reports whether the action ran to completion. | [optional] 
**Output** | Pointer to **map[string]interface{}** | Output is the action&#39;s result when ok. Its shape is the action&#39;s own. | [optional] 

## Methods

### NewRunResp

`func NewRunResp() *RunResp`

NewRunResp instantiates a new RunResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunRespWithDefaults

`func NewRunRespWithDefaults() *RunResp`

NewRunRespWithDefaults instantiates a new RunResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *RunResp) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RunResp) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RunResp) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RunResp) HasError() bool`

HasError returns a boolean if a field has been set.

### GetOk

`func (o *RunResp) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *RunResp) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *RunResp) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *RunResp) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetOutput

`func (o *RunResp) GetOutput() map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *RunResp) GetOutputOk() (*map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *RunResp) SetOutput(v map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *RunResp) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



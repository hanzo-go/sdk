# AutoRunResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error is the provider-level failure message when not ok. | [optional] 
**Ok** | Pointer to **bool** | Ok reports whether the action ran to completion. | [optional] 
**Output** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewAutoRunResp

`func NewAutoRunResp() *AutoRunResp`

NewAutoRunResp instantiates a new AutoRunResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoRunRespWithDefaults

`func NewAutoRunRespWithDefaults() *AutoRunResp`

NewAutoRunRespWithDefaults instantiates a new AutoRunResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *AutoRunResp) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AutoRunResp) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AutoRunResp) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AutoRunResp) HasError() bool`

HasError returns a boolean if a field has been set.

### GetOk

`func (o *AutoRunResp) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *AutoRunResp) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *AutoRunResp) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *AutoRunResp) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetOutput

`func (o *AutoRunResp) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *AutoRunResp) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *AutoRunResp) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *AutoRunResp) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *AutoRunResp) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *AutoRunResp) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



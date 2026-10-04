# EvalItemReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedOutput** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** | ID makes the write idempotent — re-posting the same id replaces that example in place. Omit it and one is generated. An id that already exists in a DIFFERENT dataset is 409 rather than a move. | [optional] 
**Input** | Pointer to **interface{}** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata is a free-form object stored with the example. | [optional] 
**Status** | Pointer to **string** | Status is ACTIVE (the default) or ARCHIVED. Only ACTIVE examples are fed to a run, which is how an example is retired without being deleted. | [optional] 

## Methods

### NewEvalItemReq

`func NewEvalItemReq() *EvalItemReq`

NewEvalItemReq instantiates a new EvalItemReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalItemReqWithDefaults

`func NewEvalItemReqWithDefaults() *EvalItemReq`

NewEvalItemReqWithDefaults instantiates a new EvalItemReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedOutput

`func (o *EvalItemReq) GetExpectedOutput() interface{}`

GetExpectedOutput returns the ExpectedOutput field if non-nil, zero value otherwise.

### GetExpectedOutputOk

`func (o *EvalItemReq) GetExpectedOutputOk() (*interface{}, bool)`

GetExpectedOutputOk returns a tuple with the ExpectedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedOutput

`func (o *EvalItemReq) SetExpectedOutput(v interface{})`

SetExpectedOutput sets ExpectedOutput field to given value.

### HasExpectedOutput

`func (o *EvalItemReq) HasExpectedOutput() bool`

HasExpectedOutput returns a boolean if a field has been set.

### SetExpectedOutputNil

`func (o *EvalItemReq) SetExpectedOutputNil(b bool)`

 SetExpectedOutputNil sets the value for ExpectedOutput to be an explicit nil

### UnsetExpectedOutput
`func (o *EvalItemReq) UnsetExpectedOutput()`

UnsetExpectedOutput ensures that no value is present for ExpectedOutput, not even an explicit nil
### GetId

`func (o *EvalItemReq) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EvalItemReq) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EvalItemReq) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EvalItemReq) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInput

`func (o *EvalItemReq) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *EvalItemReq) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *EvalItemReq) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *EvalItemReq) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *EvalItemReq) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *EvalItemReq) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetMetadata

`func (o *EvalItemReq) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EvalItemReq) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EvalItemReq) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EvalItemReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStatus

`func (o *EvalItemReq) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EvalItemReq) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EvalItemReq) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EvalItemReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



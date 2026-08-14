# ItemReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedOutput** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** | ID makes the write idempotent — re-posting the same id replaces that example in place. Omit it and one is generated. An id that already exists in a DIFFERENT dataset is 409 rather than a move. | [optional] 
**Input** | Pointer to **interface{}** |  | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata is a free-form object stored with the example. | [optional] 
**Status** | Pointer to **string** | Status is ACTIVE (the default) or ARCHIVED. Only ACTIVE examples are fed to a run, which is how an example is retired without being deleted. | [optional] 

## Methods

### NewItemReq

`func NewItemReq() *ItemReq`

NewItemReq instantiates a new ItemReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemReqWithDefaults

`func NewItemReqWithDefaults() *ItemReq`

NewItemReqWithDefaults instantiates a new ItemReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedOutput

`func (o *ItemReq) GetExpectedOutput() interface{}`

GetExpectedOutput returns the ExpectedOutput field if non-nil, zero value otherwise.

### GetExpectedOutputOk

`func (o *ItemReq) GetExpectedOutputOk() (*interface{}, bool)`

GetExpectedOutputOk returns a tuple with the ExpectedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedOutput

`func (o *ItemReq) SetExpectedOutput(v interface{})`

SetExpectedOutput sets ExpectedOutput field to given value.

### HasExpectedOutput

`func (o *ItemReq) HasExpectedOutput() bool`

HasExpectedOutput returns a boolean if a field has been set.

### SetExpectedOutputNil

`func (o *ItemReq) SetExpectedOutputNil(b bool)`

 SetExpectedOutputNil sets the value for ExpectedOutput to be an explicit nil

### UnsetExpectedOutput
`func (o *ItemReq) UnsetExpectedOutput()`

UnsetExpectedOutput ensures that no value is present for ExpectedOutput, not even an explicit nil
### GetId

`func (o *ItemReq) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemReq) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemReq) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemReq) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInput

`func (o *ItemReq) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ItemReq) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ItemReq) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ItemReq) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *ItemReq) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *ItemReq) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetMetadata

`func (o *ItemReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ItemReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ItemReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ItemReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStatus

`func (o *ItemReq) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ItemReq) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ItemReq) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ItemReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



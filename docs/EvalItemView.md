# EvalItemView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is when the example was first written. | [optional] 
**DatasetName** | Pointer to **string** | Dataset is the set this example belongs to. | [optional] 
**ExpectedOutput** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** | ID is the example&#39;s handle, unique within the caller&#39;s org. | [optional] 
**Input** | Pointer to **interface{}** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata is the free-form object stored with the example. | [optional] 
**Status** | Pointer to **string** | Status is ACTIVE or ARCHIVED. Only ACTIVE examples are fed to a run, which is how one is retired without being deleted. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when it last changed. | [optional] 

## Methods

### NewEvalItemView

`func NewEvalItemView() *EvalItemView`

NewEvalItemView instantiates a new EvalItemView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalItemViewWithDefaults

`func NewEvalItemViewWithDefaults() *EvalItemView`

NewEvalItemViewWithDefaults instantiates a new EvalItemView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EvalItemView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EvalItemView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EvalItemView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EvalItemView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDatasetName

`func (o *EvalItemView) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *EvalItemView) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *EvalItemView) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.

### HasDatasetName

`func (o *EvalItemView) HasDatasetName() bool`

HasDatasetName returns a boolean if a field has been set.

### GetExpectedOutput

`func (o *EvalItemView) GetExpectedOutput() interface{}`

GetExpectedOutput returns the ExpectedOutput field if non-nil, zero value otherwise.

### GetExpectedOutputOk

`func (o *EvalItemView) GetExpectedOutputOk() (*interface{}, bool)`

GetExpectedOutputOk returns a tuple with the ExpectedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedOutput

`func (o *EvalItemView) SetExpectedOutput(v interface{})`

SetExpectedOutput sets ExpectedOutput field to given value.

### HasExpectedOutput

`func (o *EvalItemView) HasExpectedOutput() bool`

HasExpectedOutput returns a boolean if a field has been set.

### SetExpectedOutputNil

`func (o *EvalItemView) SetExpectedOutputNil(b bool)`

 SetExpectedOutputNil sets the value for ExpectedOutput to be an explicit nil

### UnsetExpectedOutput
`func (o *EvalItemView) UnsetExpectedOutput()`

UnsetExpectedOutput ensures that no value is present for ExpectedOutput, not even an explicit nil
### GetId

`func (o *EvalItemView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EvalItemView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EvalItemView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EvalItemView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInput

`func (o *EvalItemView) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *EvalItemView) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *EvalItemView) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *EvalItemView) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *EvalItemView) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *EvalItemView) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetMetadata

`func (o *EvalItemView) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EvalItemView) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EvalItemView) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EvalItemView) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStatus

`func (o *EvalItemView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EvalItemView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EvalItemView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EvalItemView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EvalItemView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EvalItemView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EvalItemView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EvalItemView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ItemView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is when the example was first written. | [optional] 
**DatasetName** | Pointer to **string** | Dataset is the set this example belongs to. | [optional] 
**ExpectedOutput** | Pointer to **map[string]interface{}** | Expected is the answer a correct model produces, which the judge grades against. | [optional] 
**Id** | Pointer to **string** | ID is the example&#39;s handle, unique within the caller&#39;s org. | [optional] 
**Input** | Pointer to **map[string]interface{}** | Input is what the model under test is given, as it was written. | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata is the free-form object stored with the example. | [optional] 
**Status** | Pointer to **string** | Status is ACTIVE or ARCHIVED. Only ACTIVE examples are fed to a run, which is how one is retired without being deleted. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when it last changed. | [optional] 

## Methods

### NewItemView

`func NewItemView() *ItemView`

NewItemView instantiates a new ItemView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemViewWithDefaults

`func NewItemViewWithDefaults() *ItemView`

NewItemViewWithDefaults instantiates a new ItemView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ItemView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ItemView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ItemView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ItemView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDatasetName

`func (o *ItemView) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *ItemView) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *ItemView) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.

### HasDatasetName

`func (o *ItemView) HasDatasetName() bool`

HasDatasetName returns a boolean if a field has been set.

### GetExpectedOutput

`func (o *ItemView) GetExpectedOutput() map[string]interface{}`

GetExpectedOutput returns the ExpectedOutput field if non-nil, zero value otherwise.

### GetExpectedOutputOk

`func (o *ItemView) GetExpectedOutputOk() (*map[string]interface{}, bool)`

GetExpectedOutputOk returns a tuple with the ExpectedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedOutput

`func (o *ItemView) SetExpectedOutput(v map[string]interface{})`

SetExpectedOutput sets ExpectedOutput field to given value.

### HasExpectedOutput

`func (o *ItemView) HasExpectedOutput() bool`

HasExpectedOutput returns a boolean if a field has been set.

### GetId

`func (o *ItemView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInput

`func (o *ItemView) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ItemView) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ItemView) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ItemView) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetMetadata

`func (o *ItemView) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ItemView) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ItemView) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ItemView) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStatus

`func (o *ItemView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ItemView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ItemView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ItemView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ItemView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ItemView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ItemView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ItemView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



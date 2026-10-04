# FlowFlowUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **interface{}** |  | [optional] 
**Description** | Pointer to **string** | Description replaces the description when present. | [optional] 
**Locked** | Pointer to **bool** | Locked freezes or unfreezes the workflow against edits when present. | [optional] 
**Name** | Pointer to **string** | Name renames the workflow when present. | [optional] 
**Workflow** | Pointer to **string** | Workflow is the workflow&#39;s UUID, taken from the path. | [optional] 

## Methods

### NewFlowFlowUpdate

`func NewFlowFlowUpdate() *FlowFlowUpdate`

NewFlowFlowUpdate instantiates a new FlowFlowUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowFlowUpdateWithDefaults

`func NewFlowFlowUpdateWithDefaults() *FlowFlowUpdate`

NewFlowFlowUpdateWithDefaults instantiates a new FlowFlowUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FlowFlowUpdate) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FlowFlowUpdate) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FlowFlowUpdate) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *FlowFlowUpdate) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *FlowFlowUpdate) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *FlowFlowUpdate) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDescription

`func (o *FlowFlowUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlowFlowUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlowFlowUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlowFlowUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocked

`func (o *FlowFlowUpdate) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *FlowFlowUpdate) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *FlowFlowUpdate) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *FlowFlowUpdate) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetName

`func (o *FlowFlowUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowFlowUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowFlowUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlowFlowUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWorkflow

`func (o *FlowFlowUpdate) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *FlowFlowUpdate) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *FlowFlowUpdate) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *FlowFlowUpdate) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



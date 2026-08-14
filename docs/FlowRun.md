# FlowRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to **string** | Input is the run&#39;s chat input value, handed to the graph&#39;s input node. | [optional] 
**Session** | Pointer to **string** | Session groups runs into one conversation; the product mints one when absent and returns it in the response. | [optional] 
**Tweaks** | Pointer to **interface{}** |  | [optional] 
**Workflow** | Pointer to **string** | Workflow is the UUID of the workflow to run. | [optional] 

## Methods

### NewFlowRun

`func NewFlowRun() *FlowRun`

NewFlowRun instantiates a new FlowRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowRunWithDefaults

`func NewFlowRunWithDefaults() *FlowRun`

NewFlowRunWithDefaults instantiates a new FlowRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *FlowRun) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *FlowRun) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *FlowRun) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *FlowRun) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetSession

`func (o *FlowRun) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *FlowRun) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *FlowRun) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *FlowRun) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetTweaks

`func (o *FlowRun) GetTweaks() interface{}`

GetTweaks returns the Tweaks field if non-nil, zero value otherwise.

### GetTweaksOk

`func (o *FlowRun) GetTweaksOk() (*interface{}, bool)`

GetTweaksOk returns a tuple with the Tweaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweaks

`func (o *FlowRun) SetTweaks(v interface{})`

SetTweaks sets Tweaks field to given value.

### HasTweaks

`func (o *FlowRun) HasTweaks() bool`

HasTweaks returns a boolean if a field has been set.

### SetTweaksNil

`func (o *FlowRun) SetTweaksNil(b bool)`

 SetTweaksNil sets the value for Tweaks to be an explicit nil

### UnsetTweaks
`func (o *FlowRun) UnsetTweaks()`

UnsetTweaks ensures that no value is present for Tweaks, not even an explicit nil
### GetWorkflow

`func (o *FlowRun) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *FlowRun) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *FlowRun) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *FlowRun) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



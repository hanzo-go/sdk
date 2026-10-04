# AgentActivityView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** | agent name | [optional] 
**At** | Pointer to **string** | RFC3339 UTC | [optional] 
**Id** | Pointer to **string** | ID identifies the event, and its shape says which kind it is: a run event carries the run&#39;s own id, while an agent event is the agent id suffixed \&quot;:created\&quot; or \&quot;:updated\&quot;. Unique within a feed, and not an address — there is nothing to fetch it by. | [optional] 
**Kind** | Pointer to **string** | invoked|failed|created|updated (from real events) | [optional] 
**Message** | Pointer to **string** | Message is the line to render, already bounded: \&quot;Invoked &lt;model&gt;\&quot; for a run that worked, the run&#39;s own error truncated to 200 characters for one that did not (or \&quot;Run failed\&quot; when it said nothing), and a fixed phrase for the two agent events. Nothing here is invented — every event is a row that exists. | [optional] 

## Methods

### NewAgentActivityView

`func NewAgentActivityView() *AgentActivityView`

NewAgentActivityView instantiates a new AgentActivityView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityViewWithDefaults

`func NewAgentActivityViewWithDefaults() *AgentActivityView`

NewAgentActivityViewWithDefaults instantiates a new AgentActivityView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *AgentActivityView) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentActivityView) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentActivityView) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *AgentActivityView) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetAt

`func (o *AgentActivityView) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *AgentActivityView) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *AgentActivityView) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *AgentActivityView) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetId

`func (o *AgentActivityView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentActivityView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentActivityView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentActivityView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *AgentActivityView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AgentActivityView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AgentActivityView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AgentActivityView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMessage

`func (o *AgentActivityView) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AgentActivityView) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AgentActivityView) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AgentActivityView) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



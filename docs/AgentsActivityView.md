# AgentsActivityView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Agent** | Pointer to **string** | Agent name. | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**At** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAgentsActivityView

`func NewAgentsActivityView() *AgentsActivityView`

NewAgentsActivityView instantiates a new AgentsActivityView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsActivityViewWithDefaults

`func NewAgentsActivityViewWithDefaults() *AgentsActivityView`

NewAgentsActivityViewWithDefaults instantiates a new AgentsActivityView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentsActivityView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentsActivityView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentsActivityView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentsActivityView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *AgentsActivityView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AgentsActivityView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AgentsActivityView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AgentsActivityView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAgent

`func (o *AgentsActivityView) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentsActivityView) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentsActivityView) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *AgentsActivityView) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetMessage

`func (o *AgentsActivityView) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AgentsActivityView) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AgentsActivityView) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AgentsActivityView) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetAt

`func (o *AgentsActivityView) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *AgentsActivityView) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *AgentsActivityView) SetAt(v time.Time)`

SetAt sets At field to given value.

### HasAt

`func (o *AgentsActivityView) HasAt() bool`

HasAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



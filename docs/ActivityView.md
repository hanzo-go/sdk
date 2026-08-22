# ActivityView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** | agent name | [optional] 
**At** | Pointer to **string** | RFC3339 UTC | [optional] 
**Id** | Pointer to **string** | ID identifies the event, and its shape says which kind it is: a run event carries the run&#39;s own id, while an agent event is the agent id suffixed \&quot;:created\&quot; or \&quot;:updated\&quot;. Unique within a feed, and not an address — there is nothing to fetch it by. | [optional] 
**Kind** | Pointer to **string** | invoked|failed|created|updated (from real events) | [optional] 
**Message** | Pointer to **string** | Message is the line to render, already bounded: \&quot;Invoked &lt;model&gt;\&quot; for a run that worked, the run&#39;s own error truncated to 200 characters for one that did not (or \&quot;Run failed\&quot; when it said nothing), and a fixed phrase for the two agent events. Nothing here is invented — every event is a row that exists. | [optional] 

## Methods

### NewActivityView

`func NewActivityView() *ActivityView`

NewActivityView instantiates a new ActivityView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityViewWithDefaults

`func NewActivityViewWithDefaults() *ActivityView`

NewActivityViewWithDefaults instantiates a new ActivityView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *ActivityView) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *ActivityView) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *ActivityView) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *ActivityView) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetAt

`func (o *ActivityView) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *ActivityView) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *ActivityView) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *ActivityView) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetId

`func (o *ActivityView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ActivityView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ActivityView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ActivityView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ActivityView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMessage

`func (o *ActivityView) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ActivityView) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ActivityView) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ActivityView) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



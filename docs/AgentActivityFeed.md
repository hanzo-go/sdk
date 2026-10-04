# AgentActivityFeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to [**[]AgentActivityView**](AgentActivityView.md) | Activity is the merged run/create/update events, newest first, capped at 50. | [optional] 

## Methods

### NewAgentActivityFeed

`func NewAgentActivityFeed() *AgentActivityFeed`

NewAgentActivityFeed instantiates a new AgentActivityFeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityFeedWithDefaults

`func NewAgentActivityFeedWithDefaults() *AgentActivityFeed`

NewAgentActivityFeedWithDefaults instantiates a new AgentActivityFeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *AgentActivityFeed) GetActivity() []AgentActivityView`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *AgentActivityFeed) GetActivityOk() (*[]AgentActivityView, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *AgentActivityFeed) SetActivity(v []AgentActivityView)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *AgentActivityFeed) HasActivity() bool`

HasActivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



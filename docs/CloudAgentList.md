# CloudAgentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**[]CloudAgentView**](CloudAgentView.md) | Agents is the org&#39;s agents, each carrying its recorded run count. | [optional] 

## Methods

### NewCloudAgentList

`func NewCloudAgentList() *CloudAgentList`

NewCloudAgentList instantiates a new CloudAgentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAgentListWithDefaults

`func NewCloudAgentListWithDefaults() *CloudAgentList`

NewCloudAgentListWithDefaults instantiates a new CloudAgentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *CloudAgentList) GetAgents() []CloudAgentView`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *CloudAgentList) GetAgentsOk() (*[]CloudAgentView, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *CloudAgentList) SetAgents(v []CloudAgentView)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *CloudAgentList) HasAgents() bool`

HasAgents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



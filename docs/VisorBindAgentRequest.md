# VisorBindAgentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentName** | **string** |  | 
**BotVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewVisorBindAgentRequest

`func NewVisorBindAgentRequest(agentName string, ) *VisorBindAgentRequest`

NewVisorBindAgentRequest instantiates a new VisorBindAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisorBindAgentRequestWithDefaults

`func NewVisorBindAgentRequestWithDefaults() *VisorBindAgentRequest`

NewVisorBindAgentRequestWithDefaults instantiates a new VisorBindAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentName

`func (o *VisorBindAgentRequest) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *VisorBindAgentRequest) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *VisorBindAgentRequest) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.


### GetBotVersion

`func (o *VisorBindAgentRequest) GetBotVersion() string`

GetBotVersion returns the BotVersion field if non-nil, zero value otherwise.

### GetBotVersionOk

`func (o *VisorBindAgentRequest) GetBotVersionOk() (*string, bool)`

GetBotVersionOk returns a tuple with the BotVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotVersion

`func (o *VisorBindAgentRequest) SetBotVersion(v string)`

SetBotVersion sets BotVersion field to given value.

### HasBotVersion

`func (o *VisorBindAgentRequest) HasBotVersion() bool`

HasBotVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



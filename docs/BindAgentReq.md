# BindAgentReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentName** | Pointer to **string** | AgentName is the cloud Agent (/v1/agents) the machine will run. Required. | [optional] 
**BotVersion** | Pointer to **string** | BotVersion pins the @hanzo/bot runtime version; empty takes the default. | [optional] 
**Id** | Pointer to **string** | ID is the machine to bind, from the URL path. | [optional] 

## Methods

### NewBindAgentReq

`func NewBindAgentReq() *BindAgentReq`

NewBindAgentReq instantiates a new BindAgentReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindAgentReqWithDefaults

`func NewBindAgentReqWithDefaults() *BindAgentReq`

NewBindAgentReqWithDefaults instantiates a new BindAgentReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentName

`func (o *BindAgentReq) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *BindAgentReq) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *BindAgentReq) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.

### HasAgentName

`func (o *BindAgentReq) HasAgentName() bool`

HasAgentName returns a boolean if a field has been set.

### GetBotVersion

`func (o *BindAgentReq) GetBotVersion() string`

GetBotVersion returns the BotVersion field if non-nil, zero value otherwise.

### GetBotVersionOk

`func (o *BindAgentReq) GetBotVersionOk() (*string, bool)`

GetBotVersionOk returns a tuple with the BotVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotVersion

`func (o *BindAgentReq) SetBotVersion(v string)`

SetBotVersion sets BotVersion field to given value.

### HasBotVersion

`func (o *BindAgentReq) HasBotVersion() bool`

HasBotVersion returns a boolean if a field has been set.

### GetId

`func (o *BindAgentReq) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BindAgentReq) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BindAgentReq) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BindAgentReq) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



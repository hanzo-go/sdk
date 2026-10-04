# ComputeBindAgentReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentName** | Pointer to **string** | AgentName is the cloud Agent (/v1/agent) the machine will run. Required. | [optional] 
**BotVersion** | Pointer to **string** | BotVersion pins the @hanzo/bot runtime version; empty takes the default. | [optional] 
**Id** | Pointer to **string** | ID is the machine to bind, from the URL path. | [optional] 

## Methods

### NewComputeBindAgentReq

`func NewComputeBindAgentReq() *ComputeBindAgentReq`

NewComputeBindAgentReq instantiates a new ComputeBindAgentReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeBindAgentReqWithDefaults

`func NewComputeBindAgentReqWithDefaults() *ComputeBindAgentReq`

NewComputeBindAgentReqWithDefaults instantiates a new ComputeBindAgentReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentName

`func (o *ComputeBindAgentReq) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *ComputeBindAgentReq) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *ComputeBindAgentReq) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.

### HasAgentName

`func (o *ComputeBindAgentReq) HasAgentName() bool`

HasAgentName returns a boolean if a field has been set.

### GetBotVersion

`func (o *ComputeBindAgentReq) GetBotVersion() string`

GetBotVersion returns the BotVersion field if non-nil, zero value otherwise.

### GetBotVersionOk

`func (o *ComputeBindAgentReq) GetBotVersionOk() (*string, bool)`

GetBotVersionOk returns a tuple with the BotVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotVersion

`func (o *ComputeBindAgentReq) SetBotVersion(v string)`

SetBotVersion sets BotVersion field to given value.

### HasBotVersion

`func (o *ComputeBindAgentReq) HasBotVersion() bool`

HasBotVersion returns a boolean if a field has been set.

### GetId

`func (o *ComputeBindAgentReq) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputeBindAgentReq) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputeBindAgentReq) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComputeBindAgentReq) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



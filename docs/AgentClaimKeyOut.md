# AgentClaimKeyOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimKey** | Pointer to **string** | ClaimKey is the capability itself. It is returned ONCE and never again — only its SHA-256 hash is stored — so a daemon that loses it mints a new one. | [optional] 
**TargetId** | Pointer to **string** | TargetID is the machine the key authenticates. | [optional] 

## Methods

### NewAgentClaimKeyOut

`func NewAgentClaimKeyOut() *AgentClaimKeyOut`

NewAgentClaimKeyOut instantiates a new AgentClaimKeyOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentClaimKeyOutWithDefaults

`func NewAgentClaimKeyOutWithDefaults() *AgentClaimKeyOut`

NewAgentClaimKeyOutWithDefaults instantiates a new AgentClaimKeyOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimKey

`func (o *AgentClaimKeyOut) GetClaimKey() string`

GetClaimKey returns the ClaimKey field if non-nil, zero value otherwise.

### GetClaimKeyOk

`func (o *AgentClaimKeyOut) GetClaimKeyOk() (*string, bool)`

GetClaimKeyOk returns a tuple with the ClaimKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimKey

`func (o *AgentClaimKeyOut) SetClaimKey(v string)`

SetClaimKey sets ClaimKey field to given value.

### HasClaimKey

`func (o *AgentClaimKeyOut) HasClaimKey() bool`

HasClaimKey returns a boolean if a field has been set.

### GetTargetId

`func (o *AgentClaimKeyOut) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *AgentClaimKeyOut) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *AgentClaimKeyOut) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *AgentClaimKeyOut) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



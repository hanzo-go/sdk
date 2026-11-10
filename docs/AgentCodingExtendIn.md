# AgentCodingExtendIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seconds** | Pointer to **int64** | Seconds is how much to add to the time the run is held to. It is added to what the run has left, up to the longest a run may be given, so asking for more than that is answered with what the run got. | [optional] 
**Session** | Pointer to **string** | Session is the run&#39;s handle, from the path. | [optional] 

## Methods

### NewAgentCodingExtendIn

`func NewAgentCodingExtendIn() *AgentCodingExtendIn`

NewAgentCodingExtendIn instantiates a new AgentCodingExtendIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCodingExtendInWithDefaults

`func NewAgentCodingExtendInWithDefaults() *AgentCodingExtendIn`

NewAgentCodingExtendInWithDefaults instantiates a new AgentCodingExtendIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeconds

`func (o *AgentCodingExtendIn) GetSeconds() int64`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *AgentCodingExtendIn) GetSecondsOk() (*int64, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *AgentCodingExtendIn) SetSeconds(v int64)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *AgentCodingExtendIn) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.

### GetSession

`func (o *AgentCodingExtendIn) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *AgentCodingExtendIn) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *AgentCodingExtendIn) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *AgentCodingExtendIn) HasSession() bool`

HasSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



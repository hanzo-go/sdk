# AgentCodingExtended

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seconds** | Pointer to **int64** | Seconds is how long the run has left. | [optional] 
**Until** | Pointer to **string** | Until is when the run is ended for time now, as an RFC 3339 time. | [optional] 

## Methods

### NewAgentCodingExtended

`func NewAgentCodingExtended() *AgentCodingExtended`

NewAgentCodingExtended instantiates a new AgentCodingExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCodingExtendedWithDefaults

`func NewAgentCodingExtendedWithDefaults() *AgentCodingExtended`

NewAgentCodingExtendedWithDefaults instantiates a new AgentCodingExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeconds

`func (o *AgentCodingExtended) GetSeconds() int64`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *AgentCodingExtended) GetSecondsOk() (*int64, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *AgentCodingExtended) SetSeconds(v int64)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *AgentCodingExtended) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.

### GetUntil

`func (o *AgentCodingExtended) GetUntil() string`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *AgentCodingExtended) GetUntilOk() (*string, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *AgentCodingExtended) SetUntil(v string)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *AgentCodingExtended) HasUntil() bool`

HasUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



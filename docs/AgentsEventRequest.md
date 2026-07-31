# AgentsEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Actor** | Pointer to **string** | Defaults to the validated principal (org/sub). | [optional] 
**Payload** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewAgentsEventRequest

`func NewAgentsEventRequest(kind string, ) *AgentsEventRequest`

NewAgentsEventRequest instantiates a new AgentsEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsEventRequestWithDefaults

`func NewAgentsEventRequestWithDefaults() *AgentsEventRequest`

NewAgentsEventRequestWithDefaults instantiates a new AgentsEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AgentsEventRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AgentsEventRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AgentsEventRequest) SetKind(v string)`

SetKind sets Kind field to given value.


### GetActor

`func (o *AgentsEventRequest) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AgentsEventRequest) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AgentsEventRequest) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *AgentsEventRequest) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetPayload

`func (o *AgentsEventRequest) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *AgentsEventRequest) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *AgentsEventRequest) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *AgentsEventRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *AgentsEventRequest) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *AgentsEventRequest) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



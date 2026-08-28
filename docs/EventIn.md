# EventIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** | Actor is who produced the turn. Empty takes the validated caller, which is what an agent writing its own transcript wants; naming one is for a surface recording on somebody else&#39;s behalf. | [optional] 
**Id** | Pointer to **string** | ID is the session to append to, from the path. | [optional] 
**Kind** | Pointer to **string** | Kind is what this turn IS: message, tool-call, spawn, log, status, control or progress. Anything else is refused — the vocabulary is closed so a reader can branch on it. | [optional] 
**Payload** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewEventIn

`func NewEventIn() *EventIn`

NewEventIn instantiates a new EventIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventInWithDefaults

`func NewEventInWithDefaults() *EventIn`

NewEventInWithDefaults instantiates a new EventIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *EventIn) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *EventIn) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *EventIn) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *EventIn) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetId

`func (o *EventIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *EventIn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EventIn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EventIn) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EventIn) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPayload

`func (o *EventIn) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *EventIn) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *EventIn) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *EventIn) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *EventIn) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *EventIn) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



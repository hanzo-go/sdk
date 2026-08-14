# EventView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **interface{}** |  | [optional] 
**Seq** | Pointer to **int32** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 

## Methods

### NewEventView

`func NewEventView() *EventView`

NewEventView instantiates a new EventView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventViewWithDefaults

`func NewEventViewWithDefaults() *EventView`

NewEventViewWithDefaults instantiates a new EventView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *EventView) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *EventView) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *EventView) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *EventView) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EventView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EventView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *EventView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *EventView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EventView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EventView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EventView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPayload

`func (o *EventView) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *EventView) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *EventView) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *EventView) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *EventView) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *EventView) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetSeq

`func (o *EventView) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *EventView) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *EventView) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *EventView) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetSessionId

`func (o *EventView) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *EventView) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *EventView) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *EventView) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



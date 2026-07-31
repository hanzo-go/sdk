# CloudAgentsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**Seq** | Pointer to **int64** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Actor** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **map[string]interface{}** | Opaque, well-formed, size-bounded JSON blob. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCloudAgentsEvent

`func NewCloudAgentsEvent() *CloudAgentsEvent`

NewCloudAgentsEvent instantiates a new CloudAgentsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAgentsEventWithDefaults

`func NewCloudAgentsEventWithDefaults() *CloudAgentsEvent`

NewCloudAgentsEventWithDefaults instantiates a new CloudAgentsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudAgentsEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAgentsEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAgentsEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAgentsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSessionId

`func (o *CloudAgentsEvent) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CloudAgentsEvent) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CloudAgentsEvent) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CloudAgentsEvent) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSeq

`func (o *CloudAgentsEvent) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *CloudAgentsEvent) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *CloudAgentsEvent) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *CloudAgentsEvent) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetKind

`func (o *CloudAgentsEvent) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudAgentsEvent) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudAgentsEvent) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudAgentsEvent) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetActor

`func (o *CloudAgentsEvent) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *CloudAgentsEvent) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *CloudAgentsEvent) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *CloudAgentsEvent) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetPayload

`func (o *CloudAgentsEvent) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CloudAgentsEvent) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CloudAgentsEvent) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CloudAgentsEvent) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudAgentsEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudAgentsEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudAgentsEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudAgentsEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



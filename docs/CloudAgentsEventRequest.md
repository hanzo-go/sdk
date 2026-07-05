# CloudAgentsEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Actor** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **map[string]interface{}** | Opaque JSON blob (validated well-formed, size-bounded). | [optional] 

## Methods

### NewCloudAgentsEventRequest

`func NewCloudAgentsEventRequest(kind string, ) *CloudAgentsEventRequest`

NewCloudAgentsEventRequest instantiates a new CloudAgentsEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAgentsEventRequestWithDefaults

`func NewCloudAgentsEventRequestWithDefaults() *CloudAgentsEventRequest`

NewCloudAgentsEventRequestWithDefaults instantiates a new CloudAgentsEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CloudAgentsEventRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudAgentsEventRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudAgentsEventRequest) SetKind(v string)`

SetKind sets Kind field to given value.


### GetActor

`func (o *CloudAgentsEventRequest) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *CloudAgentsEventRequest) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *CloudAgentsEventRequest) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *CloudAgentsEventRequest) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetPayload

`func (o *CloudAgentsEventRequest) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CloudAgentsEventRequest) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CloudAgentsEventRequest) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CloudAgentsEventRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# EventProductEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinctId** | Pointer to **string** | DistinctID is the person/visitor the event is attributed to. | [optional] 
**Event** | Pointer to **string** | Event is the event name, e.g. page_viewed or signup_completed. | [optional] 
**Id** | Pointer to **string** | ID is the row&#39;s stable event id — the client&#39;s own idempotency id when it sent one, else the server-minted one. | [optional] 
**Path** | Pointer to **string** | Path is the URL&#39;s path component, the key the topPages lens groups by. | [optional] 
**Product** | Pointer to **string** | Product is the surface that emitted the event. Omitted when absent. | [optional] 
**Properties** | Pointer to **interface{}** |  | [optional] 
**SessionId** | Pointer to **string** | SessionID groups the events of one visit. Omitted when the client sent none. | [optional] 
**Timestamp** | Pointer to **string** | Timestamp is when the event happened, RFC3339 UTC. | [optional] 
**Type** | Pointer to **string** | Type is the row&#39;s kind — the plane&#39;s discriminator: page, track, identify or group. (Errors are not here at all: they land on event.error and are read at /v1/event/errors.) | [optional] 
**Url** | Pointer to **string** | URL is the full page address the event fired on. Omitted when absent. | [optional] 

## Methods

### NewEventProductEvent

`func NewEventProductEvent() *EventProductEvent`

NewEventProductEvent instantiates a new EventProductEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventProductEventWithDefaults

`func NewEventProductEventWithDefaults() *EventProductEvent`

NewEventProductEventWithDefaults instantiates a new EventProductEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinctId

`func (o *EventProductEvent) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *EventProductEvent) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *EventProductEvent) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *EventProductEvent) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### GetEvent

`func (o *EventProductEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventProductEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventProductEvent) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *EventProductEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetId

`func (o *EventProductEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventProductEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventProductEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventProductEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *EventProductEvent) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EventProductEvent) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EventProductEvent) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *EventProductEvent) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProduct

`func (o *EventProductEvent) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *EventProductEvent) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *EventProductEvent) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *EventProductEvent) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProperties

`func (o *EventProductEvent) GetProperties() interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EventProductEvent) GetPropertiesOk() (*interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EventProductEvent) SetProperties(v interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *EventProductEvent) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *EventProductEvent) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *EventProductEvent) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetSessionId

`func (o *EventProductEvent) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *EventProductEvent) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *EventProductEvent) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *EventProductEvent) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetTimestamp

`func (o *EventProductEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventProductEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventProductEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EventProductEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *EventProductEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventProductEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventProductEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventProductEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *EventProductEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EventProductEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EventProductEvent) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EventProductEvent) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



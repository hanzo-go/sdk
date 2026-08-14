# ProductEvent

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
**Type** | Pointer to **string** | Type is the row&#39;s kind — the plane&#39;s discriminator: page, track, identify or group. (Errors are not here at all: they land on event.error and are read at /v1/errors.) | [optional] 
**Url** | Pointer to **string** | URL is the full page address the event fired on. Omitted when absent. | [optional] 

## Methods

### NewProductEvent

`func NewProductEvent() *ProductEvent`

NewProductEvent instantiates a new ProductEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductEventWithDefaults

`func NewProductEventWithDefaults() *ProductEvent`

NewProductEventWithDefaults instantiates a new ProductEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinctId

`func (o *ProductEvent) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *ProductEvent) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *ProductEvent) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *ProductEvent) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### GetEvent

`func (o *ProductEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ProductEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ProductEvent) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *ProductEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetId

`func (o *ProductEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *ProductEvent) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ProductEvent) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ProductEvent) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ProductEvent) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProduct

`func (o *ProductEvent) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ProductEvent) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ProductEvent) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ProductEvent) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProperties

`func (o *ProductEvent) GetProperties() interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ProductEvent) GetPropertiesOk() (*interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ProductEvent) SetProperties(v interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ProductEvent) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *ProductEvent) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *ProductEvent) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetSessionId

`func (o *ProductEvent) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ProductEvent) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ProductEvent) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ProductEvent) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ProductEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProductEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProductEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProductEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *ProductEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProductEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *ProductEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProductEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProductEvent) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProductEvent) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



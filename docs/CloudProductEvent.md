# CloudProductEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinctId** | Pointer to **string** | DistinctID is the person/visitor the event is attributed to. | [optional] 
**Event** | Pointer to **string** | Event is the event name, e.g. $pageview. | [optional] 
**Id** | Pointer to **string** | ID is the row&#39;s stable event id — the client&#39;s own idempotency id when it sent one, else the server-minted one. | [optional] 
**Path** | Pointer to **string** | Path is the URL&#39;s path component, the key the topPages lens groups by. | [optional] 
**Product** | Pointer to **string** | Product is the surface that emitted the event. Omitted when absent. | [optional] 
**Properties** | Pointer to **interface{}** |  | [optional] 
**SessionId** | Pointer to **string** | SessionID groups the events of one visit. Omitted when the client sent none. | [optional] 
**Timestamp** | Pointer to **string** | Timestamp is when the event happened, RFC3339 UTC. | [optional] 
**Type** | Pointer to **string** | Type is the canonical kind: pageview, error, identify, group or event. | [optional] 
**Url** | Pointer to **string** | URL is the full page address the event fired on. Omitted when absent. | [optional] 

## Methods

### NewCloudProductEvent

`func NewCloudProductEvent() *CloudProductEvent`

NewCloudProductEvent instantiates a new CloudProductEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProductEventWithDefaults

`func NewCloudProductEventWithDefaults() *CloudProductEvent`

NewCloudProductEventWithDefaults instantiates a new CloudProductEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinctId

`func (o *CloudProductEvent) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *CloudProductEvent) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *CloudProductEvent) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *CloudProductEvent) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### GetEvent

`func (o *CloudProductEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CloudProductEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CloudProductEvent) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *CloudProductEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetId

`func (o *CloudProductEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudProductEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudProductEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudProductEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *CloudProductEvent) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CloudProductEvent) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CloudProductEvent) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CloudProductEvent) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProduct

`func (o *CloudProductEvent) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CloudProductEvent) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CloudProductEvent) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CloudProductEvent) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProperties

`func (o *CloudProductEvent) GetProperties() interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CloudProductEvent) GetPropertiesOk() (*interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CloudProductEvent) SetProperties(v interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CloudProductEvent) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *CloudProductEvent) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *CloudProductEvent) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetSessionId

`func (o *CloudProductEvent) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CloudProductEvent) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CloudProductEvent) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CloudProductEvent) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CloudProductEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CloudProductEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CloudProductEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CloudProductEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *CloudProductEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudProductEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudProductEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudProductEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *CloudProductEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudProductEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudProductEvent) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CloudProductEvent) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



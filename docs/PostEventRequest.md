# PostEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinctId** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Batch** | Pointer to [**[]InsightsEvent**](InsightsEvent.md) |  | [optional] 
**Events** | Pointer to [**[]CaptureEvent**](CaptureEvent.md) |  | [optional] 
**DistinctId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewPostEventRequest

`func NewPostEventRequest() *PostEventRequest`

NewPostEventRequest instantiates a new PostEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostEventRequestWithDefaults

`func NewPostEventRequestWithDefaults() *PostEventRequest`

NewPostEventRequestWithDefaults instantiates a new PostEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinctId

`func (o *PostEventRequest) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *PostEventRequest) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *PostEventRequest) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *PostEventRequest) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### GetEvent

`func (o *PostEventRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PostEventRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PostEventRequest) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *PostEventRequest) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetProperties

`func (o *PostEventRequest) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PostEventRequest) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PostEventRequest) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PostEventRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTime

`func (o *PostEventRequest) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PostEventRequest) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PostEventRequest) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *PostEventRequest) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetType

`func (o *PostEventRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostEventRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostEventRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostEventRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBatch

`func (o *PostEventRequest) GetBatch() []InsightsEvent`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *PostEventRequest) GetBatchOk() (*[]InsightsEvent, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *PostEventRequest) SetBatch(v []InsightsEvent)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *PostEventRequest) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### GetEvents

`func (o *PostEventRequest) GetEvents() []CaptureEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *PostEventRequest) GetEventsOk() (*[]CaptureEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *PostEventRequest) SetEvents(v []CaptureEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *PostEventRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetDistinctId

`func (o *PostEventRequest) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *PostEventRequest) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *PostEventRequest) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *PostEventRequest) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### GetTimestamp

`func (o *PostEventRequest) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PostEventRequest) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PostEventRequest) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PostEventRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUuid

`func (o *PostEventRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PostEventRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PostEventRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PostEventRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



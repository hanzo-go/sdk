# CloudPostV1EventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinctId** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**Batch** | Pointer to [**[]CloudInsightsEvent**](CloudInsightsEvent.md) |  | [optional] 
**Events** | Pointer to [**[]CloudCaptureEvent**](CloudCaptureEvent.md) |  | [optional] 
**DistinctId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudPostV1EventRequest

`func NewCloudPostV1EventRequest() *CloudPostV1EventRequest`

NewCloudPostV1EventRequest instantiates a new CloudPostV1EventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPostV1EventRequestWithDefaults

`func NewCloudPostV1EventRequestWithDefaults() *CloudPostV1EventRequest`

NewCloudPostV1EventRequestWithDefaults instantiates a new CloudPostV1EventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinctId

`func (o *CloudPostV1EventRequest) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *CloudPostV1EventRequest) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *CloudPostV1EventRequest) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *CloudPostV1EventRequest) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### GetEvent

`func (o *CloudPostV1EventRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CloudPostV1EventRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CloudPostV1EventRequest) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *CloudPostV1EventRequest) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetProperties

`func (o *CloudPostV1EventRequest) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CloudPostV1EventRequest) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CloudPostV1EventRequest) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CloudPostV1EventRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTime

`func (o *CloudPostV1EventRequest) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CloudPostV1EventRequest) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CloudPostV1EventRequest) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *CloudPostV1EventRequest) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetBatch

`func (o *CloudPostV1EventRequest) GetBatch() []CloudInsightsEvent`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *CloudPostV1EventRequest) GetBatchOk() (*[]CloudInsightsEvent, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *CloudPostV1EventRequest) SetBatch(v []CloudInsightsEvent)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *CloudPostV1EventRequest) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### GetEvents

`func (o *CloudPostV1EventRequest) GetEvents() []CloudCaptureEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CloudPostV1EventRequest) GetEventsOk() (*[]CloudCaptureEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CloudPostV1EventRequest) SetEvents(v []CloudCaptureEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CloudPostV1EventRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetDistinctId

`func (o *CloudPostV1EventRequest) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *CloudPostV1EventRequest) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *CloudPostV1EventRequest) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *CloudPostV1EventRequest) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CloudPostV1EventRequest) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CloudPostV1EventRequest) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CloudPostV1EventRequest) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CloudPostV1EventRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUuid

`func (o *CloudPostV1EventRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CloudPostV1EventRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CloudPostV1EventRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CloudPostV1EventRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



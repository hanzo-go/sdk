# InsightsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batch** | Pointer to [**[]InsightsEvent**](InsightsEvent.md) |  | [optional] 
**DistinctId** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewInsightsBody

`func NewInsightsBody() *InsightsBody`

NewInsightsBody instantiates a new InsightsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsBodyWithDefaults

`func NewInsightsBodyWithDefaults() *InsightsBody`

NewInsightsBodyWithDefaults instantiates a new InsightsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatch

`func (o *InsightsBody) GetBatch() []InsightsEvent`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *InsightsBody) GetBatchOk() (*[]InsightsEvent, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *InsightsBody) SetBatch(v []InsightsEvent)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *InsightsBody) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### GetDistinctId

`func (o *InsightsBody) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *InsightsBody) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *InsightsBody) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *InsightsBody) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### GetEvent

`func (o *InsightsBody) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *InsightsBody) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *InsightsBody) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *InsightsBody) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetProperties

`func (o *InsightsBody) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *InsightsBody) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *InsightsBody) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *InsightsBody) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTimestamp

`func (o *InsightsBody) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InsightsBody) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InsightsBody) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InsightsBody) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUuid

`func (o *InsightsBody) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InsightsBody) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InsightsBody) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *InsightsBody) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



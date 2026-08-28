# AiHistoryRetrain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**Events** | Pointer to **int32** |  | [optional] 
**GateBase** | Pointer to **float32** |  | [optional] 
**GateMetric** | Pointer to **string** |  | [optional] 
**GatePass** | Pointer to **bool** |  | [optional] 
**GateValue** | Pointer to **float32** |  | [optional] 
**HoldoutAccuracy** | Pointer to **float32** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**TrainedTime** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewAiHistoryRetrain

`func NewAiHistoryRetrain() *AiHistoryRetrain`

NewAiHistoryRetrain instantiates a new AiHistoryRetrain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiHistoryRetrainWithDefaults

`func NewAiHistoryRetrainWithDefaults() *AiHistoryRetrain`

NewAiHistoryRetrainWithDefaults instantiates a new AiHistoryRetrain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *AiHistoryRetrain) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AiHistoryRetrain) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AiHistoryRetrain) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AiHistoryRetrain) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetEvents

`func (o *AiHistoryRetrain) GetEvents() int32`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AiHistoryRetrain) GetEventsOk() (*int32, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AiHistoryRetrain) SetEvents(v int32)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AiHistoryRetrain) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetGateBase

`func (o *AiHistoryRetrain) GetGateBase() float32`

GetGateBase returns the GateBase field if non-nil, zero value otherwise.

### GetGateBaseOk

`func (o *AiHistoryRetrain) GetGateBaseOk() (*float32, bool)`

GetGateBaseOk returns a tuple with the GateBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateBase

`func (o *AiHistoryRetrain) SetGateBase(v float32)`

SetGateBase sets GateBase field to given value.

### HasGateBase

`func (o *AiHistoryRetrain) HasGateBase() bool`

HasGateBase returns a boolean if a field has been set.

### GetGateMetric

`func (o *AiHistoryRetrain) GetGateMetric() string`

GetGateMetric returns the GateMetric field if non-nil, zero value otherwise.

### GetGateMetricOk

`func (o *AiHistoryRetrain) GetGateMetricOk() (*string, bool)`

GetGateMetricOk returns a tuple with the GateMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateMetric

`func (o *AiHistoryRetrain) SetGateMetric(v string)`

SetGateMetric sets GateMetric field to given value.

### HasGateMetric

`func (o *AiHistoryRetrain) HasGateMetric() bool`

HasGateMetric returns a boolean if a field has been set.

### GetGatePass

`func (o *AiHistoryRetrain) GetGatePass() bool`

GetGatePass returns the GatePass field if non-nil, zero value otherwise.

### GetGatePassOk

`func (o *AiHistoryRetrain) GetGatePassOk() (*bool, bool)`

GetGatePassOk returns a tuple with the GatePass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatePass

`func (o *AiHistoryRetrain) SetGatePass(v bool)`

SetGatePass sets GatePass field to given value.

### HasGatePass

`func (o *AiHistoryRetrain) HasGatePass() bool`

HasGatePass returns a boolean if a field has been set.

### GetGateValue

`func (o *AiHistoryRetrain) GetGateValue() float32`

GetGateValue returns the GateValue field if non-nil, zero value otherwise.

### GetGateValueOk

`func (o *AiHistoryRetrain) GetGateValueOk() (*float32, bool)`

GetGateValueOk returns a tuple with the GateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateValue

`func (o *AiHistoryRetrain) SetGateValue(v float32)`

SetGateValue sets GateValue field to given value.

### HasGateValue

`func (o *AiHistoryRetrain) HasGateValue() bool`

HasGateValue returns a boolean if a field has been set.

### GetHoldoutAccuracy

`func (o *AiHistoryRetrain) GetHoldoutAccuracy() float32`

GetHoldoutAccuracy returns the HoldoutAccuracy field if non-nil, zero value otherwise.

### GetHoldoutAccuracyOk

`func (o *AiHistoryRetrain) GetHoldoutAccuracyOk() (*float32, bool)`

GetHoldoutAccuracyOk returns a tuple with the HoldoutAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldoutAccuracy

`func (o *AiHistoryRetrain) SetHoldoutAccuracy(v float32)`

SetHoldoutAccuracy sets HoldoutAccuracy field to given value.

### HasHoldoutAccuracy

`func (o *AiHistoryRetrain) HasHoldoutAccuracy() bool`

HasHoldoutAccuracy returns a boolean if a field has been set.

### GetPublished

`func (o *AiHistoryRetrain) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *AiHistoryRetrain) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *AiHistoryRetrain) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *AiHistoryRetrain) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetTrainedTime

`func (o *AiHistoryRetrain) GetTrainedTime() string`

GetTrainedTime returns the TrainedTime field if non-nil, zero value otherwise.

### GetTrainedTimeOk

`func (o *AiHistoryRetrain) GetTrainedTimeOk() (*string, bool)`

GetTrainedTimeOk returns a tuple with the TrainedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainedTime

`func (o *AiHistoryRetrain) SetTrainedTime(v string)`

SetTrainedTime sets TrainedTime field to given value.

### HasTrainedTime

`func (o *AiHistoryRetrain) HasTrainedTime() bool`

HasTrainedTime returns a boolean if a field has been set.

### GetVersion

`func (o *AiHistoryRetrain) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AiHistoryRetrain) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AiHistoryRetrain) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AiHistoryRetrain) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AiRetrainMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to **int32** |  | [optional] 
**GateBase** | Pointer to **float32** |  | [optional] 
**GateKind** | Pointer to **string** |  | [optional] 
**GateMetric** | Pointer to **string** |  | [optional] 
**GatePassed** | Pointer to **bool** |  | [optional] 
**GateValue** | Pointer to **float32** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**TrainedTime** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewAiRetrainMeta

`func NewAiRetrainMeta() *AiRetrainMeta`

NewAiRetrainMeta instantiates a new AiRetrainMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiRetrainMetaWithDefaults

`func NewAiRetrainMetaWithDefaults() *AiRetrainMeta`

NewAiRetrainMetaWithDefaults instantiates a new AiRetrainMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *AiRetrainMeta) GetEvents() int32`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AiRetrainMeta) GetEventsOk() (*int32, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AiRetrainMeta) SetEvents(v int32)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AiRetrainMeta) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetGateBase

`func (o *AiRetrainMeta) GetGateBase() float32`

GetGateBase returns the GateBase field if non-nil, zero value otherwise.

### GetGateBaseOk

`func (o *AiRetrainMeta) GetGateBaseOk() (*float32, bool)`

GetGateBaseOk returns a tuple with the GateBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateBase

`func (o *AiRetrainMeta) SetGateBase(v float32)`

SetGateBase sets GateBase field to given value.

### HasGateBase

`func (o *AiRetrainMeta) HasGateBase() bool`

HasGateBase returns a boolean if a field has been set.

### GetGateKind

`func (o *AiRetrainMeta) GetGateKind() string`

GetGateKind returns the GateKind field if non-nil, zero value otherwise.

### GetGateKindOk

`func (o *AiRetrainMeta) GetGateKindOk() (*string, bool)`

GetGateKindOk returns a tuple with the GateKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateKind

`func (o *AiRetrainMeta) SetGateKind(v string)`

SetGateKind sets GateKind field to given value.

### HasGateKind

`func (o *AiRetrainMeta) HasGateKind() bool`

HasGateKind returns a boolean if a field has been set.

### GetGateMetric

`func (o *AiRetrainMeta) GetGateMetric() string`

GetGateMetric returns the GateMetric field if non-nil, zero value otherwise.

### GetGateMetricOk

`func (o *AiRetrainMeta) GetGateMetricOk() (*string, bool)`

GetGateMetricOk returns a tuple with the GateMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateMetric

`func (o *AiRetrainMeta) SetGateMetric(v string)`

SetGateMetric sets GateMetric field to given value.

### HasGateMetric

`func (o *AiRetrainMeta) HasGateMetric() bool`

HasGateMetric returns a boolean if a field has been set.

### GetGatePassed

`func (o *AiRetrainMeta) GetGatePassed() bool`

GetGatePassed returns the GatePassed field if non-nil, zero value otherwise.

### GetGatePassedOk

`func (o *AiRetrainMeta) GetGatePassedOk() (*bool, bool)`

GetGatePassedOk returns a tuple with the GatePassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatePassed

`func (o *AiRetrainMeta) SetGatePassed(v bool)`

SetGatePassed sets GatePassed field to given value.

### HasGatePassed

`func (o *AiRetrainMeta) HasGatePassed() bool`

HasGatePassed returns a boolean if a field has been set.

### GetGateValue

`func (o *AiRetrainMeta) GetGateValue() float32`

GetGateValue returns the GateValue field if non-nil, zero value otherwise.

### GetGateValueOk

`func (o *AiRetrainMeta) GetGateValueOk() (*float32, bool)`

GetGateValueOk returns a tuple with the GateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateValue

`func (o *AiRetrainMeta) SetGateValue(v float32)`

SetGateValue sets GateValue field to given value.

### HasGateValue

`func (o *AiRetrainMeta) HasGateValue() bool`

HasGateValue returns a boolean if a field has been set.

### GetNote

`func (o *AiRetrainMeta) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AiRetrainMeta) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AiRetrainMeta) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *AiRetrainMeta) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPublished

`func (o *AiRetrainMeta) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *AiRetrainMeta) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *AiRetrainMeta) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *AiRetrainMeta) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetTrainedTime

`func (o *AiRetrainMeta) GetTrainedTime() string`

GetTrainedTime returns the TrainedTime field if non-nil, zero value otherwise.

### GetTrainedTimeOk

`func (o *AiRetrainMeta) GetTrainedTimeOk() (*string, bool)`

GetTrainedTimeOk returns a tuple with the TrainedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainedTime

`func (o *AiRetrainMeta) SetTrainedTime(v string)`

SetTrainedTime sets TrainedTime field to given value.

### HasTrainedTime

`func (o *AiRetrainMeta) HasTrainedTime() bool`

HasTrainedTime returns a boolean if a field has been set.

### GetVersion

`func (o *AiRetrainMeta) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AiRetrainMeta) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AiRetrainMeta) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AiRetrainMeta) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



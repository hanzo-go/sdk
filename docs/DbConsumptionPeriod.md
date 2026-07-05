# DbConsumptionPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodId** | Pointer to **string** |  | [optional] 
**PeriodStart** | Pointer to **time.Time** |  | [optional] 
**PeriodEnd** | Pointer to **time.Time** |  | [optional] 
**ActiveTimeSeconds** | Pointer to **int64** |  | [optional] 
**ComputeTimeSeconds** | Pointer to **int64** |  | [optional] 
**DataStorageBytesHour** | Pointer to **int64** |  | [optional] 
**WrittenDataBytes** | Pointer to **int64** |  | [optional] 
**DataTransferBytes** | Pointer to **int64** |  | [optional] 

## Methods

### NewDbConsumptionPeriod

`func NewDbConsumptionPeriod() *DbConsumptionPeriod`

NewDbConsumptionPeriod instantiates a new DbConsumptionPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbConsumptionPeriodWithDefaults

`func NewDbConsumptionPeriodWithDefaults() *DbConsumptionPeriod`

NewDbConsumptionPeriodWithDefaults instantiates a new DbConsumptionPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodId

`func (o *DbConsumptionPeriod) GetPeriodId() string`

GetPeriodId returns the PeriodId field if non-nil, zero value otherwise.

### GetPeriodIdOk

`func (o *DbConsumptionPeriod) GetPeriodIdOk() (*string, bool)`

GetPeriodIdOk returns a tuple with the PeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodId

`func (o *DbConsumptionPeriod) SetPeriodId(v string)`

SetPeriodId sets PeriodId field to given value.

### HasPeriodId

`func (o *DbConsumptionPeriod) HasPeriodId() bool`

HasPeriodId returns a boolean if a field has been set.

### GetPeriodStart

`func (o *DbConsumptionPeriod) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *DbConsumptionPeriod) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *DbConsumptionPeriod) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *DbConsumptionPeriod) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *DbConsumptionPeriod) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *DbConsumptionPeriod) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *DbConsumptionPeriod) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *DbConsumptionPeriod) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetActiveTimeSeconds

`func (o *DbConsumptionPeriod) GetActiveTimeSeconds() int64`

GetActiveTimeSeconds returns the ActiveTimeSeconds field if non-nil, zero value otherwise.

### GetActiveTimeSecondsOk

`func (o *DbConsumptionPeriod) GetActiveTimeSecondsOk() (*int64, bool)`

GetActiveTimeSecondsOk returns a tuple with the ActiveTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTimeSeconds

`func (o *DbConsumptionPeriod) SetActiveTimeSeconds(v int64)`

SetActiveTimeSeconds sets ActiveTimeSeconds field to given value.

### HasActiveTimeSeconds

`func (o *DbConsumptionPeriod) HasActiveTimeSeconds() bool`

HasActiveTimeSeconds returns a boolean if a field has been set.

### GetComputeTimeSeconds

`func (o *DbConsumptionPeriod) GetComputeTimeSeconds() int64`

GetComputeTimeSeconds returns the ComputeTimeSeconds field if non-nil, zero value otherwise.

### GetComputeTimeSecondsOk

`func (o *DbConsumptionPeriod) GetComputeTimeSecondsOk() (*int64, bool)`

GetComputeTimeSecondsOk returns a tuple with the ComputeTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeTimeSeconds

`func (o *DbConsumptionPeriod) SetComputeTimeSeconds(v int64)`

SetComputeTimeSeconds sets ComputeTimeSeconds field to given value.

### HasComputeTimeSeconds

`func (o *DbConsumptionPeriod) HasComputeTimeSeconds() bool`

HasComputeTimeSeconds returns a boolean if a field has been set.

### GetDataStorageBytesHour

`func (o *DbConsumptionPeriod) GetDataStorageBytesHour() int64`

GetDataStorageBytesHour returns the DataStorageBytesHour field if non-nil, zero value otherwise.

### GetDataStorageBytesHourOk

`func (o *DbConsumptionPeriod) GetDataStorageBytesHourOk() (*int64, bool)`

GetDataStorageBytesHourOk returns a tuple with the DataStorageBytesHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStorageBytesHour

`func (o *DbConsumptionPeriod) SetDataStorageBytesHour(v int64)`

SetDataStorageBytesHour sets DataStorageBytesHour field to given value.

### HasDataStorageBytesHour

`func (o *DbConsumptionPeriod) HasDataStorageBytesHour() bool`

HasDataStorageBytesHour returns a boolean if a field has been set.

### GetWrittenDataBytes

`func (o *DbConsumptionPeriod) GetWrittenDataBytes() int64`

GetWrittenDataBytes returns the WrittenDataBytes field if non-nil, zero value otherwise.

### GetWrittenDataBytesOk

`func (o *DbConsumptionPeriod) GetWrittenDataBytesOk() (*int64, bool)`

GetWrittenDataBytesOk returns a tuple with the WrittenDataBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrittenDataBytes

`func (o *DbConsumptionPeriod) SetWrittenDataBytes(v int64)`

SetWrittenDataBytes sets WrittenDataBytes field to given value.

### HasWrittenDataBytes

`func (o *DbConsumptionPeriod) HasWrittenDataBytes() bool`

HasWrittenDataBytes returns a boolean if a field has been set.

### GetDataTransferBytes

`func (o *DbConsumptionPeriod) GetDataTransferBytes() int64`

GetDataTransferBytes returns the DataTransferBytes field if non-nil, zero value otherwise.

### GetDataTransferBytesOk

`func (o *DbConsumptionPeriod) GetDataTransferBytesOk() (*int64, bool)`

GetDataTransferBytesOk returns a tuple with the DataTransferBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferBytes

`func (o *DbConsumptionPeriod) SetDataTransferBytes(v int64)`

SetDataTransferBytes sets DataTransferBytes field to given value.

### HasDataTransferBytes

`func (o *DbConsumptionPeriod) HasDataTransferBytes() bool`

HasDataTransferBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



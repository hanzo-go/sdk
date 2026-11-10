# TrainRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | Pointer to **[]string** | Devices are the accelerator kinds it runs on there; none is the executor&#39;s own. | [optional] 
**PriceNanoUsd** | Pointer to **int64** | Price is what one device-second costs there, in nano-USD: 0 on the org&#39;s machines. | [optional] 
**Where** | Pointer to **string** | Where is hanzo, Hanzo&#39;s devices, charged by the device-second to the payer inside the job&#39;s budget; or machines, the org&#39;s linked machines, charged nothing. | [optional] 

## Methods

### NewTrainRun

`func NewTrainRun() *TrainRun`

NewTrainRun instantiates a new TrainRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainRunWithDefaults

`func NewTrainRunWithDefaults() *TrainRun`

NewTrainRunWithDefaults instantiates a new TrainRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *TrainRun) GetDevices() []string`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *TrainRun) GetDevicesOk() (*[]string, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *TrainRun) SetDevices(v []string)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *TrainRun) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetPriceNanoUsd

`func (o *TrainRun) GetPriceNanoUsd() int64`

GetPriceNanoUsd returns the PriceNanoUsd field if non-nil, zero value otherwise.

### GetPriceNanoUsdOk

`func (o *TrainRun) GetPriceNanoUsdOk() (*int64, bool)`

GetPriceNanoUsdOk returns a tuple with the PriceNanoUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNanoUsd

`func (o *TrainRun) SetPriceNanoUsd(v int64)`

SetPriceNanoUsd sets PriceNanoUsd field to given value.

### HasPriceNanoUsd

`func (o *TrainRun) HasPriceNanoUsd() bool`

HasPriceNanoUsd returns a boolean if a field has been set.

### GetWhere

`func (o *TrainRun) GetWhere() string`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *TrainRun) GetWhereOk() (*string, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *TrainRun) SetWhere(v string)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *TrainRun) HasWhere() bool`

HasWhere returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



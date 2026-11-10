# TrainClaimed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coordinator** | Pointer to **string** | Coordinator is where a join task reaches its lead. | [optional] 
**Device** | Pointer to **string** | Device is the accelerator kind it runs on, and Devices how many of them. | [optional] 
**Devices** | Pointer to **int64** |  | [optional] 
**Held** | Pointer to **int64** | Held is the device-seconds held for it. The cloud meters the task by the clock and stops its job before it runs past them. | [optional] 
**Id** | Pointer to **string** | ID names the task. | [optional] 
**Job** | Pointer to **string** | Job names its job. | [optional] 
**Lease** | Pointer to **string** | Lease authenticates every report for the task; it is answered once. | [optional] 
**Report** | Pointer to **int64** | Report is the most seconds that may pass between its reports. | [optional] 
**Role** | Pointer to **string** | Role is lead or join. | [optional] 

## Methods

### NewTrainClaimed

`func NewTrainClaimed() *TrainClaimed`

NewTrainClaimed instantiates a new TrainClaimed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainClaimedWithDefaults

`func NewTrainClaimedWithDefaults() *TrainClaimed`

NewTrainClaimedWithDefaults instantiates a new TrainClaimed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoordinator

`func (o *TrainClaimed) GetCoordinator() string`

GetCoordinator returns the Coordinator field if non-nil, zero value otherwise.

### GetCoordinatorOk

`func (o *TrainClaimed) GetCoordinatorOk() (*string, bool)`

GetCoordinatorOk returns a tuple with the Coordinator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinator

`func (o *TrainClaimed) SetCoordinator(v string)`

SetCoordinator sets Coordinator field to given value.

### HasCoordinator

`func (o *TrainClaimed) HasCoordinator() bool`

HasCoordinator returns a boolean if a field has been set.

### GetDevice

`func (o *TrainClaimed) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *TrainClaimed) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *TrainClaimed) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *TrainClaimed) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDevices

`func (o *TrainClaimed) GetDevices() int64`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *TrainClaimed) GetDevicesOk() (*int64, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *TrainClaimed) SetDevices(v int64)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *TrainClaimed) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetHeld

`func (o *TrainClaimed) GetHeld() int64`

GetHeld returns the Held field if non-nil, zero value otherwise.

### GetHeldOk

`func (o *TrainClaimed) GetHeldOk() (*int64, bool)`

GetHeldOk returns a tuple with the Held field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeld

`func (o *TrainClaimed) SetHeld(v int64)`

SetHeld sets Held field to given value.

### HasHeld

`func (o *TrainClaimed) HasHeld() bool`

HasHeld returns a boolean if a field has been set.

### GetId

`func (o *TrainClaimed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrainClaimed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrainClaimed) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrainClaimed) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJob

`func (o *TrainClaimed) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *TrainClaimed) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *TrainClaimed) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *TrainClaimed) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetLease

`func (o *TrainClaimed) GetLease() string`

GetLease returns the Lease field if non-nil, zero value otherwise.

### GetLeaseOk

`func (o *TrainClaimed) GetLeaseOk() (*string, bool)`

GetLeaseOk returns a tuple with the Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease

`func (o *TrainClaimed) SetLease(v string)`

SetLease sets Lease field to given value.

### HasLease

`func (o *TrainClaimed) HasLease() bool`

HasLease returns a boolean if a field has been set.

### GetReport

`func (o *TrainClaimed) GetReport() int64`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *TrainClaimed) GetReportOk() (*int64, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *TrainClaimed) SetReport(v int64)`

SetReport sets Report field to given value.

### HasReport

`func (o *TrainClaimed) HasReport() bool`

HasReport returns a boolean if a field has been set.

### GetRole

`func (o *TrainClaimed) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TrainClaimed) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TrainClaimed) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *TrainClaimed) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



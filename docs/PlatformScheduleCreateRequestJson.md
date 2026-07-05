# PlatformScheduleCreateRequestJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** |  | 
**Schedule** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 

## Methods

### NewPlatformScheduleCreateRequestJson

`func NewPlatformScheduleCreateRequestJson(command string, schedule string, ) *PlatformScheduleCreateRequestJson`

NewPlatformScheduleCreateRequestJson instantiates a new PlatformScheduleCreateRequestJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformScheduleCreateRequestJsonWithDefaults

`func NewPlatformScheduleCreateRequestJsonWithDefaults() *PlatformScheduleCreateRequestJson`

NewPlatformScheduleCreateRequestJsonWithDefaults instantiates a new PlatformScheduleCreateRequestJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *PlatformScheduleCreateRequestJson) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *PlatformScheduleCreateRequestJson) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *PlatformScheduleCreateRequestJson) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetSchedule

`func (o *PlatformScheduleCreateRequestJson) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *PlatformScheduleCreateRequestJson) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *PlatformScheduleCreateRequestJson) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetEnabled

`func (o *PlatformScheduleCreateRequestJson) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PlatformScheduleCreateRequestJson) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PlatformScheduleCreateRequestJson) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PlatformScheduleCreateRequestJson) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServerId

`func (o *PlatformScheduleCreateRequestJson) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *PlatformScheduleCreateRequestJson) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *PlatformScheduleCreateRequestJson) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *PlatformScheduleCreateRequestJson) HasServerId() bool`

HasServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



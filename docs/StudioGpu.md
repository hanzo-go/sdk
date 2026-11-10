# StudioGpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Machine** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name is the device, as its worker reports it. | [optional] 
**Owner** | Pointer to **string** | Owner is org (linked by your org with hanzo link) or platform. | [optional] 
**Seen** | Pointer to **string** | Seen is its worker&#39;s last heartbeat. | [optional] 
**Status** | Pointer to **string** | Status is idle, busy or offline. | [optional] 
**VramMb** | Pointer to **int64** |  | [optional] 

## Methods

### NewStudioGpu

`func NewStudioGpu() *StudioGpu`

NewStudioGpu instantiates a new StudioGpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioGpuWithDefaults

`func NewStudioGpuWithDefaults() *StudioGpu`

NewStudioGpuWithDefaults instantiates a new StudioGpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StudioGpu) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StudioGpu) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StudioGpu) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StudioGpu) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMachine

`func (o *StudioGpu) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *StudioGpu) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *StudioGpu) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *StudioGpu) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetName

`func (o *StudioGpu) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StudioGpu) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StudioGpu) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StudioGpu) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *StudioGpu) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *StudioGpu) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *StudioGpu) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *StudioGpu) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSeen

`func (o *StudioGpu) GetSeen() string`

GetSeen returns the Seen field if non-nil, zero value otherwise.

### GetSeenOk

`func (o *StudioGpu) GetSeenOk() (*string, bool)`

GetSeenOk returns a tuple with the Seen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeen

`func (o *StudioGpu) SetSeen(v string)`

SetSeen sets Seen field to given value.

### HasSeen

`func (o *StudioGpu) HasSeen() bool`

HasSeen returns a boolean if a field has been set.

### GetStatus

`func (o *StudioGpu) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StudioGpu) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StudioGpu) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StudioGpu) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVramMb

`func (o *StudioGpu) GetVramMb() int64`

GetVramMb returns the VramMb field if non-nil, zero value otherwise.

### GetVramMbOk

`func (o *StudioGpu) GetVramMbOk() (*int64, bool)`

GetVramMbOk returns a tuple with the VramMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVramMb

`func (o *StudioGpu) SetVramMb(v int64)`

SetVramMb sets VramMb field to given value.

### HasVramMb

`func (o *StudioGpu) HasVramMb() bool`

HasVramMb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



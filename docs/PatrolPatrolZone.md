# PatrolPatrolZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **string** | Device is the detector type. | [optional] 
**Location** | Pointer to **string** | Location is where in the building the zone is. | [optional] 
**Ref** | Pointer to **string** | Ref is the zone number the panel reports. | [optional] 
**State** | Pointer to **string** | State is the zone&#39;s last known condition. | [optional] 

## Methods

### NewPatrolPatrolZone

`func NewPatrolPatrolZone() *PatrolPatrolZone`

NewPatrolPatrolZone instantiates a new PatrolPatrolZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolZoneWithDefaults

`func NewPatrolPatrolZoneWithDefaults() *PatrolPatrolZone`

NewPatrolPatrolZoneWithDefaults instantiates a new PatrolPatrolZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *PatrolPatrolZone) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatrolPatrolZone) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatrolPatrolZone) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatrolPatrolZone) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetLocation

`func (o *PatrolPatrolZone) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatrolPatrolZone) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatrolPatrolZone) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatrolPatrolZone) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRef

`func (o *PatrolPatrolZone) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *PatrolPatrolZone) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *PatrolPatrolZone) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *PatrolPatrolZone) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetState

`func (o *PatrolPatrolZone) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PatrolPatrolZone) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PatrolPatrolZone) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PatrolPatrolZone) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



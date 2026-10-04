# PatrolPatrolUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Due** | Pointer to **string** | Due is the lone-worker deadline, or empty when none is set. | [optional] 
**Lat** | Pointer to **float64** | Lat is the unit&#39;s last latitude in degrees. | [optional] 
**Lon** | Pointer to **float64** | Lon is the unit&#39;s last longitude in degrees. | [optional] 
**Name** | Pointer to **string** | Name is the unit&#39;s call sign, which is its document name and the segment /v1/patrol/unit/{name}/fix addresses it by. | [optional] 
**Officer** | Pointer to **string** | Officer is who is crewing it. | [optional] 
**Sector** | Pointer to **string** | Sector is the patrol sector it works. | [optional] 
**Seen** | Pointer to **string** | Seen is when the unit last reported. | [optional] 
**State** | Pointer to **string** | State is available, assigned, enroute, onsite, offline, break or emergency. | [optional] 
**Trail** | Pointer to [**[]PatrolPatrolFix**](PatrolPatrolFix.md) | Trail is the tail of the unit&#39;s recorded positions, oldest first. | [optional] 

## Methods

### NewPatrolPatrolUnit

`func NewPatrolPatrolUnit() *PatrolPatrolUnit`

NewPatrolPatrolUnit instantiates a new PatrolPatrolUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolUnitWithDefaults

`func NewPatrolPatrolUnitWithDefaults() *PatrolPatrolUnit`

NewPatrolPatrolUnitWithDefaults instantiates a new PatrolPatrolUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDue

`func (o *PatrolPatrolUnit) GetDue() string`

GetDue returns the Due field if non-nil, zero value otherwise.

### GetDueOk

`func (o *PatrolPatrolUnit) GetDueOk() (*string, bool)`

GetDueOk returns a tuple with the Due field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDue

`func (o *PatrolPatrolUnit) SetDue(v string)`

SetDue sets Due field to given value.

### HasDue

`func (o *PatrolPatrolUnit) HasDue() bool`

HasDue returns a boolean if a field has been set.

### GetLat

`func (o *PatrolPatrolUnit) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *PatrolPatrolUnit) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *PatrolPatrolUnit) SetLat(v float64)`

SetLat sets Lat field to given value.

### HasLat

`func (o *PatrolPatrolUnit) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLon

`func (o *PatrolPatrolUnit) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *PatrolPatrolUnit) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *PatrolPatrolUnit) SetLon(v float64)`

SetLon sets Lon field to given value.

### HasLon

`func (o *PatrolPatrolUnit) HasLon() bool`

HasLon returns a boolean if a field has been set.

### GetName

`func (o *PatrolPatrolUnit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatrolPatrolUnit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatrolPatrolUnit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatrolPatrolUnit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOfficer

`func (o *PatrolPatrolUnit) GetOfficer() string`

GetOfficer returns the Officer field if non-nil, zero value otherwise.

### GetOfficerOk

`func (o *PatrolPatrolUnit) GetOfficerOk() (*string, bool)`

GetOfficerOk returns a tuple with the Officer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficer

`func (o *PatrolPatrolUnit) SetOfficer(v string)`

SetOfficer sets Officer field to given value.

### HasOfficer

`func (o *PatrolPatrolUnit) HasOfficer() bool`

HasOfficer returns a boolean if a field has been set.

### GetSector

`func (o *PatrolPatrolUnit) GetSector() string`

GetSector returns the Sector field if non-nil, zero value otherwise.

### GetSectorOk

`func (o *PatrolPatrolUnit) GetSectorOk() (*string, bool)`

GetSectorOk returns a tuple with the Sector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSector

`func (o *PatrolPatrolUnit) SetSector(v string)`

SetSector sets Sector field to given value.

### HasSector

`func (o *PatrolPatrolUnit) HasSector() bool`

HasSector returns a boolean if a field has been set.

### GetSeen

`func (o *PatrolPatrolUnit) GetSeen() string`

GetSeen returns the Seen field if non-nil, zero value otherwise.

### GetSeenOk

`func (o *PatrolPatrolUnit) GetSeenOk() (*string, bool)`

GetSeenOk returns a tuple with the Seen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeen

`func (o *PatrolPatrolUnit) SetSeen(v string)`

SetSeen sets Seen field to given value.

### HasSeen

`func (o *PatrolPatrolUnit) HasSeen() bool`

HasSeen returns a boolean if a field has been set.

### GetState

`func (o *PatrolPatrolUnit) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PatrolPatrolUnit) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PatrolPatrolUnit) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PatrolPatrolUnit) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTrail

`func (o *PatrolPatrolUnit) GetTrail() []PatrolPatrolFix`

GetTrail returns the Trail field if non-nil, zero value otherwise.

### GetTrailOk

`func (o *PatrolPatrolUnit) GetTrailOk() (*[]PatrolPatrolFix, bool)`

GetTrailOk returns a tuple with the Trail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrail

`func (o *PatrolPatrolUnit) SetTrail(v []PatrolPatrolFix)`

SetTrail sets Trail field to given value.

### HasTrail

`func (o *PatrolPatrolUnit) HasTrail() bool`

HasTrail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



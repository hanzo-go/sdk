# PatrolPatrolFixIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is not a caller&#39;s to fill. A position report is stamped by the server, and a body naming it is refused rather than silently restamped: the trail is the proof of where a unit was, so a report that thought it was dating itself must learn otherwise. | [optional] 
**Lat** | Pointer to **float64** | Lat is the latitude in degrees. | [optional] 
**Lon** | Pointer to **float64** | Lon is the longitude in degrees. | [optional] 
**Name** | Pointer to **string** | Name is the unit&#39;s call sign, from the path. | [optional] 

## Methods

### NewPatrolPatrolFixIn

`func NewPatrolPatrolFixIn() *PatrolPatrolFixIn`

NewPatrolPatrolFixIn instantiates a new PatrolPatrolFixIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolFixInWithDefaults

`func NewPatrolPatrolFixInWithDefaults() *PatrolPatrolFixIn`

NewPatrolPatrolFixInWithDefaults instantiates a new PatrolPatrolFixIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *PatrolPatrolFixIn) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *PatrolPatrolFixIn) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *PatrolPatrolFixIn) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *PatrolPatrolFixIn) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetLat

`func (o *PatrolPatrolFixIn) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *PatrolPatrolFixIn) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *PatrolPatrolFixIn) SetLat(v float64)`

SetLat sets Lat field to given value.

### HasLat

`func (o *PatrolPatrolFixIn) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLon

`func (o *PatrolPatrolFixIn) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *PatrolPatrolFixIn) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *PatrolPatrolFixIn) SetLon(v float64)`

SetLon sets Lon field to given value.

### HasLon

`func (o *PatrolPatrolFixIn) HasLon() bool`

HasLon returns a boolean if a field has been set.

### GetName

`func (o *PatrolPatrolFixIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatrolPatrolFixIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatrolPatrolFixIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatrolPatrolFixIn) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



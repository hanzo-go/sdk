# PatrolPatrolConfirmIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lat** | Pointer to **float64** | Lat is the latitude the confirmation was taken at. | [optional] 
**Lon** | Pointer to **float64** | Lon is the longitude the confirmation was taken at. | [optional] 
**Method** | Pointer to **string** | Method is how presence was proved: tag, scan, fence, manual or photo. | [optional] 
**Name** | Pointer to **string** | Name is the checkpoint&#39;s document name, from the path. | [optional] 
**Photo** | Pointer to **string** | Photo is the object key of a photo taken there, never the bytes. | [optional] 

## Methods

### NewPatrolPatrolConfirmIn

`func NewPatrolPatrolConfirmIn() *PatrolPatrolConfirmIn`

NewPatrolPatrolConfirmIn instantiates a new PatrolPatrolConfirmIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolConfirmInWithDefaults

`func NewPatrolPatrolConfirmInWithDefaults() *PatrolPatrolConfirmIn`

NewPatrolPatrolConfirmInWithDefaults instantiates a new PatrolPatrolConfirmIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLat

`func (o *PatrolPatrolConfirmIn) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *PatrolPatrolConfirmIn) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *PatrolPatrolConfirmIn) SetLat(v float64)`

SetLat sets Lat field to given value.

### HasLat

`func (o *PatrolPatrolConfirmIn) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLon

`func (o *PatrolPatrolConfirmIn) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *PatrolPatrolConfirmIn) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *PatrolPatrolConfirmIn) SetLon(v float64)`

SetLon sets Lon field to given value.

### HasLon

`func (o *PatrolPatrolConfirmIn) HasLon() bool`

HasLon returns a boolean if a field has been set.

### GetMethod

`func (o *PatrolPatrolConfirmIn) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PatrolPatrolConfirmIn) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PatrolPatrolConfirmIn) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PatrolPatrolConfirmIn) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetName

`func (o *PatrolPatrolConfirmIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatrolPatrolConfirmIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatrolPatrolConfirmIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatrolPatrolConfirmIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhoto

`func (o *PatrolPatrolConfirmIn) GetPhoto() string`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *PatrolPatrolConfirmIn) GetPhotoOk() (*string, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *PatrolPatrolConfirmIn) SetPhoto(v string)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *PatrolPatrolConfirmIn) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



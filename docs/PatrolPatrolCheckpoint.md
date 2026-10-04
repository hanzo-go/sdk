# PatrolPatrolCheckpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** | Actor is who confirmed it. | [optional] 
**At** | Pointer to **string** | At is when it was confirmed, or empty. | [optional] 
**Close** | Pointer to **string** | Close is when that window shuts. | [optional] 
**Lat** | Pointer to **float64** | Lat is the latitude the confirmation was taken at. | [optional] 
**Lon** | Pointer to **float64** | Lon is the longitude the confirmation was taken at. | [optional] 
**Method** | Pointer to **string** | Method is how it was confirmed: tag, scan, fence, manual or photo. | [optional] 
**Name** | Pointer to **string** | Name is the checkpoint&#39;s document name and the segment /v1/patrol/point/{name}/confirm addresses it by. | [optional] 
**Open** | Pointer to **string** | Open is when the window to confirm it opens. | [optional] 
**Photo** | Pointer to **string** | Photo is the object key of the photo taken, or empty. | [optional] 
**Site** | Pointer to **string** | Site is the site to attend. | [optional] 
**State** | Pointer to **string** | State is pending, done, late or skipped. | [optional] 
**Tag** | Pointer to **string** | Tag is the tag or marker to read there. | [optional] 
**Tour** | Pointer to **string** | Tour is the round this checkpoint belongs to. | [optional] 

## Methods

### NewPatrolPatrolCheckpoint

`func NewPatrolPatrolCheckpoint() *PatrolPatrolCheckpoint`

NewPatrolPatrolCheckpoint instantiates a new PatrolPatrolCheckpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolCheckpointWithDefaults

`func NewPatrolPatrolCheckpointWithDefaults() *PatrolPatrolCheckpoint`

NewPatrolPatrolCheckpointWithDefaults instantiates a new PatrolPatrolCheckpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *PatrolPatrolCheckpoint) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *PatrolPatrolCheckpoint) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *PatrolPatrolCheckpoint) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *PatrolPatrolCheckpoint) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAt

`func (o *PatrolPatrolCheckpoint) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *PatrolPatrolCheckpoint) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *PatrolPatrolCheckpoint) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *PatrolPatrolCheckpoint) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetClose

`func (o *PatrolPatrolCheckpoint) GetClose() string`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *PatrolPatrolCheckpoint) GetCloseOk() (*string, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *PatrolPatrolCheckpoint) SetClose(v string)`

SetClose sets Close field to given value.

### HasClose

`func (o *PatrolPatrolCheckpoint) HasClose() bool`

HasClose returns a boolean if a field has been set.

### GetLat

`func (o *PatrolPatrolCheckpoint) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *PatrolPatrolCheckpoint) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *PatrolPatrolCheckpoint) SetLat(v float64)`

SetLat sets Lat field to given value.

### HasLat

`func (o *PatrolPatrolCheckpoint) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLon

`func (o *PatrolPatrolCheckpoint) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *PatrolPatrolCheckpoint) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *PatrolPatrolCheckpoint) SetLon(v float64)`

SetLon sets Lon field to given value.

### HasLon

`func (o *PatrolPatrolCheckpoint) HasLon() bool`

HasLon returns a boolean if a field has been set.

### GetMethod

`func (o *PatrolPatrolCheckpoint) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PatrolPatrolCheckpoint) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PatrolPatrolCheckpoint) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PatrolPatrolCheckpoint) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetName

`func (o *PatrolPatrolCheckpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatrolPatrolCheckpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatrolPatrolCheckpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatrolPatrolCheckpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpen

`func (o *PatrolPatrolCheckpoint) GetOpen() string`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *PatrolPatrolCheckpoint) GetOpenOk() (*string, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *PatrolPatrolCheckpoint) SetOpen(v string)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *PatrolPatrolCheckpoint) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetPhoto

`func (o *PatrolPatrolCheckpoint) GetPhoto() string`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *PatrolPatrolCheckpoint) GetPhotoOk() (*string, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *PatrolPatrolCheckpoint) SetPhoto(v string)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *PatrolPatrolCheckpoint) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### GetSite

`func (o *PatrolPatrolCheckpoint) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatrolPatrolCheckpoint) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatrolPatrolCheckpoint) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatrolPatrolCheckpoint) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetState

`func (o *PatrolPatrolCheckpoint) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PatrolPatrolCheckpoint) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PatrolPatrolCheckpoint) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PatrolPatrolCheckpoint) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTag

`func (o *PatrolPatrolCheckpoint) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PatrolPatrolCheckpoint) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PatrolPatrolCheckpoint) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PatrolPatrolCheckpoint) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTour

`func (o *PatrolPatrolCheckpoint) GetTour() string`

GetTour returns the Tour field if non-nil, zero value otherwise.

### GetTourOk

`func (o *PatrolPatrolCheckpoint) GetTourOk() (*string, bool)`

GetTourOk returns a tuple with the Tour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTour

`func (o *PatrolPatrolCheckpoint) SetTour(v string)`

SetTour sets Tour field to given value.

### HasTour

`func (o *PatrolPatrolCheckpoint) HasTour() bool`

HasTour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



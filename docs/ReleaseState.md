# ReleaseState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndedAt** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **string** | Error is why it stopped, when it failed. | [optional] 
**Id** | Pointer to **string** | ID is the build id returned by the 202. | [optional] 
**Image** | Pointer to **string** | Image is the tag the release publishes on success. | [optional] 
**Reached** | Pointer to **string** | Reached is the last pipeline step completed: built, smoked, tagged, pinned. | [optional] 
**Sha** | Pointer to **string** | SHA is the commit the release pinned. | [optional] 
**StartedAt** | Pointer to **int32** | StartedAt / EndedAt are unix seconds. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;releasing\&quot;, \&quot;released\&quot; or \&quot;failed\&quot;. | [optional] 
**Version** | Pointer to **string** | Version is that tag without the leading \&quot;v\&quot;. | [optional] 

## Methods

### NewReleaseState

`func NewReleaseState() *ReleaseState`

NewReleaseState instantiates a new ReleaseState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseStateWithDefaults

`func NewReleaseStateWithDefaults() *ReleaseState`

NewReleaseStateWithDefaults instantiates a new ReleaseState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndedAt

`func (o *ReleaseState) GetEndedAt() int32`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *ReleaseState) GetEndedAtOk() (*int32, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *ReleaseState) SetEndedAt(v int32)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *ReleaseState) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetError

`func (o *ReleaseState) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ReleaseState) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ReleaseState) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ReleaseState) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *ReleaseState) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReleaseState) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReleaseState) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReleaseState) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *ReleaseState) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ReleaseState) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ReleaseState) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ReleaseState) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetReached

`func (o *ReleaseState) GetReached() string`

GetReached returns the Reached field if non-nil, zero value otherwise.

### GetReachedOk

`func (o *ReleaseState) GetReachedOk() (*string, bool)`

GetReachedOk returns a tuple with the Reached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReached

`func (o *ReleaseState) SetReached(v string)`

SetReached sets Reached field to given value.

### HasReached

`func (o *ReleaseState) HasReached() bool`

HasReached returns a boolean if a field has been set.

### GetSha

`func (o *ReleaseState) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *ReleaseState) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *ReleaseState) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *ReleaseState) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetStartedAt

`func (o *ReleaseState) GetStartedAt() int32`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ReleaseState) GetStartedAtOk() (*int32, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ReleaseState) SetStartedAt(v int32)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ReleaseState) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ReleaseState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReleaseState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReleaseState) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReleaseState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *ReleaseState) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ReleaseState) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ReleaseState) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ReleaseState) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



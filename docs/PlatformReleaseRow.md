# PlatformReleaseRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **string** | Environment is the deploy target the application names. | [optional] 
**Id** | Pointer to **string** | ID is the deployment&#39;s id — a release IS a deployment that reached the cluster. | [optional] 
**Name** | Pointer to **string** | Name is the application the release belongs to. | [optional] 
**ReleasedAt** | Pointer to **string** | ReleasedAt is when the deployment last changed, RFC3339 UTC. | [optional] 
**Status** | Pointer to **string** | Status is deploying or live — the two states that mean released. | [optional] 
**Version** | Pointer to **string** | Version is the released image tag, or v&lt;n&gt; when the image carries none. | [optional] 

## Methods

### NewPlatformReleaseRow

`func NewPlatformReleaseRow() *PlatformReleaseRow`

NewPlatformReleaseRow instantiates a new PlatformReleaseRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformReleaseRowWithDefaults

`func NewPlatformReleaseRowWithDefaults() *PlatformReleaseRow`

NewPlatformReleaseRowWithDefaults instantiates a new PlatformReleaseRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *PlatformReleaseRow) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PlatformReleaseRow) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PlatformReleaseRow) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PlatformReleaseRow) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *PlatformReleaseRow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlatformReleaseRow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlatformReleaseRow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlatformReleaseRow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PlatformReleaseRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformReleaseRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformReleaseRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformReleaseRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReleasedAt

`func (o *PlatformReleaseRow) GetReleasedAt() string`

GetReleasedAt returns the ReleasedAt field if non-nil, zero value otherwise.

### GetReleasedAtOk

`func (o *PlatformReleaseRow) GetReleasedAtOk() (*string, bool)`

GetReleasedAtOk returns a tuple with the ReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedAt

`func (o *PlatformReleaseRow) SetReleasedAt(v string)`

SetReleasedAt sets ReleasedAt field to given value.

### HasReleasedAt

`func (o *PlatformReleaseRow) HasReleasedAt() bool`

HasReleasedAt returns a boolean if a field has been set.

### GetStatus

`func (o *PlatformReleaseRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlatformReleaseRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlatformReleaseRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlatformReleaseRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *PlatformReleaseRow) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PlatformReleaseRow) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PlatformReleaseRow) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PlatformReleaseRow) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



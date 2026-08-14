# BuildRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | Pointer to **string** | Commit is the short git ref the build pinned. | [optional] 
**Duration** | Pointer to **string** | Duration is the wall time of a TERMINAL build; empty while it still runs. | [optional] 
**Id** | Pointer to **string** | ID is the build record&#39;s id. | [optional] 
**Repo** | Pointer to **string** | Repo is the repo the build built, or the image it produced. | [optional] 
**StartedAt** | Pointer to **string** | StartedAt is when the build was recorded, RFC3339 UTC. | [optional] 
**Status** | Pointer to **string** | Status is the build&#39;s real state: queued, building, succeeded or failed. | [optional] 

## Methods

### NewBuildRow

`func NewBuildRow() *BuildRow`

NewBuildRow instantiates a new BuildRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildRowWithDefaults

`func NewBuildRowWithDefaults() *BuildRow`

NewBuildRowWithDefaults instantiates a new BuildRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *BuildRow) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *BuildRow) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *BuildRow) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *BuildRow) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetDuration

`func (o *BuildRow) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *BuildRow) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *BuildRow) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *BuildRow) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetId

`func (o *BuildRow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BuildRow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BuildRow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BuildRow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRepo

`func (o *BuildRow) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *BuildRow) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *BuildRow) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *BuildRow) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetStartedAt

`func (o *BuildRow) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BuildRow) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BuildRow) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *BuildRow) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *BuildRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BuildRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BuildRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BuildRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



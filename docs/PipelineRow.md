# PipelineRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **string** | Duration is how long that run took; empty while it is still queued or building. | [optional] 
**Id** | Pointer to **string** | ID is the application id — one pipeline is one application. | [optional] 
**LastRun** | Pointer to **string** | LastRun is when the most recent deployment started, RFC3339 UTC. | [optional] 
**Name** | Pointer to **string** | Name is the application&#39;s name. | [optional] 
**Repo** | Pointer to **string** | Repo is the git repo or image the pipeline builds from. | [optional] 
**Status** | Pointer to **string** | Status is the latest deployment&#39;s status, or the app&#39;s when it has none. | [optional] 

## Methods

### NewPipelineRow

`func NewPipelineRow() *PipelineRow`

NewPipelineRow instantiates a new PipelineRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineRowWithDefaults

`func NewPipelineRowWithDefaults() *PipelineRow`

NewPipelineRowWithDefaults instantiates a new PipelineRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *PipelineRow) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PipelineRow) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PipelineRow) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PipelineRow) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetId

`func (o *PipelineRow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PipelineRow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PipelineRow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PipelineRow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastRun

`func (o *PipelineRow) GetLastRun() string`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *PipelineRow) GetLastRunOk() (*string, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *PipelineRow) SetLastRun(v string)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *PipelineRow) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetName

`func (o *PipelineRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelineRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelineRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PipelineRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepo

`func (o *PipelineRow) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *PipelineRow) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *PipelineRow) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *PipelineRow) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetStatus

`func (o *PipelineRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PipelineRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PipelineRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PipelineRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AutoFlowRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**FlowId** | Pointer to **string** |  | [optional] 
**FlowVersionId** | Pointer to **string** |  | [optional] 
**FlowDisplayName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**FinishTime** | Pointer to **time.Time** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAutoFlowRun

`func NewAutoFlowRun() *AutoFlowRun`

NewAutoFlowRun instantiates a new AutoFlowRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoFlowRunWithDefaults

`func NewAutoFlowRunWithDefaults() *AutoFlowRun`

NewAutoFlowRunWithDefaults instantiates a new AutoFlowRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoFlowRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoFlowRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoFlowRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoFlowRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *AutoFlowRun) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AutoFlowRun) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AutoFlowRun) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AutoFlowRun) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetFlowId

`func (o *AutoFlowRun) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *AutoFlowRun) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *AutoFlowRun) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *AutoFlowRun) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetFlowVersionId

`func (o *AutoFlowRun) GetFlowVersionId() string`

GetFlowVersionId returns the FlowVersionId field if non-nil, zero value otherwise.

### GetFlowVersionIdOk

`func (o *AutoFlowRun) GetFlowVersionIdOk() (*string, bool)`

GetFlowVersionIdOk returns a tuple with the FlowVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowVersionId

`func (o *AutoFlowRun) SetFlowVersionId(v string)`

SetFlowVersionId sets FlowVersionId field to given value.

### HasFlowVersionId

`func (o *AutoFlowRun) HasFlowVersionId() bool`

HasFlowVersionId returns a boolean if a field has been set.

### GetFlowDisplayName

`func (o *AutoFlowRun) GetFlowDisplayName() string`

GetFlowDisplayName returns the FlowDisplayName field if non-nil, zero value otherwise.

### GetFlowDisplayNameOk

`func (o *AutoFlowRun) GetFlowDisplayNameOk() (*string, bool)`

GetFlowDisplayNameOk returns a tuple with the FlowDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDisplayName

`func (o *AutoFlowRun) SetFlowDisplayName(v string)`

SetFlowDisplayName sets FlowDisplayName field to given value.

### HasFlowDisplayName

`func (o *AutoFlowRun) HasFlowDisplayName() bool`

HasFlowDisplayName returns a boolean if a field has been set.

### GetStatus

`func (o *AutoFlowRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutoFlowRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutoFlowRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutoFlowRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartTime

`func (o *AutoFlowRun) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AutoFlowRun) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AutoFlowRun) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *AutoFlowRun) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetFinishTime

`func (o *AutoFlowRun) GetFinishTime() time.Time`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *AutoFlowRun) GetFinishTimeOk() (*time.Time, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *AutoFlowRun) SetFinishTime(v time.Time)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *AutoFlowRun) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetDuration

`func (o *AutoFlowRun) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AutoFlowRun) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AutoFlowRun) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AutoFlowRun) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTags

`func (o *AutoFlowRun) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AutoFlowRun) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AutoFlowRun) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AutoFlowRun) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreated

`func (o *AutoFlowRun) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AutoFlowRun) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AutoFlowRun) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AutoFlowRun) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *AutoFlowRun) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AutoFlowRun) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AutoFlowRun) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AutoFlowRun) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



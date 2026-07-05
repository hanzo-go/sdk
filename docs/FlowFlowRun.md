# FlowFlowRun

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

### NewFlowFlowRun

`func NewFlowFlowRun() *FlowFlowRun`

NewFlowFlowRun instantiates a new FlowFlowRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowFlowRunWithDefaults

`func NewFlowFlowRunWithDefaults() *FlowFlowRun`

NewFlowFlowRunWithDefaults instantiates a new FlowFlowRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlowFlowRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowFlowRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowFlowRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlowFlowRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *FlowFlowRun) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *FlowFlowRun) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *FlowFlowRun) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *FlowFlowRun) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetFlowId

`func (o *FlowFlowRun) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *FlowFlowRun) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *FlowFlowRun) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *FlowFlowRun) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetFlowVersionId

`func (o *FlowFlowRun) GetFlowVersionId() string`

GetFlowVersionId returns the FlowVersionId field if non-nil, zero value otherwise.

### GetFlowVersionIdOk

`func (o *FlowFlowRun) GetFlowVersionIdOk() (*string, bool)`

GetFlowVersionIdOk returns a tuple with the FlowVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowVersionId

`func (o *FlowFlowRun) SetFlowVersionId(v string)`

SetFlowVersionId sets FlowVersionId field to given value.

### HasFlowVersionId

`func (o *FlowFlowRun) HasFlowVersionId() bool`

HasFlowVersionId returns a boolean if a field has been set.

### GetFlowDisplayName

`func (o *FlowFlowRun) GetFlowDisplayName() string`

GetFlowDisplayName returns the FlowDisplayName field if non-nil, zero value otherwise.

### GetFlowDisplayNameOk

`func (o *FlowFlowRun) GetFlowDisplayNameOk() (*string, bool)`

GetFlowDisplayNameOk returns a tuple with the FlowDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDisplayName

`func (o *FlowFlowRun) SetFlowDisplayName(v string)`

SetFlowDisplayName sets FlowDisplayName field to given value.

### HasFlowDisplayName

`func (o *FlowFlowRun) HasFlowDisplayName() bool`

HasFlowDisplayName returns a boolean if a field has been set.

### GetStatus

`func (o *FlowFlowRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlowFlowRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlowFlowRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlowFlowRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartTime

`func (o *FlowFlowRun) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *FlowFlowRun) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *FlowFlowRun) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *FlowFlowRun) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetFinishTime

`func (o *FlowFlowRun) GetFinishTime() time.Time`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *FlowFlowRun) GetFinishTimeOk() (*time.Time, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *FlowFlowRun) SetFinishTime(v time.Time)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *FlowFlowRun) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetDuration

`func (o *FlowFlowRun) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *FlowFlowRun) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *FlowFlowRun) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *FlowFlowRun) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTags

`func (o *FlowFlowRun) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FlowFlowRun) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FlowFlowRun) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FlowFlowRun) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreated

`func (o *FlowFlowRun) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FlowFlowRun) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FlowFlowRun) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FlowFlowRun) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *FlowFlowRun) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FlowFlowRun) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FlowFlowRun) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FlowFlowRun) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CiJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conclusion** | Pointer to **string** | Conclusion is how the job ended; empty until it has. | [optional] 
**EndedAt** | Pointer to **time.Time** | EndedAt is when it finished; absent until it has. | [optional] 
**Id** | Pointer to **int64** | ID is the job&#39;s id on the forge. | [optional] 
**Name** | Pointer to **string** | Name is the job&#39;s name as the workflow declares it. | [optional] 
**StartedAt** | Pointer to **time.Time** | StartedAt is when the job started; absent until it has. | [optional] 
**Status** | Pointer to **string** | Status is whether the job is over, in the same words as the run&#39;s. | [optional] 
**Steps** | Pointer to [**[]CiStep**](CiStep.md) | Steps are the job&#39;s steps, in order. | [optional] 
**Url** | Pointer to **string** | URL is the job&#39;s page on the forge. | [optional] 

## Methods

### NewCiJob

`func NewCiJob() *CiJob`

NewCiJob instantiates a new CiJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiJobWithDefaults

`func NewCiJobWithDefaults() *CiJob`

NewCiJobWithDefaults instantiates a new CiJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConclusion

`func (o *CiJob) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *CiJob) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *CiJob) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.

### HasConclusion

`func (o *CiJob) HasConclusion() bool`

HasConclusion returns a boolean if a field has been set.

### GetEndedAt

`func (o *CiJob) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *CiJob) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *CiJob) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *CiJob) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetId

`func (o *CiJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CiJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CiJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CiJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CiJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CiJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CiJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CiJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartedAt

`func (o *CiJob) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CiJob) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CiJob) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CiJob) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *CiJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CiJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CiJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CiJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSteps

`func (o *CiJob) GetSteps() []CiStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *CiJob) GetStepsOk() (*[]CiStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *CiJob) SetSteps(v []CiStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *CiJob) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetUrl

`func (o *CiJob) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CiJob) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CiJob) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CiJob) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



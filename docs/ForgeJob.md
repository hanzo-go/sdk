# ForgeJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to [**ForgeJobRepository**](ForgeJobRepository.md) |  | [optional] 
**WorkflowJob** | Pointer to [**ForgeJobWorkflowJob**](ForgeJobWorkflowJob.md) |  | [optional] 

## Methods

### NewForgeJob

`func NewForgeJob() *ForgeJob`

NewForgeJob instantiates a new ForgeJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForgeJobWithDefaults

`func NewForgeJobWithDefaults() *ForgeJob`

NewForgeJobWithDefaults instantiates a new ForgeJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ForgeJob) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ForgeJob) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ForgeJob) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ForgeJob) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRepository

`func (o *ForgeJob) GetRepository() ForgeJobRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ForgeJob) GetRepositoryOk() (*ForgeJobRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ForgeJob) SetRepository(v ForgeJobRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ForgeJob) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetWorkflowJob

`func (o *ForgeJob) GetWorkflowJob() ForgeJobWorkflowJob`

GetWorkflowJob returns the WorkflowJob field if non-nil, zero value otherwise.

### GetWorkflowJobOk

`func (o *ForgeJob) GetWorkflowJobOk() (*ForgeJobWorkflowJob, bool)`

GetWorkflowJobOk returns a tuple with the WorkflowJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowJob

`func (o *ForgeJob) SetWorkflowJob(v ForgeJobWorkflowJob)`

SetWorkflowJob sets WorkflowJob field to given value.

### HasWorkflowJob

`func (o *ForgeJob) HasWorkflowJob() bool`

HasWorkflowJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



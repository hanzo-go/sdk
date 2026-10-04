# GitWorkflowRuns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GitWorkflowRun**](GitWorkflowRun.md) | Data is the runs in scope, newest first. | [optional] 

## Methods

### NewGitWorkflowRuns

`func NewGitWorkflowRuns() *GitWorkflowRuns`

NewGitWorkflowRuns instantiates a new GitWorkflowRuns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitWorkflowRunsWithDefaults

`func NewGitWorkflowRunsWithDefaults() *GitWorkflowRuns`

NewGitWorkflowRunsWithDefaults instantiates a new GitWorkflowRuns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GitWorkflowRuns) GetData() []GitWorkflowRun`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GitWorkflowRuns) GetDataOk() (*[]GitWorkflowRun, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GitWorkflowRuns) SetData(v []GitWorkflowRun)`

SetData sets Data field to given value.

### HasData

`func (o *GitWorkflowRuns) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



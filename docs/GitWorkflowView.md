# GitWorkflowView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fault** | Pointer to **string** | Fault is why this document cannot run at all, when it cannot be read. | [optional] 
**Jobs** | Pointer to [**[]GitJobView**](GitJobView.md) | Jobs is what it declares, and where each would run. | [optional] 
**Name** | Pointer to **string** | Name is the document&#39;s path in the repository. | [optional] 

## Methods

### NewGitWorkflowView

`func NewGitWorkflowView() *GitWorkflowView`

NewGitWorkflowView instantiates a new GitWorkflowView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitWorkflowViewWithDefaults

`func NewGitWorkflowViewWithDefaults() *GitWorkflowView`

NewGitWorkflowViewWithDefaults instantiates a new GitWorkflowView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFault

`func (o *GitWorkflowView) GetFault() string`

GetFault returns the Fault field if non-nil, zero value otherwise.

### GetFaultOk

`func (o *GitWorkflowView) GetFaultOk() (*string, bool)`

GetFaultOk returns a tuple with the Fault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFault

`func (o *GitWorkflowView) SetFault(v string)`

SetFault sets Fault field to given value.

### HasFault

`func (o *GitWorkflowView) HasFault() bool`

HasFault returns a boolean if a field has been set.

### GetJobs

`func (o *GitWorkflowView) GetJobs() []GitJobView`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *GitWorkflowView) GetJobsOk() (*[]GitJobView, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *GitWorkflowView) SetJobs(v []GitJobView)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *GitWorkflowView) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetName

`func (o *GitWorkflowView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitWorkflowView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitWorkflowView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitWorkflowView) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



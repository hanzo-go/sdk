# GitJobView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the job&#39;s id in the workflow document. | [optional] 
**Pool** | Pointer to **string** | Pool is the declared pool carrying every one of those labels, or empty when no pool does — which is the answer to \&quot;why is nothing running\&quot;. | [optional] 
**RunsOn** | Pointer to **[]string** | RunsOn is the labels the job asked for. | [optional] 

## Methods

### NewGitJobView

`func NewGitJobView() *GitJobView`

NewGitJobView instantiates a new GitJobView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitJobViewWithDefaults

`func NewGitJobViewWithDefaults() *GitJobView`

NewGitJobViewWithDefaults instantiates a new GitJobView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GitJobView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitJobView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitJobView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitJobView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPool

`func (o *GitJobView) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *GitJobView) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *GitJobView) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *GitJobView) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetRunsOn

`func (o *GitJobView) GetRunsOn() []string`

GetRunsOn returns the RunsOn field if non-nil, zero value otherwise.

### GetRunsOnOk

`func (o *GitJobView) GetRunsOnOk() (*[]string, bool)`

GetRunsOnOk returns a tuple with the RunsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunsOn

`func (o *GitJobView) SetRunsOn(v []string)`

SetRunsOn sets RunsOn field to given value.

### HasRunsOn

`func (o *GitJobView) HasRunsOn() bool`

HasRunsOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CloudGithubBackfillResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **int32** | Created is how many native issues this pass created. | [optional] 
**Failed** | Pointer to **int32** | Failed is how many repos or issues errored; the pass continues past each. | [optional] 
**Issues** | Pointer to **int32** | Issues is how many upstream issues were seen. | [optional] 
**Repos** | Pointer to **int32** | Repos is how many granted repos were walked (archived/disabled are skipped). | [optional] 
**Truncated** | Pointer to **bool** | Truncated is set when the time budget or the issue cap stopped the pass early. Re-run to continue — the mirror is idempotent by ExtRef, so nothing duplicates. | [optional] 
**Updated** | Pointer to **int32** | Updated is how many existing native issues this pass refreshed. | [optional] 

## Methods

### NewCloudGithubBackfillResult

`func NewCloudGithubBackfillResult() *CloudGithubBackfillResult`

NewCloudGithubBackfillResult instantiates a new CloudGithubBackfillResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGithubBackfillResultWithDefaults

`func NewCloudGithubBackfillResultWithDefaults() *CloudGithubBackfillResult`

NewCloudGithubBackfillResultWithDefaults instantiates a new CloudGithubBackfillResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CloudGithubBackfillResult) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CloudGithubBackfillResult) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CloudGithubBackfillResult) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CloudGithubBackfillResult) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFailed

`func (o *CloudGithubBackfillResult) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *CloudGithubBackfillResult) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *CloudGithubBackfillResult) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *CloudGithubBackfillResult) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetIssues

`func (o *CloudGithubBackfillResult) GetIssues() int32`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *CloudGithubBackfillResult) GetIssuesOk() (*int32, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *CloudGithubBackfillResult) SetIssues(v int32)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *CloudGithubBackfillResult) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetRepos

`func (o *CloudGithubBackfillResult) GetRepos() int32`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *CloudGithubBackfillResult) GetReposOk() (*int32, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *CloudGithubBackfillResult) SetRepos(v int32)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *CloudGithubBackfillResult) HasRepos() bool`

HasRepos returns a boolean if a field has been set.

### GetTruncated

`func (o *CloudGithubBackfillResult) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *CloudGithubBackfillResult) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *CloudGithubBackfillResult) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *CloudGithubBackfillResult) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetUpdated

`func (o *CloudGithubBackfillResult) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CloudGithubBackfillResult) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CloudGithubBackfillResult) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CloudGithubBackfillResult) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



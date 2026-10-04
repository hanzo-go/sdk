# ProviderGithubBackfillResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **int64** | Created is how many native issues this pass created. | [optional] 
**Failed** | Pointer to **int64** | Failed is how many repos or issues errored; the pass continues past each. | [optional] 
**Issues** | Pointer to **int64** | Issues is how many upstream issues were seen. | [optional] 
**Repos** | Pointer to **int64** | Repos is how many granted repos were walked (archived/disabled are skipped). | [optional] 
**Truncated** | Pointer to **bool** | Truncated is set when the time budget or the issue cap stopped the pass early. Re-run to continue — the mirror is idempotent by ExtRef, so nothing duplicates. | [optional] 
**Updated** | Pointer to **int64** | Updated is how many existing native issues this pass refreshed. | [optional] 

## Methods

### NewProviderGithubBackfillResult

`func NewProviderGithubBackfillResult() *ProviderGithubBackfillResult`

NewProviderGithubBackfillResult instantiates a new ProviderGithubBackfillResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGithubBackfillResultWithDefaults

`func NewProviderGithubBackfillResultWithDefaults() *ProviderGithubBackfillResult`

NewProviderGithubBackfillResultWithDefaults instantiates a new ProviderGithubBackfillResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ProviderGithubBackfillResult) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ProviderGithubBackfillResult) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ProviderGithubBackfillResult) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ProviderGithubBackfillResult) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFailed

`func (o *ProviderGithubBackfillResult) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ProviderGithubBackfillResult) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ProviderGithubBackfillResult) SetFailed(v int64)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ProviderGithubBackfillResult) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetIssues

`func (o *ProviderGithubBackfillResult) GetIssues() int64`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *ProviderGithubBackfillResult) GetIssuesOk() (*int64, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *ProviderGithubBackfillResult) SetIssues(v int64)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *ProviderGithubBackfillResult) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetRepos

`func (o *ProviderGithubBackfillResult) GetRepos() int64`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *ProviderGithubBackfillResult) GetReposOk() (*int64, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *ProviderGithubBackfillResult) SetRepos(v int64)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *ProviderGithubBackfillResult) HasRepos() bool`

HasRepos returns a boolean if a field has been set.

### GetTruncated

`func (o *ProviderGithubBackfillResult) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *ProviderGithubBackfillResult) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *ProviderGithubBackfillResult) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *ProviderGithubBackfillResult) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetUpdated

`func (o *ProviderGithubBackfillResult) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ProviderGithubBackfillResult) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ProviderGithubBackfillResult) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ProviderGithubBackfillResult) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



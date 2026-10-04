# CiExecutions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FetchedAt** | Pointer to **time.Time** |  | [optional] 
**Orgs** | Pointer to **[]string** |  | [optional] 
**Repos** | Pointer to **int64** |  | [optional] 
**Runs** | Pointer to [**[]CiExecution**](CiExecution.md) |  | [optional] 
**SourceErr** | Pointer to **string** |  | [optional] 
**Stale** | Pointer to **bool** |  | [optional] 

## Methods

### NewCiExecutions

`func NewCiExecutions() *CiExecutions`

NewCiExecutions instantiates a new CiExecutions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiExecutionsWithDefaults

`func NewCiExecutionsWithDefaults() *CiExecutions`

NewCiExecutionsWithDefaults instantiates a new CiExecutions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFetchedAt

`func (o *CiExecutions) GetFetchedAt() time.Time`

GetFetchedAt returns the FetchedAt field if non-nil, zero value otherwise.

### GetFetchedAtOk

`func (o *CiExecutions) GetFetchedAtOk() (*time.Time, bool)`

GetFetchedAtOk returns a tuple with the FetchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchedAt

`func (o *CiExecutions) SetFetchedAt(v time.Time)`

SetFetchedAt sets FetchedAt field to given value.

### HasFetchedAt

`func (o *CiExecutions) HasFetchedAt() bool`

HasFetchedAt returns a boolean if a field has been set.

### GetOrgs

`func (o *CiExecutions) GetOrgs() []string`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *CiExecutions) GetOrgsOk() (*[]string, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *CiExecutions) SetOrgs(v []string)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *CiExecutions) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetRepos

`func (o *CiExecutions) GetRepos() int64`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *CiExecutions) GetReposOk() (*int64, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *CiExecutions) SetRepos(v int64)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *CiExecutions) HasRepos() bool`

HasRepos returns a boolean if a field has been set.

### GetRuns

`func (o *CiExecutions) GetRuns() []CiExecution`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *CiExecutions) GetRunsOk() (*[]CiExecution, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *CiExecutions) SetRuns(v []CiExecution)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *CiExecutions) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetSourceErr

`func (o *CiExecutions) GetSourceErr() string`

GetSourceErr returns the SourceErr field if non-nil, zero value otherwise.

### GetSourceErrOk

`func (o *CiExecutions) GetSourceErrOk() (*string, bool)`

GetSourceErrOk returns a tuple with the SourceErr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceErr

`func (o *CiExecutions) SetSourceErr(v string)`

SetSourceErr sets SourceErr field to given value.

### HasSourceErr

`func (o *CiExecutions) HasSourceErr() bool`

HasSourceErr returns a boolean if a field has been set.

### GetStale

`func (o *CiExecutions) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *CiExecutions) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *CiExecutions) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *CiExecutions) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



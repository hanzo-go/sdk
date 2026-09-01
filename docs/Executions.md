# Executions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FetchedAt** | Pointer to **time.Time** |  | [optional] 
**Orgs** | Pointer to **[]string** |  | [optional] 
**Repos** | Pointer to **int32** |  | [optional] 
**Runs** | Pointer to [**[]Execution**](Execution.md) |  | [optional] 
**SourceErr** | Pointer to **string** |  | [optional] 
**Stale** | Pointer to **bool** |  | [optional] 

## Methods

### NewExecutions

`func NewExecutions() *Executions`

NewExecutions instantiates a new Executions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionsWithDefaults

`func NewExecutionsWithDefaults() *Executions`

NewExecutionsWithDefaults instantiates a new Executions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFetchedAt

`func (o *Executions) GetFetchedAt() time.Time`

GetFetchedAt returns the FetchedAt field if non-nil, zero value otherwise.

### GetFetchedAtOk

`func (o *Executions) GetFetchedAtOk() (*time.Time, bool)`

GetFetchedAtOk returns a tuple with the FetchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchedAt

`func (o *Executions) SetFetchedAt(v time.Time)`

SetFetchedAt sets FetchedAt field to given value.

### HasFetchedAt

`func (o *Executions) HasFetchedAt() bool`

HasFetchedAt returns a boolean if a field has been set.

### GetOrgs

`func (o *Executions) GetOrgs() []string`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *Executions) GetOrgsOk() (*[]string, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *Executions) SetOrgs(v []string)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *Executions) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetRepos

`func (o *Executions) GetRepos() int32`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *Executions) GetReposOk() (*int32, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *Executions) SetRepos(v int32)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *Executions) HasRepos() bool`

HasRepos returns a boolean if a field has been set.

### GetRuns

`func (o *Executions) GetRuns() []Execution`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *Executions) GetRunsOk() (*[]Execution, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *Executions) SetRuns(v []Execution)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *Executions) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetSourceErr

`func (o *Executions) GetSourceErr() string`

GetSourceErr returns the SourceErr field if non-nil, zero value otherwise.

### GetSourceErrOk

`func (o *Executions) GetSourceErrOk() (*string, bool)`

GetSourceErrOk returns a tuple with the SourceErr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceErr

`func (o *Executions) SetSourceErr(v string)`

SetSourceErr sets SourceErr field to given value.

### HasSourceErr

`func (o *Executions) HasSourceErr() bool`

HasSourceErr returns a boolean if a field has been set.

### GetStale

`func (o *Executions) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *Executions) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *Executions) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *Executions) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



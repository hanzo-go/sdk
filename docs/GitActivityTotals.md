# GitActivityTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int64** | Active is how many have a commit in the window. | [optional] 
**Active30** | Pointer to **int64** | Active30 is how many have a commit in the last 30 days. | [optional] 
**Commits** | Pointer to **int64** | Commits is distinct commits in the window, each counted once in this scope: every commit the scope&#39;s repositories hold, the upstream projects carried and kept current included. | [optional] 
**Contributors** | Pointer to **int64** | Contributors is distinct author addresses in the window, counted and never named. | [optional] 
**Empty** | Pointer to **int64** | Empty is how many hold no commit. | [optional] 
**Forks** | Pointer to **int64** | Forks is commits found only in forks, upstream history included. | [optional] 
**Last** | Pointer to **string** | Last is the newest commit time in the scope, RFC 3339 UTC; empty when none. | [optional] 
**MergeMonths** | Pointer to **[]int64** | MergeMonths is Merges per calendar month, as Months. | [optional] 
**Merges** | Pointer to **int64** | Merges is how many of those have more than one parent. | [optional] 
**Mirrors** | Pointer to **int64** | Mirrors is how many are pull mirrors of github.com. | [optional] 
**Months** | Pointer to **[]int64** | Months is Commits per calendar month, oldest first, one per entry of the document&#39;s Months. | [optional] 
**Orgs** | Pointer to **int64** | Orgs is how many orgs the estate holds on the forge. | [optional] 
**Ours** | Pointer to **int64** | Ours is how many of those the estate&#39;s people authored: an address a member of an estate org commits with, or one at an estate domain. | [optional] 
**OursMonths** | Pointer to **[]int64** | OursMonths is Ours per calendar month, as Months. | [optional] 
**Own** | Pointer to **int64** | Own is commits found in at least one repository GitHub does not call a fork. | [optional] 
**People** | Pointer to **int64** | People is how many of those are the estate&#39;s own. | [optional] 
**Private** | Pointer to **int64** | Private is how many only their members can read. They are counted, never named. | [optional] 
**Public** | Pointer to **int64** | Public is how many of them anyone can read. | [optional] 
**Repos** | Pointer to **int64** | Repos is how many repositories the scope holds on the forge. | [optional] 

## Methods

### NewGitActivityTotals

`func NewGitActivityTotals() *GitActivityTotals`

NewGitActivityTotals instantiates a new GitActivityTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitActivityTotalsWithDefaults

`func NewGitActivityTotalsWithDefaults() *GitActivityTotals`

NewGitActivityTotalsWithDefaults instantiates a new GitActivityTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *GitActivityTotals) GetActive() int64`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GitActivityTotals) GetActiveOk() (*int64, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GitActivityTotals) SetActive(v int64)`

SetActive sets Active field to given value.

### HasActive

`func (o *GitActivityTotals) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetActive30

`func (o *GitActivityTotals) GetActive30() int64`

GetActive30 returns the Active30 field if non-nil, zero value otherwise.

### GetActive30Ok

`func (o *GitActivityTotals) GetActive30Ok() (*int64, bool)`

GetActive30Ok returns a tuple with the Active30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive30

`func (o *GitActivityTotals) SetActive30(v int64)`

SetActive30 sets Active30 field to given value.

### HasActive30

`func (o *GitActivityTotals) HasActive30() bool`

HasActive30 returns a boolean if a field has been set.

### GetCommits

`func (o *GitActivityTotals) GetCommits() int64`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *GitActivityTotals) GetCommitsOk() (*int64, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *GitActivityTotals) SetCommits(v int64)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *GitActivityTotals) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### GetContributors

`func (o *GitActivityTotals) GetContributors() int64`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *GitActivityTotals) GetContributorsOk() (*int64, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *GitActivityTotals) SetContributors(v int64)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *GitActivityTotals) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetEmpty

`func (o *GitActivityTotals) GetEmpty() int64`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *GitActivityTotals) GetEmptyOk() (*int64, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *GitActivityTotals) SetEmpty(v int64)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *GitActivityTotals) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### GetForks

`func (o *GitActivityTotals) GetForks() int64`

GetForks returns the Forks field if non-nil, zero value otherwise.

### GetForksOk

`func (o *GitActivityTotals) GetForksOk() (*int64, bool)`

GetForksOk returns a tuple with the Forks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForks

`func (o *GitActivityTotals) SetForks(v int64)`

SetForks sets Forks field to given value.

### HasForks

`func (o *GitActivityTotals) HasForks() bool`

HasForks returns a boolean if a field has been set.

### GetLast

`func (o *GitActivityTotals) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GitActivityTotals) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GitActivityTotals) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *GitActivityTotals) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetMergeMonths

`func (o *GitActivityTotals) GetMergeMonths() []int64`

GetMergeMonths returns the MergeMonths field if non-nil, zero value otherwise.

### GetMergeMonthsOk

`func (o *GitActivityTotals) GetMergeMonthsOk() (*[]int64, bool)`

GetMergeMonthsOk returns a tuple with the MergeMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeMonths

`func (o *GitActivityTotals) SetMergeMonths(v []int64)`

SetMergeMonths sets MergeMonths field to given value.

### HasMergeMonths

`func (o *GitActivityTotals) HasMergeMonths() bool`

HasMergeMonths returns a boolean if a field has been set.

### GetMerges

`func (o *GitActivityTotals) GetMerges() int64`

GetMerges returns the Merges field if non-nil, zero value otherwise.

### GetMergesOk

`func (o *GitActivityTotals) GetMergesOk() (*int64, bool)`

GetMergesOk returns a tuple with the Merges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerges

`func (o *GitActivityTotals) SetMerges(v int64)`

SetMerges sets Merges field to given value.

### HasMerges

`func (o *GitActivityTotals) HasMerges() bool`

HasMerges returns a boolean if a field has been set.

### GetMirrors

`func (o *GitActivityTotals) GetMirrors() int64`

GetMirrors returns the Mirrors field if non-nil, zero value otherwise.

### GetMirrorsOk

`func (o *GitActivityTotals) GetMirrorsOk() (*int64, bool)`

GetMirrorsOk returns a tuple with the Mirrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrors

`func (o *GitActivityTotals) SetMirrors(v int64)`

SetMirrors sets Mirrors field to given value.

### HasMirrors

`func (o *GitActivityTotals) HasMirrors() bool`

HasMirrors returns a boolean if a field has been set.

### GetMonths

`func (o *GitActivityTotals) GetMonths() []int64`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *GitActivityTotals) GetMonthsOk() (*[]int64, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *GitActivityTotals) SetMonths(v []int64)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *GitActivityTotals) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### GetOrgs

`func (o *GitActivityTotals) GetOrgs() int64`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *GitActivityTotals) GetOrgsOk() (*int64, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *GitActivityTotals) SetOrgs(v int64)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *GitActivityTotals) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetOurs

`func (o *GitActivityTotals) GetOurs() int64`

GetOurs returns the Ours field if non-nil, zero value otherwise.

### GetOursOk

`func (o *GitActivityTotals) GetOursOk() (*int64, bool)`

GetOursOk returns a tuple with the Ours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOurs

`func (o *GitActivityTotals) SetOurs(v int64)`

SetOurs sets Ours field to given value.

### HasOurs

`func (o *GitActivityTotals) HasOurs() bool`

HasOurs returns a boolean if a field has been set.

### GetOursMonths

`func (o *GitActivityTotals) GetOursMonths() []int64`

GetOursMonths returns the OursMonths field if non-nil, zero value otherwise.

### GetOursMonthsOk

`func (o *GitActivityTotals) GetOursMonthsOk() (*[]int64, bool)`

GetOursMonthsOk returns a tuple with the OursMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOursMonths

`func (o *GitActivityTotals) SetOursMonths(v []int64)`

SetOursMonths sets OursMonths field to given value.

### HasOursMonths

`func (o *GitActivityTotals) HasOursMonths() bool`

HasOursMonths returns a boolean if a field has been set.

### GetOwn

`func (o *GitActivityTotals) GetOwn() int64`

GetOwn returns the Own field if non-nil, zero value otherwise.

### GetOwnOk

`func (o *GitActivityTotals) GetOwnOk() (*int64, bool)`

GetOwnOk returns a tuple with the Own field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwn

`func (o *GitActivityTotals) SetOwn(v int64)`

SetOwn sets Own field to given value.

### HasOwn

`func (o *GitActivityTotals) HasOwn() bool`

HasOwn returns a boolean if a field has been set.

### GetPeople

`func (o *GitActivityTotals) GetPeople() int64`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *GitActivityTotals) GetPeopleOk() (*int64, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *GitActivityTotals) SetPeople(v int64)`

SetPeople sets People field to given value.

### HasPeople

`func (o *GitActivityTotals) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetPrivate

`func (o *GitActivityTotals) GetPrivate() int64`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *GitActivityTotals) GetPrivateOk() (*int64, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *GitActivityTotals) SetPrivate(v int64)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *GitActivityTotals) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPublic

`func (o *GitActivityTotals) GetPublic() int64`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *GitActivityTotals) GetPublicOk() (*int64, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *GitActivityTotals) SetPublic(v int64)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *GitActivityTotals) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRepos

`func (o *GitActivityTotals) GetRepos() int64`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *GitActivityTotals) GetReposOk() (*int64, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *GitActivityTotals) SetRepos(v int64)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *GitActivityTotals) HasRepos() bool`

HasRepos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



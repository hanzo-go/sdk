# GitActivityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **map[string]int64** | Active is how many of the group&#39;s repositories have a commit on each UTC day. | [optional] 
**Active30** | Pointer to **int64** | Active30 is how many have a commit in the last 30 days. | [optional] 
**Commits** | Pointer to **int64** | Commits is distinct commits in the window, each counted once in this scope: every commit the scope&#39;s repositories hold, the upstream projects carried and kept current included. | [optional] 
**Contributors** | Pointer to **int64** | Contributors is distinct author addresses in the window, counted and never named. | [optional] 
**Daily** | Pointer to **map[string]int64** | Daily is the group&#39;s distinct commits per UTC day, YYYY-MM-DD. | [optional] 
**Empty** | Pointer to **int64** | Empty is how many hold no commit. | [optional] 
**Last** | Pointer to **string** | Last is the newest commit time in the scope, RFC 3339 UTC; empty when none. | [optional] 
**Latest** | Pointer to [**[]GitActivityCommit**](GitActivityCommit.md) | Latest is the group&#39;s newest commits of public repositories, newest first. | [optional] 
**MergeMonths** | Pointer to **[]int64** | MergeMonths is Merges per calendar month, as Months. | [optional] 
**Merges** | Pointer to **int64** | Merges is how many of those have more than one parent. | [optional] 
**Mirrors** | Pointer to **int64** | Mirrors is how many are pull mirrors of github.com. | [optional] 
**Months** | Pointer to **[]int64** | Months is Commits per calendar month, oldest first, one per entry of the document&#39;s Months. | [optional] 
**Name** | Pointer to **string** | Name is the group. | [optional] 
**Orgs** | Pointer to **int64** | Orgs is how many orgs the group holds. | [optional] 
**Ours** | Pointer to **int64** | Ours is how many of those the estate&#39;s people authored: an address a member of an estate org commits with, or one at an estate domain. | [optional] 
**OursDaily** | Pointer to **map[string]int64** | OursDaily is the group&#39;s own commits per UTC day. | [optional] 
**OursMonths** | Pointer to **[]int64** | OursMonths is Ours per calendar month, as Months. | [optional] 
**People** | Pointer to **int64** | People is how many of those are the estate&#39;s own. | [optional] 
**Private** | Pointer to **int64** | Private is how many only their members can read. They are counted, never named. | [optional] 
**Public** | Pointer to **int64** | Public is how many of them anyone can read. | [optional] 
**Repos** | Pointer to **int64** | Repos is how many repositories the scope holds on the forge. | [optional] 

## Methods

### NewGitActivityGroup

`func NewGitActivityGroup() *GitActivityGroup`

NewGitActivityGroup instantiates a new GitActivityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitActivityGroupWithDefaults

`func NewGitActivityGroupWithDefaults() *GitActivityGroup`

NewGitActivityGroupWithDefaults instantiates a new GitActivityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *GitActivityGroup) GetActive() map[string]int64`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GitActivityGroup) GetActiveOk() (*map[string]int64, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GitActivityGroup) SetActive(v map[string]int64)`

SetActive sets Active field to given value.

### HasActive

`func (o *GitActivityGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetActive30

`func (o *GitActivityGroup) GetActive30() int64`

GetActive30 returns the Active30 field if non-nil, zero value otherwise.

### GetActive30Ok

`func (o *GitActivityGroup) GetActive30Ok() (*int64, bool)`

GetActive30Ok returns a tuple with the Active30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive30

`func (o *GitActivityGroup) SetActive30(v int64)`

SetActive30 sets Active30 field to given value.

### HasActive30

`func (o *GitActivityGroup) HasActive30() bool`

HasActive30 returns a boolean if a field has been set.

### GetCommits

`func (o *GitActivityGroup) GetCommits() int64`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *GitActivityGroup) GetCommitsOk() (*int64, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *GitActivityGroup) SetCommits(v int64)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *GitActivityGroup) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### GetContributors

`func (o *GitActivityGroup) GetContributors() int64`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *GitActivityGroup) GetContributorsOk() (*int64, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *GitActivityGroup) SetContributors(v int64)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *GitActivityGroup) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetDaily

`func (o *GitActivityGroup) GetDaily() map[string]int64`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *GitActivityGroup) GetDailyOk() (*map[string]int64, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *GitActivityGroup) SetDaily(v map[string]int64)`

SetDaily sets Daily field to given value.

### HasDaily

`func (o *GitActivityGroup) HasDaily() bool`

HasDaily returns a boolean if a field has been set.

### GetEmpty

`func (o *GitActivityGroup) GetEmpty() int64`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *GitActivityGroup) GetEmptyOk() (*int64, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *GitActivityGroup) SetEmpty(v int64)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *GitActivityGroup) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### GetLast

`func (o *GitActivityGroup) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GitActivityGroup) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GitActivityGroup) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *GitActivityGroup) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLatest

`func (o *GitActivityGroup) GetLatest() []GitActivityCommit`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *GitActivityGroup) GetLatestOk() (*[]GitActivityCommit, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *GitActivityGroup) SetLatest(v []GitActivityCommit)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *GitActivityGroup) HasLatest() bool`

HasLatest returns a boolean if a field has been set.

### GetMergeMonths

`func (o *GitActivityGroup) GetMergeMonths() []int64`

GetMergeMonths returns the MergeMonths field if non-nil, zero value otherwise.

### GetMergeMonthsOk

`func (o *GitActivityGroup) GetMergeMonthsOk() (*[]int64, bool)`

GetMergeMonthsOk returns a tuple with the MergeMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeMonths

`func (o *GitActivityGroup) SetMergeMonths(v []int64)`

SetMergeMonths sets MergeMonths field to given value.

### HasMergeMonths

`func (o *GitActivityGroup) HasMergeMonths() bool`

HasMergeMonths returns a boolean if a field has been set.

### GetMerges

`func (o *GitActivityGroup) GetMerges() int64`

GetMerges returns the Merges field if non-nil, zero value otherwise.

### GetMergesOk

`func (o *GitActivityGroup) GetMergesOk() (*int64, bool)`

GetMergesOk returns a tuple with the Merges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerges

`func (o *GitActivityGroup) SetMerges(v int64)`

SetMerges sets Merges field to given value.

### HasMerges

`func (o *GitActivityGroup) HasMerges() bool`

HasMerges returns a boolean if a field has been set.

### GetMirrors

`func (o *GitActivityGroup) GetMirrors() int64`

GetMirrors returns the Mirrors field if non-nil, zero value otherwise.

### GetMirrorsOk

`func (o *GitActivityGroup) GetMirrorsOk() (*int64, bool)`

GetMirrorsOk returns a tuple with the Mirrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrors

`func (o *GitActivityGroup) SetMirrors(v int64)`

SetMirrors sets Mirrors field to given value.

### HasMirrors

`func (o *GitActivityGroup) HasMirrors() bool`

HasMirrors returns a boolean if a field has been set.

### GetMonths

`func (o *GitActivityGroup) GetMonths() []int64`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *GitActivityGroup) GetMonthsOk() (*[]int64, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *GitActivityGroup) SetMonths(v []int64)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *GitActivityGroup) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### GetName

`func (o *GitActivityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitActivityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitActivityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitActivityGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgs

`func (o *GitActivityGroup) GetOrgs() int64`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *GitActivityGroup) GetOrgsOk() (*int64, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *GitActivityGroup) SetOrgs(v int64)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *GitActivityGroup) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetOurs

`func (o *GitActivityGroup) GetOurs() int64`

GetOurs returns the Ours field if non-nil, zero value otherwise.

### GetOursOk

`func (o *GitActivityGroup) GetOursOk() (*int64, bool)`

GetOursOk returns a tuple with the Ours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOurs

`func (o *GitActivityGroup) SetOurs(v int64)`

SetOurs sets Ours field to given value.

### HasOurs

`func (o *GitActivityGroup) HasOurs() bool`

HasOurs returns a boolean if a field has been set.

### GetOursDaily

`func (o *GitActivityGroup) GetOursDaily() map[string]int64`

GetOursDaily returns the OursDaily field if non-nil, zero value otherwise.

### GetOursDailyOk

`func (o *GitActivityGroup) GetOursDailyOk() (*map[string]int64, bool)`

GetOursDailyOk returns a tuple with the OursDaily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOursDaily

`func (o *GitActivityGroup) SetOursDaily(v map[string]int64)`

SetOursDaily sets OursDaily field to given value.

### HasOursDaily

`func (o *GitActivityGroup) HasOursDaily() bool`

HasOursDaily returns a boolean if a field has been set.

### GetOursMonths

`func (o *GitActivityGroup) GetOursMonths() []int64`

GetOursMonths returns the OursMonths field if non-nil, zero value otherwise.

### GetOursMonthsOk

`func (o *GitActivityGroup) GetOursMonthsOk() (*[]int64, bool)`

GetOursMonthsOk returns a tuple with the OursMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOursMonths

`func (o *GitActivityGroup) SetOursMonths(v []int64)`

SetOursMonths sets OursMonths field to given value.

### HasOursMonths

`func (o *GitActivityGroup) HasOursMonths() bool`

HasOursMonths returns a boolean if a field has been set.

### GetPeople

`func (o *GitActivityGroup) GetPeople() int64`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *GitActivityGroup) GetPeopleOk() (*int64, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *GitActivityGroup) SetPeople(v int64)`

SetPeople sets People field to given value.

### HasPeople

`func (o *GitActivityGroup) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetPrivate

`func (o *GitActivityGroup) GetPrivate() int64`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *GitActivityGroup) GetPrivateOk() (*int64, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *GitActivityGroup) SetPrivate(v int64)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *GitActivityGroup) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPublic

`func (o *GitActivityGroup) GetPublic() int64`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *GitActivityGroup) GetPublicOk() (*int64, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *GitActivityGroup) SetPublic(v int64)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *GitActivityGroup) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRepos

`func (o *GitActivityGroup) GetRepos() int64`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *GitActivityGroup) GetReposOk() (*int64, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *GitActivityGroup) SetRepos(v int64)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *GitActivityGroup) HasRepos() bool`

HasRepos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



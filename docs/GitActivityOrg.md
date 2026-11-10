# GitActivityOrg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int64** | Active is how many have a commit in the window. | [optional] 
**Active30** | Pointer to **int64** | Active30 is how many have a commit in the last 30 days. | [optional] 
**Commits** | Pointer to **int64** | Commits is distinct commits in the window, each counted once in this scope: every commit the scope&#39;s repositories hold, the upstream projects carried and kept current included. | [optional] 
**Contributors** | Pointer to **int64** | Contributors is distinct author addresses in the window, counted and never named. | [optional] 
**Empty** | Pointer to **int64** | Empty is how many hold no commit. | [optional] 
**Forge** | Pointer to **string** | Forge is the org&#39;s page on the forge. | [optional] 
**Github** | Pointer to **string** | GitHub is the org&#39;s page on GitHub; empty for a namespace only the forge has. | [optional] 
**Group** | Pointer to **string** | Group is the group that owns it. | [optional] 
**Last** | Pointer to **string** | Last is the newest commit time in the scope, RFC 3339 UTC; empty when none. | [optional] 
**MergeMonths** | Pointer to **[]int64** | MergeMonths is Merges per calendar month, as Months. | [optional] 
**Merges** | Pointer to **int64** | Merges is how many of those have more than one parent. | [optional] 
**Mirrors** | Pointer to **int64** | Mirrors is how many are pull mirrors of github.com. | [optional] 
**Months** | Pointer to **[]int64** | Months is Commits per calendar month, oldest first, one per entry of the document&#39;s Months. | [optional] 
**Name** | Pointer to **string** | Name is the org, as GitHub and the forge both spell it. | [optional] 
**Ours** | Pointer to **int64** | Ours is how many of those the estate&#39;s people authored: an address a member of an estate org commits with, or one at an estate domain. | [optional] 
**OursMonths** | Pointer to **[]int64** | OursMonths is Ours per calendar month, as Months. | [optional] 
**People** | Pointer to **int64** | People is how many of those are the estate&#39;s own. | [optional] 
**Private** | Pointer to **int64** | Private is how many only their members can read. They are counted, never named. | [optional] 
**Public** | Pointer to **int64** | Public is how many of them anyone can read. | [optional] 
**Repos** | Pointer to **int64** | Repos is how many repositories the scope holds on the forge. | [optional] 

## Methods

### NewGitActivityOrg

`func NewGitActivityOrg() *GitActivityOrg`

NewGitActivityOrg instantiates a new GitActivityOrg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitActivityOrgWithDefaults

`func NewGitActivityOrgWithDefaults() *GitActivityOrg`

NewGitActivityOrgWithDefaults instantiates a new GitActivityOrg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *GitActivityOrg) GetActive() int64`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GitActivityOrg) GetActiveOk() (*int64, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GitActivityOrg) SetActive(v int64)`

SetActive sets Active field to given value.

### HasActive

`func (o *GitActivityOrg) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetActive30

`func (o *GitActivityOrg) GetActive30() int64`

GetActive30 returns the Active30 field if non-nil, zero value otherwise.

### GetActive30Ok

`func (o *GitActivityOrg) GetActive30Ok() (*int64, bool)`

GetActive30Ok returns a tuple with the Active30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive30

`func (o *GitActivityOrg) SetActive30(v int64)`

SetActive30 sets Active30 field to given value.

### HasActive30

`func (o *GitActivityOrg) HasActive30() bool`

HasActive30 returns a boolean if a field has been set.

### GetCommits

`func (o *GitActivityOrg) GetCommits() int64`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *GitActivityOrg) GetCommitsOk() (*int64, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *GitActivityOrg) SetCommits(v int64)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *GitActivityOrg) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### GetContributors

`func (o *GitActivityOrg) GetContributors() int64`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *GitActivityOrg) GetContributorsOk() (*int64, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *GitActivityOrg) SetContributors(v int64)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *GitActivityOrg) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetEmpty

`func (o *GitActivityOrg) GetEmpty() int64`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *GitActivityOrg) GetEmptyOk() (*int64, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *GitActivityOrg) SetEmpty(v int64)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *GitActivityOrg) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### GetForge

`func (o *GitActivityOrg) GetForge() string`

GetForge returns the Forge field if non-nil, zero value otherwise.

### GetForgeOk

`func (o *GitActivityOrg) GetForgeOk() (*string, bool)`

GetForgeOk returns a tuple with the Forge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForge

`func (o *GitActivityOrg) SetForge(v string)`

SetForge sets Forge field to given value.

### HasForge

`func (o *GitActivityOrg) HasForge() bool`

HasForge returns a boolean if a field has been set.

### GetGithub

`func (o *GitActivityOrg) GetGithub() string`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *GitActivityOrg) GetGithubOk() (*string, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *GitActivityOrg) SetGithub(v string)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *GitActivityOrg) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetGroup

`func (o *GitActivityOrg) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GitActivityOrg) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GitActivityOrg) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GitActivityOrg) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetLast

`func (o *GitActivityOrg) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GitActivityOrg) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GitActivityOrg) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *GitActivityOrg) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetMergeMonths

`func (o *GitActivityOrg) GetMergeMonths() []int64`

GetMergeMonths returns the MergeMonths field if non-nil, zero value otherwise.

### GetMergeMonthsOk

`func (o *GitActivityOrg) GetMergeMonthsOk() (*[]int64, bool)`

GetMergeMonthsOk returns a tuple with the MergeMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeMonths

`func (o *GitActivityOrg) SetMergeMonths(v []int64)`

SetMergeMonths sets MergeMonths field to given value.

### HasMergeMonths

`func (o *GitActivityOrg) HasMergeMonths() bool`

HasMergeMonths returns a boolean if a field has been set.

### GetMerges

`func (o *GitActivityOrg) GetMerges() int64`

GetMerges returns the Merges field if non-nil, zero value otherwise.

### GetMergesOk

`func (o *GitActivityOrg) GetMergesOk() (*int64, bool)`

GetMergesOk returns a tuple with the Merges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerges

`func (o *GitActivityOrg) SetMerges(v int64)`

SetMerges sets Merges field to given value.

### HasMerges

`func (o *GitActivityOrg) HasMerges() bool`

HasMerges returns a boolean if a field has been set.

### GetMirrors

`func (o *GitActivityOrg) GetMirrors() int64`

GetMirrors returns the Mirrors field if non-nil, zero value otherwise.

### GetMirrorsOk

`func (o *GitActivityOrg) GetMirrorsOk() (*int64, bool)`

GetMirrorsOk returns a tuple with the Mirrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrors

`func (o *GitActivityOrg) SetMirrors(v int64)`

SetMirrors sets Mirrors field to given value.

### HasMirrors

`func (o *GitActivityOrg) HasMirrors() bool`

HasMirrors returns a boolean if a field has been set.

### GetMonths

`func (o *GitActivityOrg) GetMonths() []int64`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *GitActivityOrg) GetMonthsOk() (*[]int64, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *GitActivityOrg) SetMonths(v []int64)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *GitActivityOrg) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### GetName

`func (o *GitActivityOrg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitActivityOrg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitActivityOrg) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitActivityOrg) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOurs

`func (o *GitActivityOrg) GetOurs() int64`

GetOurs returns the Ours field if non-nil, zero value otherwise.

### GetOursOk

`func (o *GitActivityOrg) GetOursOk() (*int64, bool)`

GetOursOk returns a tuple with the Ours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOurs

`func (o *GitActivityOrg) SetOurs(v int64)`

SetOurs sets Ours field to given value.

### HasOurs

`func (o *GitActivityOrg) HasOurs() bool`

HasOurs returns a boolean if a field has been set.

### GetOursMonths

`func (o *GitActivityOrg) GetOursMonths() []int64`

GetOursMonths returns the OursMonths field if non-nil, zero value otherwise.

### GetOursMonthsOk

`func (o *GitActivityOrg) GetOursMonthsOk() (*[]int64, bool)`

GetOursMonthsOk returns a tuple with the OursMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOursMonths

`func (o *GitActivityOrg) SetOursMonths(v []int64)`

SetOursMonths sets OursMonths field to given value.

### HasOursMonths

`func (o *GitActivityOrg) HasOursMonths() bool`

HasOursMonths returns a boolean if a field has been set.

### GetPeople

`func (o *GitActivityOrg) GetPeople() int64`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *GitActivityOrg) GetPeopleOk() (*int64, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *GitActivityOrg) SetPeople(v int64)`

SetPeople sets People field to given value.

### HasPeople

`func (o *GitActivityOrg) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetPrivate

`func (o *GitActivityOrg) GetPrivate() int64`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *GitActivityOrg) GetPrivateOk() (*int64, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *GitActivityOrg) SetPrivate(v int64)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *GitActivityOrg) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPublic

`func (o *GitActivityOrg) GetPublic() int64`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *GitActivityOrg) GetPublicOk() (*int64, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *GitActivityOrg) SetPublic(v int64)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *GitActivityOrg) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRepos

`func (o *GitActivityOrg) GetRepos() int64`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *GitActivityOrg) GetReposOk() (*int64, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *GitActivityOrg) SetRepos(v int64)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *GitActivityOrg) HasRepos() bool`

HasRepos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



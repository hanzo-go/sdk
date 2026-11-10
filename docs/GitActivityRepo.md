# GitActivityRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active30** | Pointer to **bool** | Active30 is whether it has a commit in the last 30 days. | [optional] 
**Commits** | Pointer to **int64** | Commits is its distinct commits in the window. | [optional] 
**Description** | Pointer to **string** | Description is the repository&#39;s own. | [optional] 
**Forge** | Pointer to **string** | Forge is its page on the forge. | [optional] 
**Fork** | Pointer to **bool** | Fork is whether GitHub calls it a fork. | [optional] 
**Github** | Pointer to **string** | GitHub is its page on GitHub when GitHub publishes it. | [optional] 
**Group** | Pointer to **string** | Group is the group that owns the org. | [optional] 
**Last** | Pointer to **string** | Last is its newest commit time, RFC 3339 UTC; empty when none in the window. | [optional] 
**Merges** | Pointer to **int64** | Merges is how many of those are merges. | [optional] 
**Mirror** | Pointer to **bool** | Mirror is whether the forge pulls it from github.com. | [optional] 
**Months** | Pointer to **[]int64** | Months is Commits per calendar month, oldest first, one per entry of the document&#39;s Months. | [optional] 
**Org** | Pointer to **string** | Org is the org that holds it. | [optional] 
**Ours** | Pointer to **int64** | Ours is how many of those the estate&#39;s people authored. | [optional] 
**Repo** | Pointer to **string** | Repo is org/name. | [optional] 
**Stars** | Pointer to **int64** | Stars is its stars on the forge. | [optional] 

## Methods

### NewGitActivityRepo

`func NewGitActivityRepo() *GitActivityRepo`

NewGitActivityRepo instantiates a new GitActivityRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitActivityRepoWithDefaults

`func NewGitActivityRepoWithDefaults() *GitActivityRepo`

NewGitActivityRepoWithDefaults instantiates a new GitActivityRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive30

`func (o *GitActivityRepo) GetActive30() bool`

GetActive30 returns the Active30 field if non-nil, zero value otherwise.

### GetActive30Ok

`func (o *GitActivityRepo) GetActive30Ok() (*bool, bool)`

GetActive30Ok returns a tuple with the Active30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive30

`func (o *GitActivityRepo) SetActive30(v bool)`

SetActive30 sets Active30 field to given value.

### HasActive30

`func (o *GitActivityRepo) HasActive30() bool`

HasActive30 returns a boolean if a field has been set.

### GetCommits

`func (o *GitActivityRepo) GetCommits() int64`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *GitActivityRepo) GetCommitsOk() (*int64, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *GitActivityRepo) SetCommits(v int64)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *GitActivityRepo) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### GetDescription

`func (o *GitActivityRepo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GitActivityRepo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GitActivityRepo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GitActivityRepo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForge

`func (o *GitActivityRepo) GetForge() string`

GetForge returns the Forge field if non-nil, zero value otherwise.

### GetForgeOk

`func (o *GitActivityRepo) GetForgeOk() (*string, bool)`

GetForgeOk returns a tuple with the Forge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForge

`func (o *GitActivityRepo) SetForge(v string)`

SetForge sets Forge field to given value.

### HasForge

`func (o *GitActivityRepo) HasForge() bool`

HasForge returns a boolean if a field has been set.

### GetFork

`func (o *GitActivityRepo) GetFork() bool`

GetFork returns the Fork field if non-nil, zero value otherwise.

### GetForkOk

`func (o *GitActivityRepo) GetForkOk() (*bool, bool)`

GetForkOk returns a tuple with the Fork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFork

`func (o *GitActivityRepo) SetFork(v bool)`

SetFork sets Fork field to given value.

### HasFork

`func (o *GitActivityRepo) HasFork() bool`

HasFork returns a boolean if a field has been set.

### GetGithub

`func (o *GitActivityRepo) GetGithub() string`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *GitActivityRepo) GetGithubOk() (*string, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *GitActivityRepo) SetGithub(v string)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *GitActivityRepo) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetGroup

`func (o *GitActivityRepo) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GitActivityRepo) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GitActivityRepo) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GitActivityRepo) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetLast

`func (o *GitActivityRepo) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GitActivityRepo) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GitActivityRepo) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *GitActivityRepo) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetMerges

`func (o *GitActivityRepo) GetMerges() int64`

GetMerges returns the Merges field if non-nil, zero value otherwise.

### GetMergesOk

`func (o *GitActivityRepo) GetMergesOk() (*int64, bool)`

GetMergesOk returns a tuple with the Merges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerges

`func (o *GitActivityRepo) SetMerges(v int64)`

SetMerges sets Merges field to given value.

### HasMerges

`func (o *GitActivityRepo) HasMerges() bool`

HasMerges returns a boolean if a field has been set.

### GetMirror

`func (o *GitActivityRepo) GetMirror() bool`

GetMirror returns the Mirror field if non-nil, zero value otherwise.

### GetMirrorOk

`func (o *GitActivityRepo) GetMirrorOk() (*bool, bool)`

GetMirrorOk returns a tuple with the Mirror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirror

`func (o *GitActivityRepo) SetMirror(v bool)`

SetMirror sets Mirror field to given value.

### HasMirror

`func (o *GitActivityRepo) HasMirror() bool`

HasMirror returns a boolean if a field has been set.

### GetMonths

`func (o *GitActivityRepo) GetMonths() []int64`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *GitActivityRepo) GetMonthsOk() (*[]int64, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *GitActivityRepo) SetMonths(v []int64)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *GitActivityRepo) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### GetOrg

`func (o *GitActivityRepo) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *GitActivityRepo) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *GitActivityRepo) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *GitActivityRepo) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetOurs

`func (o *GitActivityRepo) GetOurs() int64`

GetOurs returns the Ours field if non-nil, zero value otherwise.

### GetOursOk

`func (o *GitActivityRepo) GetOursOk() (*int64, bool)`

GetOursOk returns a tuple with the Ours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOurs

`func (o *GitActivityRepo) SetOurs(v int64)`

SetOurs sets Ours field to given value.

### HasOurs

`func (o *GitActivityRepo) HasOurs() bool`

HasOurs returns a boolean if a field has been set.

### GetRepo

`func (o *GitActivityRepo) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *GitActivityRepo) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *GitActivityRepo) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *GitActivityRepo) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetStars

`func (o *GitActivityRepo) GetStars() int64`

GetStars returns the Stars field if non-nil, zero value otherwise.

### GetStarsOk

`func (o *GitActivityRepo) GetStarsOk() (*int64, bool)`

GetStarsOk returns a tuple with the Stars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStars

`func (o *GitActivityRepo) SetStars(v int64)`

SetStars sets Stars field to given value.

### HasStars

`func (o *GitActivityRepo) HasStars() bool`

HasStars returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



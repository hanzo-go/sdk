# GitActivityOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **map[string]int64** | Active is how many repositories have a commit on each UTC day. | [optional] 
**Complete** | Pointer to **bool** | Complete is false while a repository the forge lists has never been read. | [optional] 
**Daily** | Pointer to **map[string]int64** | Daily is the estate&#39;s distinct commits per UTC day, YYYY-MM-DD. | [optional] 
**Github** | Pointer to **string** | GitHub is when the mirror last read GitHub, RFC 3339 UTC. | [optional] 
**Groups** | Pointer to [**[]GitActivityGroup**](GitActivityGroup.md) | Groups is each group, most commits first. | [optional] 
**Latest** | Pointer to [**[]GitActivityCommit**](GitActivityCommit.md) | Latest is the newest commits of public repositories, newest first. | [optional] 
**Months** | Pointer to **[]string** | Months names the calendar months every Months series counts, YYYY-MM, oldest first. | [optional] 
**Orgs** | Pointer to [**[]GitActivityOrg**](GitActivityOrg.md) | Orgs is each org, most commits first. | [optional] 
**OursDaily** | Pointer to **map[string]int64** | OursDaily is the estate&#39;s own commits per UTC day. | [optional] 
**Read** | Pointer to **string** | Read is when the index wrote this, RFC 3339 UTC. | [optional] 
**Repos** | Pointer to [**GitActivityPage**](GitActivityPage.md) | Repos is the page of public repositories asked for. | [optional] 
**Source** | Pointer to **string** | Source names what was read: \&quot;forge\&quot;, its repositories and database. | [optional] 
**Totals** | Pointer to [**GitActivityTotals**](GitActivityTotals.md) | Totals is the whole estate. | [optional] 
**Unread** | Pointer to **int64** | Unread is how many such repositories there are. | [optional] 
**Window** | Pointer to [**GitActivityWindow**](GitActivityWindow.md) | Window is the span counted. | [optional] 

## Methods

### NewGitActivityOut

`func NewGitActivityOut() *GitActivityOut`

NewGitActivityOut instantiates a new GitActivityOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitActivityOutWithDefaults

`func NewGitActivityOutWithDefaults() *GitActivityOut`

NewGitActivityOutWithDefaults instantiates a new GitActivityOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *GitActivityOut) GetActive() map[string]int64`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GitActivityOut) GetActiveOk() (*map[string]int64, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GitActivityOut) SetActive(v map[string]int64)`

SetActive sets Active field to given value.

### HasActive

`func (o *GitActivityOut) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetComplete

`func (o *GitActivityOut) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *GitActivityOut) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *GitActivityOut) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *GitActivityOut) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetDaily

`func (o *GitActivityOut) GetDaily() map[string]int64`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *GitActivityOut) GetDailyOk() (*map[string]int64, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *GitActivityOut) SetDaily(v map[string]int64)`

SetDaily sets Daily field to given value.

### HasDaily

`func (o *GitActivityOut) HasDaily() bool`

HasDaily returns a boolean if a field has been set.

### GetGithub

`func (o *GitActivityOut) GetGithub() string`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *GitActivityOut) GetGithubOk() (*string, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *GitActivityOut) SetGithub(v string)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *GitActivityOut) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetGroups

`func (o *GitActivityOut) GetGroups() []GitActivityGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GitActivityOut) GetGroupsOk() (*[]GitActivityGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GitActivityOut) SetGroups(v []GitActivityGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *GitActivityOut) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetLatest

`func (o *GitActivityOut) GetLatest() []GitActivityCommit`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *GitActivityOut) GetLatestOk() (*[]GitActivityCommit, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *GitActivityOut) SetLatest(v []GitActivityCommit)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *GitActivityOut) HasLatest() bool`

HasLatest returns a boolean if a field has been set.

### GetMonths

`func (o *GitActivityOut) GetMonths() []string`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *GitActivityOut) GetMonthsOk() (*[]string, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *GitActivityOut) SetMonths(v []string)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *GitActivityOut) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### GetOrgs

`func (o *GitActivityOut) GetOrgs() []GitActivityOrg`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *GitActivityOut) GetOrgsOk() (*[]GitActivityOrg, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *GitActivityOut) SetOrgs(v []GitActivityOrg)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *GitActivityOut) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetOursDaily

`func (o *GitActivityOut) GetOursDaily() map[string]int64`

GetOursDaily returns the OursDaily field if non-nil, zero value otherwise.

### GetOursDailyOk

`func (o *GitActivityOut) GetOursDailyOk() (*map[string]int64, bool)`

GetOursDailyOk returns a tuple with the OursDaily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOursDaily

`func (o *GitActivityOut) SetOursDaily(v map[string]int64)`

SetOursDaily sets OursDaily field to given value.

### HasOursDaily

`func (o *GitActivityOut) HasOursDaily() bool`

HasOursDaily returns a boolean if a field has been set.

### GetRead

`func (o *GitActivityOut) GetRead() string`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *GitActivityOut) GetReadOk() (*string, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *GitActivityOut) SetRead(v string)`

SetRead sets Read field to given value.

### HasRead

`func (o *GitActivityOut) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetRepos

`func (o *GitActivityOut) GetRepos() GitActivityPage`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *GitActivityOut) GetReposOk() (*GitActivityPage, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *GitActivityOut) SetRepos(v GitActivityPage)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *GitActivityOut) HasRepos() bool`

HasRepos returns a boolean if a field has been set.

### GetSource

`func (o *GitActivityOut) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GitActivityOut) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GitActivityOut) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GitActivityOut) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTotals

`func (o *GitActivityOut) GetTotals() GitActivityTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *GitActivityOut) GetTotalsOk() (*GitActivityTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *GitActivityOut) SetTotals(v GitActivityTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *GitActivityOut) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetUnread

`func (o *GitActivityOut) GetUnread() int64`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *GitActivityOut) GetUnreadOk() (*int64, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *GitActivityOut) SetUnread(v int64)`

SetUnread sets Unread field to given value.

### HasUnread

`func (o *GitActivityOut) HasUnread() bool`

HasUnread returns a boolean if a field has been set.

### GetWindow

`func (o *GitActivityOut) GetWindow() GitActivityWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *GitActivityOut) GetWindowOk() (*GitActivityWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *GitActivityOut) SetWindow(v GitActivityWindow)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *GitActivityOut) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



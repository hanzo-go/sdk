# GitPullView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** | Author is the user who opened it; empty for a caller with no user. | [optional] 
**Base** | Pointer to **string** | Base is the branch the work is proposed into. | [optional] 
**Body** | Pointer to **string** | Body is the longer description; empty when none was given. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is RFC 3339 UTC. | [optional] 
**Head** | Pointer to **string** | Head is the branch holding the work. | [optional] 
**MergedRev** | Pointer to **string** | MergedRev is what base points at now that the merge landed. Empty while the proposal is open. | [optional] 
**Number** | Pointer to **int64** | Number is the proposal&#39;s per-repo handle, dense from 1. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository the proposal belongs to. | [optional] 
**State** | Pointer to **string** | State is \&quot;open\&quot; or \&quot;merged\&quot;. | [optional] 
**Title** | Pointer to **string** | Title is the one-line summary. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is RFC 3339 UTC. | [optional] 

## Methods

### NewGitPullView

`func NewGitPullView() *GitPullView`

NewGitPullView instantiates a new GitPullView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitPullViewWithDefaults

`func NewGitPullViewWithDefaults() *GitPullView`

NewGitPullViewWithDefaults instantiates a new GitPullView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *GitPullView) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GitPullView) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GitPullView) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *GitPullView) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBase

`func (o *GitPullView) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *GitPullView) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *GitPullView) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *GitPullView) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetBody

`func (o *GitPullView) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GitPullView) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GitPullView) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *GitPullView) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GitPullView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GitPullView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GitPullView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GitPullView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHead

`func (o *GitPullView) GetHead() string`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *GitPullView) GetHeadOk() (*string, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *GitPullView) SetHead(v string)`

SetHead sets Head field to given value.

### HasHead

`func (o *GitPullView) HasHead() bool`

HasHead returns a boolean if a field has been set.

### GetMergedRev

`func (o *GitPullView) GetMergedRev() string`

GetMergedRev returns the MergedRev field if non-nil, zero value otherwise.

### GetMergedRevOk

`func (o *GitPullView) GetMergedRevOk() (*string, bool)`

GetMergedRevOk returns a tuple with the MergedRev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedRev

`func (o *GitPullView) SetMergedRev(v string)`

SetMergedRev sets MergedRev field to given value.

### HasMergedRev

`func (o *GitPullView) HasMergedRev() bool`

HasMergedRev returns a boolean if a field has been set.

### GetNumber

`func (o *GitPullView) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GitPullView) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GitPullView) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GitPullView) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetRepo

`func (o *GitPullView) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *GitPullView) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *GitPullView) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *GitPullView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetState

`func (o *GitPullView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GitPullView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GitPullView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GitPullView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTitle

`func (o *GitPullView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GitPullView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GitPullView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GitPullView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GitPullView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GitPullView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GitPullView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GitPullView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



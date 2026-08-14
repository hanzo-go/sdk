# PullView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** | Author is the user who opened it; empty for a caller with no user. | [optional] 
**Base** | Pointer to **string** | Base is the branch the work is proposed into. | [optional] 
**Body** | Pointer to **string** | Body is the longer description; empty when none was given. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is RFC 3339 UTC. | [optional] 
**Head** | Pointer to **string** | Head is the branch holding the work. | [optional] 
**MergedRev** | Pointer to **string** | MergedRev is what base points at now that the merge landed. Empty while the proposal is open. | [optional] 
**Number** | Pointer to **int32** | Number is the proposal&#39;s per-repo handle, dense from 1. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository the proposal belongs to. | [optional] 
**State** | Pointer to **string** | State is \&quot;open\&quot; or \&quot;merged\&quot;. | [optional] 
**Title** | Pointer to **string** | Title is the one-line summary. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is RFC 3339 UTC. | [optional] 

## Methods

### NewPullView

`func NewPullView() *PullView`

NewPullView instantiates a new PullView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullViewWithDefaults

`func NewPullViewWithDefaults() *PullView`

NewPullViewWithDefaults instantiates a new PullView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *PullView) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *PullView) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *PullView) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *PullView) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBase

`func (o *PullView) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *PullView) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *PullView) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *PullView) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetBody

`func (o *PullView) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PullView) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PullView) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PullView) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PullView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PullView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PullView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PullView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHead

`func (o *PullView) GetHead() string`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *PullView) GetHeadOk() (*string, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *PullView) SetHead(v string)`

SetHead sets Head field to given value.

### HasHead

`func (o *PullView) HasHead() bool`

HasHead returns a boolean if a field has been set.

### GetMergedRev

`func (o *PullView) GetMergedRev() string`

GetMergedRev returns the MergedRev field if non-nil, zero value otherwise.

### GetMergedRevOk

`func (o *PullView) GetMergedRevOk() (*string, bool)`

GetMergedRevOk returns a tuple with the MergedRev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedRev

`func (o *PullView) SetMergedRev(v string)`

SetMergedRev sets MergedRev field to given value.

### HasMergedRev

`func (o *PullView) HasMergedRev() bool`

HasMergedRev returns a boolean if a field has been set.

### GetNumber

`func (o *PullView) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PullView) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PullView) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PullView) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetRepo

`func (o *PullView) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *PullView) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *PullView) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *PullView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetState

`func (o *PullView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PullView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PullView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PullView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTitle

`func (o *PullView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PullView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PullView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PullView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PullView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PullView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PullView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PullView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



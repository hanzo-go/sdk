# AuthorsAuthorDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAuthor** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**GithubLogin** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**VerifyCode** | Pointer to **string** |  | [optional] 
**VerifyFile** | Pointer to **string** |  | [optional] 
**VerifySnippet** | Pointer to **string** | The hanzo.json body to place on the default branch for file verification. | [optional] 
**ShareBps** | Pointer to **int64** |  | [optional] 
**BadgeBase** | Pointer to **string** |  | [optional] 
**Repos** | Pointer to [**[]AuthorsRepoView**](AuthorsRepoView.md) |  | [optional] 
**Deploys** | Pointer to [**[]AuthorsDeployView**](AuthorsDeployView.md) |  | [optional] 
**AccruedCents** | Pointer to **int64** |  | [optional] 
**PendingCents** | Pointer to **int64** |  | [optional] 
**PaidCents** | Pointer to **int64** |  | [optional] 
**Payouts** | Pointer to [**[]AuthorsPayoutView**](AuthorsPayoutView.md) |  | [optional] 

## Methods

### NewAuthorsAuthorDashboard

`func NewAuthorsAuthorDashboard() *AuthorsAuthorDashboard`

NewAuthorsAuthorDashboard instantiates a new AuthorsAuthorDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsAuthorDashboardWithDefaults

`func NewAuthorsAuthorDashboardWithDefaults() *AuthorsAuthorDashboard`

NewAuthorsAuthorDashboardWithDefaults instantiates a new AuthorsAuthorDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAuthor

`func (o *AuthorsAuthorDashboard) GetIsAuthor() bool`

GetIsAuthor returns the IsAuthor field if non-nil, zero value otherwise.

### GetIsAuthorOk

`func (o *AuthorsAuthorDashboard) GetIsAuthorOk() (*bool, bool)`

GetIsAuthorOk returns a tuple with the IsAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthor

`func (o *AuthorsAuthorDashboard) SetIsAuthor(v bool)`

SetIsAuthor sets IsAuthor field to given value.

### HasIsAuthor

`func (o *AuthorsAuthorDashboard) HasIsAuthor() bool`

HasIsAuthor returns a boolean if a field has been set.

### GetId

`func (o *AuthorsAuthorDashboard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorsAuthorDashboard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorsAuthorDashboard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorsAuthorDashboard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AuthorsAuthorDashboard) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthorsAuthorDashboard) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthorsAuthorDashboard) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthorsAuthorDashboard) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGithubLogin

`func (o *AuthorsAuthorDashboard) GetGithubLogin() string`

GetGithubLogin returns the GithubLogin field if non-nil, zero value otherwise.

### GetGithubLoginOk

`func (o *AuthorsAuthorDashboard) GetGithubLoginOk() (*string, bool)`

GetGithubLoginOk returns a tuple with the GithubLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubLogin

`func (o *AuthorsAuthorDashboard) SetGithubLogin(v string)`

SetGithubLogin sets GithubLogin field to given value.

### HasGithubLogin

`func (o *AuthorsAuthorDashboard) HasGithubLogin() bool`

HasGithubLogin returns a boolean if a field has been set.

### GetVerified

`func (o *AuthorsAuthorDashboard) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *AuthorsAuthorDashboard) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *AuthorsAuthorDashboard) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *AuthorsAuthorDashboard) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetVerifyCode

`func (o *AuthorsAuthorDashboard) GetVerifyCode() string`

GetVerifyCode returns the VerifyCode field if non-nil, zero value otherwise.

### GetVerifyCodeOk

`func (o *AuthorsAuthorDashboard) GetVerifyCodeOk() (*string, bool)`

GetVerifyCodeOk returns a tuple with the VerifyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCode

`func (o *AuthorsAuthorDashboard) SetVerifyCode(v string)`

SetVerifyCode sets VerifyCode field to given value.

### HasVerifyCode

`func (o *AuthorsAuthorDashboard) HasVerifyCode() bool`

HasVerifyCode returns a boolean if a field has been set.

### GetVerifyFile

`func (o *AuthorsAuthorDashboard) GetVerifyFile() string`

GetVerifyFile returns the VerifyFile field if non-nil, zero value otherwise.

### GetVerifyFileOk

`func (o *AuthorsAuthorDashboard) GetVerifyFileOk() (*string, bool)`

GetVerifyFileOk returns a tuple with the VerifyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyFile

`func (o *AuthorsAuthorDashboard) SetVerifyFile(v string)`

SetVerifyFile sets VerifyFile field to given value.

### HasVerifyFile

`func (o *AuthorsAuthorDashboard) HasVerifyFile() bool`

HasVerifyFile returns a boolean if a field has been set.

### GetVerifySnippet

`func (o *AuthorsAuthorDashboard) GetVerifySnippet() string`

GetVerifySnippet returns the VerifySnippet field if non-nil, zero value otherwise.

### GetVerifySnippetOk

`func (o *AuthorsAuthorDashboard) GetVerifySnippetOk() (*string, bool)`

GetVerifySnippetOk returns a tuple with the VerifySnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySnippet

`func (o *AuthorsAuthorDashboard) SetVerifySnippet(v string)`

SetVerifySnippet sets VerifySnippet field to given value.

### HasVerifySnippet

`func (o *AuthorsAuthorDashboard) HasVerifySnippet() bool`

HasVerifySnippet returns a boolean if a field has been set.

### GetShareBps

`func (o *AuthorsAuthorDashboard) GetShareBps() int64`

GetShareBps returns the ShareBps field if non-nil, zero value otherwise.

### GetShareBpsOk

`func (o *AuthorsAuthorDashboard) GetShareBpsOk() (*int64, bool)`

GetShareBpsOk returns a tuple with the ShareBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareBps

`func (o *AuthorsAuthorDashboard) SetShareBps(v int64)`

SetShareBps sets ShareBps field to given value.

### HasShareBps

`func (o *AuthorsAuthorDashboard) HasShareBps() bool`

HasShareBps returns a boolean if a field has been set.

### GetBadgeBase

`func (o *AuthorsAuthorDashboard) GetBadgeBase() string`

GetBadgeBase returns the BadgeBase field if non-nil, zero value otherwise.

### GetBadgeBaseOk

`func (o *AuthorsAuthorDashboard) GetBadgeBaseOk() (*string, bool)`

GetBadgeBaseOk returns a tuple with the BadgeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadgeBase

`func (o *AuthorsAuthorDashboard) SetBadgeBase(v string)`

SetBadgeBase sets BadgeBase field to given value.

### HasBadgeBase

`func (o *AuthorsAuthorDashboard) HasBadgeBase() bool`

HasBadgeBase returns a boolean if a field has been set.

### GetRepos

`func (o *AuthorsAuthorDashboard) GetRepos() []AuthorsRepoView`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *AuthorsAuthorDashboard) GetReposOk() (*[]AuthorsRepoView, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *AuthorsAuthorDashboard) SetRepos(v []AuthorsRepoView)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *AuthorsAuthorDashboard) HasRepos() bool`

HasRepos returns a boolean if a field has been set.

### GetDeploys

`func (o *AuthorsAuthorDashboard) GetDeploys() []AuthorsDeployView`

GetDeploys returns the Deploys field if non-nil, zero value otherwise.

### GetDeploysOk

`func (o *AuthorsAuthorDashboard) GetDeploysOk() (*[]AuthorsDeployView, bool)`

GetDeploysOk returns a tuple with the Deploys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploys

`func (o *AuthorsAuthorDashboard) SetDeploys(v []AuthorsDeployView)`

SetDeploys sets Deploys field to given value.

### HasDeploys

`func (o *AuthorsAuthorDashboard) HasDeploys() bool`

HasDeploys returns a boolean if a field has been set.

### GetAccruedCents

`func (o *AuthorsAuthorDashboard) GetAccruedCents() int64`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *AuthorsAuthorDashboard) GetAccruedCentsOk() (*int64, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *AuthorsAuthorDashboard) SetAccruedCents(v int64)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *AuthorsAuthorDashboard) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetPendingCents

`func (o *AuthorsAuthorDashboard) GetPendingCents() int64`

GetPendingCents returns the PendingCents field if non-nil, zero value otherwise.

### GetPendingCentsOk

`func (o *AuthorsAuthorDashboard) GetPendingCentsOk() (*int64, bool)`

GetPendingCentsOk returns a tuple with the PendingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCents

`func (o *AuthorsAuthorDashboard) SetPendingCents(v int64)`

SetPendingCents sets PendingCents field to given value.

### HasPendingCents

`func (o *AuthorsAuthorDashboard) HasPendingCents() bool`

HasPendingCents returns a boolean if a field has been set.

### GetPaidCents

`func (o *AuthorsAuthorDashboard) GetPaidCents() int64`

GetPaidCents returns the PaidCents field if non-nil, zero value otherwise.

### GetPaidCentsOk

`func (o *AuthorsAuthorDashboard) GetPaidCentsOk() (*int64, bool)`

GetPaidCentsOk returns a tuple with the PaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCents

`func (o *AuthorsAuthorDashboard) SetPaidCents(v int64)`

SetPaidCents sets PaidCents field to given value.

### HasPaidCents

`func (o *AuthorsAuthorDashboard) HasPaidCents() bool`

HasPaidCents returns a boolean if a field has been set.

### GetPayouts

`func (o *AuthorsAuthorDashboard) GetPayouts() []AuthorsPayoutView`

GetPayouts returns the Payouts field if non-nil, zero value otherwise.

### GetPayoutsOk

`func (o *AuthorsAuthorDashboard) GetPayoutsOk() (*[]AuthorsPayoutView, bool)`

GetPayoutsOk returns a tuple with the Payouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayouts

`func (o *AuthorsAuthorDashboard) SetPayouts(v []AuthorsPayoutView)`

SetPayouts sets Payouts field to given value.

### HasPayouts

`func (o *AuthorsAuthorDashboard) HasPayouts() bool`

HasPayouts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



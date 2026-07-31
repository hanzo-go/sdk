# AuthorsGetMyAuthors200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAuthor** | Pointer to **bool** |  | [optional] 
**DefaultShareBps** | Pointer to **int64** |  | [optional] 
**BadgeBase** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**GithubLogin** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**VerifyCode** | Pointer to **string** |  | [optional] 
**VerifyFile** | Pointer to **string** |  | [optional] 
**VerifySnippet** | Pointer to **string** | The hanzo.json body to place on the default branch for file verification. | [optional] 
**ShareBps** | Pointer to **int64** |  | [optional] 
**Repos** | Pointer to [**[]AuthorsRepoView**](AuthorsRepoView.md) |  | [optional] 
**Deploys** | Pointer to [**[]AuthorsDeployView**](AuthorsDeployView.md) |  | [optional] 
**AccruedCents** | Pointer to **int64** |  | [optional] 
**PendingCents** | Pointer to **int64** |  | [optional] 
**PaidCents** | Pointer to **int64** |  | [optional] 
**Payouts** | Pointer to [**[]AuthorsPayoutView**](AuthorsPayoutView.md) |  | [optional] 

## Methods

### NewAuthorsGetMyAuthors200Response

`func NewAuthorsGetMyAuthors200Response() *AuthorsGetMyAuthors200Response`

NewAuthorsGetMyAuthors200Response instantiates a new AuthorsGetMyAuthors200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsGetMyAuthors200ResponseWithDefaults

`func NewAuthorsGetMyAuthors200ResponseWithDefaults() *AuthorsGetMyAuthors200Response`

NewAuthorsGetMyAuthors200ResponseWithDefaults instantiates a new AuthorsGetMyAuthors200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAuthor

`func (o *AuthorsGetMyAuthors200Response) GetIsAuthor() bool`

GetIsAuthor returns the IsAuthor field if non-nil, zero value otherwise.

### GetIsAuthorOk

`func (o *AuthorsGetMyAuthors200Response) GetIsAuthorOk() (*bool, bool)`

GetIsAuthorOk returns a tuple with the IsAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthor

`func (o *AuthorsGetMyAuthors200Response) SetIsAuthor(v bool)`

SetIsAuthor sets IsAuthor field to given value.

### HasIsAuthor

`func (o *AuthorsGetMyAuthors200Response) HasIsAuthor() bool`

HasIsAuthor returns a boolean if a field has been set.

### GetDefaultShareBps

`func (o *AuthorsGetMyAuthors200Response) GetDefaultShareBps() int64`

GetDefaultShareBps returns the DefaultShareBps field if non-nil, zero value otherwise.

### GetDefaultShareBpsOk

`func (o *AuthorsGetMyAuthors200Response) GetDefaultShareBpsOk() (*int64, bool)`

GetDefaultShareBpsOk returns a tuple with the DefaultShareBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultShareBps

`func (o *AuthorsGetMyAuthors200Response) SetDefaultShareBps(v int64)`

SetDefaultShareBps sets DefaultShareBps field to given value.

### HasDefaultShareBps

`func (o *AuthorsGetMyAuthors200Response) HasDefaultShareBps() bool`

HasDefaultShareBps returns a boolean if a field has been set.

### GetBadgeBase

`func (o *AuthorsGetMyAuthors200Response) GetBadgeBase() string`

GetBadgeBase returns the BadgeBase field if non-nil, zero value otherwise.

### GetBadgeBaseOk

`func (o *AuthorsGetMyAuthors200Response) GetBadgeBaseOk() (*string, bool)`

GetBadgeBaseOk returns a tuple with the BadgeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadgeBase

`func (o *AuthorsGetMyAuthors200Response) SetBadgeBase(v string)`

SetBadgeBase sets BadgeBase field to given value.

### HasBadgeBase

`func (o *AuthorsGetMyAuthors200Response) HasBadgeBase() bool`

HasBadgeBase returns a boolean if a field has been set.

### GetId

`func (o *AuthorsGetMyAuthors200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorsGetMyAuthors200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorsGetMyAuthors200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorsGetMyAuthors200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AuthorsGetMyAuthors200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthorsGetMyAuthors200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthorsGetMyAuthors200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthorsGetMyAuthors200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGithubLogin

`func (o *AuthorsGetMyAuthors200Response) GetGithubLogin() string`

GetGithubLogin returns the GithubLogin field if non-nil, zero value otherwise.

### GetGithubLoginOk

`func (o *AuthorsGetMyAuthors200Response) GetGithubLoginOk() (*string, bool)`

GetGithubLoginOk returns a tuple with the GithubLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubLogin

`func (o *AuthorsGetMyAuthors200Response) SetGithubLogin(v string)`

SetGithubLogin sets GithubLogin field to given value.

### HasGithubLogin

`func (o *AuthorsGetMyAuthors200Response) HasGithubLogin() bool`

HasGithubLogin returns a boolean if a field has been set.

### GetVerified

`func (o *AuthorsGetMyAuthors200Response) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *AuthorsGetMyAuthors200Response) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *AuthorsGetMyAuthors200Response) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *AuthorsGetMyAuthors200Response) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetVerifyCode

`func (o *AuthorsGetMyAuthors200Response) GetVerifyCode() string`

GetVerifyCode returns the VerifyCode field if non-nil, zero value otherwise.

### GetVerifyCodeOk

`func (o *AuthorsGetMyAuthors200Response) GetVerifyCodeOk() (*string, bool)`

GetVerifyCodeOk returns a tuple with the VerifyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCode

`func (o *AuthorsGetMyAuthors200Response) SetVerifyCode(v string)`

SetVerifyCode sets VerifyCode field to given value.

### HasVerifyCode

`func (o *AuthorsGetMyAuthors200Response) HasVerifyCode() bool`

HasVerifyCode returns a boolean if a field has been set.

### GetVerifyFile

`func (o *AuthorsGetMyAuthors200Response) GetVerifyFile() string`

GetVerifyFile returns the VerifyFile field if non-nil, zero value otherwise.

### GetVerifyFileOk

`func (o *AuthorsGetMyAuthors200Response) GetVerifyFileOk() (*string, bool)`

GetVerifyFileOk returns a tuple with the VerifyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyFile

`func (o *AuthorsGetMyAuthors200Response) SetVerifyFile(v string)`

SetVerifyFile sets VerifyFile field to given value.

### HasVerifyFile

`func (o *AuthorsGetMyAuthors200Response) HasVerifyFile() bool`

HasVerifyFile returns a boolean if a field has been set.

### GetVerifySnippet

`func (o *AuthorsGetMyAuthors200Response) GetVerifySnippet() string`

GetVerifySnippet returns the VerifySnippet field if non-nil, zero value otherwise.

### GetVerifySnippetOk

`func (o *AuthorsGetMyAuthors200Response) GetVerifySnippetOk() (*string, bool)`

GetVerifySnippetOk returns a tuple with the VerifySnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySnippet

`func (o *AuthorsGetMyAuthors200Response) SetVerifySnippet(v string)`

SetVerifySnippet sets VerifySnippet field to given value.

### HasVerifySnippet

`func (o *AuthorsGetMyAuthors200Response) HasVerifySnippet() bool`

HasVerifySnippet returns a boolean if a field has been set.

### GetShareBps

`func (o *AuthorsGetMyAuthors200Response) GetShareBps() int64`

GetShareBps returns the ShareBps field if non-nil, zero value otherwise.

### GetShareBpsOk

`func (o *AuthorsGetMyAuthors200Response) GetShareBpsOk() (*int64, bool)`

GetShareBpsOk returns a tuple with the ShareBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareBps

`func (o *AuthorsGetMyAuthors200Response) SetShareBps(v int64)`

SetShareBps sets ShareBps field to given value.

### HasShareBps

`func (o *AuthorsGetMyAuthors200Response) HasShareBps() bool`

HasShareBps returns a boolean if a field has been set.

### GetRepos

`func (o *AuthorsGetMyAuthors200Response) GetRepos() []AuthorsRepoView`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *AuthorsGetMyAuthors200Response) GetReposOk() (*[]AuthorsRepoView, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *AuthorsGetMyAuthors200Response) SetRepos(v []AuthorsRepoView)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *AuthorsGetMyAuthors200Response) HasRepos() bool`

HasRepos returns a boolean if a field has been set.

### GetDeploys

`func (o *AuthorsGetMyAuthors200Response) GetDeploys() []AuthorsDeployView`

GetDeploys returns the Deploys field if non-nil, zero value otherwise.

### GetDeploysOk

`func (o *AuthorsGetMyAuthors200Response) GetDeploysOk() (*[]AuthorsDeployView, bool)`

GetDeploysOk returns a tuple with the Deploys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploys

`func (o *AuthorsGetMyAuthors200Response) SetDeploys(v []AuthorsDeployView)`

SetDeploys sets Deploys field to given value.

### HasDeploys

`func (o *AuthorsGetMyAuthors200Response) HasDeploys() bool`

HasDeploys returns a boolean if a field has been set.

### GetAccruedCents

`func (o *AuthorsGetMyAuthors200Response) GetAccruedCents() int64`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *AuthorsGetMyAuthors200Response) GetAccruedCentsOk() (*int64, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *AuthorsGetMyAuthors200Response) SetAccruedCents(v int64)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *AuthorsGetMyAuthors200Response) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetPendingCents

`func (o *AuthorsGetMyAuthors200Response) GetPendingCents() int64`

GetPendingCents returns the PendingCents field if non-nil, zero value otherwise.

### GetPendingCentsOk

`func (o *AuthorsGetMyAuthors200Response) GetPendingCentsOk() (*int64, bool)`

GetPendingCentsOk returns a tuple with the PendingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCents

`func (o *AuthorsGetMyAuthors200Response) SetPendingCents(v int64)`

SetPendingCents sets PendingCents field to given value.

### HasPendingCents

`func (o *AuthorsGetMyAuthors200Response) HasPendingCents() bool`

HasPendingCents returns a boolean if a field has been set.

### GetPaidCents

`func (o *AuthorsGetMyAuthors200Response) GetPaidCents() int64`

GetPaidCents returns the PaidCents field if non-nil, zero value otherwise.

### GetPaidCentsOk

`func (o *AuthorsGetMyAuthors200Response) GetPaidCentsOk() (*int64, bool)`

GetPaidCentsOk returns a tuple with the PaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCents

`func (o *AuthorsGetMyAuthors200Response) SetPaidCents(v int64)`

SetPaidCents sets PaidCents field to given value.

### HasPaidCents

`func (o *AuthorsGetMyAuthors200Response) HasPaidCents() bool`

HasPaidCents returns a boolean if a field has been set.

### GetPayouts

`func (o *AuthorsGetMyAuthors200Response) GetPayouts() []AuthorsPayoutView`

GetPayouts returns the Payouts field if non-nil, zero value otherwise.

### GetPayoutsOk

`func (o *AuthorsGetMyAuthors200Response) GetPayoutsOk() (*[]AuthorsPayoutView, bool)`

GetPayoutsOk returns a tuple with the Payouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayouts

`func (o *AuthorsGetMyAuthors200Response) SetPayouts(v []AuthorsPayoutView)`

SetPayouts sets Payouts field to given value.

### HasPayouts

`func (o *AuthorsGetMyAuthors200Response) HasPayouts() bool`

HasPayouts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AuthorsAdminAuthorView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**GithubLogin** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**ShareBps** | Pointer to **int64** |  | [optional] 
**RepoCount** | Pointer to **int32** |  | [optional] 
**DeployCount** | Pointer to **int32** |  | [optional] 
**AccruedCents** | Pointer to **int64** |  | [optional] 
**PendingCents** | Pointer to **int64** |  | [optional] 
**PaidCents** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**ApprovedAt** | Pointer to **int64** |  | [optional] 
**SuspendedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewAuthorsAdminAuthorView

`func NewAuthorsAdminAuthorView() *AuthorsAdminAuthorView`

NewAuthorsAdminAuthorView instantiates a new AuthorsAdminAuthorView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsAdminAuthorViewWithDefaults

`func NewAuthorsAdminAuthorViewWithDefaults() *AuthorsAdminAuthorView`

NewAuthorsAdminAuthorViewWithDefaults instantiates a new AuthorsAdminAuthorView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthorsAdminAuthorView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorsAdminAuthorView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorsAdminAuthorView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorsAdminAuthorView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrg

`func (o *AuthorsAdminAuthorView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AuthorsAdminAuthorView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AuthorsAdminAuthorView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AuthorsAdminAuthorView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetGithubLogin

`func (o *AuthorsAdminAuthorView) GetGithubLogin() string`

GetGithubLogin returns the GithubLogin field if non-nil, zero value otherwise.

### GetGithubLoginOk

`func (o *AuthorsAdminAuthorView) GetGithubLoginOk() (*string, bool)`

GetGithubLoginOk returns a tuple with the GithubLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubLogin

`func (o *AuthorsAdminAuthorView) SetGithubLogin(v string)`

SetGithubLogin sets GithubLogin field to given value.

### HasGithubLogin

`func (o *AuthorsAdminAuthorView) HasGithubLogin() bool`

HasGithubLogin returns a boolean if a field has been set.

### GetStatus

`func (o *AuthorsAdminAuthorView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthorsAdminAuthorView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthorsAdminAuthorView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthorsAdminAuthorView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerified

`func (o *AuthorsAdminAuthorView) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *AuthorsAdminAuthorView) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *AuthorsAdminAuthorView) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *AuthorsAdminAuthorView) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetShareBps

`func (o *AuthorsAdminAuthorView) GetShareBps() int64`

GetShareBps returns the ShareBps field if non-nil, zero value otherwise.

### GetShareBpsOk

`func (o *AuthorsAdminAuthorView) GetShareBpsOk() (*int64, bool)`

GetShareBpsOk returns a tuple with the ShareBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareBps

`func (o *AuthorsAdminAuthorView) SetShareBps(v int64)`

SetShareBps sets ShareBps field to given value.

### HasShareBps

`func (o *AuthorsAdminAuthorView) HasShareBps() bool`

HasShareBps returns a boolean if a field has been set.

### GetRepoCount

`func (o *AuthorsAdminAuthorView) GetRepoCount() int32`

GetRepoCount returns the RepoCount field if non-nil, zero value otherwise.

### GetRepoCountOk

`func (o *AuthorsAdminAuthorView) GetRepoCountOk() (*int32, bool)`

GetRepoCountOk returns a tuple with the RepoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoCount

`func (o *AuthorsAdminAuthorView) SetRepoCount(v int32)`

SetRepoCount sets RepoCount field to given value.

### HasRepoCount

`func (o *AuthorsAdminAuthorView) HasRepoCount() bool`

HasRepoCount returns a boolean if a field has been set.

### GetDeployCount

`func (o *AuthorsAdminAuthorView) GetDeployCount() int32`

GetDeployCount returns the DeployCount field if non-nil, zero value otherwise.

### GetDeployCountOk

`func (o *AuthorsAdminAuthorView) GetDeployCountOk() (*int32, bool)`

GetDeployCountOk returns a tuple with the DeployCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployCount

`func (o *AuthorsAdminAuthorView) SetDeployCount(v int32)`

SetDeployCount sets DeployCount field to given value.

### HasDeployCount

`func (o *AuthorsAdminAuthorView) HasDeployCount() bool`

HasDeployCount returns a boolean if a field has been set.

### GetAccruedCents

`func (o *AuthorsAdminAuthorView) GetAccruedCents() int64`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *AuthorsAdminAuthorView) GetAccruedCentsOk() (*int64, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *AuthorsAdminAuthorView) SetAccruedCents(v int64)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *AuthorsAdminAuthorView) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetPendingCents

`func (o *AuthorsAdminAuthorView) GetPendingCents() int64`

GetPendingCents returns the PendingCents field if non-nil, zero value otherwise.

### GetPendingCentsOk

`func (o *AuthorsAdminAuthorView) GetPendingCentsOk() (*int64, bool)`

GetPendingCentsOk returns a tuple with the PendingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCents

`func (o *AuthorsAdminAuthorView) SetPendingCents(v int64)`

SetPendingCents sets PendingCents field to given value.

### HasPendingCents

`func (o *AuthorsAdminAuthorView) HasPendingCents() bool`

HasPendingCents returns a boolean if a field has been set.

### GetPaidCents

`func (o *AuthorsAdminAuthorView) GetPaidCents() int64`

GetPaidCents returns the PaidCents field if non-nil, zero value otherwise.

### GetPaidCentsOk

`func (o *AuthorsAdminAuthorView) GetPaidCentsOk() (*int64, bool)`

GetPaidCentsOk returns a tuple with the PaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCents

`func (o *AuthorsAdminAuthorView) SetPaidCents(v int64)`

SetPaidCents sets PaidCents field to given value.

### HasPaidCents

`func (o *AuthorsAdminAuthorView) HasPaidCents() bool`

HasPaidCents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthorsAdminAuthorView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthorsAdminAuthorView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthorsAdminAuthorView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthorsAdminAuthorView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetApprovedAt

`func (o *AuthorsAdminAuthorView) GetApprovedAt() int64`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *AuthorsAdminAuthorView) GetApprovedAtOk() (*int64, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *AuthorsAdminAuthorView) SetApprovedAt(v int64)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *AuthorsAdminAuthorView) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### GetSuspendedAt

`func (o *AuthorsAdminAuthorView) GetSuspendedAt() int64`

GetSuspendedAt returns the SuspendedAt field if non-nil, zero value otherwise.

### GetSuspendedAtOk

`func (o *AuthorsAdminAuthorView) GetSuspendedAtOk() (*int64, bool)`

GetSuspendedAtOk returns a tuple with the SuspendedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedAt

`func (o *AuthorsAdminAuthorView) SetSuspendedAt(v int64)`

SetSuspendedAt sets SuspendedAt field to given value.

### HasSuspendedAt

`func (o *AuthorsAdminAuthorView) HasSuspendedAt() bool`

HasSuspendedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



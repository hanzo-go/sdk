# AdminAuthorView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedCents** | Pointer to **int32** | AccruedCents is lifetime royalty accrued, in integer USD cents: the sum of every latched accrual (spend × shareBps / 10000). It only ever rises — a payout is recorded against paidCents and never reduces this. | [optional] 
**ApprovedAt** | Pointer to **int32** | ApprovedAt is unix seconds of the first approval, and 0 means never approved — which is also \&quot;has never been able to accrue\&quot;. Re-approving to renegotiate the share leaves it at the original date. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is unix seconds at the FIRST connect. Re-connecting re-links the login and leaves this alone, so it dates the enrolment, not the latest link. | [optional] 
**DeployCount** | Pointer to **int32** | DeployCount is how many attribution edges point at this author — one per (repository, project, deploying org), so re-deploying the same project adds none. It includes self-deploys, which are recorded for provenance and excluded from accrual, so it measures reach, not the earning set. | [optional] 
**GithubLogin** | Pointer to **string** | GithubLogin is the linked forge account, lowercased. It comes from IAM&#39;s linked account when the connect had one — which is also what sets verified — and otherwise from the login the caller declared. The treasury author carries \&quot;&lt;brand&gt;-maintainers\&quot;. | [optional] 
**Id** | Pointer to **string** | ID is the author record&#39;s server-minted handle, \&quot;aut_\&quot;-prefixed. It is the id the approve, suspend, payout and admin-basis routes address. | [optional] 
**Org** | Pointer to **string** | Org is the tenant org that owns this author record — UNIQUE, one author per org. It is exposed HERE and nowhere else (Author.Org is json:\&quot;-\&quot; on the tenant surface), and it is the org excluded from this author&#39;s own accrual: deploying your own repo earns you nothing. | [optional] 
**PaidCents** | Pointer to **int32** | PaidCents is lifetime royalty RECORDED as paid, in integer USD cents. It rises the moment a payout reserves against pending — recording, not settling; a human moves the money out of band — and falls back only when a payout is voided. | [optional] 
**PendingCents** | Pointer to **int32** | PendingCents is what a payout may still draw against — accrued − paid, floored at zero. It is derived for each response, never stored, and it is the exact figure the atomic payout guard refuses to exceed. | [optional] 
**RepoCount** | Pointer to **int32** | RepoCount is how many of this author&#39;s repository claims are VERIFIED, counted for this response in one GROUP BY over the whole table rather than a query per row. The single-author replies from approve, suspend and payout report 0: they carry the mutated row, not a re-listing. | [optional] 
**ShareBps** | Pointer to **int32** | ShareBps is the royalty rate accrual applies, in basis points of a deploying org&#39;s metered spend for the period: 2000 (the platform default) is 20%, 10000 would be the entire spend. The platform keeps 10000 − shareBps. Changing it never rewrites history — each ledger row keeps the rate it was written with. | [optional] 
**Status** | Pointer to **string** | Status is connected, approved or suspended. Only an approved author accrues; a connected one may verify repos and collect deploy edges but earns nothing until a reviewer admits it. | [optional] 
**SuspendedAt** | Pointer to **int32** | SuspendedAt is unix seconds of the most recent suspension. 0 means the author is not suspended: either never was, or was and has since been approved again, which clears this back to 0. | [optional] 
**Verified** | Pointer to **bool** | Verified is IDENTITY proof of the login, NOT proof of any repository: true when the connect took the login from IAM&#39;s linked forge account (and for the seeded treasury author), false when the caller merely declared it. A false here still earns — repository ownership is proven separately, per claim. | [optional] 

## Methods

### NewAdminAuthorView

`func NewAdminAuthorView() *AdminAuthorView`

NewAdminAuthorView instantiates a new AdminAuthorView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminAuthorViewWithDefaults

`func NewAdminAuthorViewWithDefaults() *AdminAuthorView`

NewAdminAuthorViewWithDefaults instantiates a new AdminAuthorView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedCents

`func (o *AdminAuthorView) GetAccruedCents() int32`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *AdminAuthorView) GetAccruedCentsOk() (*int32, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *AdminAuthorView) SetAccruedCents(v int32)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *AdminAuthorView) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetApprovedAt

`func (o *AdminAuthorView) GetApprovedAt() int32`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *AdminAuthorView) GetApprovedAtOk() (*int32, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *AdminAuthorView) SetApprovedAt(v int32)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *AdminAuthorView) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdminAuthorView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminAuthorView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminAuthorView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdminAuthorView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeployCount

`func (o *AdminAuthorView) GetDeployCount() int32`

GetDeployCount returns the DeployCount field if non-nil, zero value otherwise.

### GetDeployCountOk

`func (o *AdminAuthorView) GetDeployCountOk() (*int32, bool)`

GetDeployCountOk returns a tuple with the DeployCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployCount

`func (o *AdminAuthorView) SetDeployCount(v int32)`

SetDeployCount sets DeployCount field to given value.

### HasDeployCount

`func (o *AdminAuthorView) HasDeployCount() bool`

HasDeployCount returns a boolean if a field has been set.

### GetGithubLogin

`func (o *AdminAuthorView) GetGithubLogin() string`

GetGithubLogin returns the GithubLogin field if non-nil, zero value otherwise.

### GetGithubLoginOk

`func (o *AdminAuthorView) GetGithubLoginOk() (*string, bool)`

GetGithubLoginOk returns a tuple with the GithubLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubLogin

`func (o *AdminAuthorView) SetGithubLogin(v string)`

SetGithubLogin sets GithubLogin field to given value.

### HasGithubLogin

`func (o *AdminAuthorView) HasGithubLogin() bool`

HasGithubLogin returns a boolean if a field has been set.

### GetId

`func (o *AdminAuthorView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminAuthorView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminAuthorView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdminAuthorView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrg

`func (o *AdminAuthorView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AdminAuthorView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AdminAuthorView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AdminAuthorView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPaidCents

`func (o *AdminAuthorView) GetPaidCents() int32`

GetPaidCents returns the PaidCents field if non-nil, zero value otherwise.

### GetPaidCentsOk

`func (o *AdminAuthorView) GetPaidCentsOk() (*int32, bool)`

GetPaidCentsOk returns a tuple with the PaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCents

`func (o *AdminAuthorView) SetPaidCents(v int32)`

SetPaidCents sets PaidCents field to given value.

### HasPaidCents

`func (o *AdminAuthorView) HasPaidCents() bool`

HasPaidCents returns a boolean if a field has been set.

### GetPendingCents

`func (o *AdminAuthorView) GetPendingCents() int32`

GetPendingCents returns the PendingCents field if non-nil, zero value otherwise.

### GetPendingCentsOk

`func (o *AdminAuthorView) GetPendingCentsOk() (*int32, bool)`

GetPendingCentsOk returns a tuple with the PendingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCents

`func (o *AdminAuthorView) SetPendingCents(v int32)`

SetPendingCents sets PendingCents field to given value.

### HasPendingCents

`func (o *AdminAuthorView) HasPendingCents() bool`

HasPendingCents returns a boolean if a field has been set.

### GetRepoCount

`func (o *AdminAuthorView) GetRepoCount() int32`

GetRepoCount returns the RepoCount field if non-nil, zero value otherwise.

### GetRepoCountOk

`func (o *AdminAuthorView) GetRepoCountOk() (*int32, bool)`

GetRepoCountOk returns a tuple with the RepoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoCount

`func (o *AdminAuthorView) SetRepoCount(v int32)`

SetRepoCount sets RepoCount field to given value.

### HasRepoCount

`func (o *AdminAuthorView) HasRepoCount() bool`

HasRepoCount returns a boolean if a field has been set.

### GetShareBps

`func (o *AdminAuthorView) GetShareBps() int32`

GetShareBps returns the ShareBps field if non-nil, zero value otherwise.

### GetShareBpsOk

`func (o *AdminAuthorView) GetShareBpsOk() (*int32, bool)`

GetShareBpsOk returns a tuple with the ShareBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareBps

`func (o *AdminAuthorView) SetShareBps(v int32)`

SetShareBps sets ShareBps field to given value.

### HasShareBps

`func (o *AdminAuthorView) HasShareBps() bool`

HasShareBps returns a boolean if a field has been set.

### GetStatus

`func (o *AdminAuthorView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminAuthorView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminAuthorView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminAuthorView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuspendedAt

`func (o *AdminAuthorView) GetSuspendedAt() int32`

GetSuspendedAt returns the SuspendedAt field if non-nil, zero value otherwise.

### GetSuspendedAtOk

`func (o *AdminAuthorView) GetSuspendedAtOk() (*int32, bool)`

GetSuspendedAtOk returns a tuple with the SuspendedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedAt

`func (o *AdminAuthorView) SetSuspendedAt(v int32)`

SetSuspendedAt sets SuspendedAt field to given value.

### HasSuspendedAt

`func (o *AdminAuthorView) HasSuspendedAt() bool`

HasSuspendedAt returns a boolean if a field has been set.

### GetVerified

`func (o *AdminAuthorView) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *AdminAuthorView) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *AdminAuthorView) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *AdminAuthorView) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



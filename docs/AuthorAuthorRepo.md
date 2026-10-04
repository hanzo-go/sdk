# AuthorAuthorRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BadgeMarkdown** | Pointer to **string** | BadgeMarkdown is the ready-to-paste README snippet, DERIVED for each response from this deployment&#39;s badge host and never stored: a \&quot;Deploy on Hanzo\&quot; image linking to the one-click import of this repository. Re-hosting the builder changes every badge without touching a row. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is unix seconds when the claim was first recorded. It equals verifiedAt on the first proof and then stays put while verifiedAt moves, so the pair reads as \&quot;claimed since / last proven\&quot;. | [optional] 
**Method** | Pointer to **string** | Method is HOW ownership was proven: \&quot;oauth\&quot; — an IAM-linked forge token showed admin or push on the repository; \&quot;file\&quot; — a hanzo.json on the default branch carried this author&#39;s verify code; or \&quot;maintainer\&quot; — the repository sits in a first-party namespace, where ownership is intrinsic and the treasury author holds it with no proof step. Omitted on a row written before the method was recorded. | [optional] 
**RepoUrl** | Pointer to **string** | RepoURL is the claim key in canonical form — lowercased \&quot;host/owner/name\&quot;, no scheme, no .git, host ∈ {github.com, gitlab.com}. A deploy&#39;s source repo is normalized through the same function before attribution, so the two sides can never miss on a cosmetic difference. UNIQUE across every author: first proven claim wins. | [optional] 
**Verified** | Pointer to **bool** | Verified reports that ownership was proven. Only a proven claim is ever written, so it is true on every row this surface returns; the deploy path re-reads it regardless, because an unverified claim attributes nothing. | [optional] 
**VerifiedAt** | Pointer to **int64** | VerifiedAt is unix seconds of the most recent successful proof. Re-verifying refreshes it, and the method beside it, in place. | [optional] 

## Methods

### NewAuthorAuthorRepo

`func NewAuthorAuthorRepo() *AuthorAuthorRepo`

NewAuthorAuthorRepo instantiates a new AuthorAuthorRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorAuthorRepoWithDefaults

`func NewAuthorAuthorRepoWithDefaults() *AuthorAuthorRepo`

NewAuthorAuthorRepoWithDefaults instantiates a new AuthorAuthorRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBadgeMarkdown

`func (o *AuthorAuthorRepo) GetBadgeMarkdown() string`

GetBadgeMarkdown returns the BadgeMarkdown field if non-nil, zero value otherwise.

### GetBadgeMarkdownOk

`func (o *AuthorAuthorRepo) GetBadgeMarkdownOk() (*string, bool)`

GetBadgeMarkdownOk returns a tuple with the BadgeMarkdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadgeMarkdown

`func (o *AuthorAuthorRepo) SetBadgeMarkdown(v string)`

SetBadgeMarkdown sets BadgeMarkdown field to given value.

### HasBadgeMarkdown

`func (o *AuthorAuthorRepo) HasBadgeMarkdown() bool`

HasBadgeMarkdown returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthorAuthorRepo) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthorAuthorRepo) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthorAuthorRepo) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthorAuthorRepo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMethod

`func (o *AuthorAuthorRepo) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AuthorAuthorRepo) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AuthorAuthorRepo) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *AuthorAuthorRepo) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRepoUrl

`func (o *AuthorAuthorRepo) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *AuthorAuthorRepo) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *AuthorAuthorRepo) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.

### HasRepoUrl

`func (o *AuthorAuthorRepo) HasRepoUrl() bool`

HasRepoUrl returns a boolean if a field has been set.

### GetVerified

`func (o *AuthorAuthorRepo) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *AuthorAuthorRepo) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *AuthorAuthorRepo) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *AuthorAuthorRepo) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetVerifiedAt

`func (o *AuthorAuthorRepo) GetVerifiedAt() int64`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *AuthorAuthorRepo) GetVerifiedAtOk() (*int64, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *AuthorAuthorRepo) SetVerifiedAt(v int64)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *AuthorAuthorRepo) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



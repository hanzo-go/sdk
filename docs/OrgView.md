# OrgView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BadgeMarkdown** | Pointer to **string** | BadgeMarkdown is the ready-to-paste README snippet, DERIVED for each response from this deployment&#39;s badge host and never stored — here it deep-links the OWNER&#39;s template import rather than one repository&#39;s. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is unix seconds when the owner claim was first recorded — equal to verifiedAt on the first proof, then fixed while verifiedAt moves. | [optional] 
**Method** | Pointer to **string** | Method is HOW the owner was proven, always against its \&quot;.github\&quot; control repository: \&quot;oauth\&quot; — an IAM-linked forge token showed admin or push on it; or \&quot;file\&quot; — a hanzo.json on its default branch carried this author&#39;s verify code. The \&quot;maintainer\&quot; shortcut is a per-repository attribution and never appears here. Omitted on a row written before the method was recorded. | [optional] 
**OwnerUrl** | Pointer to **string** | OwnerURL is the claim key in canonical form — lowercased \&quot;host/owner\&quot; with NO repository segment, host ∈ {github.com, gitlab.com}. It covers every repository under that owner, so code with no claim of its own still earns; a per-repository claim outranks it. UNIQUE across every author: first proven claim wins. | [optional] 
**Verified** | Pointer to **bool** | Verified reports that ownership of the WHOLE owner was proven — against that owner&#39;s \&quot;.github\&quot; control repository, which is exactly as strong as a per-repository claim. Only a proven claim is written, so every row returned here is true. | [optional] 
**VerifiedAt** | Pointer to **int64** | VerifiedAt is unix seconds of the most recent successful proof of the owner; re-verifying refreshes it, and the method beside it, in place. | [optional] 

## Methods

### NewOrgView

`func NewOrgView() *OrgView`

NewOrgView instantiates a new OrgView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgViewWithDefaults

`func NewOrgViewWithDefaults() *OrgView`

NewOrgViewWithDefaults instantiates a new OrgView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBadgeMarkdown

`func (o *OrgView) GetBadgeMarkdown() string`

GetBadgeMarkdown returns the BadgeMarkdown field if non-nil, zero value otherwise.

### GetBadgeMarkdownOk

`func (o *OrgView) GetBadgeMarkdownOk() (*string, bool)`

GetBadgeMarkdownOk returns a tuple with the BadgeMarkdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadgeMarkdown

`func (o *OrgView) SetBadgeMarkdown(v string)`

SetBadgeMarkdown sets BadgeMarkdown field to given value.

### HasBadgeMarkdown

`func (o *OrgView) HasBadgeMarkdown() bool`

HasBadgeMarkdown returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrgView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrgView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrgView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrgView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMethod

`func (o *OrgView) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OrgView) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OrgView) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *OrgView) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetOwnerUrl

`func (o *OrgView) GetOwnerUrl() string`

GetOwnerUrl returns the OwnerUrl field if non-nil, zero value otherwise.

### GetOwnerUrlOk

`func (o *OrgView) GetOwnerUrlOk() (*string, bool)`

GetOwnerUrlOk returns a tuple with the OwnerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUrl

`func (o *OrgView) SetOwnerUrl(v string)`

SetOwnerUrl sets OwnerUrl field to given value.

### HasOwnerUrl

`func (o *OrgView) HasOwnerUrl() bool`

HasOwnerUrl returns a boolean if a field has been set.

### GetVerified

`func (o *OrgView) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *OrgView) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *OrgView) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *OrgView) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetVerifiedAt

`func (o *OrgView) GetVerifiedAt() int64`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *OrgView) GetVerifiedAtOk() (*int64, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *OrgView) SetVerifiedAt(v int64)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *OrgView) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



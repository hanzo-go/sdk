# AuthorsRepoView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepoUrl** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**Method** | Pointer to **string** | Verification method that succeeded (oauth or file). | [optional] 
**BadgeMarkdown** | Pointer to **string** | Ready-to-paste \&quot;Deploy on Hanzo\&quot; README snippet. | [optional] 
**VerifiedAt** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewAuthorsRepoView

`func NewAuthorsRepoView() *AuthorsRepoView`

NewAuthorsRepoView instantiates a new AuthorsRepoView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsRepoViewWithDefaults

`func NewAuthorsRepoViewWithDefaults() *AuthorsRepoView`

NewAuthorsRepoViewWithDefaults instantiates a new AuthorsRepoView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepoUrl

`func (o *AuthorsRepoView) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *AuthorsRepoView) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *AuthorsRepoView) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.

### HasRepoUrl

`func (o *AuthorsRepoView) HasRepoUrl() bool`

HasRepoUrl returns a boolean if a field has been set.

### GetVerified

`func (o *AuthorsRepoView) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *AuthorsRepoView) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *AuthorsRepoView) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *AuthorsRepoView) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetMethod

`func (o *AuthorsRepoView) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AuthorsRepoView) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AuthorsRepoView) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *AuthorsRepoView) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetBadgeMarkdown

`func (o *AuthorsRepoView) GetBadgeMarkdown() string`

GetBadgeMarkdown returns the BadgeMarkdown field if non-nil, zero value otherwise.

### GetBadgeMarkdownOk

`func (o *AuthorsRepoView) GetBadgeMarkdownOk() (*string, bool)`

GetBadgeMarkdownOk returns a tuple with the BadgeMarkdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadgeMarkdown

`func (o *AuthorsRepoView) SetBadgeMarkdown(v string)`

SetBadgeMarkdown sets BadgeMarkdown field to given value.

### HasBadgeMarkdown

`func (o *AuthorsRepoView) HasBadgeMarkdown() bool`

HasBadgeMarkdown returns a boolean if a field has been set.

### GetVerifiedAt

`func (o *AuthorsRepoView) GetVerifiedAt() int64`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *AuthorsRepoView) GetVerifiedAtOk() (*int64, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *AuthorsRepoView) SetVerifiedAt(v int64)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *AuthorsRepoView) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthorsRepoView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthorsRepoView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthorsRepoView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthorsRepoView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



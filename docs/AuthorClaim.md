# AuthorClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **bool** | Created reports whether this call recorded a new claim (201) or found an existing one (200). | [optional] 
**Org** | Pointer to [**AuthorOrgView**](AuthorOrgView.md) | Org is the verified owner-wide claim, present when an owner was claimed. It covers every repository the author publishes under that owner. | [optional] 
**Repo** | Pointer to [**AuthorAuthorRepo**](AuthorAuthorRepo.md) | Repo is the verified repository claim, present when a repository was claimed. | [optional] 

## Methods

### NewAuthorClaim

`func NewAuthorClaim() *AuthorClaim`

NewAuthorClaim instantiates a new AuthorClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorClaimWithDefaults

`func NewAuthorClaimWithDefaults() *AuthorClaim`

NewAuthorClaimWithDefaults instantiates a new AuthorClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AuthorClaim) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AuthorClaim) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AuthorClaim) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AuthorClaim) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetOrg

`func (o *AuthorClaim) GetOrg() AuthorOrgView`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AuthorClaim) GetOrgOk() (*AuthorOrgView, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AuthorClaim) SetOrg(v AuthorOrgView)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AuthorClaim) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRepo

`func (o *AuthorClaim) GetRepo() AuthorAuthorRepo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *AuthorClaim) GetRepoOk() (*AuthorAuthorRepo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *AuthorClaim) SetRepo(v AuthorAuthorRepo)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *AuthorClaim) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



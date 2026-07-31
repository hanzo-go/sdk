# CloudClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **bool** | Created reports whether this call recorded a new claim (201) or found an existing one (200). | [optional] 
**Org** | Pointer to [**CloudOrgView**](CloudOrgView.md) | Org is the verified owner-wide claim, present when an owner was claimed. It covers every repository the author publishes under that owner. | [optional] 
**Repo** | Pointer to [**CloudAuthorRepo**](CloudAuthorRepo.md) | Repo is the verified repository claim, present when a repository was claimed. | [optional] 

## Methods

### NewCloudClaim

`func NewCloudClaim() *CloudClaim`

NewCloudClaim instantiates a new CloudClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudClaimWithDefaults

`func NewCloudClaimWithDefaults() *CloudClaim`

NewCloudClaimWithDefaults instantiates a new CloudClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CloudClaim) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CloudClaim) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CloudClaim) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CloudClaim) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetOrg

`func (o *CloudClaim) GetOrg() CloudOrgView`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudClaim) GetOrgOk() (*CloudOrgView, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudClaim) SetOrg(v CloudOrgView)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudClaim) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRepo

`func (o *CloudClaim) GetRepo() CloudAuthorRepo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudClaim) GetRepoOk() (*CloudAuthorRepo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudClaim) SetRepo(v CloudAuthorRepo)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudClaim) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



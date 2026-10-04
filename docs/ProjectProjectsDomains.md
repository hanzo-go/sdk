# ProjectProjectsDomains

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claims** | Pointer to [**[]ProjectProjectsDomain**](ProjectProjectsDomain.md) | Claims is one row per host — live, or pending with the DNS records it still owes. | [optional] 
**Domains** | Pointer to **[]string** | Domains are the hostnames that are VERIFIED and routing right now. | [optional] 
**Org** | Pointer to **string** | Org is the organisation that owns the site. | [optional] 
**Slug** | Pointer to **string** | Slug is the site the panel belongs to. | [optional] 

## Methods

### NewProjectProjectsDomains

`func NewProjectProjectsDomains() *ProjectProjectsDomains`

NewProjectProjectsDomains instantiates a new ProjectProjectsDomains object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectsDomainsWithDefaults

`func NewProjectProjectsDomainsWithDefaults() *ProjectProjectsDomains`

NewProjectProjectsDomainsWithDefaults instantiates a new ProjectProjectsDomains object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaims

`func (o *ProjectProjectsDomains) GetClaims() []ProjectProjectsDomain`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *ProjectProjectsDomains) GetClaimsOk() (*[]ProjectProjectsDomain, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *ProjectProjectsDomains) SetClaims(v []ProjectProjectsDomain)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *ProjectProjectsDomains) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetDomains

`func (o *ProjectProjectsDomains) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ProjectProjectsDomains) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ProjectProjectsDomains) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *ProjectProjectsDomains) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetOrg

`func (o *ProjectProjectsDomains) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *ProjectProjectsDomains) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *ProjectProjectsDomains) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *ProjectProjectsDomains) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectProjectsDomains) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectProjectsDomains) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectProjectsDomains) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectProjectsDomains) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



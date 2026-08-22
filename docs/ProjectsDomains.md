# ProjectsDomains

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claims** | Pointer to [**[]ProjectsDomain**](ProjectsDomain.md) | Claims is one row per host — live, or pending with the DNS records it still owes. | [optional] 
**Domains** | Pointer to **[]string** | Domains are the hostnames that are VERIFIED and routing right now. | [optional] 
**Org** | Pointer to **string** | Org is the organisation that owns the site. | [optional] 
**Slug** | Pointer to **string** | Slug is the site the panel belongs to. | [optional] 

## Methods

### NewProjectsDomains

`func NewProjectsDomains() *ProjectsDomains`

NewProjectsDomains instantiates a new ProjectsDomains object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsDomainsWithDefaults

`func NewProjectsDomainsWithDefaults() *ProjectsDomains`

NewProjectsDomainsWithDefaults instantiates a new ProjectsDomains object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaims

`func (o *ProjectsDomains) GetClaims() []ProjectsDomain`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *ProjectsDomains) GetClaimsOk() (*[]ProjectsDomain, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *ProjectsDomains) SetClaims(v []ProjectsDomain)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *ProjectsDomains) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetDomains

`func (o *ProjectsDomains) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ProjectsDomains) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ProjectsDomains) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *ProjectsDomains) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetOrg

`func (o *ProjectsDomains) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *ProjectsDomains) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *ProjectsDomains) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *ProjectsDomains) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectsDomains) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsDomains) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsDomains) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectsDomains) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



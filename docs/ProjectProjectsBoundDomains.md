# ProjectProjectsBoundDomains

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bound** | Pointer to [**[]ProjectProjectsDomain**](ProjectProjectsDomain.md) | Bound is the result of THIS call, one row per host in the request: live for an already-vouched host, pending with the DNS records to publish otherwise. | [optional] 
**Domains** | Pointer to **[]string** | Domains are the hostnames that are VERIFIED and routing right now, after this bind. | [optional] 
**Org** | Pointer to **string** | Org is the organisation that owns the site. | [optional] 
**Slug** | Pointer to **string** | Slug is the site the hosts were bound to. | [optional] 

## Methods

### NewProjectProjectsBoundDomains

`func NewProjectProjectsBoundDomains() *ProjectProjectsBoundDomains`

NewProjectProjectsBoundDomains instantiates a new ProjectProjectsBoundDomains object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectsBoundDomainsWithDefaults

`func NewProjectProjectsBoundDomainsWithDefaults() *ProjectProjectsBoundDomains`

NewProjectProjectsBoundDomainsWithDefaults instantiates a new ProjectProjectsBoundDomains object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBound

`func (o *ProjectProjectsBoundDomains) GetBound() []ProjectProjectsDomain`

GetBound returns the Bound field if non-nil, zero value otherwise.

### GetBoundOk

`func (o *ProjectProjectsBoundDomains) GetBoundOk() (*[]ProjectProjectsDomain, bool)`

GetBoundOk returns a tuple with the Bound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBound

`func (o *ProjectProjectsBoundDomains) SetBound(v []ProjectProjectsDomain)`

SetBound sets Bound field to given value.

### HasBound

`func (o *ProjectProjectsBoundDomains) HasBound() bool`

HasBound returns a boolean if a field has been set.

### GetDomains

`func (o *ProjectProjectsBoundDomains) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ProjectProjectsBoundDomains) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ProjectProjectsBoundDomains) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *ProjectProjectsBoundDomains) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetOrg

`func (o *ProjectProjectsBoundDomains) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *ProjectProjectsBoundDomains) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *ProjectProjectsBoundDomains) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *ProjectProjectsBoundDomains) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectProjectsBoundDomains) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectProjectsBoundDomains) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectProjectsBoundDomains) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectProjectsBoundDomains) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



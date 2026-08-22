# ProjectsBoundDomains

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bound** | Pointer to [**[]ProjectsDomain**](ProjectsDomain.md) | Bound is the result of THIS call, one row per host in the request: live for an already-vouched host, pending with the DNS records to publish otherwise. | [optional] 
**Domains** | Pointer to **[]string** | Domains are the hostnames that are VERIFIED and routing right now, after this bind. | [optional] 
**Org** | Pointer to **string** | Org is the organisation that owns the site. | [optional] 
**Slug** | Pointer to **string** | Slug is the site the hosts were bound to. | [optional] 

## Methods

### NewProjectsBoundDomains

`func NewProjectsBoundDomains() *ProjectsBoundDomains`

NewProjectsBoundDomains instantiates a new ProjectsBoundDomains object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsBoundDomainsWithDefaults

`func NewProjectsBoundDomainsWithDefaults() *ProjectsBoundDomains`

NewProjectsBoundDomainsWithDefaults instantiates a new ProjectsBoundDomains object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBound

`func (o *ProjectsBoundDomains) GetBound() []ProjectsDomain`

GetBound returns the Bound field if non-nil, zero value otherwise.

### GetBoundOk

`func (o *ProjectsBoundDomains) GetBoundOk() (*[]ProjectsDomain, bool)`

GetBoundOk returns a tuple with the Bound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBound

`func (o *ProjectsBoundDomains) SetBound(v []ProjectsDomain)`

SetBound sets Bound field to given value.

### HasBound

`func (o *ProjectsBoundDomains) HasBound() bool`

HasBound returns a boolean if a field has been set.

### GetDomains

`func (o *ProjectsBoundDomains) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ProjectsBoundDomains) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ProjectsBoundDomains) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *ProjectsBoundDomains) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetOrg

`func (o *ProjectsBoundDomains) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *ProjectsBoundDomains) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *ProjectsBoundDomains) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *ProjectsBoundDomains) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectsBoundDomains) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsBoundDomains) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsBoundDomains) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectsBoundDomains) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



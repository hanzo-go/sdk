# ProjectsDomainsBind

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | Pointer to **[]string** | Domains are the custom hostnames to attach, in order. An empty list is a 400 rather than a clear — releasing a host is its own call. | [optional] 
**Slug** | Pointer to **string** | Slug is the site the hosts attach to, from the path. | [optional] 

## Methods

### NewProjectsDomainsBind

`func NewProjectsDomainsBind() *ProjectsDomainsBind`

NewProjectsDomainsBind instantiates a new ProjectsDomainsBind object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsDomainsBindWithDefaults

`func NewProjectsDomainsBindWithDefaults() *ProjectsDomainsBind`

NewProjectsDomainsBindWithDefaults instantiates a new ProjectsDomainsBind object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *ProjectsDomainsBind) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ProjectsDomainsBind) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ProjectsDomainsBind) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *ProjectsDomainsBind) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectsDomainsBind) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsDomainsBind) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsDomainsBind) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectsDomainsBind) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



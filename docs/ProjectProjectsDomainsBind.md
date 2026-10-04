# ProjectProjectsDomainsBind

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | Pointer to **[]string** | Domains are the custom hostnames to attach, in order. An empty list is a 400 rather than a clear — releasing a host is its own call. | [optional] 
**Slug** | Pointer to **string** | Slug is the site the hosts attach to, from the path. | [optional] 

## Methods

### NewProjectProjectsDomainsBind

`func NewProjectProjectsDomainsBind() *ProjectProjectsDomainsBind`

NewProjectProjectsDomainsBind instantiates a new ProjectProjectsDomainsBind object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectsDomainsBindWithDefaults

`func NewProjectProjectsDomainsBindWithDefaults() *ProjectProjectsDomainsBind`

NewProjectProjectsDomainsBindWithDefaults instantiates a new ProjectProjectsDomainsBind object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *ProjectProjectsDomainsBind) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ProjectProjectsDomainsBind) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ProjectProjectsDomainsBind) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *ProjectProjectsDomainsBind) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectProjectsDomainsBind) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectProjectsDomainsBind) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectProjectsDomainsBind) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectProjectsDomainsBind) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



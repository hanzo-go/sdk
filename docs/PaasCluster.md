# PaasCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ClusterAccesssToken** | Pointer to **string** |  | [optional] 
**MasterToken** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Release** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**Ips** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPaasCluster

`func NewPaasCluster() *PaasCluster`

NewPaasCluster instantiates a new PaasCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaasClusterWithDefaults

`func NewPaasClusterWithDefaults() *PaasCluster`

NewPaasClusterWithDefaults instantiates a new PaasCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaasCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaasCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaasCluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaasCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClusterAccesssToken

`func (o *PaasCluster) GetClusterAccesssToken() string`

GetClusterAccesssToken returns the ClusterAccesssToken field if non-nil, zero value otherwise.

### GetClusterAccesssTokenOk

`func (o *PaasCluster) GetClusterAccesssTokenOk() (*string, bool)`

GetClusterAccesssTokenOk returns a tuple with the ClusterAccesssToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAccesssToken

`func (o *PaasCluster) SetClusterAccesssToken(v string)`

SetClusterAccesssToken sets ClusterAccesssToken field to given value.

### HasClusterAccesssToken

`func (o *PaasCluster) HasClusterAccesssToken() bool`

HasClusterAccesssToken returns a boolean if a field has been set.

### GetMasterToken

`func (o *PaasCluster) GetMasterToken() string`

GetMasterToken returns the MasterToken field if non-nil, zero value otherwise.

### GetMasterTokenOk

`func (o *PaasCluster) GetMasterTokenOk() (*string, bool)`

GetMasterTokenOk returns a tuple with the MasterToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterToken

`func (o *PaasCluster) SetMasterToken(v string)`

SetMasterToken sets MasterToken field to given value.

### HasMasterToken

`func (o *PaasCluster) HasMasterToken() bool`

HasMasterToken returns a boolean if a field has been set.

### GetSlug

`func (o *PaasCluster) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PaasCluster) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PaasCluster) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PaasCluster) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetRelease

`func (o *PaasCluster) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *PaasCluster) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *PaasCluster) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *PaasCluster) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetDomains

`func (o *PaasCluster) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *PaasCluster) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *PaasCluster) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *PaasCluster) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetIps

`func (o *PaasCluster) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *PaasCluster) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *PaasCluster) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *PaasCluster) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaasCluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaasCluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaasCluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaasCluster) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CloudTLSConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcmeEmail** | Pointer to **string** | ACMEEmail is the ACME account email. It binds an account for the lifetime of an edge process, so it applies only when the edge (re)starts. | [optional] 
**ExtraHosts** | Pointer to **[]string** | ExtraHosts get certificates without owning a route — at most 256. They feed the ACME HostPolicy and hot-apply on the next reload. | [optional] 
**Staging** | Pointer to **bool** | Staging issues from Let&#39;s Encrypt&#39;s staging directory (untrusted certs, high rate limits). Like ACMEEmail it applies only when the edge (re)starts. | [optional] 

## Methods

### NewCloudTLSConfig

`func NewCloudTLSConfig() *CloudTLSConfig`

NewCloudTLSConfig instantiates a new CloudTLSConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTLSConfigWithDefaults

`func NewCloudTLSConfigWithDefaults() *CloudTLSConfig`

NewCloudTLSConfigWithDefaults instantiates a new CloudTLSConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcmeEmail

`func (o *CloudTLSConfig) GetAcmeEmail() string`

GetAcmeEmail returns the AcmeEmail field if non-nil, zero value otherwise.

### GetAcmeEmailOk

`func (o *CloudTLSConfig) GetAcmeEmailOk() (*string, bool)`

GetAcmeEmailOk returns a tuple with the AcmeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeEmail

`func (o *CloudTLSConfig) SetAcmeEmail(v string)`

SetAcmeEmail sets AcmeEmail field to given value.

### HasAcmeEmail

`func (o *CloudTLSConfig) HasAcmeEmail() bool`

HasAcmeEmail returns a boolean if a field has been set.

### GetExtraHosts

`func (o *CloudTLSConfig) GetExtraHosts() []string`

GetExtraHosts returns the ExtraHosts field if non-nil, zero value otherwise.

### GetExtraHostsOk

`func (o *CloudTLSConfig) GetExtraHostsOk() (*[]string, bool)`

GetExtraHostsOk returns a tuple with the ExtraHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHosts

`func (o *CloudTLSConfig) SetExtraHosts(v []string)`

SetExtraHosts sets ExtraHosts field to given value.

### HasExtraHosts

`func (o *CloudTLSConfig) HasExtraHosts() bool`

HasExtraHosts returns a boolean if a field has been set.

### GetStaging

`func (o *CloudTLSConfig) GetStaging() bool`

GetStaging returns the Staging field if non-nil, zero value otherwise.

### GetStagingOk

`func (o *CloudTLSConfig) GetStagingOk() (*bool, bool)`

GetStagingOk returns a tuple with the Staging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaging

`func (o *CloudTLSConfig) SetStaging(v bool)`

SetStaging sets Staging field to given value.

### HasStaging

`func (o *CloudTLSConfig) HasStaging() bool`

HasStaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CloudIngressStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcmeCacheDir** | Pointer to **string** | ACMECacheDir is where autocert persists accounts and certificates. | [optional] 
**AcmeStaging** | Pointer to **bool** | ACMEStaging is true when certificates are issued from Let&#39;s Encrypt staging. | [optional] 
**EdgeEnabled** | Pointer to **bool** | EdgeEnabled is true when the edge listeners are actually bound. | [optional] 
**HttpAddr** | Pointer to **string** | HTTPAddr is the address the ACME HTTP-01 + HTTP router listens on. | [optional] 
**HttpsAddr** | Pointer to **string** | HTTPSAddr is the address the SNI TLS terminator listens on. | [optional] 
**LiveHosts** | Pointer to **int32** | LiveHosts is how many hosts the compiled table routes. | [optional] 
**Proxy** | Pointer to **string** | Proxy names the reverse-proxy implementation behind every route. | [optional] 
**Role** | Pointer to **string** | Role is \&quot;edge\&quot; when CLOUD_INGRESS_EDGE_ENABLED is set, else \&quot;app\&quot;. | [optional] 
**TlsHosts** | Pointer to **int32** | TLSHosts is how many hosts the ACME HostPolicy will issue a certificate for. NOT a subset of LiveHosts: an extraHost owns no route, and a TLS route naming a missing service is skipped while its host still wants a cert. | [optional] 

## Methods

### NewCloudIngressStatus

`func NewCloudIngressStatus() *CloudIngressStatus`

NewCloudIngressStatus instantiates a new CloudIngressStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudIngressStatusWithDefaults

`func NewCloudIngressStatusWithDefaults() *CloudIngressStatus`

NewCloudIngressStatusWithDefaults instantiates a new CloudIngressStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcmeCacheDir

`func (o *CloudIngressStatus) GetAcmeCacheDir() string`

GetAcmeCacheDir returns the AcmeCacheDir field if non-nil, zero value otherwise.

### GetAcmeCacheDirOk

`func (o *CloudIngressStatus) GetAcmeCacheDirOk() (*string, bool)`

GetAcmeCacheDirOk returns a tuple with the AcmeCacheDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeCacheDir

`func (o *CloudIngressStatus) SetAcmeCacheDir(v string)`

SetAcmeCacheDir sets AcmeCacheDir field to given value.

### HasAcmeCacheDir

`func (o *CloudIngressStatus) HasAcmeCacheDir() bool`

HasAcmeCacheDir returns a boolean if a field has been set.

### GetAcmeStaging

`func (o *CloudIngressStatus) GetAcmeStaging() bool`

GetAcmeStaging returns the AcmeStaging field if non-nil, zero value otherwise.

### GetAcmeStagingOk

`func (o *CloudIngressStatus) GetAcmeStagingOk() (*bool, bool)`

GetAcmeStagingOk returns a tuple with the AcmeStaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeStaging

`func (o *CloudIngressStatus) SetAcmeStaging(v bool)`

SetAcmeStaging sets AcmeStaging field to given value.

### HasAcmeStaging

`func (o *CloudIngressStatus) HasAcmeStaging() bool`

HasAcmeStaging returns a boolean if a field has been set.

### GetEdgeEnabled

`func (o *CloudIngressStatus) GetEdgeEnabled() bool`

GetEdgeEnabled returns the EdgeEnabled field if non-nil, zero value otherwise.

### GetEdgeEnabledOk

`func (o *CloudIngressStatus) GetEdgeEnabledOk() (*bool, bool)`

GetEdgeEnabledOk returns a tuple with the EdgeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeEnabled

`func (o *CloudIngressStatus) SetEdgeEnabled(v bool)`

SetEdgeEnabled sets EdgeEnabled field to given value.

### HasEdgeEnabled

`func (o *CloudIngressStatus) HasEdgeEnabled() bool`

HasEdgeEnabled returns a boolean if a field has been set.

### GetHttpAddr

`func (o *CloudIngressStatus) GetHttpAddr() string`

GetHttpAddr returns the HttpAddr field if non-nil, zero value otherwise.

### GetHttpAddrOk

`func (o *CloudIngressStatus) GetHttpAddrOk() (*string, bool)`

GetHttpAddrOk returns a tuple with the HttpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpAddr

`func (o *CloudIngressStatus) SetHttpAddr(v string)`

SetHttpAddr sets HttpAddr field to given value.

### HasHttpAddr

`func (o *CloudIngressStatus) HasHttpAddr() bool`

HasHttpAddr returns a boolean if a field has been set.

### GetHttpsAddr

`func (o *CloudIngressStatus) GetHttpsAddr() string`

GetHttpsAddr returns the HttpsAddr field if non-nil, zero value otherwise.

### GetHttpsAddrOk

`func (o *CloudIngressStatus) GetHttpsAddrOk() (*string, bool)`

GetHttpsAddrOk returns a tuple with the HttpsAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsAddr

`func (o *CloudIngressStatus) SetHttpsAddr(v string)`

SetHttpsAddr sets HttpsAddr field to given value.

### HasHttpsAddr

`func (o *CloudIngressStatus) HasHttpsAddr() bool`

HasHttpsAddr returns a boolean if a field has been set.

### GetLiveHosts

`func (o *CloudIngressStatus) GetLiveHosts() int32`

GetLiveHosts returns the LiveHosts field if non-nil, zero value otherwise.

### GetLiveHostsOk

`func (o *CloudIngressStatus) GetLiveHostsOk() (*int32, bool)`

GetLiveHostsOk returns a tuple with the LiveHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveHosts

`func (o *CloudIngressStatus) SetLiveHosts(v int32)`

SetLiveHosts sets LiveHosts field to given value.

### HasLiveHosts

`func (o *CloudIngressStatus) HasLiveHosts() bool`

HasLiveHosts returns a boolean if a field has been set.

### GetProxy

`func (o *CloudIngressStatus) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *CloudIngressStatus) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *CloudIngressStatus) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *CloudIngressStatus) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetRole

`func (o *CloudIngressStatus) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CloudIngressStatus) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CloudIngressStatus) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CloudIngressStatus) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTlsHosts

`func (o *CloudIngressStatus) GetTlsHosts() int32`

GetTlsHosts returns the TlsHosts field if non-nil, zero value otherwise.

### GetTlsHostsOk

`func (o *CloudIngressStatus) GetTlsHostsOk() (*int32, bool)`

GetTlsHostsOk returns a tuple with the TlsHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsHosts

`func (o *CloudIngressStatus) SetTlsHosts(v int32)`

SetTlsHosts sets TlsHosts field to given value.

### HasTlsHosts

`func (o *CloudIngressStatus) HasTlsHosts() bool`

HasTlsHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



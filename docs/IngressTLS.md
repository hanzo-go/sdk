# IngressTLS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcmeDirectory** | Pointer to **string** | ACMEDirectory is the ACME endpoint in use: the staging URL, or \&quot;letsencrypt-production\&quot;. | [optional] 
**AcmeEmail** | Pointer to **string** | ACMEEmail is the account email the PROCESS was started with (CLOUD_INGRESS_ACME_EMAIL), not the stored config&#39;s. | [optional] 
**Config** | Pointer to [**TLSConfig**](TLSConfig.md) | Config is the caller org&#39;s stored ACME intent. | [optional] 
**EdgeEnabled** | Pointer to **bool** | EdgeEnabled is true when the edge listeners are actually bound. | [optional] 
**ManagedHosts** | Pointer to **[]string** | ManagedHosts is every host the ACME HostPolicy will issue a certificate for — the union across ALL orgs of TLS-marked routes and configured extraHosts, because one process holds one certificate cache. | [optional] 
**Note** | Pointer to **string** | Note states which fields hot-apply and which need an edge restart. | [optional] 
**Role** | Pointer to **string** | Role is \&quot;edge\&quot; when this instance binds the listeners, else \&quot;app\&quot;. | [optional] 

## Methods

### NewIngressTLS

`func NewIngressTLS() *IngressTLS`

NewIngressTLS instantiates a new IngressTLS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressTLSWithDefaults

`func NewIngressTLSWithDefaults() *IngressTLS`

NewIngressTLSWithDefaults instantiates a new IngressTLS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcmeDirectory

`func (o *IngressTLS) GetAcmeDirectory() string`

GetAcmeDirectory returns the AcmeDirectory field if non-nil, zero value otherwise.

### GetAcmeDirectoryOk

`func (o *IngressTLS) GetAcmeDirectoryOk() (*string, bool)`

GetAcmeDirectoryOk returns a tuple with the AcmeDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeDirectory

`func (o *IngressTLS) SetAcmeDirectory(v string)`

SetAcmeDirectory sets AcmeDirectory field to given value.

### HasAcmeDirectory

`func (o *IngressTLS) HasAcmeDirectory() bool`

HasAcmeDirectory returns a boolean if a field has been set.

### GetAcmeEmail

`func (o *IngressTLS) GetAcmeEmail() string`

GetAcmeEmail returns the AcmeEmail field if non-nil, zero value otherwise.

### GetAcmeEmailOk

`func (o *IngressTLS) GetAcmeEmailOk() (*string, bool)`

GetAcmeEmailOk returns a tuple with the AcmeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeEmail

`func (o *IngressTLS) SetAcmeEmail(v string)`

SetAcmeEmail sets AcmeEmail field to given value.

### HasAcmeEmail

`func (o *IngressTLS) HasAcmeEmail() bool`

HasAcmeEmail returns a boolean if a field has been set.

### GetConfig

`func (o *IngressTLS) GetConfig() TLSConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IngressTLS) GetConfigOk() (*TLSConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IngressTLS) SetConfig(v TLSConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *IngressTLS) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEdgeEnabled

`func (o *IngressTLS) GetEdgeEnabled() bool`

GetEdgeEnabled returns the EdgeEnabled field if non-nil, zero value otherwise.

### GetEdgeEnabledOk

`func (o *IngressTLS) GetEdgeEnabledOk() (*bool, bool)`

GetEdgeEnabledOk returns a tuple with the EdgeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeEnabled

`func (o *IngressTLS) SetEdgeEnabled(v bool)`

SetEdgeEnabled sets EdgeEnabled field to given value.

### HasEdgeEnabled

`func (o *IngressTLS) HasEdgeEnabled() bool`

HasEdgeEnabled returns a boolean if a field has been set.

### GetManagedHosts

`func (o *IngressTLS) GetManagedHosts() []string`

GetManagedHosts returns the ManagedHosts field if non-nil, zero value otherwise.

### GetManagedHostsOk

`func (o *IngressTLS) GetManagedHostsOk() (*[]string, bool)`

GetManagedHostsOk returns a tuple with the ManagedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedHosts

`func (o *IngressTLS) SetManagedHosts(v []string)`

SetManagedHosts sets ManagedHosts field to given value.

### HasManagedHosts

`func (o *IngressTLS) HasManagedHosts() bool`

HasManagedHosts returns a boolean if a field has been set.

### GetNote

`func (o *IngressTLS) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *IngressTLS) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *IngressTLS) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *IngressTLS) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetRole

`func (o *IngressTLS) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IngressTLS) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IngressTLS) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *IngressTLS) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



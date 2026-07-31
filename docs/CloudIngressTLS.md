# CloudIngressTLS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcmeDirectory** | Pointer to **string** | ACMEDirectory is the ACME endpoint in use: the staging URL, or \&quot;letsencrypt-production\&quot;. | [optional] 
**AcmeEmail** | Pointer to **string** | ACMEEmail is the account email the PROCESS was started with (CLOUD_INGRESS_ACME_EMAIL), not the stored config&#39;s. | [optional] 
**Config** | Pointer to [**CloudTLSConfig**](CloudTLSConfig.md) | Config is the caller org&#39;s stored ACME intent. | [optional] 
**EdgeEnabled** | Pointer to **bool** | EdgeEnabled is true when the edge listeners are actually bound. | [optional] 
**ManagedHosts** | Pointer to **[]string** | ManagedHosts is every host the ACME HostPolicy will issue a certificate for — the union across ALL orgs of TLS-marked routes and configured extraHosts, because one process holds one certificate cache. | [optional] 
**Note** | Pointer to **string** | Note states which fields hot-apply and which need an edge restart. | [optional] 
**Role** | Pointer to **string** | Role is \&quot;edge\&quot; when this instance binds the listeners, else \&quot;app\&quot;. | [optional] 

## Methods

### NewCloudIngressTLS

`func NewCloudIngressTLS() *CloudIngressTLS`

NewCloudIngressTLS instantiates a new CloudIngressTLS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudIngressTLSWithDefaults

`func NewCloudIngressTLSWithDefaults() *CloudIngressTLS`

NewCloudIngressTLSWithDefaults instantiates a new CloudIngressTLS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcmeDirectory

`func (o *CloudIngressTLS) GetAcmeDirectory() string`

GetAcmeDirectory returns the AcmeDirectory field if non-nil, zero value otherwise.

### GetAcmeDirectoryOk

`func (o *CloudIngressTLS) GetAcmeDirectoryOk() (*string, bool)`

GetAcmeDirectoryOk returns a tuple with the AcmeDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeDirectory

`func (o *CloudIngressTLS) SetAcmeDirectory(v string)`

SetAcmeDirectory sets AcmeDirectory field to given value.

### HasAcmeDirectory

`func (o *CloudIngressTLS) HasAcmeDirectory() bool`

HasAcmeDirectory returns a boolean if a field has been set.

### GetAcmeEmail

`func (o *CloudIngressTLS) GetAcmeEmail() string`

GetAcmeEmail returns the AcmeEmail field if non-nil, zero value otherwise.

### GetAcmeEmailOk

`func (o *CloudIngressTLS) GetAcmeEmailOk() (*string, bool)`

GetAcmeEmailOk returns a tuple with the AcmeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeEmail

`func (o *CloudIngressTLS) SetAcmeEmail(v string)`

SetAcmeEmail sets AcmeEmail field to given value.

### HasAcmeEmail

`func (o *CloudIngressTLS) HasAcmeEmail() bool`

HasAcmeEmail returns a boolean if a field has been set.

### GetConfig

`func (o *CloudIngressTLS) GetConfig() CloudTLSConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CloudIngressTLS) GetConfigOk() (*CloudTLSConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CloudIngressTLS) SetConfig(v CloudTLSConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CloudIngressTLS) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEdgeEnabled

`func (o *CloudIngressTLS) GetEdgeEnabled() bool`

GetEdgeEnabled returns the EdgeEnabled field if non-nil, zero value otherwise.

### GetEdgeEnabledOk

`func (o *CloudIngressTLS) GetEdgeEnabledOk() (*bool, bool)`

GetEdgeEnabledOk returns a tuple with the EdgeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeEnabled

`func (o *CloudIngressTLS) SetEdgeEnabled(v bool)`

SetEdgeEnabled sets EdgeEnabled field to given value.

### HasEdgeEnabled

`func (o *CloudIngressTLS) HasEdgeEnabled() bool`

HasEdgeEnabled returns a boolean if a field has been set.

### GetManagedHosts

`func (o *CloudIngressTLS) GetManagedHosts() []string`

GetManagedHosts returns the ManagedHosts field if non-nil, zero value otherwise.

### GetManagedHostsOk

`func (o *CloudIngressTLS) GetManagedHostsOk() (*[]string, bool)`

GetManagedHostsOk returns a tuple with the ManagedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedHosts

`func (o *CloudIngressTLS) SetManagedHosts(v []string)`

SetManagedHosts sets ManagedHosts field to given value.

### HasManagedHosts

`func (o *CloudIngressTLS) HasManagedHosts() bool`

HasManagedHosts returns a boolean if a field has been set.

### GetNote

`func (o *CloudIngressTLS) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CloudIngressTLS) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CloudIngressTLS) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CloudIngressTLS) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetRole

`func (o *CloudIngressTLS) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CloudIngressTLS) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CloudIngressTLS) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CloudIngressTLS) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



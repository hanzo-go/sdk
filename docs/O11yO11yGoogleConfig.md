# O11yO11yGoogleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedGroups** | Pointer to **[]string** | AllowedGroups, when set, admits only members of these groups. | [optional] 
**ClientId** | Pointer to **string** | ClientID is the OAuth application&#39;s id. | [optional] 
**ClientSecret** | Pointer to **string** | ClientSecret is the OAuth application&#39;s secret. | [optional] 
**DomainToAdminEmail** | Pointer to **map[string]string** | DomainToAdminEmail maps each Workspace domain to the admin the service account impersonates; \&quot;*\&quot; is the fallback. | [optional] 
**FetchGroups** | Pointer to **bool** | FetchGroups reads the user&#39;s Workspace groups for role mapping. | [optional] 
**FetchTransitiveGroupMembership** | Pointer to **bool** | FetchTransitiveGroupMembership also reads groups held through other groups. | [optional] 
**InsecureSkipEmailVerified** | Pointer to **bool** | InsecureSkipEmailVerified admits addresses Google has not verified. | [optional] 
**RedirectURI** | Pointer to **string** | RedirectURI is the callback the flow returns to. | [optional] 
**ServiceAccountJson** | Pointer to **string** | ServiceAccountJSON is the service-account credential used to read groups, when FetchGroups is on. | [optional] 

## Methods

### NewO11yO11yGoogleConfig

`func NewO11yO11yGoogleConfig() *O11yO11yGoogleConfig`

NewO11yO11yGoogleConfig instantiates a new O11yO11yGoogleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yGoogleConfigWithDefaults

`func NewO11yO11yGoogleConfigWithDefaults() *O11yO11yGoogleConfig`

NewO11yO11yGoogleConfigWithDefaults instantiates a new O11yO11yGoogleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedGroups

`func (o *O11yO11yGoogleConfig) GetAllowedGroups() []string`

GetAllowedGroups returns the AllowedGroups field if non-nil, zero value otherwise.

### GetAllowedGroupsOk

`func (o *O11yO11yGoogleConfig) GetAllowedGroupsOk() (*[]string, bool)`

GetAllowedGroupsOk returns a tuple with the AllowedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedGroups

`func (o *O11yO11yGoogleConfig) SetAllowedGroups(v []string)`

SetAllowedGroups sets AllowedGroups field to given value.

### HasAllowedGroups

`func (o *O11yO11yGoogleConfig) HasAllowedGroups() bool`

HasAllowedGroups returns a boolean if a field has been set.

### GetClientId

`func (o *O11yO11yGoogleConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *O11yO11yGoogleConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *O11yO11yGoogleConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *O11yO11yGoogleConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *O11yO11yGoogleConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *O11yO11yGoogleConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *O11yO11yGoogleConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *O11yO11yGoogleConfig) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetDomainToAdminEmail

`func (o *O11yO11yGoogleConfig) GetDomainToAdminEmail() map[string]string`

GetDomainToAdminEmail returns the DomainToAdminEmail field if non-nil, zero value otherwise.

### GetDomainToAdminEmailOk

`func (o *O11yO11yGoogleConfig) GetDomainToAdminEmailOk() (*map[string]string, bool)`

GetDomainToAdminEmailOk returns a tuple with the DomainToAdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainToAdminEmail

`func (o *O11yO11yGoogleConfig) SetDomainToAdminEmail(v map[string]string)`

SetDomainToAdminEmail sets DomainToAdminEmail field to given value.

### HasDomainToAdminEmail

`func (o *O11yO11yGoogleConfig) HasDomainToAdminEmail() bool`

HasDomainToAdminEmail returns a boolean if a field has been set.

### GetFetchGroups

`func (o *O11yO11yGoogleConfig) GetFetchGroups() bool`

GetFetchGroups returns the FetchGroups field if non-nil, zero value otherwise.

### GetFetchGroupsOk

`func (o *O11yO11yGoogleConfig) GetFetchGroupsOk() (*bool, bool)`

GetFetchGroupsOk returns a tuple with the FetchGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchGroups

`func (o *O11yO11yGoogleConfig) SetFetchGroups(v bool)`

SetFetchGroups sets FetchGroups field to given value.

### HasFetchGroups

`func (o *O11yO11yGoogleConfig) HasFetchGroups() bool`

HasFetchGroups returns a boolean if a field has been set.

### GetFetchTransitiveGroupMembership

`func (o *O11yO11yGoogleConfig) GetFetchTransitiveGroupMembership() bool`

GetFetchTransitiveGroupMembership returns the FetchTransitiveGroupMembership field if non-nil, zero value otherwise.

### GetFetchTransitiveGroupMembershipOk

`func (o *O11yO11yGoogleConfig) GetFetchTransitiveGroupMembershipOk() (*bool, bool)`

GetFetchTransitiveGroupMembershipOk returns a tuple with the FetchTransitiveGroupMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchTransitiveGroupMembership

`func (o *O11yO11yGoogleConfig) SetFetchTransitiveGroupMembership(v bool)`

SetFetchTransitiveGroupMembership sets FetchTransitiveGroupMembership field to given value.

### HasFetchTransitiveGroupMembership

`func (o *O11yO11yGoogleConfig) HasFetchTransitiveGroupMembership() bool`

HasFetchTransitiveGroupMembership returns a boolean if a field has been set.

### GetInsecureSkipEmailVerified

`func (o *O11yO11yGoogleConfig) GetInsecureSkipEmailVerified() bool`

GetInsecureSkipEmailVerified returns the InsecureSkipEmailVerified field if non-nil, zero value otherwise.

### GetInsecureSkipEmailVerifiedOk

`func (o *O11yO11yGoogleConfig) GetInsecureSkipEmailVerifiedOk() (*bool, bool)`

GetInsecureSkipEmailVerifiedOk returns a tuple with the InsecureSkipEmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSkipEmailVerified

`func (o *O11yO11yGoogleConfig) SetInsecureSkipEmailVerified(v bool)`

SetInsecureSkipEmailVerified sets InsecureSkipEmailVerified field to given value.

### HasInsecureSkipEmailVerified

`func (o *O11yO11yGoogleConfig) HasInsecureSkipEmailVerified() bool`

HasInsecureSkipEmailVerified returns a boolean if a field has been set.

### GetRedirectURI

`func (o *O11yO11yGoogleConfig) GetRedirectURI() string`

GetRedirectURI returns the RedirectURI field if non-nil, zero value otherwise.

### GetRedirectURIOk

`func (o *O11yO11yGoogleConfig) GetRedirectURIOk() (*string, bool)`

GetRedirectURIOk returns a tuple with the RedirectURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectURI

`func (o *O11yO11yGoogleConfig) SetRedirectURI(v string)`

SetRedirectURI sets RedirectURI field to given value.

### HasRedirectURI

`func (o *O11yO11yGoogleConfig) HasRedirectURI() bool`

HasRedirectURI returns a boolean if a field has been set.

### GetServiceAccountJson

`func (o *O11yO11yGoogleConfig) GetServiceAccountJson() string`

GetServiceAccountJson returns the ServiceAccountJson field if non-nil, zero value otherwise.

### GetServiceAccountJsonOk

`func (o *O11yO11yGoogleConfig) GetServiceAccountJsonOk() (*string, bool)`

GetServiceAccountJsonOk returns a tuple with the ServiceAccountJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountJson

`func (o *O11yO11yGoogleConfig) SetServiceAccountJson(v string)`

SetServiceAccountJson sets ServiceAccountJson field to given value.

### HasServiceAccountJson

`func (o *O11yO11yGoogleConfig) HasServiceAccountJson() bool`

HasServiceAccountJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# VenueLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | ClientID is the Azure AD application id. Azure only. | [optional] 
**ClientSecret** | Pointer to **string** | ClientSecret selects the service-principal flow. LEAVING IT OUT selects keyless workload identity federation instead, so omitting it is a choice rather than an omission. Azure only. | [optional] 
**CredentialJson** | Pointer to **string** | CredentialJSON is a Google credentials document — an external_account (workload identity federation, keyless) or a service-account key. GCP only. | [optional] 
**ExternalId** | Pointer to **string** | ExternalID pins that role assumption to Hanzo, which is what closes the confused-deputy hole. AWS only. | [optional] 
**Label** | Pointer to **string** | Label is the org-chosen name for this account within the provider, which is how a second account at the same provider is addressed later. Empty means \&quot;default\&quot;; anything outside 1–64 of [A-Za-z0-9._-] is refused. | [optional] 
**ProjectIds** | Pointer to **[]string** | ProjectIDs bounds the GKE cluster sweep. GCP only. | [optional] 
**Provider** | Pointer to **string** | Provider is the cloud being linked, from the path: digitalocean, aws, gcp or azure. | [optional] 
**Regions** | Pointer to **[]string** | Regions bounds the AWS EKS cluster sweep. AWS only. | [optional] 
**RoleArn** | Pointer to **string** | RoleARN is the AWS role Hanzo assumes into the account — the keyless path, so no access key is ever stored. | [optional] 
**SubscriptionIds** | Pointer to **[]string** | SubscriptionIDs bounds the AKS cluster sweep. Azure only. | [optional] 
**TenantId** | Pointer to **string** | TenantID is the Azure AD tenant of the app. Azure only. | [optional] 
**Token** | Pointer to **string** | Token is the DigitalOcean personal access token. DigitalOcean only, and it is the one provider that requires storing a secret. | [optional] 

## Methods

### NewVenueLinkRequest

`func NewVenueLinkRequest() *VenueLinkRequest`

NewVenueLinkRequest instantiates a new VenueLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVenueLinkRequestWithDefaults

`func NewVenueLinkRequestWithDefaults() *VenueLinkRequest`

NewVenueLinkRequestWithDefaults instantiates a new VenueLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *VenueLinkRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *VenueLinkRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *VenueLinkRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *VenueLinkRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *VenueLinkRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *VenueLinkRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *VenueLinkRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *VenueLinkRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetCredentialJson

`func (o *VenueLinkRequest) GetCredentialJson() string`

GetCredentialJson returns the CredentialJson field if non-nil, zero value otherwise.

### GetCredentialJsonOk

`func (o *VenueLinkRequest) GetCredentialJsonOk() (*string, bool)`

GetCredentialJsonOk returns a tuple with the CredentialJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialJson

`func (o *VenueLinkRequest) SetCredentialJson(v string)`

SetCredentialJson sets CredentialJson field to given value.

### HasCredentialJson

`func (o *VenueLinkRequest) HasCredentialJson() bool`

HasCredentialJson returns a boolean if a field has been set.

### GetExternalId

`func (o *VenueLinkRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *VenueLinkRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *VenueLinkRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *VenueLinkRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLabel

`func (o *VenueLinkRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VenueLinkRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VenueLinkRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *VenueLinkRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetProjectIds

`func (o *VenueLinkRequest) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *VenueLinkRequest) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *VenueLinkRequest) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *VenueLinkRequest) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### GetProvider

`func (o *VenueLinkRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VenueLinkRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VenueLinkRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *VenueLinkRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegions

`func (o *VenueLinkRequest) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *VenueLinkRequest) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *VenueLinkRequest) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *VenueLinkRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetRoleArn

`func (o *VenueLinkRequest) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *VenueLinkRequest) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *VenueLinkRequest) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *VenueLinkRequest) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetSubscriptionIds

`func (o *VenueLinkRequest) GetSubscriptionIds() []string`

GetSubscriptionIds returns the SubscriptionIds field if non-nil, zero value otherwise.

### GetSubscriptionIdsOk

`func (o *VenueLinkRequest) GetSubscriptionIdsOk() (*[]string, bool)`

GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIds

`func (o *VenueLinkRequest) SetSubscriptionIds(v []string)`

SetSubscriptionIds sets SubscriptionIds field to given value.

### HasSubscriptionIds

`func (o *VenueLinkRequest) HasSubscriptionIds() bool`

HasSubscriptionIds returns a boolean if a field has been set.

### GetTenantId

`func (o *VenueLinkRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *VenueLinkRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *VenueLinkRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *VenueLinkRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetToken

`func (o *VenueLinkRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VenueLinkRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VenueLinkRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *VenueLinkRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



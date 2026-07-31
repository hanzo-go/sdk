# CloudVenueLinkRequest

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

### NewCloudVenueLinkRequest

`func NewCloudVenueLinkRequest() *CloudVenueLinkRequest`

NewCloudVenueLinkRequest instantiates a new CloudVenueLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVenueLinkRequestWithDefaults

`func NewCloudVenueLinkRequestWithDefaults() *CloudVenueLinkRequest`

NewCloudVenueLinkRequestWithDefaults instantiates a new CloudVenueLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *CloudVenueLinkRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CloudVenueLinkRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CloudVenueLinkRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CloudVenueLinkRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *CloudVenueLinkRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CloudVenueLinkRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CloudVenueLinkRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *CloudVenueLinkRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetCredentialJson

`func (o *CloudVenueLinkRequest) GetCredentialJson() string`

GetCredentialJson returns the CredentialJson field if non-nil, zero value otherwise.

### GetCredentialJsonOk

`func (o *CloudVenueLinkRequest) GetCredentialJsonOk() (*string, bool)`

GetCredentialJsonOk returns a tuple with the CredentialJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialJson

`func (o *CloudVenueLinkRequest) SetCredentialJson(v string)`

SetCredentialJson sets CredentialJson field to given value.

### HasCredentialJson

`func (o *CloudVenueLinkRequest) HasCredentialJson() bool`

HasCredentialJson returns a boolean if a field has been set.

### GetExternalId

`func (o *CloudVenueLinkRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CloudVenueLinkRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CloudVenueLinkRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CloudVenueLinkRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLabel

`func (o *CloudVenueLinkRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CloudVenueLinkRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CloudVenueLinkRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CloudVenueLinkRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetProjectIds

`func (o *CloudVenueLinkRequest) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *CloudVenueLinkRequest) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *CloudVenueLinkRequest) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *CloudVenueLinkRequest) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### GetProvider

`func (o *CloudVenueLinkRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudVenueLinkRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudVenueLinkRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudVenueLinkRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegions

`func (o *CloudVenueLinkRequest) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *CloudVenueLinkRequest) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *CloudVenueLinkRequest) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *CloudVenueLinkRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetRoleArn

`func (o *CloudVenueLinkRequest) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *CloudVenueLinkRequest) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *CloudVenueLinkRequest) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *CloudVenueLinkRequest) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetSubscriptionIds

`func (o *CloudVenueLinkRequest) GetSubscriptionIds() []string`

GetSubscriptionIds returns the SubscriptionIds field if non-nil, zero value otherwise.

### GetSubscriptionIdsOk

`func (o *CloudVenueLinkRequest) GetSubscriptionIdsOk() (*[]string, bool)`

GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIds

`func (o *CloudVenueLinkRequest) SetSubscriptionIds(v []string)`

SetSubscriptionIds sets SubscriptionIds field to given value.

### HasSubscriptionIds

`func (o *CloudVenueLinkRequest) HasSubscriptionIds() bool`

HasSubscriptionIds returns a boolean if a field has been set.

### GetTenantId

`func (o *CloudVenueLinkRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CloudVenueLinkRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CloudVenueLinkRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CloudVenueLinkRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetToken

`func (o *CloudVenueLinkRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CloudVenueLinkRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CloudVenueLinkRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CloudVenueLinkRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



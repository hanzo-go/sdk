# IamProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** |  | [optional] 
**Bucket** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Cert** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientId2** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**ClientSecret2** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**CustomAuthUrl** | Pointer to **string** |  | [optional] 
**CustomLogo** | Pointer to **string** |  | [optional] 
**CustomTokenUrl** | Pointer to **string** |  | [optional] 
**CustomUserInfoUrl** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DisableSsl** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**EmailRegex** | Pointer to **string** |  | [optional] 
**EnablePkce** | Pointer to **bool** |  | [optional] 
**EnableProxy** | Pointer to **bool** |  | [optional] 
**EnableSignAuthnRequest** | Pointer to **bool** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**HttpHeaders** | Pointer to **map[string]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IdP** | Pointer to **string** |  | [optional] 
**IntranetEndpoint** | Pointer to **string** |  | [optional] 
**IssuerUrl** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**PathPrefix** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**ProviderUrl** | Pointer to **string** |  | [optional] 
**Receiver** | Pointer to **string** |  | [optional] 
**RegionId** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **string** |  | [optional] 
**SignName** | Pointer to **string** |  | [optional] 
**SslMode** | Pointer to **string** |  | [optional] 
**SubType** | Pointer to **string** |  | [optional] 
**TemplateCode** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UserMapping** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewIamProvider

`func NewIamProvider() *IamProvider`

NewIamProvider instantiates a new IamProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamProviderWithDefaults

`func NewIamProviderWithDefaults() *IamProvider`

NewIamProviderWithDefaults instantiates a new IamProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *IamProvider) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *IamProvider) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *IamProvider) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *IamProvider) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetBucket

`func (o *IamProvider) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *IamProvider) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *IamProvider) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *IamProvider) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetCategory

`func (o *IamProvider) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IamProvider) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IamProvider) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IamProvider) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCert

`func (o *IamProvider) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *IamProvider) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *IamProvider) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *IamProvider) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetClientId

`func (o *IamProvider) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IamProvider) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IamProvider) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IamProvider) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientId2

`func (o *IamProvider) GetClientId2() string`

GetClientId2 returns the ClientId2 field if non-nil, zero value otherwise.

### GetClientId2Ok

`func (o *IamProvider) GetClientId2Ok() (*string, bool)`

GetClientId2Ok returns a tuple with the ClientId2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId2

`func (o *IamProvider) SetClientId2(v string)`

SetClientId2 sets ClientId2 field to given value.

### HasClientId2

`func (o *IamProvider) HasClientId2() bool`

HasClientId2 returns a boolean if a field has been set.

### GetClientSecret

`func (o *IamProvider) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IamProvider) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IamProvider) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *IamProvider) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientSecret2

`func (o *IamProvider) GetClientSecret2() string`

GetClientSecret2 returns the ClientSecret2 field if non-nil, zero value otherwise.

### GetClientSecret2Ok

`func (o *IamProvider) GetClientSecret2Ok() (*string, bool)`

GetClientSecret2Ok returns a tuple with the ClientSecret2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret2

`func (o *IamProvider) SetClientSecret2(v string)`

SetClientSecret2 sets ClientSecret2 field to given value.

### HasClientSecret2

`func (o *IamProvider) HasClientSecret2() bool`

HasClientSecret2 returns a boolean if a field has been set.

### GetContent

`func (o *IamProvider) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *IamProvider) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *IamProvider) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *IamProvider) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IamProvider) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamProvider) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamProvider) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamProvider) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamProvider) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamProvider) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamProvider) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamProvider) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCustomAuthUrl

`func (o *IamProvider) GetCustomAuthUrl() string`

GetCustomAuthUrl returns the CustomAuthUrl field if non-nil, zero value otherwise.

### GetCustomAuthUrlOk

`func (o *IamProvider) GetCustomAuthUrlOk() (*string, bool)`

GetCustomAuthUrlOk returns a tuple with the CustomAuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAuthUrl

`func (o *IamProvider) SetCustomAuthUrl(v string)`

SetCustomAuthUrl sets CustomAuthUrl field to given value.

### HasCustomAuthUrl

`func (o *IamProvider) HasCustomAuthUrl() bool`

HasCustomAuthUrl returns a boolean if a field has been set.

### GetCustomLogo

`func (o *IamProvider) GetCustomLogo() string`

GetCustomLogo returns the CustomLogo field if non-nil, zero value otherwise.

### GetCustomLogoOk

`func (o *IamProvider) GetCustomLogoOk() (*string, bool)`

GetCustomLogoOk returns a tuple with the CustomLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLogo

`func (o *IamProvider) SetCustomLogo(v string)`

SetCustomLogo sets CustomLogo field to given value.

### HasCustomLogo

`func (o *IamProvider) HasCustomLogo() bool`

HasCustomLogo returns a boolean if a field has been set.

### GetCustomTokenUrl

`func (o *IamProvider) GetCustomTokenUrl() string`

GetCustomTokenUrl returns the CustomTokenUrl field if non-nil, zero value otherwise.

### GetCustomTokenUrlOk

`func (o *IamProvider) GetCustomTokenUrlOk() (*string, bool)`

GetCustomTokenUrlOk returns a tuple with the CustomTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTokenUrl

`func (o *IamProvider) SetCustomTokenUrl(v string)`

SetCustomTokenUrl sets CustomTokenUrl field to given value.

### HasCustomTokenUrl

`func (o *IamProvider) HasCustomTokenUrl() bool`

HasCustomTokenUrl returns a boolean if a field has been set.

### GetCustomUserInfoUrl

`func (o *IamProvider) GetCustomUserInfoUrl() string`

GetCustomUserInfoUrl returns the CustomUserInfoUrl field if non-nil, zero value otherwise.

### GetCustomUserInfoUrlOk

`func (o *IamProvider) GetCustomUserInfoUrlOk() (*string, bool)`

GetCustomUserInfoUrlOk returns a tuple with the CustomUserInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserInfoUrl

`func (o *IamProvider) SetCustomUserInfoUrl(v string)`

SetCustomUserInfoUrl sets CustomUserInfoUrl field to given value.

### HasCustomUserInfoUrl

`func (o *IamProvider) HasCustomUserInfoUrl() bool`

HasCustomUserInfoUrl returns a boolean if a field has been set.

### GetDeleted

`func (o *IamProvider) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamProvider) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamProvider) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamProvider) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDisableSsl

`func (o *IamProvider) GetDisableSsl() bool`

GetDisableSsl returns the DisableSsl field if non-nil, zero value otherwise.

### GetDisableSslOk

`func (o *IamProvider) GetDisableSslOk() (*bool, bool)`

GetDisableSslOk returns a tuple with the DisableSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSsl

`func (o *IamProvider) SetDisableSsl(v bool)`

SetDisableSsl sets DisableSsl field to given value.

### HasDisableSsl

`func (o *IamProvider) HasDisableSsl() bool`

HasDisableSsl returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamProvider) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamProvider) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamProvider) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamProvider) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDomain

`func (o *IamProvider) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IamProvider) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IamProvider) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IamProvider) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEmailRegex

`func (o *IamProvider) GetEmailRegex() string`

GetEmailRegex returns the EmailRegex field if non-nil, zero value otherwise.

### GetEmailRegexOk

`func (o *IamProvider) GetEmailRegexOk() (*string, bool)`

GetEmailRegexOk returns a tuple with the EmailRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRegex

`func (o *IamProvider) SetEmailRegex(v string)`

SetEmailRegex sets EmailRegex field to given value.

### HasEmailRegex

`func (o *IamProvider) HasEmailRegex() bool`

HasEmailRegex returns a boolean if a field has been set.

### GetEnablePkce

`func (o *IamProvider) GetEnablePkce() bool`

GetEnablePkce returns the EnablePkce field if non-nil, zero value otherwise.

### GetEnablePkceOk

`func (o *IamProvider) GetEnablePkceOk() (*bool, bool)`

GetEnablePkceOk returns a tuple with the EnablePkce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePkce

`func (o *IamProvider) SetEnablePkce(v bool)`

SetEnablePkce sets EnablePkce field to given value.

### HasEnablePkce

`func (o *IamProvider) HasEnablePkce() bool`

HasEnablePkce returns a boolean if a field has been set.

### GetEnableProxy

`func (o *IamProvider) GetEnableProxy() bool`

GetEnableProxy returns the EnableProxy field if non-nil, zero value otherwise.

### GetEnableProxyOk

`func (o *IamProvider) GetEnableProxyOk() (*bool, bool)`

GetEnableProxyOk returns a tuple with the EnableProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProxy

`func (o *IamProvider) SetEnableProxy(v bool)`

SetEnableProxy sets EnableProxy field to given value.

### HasEnableProxy

`func (o *IamProvider) HasEnableProxy() bool`

HasEnableProxy returns a boolean if a field has been set.

### GetEnableSignAuthnRequest

`func (o *IamProvider) GetEnableSignAuthnRequest() bool`

GetEnableSignAuthnRequest returns the EnableSignAuthnRequest field if non-nil, zero value otherwise.

### GetEnableSignAuthnRequestOk

`func (o *IamProvider) GetEnableSignAuthnRequestOk() (*bool, bool)`

GetEnableSignAuthnRequestOk returns a tuple with the EnableSignAuthnRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSignAuthnRequest

`func (o *IamProvider) SetEnableSignAuthnRequest(v bool)`

SetEnableSignAuthnRequest sets EnableSignAuthnRequest field to given value.

### HasEnableSignAuthnRequest

`func (o *IamProvider) HasEnableSignAuthnRequest() bool`

HasEnableSignAuthnRequest returns a boolean if a field has been set.

### GetEndpoint

`func (o *IamProvider) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *IamProvider) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *IamProvider) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *IamProvider) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetHost

`func (o *IamProvider) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IamProvider) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IamProvider) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *IamProvider) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHttpHeaders

`func (o *IamProvider) GetHttpHeaders() map[string]string`

GetHttpHeaders returns the HttpHeaders field if non-nil, zero value otherwise.

### GetHttpHeadersOk

`func (o *IamProvider) GetHttpHeadersOk() (*map[string]string, bool)`

GetHttpHeadersOk returns a tuple with the HttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHeaders

`func (o *IamProvider) SetHttpHeaders(v map[string]string)`

SetHttpHeaders sets HttpHeaders field to given value.

### HasHttpHeaders

`func (o *IamProvider) HasHttpHeaders() bool`

HasHttpHeaders returns a boolean if a field has been set.

### GetId

`func (o *IamProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdP

`func (o *IamProvider) GetIdP() string`

GetIdP returns the IdP field if non-nil, zero value otherwise.

### GetIdPOk

`func (o *IamProvider) GetIdPOk() (*string, bool)`

GetIdPOk returns a tuple with the IdP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdP

`func (o *IamProvider) SetIdP(v string)`

SetIdP sets IdP field to given value.

### HasIdP

`func (o *IamProvider) HasIdP() bool`

HasIdP returns a boolean if a field has been set.

### GetIntranetEndpoint

`func (o *IamProvider) GetIntranetEndpoint() string`

GetIntranetEndpoint returns the IntranetEndpoint field if non-nil, zero value otherwise.

### GetIntranetEndpointOk

`func (o *IamProvider) GetIntranetEndpointOk() (*string, bool)`

GetIntranetEndpointOk returns a tuple with the IntranetEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntranetEndpoint

`func (o *IamProvider) SetIntranetEndpoint(v string)`

SetIntranetEndpoint sets IntranetEndpoint field to given value.

### HasIntranetEndpoint

`func (o *IamProvider) HasIntranetEndpoint() bool`

HasIntranetEndpoint returns a boolean if a field has been set.

### GetIssuerUrl

`func (o *IamProvider) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *IamProvider) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *IamProvider) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.

### HasIssuerUrl

`func (o *IamProvider) HasIssuerUrl() bool`

HasIssuerUrl returns a boolean if a field has been set.

### GetMetadata

`func (o *IamProvider) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamProvider) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamProvider) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamProvider) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMethod

`func (o *IamProvider) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *IamProvider) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *IamProvider) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *IamProvider) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetName

`func (o *IamProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamProvider) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamProvider) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamProvider) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamProvider) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPathPrefix

`func (o *IamProvider) GetPathPrefix() string`

GetPathPrefix returns the PathPrefix field if non-nil, zero value otherwise.

### GetPathPrefixOk

`func (o *IamProvider) GetPathPrefixOk() (*string, bool)`

GetPathPrefixOk returns a tuple with the PathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefix

`func (o *IamProvider) SetPathPrefix(v string)`

SetPathPrefix sets PathPrefix field to given value.

### HasPathPrefix

`func (o *IamProvider) HasPathPrefix() bool`

HasPathPrefix returns a boolean if a field has been set.

### GetPort

`func (o *IamProvider) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IamProvider) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IamProvider) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *IamProvider) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProviderUrl

`func (o *IamProvider) GetProviderUrl() string`

GetProviderUrl returns the ProviderUrl field if non-nil, zero value otherwise.

### GetProviderUrlOk

`func (o *IamProvider) GetProviderUrlOk() (*string, bool)`

GetProviderUrlOk returns a tuple with the ProviderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrl

`func (o *IamProvider) SetProviderUrl(v string)`

SetProviderUrl sets ProviderUrl field to given value.

### HasProviderUrl

`func (o *IamProvider) HasProviderUrl() bool`

HasProviderUrl returns a boolean if a field has been set.

### GetReceiver

`func (o *IamProvider) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *IamProvider) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *IamProvider) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *IamProvider) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetRegionId

`func (o *IamProvider) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *IamProvider) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *IamProvider) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *IamProvider) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetScopes

`func (o *IamProvider) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IamProvider) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IamProvider) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *IamProvider) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSignName

`func (o *IamProvider) GetSignName() string`

GetSignName returns the SignName field if non-nil, zero value otherwise.

### GetSignNameOk

`func (o *IamProvider) GetSignNameOk() (*string, bool)`

GetSignNameOk returns a tuple with the SignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignName

`func (o *IamProvider) SetSignName(v string)`

SetSignName sets SignName field to given value.

### HasSignName

`func (o *IamProvider) HasSignName() bool`

HasSignName returns a boolean if a field has been set.

### GetSslMode

`func (o *IamProvider) GetSslMode() string`

GetSslMode returns the SslMode field if non-nil, zero value otherwise.

### GetSslModeOk

`func (o *IamProvider) GetSslModeOk() (*string, bool)`

GetSslModeOk returns a tuple with the SslMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslMode

`func (o *IamProvider) SetSslMode(v string)`

SetSslMode sets SslMode field to given value.

### HasSslMode

`func (o *IamProvider) HasSslMode() bool`

HasSslMode returns a boolean if a field has been set.

### GetSubType

`func (o *IamProvider) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *IamProvider) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *IamProvider) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *IamProvider) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetTemplateCode

`func (o *IamProvider) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *IamProvider) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *IamProvider) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.

### HasTemplateCode

`func (o *IamProvider) HasTemplateCode() bool`

HasTemplateCode returns a boolean if a field has been set.

### GetTitle

`func (o *IamProvider) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IamProvider) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IamProvider) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IamProvider) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *IamProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IamProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IamProvider) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IamProvider) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamProvider) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamProvider) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamProvider) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamProvider) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserMapping

`func (o *IamProvider) GetUserMapping() map[string]string`

GetUserMapping returns the UserMapping field if non-nil, zero value otherwise.

### GetUserMappingOk

`func (o *IamProvider) GetUserMappingOk() (*map[string]string, bool)`

GetUserMappingOk returns a tuple with the UserMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMapping

`func (o *IamProvider) SetUserMapping(v map[string]string)`

SetUserMapping sets UserMapping field to given value.

### HasUserMapping

`func (o *IamProvider) HasUserMapping() bool`

HasUserMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



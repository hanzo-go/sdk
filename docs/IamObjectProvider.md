# IamObjectProvider

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
**CreatedTime** | Pointer to **string** |  | [optional] 
**CustomAuthUrl** | Pointer to **string** |  | [optional] 
**CustomLogo** | Pointer to **string** |  | [optional] 
**CustomTokenUrl** | Pointer to **string** |  | [optional] 
**CustomUserInfoUrl** | Pointer to **string** |  | [optional] 
**DisableSsl** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**EmailRegex** | Pointer to **string** |  | [optional] 
**EnablePkce** | Pointer to **bool** |  | [optional] 
**EnableProxy** | Pointer to **bool** |  | [optional] 
**EnableSignAuthnRequest** | Pointer to **bool** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**HttpHeaders** | Pointer to  |  | [optional] 
**IdP** | Pointer to **string** |  | [optional] 
**IntranetEndpoint** | Pointer to **string** |  | [optional] 
**IssuerUrl** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**PathPrefix** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**ProviderUrl** | Pointer to **string** |  | [optional] 
**Receiver** | Pointer to **string** |  | [optional] 
**RegionId** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **string** |  | [optional] 
**SignName** | Pointer to **string** |  | [optional] 
**SubType** | Pointer to **string** |  | [optional] 
**TemplateCode** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserMapping** | Pointer to  |  | [optional] 

## Methods

### NewIamObjectProvider

`func NewIamObjectProvider() *IamObjectProvider`

NewIamObjectProvider instantiates a new IamObjectProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectProviderWithDefaults

`func NewIamObjectProviderWithDefaults() *IamObjectProvider`

NewIamObjectProviderWithDefaults instantiates a new IamObjectProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *IamObjectProvider) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *IamObjectProvider) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *IamObjectProvider) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *IamObjectProvider) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetBucket

`func (o *IamObjectProvider) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *IamObjectProvider) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *IamObjectProvider) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *IamObjectProvider) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetCategory

`func (o *IamObjectProvider) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IamObjectProvider) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IamObjectProvider) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IamObjectProvider) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCert

`func (o *IamObjectProvider) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *IamObjectProvider) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *IamObjectProvider) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *IamObjectProvider) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetClientId

`func (o *IamObjectProvider) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IamObjectProvider) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IamObjectProvider) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IamObjectProvider) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientId2

`func (o *IamObjectProvider) GetClientId2() string`

GetClientId2 returns the ClientId2 field if non-nil, zero value otherwise.

### GetClientId2Ok

`func (o *IamObjectProvider) GetClientId2Ok() (*string, bool)`

GetClientId2Ok returns a tuple with the ClientId2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId2

`func (o *IamObjectProvider) SetClientId2(v string)`

SetClientId2 sets ClientId2 field to given value.

### HasClientId2

`func (o *IamObjectProvider) HasClientId2() bool`

HasClientId2 returns a boolean if a field has been set.

### GetClientSecret

`func (o *IamObjectProvider) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IamObjectProvider) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IamObjectProvider) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *IamObjectProvider) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientSecret2

`func (o *IamObjectProvider) GetClientSecret2() string`

GetClientSecret2 returns the ClientSecret2 field if non-nil, zero value otherwise.

### GetClientSecret2Ok

`func (o *IamObjectProvider) GetClientSecret2Ok() (*string, bool)`

GetClientSecret2Ok returns a tuple with the ClientSecret2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret2

`func (o *IamObjectProvider) SetClientSecret2(v string)`

SetClientSecret2 sets ClientSecret2 field to given value.

### HasClientSecret2

`func (o *IamObjectProvider) HasClientSecret2() bool`

HasClientSecret2 returns a boolean if a field has been set.

### GetContent

`func (o *IamObjectProvider) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *IamObjectProvider) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *IamObjectProvider) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *IamObjectProvider) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamObjectProvider) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectProvider) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectProvider) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectProvider) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCustomAuthUrl

`func (o *IamObjectProvider) GetCustomAuthUrl() string`

GetCustomAuthUrl returns the CustomAuthUrl field if non-nil, zero value otherwise.

### GetCustomAuthUrlOk

`func (o *IamObjectProvider) GetCustomAuthUrlOk() (*string, bool)`

GetCustomAuthUrlOk returns a tuple with the CustomAuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAuthUrl

`func (o *IamObjectProvider) SetCustomAuthUrl(v string)`

SetCustomAuthUrl sets CustomAuthUrl field to given value.

### HasCustomAuthUrl

`func (o *IamObjectProvider) HasCustomAuthUrl() bool`

HasCustomAuthUrl returns a boolean if a field has been set.

### GetCustomLogo

`func (o *IamObjectProvider) GetCustomLogo() string`

GetCustomLogo returns the CustomLogo field if non-nil, zero value otherwise.

### GetCustomLogoOk

`func (o *IamObjectProvider) GetCustomLogoOk() (*string, bool)`

GetCustomLogoOk returns a tuple with the CustomLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLogo

`func (o *IamObjectProvider) SetCustomLogo(v string)`

SetCustomLogo sets CustomLogo field to given value.

### HasCustomLogo

`func (o *IamObjectProvider) HasCustomLogo() bool`

HasCustomLogo returns a boolean if a field has been set.

### GetCustomTokenUrl

`func (o *IamObjectProvider) GetCustomTokenUrl() string`

GetCustomTokenUrl returns the CustomTokenUrl field if non-nil, zero value otherwise.

### GetCustomTokenUrlOk

`func (o *IamObjectProvider) GetCustomTokenUrlOk() (*string, bool)`

GetCustomTokenUrlOk returns a tuple with the CustomTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTokenUrl

`func (o *IamObjectProvider) SetCustomTokenUrl(v string)`

SetCustomTokenUrl sets CustomTokenUrl field to given value.

### HasCustomTokenUrl

`func (o *IamObjectProvider) HasCustomTokenUrl() bool`

HasCustomTokenUrl returns a boolean if a field has been set.

### GetCustomUserInfoUrl

`func (o *IamObjectProvider) GetCustomUserInfoUrl() string`

GetCustomUserInfoUrl returns the CustomUserInfoUrl field if non-nil, zero value otherwise.

### GetCustomUserInfoUrlOk

`func (o *IamObjectProvider) GetCustomUserInfoUrlOk() (*string, bool)`

GetCustomUserInfoUrlOk returns a tuple with the CustomUserInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserInfoUrl

`func (o *IamObjectProvider) SetCustomUserInfoUrl(v string)`

SetCustomUserInfoUrl sets CustomUserInfoUrl field to given value.

### HasCustomUserInfoUrl

`func (o *IamObjectProvider) HasCustomUserInfoUrl() bool`

HasCustomUserInfoUrl returns a boolean if a field has been set.

### GetDisableSsl

`func (o *IamObjectProvider) GetDisableSsl() bool`

GetDisableSsl returns the DisableSsl field if non-nil, zero value otherwise.

### GetDisableSslOk

`func (o *IamObjectProvider) GetDisableSslOk() (*bool, bool)`

GetDisableSslOk returns a tuple with the DisableSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSsl

`func (o *IamObjectProvider) SetDisableSsl(v bool)`

SetDisableSsl sets DisableSsl field to given value.

### HasDisableSsl

`func (o *IamObjectProvider) HasDisableSsl() bool`

HasDisableSsl returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamObjectProvider) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamObjectProvider) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamObjectProvider) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamObjectProvider) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDomain

`func (o *IamObjectProvider) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IamObjectProvider) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IamObjectProvider) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IamObjectProvider) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEmailRegex

`func (o *IamObjectProvider) GetEmailRegex() string`

GetEmailRegex returns the EmailRegex field if non-nil, zero value otherwise.

### GetEmailRegexOk

`func (o *IamObjectProvider) GetEmailRegexOk() (*string, bool)`

GetEmailRegexOk returns a tuple with the EmailRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRegex

`func (o *IamObjectProvider) SetEmailRegex(v string)`

SetEmailRegex sets EmailRegex field to given value.

### HasEmailRegex

`func (o *IamObjectProvider) HasEmailRegex() bool`

HasEmailRegex returns a boolean if a field has been set.

### GetEnablePkce

`func (o *IamObjectProvider) GetEnablePkce() bool`

GetEnablePkce returns the EnablePkce field if non-nil, zero value otherwise.

### GetEnablePkceOk

`func (o *IamObjectProvider) GetEnablePkceOk() (*bool, bool)`

GetEnablePkceOk returns a tuple with the EnablePkce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePkce

`func (o *IamObjectProvider) SetEnablePkce(v bool)`

SetEnablePkce sets EnablePkce field to given value.

### HasEnablePkce

`func (o *IamObjectProvider) HasEnablePkce() bool`

HasEnablePkce returns a boolean if a field has been set.

### GetEnableProxy

`func (o *IamObjectProvider) GetEnableProxy() bool`

GetEnableProxy returns the EnableProxy field if non-nil, zero value otherwise.

### GetEnableProxyOk

`func (o *IamObjectProvider) GetEnableProxyOk() (*bool, bool)`

GetEnableProxyOk returns a tuple with the EnableProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProxy

`func (o *IamObjectProvider) SetEnableProxy(v bool)`

SetEnableProxy sets EnableProxy field to given value.

### HasEnableProxy

`func (o *IamObjectProvider) HasEnableProxy() bool`

HasEnableProxy returns a boolean if a field has been set.

### GetEnableSignAuthnRequest

`func (o *IamObjectProvider) GetEnableSignAuthnRequest() bool`

GetEnableSignAuthnRequest returns the EnableSignAuthnRequest field if non-nil, zero value otherwise.

### GetEnableSignAuthnRequestOk

`func (o *IamObjectProvider) GetEnableSignAuthnRequestOk() (*bool, bool)`

GetEnableSignAuthnRequestOk returns a tuple with the EnableSignAuthnRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSignAuthnRequest

`func (o *IamObjectProvider) SetEnableSignAuthnRequest(v bool)`

SetEnableSignAuthnRequest sets EnableSignAuthnRequest field to given value.

### HasEnableSignAuthnRequest

`func (o *IamObjectProvider) HasEnableSignAuthnRequest() bool`

HasEnableSignAuthnRequest returns a boolean if a field has been set.

### GetEndpoint

`func (o *IamObjectProvider) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *IamObjectProvider) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *IamObjectProvider) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *IamObjectProvider) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetHost

`func (o *IamObjectProvider) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IamObjectProvider) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IamObjectProvider) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *IamObjectProvider) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHttpHeaders

`func (o *IamObjectProvider) GetHttpHeaders() map[string]string`

GetHttpHeaders returns the HttpHeaders field if non-nil, zero value otherwise.

### GetHttpHeadersOk

`func (o *IamObjectProvider) GetHttpHeadersOk() (*map[string]string, bool)`

GetHttpHeadersOk returns a tuple with the HttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHeaders

`func (o *IamObjectProvider) SetHttpHeaders(v map[string]string)`

SetHttpHeaders sets HttpHeaders field to given value.

### HasHttpHeaders

`func (o *IamObjectProvider) HasHttpHeaders() bool`

HasHttpHeaders returns a boolean if a field has been set.

### SetHttpHeadersNil

`func (o *IamObjectProvider) SetHttpHeadersNil(b bool)`

 SetHttpHeadersNil sets the value for HttpHeaders to be an explicit nil

### UnsetHttpHeaders
`func (o *IamObjectProvider) UnsetHttpHeaders()`

UnsetHttpHeaders ensures that no value is present for HttpHeaders, not even an explicit nil
### GetIdP

`func (o *IamObjectProvider) GetIdP() string`

GetIdP returns the IdP field if non-nil, zero value otherwise.

### GetIdPOk

`func (o *IamObjectProvider) GetIdPOk() (*string, bool)`

GetIdPOk returns a tuple with the IdP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdP

`func (o *IamObjectProvider) SetIdP(v string)`

SetIdP sets IdP field to given value.

### HasIdP

`func (o *IamObjectProvider) HasIdP() bool`

HasIdP returns a boolean if a field has been set.

### GetIntranetEndpoint

`func (o *IamObjectProvider) GetIntranetEndpoint() string`

GetIntranetEndpoint returns the IntranetEndpoint field if non-nil, zero value otherwise.

### GetIntranetEndpointOk

`func (o *IamObjectProvider) GetIntranetEndpointOk() (*string, bool)`

GetIntranetEndpointOk returns a tuple with the IntranetEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntranetEndpoint

`func (o *IamObjectProvider) SetIntranetEndpoint(v string)`

SetIntranetEndpoint sets IntranetEndpoint field to given value.

### HasIntranetEndpoint

`func (o *IamObjectProvider) HasIntranetEndpoint() bool`

HasIntranetEndpoint returns a boolean if a field has been set.

### GetIssuerUrl

`func (o *IamObjectProvider) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *IamObjectProvider) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *IamObjectProvider) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.

### HasIssuerUrl

`func (o *IamObjectProvider) HasIssuerUrl() bool`

HasIssuerUrl returns a boolean if a field has been set.

### GetMetadata

`func (o *IamObjectProvider) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamObjectProvider) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamObjectProvider) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamObjectProvider) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMethod

`func (o *IamObjectProvider) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *IamObjectProvider) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *IamObjectProvider) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *IamObjectProvider) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetName

`func (o *IamObjectProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectProvider) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectProvider) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectProvider) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectProvider) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPathPrefix

`func (o *IamObjectProvider) GetPathPrefix() string`

GetPathPrefix returns the PathPrefix field if non-nil, zero value otherwise.

### GetPathPrefixOk

`func (o *IamObjectProvider) GetPathPrefixOk() (*string, bool)`

GetPathPrefixOk returns a tuple with the PathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefix

`func (o *IamObjectProvider) SetPathPrefix(v string)`

SetPathPrefix sets PathPrefix field to given value.

### HasPathPrefix

`func (o *IamObjectProvider) HasPathPrefix() bool`

HasPathPrefix returns a boolean if a field has been set.

### GetPort

`func (o *IamObjectProvider) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IamObjectProvider) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IamObjectProvider) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *IamObjectProvider) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProviderUrl

`func (o *IamObjectProvider) GetProviderUrl() string`

GetProviderUrl returns the ProviderUrl field if non-nil, zero value otherwise.

### GetProviderUrlOk

`func (o *IamObjectProvider) GetProviderUrlOk() (*string, bool)`

GetProviderUrlOk returns a tuple with the ProviderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrl

`func (o *IamObjectProvider) SetProviderUrl(v string)`

SetProviderUrl sets ProviderUrl field to given value.

### HasProviderUrl

`func (o *IamObjectProvider) HasProviderUrl() bool`

HasProviderUrl returns a boolean if a field has been set.

### GetReceiver

`func (o *IamObjectProvider) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *IamObjectProvider) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *IamObjectProvider) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *IamObjectProvider) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetRegionId

`func (o *IamObjectProvider) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *IamObjectProvider) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *IamObjectProvider) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *IamObjectProvider) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetScopes

`func (o *IamObjectProvider) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IamObjectProvider) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IamObjectProvider) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *IamObjectProvider) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSignName

`func (o *IamObjectProvider) GetSignName() string`

GetSignName returns the SignName field if non-nil, zero value otherwise.

### GetSignNameOk

`func (o *IamObjectProvider) GetSignNameOk() (*string, bool)`

GetSignNameOk returns a tuple with the SignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignName

`func (o *IamObjectProvider) SetSignName(v string)`

SetSignName sets SignName field to given value.

### HasSignName

`func (o *IamObjectProvider) HasSignName() bool`

HasSignName returns a boolean if a field has been set.

### GetSubType

`func (o *IamObjectProvider) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *IamObjectProvider) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *IamObjectProvider) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *IamObjectProvider) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetTemplateCode

`func (o *IamObjectProvider) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *IamObjectProvider) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *IamObjectProvider) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.

### HasTemplateCode

`func (o *IamObjectProvider) HasTemplateCode() bool`

HasTemplateCode returns a boolean if a field has been set.

### GetTitle

`func (o *IamObjectProvider) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IamObjectProvider) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IamObjectProvider) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IamObjectProvider) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *IamObjectProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IamObjectProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IamObjectProvider) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IamObjectProvider) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserMapping

`func (o *IamObjectProvider) GetUserMapping() map[string]string`

GetUserMapping returns the UserMapping field if non-nil, zero value otherwise.

### GetUserMappingOk

`func (o *IamObjectProvider) GetUserMappingOk() (*map[string]string, bool)`

GetUserMappingOk returns a tuple with the UserMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMapping

`func (o *IamObjectProvider) SetUserMapping(v map[string]string)`

SetUserMapping sets UserMapping field to given value.

### HasUserMapping

`func (o *IamObjectProvider) HasUserMapping() bool`

HasUserMapping returns a boolean if a field has been set.

### SetUserMappingNil

`func (o *IamObjectProvider) SetUserMappingNil(b bool)`

 SetUserMappingNil sets the value for UserMapping to be an explicit nil

### UnsetUserMapping
`func (o *IamObjectProvider) UnsetUserMapping()`

UnsetUserMapping ensures that no value is present for UserMapping, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



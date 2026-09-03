# IamApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffiliationUrl** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Cert** | Pointer to **string** |  | [optional] 
**CertObj** | Pointer to [**IamCert**](IamCert.md) |  | [optional] 
**CertPublicKey** | Pointer to **string** |  | [optional] 
**ClientCert** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** | ClientId is the OAuth2/OIDC client identifier and the GLOBAL key every confidential-client resolver authenticates against (store.GetApplicationByClientId, the mint gates, Basic auth). It MUST be globally unique across ALL owners — a collision would let one app shadow another at that key. This store persists each entity as a JSON document in a shared table, so there is no per-field column to carry a DB UNIQUE index; uniqueness is enforced at the write in applications.Create/Update (ensureClientIdUnique), exactly as the (owner,name) natural key is, and store.GetApplicationByClientId resolves admin-preferring as defense-in-depth. | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**CodeResendTimeout** | Pointer to **int64** |  | [optional] 
**CookieExpireInHours** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**CustomScopes** | Pointer to [**[]IamScopeDescription**](IamScopeDescription.md) |  | [optional] 
**DefaultGroup** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisableSamlAttributes** | Pointer to **bool** |  | [optional] 
**DisableSignin** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**EnableAutoSignin** | Pointer to **bool** |  | [optional] 
**EnableCodeSignin** | Pointer to **bool** |  | [optional] 
**EnableExclusiveSignin** | Pointer to **bool** |  | [optional] 
**EnableLinkWithEmail** | Pointer to **bool** |  | [optional] 
**EnablePassword** | Pointer to **bool** |  | [optional] 
**EnableSamlAssertionSignature** | Pointer to **bool** |  | [optional] 
**EnableSamlC14n10** | Pointer to **bool** |  | [optional] 
**EnableSamlCompress** | Pointer to **bool** |  | [optional] 
**EnableSamlPostBinding** | Pointer to **bool** |  | [optional] 
**EnableSignUp** | Pointer to **bool** |  | [optional] 
**EnableSigninSession** | Pointer to **bool** |  | [optional] 
**EnableWebAuthn** | Pointer to **bool** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**ExpireInHours** | Pointer to **float64** |  | [optional] 
**FailedSigninFrozenTime** | Pointer to **int64** |  | [optional] 
**FailedSigninLimit** | Pointer to **int64** |  | [optional] 
**Favicon** | Pointer to **string** |  | [optional] 
**FooterHtml** | Pointer to **string** |  | [optional] 
**ForcedRedirectOrigin** | Pointer to **string** |  | [optional] 
**ForgetUrl** | Pointer to **string** |  | [optional] 
**FormBackgroundUrl** | Pointer to **string** |  | [optional] 
**FormBackgroundUrlMobile** | Pointer to **string** |  | [optional] 
**FormCss** | Pointer to **string** |  | [optional] 
**FormCssMobile** | Pointer to **string** |  | [optional] 
**FormOffset** | Pointer to **int64** |  | [optional] 
**FormSideHtml** | Pointer to **string** |  | [optional] 
**GrantTypes** | Pointer to **[]string** |  | [optional] 
**HeaderHtml** | Pointer to **string** |  | [optional] 
**HomepageUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IpRestriction** | Pointer to **string** |  | [optional] 
**IpWhitelist** | Pointer to **string** |  | [optional] 
**IsShared** | Pointer to **bool** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **int64** |  | [optional] 
**OrgChoiceMode** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**OrganizationObj** | Pointer to [**IamOrganization**](IamOrganization.md) |  | [optional] 
**OtherDomains** | Pointer to **[]string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Providers** | Pointer to [**[]IamProviderItem**](IamProviderItem.md) |  | [optional] 
**RedirectUris** | Pointer to **[]string** |  | [optional] 
**RefreshExpireInHours** | Pointer to **float64** |  | [optional] 
**SamlAttributes** | Pointer to [**[]IamSamlItem**](IamSamlItem.md) |  | [optional] 
**SamlHashAlgorithm** | Pointer to **string** |  | [optional] 
**SamlReplyUrl** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to [**[]IamScopeItem**](IamScopeItem.md) |  | [optional] 
**SigninHtml** | Pointer to **string** |  | [optional] 
**SigninItems** | Pointer to [**[]IamSigninItem**](IamSigninItem.md) |  | [optional] 
**SigninMethods** | Pointer to [**[]IamSigninMethod**](IamSigninMethod.md) |  | [optional] 
**SigninUrl** | Pointer to **string** |  | [optional] 
**SignupHtml** | Pointer to **string** |  | [optional] 
**SignupItems** | Pointer to [**[]IamSignupItem**](IamSignupItem.md) |  | [optional] 
**SignupUrl** | Pointer to **string** |  | [optional] 
**SslCert** | Pointer to **string** |  | [optional] 
**SslMode** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TermsOfUse** | Pointer to **string** |  | [optional] 
**ThemeData** | Pointer to [**IamThemeData**](IamThemeData.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TokenAttributes** | Pointer to [**[]IamJwtItem**](IamJwtItem.md) |  | [optional] 
**TokenFields** | Pointer to **[]string** |  | [optional] 
**TokenFormat** | Pointer to **string** |  | [optional] 
**TokenSigningMethod** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpstreamHost** | Pointer to **string** |  | [optional] 
**UseEmailAsSamlNameId** | Pointer to **bool** |  | [optional] 

## Methods

### NewIamApplication

`func NewIamApplication() *IamApplication`

NewIamApplication instantiates a new IamApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamApplicationWithDefaults

`func NewIamApplicationWithDefaults() *IamApplication`

NewIamApplicationWithDefaults instantiates a new IamApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffiliationUrl

`func (o *IamApplication) GetAffiliationUrl() string`

GetAffiliationUrl returns the AffiliationUrl field if non-nil, zero value otherwise.

### GetAffiliationUrlOk

`func (o *IamApplication) GetAffiliationUrlOk() (*string, bool)`

GetAffiliationUrlOk returns a tuple with the AffiliationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliationUrl

`func (o *IamApplication) SetAffiliationUrl(v string)`

SetAffiliationUrl sets AffiliationUrl field to given value.

### HasAffiliationUrl

`func (o *IamApplication) HasAffiliationUrl() bool`

HasAffiliationUrl returns a boolean if a field has been set.

### GetCategory

`func (o *IamApplication) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IamApplication) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IamApplication) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IamApplication) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCert

`func (o *IamApplication) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *IamApplication) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *IamApplication) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *IamApplication) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetCertObj

`func (o *IamApplication) GetCertObj() IamCert`

GetCertObj returns the CertObj field if non-nil, zero value otherwise.

### GetCertObjOk

`func (o *IamApplication) GetCertObjOk() (*IamCert, bool)`

GetCertObjOk returns a tuple with the CertObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertObj

`func (o *IamApplication) SetCertObj(v IamCert)`

SetCertObj sets CertObj field to given value.

### HasCertObj

`func (o *IamApplication) HasCertObj() bool`

HasCertObj returns a boolean if a field has been set.

### GetCertPublicKey

`func (o *IamApplication) GetCertPublicKey() string`

GetCertPublicKey returns the CertPublicKey field if non-nil, zero value otherwise.

### GetCertPublicKeyOk

`func (o *IamApplication) GetCertPublicKeyOk() (*string, bool)`

GetCertPublicKeyOk returns a tuple with the CertPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertPublicKey

`func (o *IamApplication) SetCertPublicKey(v string)`

SetCertPublicKey sets CertPublicKey field to given value.

### HasCertPublicKey

`func (o *IamApplication) HasCertPublicKey() bool`

HasCertPublicKey returns a boolean if a field has been set.

### GetClientCert

`func (o *IamApplication) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *IamApplication) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *IamApplication) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *IamApplication) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetClientId

`func (o *IamApplication) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IamApplication) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IamApplication) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IamApplication) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *IamApplication) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IamApplication) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IamApplication) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *IamApplication) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetCodeResendTimeout

`func (o *IamApplication) GetCodeResendTimeout() int64`

GetCodeResendTimeout returns the CodeResendTimeout field if non-nil, zero value otherwise.

### GetCodeResendTimeoutOk

`func (o *IamApplication) GetCodeResendTimeoutOk() (*int64, bool)`

GetCodeResendTimeoutOk returns a tuple with the CodeResendTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeResendTimeout

`func (o *IamApplication) SetCodeResendTimeout(v int64)`

SetCodeResendTimeout sets CodeResendTimeout field to given value.

### HasCodeResendTimeout

`func (o *IamApplication) HasCodeResendTimeout() bool`

HasCodeResendTimeout returns a boolean if a field has been set.

### GetCookieExpireInHours

`func (o *IamApplication) GetCookieExpireInHours() int64`

GetCookieExpireInHours returns the CookieExpireInHours field if non-nil, zero value otherwise.

### GetCookieExpireInHoursOk

`func (o *IamApplication) GetCookieExpireInHoursOk() (*int64, bool)`

GetCookieExpireInHoursOk returns a tuple with the CookieExpireInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieExpireInHours

`func (o *IamApplication) SetCookieExpireInHours(v int64)`

SetCookieExpireInHours sets CookieExpireInHours field to given value.

### HasCookieExpireInHours

`func (o *IamApplication) HasCookieExpireInHours() bool`

HasCookieExpireInHours returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IamApplication) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamApplication) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamApplication) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamApplication) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamApplication) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamApplication) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamApplication) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamApplication) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCustomScopes

`func (o *IamApplication) GetCustomScopes() []IamScopeDescription`

GetCustomScopes returns the CustomScopes field if non-nil, zero value otherwise.

### GetCustomScopesOk

`func (o *IamApplication) GetCustomScopesOk() (*[]IamScopeDescription, bool)`

GetCustomScopesOk returns a tuple with the CustomScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomScopes

`func (o *IamApplication) SetCustomScopes(v []IamScopeDescription)`

SetCustomScopes sets CustomScopes field to given value.

### HasCustomScopes

`func (o *IamApplication) HasCustomScopes() bool`

HasCustomScopes returns a boolean if a field has been set.

### GetDefaultGroup

`func (o *IamApplication) GetDefaultGroup() string`

GetDefaultGroup returns the DefaultGroup field if non-nil, zero value otherwise.

### GetDefaultGroupOk

`func (o *IamApplication) GetDefaultGroupOk() (*string, bool)`

GetDefaultGroupOk returns a tuple with the DefaultGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroup

`func (o *IamApplication) SetDefaultGroup(v string)`

SetDefaultGroup sets DefaultGroup field to given value.

### HasDefaultGroup

`func (o *IamApplication) HasDefaultGroup() bool`

HasDefaultGroup returns a boolean if a field has been set.

### GetDeleted

`func (o *IamApplication) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamApplication) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamApplication) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamApplication) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDescription

`func (o *IamApplication) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamApplication) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamApplication) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamApplication) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableSamlAttributes

`func (o *IamApplication) GetDisableSamlAttributes() bool`

GetDisableSamlAttributes returns the DisableSamlAttributes field if non-nil, zero value otherwise.

### GetDisableSamlAttributesOk

`func (o *IamApplication) GetDisableSamlAttributesOk() (*bool, bool)`

GetDisableSamlAttributesOk returns a tuple with the DisableSamlAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSamlAttributes

`func (o *IamApplication) SetDisableSamlAttributes(v bool)`

SetDisableSamlAttributes sets DisableSamlAttributes field to given value.

### HasDisableSamlAttributes

`func (o *IamApplication) HasDisableSamlAttributes() bool`

HasDisableSamlAttributes returns a boolean if a field has been set.

### GetDisableSignin

`func (o *IamApplication) GetDisableSignin() bool`

GetDisableSignin returns the DisableSignin field if non-nil, zero value otherwise.

### GetDisableSigninOk

`func (o *IamApplication) GetDisableSigninOk() (*bool, bool)`

GetDisableSigninOk returns a tuple with the DisableSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSignin

`func (o *IamApplication) SetDisableSignin(v bool)`

SetDisableSignin sets DisableSignin field to given value.

### HasDisableSignin

`func (o *IamApplication) HasDisableSignin() bool`

HasDisableSignin returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamApplication) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamApplication) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamApplication) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamApplication) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDomain

`func (o *IamApplication) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IamApplication) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IamApplication) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IamApplication) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEnableAutoSignin

`func (o *IamApplication) GetEnableAutoSignin() bool`

GetEnableAutoSignin returns the EnableAutoSignin field if non-nil, zero value otherwise.

### GetEnableAutoSigninOk

`func (o *IamApplication) GetEnableAutoSigninOk() (*bool, bool)`

GetEnableAutoSigninOk returns a tuple with the EnableAutoSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoSignin

`func (o *IamApplication) SetEnableAutoSignin(v bool)`

SetEnableAutoSignin sets EnableAutoSignin field to given value.

### HasEnableAutoSignin

`func (o *IamApplication) HasEnableAutoSignin() bool`

HasEnableAutoSignin returns a boolean if a field has been set.

### GetEnableCodeSignin

`func (o *IamApplication) GetEnableCodeSignin() bool`

GetEnableCodeSignin returns the EnableCodeSignin field if non-nil, zero value otherwise.

### GetEnableCodeSigninOk

`func (o *IamApplication) GetEnableCodeSigninOk() (*bool, bool)`

GetEnableCodeSigninOk returns a tuple with the EnableCodeSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCodeSignin

`func (o *IamApplication) SetEnableCodeSignin(v bool)`

SetEnableCodeSignin sets EnableCodeSignin field to given value.

### HasEnableCodeSignin

`func (o *IamApplication) HasEnableCodeSignin() bool`

HasEnableCodeSignin returns a boolean if a field has been set.

### GetEnableExclusiveSignin

`func (o *IamApplication) GetEnableExclusiveSignin() bool`

GetEnableExclusiveSignin returns the EnableExclusiveSignin field if non-nil, zero value otherwise.

### GetEnableExclusiveSigninOk

`func (o *IamApplication) GetEnableExclusiveSigninOk() (*bool, bool)`

GetEnableExclusiveSigninOk returns a tuple with the EnableExclusiveSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExclusiveSignin

`func (o *IamApplication) SetEnableExclusiveSignin(v bool)`

SetEnableExclusiveSignin sets EnableExclusiveSignin field to given value.

### HasEnableExclusiveSignin

`func (o *IamApplication) HasEnableExclusiveSignin() bool`

HasEnableExclusiveSignin returns a boolean if a field has been set.

### GetEnableLinkWithEmail

`func (o *IamApplication) GetEnableLinkWithEmail() bool`

GetEnableLinkWithEmail returns the EnableLinkWithEmail field if non-nil, zero value otherwise.

### GetEnableLinkWithEmailOk

`func (o *IamApplication) GetEnableLinkWithEmailOk() (*bool, bool)`

GetEnableLinkWithEmailOk returns a tuple with the EnableLinkWithEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLinkWithEmail

`func (o *IamApplication) SetEnableLinkWithEmail(v bool)`

SetEnableLinkWithEmail sets EnableLinkWithEmail field to given value.

### HasEnableLinkWithEmail

`func (o *IamApplication) HasEnableLinkWithEmail() bool`

HasEnableLinkWithEmail returns a boolean if a field has been set.

### GetEnablePassword

`func (o *IamApplication) GetEnablePassword() bool`

GetEnablePassword returns the EnablePassword field if non-nil, zero value otherwise.

### GetEnablePasswordOk

`func (o *IamApplication) GetEnablePasswordOk() (*bool, bool)`

GetEnablePasswordOk returns a tuple with the EnablePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePassword

`func (o *IamApplication) SetEnablePassword(v bool)`

SetEnablePassword sets EnablePassword field to given value.

### HasEnablePassword

`func (o *IamApplication) HasEnablePassword() bool`

HasEnablePassword returns a boolean if a field has been set.

### GetEnableSamlAssertionSignature

`func (o *IamApplication) GetEnableSamlAssertionSignature() bool`

GetEnableSamlAssertionSignature returns the EnableSamlAssertionSignature field if non-nil, zero value otherwise.

### GetEnableSamlAssertionSignatureOk

`func (o *IamApplication) GetEnableSamlAssertionSignatureOk() (*bool, bool)`

GetEnableSamlAssertionSignatureOk returns a tuple with the EnableSamlAssertionSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSamlAssertionSignature

`func (o *IamApplication) SetEnableSamlAssertionSignature(v bool)`

SetEnableSamlAssertionSignature sets EnableSamlAssertionSignature field to given value.

### HasEnableSamlAssertionSignature

`func (o *IamApplication) HasEnableSamlAssertionSignature() bool`

HasEnableSamlAssertionSignature returns a boolean if a field has been set.

### GetEnableSamlC14n10

`func (o *IamApplication) GetEnableSamlC14n10() bool`

GetEnableSamlC14n10 returns the EnableSamlC14n10 field if non-nil, zero value otherwise.

### GetEnableSamlC14n10Ok

`func (o *IamApplication) GetEnableSamlC14n10Ok() (*bool, bool)`

GetEnableSamlC14n10Ok returns a tuple with the EnableSamlC14n10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSamlC14n10

`func (o *IamApplication) SetEnableSamlC14n10(v bool)`

SetEnableSamlC14n10 sets EnableSamlC14n10 field to given value.

### HasEnableSamlC14n10

`func (o *IamApplication) HasEnableSamlC14n10() bool`

HasEnableSamlC14n10 returns a boolean if a field has been set.

### GetEnableSamlCompress

`func (o *IamApplication) GetEnableSamlCompress() bool`

GetEnableSamlCompress returns the EnableSamlCompress field if non-nil, zero value otherwise.

### GetEnableSamlCompressOk

`func (o *IamApplication) GetEnableSamlCompressOk() (*bool, bool)`

GetEnableSamlCompressOk returns a tuple with the EnableSamlCompress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSamlCompress

`func (o *IamApplication) SetEnableSamlCompress(v bool)`

SetEnableSamlCompress sets EnableSamlCompress field to given value.

### HasEnableSamlCompress

`func (o *IamApplication) HasEnableSamlCompress() bool`

HasEnableSamlCompress returns a boolean if a field has been set.

### GetEnableSamlPostBinding

`func (o *IamApplication) GetEnableSamlPostBinding() bool`

GetEnableSamlPostBinding returns the EnableSamlPostBinding field if non-nil, zero value otherwise.

### GetEnableSamlPostBindingOk

`func (o *IamApplication) GetEnableSamlPostBindingOk() (*bool, bool)`

GetEnableSamlPostBindingOk returns a tuple with the EnableSamlPostBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSamlPostBinding

`func (o *IamApplication) SetEnableSamlPostBinding(v bool)`

SetEnableSamlPostBinding sets EnableSamlPostBinding field to given value.

### HasEnableSamlPostBinding

`func (o *IamApplication) HasEnableSamlPostBinding() bool`

HasEnableSamlPostBinding returns a boolean if a field has been set.

### GetEnableSignUp

`func (o *IamApplication) GetEnableSignUp() bool`

GetEnableSignUp returns the EnableSignUp field if non-nil, zero value otherwise.

### GetEnableSignUpOk

`func (o *IamApplication) GetEnableSignUpOk() (*bool, bool)`

GetEnableSignUpOk returns a tuple with the EnableSignUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSignUp

`func (o *IamApplication) SetEnableSignUp(v bool)`

SetEnableSignUp sets EnableSignUp field to given value.

### HasEnableSignUp

`func (o *IamApplication) HasEnableSignUp() bool`

HasEnableSignUp returns a boolean if a field has been set.

### GetEnableSigninSession

`func (o *IamApplication) GetEnableSigninSession() bool`

GetEnableSigninSession returns the EnableSigninSession field if non-nil, zero value otherwise.

### GetEnableSigninSessionOk

`func (o *IamApplication) GetEnableSigninSessionOk() (*bool, bool)`

GetEnableSigninSessionOk returns a tuple with the EnableSigninSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSigninSession

`func (o *IamApplication) SetEnableSigninSession(v bool)`

SetEnableSigninSession sets EnableSigninSession field to given value.

### HasEnableSigninSession

`func (o *IamApplication) HasEnableSigninSession() bool`

HasEnableSigninSession returns a boolean if a field has been set.

### GetEnableWebAuthn

`func (o *IamApplication) GetEnableWebAuthn() bool`

GetEnableWebAuthn returns the EnableWebAuthn field if non-nil, zero value otherwise.

### GetEnableWebAuthnOk

`func (o *IamApplication) GetEnableWebAuthnOk() (*bool, bool)`

GetEnableWebAuthnOk returns a tuple with the EnableWebAuthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWebAuthn

`func (o *IamApplication) SetEnableWebAuthn(v bool)`

SetEnableWebAuthn sets EnableWebAuthn field to given value.

### HasEnableWebAuthn

`func (o *IamApplication) HasEnableWebAuthn() bool`

HasEnableWebAuthn returns a boolean if a field has been set.

### GetEnvironment

`func (o *IamApplication) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *IamApplication) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *IamApplication) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *IamApplication) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetExpireInHours

`func (o *IamApplication) GetExpireInHours() float64`

GetExpireInHours returns the ExpireInHours field if non-nil, zero value otherwise.

### GetExpireInHoursOk

`func (o *IamApplication) GetExpireInHoursOk() (*float64, bool)`

GetExpireInHoursOk returns a tuple with the ExpireInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireInHours

`func (o *IamApplication) SetExpireInHours(v float64)`

SetExpireInHours sets ExpireInHours field to given value.

### HasExpireInHours

`func (o *IamApplication) HasExpireInHours() bool`

HasExpireInHours returns a boolean if a field has been set.

### GetFailedSigninFrozenTime

`func (o *IamApplication) GetFailedSigninFrozenTime() int64`

GetFailedSigninFrozenTime returns the FailedSigninFrozenTime field if non-nil, zero value otherwise.

### GetFailedSigninFrozenTimeOk

`func (o *IamApplication) GetFailedSigninFrozenTimeOk() (*int64, bool)`

GetFailedSigninFrozenTimeOk returns a tuple with the FailedSigninFrozenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedSigninFrozenTime

`func (o *IamApplication) SetFailedSigninFrozenTime(v int64)`

SetFailedSigninFrozenTime sets FailedSigninFrozenTime field to given value.

### HasFailedSigninFrozenTime

`func (o *IamApplication) HasFailedSigninFrozenTime() bool`

HasFailedSigninFrozenTime returns a boolean if a field has been set.

### GetFailedSigninLimit

`func (o *IamApplication) GetFailedSigninLimit() int64`

GetFailedSigninLimit returns the FailedSigninLimit field if non-nil, zero value otherwise.

### GetFailedSigninLimitOk

`func (o *IamApplication) GetFailedSigninLimitOk() (*int64, bool)`

GetFailedSigninLimitOk returns a tuple with the FailedSigninLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedSigninLimit

`func (o *IamApplication) SetFailedSigninLimit(v int64)`

SetFailedSigninLimit sets FailedSigninLimit field to given value.

### HasFailedSigninLimit

`func (o *IamApplication) HasFailedSigninLimit() bool`

HasFailedSigninLimit returns a boolean if a field has been set.

### GetFavicon

`func (o *IamApplication) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *IamApplication) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *IamApplication) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *IamApplication) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### GetFooterHtml

`func (o *IamApplication) GetFooterHtml() string`

GetFooterHtml returns the FooterHtml field if non-nil, zero value otherwise.

### GetFooterHtmlOk

`func (o *IamApplication) GetFooterHtmlOk() (*string, bool)`

GetFooterHtmlOk returns a tuple with the FooterHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterHtml

`func (o *IamApplication) SetFooterHtml(v string)`

SetFooterHtml sets FooterHtml field to given value.

### HasFooterHtml

`func (o *IamApplication) HasFooterHtml() bool`

HasFooterHtml returns a boolean if a field has been set.

### GetForcedRedirectOrigin

`func (o *IamApplication) GetForcedRedirectOrigin() string`

GetForcedRedirectOrigin returns the ForcedRedirectOrigin field if non-nil, zero value otherwise.

### GetForcedRedirectOriginOk

`func (o *IamApplication) GetForcedRedirectOriginOk() (*string, bool)`

GetForcedRedirectOriginOk returns a tuple with the ForcedRedirectOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedRedirectOrigin

`func (o *IamApplication) SetForcedRedirectOrigin(v string)`

SetForcedRedirectOrigin sets ForcedRedirectOrigin field to given value.

### HasForcedRedirectOrigin

`func (o *IamApplication) HasForcedRedirectOrigin() bool`

HasForcedRedirectOrigin returns a boolean if a field has been set.

### GetForgetUrl

`func (o *IamApplication) GetForgetUrl() string`

GetForgetUrl returns the ForgetUrl field if non-nil, zero value otherwise.

### GetForgetUrlOk

`func (o *IamApplication) GetForgetUrlOk() (*string, bool)`

GetForgetUrlOk returns a tuple with the ForgetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgetUrl

`func (o *IamApplication) SetForgetUrl(v string)`

SetForgetUrl sets ForgetUrl field to given value.

### HasForgetUrl

`func (o *IamApplication) HasForgetUrl() bool`

HasForgetUrl returns a boolean if a field has been set.

### GetFormBackgroundUrl

`func (o *IamApplication) GetFormBackgroundUrl() string`

GetFormBackgroundUrl returns the FormBackgroundUrl field if non-nil, zero value otherwise.

### GetFormBackgroundUrlOk

`func (o *IamApplication) GetFormBackgroundUrlOk() (*string, bool)`

GetFormBackgroundUrlOk returns a tuple with the FormBackgroundUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormBackgroundUrl

`func (o *IamApplication) SetFormBackgroundUrl(v string)`

SetFormBackgroundUrl sets FormBackgroundUrl field to given value.

### HasFormBackgroundUrl

`func (o *IamApplication) HasFormBackgroundUrl() bool`

HasFormBackgroundUrl returns a boolean if a field has been set.

### GetFormBackgroundUrlMobile

`func (o *IamApplication) GetFormBackgroundUrlMobile() string`

GetFormBackgroundUrlMobile returns the FormBackgroundUrlMobile field if non-nil, zero value otherwise.

### GetFormBackgroundUrlMobileOk

`func (o *IamApplication) GetFormBackgroundUrlMobileOk() (*string, bool)`

GetFormBackgroundUrlMobileOk returns a tuple with the FormBackgroundUrlMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormBackgroundUrlMobile

`func (o *IamApplication) SetFormBackgroundUrlMobile(v string)`

SetFormBackgroundUrlMobile sets FormBackgroundUrlMobile field to given value.

### HasFormBackgroundUrlMobile

`func (o *IamApplication) HasFormBackgroundUrlMobile() bool`

HasFormBackgroundUrlMobile returns a boolean if a field has been set.

### GetFormCss

`func (o *IamApplication) GetFormCss() string`

GetFormCss returns the FormCss field if non-nil, zero value otherwise.

### GetFormCssOk

`func (o *IamApplication) GetFormCssOk() (*string, bool)`

GetFormCssOk returns a tuple with the FormCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormCss

`func (o *IamApplication) SetFormCss(v string)`

SetFormCss sets FormCss field to given value.

### HasFormCss

`func (o *IamApplication) HasFormCss() bool`

HasFormCss returns a boolean if a field has been set.

### GetFormCssMobile

`func (o *IamApplication) GetFormCssMobile() string`

GetFormCssMobile returns the FormCssMobile field if non-nil, zero value otherwise.

### GetFormCssMobileOk

`func (o *IamApplication) GetFormCssMobileOk() (*string, bool)`

GetFormCssMobileOk returns a tuple with the FormCssMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormCssMobile

`func (o *IamApplication) SetFormCssMobile(v string)`

SetFormCssMobile sets FormCssMobile field to given value.

### HasFormCssMobile

`func (o *IamApplication) HasFormCssMobile() bool`

HasFormCssMobile returns a boolean if a field has been set.

### GetFormOffset

`func (o *IamApplication) GetFormOffset() int64`

GetFormOffset returns the FormOffset field if non-nil, zero value otherwise.

### GetFormOffsetOk

`func (o *IamApplication) GetFormOffsetOk() (*int64, bool)`

GetFormOffsetOk returns a tuple with the FormOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormOffset

`func (o *IamApplication) SetFormOffset(v int64)`

SetFormOffset sets FormOffset field to given value.

### HasFormOffset

`func (o *IamApplication) HasFormOffset() bool`

HasFormOffset returns a boolean if a field has been set.

### GetFormSideHtml

`func (o *IamApplication) GetFormSideHtml() string`

GetFormSideHtml returns the FormSideHtml field if non-nil, zero value otherwise.

### GetFormSideHtmlOk

`func (o *IamApplication) GetFormSideHtmlOk() (*string, bool)`

GetFormSideHtmlOk returns a tuple with the FormSideHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormSideHtml

`func (o *IamApplication) SetFormSideHtml(v string)`

SetFormSideHtml sets FormSideHtml field to given value.

### HasFormSideHtml

`func (o *IamApplication) HasFormSideHtml() bool`

HasFormSideHtml returns a boolean if a field has been set.

### GetGrantTypes

`func (o *IamApplication) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *IamApplication) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *IamApplication) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *IamApplication) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetHeaderHtml

`func (o *IamApplication) GetHeaderHtml() string`

GetHeaderHtml returns the HeaderHtml field if non-nil, zero value otherwise.

### GetHeaderHtmlOk

`func (o *IamApplication) GetHeaderHtmlOk() (*string, bool)`

GetHeaderHtmlOk returns a tuple with the HeaderHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderHtml

`func (o *IamApplication) SetHeaderHtml(v string)`

SetHeaderHtml sets HeaderHtml field to given value.

### HasHeaderHtml

`func (o *IamApplication) HasHeaderHtml() bool`

HasHeaderHtml returns a boolean if a field has been set.

### GetHomepageUrl

`func (o *IamApplication) GetHomepageUrl() string`

GetHomepageUrl returns the HomepageUrl field if non-nil, zero value otherwise.

### GetHomepageUrlOk

`func (o *IamApplication) GetHomepageUrlOk() (*string, bool)`

GetHomepageUrlOk returns a tuple with the HomepageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepageUrl

`func (o *IamApplication) SetHomepageUrl(v string)`

SetHomepageUrl sets HomepageUrl field to given value.

### HasHomepageUrl

`func (o *IamApplication) HasHomepageUrl() bool`

HasHomepageUrl returns a boolean if a field has been set.

### GetId

`func (o *IamApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamApplication) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpRestriction

`func (o *IamApplication) GetIpRestriction() string`

GetIpRestriction returns the IpRestriction field if non-nil, zero value otherwise.

### GetIpRestrictionOk

`func (o *IamApplication) GetIpRestrictionOk() (*string, bool)`

GetIpRestrictionOk returns a tuple with the IpRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRestriction

`func (o *IamApplication) SetIpRestriction(v string)`

SetIpRestriction sets IpRestriction field to given value.

### HasIpRestriction

`func (o *IamApplication) HasIpRestriction() bool`

HasIpRestriction returns a boolean if a field has been set.

### GetIpWhitelist

`func (o *IamApplication) GetIpWhitelist() string`

GetIpWhitelist returns the IpWhitelist field if non-nil, zero value otherwise.

### GetIpWhitelistOk

`func (o *IamApplication) GetIpWhitelistOk() (*string, bool)`

GetIpWhitelistOk returns a tuple with the IpWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelist

`func (o *IamApplication) SetIpWhitelist(v string)`

SetIpWhitelist sets IpWhitelist field to given value.

### HasIpWhitelist

`func (o *IamApplication) HasIpWhitelist() bool`

HasIpWhitelist returns a boolean if a field has been set.

### GetIsShared

`func (o *IamApplication) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *IamApplication) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *IamApplication) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *IamApplication) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetLogo

`func (o *IamApplication) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *IamApplication) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *IamApplication) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *IamApplication) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *IamApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamApplication) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *IamApplication) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *IamApplication) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *IamApplication) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *IamApplication) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrgChoiceMode

`func (o *IamApplication) GetOrgChoiceMode() string`

GetOrgChoiceMode returns the OrgChoiceMode field if non-nil, zero value otherwise.

### GetOrgChoiceModeOk

`func (o *IamApplication) GetOrgChoiceModeOk() (*string, bool)`

GetOrgChoiceModeOk returns a tuple with the OrgChoiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgChoiceMode

`func (o *IamApplication) SetOrgChoiceMode(v string)`

SetOrgChoiceMode sets OrgChoiceMode field to given value.

### HasOrgChoiceMode

`func (o *IamApplication) HasOrgChoiceMode() bool`

HasOrgChoiceMode returns a boolean if a field has been set.

### GetOrganization

`func (o *IamApplication) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamApplication) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamApplication) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamApplication) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationObj

`func (o *IamApplication) GetOrganizationObj() IamOrganization`

GetOrganizationObj returns the OrganizationObj field if non-nil, zero value otherwise.

### GetOrganizationObjOk

`func (o *IamApplication) GetOrganizationObjOk() (*IamOrganization, bool)`

GetOrganizationObjOk returns a tuple with the OrganizationObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationObj

`func (o *IamApplication) SetOrganizationObj(v IamOrganization)`

SetOrganizationObj sets OrganizationObj field to given value.

### HasOrganizationObj

`func (o *IamApplication) HasOrganizationObj() bool`

HasOrganizationObj returns a boolean if a field has been set.

### GetOtherDomains

`func (o *IamApplication) GetOtherDomains() []string`

GetOtherDomains returns the OtherDomains field if non-nil, zero value otherwise.

### GetOtherDomainsOk

`func (o *IamApplication) GetOtherDomainsOk() (*[]string, bool)`

GetOtherDomainsOk returns a tuple with the OtherDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherDomains

`func (o *IamApplication) SetOtherDomains(v []string)`

SetOtherDomains sets OtherDomains field to given value.

### HasOtherDomains

`func (o *IamApplication) HasOtherDomains() bool`

HasOtherDomains returns a boolean if a field has been set.

### GetOwner

`func (o *IamApplication) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamApplication) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamApplication) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamApplication) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProject

`func (o *IamApplication) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *IamApplication) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *IamApplication) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *IamApplication) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProviders

`func (o *IamApplication) GetProviders() []IamProviderItem`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *IamApplication) GetProvidersOk() (*[]IamProviderItem, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *IamApplication) SetProviders(v []IamProviderItem)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *IamApplication) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetRedirectUris

`func (o *IamApplication) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *IamApplication) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *IamApplication) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *IamApplication) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetRefreshExpireInHours

`func (o *IamApplication) GetRefreshExpireInHours() float64`

GetRefreshExpireInHours returns the RefreshExpireInHours field if non-nil, zero value otherwise.

### GetRefreshExpireInHoursOk

`func (o *IamApplication) GetRefreshExpireInHoursOk() (*float64, bool)`

GetRefreshExpireInHoursOk returns a tuple with the RefreshExpireInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshExpireInHours

`func (o *IamApplication) SetRefreshExpireInHours(v float64)`

SetRefreshExpireInHours sets RefreshExpireInHours field to given value.

### HasRefreshExpireInHours

`func (o *IamApplication) HasRefreshExpireInHours() bool`

HasRefreshExpireInHours returns a boolean if a field has been set.

### GetSamlAttributes

`func (o *IamApplication) GetSamlAttributes() []IamSamlItem`

GetSamlAttributes returns the SamlAttributes field if non-nil, zero value otherwise.

### GetSamlAttributesOk

`func (o *IamApplication) GetSamlAttributesOk() (*[]IamSamlItem, bool)`

GetSamlAttributesOk returns a tuple with the SamlAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributes

`func (o *IamApplication) SetSamlAttributes(v []IamSamlItem)`

SetSamlAttributes sets SamlAttributes field to given value.

### HasSamlAttributes

`func (o *IamApplication) HasSamlAttributes() bool`

HasSamlAttributes returns a boolean if a field has been set.

### GetSamlHashAlgorithm

`func (o *IamApplication) GetSamlHashAlgorithm() string`

GetSamlHashAlgorithm returns the SamlHashAlgorithm field if non-nil, zero value otherwise.

### GetSamlHashAlgorithmOk

`func (o *IamApplication) GetSamlHashAlgorithmOk() (*string, bool)`

GetSamlHashAlgorithmOk returns a tuple with the SamlHashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlHashAlgorithm

`func (o *IamApplication) SetSamlHashAlgorithm(v string)`

SetSamlHashAlgorithm sets SamlHashAlgorithm field to given value.

### HasSamlHashAlgorithm

`func (o *IamApplication) HasSamlHashAlgorithm() bool`

HasSamlHashAlgorithm returns a boolean if a field has been set.

### GetSamlReplyUrl

`func (o *IamApplication) GetSamlReplyUrl() string`

GetSamlReplyUrl returns the SamlReplyUrl field if non-nil, zero value otherwise.

### GetSamlReplyUrlOk

`func (o *IamApplication) GetSamlReplyUrlOk() (*string, bool)`

GetSamlReplyUrlOk returns a tuple with the SamlReplyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlReplyUrl

`func (o *IamApplication) SetSamlReplyUrl(v string)`

SetSamlReplyUrl sets SamlReplyUrl field to given value.

### HasSamlReplyUrl

`func (o *IamApplication) HasSamlReplyUrl() bool`

HasSamlReplyUrl returns a boolean if a field has been set.

### GetScopes

`func (o *IamApplication) GetScopes() []IamScopeItem`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IamApplication) GetScopesOk() (*[]IamScopeItem, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IamApplication) SetScopes(v []IamScopeItem)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *IamApplication) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSigninHtml

`func (o *IamApplication) GetSigninHtml() string`

GetSigninHtml returns the SigninHtml field if non-nil, zero value otherwise.

### GetSigninHtmlOk

`func (o *IamApplication) GetSigninHtmlOk() (*string, bool)`

GetSigninHtmlOk returns a tuple with the SigninHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigninHtml

`func (o *IamApplication) SetSigninHtml(v string)`

SetSigninHtml sets SigninHtml field to given value.

### HasSigninHtml

`func (o *IamApplication) HasSigninHtml() bool`

HasSigninHtml returns a boolean if a field has been set.

### GetSigninItems

`func (o *IamApplication) GetSigninItems() []IamSigninItem`

GetSigninItems returns the SigninItems field if non-nil, zero value otherwise.

### GetSigninItemsOk

`func (o *IamApplication) GetSigninItemsOk() (*[]IamSigninItem, bool)`

GetSigninItemsOk returns a tuple with the SigninItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigninItems

`func (o *IamApplication) SetSigninItems(v []IamSigninItem)`

SetSigninItems sets SigninItems field to given value.

### HasSigninItems

`func (o *IamApplication) HasSigninItems() bool`

HasSigninItems returns a boolean if a field has been set.

### GetSigninMethods

`func (o *IamApplication) GetSigninMethods() []IamSigninMethod`

GetSigninMethods returns the SigninMethods field if non-nil, zero value otherwise.

### GetSigninMethodsOk

`func (o *IamApplication) GetSigninMethodsOk() (*[]IamSigninMethod, bool)`

GetSigninMethodsOk returns a tuple with the SigninMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigninMethods

`func (o *IamApplication) SetSigninMethods(v []IamSigninMethod)`

SetSigninMethods sets SigninMethods field to given value.

### HasSigninMethods

`func (o *IamApplication) HasSigninMethods() bool`

HasSigninMethods returns a boolean if a field has been set.

### GetSigninUrl

`func (o *IamApplication) GetSigninUrl() string`

GetSigninUrl returns the SigninUrl field if non-nil, zero value otherwise.

### GetSigninUrlOk

`func (o *IamApplication) GetSigninUrlOk() (*string, bool)`

GetSigninUrlOk returns a tuple with the SigninUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigninUrl

`func (o *IamApplication) SetSigninUrl(v string)`

SetSigninUrl sets SigninUrl field to given value.

### HasSigninUrl

`func (o *IamApplication) HasSigninUrl() bool`

HasSigninUrl returns a boolean if a field has been set.

### GetSignupHtml

`func (o *IamApplication) GetSignupHtml() string`

GetSignupHtml returns the SignupHtml field if non-nil, zero value otherwise.

### GetSignupHtmlOk

`func (o *IamApplication) GetSignupHtmlOk() (*string, bool)`

GetSignupHtmlOk returns a tuple with the SignupHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupHtml

`func (o *IamApplication) SetSignupHtml(v string)`

SetSignupHtml sets SignupHtml field to given value.

### HasSignupHtml

`func (o *IamApplication) HasSignupHtml() bool`

HasSignupHtml returns a boolean if a field has been set.

### GetSignupItems

`func (o *IamApplication) GetSignupItems() []IamSignupItem`

GetSignupItems returns the SignupItems field if non-nil, zero value otherwise.

### GetSignupItemsOk

`func (o *IamApplication) GetSignupItemsOk() (*[]IamSignupItem, bool)`

GetSignupItemsOk returns a tuple with the SignupItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupItems

`func (o *IamApplication) SetSignupItems(v []IamSignupItem)`

SetSignupItems sets SignupItems field to given value.

### HasSignupItems

`func (o *IamApplication) HasSignupItems() bool`

HasSignupItems returns a boolean if a field has been set.

### GetSignupUrl

`func (o *IamApplication) GetSignupUrl() string`

GetSignupUrl returns the SignupUrl field if non-nil, zero value otherwise.

### GetSignupUrlOk

`func (o *IamApplication) GetSignupUrlOk() (*string, bool)`

GetSignupUrlOk returns a tuple with the SignupUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupUrl

`func (o *IamApplication) SetSignupUrl(v string)`

SetSignupUrl sets SignupUrl field to given value.

### HasSignupUrl

`func (o *IamApplication) HasSignupUrl() bool`

HasSignupUrl returns a boolean if a field has been set.

### GetSslCert

`func (o *IamApplication) GetSslCert() string`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *IamApplication) GetSslCertOk() (*string, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *IamApplication) SetSslCert(v string)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *IamApplication) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### GetSslMode

`func (o *IamApplication) GetSslMode() string`

GetSslMode returns the SslMode field if non-nil, zero value otherwise.

### GetSslModeOk

`func (o *IamApplication) GetSslModeOk() (*string, bool)`

GetSslModeOk returns a tuple with the SslMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslMode

`func (o *IamApplication) SetSslMode(v string)`

SetSslMode sets SslMode field to given value.

### HasSslMode

`func (o *IamApplication) HasSslMode() bool`

HasSslMode returns a boolean if a field has been set.

### GetTags

`func (o *IamApplication) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IamApplication) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IamApplication) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IamApplication) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTermsOfUse

`func (o *IamApplication) GetTermsOfUse() string`

GetTermsOfUse returns the TermsOfUse field if non-nil, zero value otherwise.

### GetTermsOfUseOk

`func (o *IamApplication) GetTermsOfUseOk() (*string, bool)`

GetTermsOfUseOk returns a tuple with the TermsOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfUse

`func (o *IamApplication) SetTermsOfUse(v string)`

SetTermsOfUse sets TermsOfUse field to given value.

### HasTermsOfUse

`func (o *IamApplication) HasTermsOfUse() bool`

HasTermsOfUse returns a boolean if a field has been set.

### GetThemeData

`func (o *IamApplication) GetThemeData() IamThemeData`

GetThemeData returns the ThemeData field if non-nil, zero value otherwise.

### GetThemeDataOk

`func (o *IamApplication) GetThemeDataOk() (*IamThemeData, bool)`

GetThemeDataOk returns a tuple with the ThemeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeData

`func (o *IamApplication) SetThemeData(v IamThemeData)`

SetThemeData sets ThemeData field to given value.

### HasThemeData

`func (o *IamApplication) HasThemeData() bool`

HasThemeData returns a boolean if a field has been set.

### GetTitle

`func (o *IamApplication) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IamApplication) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IamApplication) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IamApplication) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTokenAttributes

`func (o *IamApplication) GetTokenAttributes() []IamJwtItem`

GetTokenAttributes returns the TokenAttributes field if non-nil, zero value otherwise.

### GetTokenAttributesOk

`func (o *IamApplication) GetTokenAttributesOk() (*[]IamJwtItem, bool)`

GetTokenAttributesOk returns a tuple with the TokenAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAttributes

`func (o *IamApplication) SetTokenAttributes(v []IamJwtItem)`

SetTokenAttributes sets TokenAttributes field to given value.

### HasTokenAttributes

`func (o *IamApplication) HasTokenAttributes() bool`

HasTokenAttributes returns a boolean if a field has been set.

### GetTokenFields

`func (o *IamApplication) GetTokenFields() []string`

GetTokenFields returns the TokenFields field if non-nil, zero value otherwise.

### GetTokenFieldsOk

`func (o *IamApplication) GetTokenFieldsOk() (*[]string, bool)`

GetTokenFieldsOk returns a tuple with the TokenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFields

`func (o *IamApplication) SetTokenFields(v []string)`

SetTokenFields sets TokenFields field to given value.

### HasTokenFields

`func (o *IamApplication) HasTokenFields() bool`

HasTokenFields returns a boolean if a field has been set.

### GetTokenFormat

`func (o *IamApplication) GetTokenFormat() string`

GetTokenFormat returns the TokenFormat field if non-nil, zero value otherwise.

### GetTokenFormatOk

`func (o *IamApplication) GetTokenFormatOk() (*string, bool)`

GetTokenFormatOk returns a tuple with the TokenFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFormat

`func (o *IamApplication) SetTokenFormat(v string)`

SetTokenFormat sets TokenFormat field to given value.

### HasTokenFormat

`func (o *IamApplication) HasTokenFormat() bool`

HasTokenFormat returns a boolean if a field has been set.

### GetTokenSigningMethod

`func (o *IamApplication) GetTokenSigningMethod() string`

GetTokenSigningMethod returns the TokenSigningMethod field if non-nil, zero value otherwise.

### GetTokenSigningMethodOk

`func (o *IamApplication) GetTokenSigningMethodOk() (*string, bool)`

GetTokenSigningMethodOk returns a tuple with the TokenSigningMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSigningMethod

`func (o *IamApplication) SetTokenSigningMethod(v string)`

SetTokenSigningMethod sets TokenSigningMethod field to given value.

### HasTokenSigningMethod

`func (o *IamApplication) HasTokenSigningMethod() bool`

HasTokenSigningMethod returns a boolean if a field has been set.

### GetType

`func (o *IamApplication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IamApplication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IamApplication) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IamApplication) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamApplication) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamApplication) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamApplication) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamApplication) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpstreamHost

`func (o *IamApplication) GetUpstreamHost() string`

GetUpstreamHost returns the UpstreamHost field if non-nil, zero value otherwise.

### GetUpstreamHostOk

`func (o *IamApplication) GetUpstreamHostOk() (*string, bool)`

GetUpstreamHostOk returns a tuple with the UpstreamHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamHost

`func (o *IamApplication) SetUpstreamHost(v string)`

SetUpstreamHost sets UpstreamHost field to given value.

### HasUpstreamHost

`func (o *IamApplication) HasUpstreamHost() bool`

HasUpstreamHost returns a boolean if a field has been set.

### GetUseEmailAsSamlNameId

`func (o *IamApplication) GetUseEmailAsSamlNameId() bool`

GetUseEmailAsSamlNameId returns the UseEmailAsSamlNameId field if non-nil, zero value otherwise.

### GetUseEmailAsSamlNameIdOk

`func (o *IamApplication) GetUseEmailAsSamlNameIdOk() (*bool, bool)`

GetUseEmailAsSamlNameIdOk returns a tuple with the UseEmailAsSamlNameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmailAsSamlNameId

`func (o *IamApplication) SetUseEmailAsSamlNameId(v bool)`

SetUseEmailAsSamlNameId sets UseEmailAsSamlNameId field to given value.

### HasUseEmailAsSamlNameId

`func (o *IamApplication) HasUseEmailAsSamlNameId() bool`

HasUseEmailAsSamlNameId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



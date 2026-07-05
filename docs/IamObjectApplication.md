# IamObjectApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffiliationUrl** | Pointer to **string** |  | [optional] 
**Cert** | Pointer to **string** |  | [optional] 
**CertPublicKey** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**CodeResendTimeout** | Pointer to **int64** |  | [optional] 
**CookieExpireInHours** | Pointer to **int64** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**DefaultGroup** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisableSamlAttributes** | Pointer to **bool** |  | [optional] 
**DisableSignin** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
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
**IpRestriction** | Pointer to **string** |  | [optional] 
**IpWhitelist** | Pointer to **string** |  | [optional] 
**IsShared** | Pointer to **bool** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **int64** |  | [optional] 
**OrgChoiceMode** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**OrganizationObj** | Pointer to [**IamObjectOrganization**](IamObjectOrganization.md) |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Providers** | Pointer to [**[]IamObjectProviderItem**](IamObjectProviderItem.md) |  | [optional] 
**RedirectUris** | Pointer to **[]string** |  | [optional] 
**RefreshExpireInHours** | Pointer to **float64** |  | [optional] 
**SamlAttributes** | Pointer to [**[]IamObjectSamlItem**](IamObjectSamlItem.md) |  | [optional] 
**SamlHashAlgorithm** | Pointer to **string** |  | [optional] 
**SamlReplyUrl** | Pointer to **string** |  | [optional] 
**SigninHtml** | Pointer to **string** |  | [optional] 
**SigninItems** | Pointer to [**[]IamObjectSigninItem**](IamObjectSigninItem.md) |  | [optional] 
**SigninMethods** | Pointer to [**[]IamObjectSigninMethod**](IamObjectSigninMethod.md) |  | [optional] 
**SigninUrl** | Pointer to **string** |  | [optional] 
**SignupHtml** | Pointer to **string** |  | [optional] 
**SignupItems** | Pointer to [**[]IamObjectSignupItem**](IamObjectSignupItem.md) |  | [optional] 
**SignupUrl** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TermsOfUse** | Pointer to **string** |  | [optional] 
**ThemeData** | Pointer to [**IamObjectThemeData**](IamObjectThemeData.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TokenAttributes** | Pointer to [**[]IamObjectJwtItem**](IamObjectJwtItem.md) |  | [optional] 
**TokenFields** | Pointer to **[]string** |  | [optional] 
**TokenFormat** | Pointer to **string** |  | [optional] 
**TokenSigningMethod** | Pointer to **string** |  | [optional] 
**UseEmailAsSamlNameId** | Pointer to **bool** |  | [optional] 

## Methods

### NewIamObjectApplication

`func NewIamObjectApplication() *IamObjectApplication`

NewIamObjectApplication instantiates a new IamObjectApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectApplicationWithDefaults

`func NewIamObjectApplicationWithDefaults() *IamObjectApplication`

NewIamObjectApplicationWithDefaults instantiates a new IamObjectApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffiliationUrl

`func (o *IamObjectApplication) GetAffiliationUrl() string`

GetAffiliationUrl returns the AffiliationUrl field if non-nil, zero value otherwise.

### GetAffiliationUrlOk

`func (o *IamObjectApplication) GetAffiliationUrlOk() (*string, bool)`

GetAffiliationUrlOk returns a tuple with the AffiliationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliationUrl

`func (o *IamObjectApplication) SetAffiliationUrl(v string)`

SetAffiliationUrl sets AffiliationUrl field to given value.

### HasAffiliationUrl

`func (o *IamObjectApplication) HasAffiliationUrl() bool`

HasAffiliationUrl returns a boolean if a field has been set.

### GetCert

`func (o *IamObjectApplication) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *IamObjectApplication) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *IamObjectApplication) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *IamObjectApplication) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetCertPublicKey

`func (o *IamObjectApplication) GetCertPublicKey() string`

GetCertPublicKey returns the CertPublicKey field if non-nil, zero value otherwise.

### GetCertPublicKeyOk

`func (o *IamObjectApplication) GetCertPublicKeyOk() (*string, bool)`

GetCertPublicKeyOk returns a tuple with the CertPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertPublicKey

`func (o *IamObjectApplication) SetCertPublicKey(v string)`

SetCertPublicKey sets CertPublicKey field to given value.

### HasCertPublicKey

`func (o *IamObjectApplication) HasCertPublicKey() bool`

HasCertPublicKey returns a boolean if a field has been set.

### GetClientId

`func (o *IamObjectApplication) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IamObjectApplication) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IamObjectApplication) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IamObjectApplication) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *IamObjectApplication) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IamObjectApplication) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IamObjectApplication) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *IamObjectApplication) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetCodeResendTimeout

`func (o *IamObjectApplication) GetCodeResendTimeout() int64`

GetCodeResendTimeout returns the CodeResendTimeout field if non-nil, zero value otherwise.

### GetCodeResendTimeoutOk

`func (o *IamObjectApplication) GetCodeResendTimeoutOk() (*int64, bool)`

GetCodeResendTimeoutOk returns a tuple with the CodeResendTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeResendTimeout

`func (o *IamObjectApplication) SetCodeResendTimeout(v int64)`

SetCodeResendTimeout sets CodeResendTimeout field to given value.

### HasCodeResendTimeout

`func (o *IamObjectApplication) HasCodeResendTimeout() bool`

HasCodeResendTimeout returns a boolean if a field has been set.

### GetCookieExpireInHours

`func (o *IamObjectApplication) GetCookieExpireInHours() int64`

GetCookieExpireInHours returns the CookieExpireInHours field if non-nil, zero value otherwise.

### GetCookieExpireInHoursOk

`func (o *IamObjectApplication) GetCookieExpireInHoursOk() (*int64, bool)`

GetCookieExpireInHoursOk returns a tuple with the CookieExpireInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieExpireInHours

`func (o *IamObjectApplication) SetCookieExpireInHours(v int64)`

SetCookieExpireInHours sets CookieExpireInHours field to given value.

### HasCookieExpireInHours

`func (o *IamObjectApplication) HasCookieExpireInHours() bool`

HasCookieExpireInHours returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamObjectApplication) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectApplication) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectApplication) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectApplication) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDefaultGroup

`func (o *IamObjectApplication) GetDefaultGroup() string`

GetDefaultGroup returns the DefaultGroup field if non-nil, zero value otherwise.

### GetDefaultGroupOk

`func (o *IamObjectApplication) GetDefaultGroupOk() (*string, bool)`

GetDefaultGroupOk returns a tuple with the DefaultGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroup

`func (o *IamObjectApplication) SetDefaultGroup(v string)`

SetDefaultGroup sets DefaultGroup field to given value.

### HasDefaultGroup

`func (o *IamObjectApplication) HasDefaultGroup() bool`

HasDefaultGroup returns a boolean if a field has been set.

### GetDescription

`func (o *IamObjectApplication) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamObjectApplication) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamObjectApplication) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamObjectApplication) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableSamlAttributes

`func (o *IamObjectApplication) GetDisableSamlAttributes() bool`

GetDisableSamlAttributes returns the DisableSamlAttributes field if non-nil, zero value otherwise.

### GetDisableSamlAttributesOk

`func (o *IamObjectApplication) GetDisableSamlAttributesOk() (*bool, bool)`

GetDisableSamlAttributesOk returns a tuple with the DisableSamlAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSamlAttributes

`func (o *IamObjectApplication) SetDisableSamlAttributes(v bool)`

SetDisableSamlAttributes sets DisableSamlAttributes field to given value.

### HasDisableSamlAttributes

`func (o *IamObjectApplication) HasDisableSamlAttributes() bool`

HasDisableSamlAttributes returns a boolean if a field has been set.

### GetDisableSignin

`func (o *IamObjectApplication) GetDisableSignin() bool`

GetDisableSignin returns the DisableSignin field if non-nil, zero value otherwise.

### GetDisableSigninOk

`func (o *IamObjectApplication) GetDisableSigninOk() (*bool, bool)`

GetDisableSigninOk returns a tuple with the DisableSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSignin

`func (o *IamObjectApplication) SetDisableSignin(v bool)`

SetDisableSignin sets DisableSignin field to given value.

### HasDisableSignin

`func (o *IamObjectApplication) HasDisableSignin() bool`

HasDisableSignin returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamObjectApplication) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamObjectApplication) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamObjectApplication) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamObjectApplication) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableAutoSignin

`func (o *IamObjectApplication) GetEnableAutoSignin() bool`

GetEnableAutoSignin returns the EnableAutoSignin field if non-nil, zero value otherwise.

### GetEnableAutoSigninOk

`func (o *IamObjectApplication) GetEnableAutoSigninOk() (*bool, bool)`

GetEnableAutoSigninOk returns a tuple with the EnableAutoSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoSignin

`func (o *IamObjectApplication) SetEnableAutoSignin(v bool)`

SetEnableAutoSignin sets EnableAutoSignin field to given value.

### HasEnableAutoSignin

`func (o *IamObjectApplication) HasEnableAutoSignin() bool`

HasEnableAutoSignin returns a boolean if a field has been set.

### GetEnableCodeSignin

`func (o *IamObjectApplication) GetEnableCodeSignin() bool`

GetEnableCodeSignin returns the EnableCodeSignin field if non-nil, zero value otherwise.

### GetEnableCodeSigninOk

`func (o *IamObjectApplication) GetEnableCodeSigninOk() (*bool, bool)`

GetEnableCodeSigninOk returns a tuple with the EnableCodeSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCodeSignin

`func (o *IamObjectApplication) SetEnableCodeSignin(v bool)`

SetEnableCodeSignin sets EnableCodeSignin field to given value.

### HasEnableCodeSignin

`func (o *IamObjectApplication) HasEnableCodeSignin() bool`

HasEnableCodeSignin returns a boolean if a field has been set.

### GetEnableExclusiveSignin

`func (o *IamObjectApplication) GetEnableExclusiveSignin() bool`

GetEnableExclusiveSignin returns the EnableExclusiveSignin field if non-nil, zero value otherwise.

### GetEnableExclusiveSigninOk

`func (o *IamObjectApplication) GetEnableExclusiveSigninOk() (*bool, bool)`

GetEnableExclusiveSigninOk returns a tuple with the EnableExclusiveSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExclusiveSignin

`func (o *IamObjectApplication) SetEnableExclusiveSignin(v bool)`

SetEnableExclusiveSignin sets EnableExclusiveSignin field to given value.

### HasEnableExclusiveSignin

`func (o *IamObjectApplication) HasEnableExclusiveSignin() bool`

HasEnableExclusiveSignin returns a boolean if a field has been set.

### GetEnableLinkWithEmail

`func (o *IamObjectApplication) GetEnableLinkWithEmail() bool`

GetEnableLinkWithEmail returns the EnableLinkWithEmail field if non-nil, zero value otherwise.

### GetEnableLinkWithEmailOk

`func (o *IamObjectApplication) GetEnableLinkWithEmailOk() (*bool, bool)`

GetEnableLinkWithEmailOk returns a tuple with the EnableLinkWithEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLinkWithEmail

`func (o *IamObjectApplication) SetEnableLinkWithEmail(v bool)`

SetEnableLinkWithEmail sets EnableLinkWithEmail field to given value.

### HasEnableLinkWithEmail

`func (o *IamObjectApplication) HasEnableLinkWithEmail() bool`

HasEnableLinkWithEmail returns a boolean if a field has been set.

### GetEnablePassword

`func (o *IamObjectApplication) GetEnablePassword() bool`

GetEnablePassword returns the EnablePassword field if non-nil, zero value otherwise.

### GetEnablePasswordOk

`func (o *IamObjectApplication) GetEnablePasswordOk() (*bool, bool)`

GetEnablePasswordOk returns a tuple with the EnablePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePassword

`func (o *IamObjectApplication) SetEnablePassword(v bool)`

SetEnablePassword sets EnablePassword field to given value.

### HasEnablePassword

`func (o *IamObjectApplication) HasEnablePassword() bool`

HasEnablePassword returns a boolean if a field has been set.

### GetEnableSamlAssertionSignature

`func (o *IamObjectApplication) GetEnableSamlAssertionSignature() bool`

GetEnableSamlAssertionSignature returns the EnableSamlAssertionSignature field if non-nil, zero value otherwise.

### GetEnableSamlAssertionSignatureOk

`func (o *IamObjectApplication) GetEnableSamlAssertionSignatureOk() (*bool, bool)`

GetEnableSamlAssertionSignatureOk returns a tuple with the EnableSamlAssertionSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSamlAssertionSignature

`func (o *IamObjectApplication) SetEnableSamlAssertionSignature(v bool)`

SetEnableSamlAssertionSignature sets EnableSamlAssertionSignature field to given value.

### HasEnableSamlAssertionSignature

`func (o *IamObjectApplication) HasEnableSamlAssertionSignature() bool`

HasEnableSamlAssertionSignature returns a boolean if a field has been set.

### GetEnableSamlC14n10

`func (o *IamObjectApplication) GetEnableSamlC14n10() bool`

GetEnableSamlC14n10 returns the EnableSamlC14n10 field if non-nil, zero value otherwise.

### GetEnableSamlC14n10Ok

`func (o *IamObjectApplication) GetEnableSamlC14n10Ok() (*bool, bool)`

GetEnableSamlC14n10Ok returns a tuple with the EnableSamlC14n10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSamlC14n10

`func (o *IamObjectApplication) SetEnableSamlC14n10(v bool)`

SetEnableSamlC14n10 sets EnableSamlC14n10 field to given value.

### HasEnableSamlC14n10

`func (o *IamObjectApplication) HasEnableSamlC14n10() bool`

HasEnableSamlC14n10 returns a boolean if a field has been set.

### GetEnableSamlCompress

`func (o *IamObjectApplication) GetEnableSamlCompress() bool`

GetEnableSamlCompress returns the EnableSamlCompress field if non-nil, zero value otherwise.

### GetEnableSamlCompressOk

`func (o *IamObjectApplication) GetEnableSamlCompressOk() (*bool, bool)`

GetEnableSamlCompressOk returns a tuple with the EnableSamlCompress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSamlCompress

`func (o *IamObjectApplication) SetEnableSamlCompress(v bool)`

SetEnableSamlCompress sets EnableSamlCompress field to given value.

### HasEnableSamlCompress

`func (o *IamObjectApplication) HasEnableSamlCompress() bool`

HasEnableSamlCompress returns a boolean if a field has been set.

### GetEnableSamlPostBinding

`func (o *IamObjectApplication) GetEnableSamlPostBinding() bool`

GetEnableSamlPostBinding returns the EnableSamlPostBinding field if non-nil, zero value otherwise.

### GetEnableSamlPostBindingOk

`func (o *IamObjectApplication) GetEnableSamlPostBindingOk() (*bool, bool)`

GetEnableSamlPostBindingOk returns a tuple with the EnableSamlPostBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSamlPostBinding

`func (o *IamObjectApplication) SetEnableSamlPostBinding(v bool)`

SetEnableSamlPostBinding sets EnableSamlPostBinding field to given value.

### HasEnableSamlPostBinding

`func (o *IamObjectApplication) HasEnableSamlPostBinding() bool`

HasEnableSamlPostBinding returns a boolean if a field has been set.

### GetEnableSignUp

`func (o *IamObjectApplication) GetEnableSignUp() bool`

GetEnableSignUp returns the EnableSignUp field if non-nil, zero value otherwise.

### GetEnableSignUpOk

`func (o *IamObjectApplication) GetEnableSignUpOk() (*bool, bool)`

GetEnableSignUpOk returns a tuple with the EnableSignUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSignUp

`func (o *IamObjectApplication) SetEnableSignUp(v bool)`

SetEnableSignUp sets EnableSignUp field to given value.

### HasEnableSignUp

`func (o *IamObjectApplication) HasEnableSignUp() bool`

HasEnableSignUp returns a boolean if a field has been set.

### GetEnableSigninSession

`func (o *IamObjectApplication) GetEnableSigninSession() bool`

GetEnableSigninSession returns the EnableSigninSession field if non-nil, zero value otherwise.

### GetEnableSigninSessionOk

`func (o *IamObjectApplication) GetEnableSigninSessionOk() (*bool, bool)`

GetEnableSigninSessionOk returns a tuple with the EnableSigninSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSigninSession

`func (o *IamObjectApplication) SetEnableSigninSession(v bool)`

SetEnableSigninSession sets EnableSigninSession field to given value.

### HasEnableSigninSession

`func (o *IamObjectApplication) HasEnableSigninSession() bool`

HasEnableSigninSession returns a boolean if a field has been set.

### GetEnableWebAuthn

`func (o *IamObjectApplication) GetEnableWebAuthn() bool`

GetEnableWebAuthn returns the EnableWebAuthn field if non-nil, zero value otherwise.

### GetEnableWebAuthnOk

`func (o *IamObjectApplication) GetEnableWebAuthnOk() (*bool, bool)`

GetEnableWebAuthnOk returns a tuple with the EnableWebAuthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWebAuthn

`func (o *IamObjectApplication) SetEnableWebAuthn(v bool)`

SetEnableWebAuthn sets EnableWebAuthn field to given value.

### HasEnableWebAuthn

`func (o *IamObjectApplication) HasEnableWebAuthn() bool`

HasEnableWebAuthn returns a boolean if a field has been set.

### GetExpireInHours

`func (o *IamObjectApplication) GetExpireInHours() float64`

GetExpireInHours returns the ExpireInHours field if non-nil, zero value otherwise.

### GetExpireInHoursOk

`func (o *IamObjectApplication) GetExpireInHoursOk() (*float64, bool)`

GetExpireInHoursOk returns a tuple with the ExpireInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireInHours

`func (o *IamObjectApplication) SetExpireInHours(v float64)`

SetExpireInHours sets ExpireInHours field to given value.

### HasExpireInHours

`func (o *IamObjectApplication) HasExpireInHours() bool`

HasExpireInHours returns a boolean if a field has been set.

### GetFailedSigninFrozenTime

`func (o *IamObjectApplication) GetFailedSigninFrozenTime() int64`

GetFailedSigninFrozenTime returns the FailedSigninFrozenTime field if non-nil, zero value otherwise.

### GetFailedSigninFrozenTimeOk

`func (o *IamObjectApplication) GetFailedSigninFrozenTimeOk() (*int64, bool)`

GetFailedSigninFrozenTimeOk returns a tuple with the FailedSigninFrozenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedSigninFrozenTime

`func (o *IamObjectApplication) SetFailedSigninFrozenTime(v int64)`

SetFailedSigninFrozenTime sets FailedSigninFrozenTime field to given value.

### HasFailedSigninFrozenTime

`func (o *IamObjectApplication) HasFailedSigninFrozenTime() bool`

HasFailedSigninFrozenTime returns a boolean if a field has been set.

### GetFailedSigninLimit

`func (o *IamObjectApplication) GetFailedSigninLimit() int64`

GetFailedSigninLimit returns the FailedSigninLimit field if non-nil, zero value otherwise.

### GetFailedSigninLimitOk

`func (o *IamObjectApplication) GetFailedSigninLimitOk() (*int64, bool)`

GetFailedSigninLimitOk returns a tuple with the FailedSigninLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedSigninLimit

`func (o *IamObjectApplication) SetFailedSigninLimit(v int64)`

SetFailedSigninLimit sets FailedSigninLimit field to given value.

### HasFailedSigninLimit

`func (o *IamObjectApplication) HasFailedSigninLimit() bool`

HasFailedSigninLimit returns a boolean if a field has been set.

### GetFavicon

`func (o *IamObjectApplication) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *IamObjectApplication) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *IamObjectApplication) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *IamObjectApplication) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### GetFooterHtml

`func (o *IamObjectApplication) GetFooterHtml() string`

GetFooterHtml returns the FooterHtml field if non-nil, zero value otherwise.

### GetFooterHtmlOk

`func (o *IamObjectApplication) GetFooterHtmlOk() (*string, bool)`

GetFooterHtmlOk returns a tuple with the FooterHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterHtml

`func (o *IamObjectApplication) SetFooterHtml(v string)`

SetFooterHtml sets FooterHtml field to given value.

### HasFooterHtml

`func (o *IamObjectApplication) HasFooterHtml() bool`

HasFooterHtml returns a boolean if a field has been set.

### GetForcedRedirectOrigin

`func (o *IamObjectApplication) GetForcedRedirectOrigin() string`

GetForcedRedirectOrigin returns the ForcedRedirectOrigin field if non-nil, zero value otherwise.

### GetForcedRedirectOriginOk

`func (o *IamObjectApplication) GetForcedRedirectOriginOk() (*string, bool)`

GetForcedRedirectOriginOk returns a tuple with the ForcedRedirectOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedRedirectOrigin

`func (o *IamObjectApplication) SetForcedRedirectOrigin(v string)`

SetForcedRedirectOrigin sets ForcedRedirectOrigin field to given value.

### HasForcedRedirectOrigin

`func (o *IamObjectApplication) HasForcedRedirectOrigin() bool`

HasForcedRedirectOrigin returns a boolean if a field has been set.

### GetForgetUrl

`func (o *IamObjectApplication) GetForgetUrl() string`

GetForgetUrl returns the ForgetUrl field if non-nil, zero value otherwise.

### GetForgetUrlOk

`func (o *IamObjectApplication) GetForgetUrlOk() (*string, bool)`

GetForgetUrlOk returns a tuple with the ForgetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgetUrl

`func (o *IamObjectApplication) SetForgetUrl(v string)`

SetForgetUrl sets ForgetUrl field to given value.

### HasForgetUrl

`func (o *IamObjectApplication) HasForgetUrl() bool`

HasForgetUrl returns a boolean if a field has been set.

### GetFormBackgroundUrl

`func (o *IamObjectApplication) GetFormBackgroundUrl() string`

GetFormBackgroundUrl returns the FormBackgroundUrl field if non-nil, zero value otherwise.

### GetFormBackgroundUrlOk

`func (o *IamObjectApplication) GetFormBackgroundUrlOk() (*string, bool)`

GetFormBackgroundUrlOk returns a tuple with the FormBackgroundUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormBackgroundUrl

`func (o *IamObjectApplication) SetFormBackgroundUrl(v string)`

SetFormBackgroundUrl sets FormBackgroundUrl field to given value.

### HasFormBackgroundUrl

`func (o *IamObjectApplication) HasFormBackgroundUrl() bool`

HasFormBackgroundUrl returns a boolean if a field has been set.

### GetFormBackgroundUrlMobile

`func (o *IamObjectApplication) GetFormBackgroundUrlMobile() string`

GetFormBackgroundUrlMobile returns the FormBackgroundUrlMobile field if non-nil, zero value otherwise.

### GetFormBackgroundUrlMobileOk

`func (o *IamObjectApplication) GetFormBackgroundUrlMobileOk() (*string, bool)`

GetFormBackgroundUrlMobileOk returns a tuple with the FormBackgroundUrlMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormBackgroundUrlMobile

`func (o *IamObjectApplication) SetFormBackgroundUrlMobile(v string)`

SetFormBackgroundUrlMobile sets FormBackgroundUrlMobile field to given value.

### HasFormBackgroundUrlMobile

`func (o *IamObjectApplication) HasFormBackgroundUrlMobile() bool`

HasFormBackgroundUrlMobile returns a boolean if a field has been set.

### GetFormCss

`func (o *IamObjectApplication) GetFormCss() string`

GetFormCss returns the FormCss field if non-nil, zero value otherwise.

### GetFormCssOk

`func (o *IamObjectApplication) GetFormCssOk() (*string, bool)`

GetFormCssOk returns a tuple with the FormCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormCss

`func (o *IamObjectApplication) SetFormCss(v string)`

SetFormCss sets FormCss field to given value.

### HasFormCss

`func (o *IamObjectApplication) HasFormCss() bool`

HasFormCss returns a boolean if a field has been set.

### GetFormCssMobile

`func (o *IamObjectApplication) GetFormCssMobile() string`

GetFormCssMobile returns the FormCssMobile field if non-nil, zero value otherwise.

### GetFormCssMobileOk

`func (o *IamObjectApplication) GetFormCssMobileOk() (*string, bool)`

GetFormCssMobileOk returns a tuple with the FormCssMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormCssMobile

`func (o *IamObjectApplication) SetFormCssMobile(v string)`

SetFormCssMobile sets FormCssMobile field to given value.

### HasFormCssMobile

`func (o *IamObjectApplication) HasFormCssMobile() bool`

HasFormCssMobile returns a boolean if a field has been set.

### GetFormOffset

`func (o *IamObjectApplication) GetFormOffset() int64`

GetFormOffset returns the FormOffset field if non-nil, zero value otherwise.

### GetFormOffsetOk

`func (o *IamObjectApplication) GetFormOffsetOk() (*int64, bool)`

GetFormOffsetOk returns a tuple with the FormOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormOffset

`func (o *IamObjectApplication) SetFormOffset(v int64)`

SetFormOffset sets FormOffset field to given value.

### HasFormOffset

`func (o *IamObjectApplication) HasFormOffset() bool`

HasFormOffset returns a boolean if a field has been set.

### GetFormSideHtml

`func (o *IamObjectApplication) GetFormSideHtml() string`

GetFormSideHtml returns the FormSideHtml field if non-nil, zero value otherwise.

### GetFormSideHtmlOk

`func (o *IamObjectApplication) GetFormSideHtmlOk() (*string, bool)`

GetFormSideHtmlOk returns a tuple with the FormSideHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormSideHtml

`func (o *IamObjectApplication) SetFormSideHtml(v string)`

SetFormSideHtml sets FormSideHtml field to given value.

### HasFormSideHtml

`func (o *IamObjectApplication) HasFormSideHtml() bool`

HasFormSideHtml returns a boolean if a field has been set.

### GetGrantTypes

`func (o *IamObjectApplication) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *IamObjectApplication) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *IamObjectApplication) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *IamObjectApplication) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetHeaderHtml

`func (o *IamObjectApplication) GetHeaderHtml() string`

GetHeaderHtml returns the HeaderHtml field if non-nil, zero value otherwise.

### GetHeaderHtmlOk

`func (o *IamObjectApplication) GetHeaderHtmlOk() (*string, bool)`

GetHeaderHtmlOk returns a tuple with the HeaderHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderHtml

`func (o *IamObjectApplication) SetHeaderHtml(v string)`

SetHeaderHtml sets HeaderHtml field to given value.

### HasHeaderHtml

`func (o *IamObjectApplication) HasHeaderHtml() bool`

HasHeaderHtml returns a boolean if a field has been set.

### GetHomepageUrl

`func (o *IamObjectApplication) GetHomepageUrl() string`

GetHomepageUrl returns the HomepageUrl field if non-nil, zero value otherwise.

### GetHomepageUrlOk

`func (o *IamObjectApplication) GetHomepageUrlOk() (*string, bool)`

GetHomepageUrlOk returns a tuple with the HomepageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepageUrl

`func (o *IamObjectApplication) SetHomepageUrl(v string)`

SetHomepageUrl sets HomepageUrl field to given value.

### HasHomepageUrl

`func (o *IamObjectApplication) HasHomepageUrl() bool`

HasHomepageUrl returns a boolean if a field has been set.

### GetIpRestriction

`func (o *IamObjectApplication) GetIpRestriction() string`

GetIpRestriction returns the IpRestriction field if non-nil, zero value otherwise.

### GetIpRestrictionOk

`func (o *IamObjectApplication) GetIpRestrictionOk() (*string, bool)`

GetIpRestrictionOk returns a tuple with the IpRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRestriction

`func (o *IamObjectApplication) SetIpRestriction(v string)`

SetIpRestriction sets IpRestriction field to given value.

### HasIpRestriction

`func (o *IamObjectApplication) HasIpRestriction() bool`

HasIpRestriction returns a boolean if a field has been set.

### GetIpWhitelist

`func (o *IamObjectApplication) GetIpWhitelist() string`

GetIpWhitelist returns the IpWhitelist field if non-nil, zero value otherwise.

### GetIpWhitelistOk

`func (o *IamObjectApplication) GetIpWhitelistOk() (*string, bool)`

GetIpWhitelistOk returns a tuple with the IpWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelist

`func (o *IamObjectApplication) SetIpWhitelist(v string)`

SetIpWhitelist sets IpWhitelist field to given value.

### HasIpWhitelist

`func (o *IamObjectApplication) HasIpWhitelist() bool`

HasIpWhitelist returns a boolean if a field has been set.

### GetIsShared

`func (o *IamObjectApplication) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *IamObjectApplication) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *IamObjectApplication) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *IamObjectApplication) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetLogo

`func (o *IamObjectApplication) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *IamObjectApplication) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *IamObjectApplication) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *IamObjectApplication) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *IamObjectApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectApplication) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *IamObjectApplication) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *IamObjectApplication) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *IamObjectApplication) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *IamObjectApplication) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrgChoiceMode

`func (o *IamObjectApplication) GetOrgChoiceMode() string`

GetOrgChoiceMode returns the OrgChoiceMode field if non-nil, zero value otherwise.

### GetOrgChoiceModeOk

`func (o *IamObjectApplication) GetOrgChoiceModeOk() (*string, bool)`

GetOrgChoiceModeOk returns a tuple with the OrgChoiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgChoiceMode

`func (o *IamObjectApplication) SetOrgChoiceMode(v string)`

SetOrgChoiceMode sets OrgChoiceMode field to given value.

### HasOrgChoiceMode

`func (o *IamObjectApplication) HasOrgChoiceMode() bool`

HasOrgChoiceMode returns a boolean if a field has been set.

### GetOrganization

`func (o *IamObjectApplication) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamObjectApplication) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamObjectApplication) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamObjectApplication) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationObj

`func (o *IamObjectApplication) GetOrganizationObj() IamObjectOrganization`

GetOrganizationObj returns the OrganizationObj field if non-nil, zero value otherwise.

### GetOrganizationObjOk

`func (o *IamObjectApplication) GetOrganizationObjOk() (*IamObjectOrganization, bool)`

GetOrganizationObjOk returns a tuple with the OrganizationObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationObj

`func (o *IamObjectApplication) SetOrganizationObj(v IamObjectOrganization)`

SetOrganizationObj sets OrganizationObj field to given value.

### HasOrganizationObj

`func (o *IamObjectApplication) HasOrganizationObj() bool`

HasOrganizationObj returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectApplication) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectApplication) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectApplication) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectApplication) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProviders

`func (o *IamObjectApplication) GetProviders() []IamObjectProviderItem`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *IamObjectApplication) GetProvidersOk() (*[]IamObjectProviderItem, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *IamObjectApplication) SetProviders(v []IamObjectProviderItem)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *IamObjectApplication) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetRedirectUris

`func (o *IamObjectApplication) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *IamObjectApplication) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *IamObjectApplication) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *IamObjectApplication) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetRefreshExpireInHours

`func (o *IamObjectApplication) GetRefreshExpireInHours() float64`

GetRefreshExpireInHours returns the RefreshExpireInHours field if non-nil, zero value otherwise.

### GetRefreshExpireInHoursOk

`func (o *IamObjectApplication) GetRefreshExpireInHoursOk() (*float64, bool)`

GetRefreshExpireInHoursOk returns a tuple with the RefreshExpireInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshExpireInHours

`func (o *IamObjectApplication) SetRefreshExpireInHours(v float64)`

SetRefreshExpireInHours sets RefreshExpireInHours field to given value.

### HasRefreshExpireInHours

`func (o *IamObjectApplication) HasRefreshExpireInHours() bool`

HasRefreshExpireInHours returns a boolean if a field has been set.

### GetSamlAttributes

`func (o *IamObjectApplication) GetSamlAttributes() []IamObjectSamlItem`

GetSamlAttributes returns the SamlAttributes field if non-nil, zero value otherwise.

### GetSamlAttributesOk

`func (o *IamObjectApplication) GetSamlAttributesOk() (*[]IamObjectSamlItem, bool)`

GetSamlAttributesOk returns a tuple with the SamlAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributes

`func (o *IamObjectApplication) SetSamlAttributes(v []IamObjectSamlItem)`

SetSamlAttributes sets SamlAttributes field to given value.

### HasSamlAttributes

`func (o *IamObjectApplication) HasSamlAttributes() bool`

HasSamlAttributes returns a boolean if a field has been set.

### GetSamlHashAlgorithm

`func (o *IamObjectApplication) GetSamlHashAlgorithm() string`

GetSamlHashAlgorithm returns the SamlHashAlgorithm field if non-nil, zero value otherwise.

### GetSamlHashAlgorithmOk

`func (o *IamObjectApplication) GetSamlHashAlgorithmOk() (*string, bool)`

GetSamlHashAlgorithmOk returns a tuple with the SamlHashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlHashAlgorithm

`func (o *IamObjectApplication) SetSamlHashAlgorithm(v string)`

SetSamlHashAlgorithm sets SamlHashAlgorithm field to given value.

### HasSamlHashAlgorithm

`func (o *IamObjectApplication) HasSamlHashAlgorithm() bool`

HasSamlHashAlgorithm returns a boolean if a field has been set.

### GetSamlReplyUrl

`func (o *IamObjectApplication) GetSamlReplyUrl() string`

GetSamlReplyUrl returns the SamlReplyUrl field if non-nil, zero value otherwise.

### GetSamlReplyUrlOk

`func (o *IamObjectApplication) GetSamlReplyUrlOk() (*string, bool)`

GetSamlReplyUrlOk returns a tuple with the SamlReplyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlReplyUrl

`func (o *IamObjectApplication) SetSamlReplyUrl(v string)`

SetSamlReplyUrl sets SamlReplyUrl field to given value.

### HasSamlReplyUrl

`func (o *IamObjectApplication) HasSamlReplyUrl() bool`

HasSamlReplyUrl returns a boolean if a field has been set.

### GetSigninHtml

`func (o *IamObjectApplication) GetSigninHtml() string`

GetSigninHtml returns the SigninHtml field if non-nil, zero value otherwise.

### GetSigninHtmlOk

`func (o *IamObjectApplication) GetSigninHtmlOk() (*string, bool)`

GetSigninHtmlOk returns a tuple with the SigninHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigninHtml

`func (o *IamObjectApplication) SetSigninHtml(v string)`

SetSigninHtml sets SigninHtml field to given value.

### HasSigninHtml

`func (o *IamObjectApplication) HasSigninHtml() bool`

HasSigninHtml returns a boolean if a field has been set.

### GetSigninItems

`func (o *IamObjectApplication) GetSigninItems() []IamObjectSigninItem`

GetSigninItems returns the SigninItems field if non-nil, zero value otherwise.

### GetSigninItemsOk

`func (o *IamObjectApplication) GetSigninItemsOk() (*[]IamObjectSigninItem, bool)`

GetSigninItemsOk returns a tuple with the SigninItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigninItems

`func (o *IamObjectApplication) SetSigninItems(v []IamObjectSigninItem)`

SetSigninItems sets SigninItems field to given value.

### HasSigninItems

`func (o *IamObjectApplication) HasSigninItems() bool`

HasSigninItems returns a boolean if a field has been set.

### GetSigninMethods

`func (o *IamObjectApplication) GetSigninMethods() []IamObjectSigninMethod`

GetSigninMethods returns the SigninMethods field if non-nil, zero value otherwise.

### GetSigninMethodsOk

`func (o *IamObjectApplication) GetSigninMethodsOk() (*[]IamObjectSigninMethod, bool)`

GetSigninMethodsOk returns a tuple with the SigninMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigninMethods

`func (o *IamObjectApplication) SetSigninMethods(v []IamObjectSigninMethod)`

SetSigninMethods sets SigninMethods field to given value.

### HasSigninMethods

`func (o *IamObjectApplication) HasSigninMethods() bool`

HasSigninMethods returns a boolean if a field has been set.

### GetSigninUrl

`func (o *IamObjectApplication) GetSigninUrl() string`

GetSigninUrl returns the SigninUrl field if non-nil, zero value otherwise.

### GetSigninUrlOk

`func (o *IamObjectApplication) GetSigninUrlOk() (*string, bool)`

GetSigninUrlOk returns a tuple with the SigninUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigninUrl

`func (o *IamObjectApplication) SetSigninUrl(v string)`

SetSigninUrl sets SigninUrl field to given value.

### HasSigninUrl

`func (o *IamObjectApplication) HasSigninUrl() bool`

HasSigninUrl returns a boolean if a field has been set.

### GetSignupHtml

`func (o *IamObjectApplication) GetSignupHtml() string`

GetSignupHtml returns the SignupHtml field if non-nil, zero value otherwise.

### GetSignupHtmlOk

`func (o *IamObjectApplication) GetSignupHtmlOk() (*string, bool)`

GetSignupHtmlOk returns a tuple with the SignupHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupHtml

`func (o *IamObjectApplication) SetSignupHtml(v string)`

SetSignupHtml sets SignupHtml field to given value.

### HasSignupHtml

`func (o *IamObjectApplication) HasSignupHtml() bool`

HasSignupHtml returns a boolean if a field has been set.

### GetSignupItems

`func (o *IamObjectApplication) GetSignupItems() []IamObjectSignupItem`

GetSignupItems returns the SignupItems field if non-nil, zero value otherwise.

### GetSignupItemsOk

`func (o *IamObjectApplication) GetSignupItemsOk() (*[]IamObjectSignupItem, bool)`

GetSignupItemsOk returns a tuple with the SignupItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupItems

`func (o *IamObjectApplication) SetSignupItems(v []IamObjectSignupItem)`

SetSignupItems sets SignupItems field to given value.

### HasSignupItems

`func (o *IamObjectApplication) HasSignupItems() bool`

HasSignupItems returns a boolean if a field has been set.

### GetSignupUrl

`func (o *IamObjectApplication) GetSignupUrl() string`

GetSignupUrl returns the SignupUrl field if non-nil, zero value otherwise.

### GetSignupUrlOk

`func (o *IamObjectApplication) GetSignupUrlOk() (*string, bool)`

GetSignupUrlOk returns a tuple with the SignupUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupUrl

`func (o *IamObjectApplication) SetSignupUrl(v string)`

SetSignupUrl sets SignupUrl field to given value.

### HasSignupUrl

`func (o *IamObjectApplication) HasSignupUrl() bool`

HasSignupUrl returns a boolean if a field has been set.

### GetTags

`func (o *IamObjectApplication) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IamObjectApplication) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IamObjectApplication) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IamObjectApplication) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTermsOfUse

`func (o *IamObjectApplication) GetTermsOfUse() string`

GetTermsOfUse returns the TermsOfUse field if non-nil, zero value otherwise.

### GetTermsOfUseOk

`func (o *IamObjectApplication) GetTermsOfUseOk() (*string, bool)`

GetTermsOfUseOk returns a tuple with the TermsOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfUse

`func (o *IamObjectApplication) SetTermsOfUse(v string)`

SetTermsOfUse sets TermsOfUse field to given value.

### HasTermsOfUse

`func (o *IamObjectApplication) HasTermsOfUse() bool`

HasTermsOfUse returns a boolean if a field has been set.

### GetThemeData

`func (o *IamObjectApplication) GetThemeData() IamObjectThemeData`

GetThemeData returns the ThemeData field if non-nil, zero value otherwise.

### GetThemeDataOk

`func (o *IamObjectApplication) GetThemeDataOk() (*IamObjectThemeData, bool)`

GetThemeDataOk returns a tuple with the ThemeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeData

`func (o *IamObjectApplication) SetThemeData(v IamObjectThemeData)`

SetThemeData sets ThemeData field to given value.

### HasThemeData

`func (o *IamObjectApplication) HasThemeData() bool`

HasThemeData returns a boolean if a field has been set.

### GetTitle

`func (o *IamObjectApplication) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IamObjectApplication) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IamObjectApplication) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IamObjectApplication) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTokenAttributes

`func (o *IamObjectApplication) GetTokenAttributes() []IamObjectJwtItem`

GetTokenAttributes returns the TokenAttributes field if non-nil, zero value otherwise.

### GetTokenAttributesOk

`func (o *IamObjectApplication) GetTokenAttributesOk() (*[]IamObjectJwtItem, bool)`

GetTokenAttributesOk returns a tuple with the TokenAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAttributes

`func (o *IamObjectApplication) SetTokenAttributes(v []IamObjectJwtItem)`

SetTokenAttributes sets TokenAttributes field to given value.

### HasTokenAttributes

`func (o *IamObjectApplication) HasTokenAttributes() bool`

HasTokenAttributes returns a boolean if a field has been set.

### GetTokenFields

`func (o *IamObjectApplication) GetTokenFields() []string`

GetTokenFields returns the TokenFields field if non-nil, zero value otherwise.

### GetTokenFieldsOk

`func (o *IamObjectApplication) GetTokenFieldsOk() (*[]string, bool)`

GetTokenFieldsOk returns a tuple with the TokenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFields

`func (o *IamObjectApplication) SetTokenFields(v []string)`

SetTokenFields sets TokenFields field to given value.

### HasTokenFields

`func (o *IamObjectApplication) HasTokenFields() bool`

HasTokenFields returns a boolean if a field has been set.

### GetTokenFormat

`func (o *IamObjectApplication) GetTokenFormat() string`

GetTokenFormat returns the TokenFormat field if non-nil, zero value otherwise.

### GetTokenFormatOk

`func (o *IamObjectApplication) GetTokenFormatOk() (*string, bool)`

GetTokenFormatOk returns a tuple with the TokenFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFormat

`func (o *IamObjectApplication) SetTokenFormat(v string)`

SetTokenFormat sets TokenFormat field to given value.

### HasTokenFormat

`func (o *IamObjectApplication) HasTokenFormat() bool`

HasTokenFormat returns a boolean if a field has been set.

### GetTokenSigningMethod

`func (o *IamObjectApplication) GetTokenSigningMethod() string`

GetTokenSigningMethod returns the TokenSigningMethod field if non-nil, zero value otherwise.

### GetTokenSigningMethodOk

`func (o *IamObjectApplication) GetTokenSigningMethodOk() (*string, bool)`

GetTokenSigningMethodOk returns a tuple with the TokenSigningMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSigningMethod

`func (o *IamObjectApplication) SetTokenSigningMethod(v string)`

SetTokenSigningMethod sets TokenSigningMethod field to given value.

### HasTokenSigningMethod

`func (o *IamObjectApplication) HasTokenSigningMethod() bool`

HasTokenSigningMethod returns a boolean if a field has been set.

### GetUseEmailAsSamlNameId

`func (o *IamObjectApplication) GetUseEmailAsSamlNameId() bool`

GetUseEmailAsSamlNameId returns the UseEmailAsSamlNameId field if non-nil, zero value otherwise.

### GetUseEmailAsSamlNameIdOk

`func (o *IamObjectApplication) GetUseEmailAsSamlNameIdOk() (*bool, bool)`

GetUseEmailAsSamlNameIdOk returns a tuple with the UseEmailAsSamlNameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmailAsSamlNameId

`func (o *IamObjectApplication) SetUseEmailAsSamlNameId(v bool)`

SetUseEmailAsSamlNameId sets UseEmailAsSamlNameId field to given value.

### HasUseEmailAsSamlNameId

`func (o *IamObjectApplication) HasUseEmailAsSamlNameId() bool`

HasUseEmailAsSamlNameId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



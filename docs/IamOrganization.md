# IamOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItems** | Pointer to [**[]IamAccountItem**](IamAccountItem.md) |  | [optional] 
**AccountMenu** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **string** | How the organization appears across Hanzo — the square mark beside its name — as an image or as one emoji, never both. It is the pair a person carries (User.Avatar) under the same names, resolved the same way, so a screen draws a subject without asking which kind of subject it has. Both halves live on the row: a mark that appears everywhere cannot be kept on one device. Written through schema.MarkOf; Logo and LogoDark above are a different thing, the wordmark a login screen draws. | [optional] 
**BalanceCredit** | Pointer to **float32** |  | [optional] 
**BalanceCurrency** | Pointer to **string** |  | [optional] 
**CountryCodes** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**DcrPolicy** | Pointer to **string** |  | [optional] 
**DefaultApplication** | Pointer to **string** |  | [optional] 
**DefaultAvatar** | Pointer to **string** |  | [optional] 
**DefaultPassword** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DisableSignin** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Emoji** | Pointer to **string** |  | [optional] 
**EnableSoftDeletion** | Pointer to **bool** |  | [optional] 
**EnableTour** | Pointer to **bool** |  | [optional] 
**FailedSigninFrozenTime** | Pointer to **int32** |  | [optional] 
**FailedSigninLimit** | Pointer to **int32** | Per-organization signin throttle. Zero means \&quot;inherit the application default\&quot;; a non-zero value overrides it. Safe bounds are clamped by the resource service before persistence. | [optional] 
**Favicon** | Pointer to **string** |  | [optional] 
**Founder** | Pointer to **string** | Founder is the stable storage id of the identity that provisioned this org (self-service onboarding). It is the resume token that makes provisioning converge on a backend where each write autocommits independently (no transaction rollback): after a partial failure that created the org but did not move the founder in, a retry recognises the org as the founder&#39;s own and completes it, instead of refusing it as \&quot;already taken\&quot;. It also fences the org to ONE tenant — a different identity can never complete or join it. | [optional] 
**HasPrivilegeConsent** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InitScore** | Pointer to **int32** |  | [optional] 
**IpRestriction** | Pointer to **string** |  | [optional] 
**IpWhitelist** | Pointer to **string** |  | [optional] 
**IsPersonal** | Pointer to **bool** |  | [optional] 
**IsProfilePublic** | Pointer to **bool** |  | [optional] 
**KerberosKdcHost** | Pointer to **string** |  | [optional] 
**KerberosKeytab** | Pointer to **string** |  | [optional] 
**KerberosRealm** | Pointer to **string** |  | [optional] 
**KerberosServiceName** | Pointer to **string** |  | [optional] 
**Languages** | Pointer to **[]string** |  | [optional] 
**LdapAttributes** | Pointer to **[]string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**LogoDark** | Pointer to **string** |  | [optional] 
**MasterPassword** | Pointer to **string** |  | [optional] 
**MasterVerificationCode** | Pointer to **string** |  | [optional] 
**MfaItems** | Pointer to [**[]IamMfaItem**](IamMfaItem.md) |  | [optional] 
**MfaRememberInHours** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NavItems** | Pointer to **[]string** |  | [optional] 
**OrgBalance** | Pointer to **float32** | Balance fields are read-only mirrors; authoritative balances live in Commerce (billing.hanzo.ai). Carried for field-complete v1 parity. | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**PasswordExpireDays** | Pointer to **int32** |  | [optional] 
**PasswordObfuscatorKey** | Pointer to **string** |  | [optional] 
**PasswordObfuscatorType** | Pointer to **string** |  | [optional] 
**PasswordOptions** | Pointer to **[]string** |  | [optional] 
**PasswordSalt** | Pointer to **string** |  | [optional] 
**PasswordType** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ThemeData** | Pointer to [**IamThemeData**](IamThemeData.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UseEmailAsUsername** | Pointer to **bool** |  | [optional] 
**UsePermanentAvatar** | Pointer to **bool** |  | [optional] 
**UserBalance** | Pointer to **float32** |  | [optional] 
**UserNavItems** | Pointer to **[]string** |  | [optional] 
**UserTypes** | Pointer to **[]string** |  | [optional] 
**WebsiteUrl** | Pointer to **string** |  | [optional] 
**WidgetItems** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIamOrganization

`func NewIamOrganization() *IamOrganization`

NewIamOrganization instantiates a new IamOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamOrganizationWithDefaults

`func NewIamOrganizationWithDefaults() *IamOrganization`

NewIamOrganizationWithDefaults instantiates a new IamOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItems

`func (o *IamOrganization) GetAccountItems() []IamAccountItem`

GetAccountItems returns the AccountItems field if non-nil, zero value otherwise.

### GetAccountItemsOk

`func (o *IamOrganization) GetAccountItemsOk() (*[]IamAccountItem, bool)`

GetAccountItemsOk returns a tuple with the AccountItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItems

`func (o *IamOrganization) SetAccountItems(v []IamAccountItem)`

SetAccountItems sets AccountItems field to given value.

### HasAccountItems

`func (o *IamOrganization) HasAccountItems() bool`

HasAccountItems returns a boolean if a field has been set.

### GetAccountMenu

`func (o *IamOrganization) GetAccountMenu() string`

GetAccountMenu returns the AccountMenu field if non-nil, zero value otherwise.

### GetAccountMenuOk

`func (o *IamOrganization) GetAccountMenuOk() (*string, bool)`

GetAccountMenuOk returns a tuple with the AccountMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMenu

`func (o *IamOrganization) SetAccountMenu(v string)`

SetAccountMenu sets AccountMenu field to given value.

### HasAccountMenu

`func (o *IamOrganization) HasAccountMenu() bool`

HasAccountMenu returns a boolean if a field has been set.

### GetAvatar

`func (o *IamOrganization) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *IamOrganization) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *IamOrganization) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *IamOrganization) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetBalanceCredit

`func (o *IamOrganization) GetBalanceCredit() float32`

GetBalanceCredit returns the BalanceCredit field if non-nil, zero value otherwise.

### GetBalanceCreditOk

`func (o *IamOrganization) GetBalanceCreditOk() (*float32, bool)`

GetBalanceCreditOk returns a tuple with the BalanceCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCredit

`func (o *IamOrganization) SetBalanceCredit(v float32)`

SetBalanceCredit sets BalanceCredit field to given value.

### HasBalanceCredit

`func (o *IamOrganization) HasBalanceCredit() bool`

HasBalanceCredit returns a boolean if a field has been set.

### GetBalanceCurrency

`func (o *IamOrganization) GetBalanceCurrency() string`

GetBalanceCurrency returns the BalanceCurrency field if non-nil, zero value otherwise.

### GetBalanceCurrencyOk

`func (o *IamOrganization) GetBalanceCurrencyOk() (*string, bool)`

GetBalanceCurrencyOk returns a tuple with the BalanceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCurrency

`func (o *IamOrganization) SetBalanceCurrency(v string)`

SetBalanceCurrency sets BalanceCurrency field to given value.

### HasBalanceCurrency

`func (o *IamOrganization) HasBalanceCurrency() bool`

HasBalanceCurrency returns a boolean if a field has been set.

### GetCountryCodes

`func (o *IamOrganization) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *IamOrganization) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *IamOrganization) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *IamOrganization) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IamOrganization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamOrganization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamOrganization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamOrganization) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamOrganization) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamOrganization) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamOrganization) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamOrganization) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDcrPolicy

`func (o *IamOrganization) GetDcrPolicy() string`

GetDcrPolicy returns the DcrPolicy field if non-nil, zero value otherwise.

### GetDcrPolicyOk

`func (o *IamOrganization) GetDcrPolicyOk() (*string, bool)`

GetDcrPolicyOk returns a tuple with the DcrPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcrPolicy

`func (o *IamOrganization) SetDcrPolicy(v string)`

SetDcrPolicy sets DcrPolicy field to given value.

### HasDcrPolicy

`func (o *IamOrganization) HasDcrPolicy() bool`

HasDcrPolicy returns a boolean if a field has been set.

### GetDefaultApplication

`func (o *IamOrganization) GetDefaultApplication() string`

GetDefaultApplication returns the DefaultApplication field if non-nil, zero value otherwise.

### GetDefaultApplicationOk

`func (o *IamOrganization) GetDefaultApplicationOk() (*string, bool)`

GetDefaultApplicationOk returns a tuple with the DefaultApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApplication

`func (o *IamOrganization) SetDefaultApplication(v string)`

SetDefaultApplication sets DefaultApplication field to given value.

### HasDefaultApplication

`func (o *IamOrganization) HasDefaultApplication() bool`

HasDefaultApplication returns a boolean if a field has been set.

### GetDefaultAvatar

`func (o *IamOrganization) GetDefaultAvatar() string`

GetDefaultAvatar returns the DefaultAvatar field if non-nil, zero value otherwise.

### GetDefaultAvatarOk

`func (o *IamOrganization) GetDefaultAvatarOk() (*string, bool)`

GetDefaultAvatarOk returns a tuple with the DefaultAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAvatar

`func (o *IamOrganization) SetDefaultAvatar(v string)`

SetDefaultAvatar sets DefaultAvatar field to given value.

### HasDefaultAvatar

`func (o *IamOrganization) HasDefaultAvatar() bool`

HasDefaultAvatar returns a boolean if a field has been set.

### GetDefaultPassword

`func (o *IamOrganization) GetDefaultPassword() string`

GetDefaultPassword returns the DefaultPassword field if non-nil, zero value otherwise.

### GetDefaultPasswordOk

`func (o *IamOrganization) GetDefaultPasswordOk() (*string, bool)`

GetDefaultPasswordOk returns a tuple with the DefaultPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPassword

`func (o *IamOrganization) SetDefaultPassword(v string)`

SetDefaultPassword sets DefaultPassword field to given value.

### HasDefaultPassword

`func (o *IamOrganization) HasDefaultPassword() bool`

HasDefaultPassword returns a boolean if a field has been set.

### GetDeleted

`func (o *IamOrganization) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamOrganization) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamOrganization) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamOrganization) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDisableSignin

`func (o *IamOrganization) GetDisableSignin() bool`

GetDisableSignin returns the DisableSignin field if non-nil, zero value otherwise.

### GetDisableSigninOk

`func (o *IamOrganization) GetDisableSigninOk() (*bool, bool)`

GetDisableSigninOk returns a tuple with the DisableSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSignin

`func (o *IamOrganization) SetDisableSignin(v bool)`

SetDisableSignin sets DisableSignin field to given value.

### HasDisableSignin

`func (o *IamOrganization) HasDisableSignin() bool`

HasDisableSignin returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamOrganization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamOrganization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamOrganization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamOrganization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmoji

`func (o *IamOrganization) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *IamOrganization) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *IamOrganization) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *IamOrganization) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetEnableSoftDeletion

`func (o *IamOrganization) GetEnableSoftDeletion() bool`

GetEnableSoftDeletion returns the EnableSoftDeletion field if non-nil, zero value otherwise.

### GetEnableSoftDeletionOk

`func (o *IamOrganization) GetEnableSoftDeletionOk() (*bool, bool)`

GetEnableSoftDeletionOk returns a tuple with the EnableSoftDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSoftDeletion

`func (o *IamOrganization) SetEnableSoftDeletion(v bool)`

SetEnableSoftDeletion sets EnableSoftDeletion field to given value.

### HasEnableSoftDeletion

`func (o *IamOrganization) HasEnableSoftDeletion() bool`

HasEnableSoftDeletion returns a boolean if a field has been set.

### GetEnableTour

`func (o *IamOrganization) GetEnableTour() bool`

GetEnableTour returns the EnableTour field if non-nil, zero value otherwise.

### GetEnableTourOk

`func (o *IamOrganization) GetEnableTourOk() (*bool, bool)`

GetEnableTourOk returns a tuple with the EnableTour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTour

`func (o *IamOrganization) SetEnableTour(v bool)`

SetEnableTour sets EnableTour field to given value.

### HasEnableTour

`func (o *IamOrganization) HasEnableTour() bool`

HasEnableTour returns a boolean if a field has been set.

### GetFailedSigninFrozenTime

`func (o *IamOrganization) GetFailedSigninFrozenTime() int32`

GetFailedSigninFrozenTime returns the FailedSigninFrozenTime field if non-nil, zero value otherwise.

### GetFailedSigninFrozenTimeOk

`func (o *IamOrganization) GetFailedSigninFrozenTimeOk() (*int32, bool)`

GetFailedSigninFrozenTimeOk returns a tuple with the FailedSigninFrozenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedSigninFrozenTime

`func (o *IamOrganization) SetFailedSigninFrozenTime(v int32)`

SetFailedSigninFrozenTime sets FailedSigninFrozenTime field to given value.

### HasFailedSigninFrozenTime

`func (o *IamOrganization) HasFailedSigninFrozenTime() bool`

HasFailedSigninFrozenTime returns a boolean if a field has been set.

### GetFailedSigninLimit

`func (o *IamOrganization) GetFailedSigninLimit() int32`

GetFailedSigninLimit returns the FailedSigninLimit field if non-nil, zero value otherwise.

### GetFailedSigninLimitOk

`func (o *IamOrganization) GetFailedSigninLimitOk() (*int32, bool)`

GetFailedSigninLimitOk returns a tuple with the FailedSigninLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedSigninLimit

`func (o *IamOrganization) SetFailedSigninLimit(v int32)`

SetFailedSigninLimit sets FailedSigninLimit field to given value.

### HasFailedSigninLimit

`func (o *IamOrganization) HasFailedSigninLimit() bool`

HasFailedSigninLimit returns a boolean if a field has been set.

### GetFavicon

`func (o *IamOrganization) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *IamOrganization) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *IamOrganization) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *IamOrganization) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### GetFounder

`func (o *IamOrganization) GetFounder() string`

GetFounder returns the Founder field if non-nil, zero value otherwise.

### GetFounderOk

`func (o *IamOrganization) GetFounderOk() (*string, bool)`

GetFounderOk returns a tuple with the Founder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFounder

`func (o *IamOrganization) SetFounder(v string)`

SetFounder sets Founder field to given value.

### HasFounder

`func (o *IamOrganization) HasFounder() bool`

HasFounder returns a boolean if a field has been set.

### GetHasPrivilegeConsent

`func (o *IamOrganization) GetHasPrivilegeConsent() bool`

GetHasPrivilegeConsent returns the HasPrivilegeConsent field if non-nil, zero value otherwise.

### GetHasPrivilegeConsentOk

`func (o *IamOrganization) GetHasPrivilegeConsentOk() (*bool, bool)`

GetHasPrivilegeConsentOk returns a tuple with the HasPrivilegeConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivilegeConsent

`func (o *IamOrganization) SetHasPrivilegeConsent(v bool)`

SetHasPrivilegeConsent sets HasPrivilegeConsent field to given value.

### HasHasPrivilegeConsent

`func (o *IamOrganization) HasHasPrivilegeConsent() bool`

HasHasPrivilegeConsent returns a boolean if a field has been set.

### GetId

`func (o *IamOrganization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamOrganization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamOrganization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamOrganization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInitScore

`func (o *IamOrganization) GetInitScore() int32`

GetInitScore returns the InitScore field if non-nil, zero value otherwise.

### GetInitScoreOk

`func (o *IamOrganization) GetInitScoreOk() (*int32, bool)`

GetInitScoreOk returns a tuple with the InitScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitScore

`func (o *IamOrganization) SetInitScore(v int32)`

SetInitScore sets InitScore field to given value.

### HasInitScore

`func (o *IamOrganization) HasInitScore() bool`

HasInitScore returns a boolean if a field has been set.

### GetIpRestriction

`func (o *IamOrganization) GetIpRestriction() string`

GetIpRestriction returns the IpRestriction field if non-nil, zero value otherwise.

### GetIpRestrictionOk

`func (o *IamOrganization) GetIpRestrictionOk() (*string, bool)`

GetIpRestrictionOk returns a tuple with the IpRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRestriction

`func (o *IamOrganization) SetIpRestriction(v string)`

SetIpRestriction sets IpRestriction field to given value.

### HasIpRestriction

`func (o *IamOrganization) HasIpRestriction() bool`

HasIpRestriction returns a boolean if a field has been set.

### GetIpWhitelist

`func (o *IamOrganization) GetIpWhitelist() string`

GetIpWhitelist returns the IpWhitelist field if non-nil, zero value otherwise.

### GetIpWhitelistOk

`func (o *IamOrganization) GetIpWhitelistOk() (*string, bool)`

GetIpWhitelistOk returns a tuple with the IpWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelist

`func (o *IamOrganization) SetIpWhitelist(v string)`

SetIpWhitelist sets IpWhitelist field to given value.

### HasIpWhitelist

`func (o *IamOrganization) HasIpWhitelist() bool`

HasIpWhitelist returns a boolean if a field has been set.

### GetIsPersonal

`func (o *IamOrganization) GetIsPersonal() bool`

GetIsPersonal returns the IsPersonal field if non-nil, zero value otherwise.

### GetIsPersonalOk

`func (o *IamOrganization) GetIsPersonalOk() (*bool, bool)`

GetIsPersonalOk returns a tuple with the IsPersonal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersonal

`func (o *IamOrganization) SetIsPersonal(v bool)`

SetIsPersonal sets IsPersonal field to given value.

### HasIsPersonal

`func (o *IamOrganization) HasIsPersonal() bool`

HasIsPersonal returns a boolean if a field has been set.

### GetIsProfilePublic

`func (o *IamOrganization) GetIsProfilePublic() bool`

GetIsProfilePublic returns the IsProfilePublic field if non-nil, zero value otherwise.

### GetIsProfilePublicOk

`func (o *IamOrganization) GetIsProfilePublicOk() (*bool, bool)`

GetIsProfilePublicOk returns a tuple with the IsProfilePublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProfilePublic

`func (o *IamOrganization) SetIsProfilePublic(v bool)`

SetIsProfilePublic sets IsProfilePublic field to given value.

### HasIsProfilePublic

`func (o *IamOrganization) HasIsProfilePublic() bool`

HasIsProfilePublic returns a boolean if a field has been set.

### GetKerberosKdcHost

`func (o *IamOrganization) GetKerberosKdcHost() string`

GetKerberosKdcHost returns the KerberosKdcHost field if non-nil, zero value otherwise.

### GetKerberosKdcHostOk

`func (o *IamOrganization) GetKerberosKdcHostOk() (*string, bool)`

GetKerberosKdcHostOk returns a tuple with the KerberosKdcHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKdcHost

`func (o *IamOrganization) SetKerberosKdcHost(v string)`

SetKerberosKdcHost sets KerberosKdcHost field to given value.

### HasKerberosKdcHost

`func (o *IamOrganization) HasKerberosKdcHost() bool`

HasKerberosKdcHost returns a boolean if a field has been set.

### GetKerberosKeytab

`func (o *IamOrganization) GetKerberosKeytab() string`

GetKerberosKeytab returns the KerberosKeytab field if non-nil, zero value otherwise.

### GetKerberosKeytabOk

`func (o *IamOrganization) GetKerberosKeytabOk() (*string, bool)`

GetKerberosKeytabOk returns a tuple with the KerberosKeytab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKeytab

`func (o *IamOrganization) SetKerberosKeytab(v string)`

SetKerberosKeytab sets KerberosKeytab field to given value.

### HasKerberosKeytab

`func (o *IamOrganization) HasKerberosKeytab() bool`

HasKerberosKeytab returns a boolean if a field has been set.

### GetKerberosRealm

`func (o *IamOrganization) GetKerberosRealm() string`

GetKerberosRealm returns the KerberosRealm field if non-nil, zero value otherwise.

### GetKerberosRealmOk

`func (o *IamOrganization) GetKerberosRealmOk() (*string, bool)`

GetKerberosRealmOk returns a tuple with the KerberosRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosRealm

`func (o *IamOrganization) SetKerberosRealm(v string)`

SetKerberosRealm sets KerberosRealm field to given value.

### HasKerberosRealm

`func (o *IamOrganization) HasKerberosRealm() bool`

HasKerberosRealm returns a boolean if a field has been set.

### GetKerberosServiceName

`func (o *IamOrganization) GetKerberosServiceName() string`

GetKerberosServiceName returns the KerberosServiceName field if non-nil, zero value otherwise.

### GetKerberosServiceNameOk

`func (o *IamOrganization) GetKerberosServiceNameOk() (*string, bool)`

GetKerberosServiceNameOk returns a tuple with the KerberosServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosServiceName

`func (o *IamOrganization) SetKerberosServiceName(v string)`

SetKerberosServiceName sets KerberosServiceName field to given value.

### HasKerberosServiceName

`func (o *IamOrganization) HasKerberosServiceName() bool`

HasKerberosServiceName returns a boolean if a field has been set.

### GetLanguages

`func (o *IamOrganization) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *IamOrganization) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *IamOrganization) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *IamOrganization) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLdapAttributes

`func (o *IamOrganization) GetLdapAttributes() []string`

GetLdapAttributes returns the LdapAttributes field if non-nil, zero value otherwise.

### GetLdapAttributesOk

`func (o *IamOrganization) GetLdapAttributesOk() (*[]string, bool)`

GetLdapAttributesOk returns a tuple with the LdapAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAttributes

`func (o *IamOrganization) SetLdapAttributes(v []string)`

SetLdapAttributes sets LdapAttributes field to given value.

### HasLdapAttributes

`func (o *IamOrganization) HasLdapAttributes() bool`

HasLdapAttributes returns a boolean if a field has been set.

### GetLogo

`func (o *IamOrganization) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *IamOrganization) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *IamOrganization) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *IamOrganization) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetLogoDark

`func (o *IamOrganization) GetLogoDark() string`

GetLogoDark returns the LogoDark field if non-nil, zero value otherwise.

### GetLogoDarkOk

`func (o *IamOrganization) GetLogoDarkOk() (*string, bool)`

GetLogoDarkOk returns a tuple with the LogoDark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoDark

`func (o *IamOrganization) SetLogoDark(v string)`

SetLogoDark sets LogoDark field to given value.

### HasLogoDark

`func (o *IamOrganization) HasLogoDark() bool`

HasLogoDark returns a boolean if a field has been set.

### GetMasterPassword

`func (o *IamOrganization) GetMasterPassword() string`

GetMasterPassword returns the MasterPassword field if non-nil, zero value otherwise.

### GetMasterPasswordOk

`func (o *IamOrganization) GetMasterPasswordOk() (*string, bool)`

GetMasterPasswordOk returns a tuple with the MasterPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterPassword

`func (o *IamOrganization) SetMasterPassword(v string)`

SetMasterPassword sets MasterPassword field to given value.

### HasMasterPassword

`func (o *IamOrganization) HasMasterPassword() bool`

HasMasterPassword returns a boolean if a field has been set.

### GetMasterVerificationCode

`func (o *IamOrganization) GetMasterVerificationCode() string`

GetMasterVerificationCode returns the MasterVerificationCode field if non-nil, zero value otherwise.

### GetMasterVerificationCodeOk

`func (o *IamOrganization) GetMasterVerificationCodeOk() (*string, bool)`

GetMasterVerificationCodeOk returns a tuple with the MasterVerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterVerificationCode

`func (o *IamOrganization) SetMasterVerificationCode(v string)`

SetMasterVerificationCode sets MasterVerificationCode field to given value.

### HasMasterVerificationCode

`func (o *IamOrganization) HasMasterVerificationCode() bool`

HasMasterVerificationCode returns a boolean if a field has been set.

### GetMfaItems

`func (o *IamOrganization) GetMfaItems() []IamMfaItem`

GetMfaItems returns the MfaItems field if non-nil, zero value otherwise.

### GetMfaItemsOk

`func (o *IamOrganization) GetMfaItemsOk() (*[]IamMfaItem, bool)`

GetMfaItemsOk returns a tuple with the MfaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaItems

`func (o *IamOrganization) SetMfaItems(v []IamMfaItem)`

SetMfaItems sets MfaItems field to given value.

### HasMfaItems

`func (o *IamOrganization) HasMfaItems() bool`

HasMfaItems returns a boolean if a field has been set.

### GetMfaRememberInHours

`func (o *IamOrganization) GetMfaRememberInHours() int32`

GetMfaRememberInHours returns the MfaRememberInHours field if non-nil, zero value otherwise.

### GetMfaRememberInHoursOk

`func (o *IamOrganization) GetMfaRememberInHoursOk() (*int32, bool)`

GetMfaRememberInHoursOk returns a tuple with the MfaRememberInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRememberInHours

`func (o *IamOrganization) SetMfaRememberInHours(v int32)`

SetMfaRememberInHours sets MfaRememberInHours field to given value.

### HasMfaRememberInHours

`func (o *IamOrganization) HasMfaRememberInHours() bool`

HasMfaRememberInHours returns a boolean if a field has been set.

### GetName

`func (o *IamOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamOrganization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamOrganization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNavItems

`func (o *IamOrganization) GetNavItems() []string`

GetNavItems returns the NavItems field if non-nil, zero value otherwise.

### GetNavItemsOk

`func (o *IamOrganization) GetNavItemsOk() (*[]string, bool)`

GetNavItemsOk returns a tuple with the NavItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavItems

`func (o *IamOrganization) SetNavItems(v []string)`

SetNavItems sets NavItems field to given value.

### HasNavItems

`func (o *IamOrganization) HasNavItems() bool`

HasNavItems returns a boolean if a field has been set.

### GetOrgBalance

`func (o *IamOrganization) GetOrgBalance() float32`

GetOrgBalance returns the OrgBalance field if non-nil, zero value otherwise.

### GetOrgBalanceOk

`func (o *IamOrganization) GetOrgBalanceOk() (*float32, bool)`

GetOrgBalanceOk returns a tuple with the OrgBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgBalance

`func (o *IamOrganization) SetOrgBalance(v float32)`

SetOrgBalance sets OrgBalance field to given value.

### HasOrgBalance

`func (o *IamOrganization) HasOrgBalance() bool`

HasOrgBalance returns a boolean if a field has been set.

### GetOwner

`func (o *IamOrganization) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamOrganization) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamOrganization) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamOrganization) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPasswordExpireDays

`func (o *IamOrganization) GetPasswordExpireDays() int32`

GetPasswordExpireDays returns the PasswordExpireDays field if non-nil, zero value otherwise.

### GetPasswordExpireDaysOk

`func (o *IamOrganization) GetPasswordExpireDaysOk() (*int32, bool)`

GetPasswordExpireDaysOk returns a tuple with the PasswordExpireDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpireDays

`func (o *IamOrganization) SetPasswordExpireDays(v int32)`

SetPasswordExpireDays sets PasswordExpireDays field to given value.

### HasPasswordExpireDays

`func (o *IamOrganization) HasPasswordExpireDays() bool`

HasPasswordExpireDays returns a boolean if a field has been set.

### GetPasswordObfuscatorKey

`func (o *IamOrganization) GetPasswordObfuscatorKey() string`

GetPasswordObfuscatorKey returns the PasswordObfuscatorKey field if non-nil, zero value otherwise.

### GetPasswordObfuscatorKeyOk

`func (o *IamOrganization) GetPasswordObfuscatorKeyOk() (*string, bool)`

GetPasswordObfuscatorKeyOk returns a tuple with the PasswordObfuscatorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordObfuscatorKey

`func (o *IamOrganization) SetPasswordObfuscatorKey(v string)`

SetPasswordObfuscatorKey sets PasswordObfuscatorKey field to given value.

### HasPasswordObfuscatorKey

`func (o *IamOrganization) HasPasswordObfuscatorKey() bool`

HasPasswordObfuscatorKey returns a boolean if a field has been set.

### GetPasswordObfuscatorType

`func (o *IamOrganization) GetPasswordObfuscatorType() string`

GetPasswordObfuscatorType returns the PasswordObfuscatorType field if non-nil, zero value otherwise.

### GetPasswordObfuscatorTypeOk

`func (o *IamOrganization) GetPasswordObfuscatorTypeOk() (*string, bool)`

GetPasswordObfuscatorTypeOk returns a tuple with the PasswordObfuscatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordObfuscatorType

`func (o *IamOrganization) SetPasswordObfuscatorType(v string)`

SetPasswordObfuscatorType sets PasswordObfuscatorType field to given value.

### HasPasswordObfuscatorType

`func (o *IamOrganization) HasPasswordObfuscatorType() bool`

HasPasswordObfuscatorType returns a boolean if a field has been set.

### GetPasswordOptions

`func (o *IamOrganization) GetPasswordOptions() []string`

GetPasswordOptions returns the PasswordOptions field if non-nil, zero value otherwise.

### GetPasswordOptionsOk

`func (o *IamOrganization) GetPasswordOptionsOk() (*[]string, bool)`

GetPasswordOptionsOk returns a tuple with the PasswordOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordOptions

`func (o *IamOrganization) SetPasswordOptions(v []string)`

SetPasswordOptions sets PasswordOptions field to given value.

### HasPasswordOptions

`func (o *IamOrganization) HasPasswordOptions() bool`

HasPasswordOptions returns a boolean if a field has been set.

### GetPasswordSalt

`func (o *IamOrganization) GetPasswordSalt() string`

GetPasswordSalt returns the PasswordSalt field if non-nil, zero value otherwise.

### GetPasswordSaltOk

`func (o *IamOrganization) GetPasswordSaltOk() (*string, bool)`

GetPasswordSaltOk returns a tuple with the PasswordSalt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSalt

`func (o *IamOrganization) SetPasswordSalt(v string)`

SetPasswordSalt sets PasswordSalt field to given value.

### HasPasswordSalt

`func (o *IamOrganization) HasPasswordSalt() bool`

HasPasswordSalt returns a boolean if a field has been set.

### GetPasswordType

`func (o *IamOrganization) GetPasswordType() string`

GetPasswordType returns the PasswordType field if non-nil, zero value otherwise.

### GetPasswordTypeOk

`func (o *IamOrganization) GetPasswordTypeOk() (*string, bool)`

GetPasswordTypeOk returns a tuple with the PasswordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordType

`func (o *IamOrganization) SetPasswordType(v string)`

SetPasswordType sets PasswordType field to given value.

### HasPasswordType

`func (o *IamOrganization) HasPasswordType() bool`

HasPasswordType returns a boolean if a field has been set.

### GetTags

`func (o *IamOrganization) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IamOrganization) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IamOrganization) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IamOrganization) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThemeData

`func (o *IamOrganization) GetThemeData() IamThemeData`

GetThemeData returns the ThemeData field if non-nil, zero value otherwise.

### GetThemeDataOk

`func (o *IamOrganization) GetThemeDataOk() (*IamThemeData, bool)`

GetThemeDataOk returns a tuple with the ThemeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeData

`func (o *IamOrganization) SetThemeData(v IamThemeData)`

SetThemeData sets ThemeData field to given value.

### HasThemeData

`func (o *IamOrganization) HasThemeData() bool`

HasThemeData returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamOrganization) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamOrganization) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamOrganization) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamOrganization) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUseEmailAsUsername

`func (o *IamOrganization) GetUseEmailAsUsername() bool`

GetUseEmailAsUsername returns the UseEmailAsUsername field if non-nil, zero value otherwise.

### GetUseEmailAsUsernameOk

`func (o *IamOrganization) GetUseEmailAsUsernameOk() (*bool, bool)`

GetUseEmailAsUsernameOk returns a tuple with the UseEmailAsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmailAsUsername

`func (o *IamOrganization) SetUseEmailAsUsername(v bool)`

SetUseEmailAsUsername sets UseEmailAsUsername field to given value.

### HasUseEmailAsUsername

`func (o *IamOrganization) HasUseEmailAsUsername() bool`

HasUseEmailAsUsername returns a boolean if a field has been set.

### GetUsePermanentAvatar

`func (o *IamOrganization) GetUsePermanentAvatar() bool`

GetUsePermanentAvatar returns the UsePermanentAvatar field if non-nil, zero value otherwise.

### GetUsePermanentAvatarOk

`func (o *IamOrganization) GetUsePermanentAvatarOk() (*bool, bool)`

GetUsePermanentAvatarOk returns a tuple with the UsePermanentAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePermanentAvatar

`func (o *IamOrganization) SetUsePermanentAvatar(v bool)`

SetUsePermanentAvatar sets UsePermanentAvatar field to given value.

### HasUsePermanentAvatar

`func (o *IamOrganization) HasUsePermanentAvatar() bool`

HasUsePermanentAvatar returns a boolean if a field has been set.

### GetUserBalance

`func (o *IamOrganization) GetUserBalance() float32`

GetUserBalance returns the UserBalance field if non-nil, zero value otherwise.

### GetUserBalanceOk

`func (o *IamOrganization) GetUserBalanceOk() (*float32, bool)`

GetUserBalanceOk returns a tuple with the UserBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBalance

`func (o *IamOrganization) SetUserBalance(v float32)`

SetUserBalance sets UserBalance field to given value.

### HasUserBalance

`func (o *IamOrganization) HasUserBalance() bool`

HasUserBalance returns a boolean if a field has been set.

### GetUserNavItems

`func (o *IamOrganization) GetUserNavItems() []string`

GetUserNavItems returns the UserNavItems field if non-nil, zero value otherwise.

### GetUserNavItemsOk

`func (o *IamOrganization) GetUserNavItemsOk() (*[]string, bool)`

GetUserNavItemsOk returns a tuple with the UserNavItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNavItems

`func (o *IamOrganization) SetUserNavItems(v []string)`

SetUserNavItems sets UserNavItems field to given value.

### HasUserNavItems

`func (o *IamOrganization) HasUserNavItems() bool`

HasUserNavItems returns a boolean if a field has been set.

### GetUserTypes

`func (o *IamOrganization) GetUserTypes() []string`

GetUserTypes returns the UserTypes field if non-nil, zero value otherwise.

### GetUserTypesOk

`func (o *IamOrganization) GetUserTypesOk() (*[]string, bool)`

GetUserTypesOk returns a tuple with the UserTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTypes

`func (o *IamOrganization) SetUserTypes(v []string)`

SetUserTypes sets UserTypes field to given value.

### HasUserTypes

`func (o *IamOrganization) HasUserTypes() bool`

HasUserTypes returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *IamOrganization) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *IamOrganization) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *IamOrganization) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *IamOrganization) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetWidgetItems

`func (o *IamOrganization) GetWidgetItems() []string`

GetWidgetItems returns the WidgetItems field if non-nil, zero value otherwise.

### GetWidgetItemsOk

`func (o *IamOrganization) GetWidgetItemsOk() (*[]string, bool)`

GetWidgetItemsOk returns a tuple with the WidgetItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetItems

`func (o *IamOrganization) SetWidgetItems(v []string)`

SetWidgetItems sets WidgetItems field to given value.

### HasWidgetItems

`func (o *IamOrganization) HasWidgetItems() bool`

HasWidgetItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IamUpdateOrganizationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItems** | Pointer to [**[]IamAccountItem**](IamAccountItem.md) |  | [optional] 
**AccountMenu** | Pointer to **string** |  | [optional] 
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
**EnableSoftDeletion** | Pointer to **bool** |  | [optional] 
**EnableTour** | Pointer to **bool** |  | [optional] 
**FailedSigninFrozenTime** | Pointer to **int32** |  | [optional] 
**FailedSigninLimit** | Pointer to **int32** |  | [optional] 
**Favicon** | Pointer to **string** |  | [optional] 
**Founder** | Pointer to **string** |  | [optional] 
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
**OrgBalance** | Pointer to **float32** |  | [optional] 
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

### NewIamUpdateOrganizationInput

`func NewIamUpdateOrganizationInput() *IamUpdateOrganizationInput`

NewIamUpdateOrganizationInput instantiates a new IamUpdateOrganizationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUpdateOrganizationInputWithDefaults

`func NewIamUpdateOrganizationInputWithDefaults() *IamUpdateOrganizationInput`

NewIamUpdateOrganizationInputWithDefaults instantiates a new IamUpdateOrganizationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItems

`func (o *IamUpdateOrganizationInput) GetAccountItems() []IamAccountItem`

GetAccountItems returns the AccountItems field if non-nil, zero value otherwise.

### GetAccountItemsOk

`func (o *IamUpdateOrganizationInput) GetAccountItemsOk() (*[]IamAccountItem, bool)`

GetAccountItemsOk returns a tuple with the AccountItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItems

`func (o *IamUpdateOrganizationInput) SetAccountItems(v []IamAccountItem)`

SetAccountItems sets AccountItems field to given value.

### HasAccountItems

`func (o *IamUpdateOrganizationInput) HasAccountItems() bool`

HasAccountItems returns a boolean if a field has been set.

### GetAccountMenu

`func (o *IamUpdateOrganizationInput) GetAccountMenu() string`

GetAccountMenu returns the AccountMenu field if non-nil, zero value otherwise.

### GetAccountMenuOk

`func (o *IamUpdateOrganizationInput) GetAccountMenuOk() (*string, bool)`

GetAccountMenuOk returns a tuple with the AccountMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMenu

`func (o *IamUpdateOrganizationInput) SetAccountMenu(v string)`

SetAccountMenu sets AccountMenu field to given value.

### HasAccountMenu

`func (o *IamUpdateOrganizationInput) HasAccountMenu() bool`

HasAccountMenu returns a boolean if a field has been set.

### GetBalanceCredit

`func (o *IamUpdateOrganizationInput) GetBalanceCredit() float32`

GetBalanceCredit returns the BalanceCredit field if non-nil, zero value otherwise.

### GetBalanceCreditOk

`func (o *IamUpdateOrganizationInput) GetBalanceCreditOk() (*float32, bool)`

GetBalanceCreditOk returns a tuple with the BalanceCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCredit

`func (o *IamUpdateOrganizationInput) SetBalanceCredit(v float32)`

SetBalanceCredit sets BalanceCredit field to given value.

### HasBalanceCredit

`func (o *IamUpdateOrganizationInput) HasBalanceCredit() bool`

HasBalanceCredit returns a boolean if a field has been set.

### GetBalanceCurrency

`func (o *IamUpdateOrganizationInput) GetBalanceCurrency() string`

GetBalanceCurrency returns the BalanceCurrency field if non-nil, zero value otherwise.

### GetBalanceCurrencyOk

`func (o *IamUpdateOrganizationInput) GetBalanceCurrencyOk() (*string, bool)`

GetBalanceCurrencyOk returns a tuple with the BalanceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCurrency

`func (o *IamUpdateOrganizationInput) SetBalanceCurrency(v string)`

SetBalanceCurrency sets BalanceCurrency field to given value.

### HasBalanceCurrency

`func (o *IamUpdateOrganizationInput) HasBalanceCurrency() bool`

HasBalanceCurrency returns a boolean if a field has been set.

### GetCountryCodes

`func (o *IamUpdateOrganizationInput) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *IamUpdateOrganizationInput) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *IamUpdateOrganizationInput) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *IamUpdateOrganizationInput) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IamUpdateOrganizationInput) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamUpdateOrganizationInput) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamUpdateOrganizationInput) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamUpdateOrganizationInput) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamUpdateOrganizationInput) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamUpdateOrganizationInput) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamUpdateOrganizationInput) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamUpdateOrganizationInput) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDcrPolicy

`func (o *IamUpdateOrganizationInput) GetDcrPolicy() string`

GetDcrPolicy returns the DcrPolicy field if non-nil, zero value otherwise.

### GetDcrPolicyOk

`func (o *IamUpdateOrganizationInput) GetDcrPolicyOk() (*string, bool)`

GetDcrPolicyOk returns a tuple with the DcrPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcrPolicy

`func (o *IamUpdateOrganizationInput) SetDcrPolicy(v string)`

SetDcrPolicy sets DcrPolicy field to given value.

### HasDcrPolicy

`func (o *IamUpdateOrganizationInput) HasDcrPolicy() bool`

HasDcrPolicy returns a boolean if a field has been set.

### GetDefaultApplication

`func (o *IamUpdateOrganizationInput) GetDefaultApplication() string`

GetDefaultApplication returns the DefaultApplication field if non-nil, zero value otherwise.

### GetDefaultApplicationOk

`func (o *IamUpdateOrganizationInput) GetDefaultApplicationOk() (*string, bool)`

GetDefaultApplicationOk returns a tuple with the DefaultApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApplication

`func (o *IamUpdateOrganizationInput) SetDefaultApplication(v string)`

SetDefaultApplication sets DefaultApplication field to given value.

### HasDefaultApplication

`func (o *IamUpdateOrganizationInput) HasDefaultApplication() bool`

HasDefaultApplication returns a boolean if a field has been set.

### GetDefaultAvatar

`func (o *IamUpdateOrganizationInput) GetDefaultAvatar() string`

GetDefaultAvatar returns the DefaultAvatar field if non-nil, zero value otherwise.

### GetDefaultAvatarOk

`func (o *IamUpdateOrganizationInput) GetDefaultAvatarOk() (*string, bool)`

GetDefaultAvatarOk returns a tuple with the DefaultAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAvatar

`func (o *IamUpdateOrganizationInput) SetDefaultAvatar(v string)`

SetDefaultAvatar sets DefaultAvatar field to given value.

### HasDefaultAvatar

`func (o *IamUpdateOrganizationInput) HasDefaultAvatar() bool`

HasDefaultAvatar returns a boolean if a field has been set.

### GetDefaultPassword

`func (o *IamUpdateOrganizationInput) GetDefaultPassword() string`

GetDefaultPassword returns the DefaultPassword field if non-nil, zero value otherwise.

### GetDefaultPasswordOk

`func (o *IamUpdateOrganizationInput) GetDefaultPasswordOk() (*string, bool)`

GetDefaultPasswordOk returns a tuple with the DefaultPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPassword

`func (o *IamUpdateOrganizationInput) SetDefaultPassword(v string)`

SetDefaultPassword sets DefaultPassword field to given value.

### HasDefaultPassword

`func (o *IamUpdateOrganizationInput) HasDefaultPassword() bool`

HasDefaultPassword returns a boolean if a field has been set.

### GetDeleted

`func (o *IamUpdateOrganizationInput) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamUpdateOrganizationInput) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamUpdateOrganizationInput) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamUpdateOrganizationInput) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDisableSignin

`func (o *IamUpdateOrganizationInput) GetDisableSignin() bool`

GetDisableSignin returns the DisableSignin field if non-nil, zero value otherwise.

### GetDisableSigninOk

`func (o *IamUpdateOrganizationInput) GetDisableSigninOk() (*bool, bool)`

GetDisableSigninOk returns a tuple with the DisableSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSignin

`func (o *IamUpdateOrganizationInput) SetDisableSignin(v bool)`

SetDisableSignin sets DisableSignin field to given value.

### HasDisableSignin

`func (o *IamUpdateOrganizationInput) HasDisableSignin() bool`

HasDisableSignin returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamUpdateOrganizationInput) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamUpdateOrganizationInput) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamUpdateOrganizationInput) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamUpdateOrganizationInput) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableSoftDeletion

`func (o *IamUpdateOrganizationInput) GetEnableSoftDeletion() bool`

GetEnableSoftDeletion returns the EnableSoftDeletion field if non-nil, zero value otherwise.

### GetEnableSoftDeletionOk

`func (o *IamUpdateOrganizationInput) GetEnableSoftDeletionOk() (*bool, bool)`

GetEnableSoftDeletionOk returns a tuple with the EnableSoftDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSoftDeletion

`func (o *IamUpdateOrganizationInput) SetEnableSoftDeletion(v bool)`

SetEnableSoftDeletion sets EnableSoftDeletion field to given value.

### HasEnableSoftDeletion

`func (o *IamUpdateOrganizationInput) HasEnableSoftDeletion() bool`

HasEnableSoftDeletion returns a boolean if a field has been set.

### GetEnableTour

`func (o *IamUpdateOrganizationInput) GetEnableTour() bool`

GetEnableTour returns the EnableTour field if non-nil, zero value otherwise.

### GetEnableTourOk

`func (o *IamUpdateOrganizationInput) GetEnableTourOk() (*bool, bool)`

GetEnableTourOk returns a tuple with the EnableTour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTour

`func (o *IamUpdateOrganizationInput) SetEnableTour(v bool)`

SetEnableTour sets EnableTour field to given value.

### HasEnableTour

`func (o *IamUpdateOrganizationInput) HasEnableTour() bool`

HasEnableTour returns a boolean if a field has been set.

### GetFailedSigninFrozenTime

`func (o *IamUpdateOrganizationInput) GetFailedSigninFrozenTime() int32`

GetFailedSigninFrozenTime returns the FailedSigninFrozenTime field if non-nil, zero value otherwise.

### GetFailedSigninFrozenTimeOk

`func (o *IamUpdateOrganizationInput) GetFailedSigninFrozenTimeOk() (*int32, bool)`

GetFailedSigninFrozenTimeOk returns a tuple with the FailedSigninFrozenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedSigninFrozenTime

`func (o *IamUpdateOrganizationInput) SetFailedSigninFrozenTime(v int32)`

SetFailedSigninFrozenTime sets FailedSigninFrozenTime field to given value.

### HasFailedSigninFrozenTime

`func (o *IamUpdateOrganizationInput) HasFailedSigninFrozenTime() bool`

HasFailedSigninFrozenTime returns a boolean if a field has been set.

### GetFailedSigninLimit

`func (o *IamUpdateOrganizationInput) GetFailedSigninLimit() int32`

GetFailedSigninLimit returns the FailedSigninLimit field if non-nil, zero value otherwise.

### GetFailedSigninLimitOk

`func (o *IamUpdateOrganizationInput) GetFailedSigninLimitOk() (*int32, bool)`

GetFailedSigninLimitOk returns a tuple with the FailedSigninLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedSigninLimit

`func (o *IamUpdateOrganizationInput) SetFailedSigninLimit(v int32)`

SetFailedSigninLimit sets FailedSigninLimit field to given value.

### HasFailedSigninLimit

`func (o *IamUpdateOrganizationInput) HasFailedSigninLimit() bool`

HasFailedSigninLimit returns a boolean if a field has been set.

### GetFavicon

`func (o *IamUpdateOrganizationInput) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *IamUpdateOrganizationInput) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *IamUpdateOrganizationInput) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *IamUpdateOrganizationInput) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### GetFounder

`func (o *IamUpdateOrganizationInput) GetFounder() string`

GetFounder returns the Founder field if non-nil, zero value otherwise.

### GetFounderOk

`func (o *IamUpdateOrganizationInput) GetFounderOk() (*string, bool)`

GetFounderOk returns a tuple with the Founder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFounder

`func (o *IamUpdateOrganizationInput) SetFounder(v string)`

SetFounder sets Founder field to given value.

### HasFounder

`func (o *IamUpdateOrganizationInput) HasFounder() bool`

HasFounder returns a boolean if a field has been set.

### GetHasPrivilegeConsent

`func (o *IamUpdateOrganizationInput) GetHasPrivilegeConsent() bool`

GetHasPrivilegeConsent returns the HasPrivilegeConsent field if non-nil, zero value otherwise.

### GetHasPrivilegeConsentOk

`func (o *IamUpdateOrganizationInput) GetHasPrivilegeConsentOk() (*bool, bool)`

GetHasPrivilegeConsentOk returns a tuple with the HasPrivilegeConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivilegeConsent

`func (o *IamUpdateOrganizationInput) SetHasPrivilegeConsent(v bool)`

SetHasPrivilegeConsent sets HasPrivilegeConsent field to given value.

### HasHasPrivilegeConsent

`func (o *IamUpdateOrganizationInput) HasHasPrivilegeConsent() bool`

HasHasPrivilegeConsent returns a boolean if a field has been set.

### GetId

`func (o *IamUpdateOrganizationInput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamUpdateOrganizationInput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamUpdateOrganizationInput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamUpdateOrganizationInput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInitScore

`func (o *IamUpdateOrganizationInput) GetInitScore() int32`

GetInitScore returns the InitScore field if non-nil, zero value otherwise.

### GetInitScoreOk

`func (o *IamUpdateOrganizationInput) GetInitScoreOk() (*int32, bool)`

GetInitScoreOk returns a tuple with the InitScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitScore

`func (o *IamUpdateOrganizationInput) SetInitScore(v int32)`

SetInitScore sets InitScore field to given value.

### HasInitScore

`func (o *IamUpdateOrganizationInput) HasInitScore() bool`

HasInitScore returns a boolean if a field has been set.

### GetIpRestriction

`func (o *IamUpdateOrganizationInput) GetIpRestriction() string`

GetIpRestriction returns the IpRestriction field if non-nil, zero value otherwise.

### GetIpRestrictionOk

`func (o *IamUpdateOrganizationInput) GetIpRestrictionOk() (*string, bool)`

GetIpRestrictionOk returns a tuple with the IpRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRestriction

`func (o *IamUpdateOrganizationInput) SetIpRestriction(v string)`

SetIpRestriction sets IpRestriction field to given value.

### HasIpRestriction

`func (o *IamUpdateOrganizationInput) HasIpRestriction() bool`

HasIpRestriction returns a boolean if a field has been set.

### GetIpWhitelist

`func (o *IamUpdateOrganizationInput) GetIpWhitelist() string`

GetIpWhitelist returns the IpWhitelist field if non-nil, zero value otherwise.

### GetIpWhitelistOk

`func (o *IamUpdateOrganizationInput) GetIpWhitelistOk() (*string, bool)`

GetIpWhitelistOk returns a tuple with the IpWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelist

`func (o *IamUpdateOrganizationInput) SetIpWhitelist(v string)`

SetIpWhitelist sets IpWhitelist field to given value.

### HasIpWhitelist

`func (o *IamUpdateOrganizationInput) HasIpWhitelist() bool`

HasIpWhitelist returns a boolean if a field has been set.

### GetIsPersonal

`func (o *IamUpdateOrganizationInput) GetIsPersonal() bool`

GetIsPersonal returns the IsPersonal field if non-nil, zero value otherwise.

### GetIsPersonalOk

`func (o *IamUpdateOrganizationInput) GetIsPersonalOk() (*bool, bool)`

GetIsPersonalOk returns a tuple with the IsPersonal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersonal

`func (o *IamUpdateOrganizationInput) SetIsPersonal(v bool)`

SetIsPersonal sets IsPersonal field to given value.

### HasIsPersonal

`func (o *IamUpdateOrganizationInput) HasIsPersonal() bool`

HasIsPersonal returns a boolean if a field has been set.

### GetIsProfilePublic

`func (o *IamUpdateOrganizationInput) GetIsProfilePublic() bool`

GetIsProfilePublic returns the IsProfilePublic field if non-nil, zero value otherwise.

### GetIsProfilePublicOk

`func (o *IamUpdateOrganizationInput) GetIsProfilePublicOk() (*bool, bool)`

GetIsProfilePublicOk returns a tuple with the IsProfilePublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProfilePublic

`func (o *IamUpdateOrganizationInput) SetIsProfilePublic(v bool)`

SetIsProfilePublic sets IsProfilePublic field to given value.

### HasIsProfilePublic

`func (o *IamUpdateOrganizationInput) HasIsProfilePublic() bool`

HasIsProfilePublic returns a boolean if a field has been set.

### GetKerberosKdcHost

`func (o *IamUpdateOrganizationInput) GetKerberosKdcHost() string`

GetKerberosKdcHost returns the KerberosKdcHost field if non-nil, zero value otherwise.

### GetKerberosKdcHostOk

`func (o *IamUpdateOrganizationInput) GetKerberosKdcHostOk() (*string, bool)`

GetKerberosKdcHostOk returns a tuple with the KerberosKdcHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKdcHost

`func (o *IamUpdateOrganizationInput) SetKerberosKdcHost(v string)`

SetKerberosKdcHost sets KerberosKdcHost field to given value.

### HasKerberosKdcHost

`func (o *IamUpdateOrganizationInput) HasKerberosKdcHost() bool`

HasKerberosKdcHost returns a boolean if a field has been set.

### GetKerberosKeytab

`func (o *IamUpdateOrganizationInput) GetKerberosKeytab() string`

GetKerberosKeytab returns the KerberosKeytab field if non-nil, zero value otherwise.

### GetKerberosKeytabOk

`func (o *IamUpdateOrganizationInput) GetKerberosKeytabOk() (*string, bool)`

GetKerberosKeytabOk returns a tuple with the KerberosKeytab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKeytab

`func (o *IamUpdateOrganizationInput) SetKerberosKeytab(v string)`

SetKerberosKeytab sets KerberosKeytab field to given value.

### HasKerberosKeytab

`func (o *IamUpdateOrganizationInput) HasKerberosKeytab() bool`

HasKerberosKeytab returns a boolean if a field has been set.

### GetKerberosRealm

`func (o *IamUpdateOrganizationInput) GetKerberosRealm() string`

GetKerberosRealm returns the KerberosRealm field if non-nil, zero value otherwise.

### GetKerberosRealmOk

`func (o *IamUpdateOrganizationInput) GetKerberosRealmOk() (*string, bool)`

GetKerberosRealmOk returns a tuple with the KerberosRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosRealm

`func (o *IamUpdateOrganizationInput) SetKerberosRealm(v string)`

SetKerberosRealm sets KerberosRealm field to given value.

### HasKerberosRealm

`func (o *IamUpdateOrganizationInput) HasKerberosRealm() bool`

HasKerberosRealm returns a boolean if a field has been set.

### GetKerberosServiceName

`func (o *IamUpdateOrganizationInput) GetKerberosServiceName() string`

GetKerberosServiceName returns the KerberosServiceName field if non-nil, zero value otherwise.

### GetKerberosServiceNameOk

`func (o *IamUpdateOrganizationInput) GetKerberosServiceNameOk() (*string, bool)`

GetKerberosServiceNameOk returns a tuple with the KerberosServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosServiceName

`func (o *IamUpdateOrganizationInput) SetKerberosServiceName(v string)`

SetKerberosServiceName sets KerberosServiceName field to given value.

### HasKerberosServiceName

`func (o *IamUpdateOrganizationInput) HasKerberosServiceName() bool`

HasKerberosServiceName returns a boolean if a field has been set.

### GetLanguages

`func (o *IamUpdateOrganizationInput) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *IamUpdateOrganizationInput) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *IamUpdateOrganizationInput) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *IamUpdateOrganizationInput) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLdapAttributes

`func (o *IamUpdateOrganizationInput) GetLdapAttributes() []string`

GetLdapAttributes returns the LdapAttributes field if non-nil, zero value otherwise.

### GetLdapAttributesOk

`func (o *IamUpdateOrganizationInput) GetLdapAttributesOk() (*[]string, bool)`

GetLdapAttributesOk returns a tuple with the LdapAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAttributes

`func (o *IamUpdateOrganizationInput) SetLdapAttributes(v []string)`

SetLdapAttributes sets LdapAttributes field to given value.

### HasLdapAttributes

`func (o *IamUpdateOrganizationInput) HasLdapAttributes() bool`

HasLdapAttributes returns a boolean if a field has been set.

### GetLogo

`func (o *IamUpdateOrganizationInput) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *IamUpdateOrganizationInput) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *IamUpdateOrganizationInput) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *IamUpdateOrganizationInput) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetLogoDark

`func (o *IamUpdateOrganizationInput) GetLogoDark() string`

GetLogoDark returns the LogoDark field if non-nil, zero value otherwise.

### GetLogoDarkOk

`func (o *IamUpdateOrganizationInput) GetLogoDarkOk() (*string, bool)`

GetLogoDarkOk returns a tuple with the LogoDark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoDark

`func (o *IamUpdateOrganizationInput) SetLogoDark(v string)`

SetLogoDark sets LogoDark field to given value.

### HasLogoDark

`func (o *IamUpdateOrganizationInput) HasLogoDark() bool`

HasLogoDark returns a boolean if a field has been set.

### GetMasterPassword

`func (o *IamUpdateOrganizationInput) GetMasterPassword() string`

GetMasterPassword returns the MasterPassword field if non-nil, zero value otherwise.

### GetMasterPasswordOk

`func (o *IamUpdateOrganizationInput) GetMasterPasswordOk() (*string, bool)`

GetMasterPasswordOk returns a tuple with the MasterPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterPassword

`func (o *IamUpdateOrganizationInput) SetMasterPassword(v string)`

SetMasterPassword sets MasterPassword field to given value.

### HasMasterPassword

`func (o *IamUpdateOrganizationInput) HasMasterPassword() bool`

HasMasterPassword returns a boolean if a field has been set.

### GetMasterVerificationCode

`func (o *IamUpdateOrganizationInput) GetMasterVerificationCode() string`

GetMasterVerificationCode returns the MasterVerificationCode field if non-nil, zero value otherwise.

### GetMasterVerificationCodeOk

`func (o *IamUpdateOrganizationInput) GetMasterVerificationCodeOk() (*string, bool)`

GetMasterVerificationCodeOk returns a tuple with the MasterVerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterVerificationCode

`func (o *IamUpdateOrganizationInput) SetMasterVerificationCode(v string)`

SetMasterVerificationCode sets MasterVerificationCode field to given value.

### HasMasterVerificationCode

`func (o *IamUpdateOrganizationInput) HasMasterVerificationCode() bool`

HasMasterVerificationCode returns a boolean if a field has been set.

### GetMfaItems

`func (o *IamUpdateOrganizationInput) GetMfaItems() []IamMfaItem`

GetMfaItems returns the MfaItems field if non-nil, zero value otherwise.

### GetMfaItemsOk

`func (o *IamUpdateOrganizationInput) GetMfaItemsOk() (*[]IamMfaItem, bool)`

GetMfaItemsOk returns a tuple with the MfaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaItems

`func (o *IamUpdateOrganizationInput) SetMfaItems(v []IamMfaItem)`

SetMfaItems sets MfaItems field to given value.

### HasMfaItems

`func (o *IamUpdateOrganizationInput) HasMfaItems() bool`

HasMfaItems returns a boolean if a field has been set.

### GetMfaRememberInHours

`func (o *IamUpdateOrganizationInput) GetMfaRememberInHours() int32`

GetMfaRememberInHours returns the MfaRememberInHours field if non-nil, zero value otherwise.

### GetMfaRememberInHoursOk

`func (o *IamUpdateOrganizationInput) GetMfaRememberInHoursOk() (*int32, bool)`

GetMfaRememberInHoursOk returns a tuple with the MfaRememberInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRememberInHours

`func (o *IamUpdateOrganizationInput) SetMfaRememberInHours(v int32)`

SetMfaRememberInHours sets MfaRememberInHours field to given value.

### HasMfaRememberInHours

`func (o *IamUpdateOrganizationInput) HasMfaRememberInHours() bool`

HasMfaRememberInHours returns a boolean if a field has been set.

### GetName

`func (o *IamUpdateOrganizationInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamUpdateOrganizationInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamUpdateOrganizationInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamUpdateOrganizationInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNavItems

`func (o *IamUpdateOrganizationInput) GetNavItems() []string`

GetNavItems returns the NavItems field if non-nil, zero value otherwise.

### GetNavItemsOk

`func (o *IamUpdateOrganizationInput) GetNavItemsOk() (*[]string, bool)`

GetNavItemsOk returns a tuple with the NavItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavItems

`func (o *IamUpdateOrganizationInput) SetNavItems(v []string)`

SetNavItems sets NavItems field to given value.

### HasNavItems

`func (o *IamUpdateOrganizationInput) HasNavItems() bool`

HasNavItems returns a boolean if a field has been set.

### GetOrgBalance

`func (o *IamUpdateOrganizationInput) GetOrgBalance() float32`

GetOrgBalance returns the OrgBalance field if non-nil, zero value otherwise.

### GetOrgBalanceOk

`func (o *IamUpdateOrganizationInput) GetOrgBalanceOk() (*float32, bool)`

GetOrgBalanceOk returns a tuple with the OrgBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgBalance

`func (o *IamUpdateOrganizationInput) SetOrgBalance(v float32)`

SetOrgBalance sets OrgBalance field to given value.

### HasOrgBalance

`func (o *IamUpdateOrganizationInput) HasOrgBalance() bool`

HasOrgBalance returns a boolean if a field has been set.

### GetOwner

`func (o *IamUpdateOrganizationInput) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamUpdateOrganizationInput) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamUpdateOrganizationInput) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamUpdateOrganizationInput) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPasswordExpireDays

`func (o *IamUpdateOrganizationInput) GetPasswordExpireDays() int32`

GetPasswordExpireDays returns the PasswordExpireDays field if non-nil, zero value otherwise.

### GetPasswordExpireDaysOk

`func (o *IamUpdateOrganizationInput) GetPasswordExpireDaysOk() (*int32, bool)`

GetPasswordExpireDaysOk returns a tuple with the PasswordExpireDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpireDays

`func (o *IamUpdateOrganizationInput) SetPasswordExpireDays(v int32)`

SetPasswordExpireDays sets PasswordExpireDays field to given value.

### HasPasswordExpireDays

`func (o *IamUpdateOrganizationInput) HasPasswordExpireDays() bool`

HasPasswordExpireDays returns a boolean if a field has been set.

### GetPasswordObfuscatorKey

`func (o *IamUpdateOrganizationInput) GetPasswordObfuscatorKey() string`

GetPasswordObfuscatorKey returns the PasswordObfuscatorKey field if non-nil, zero value otherwise.

### GetPasswordObfuscatorKeyOk

`func (o *IamUpdateOrganizationInput) GetPasswordObfuscatorKeyOk() (*string, bool)`

GetPasswordObfuscatorKeyOk returns a tuple with the PasswordObfuscatorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordObfuscatorKey

`func (o *IamUpdateOrganizationInput) SetPasswordObfuscatorKey(v string)`

SetPasswordObfuscatorKey sets PasswordObfuscatorKey field to given value.

### HasPasswordObfuscatorKey

`func (o *IamUpdateOrganizationInput) HasPasswordObfuscatorKey() bool`

HasPasswordObfuscatorKey returns a boolean if a field has been set.

### GetPasswordObfuscatorType

`func (o *IamUpdateOrganizationInput) GetPasswordObfuscatorType() string`

GetPasswordObfuscatorType returns the PasswordObfuscatorType field if non-nil, zero value otherwise.

### GetPasswordObfuscatorTypeOk

`func (o *IamUpdateOrganizationInput) GetPasswordObfuscatorTypeOk() (*string, bool)`

GetPasswordObfuscatorTypeOk returns a tuple with the PasswordObfuscatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordObfuscatorType

`func (o *IamUpdateOrganizationInput) SetPasswordObfuscatorType(v string)`

SetPasswordObfuscatorType sets PasswordObfuscatorType field to given value.

### HasPasswordObfuscatorType

`func (o *IamUpdateOrganizationInput) HasPasswordObfuscatorType() bool`

HasPasswordObfuscatorType returns a boolean if a field has been set.

### GetPasswordOptions

`func (o *IamUpdateOrganizationInput) GetPasswordOptions() []string`

GetPasswordOptions returns the PasswordOptions field if non-nil, zero value otherwise.

### GetPasswordOptionsOk

`func (o *IamUpdateOrganizationInput) GetPasswordOptionsOk() (*[]string, bool)`

GetPasswordOptionsOk returns a tuple with the PasswordOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordOptions

`func (o *IamUpdateOrganizationInput) SetPasswordOptions(v []string)`

SetPasswordOptions sets PasswordOptions field to given value.

### HasPasswordOptions

`func (o *IamUpdateOrganizationInput) HasPasswordOptions() bool`

HasPasswordOptions returns a boolean if a field has been set.

### GetPasswordSalt

`func (o *IamUpdateOrganizationInput) GetPasswordSalt() string`

GetPasswordSalt returns the PasswordSalt field if non-nil, zero value otherwise.

### GetPasswordSaltOk

`func (o *IamUpdateOrganizationInput) GetPasswordSaltOk() (*string, bool)`

GetPasswordSaltOk returns a tuple with the PasswordSalt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSalt

`func (o *IamUpdateOrganizationInput) SetPasswordSalt(v string)`

SetPasswordSalt sets PasswordSalt field to given value.

### HasPasswordSalt

`func (o *IamUpdateOrganizationInput) HasPasswordSalt() bool`

HasPasswordSalt returns a boolean if a field has been set.

### GetPasswordType

`func (o *IamUpdateOrganizationInput) GetPasswordType() string`

GetPasswordType returns the PasswordType field if non-nil, zero value otherwise.

### GetPasswordTypeOk

`func (o *IamUpdateOrganizationInput) GetPasswordTypeOk() (*string, bool)`

GetPasswordTypeOk returns a tuple with the PasswordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordType

`func (o *IamUpdateOrganizationInput) SetPasswordType(v string)`

SetPasswordType sets PasswordType field to given value.

### HasPasswordType

`func (o *IamUpdateOrganizationInput) HasPasswordType() bool`

HasPasswordType returns a boolean if a field has been set.

### GetTags

`func (o *IamUpdateOrganizationInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IamUpdateOrganizationInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IamUpdateOrganizationInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IamUpdateOrganizationInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThemeData

`func (o *IamUpdateOrganizationInput) GetThemeData() IamThemeData`

GetThemeData returns the ThemeData field if non-nil, zero value otherwise.

### GetThemeDataOk

`func (o *IamUpdateOrganizationInput) GetThemeDataOk() (*IamThemeData, bool)`

GetThemeDataOk returns a tuple with the ThemeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeData

`func (o *IamUpdateOrganizationInput) SetThemeData(v IamThemeData)`

SetThemeData sets ThemeData field to given value.

### HasThemeData

`func (o *IamUpdateOrganizationInput) HasThemeData() bool`

HasThemeData returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamUpdateOrganizationInput) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamUpdateOrganizationInput) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamUpdateOrganizationInput) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamUpdateOrganizationInput) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUseEmailAsUsername

`func (o *IamUpdateOrganizationInput) GetUseEmailAsUsername() bool`

GetUseEmailAsUsername returns the UseEmailAsUsername field if non-nil, zero value otherwise.

### GetUseEmailAsUsernameOk

`func (o *IamUpdateOrganizationInput) GetUseEmailAsUsernameOk() (*bool, bool)`

GetUseEmailAsUsernameOk returns a tuple with the UseEmailAsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmailAsUsername

`func (o *IamUpdateOrganizationInput) SetUseEmailAsUsername(v bool)`

SetUseEmailAsUsername sets UseEmailAsUsername field to given value.

### HasUseEmailAsUsername

`func (o *IamUpdateOrganizationInput) HasUseEmailAsUsername() bool`

HasUseEmailAsUsername returns a boolean if a field has been set.

### GetUsePermanentAvatar

`func (o *IamUpdateOrganizationInput) GetUsePermanentAvatar() bool`

GetUsePermanentAvatar returns the UsePermanentAvatar field if non-nil, zero value otherwise.

### GetUsePermanentAvatarOk

`func (o *IamUpdateOrganizationInput) GetUsePermanentAvatarOk() (*bool, bool)`

GetUsePermanentAvatarOk returns a tuple with the UsePermanentAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePermanentAvatar

`func (o *IamUpdateOrganizationInput) SetUsePermanentAvatar(v bool)`

SetUsePermanentAvatar sets UsePermanentAvatar field to given value.

### HasUsePermanentAvatar

`func (o *IamUpdateOrganizationInput) HasUsePermanentAvatar() bool`

HasUsePermanentAvatar returns a boolean if a field has been set.

### GetUserBalance

`func (o *IamUpdateOrganizationInput) GetUserBalance() float32`

GetUserBalance returns the UserBalance field if non-nil, zero value otherwise.

### GetUserBalanceOk

`func (o *IamUpdateOrganizationInput) GetUserBalanceOk() (*float32, bool)`

GetUserBalanceOk returns a tuple with the UserBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBalance

`func (o *IamUpdateOrganizationInput) SetUserBalance(v float32)`

SetUserBalance sets UserBalance field to given value.

### HasUserBalance

`func (o *IamUpdateOrganizationInput) HasUserBalance() bool`

HasUserBalance returns a boolean if a field has been set.

### GetUserNavItems

`func (o *IamUpdateOrganizationInput) GetUserNavItems() []string`

GetUserNavItems returns the UserNavItems field if non-nil, zero value otherwise.

### GetUserNavItemsOk

`func (o *IamUpdateOrganizationInput) GetUserNavItemsOk() (*[]string, bool)`

GetUserNavItemsOk returns a tuple with the UserNavItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNavItems

`func (o *IamUpdateOrganizationInput) SetUserNavItems(v []string)`

SetUserNavItems sets UserNavItems field to given value.

### HasUserNavItems

`func (o *IamUpdateOrganizationInput) HasUserNavItems() bool`

HasUserNavItems returns a boolean if a field has been set.

### GetUserTypes

`func (o *IamUpdateOrganizationInput) GetUserTypes() []string`

GetUserTypes returns the UserTypes field if non-nil, zero value otherwise.

### GetUserTypesOk

`func (o *IamUpdateOrganizationInput) GetUserTypesOk() (*[]string, bool)`

GetUserTypesOk returns a tuple with the UserTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTypes

`func (o *IamUpdateOrganizationInput) SetUserTypes(v []string)`

SetUserTypes sets UserTypes field to given value.

### HasUserTypes

`func (o *IamUpdateOrganizationInput) HasUserTypes() bool`

HasUserTypes returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *IamUpdateOrganizationInput) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *IamUpdateOrganizationInput) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *IamUpdateOrganizationInput) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *IamUpdateOrganizationInput) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetWidgetItems

`func (o *IamUpdateOrganizationInput) GetWidgetItems() []string`

GetWidgetItems returns the WidgetItems field if non-nil, zero value otherwise.

### GetWidgetItemsOk

`func (o *IamUpdateOrganizationInput) GetWidgetItemsOk() (*[]string, bool)`

GetWidgetItemsOk returns a tuple with the WidgetItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetItems

`func (o *IamUpdateOrganizationInput) SetWidgetItems(v []string)`

SetWidgetItems sets WidgetItems field to given value.

### HasWidgetItems

`func (o *IamUpdateOrganizationInput) HasWidgetItems() bool`

HasWidgetItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



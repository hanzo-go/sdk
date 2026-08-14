# IamCreateOrganizationInput

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

### NewIamCreateOrganizationInput

`func NewIamCreateOrganizationInput() *IamCreateOrganizationInput`

NewIamCreateOrganizationInput instantiates a new IamCreateOrganizationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamCreateOrganizationInputWithDefaults

`func NewIamCreateOrganizationInputWithDefaults() *IamCreateOrganizationInput`

NewIamCreateOrganizationInputWithDefaults instantiates a new IamCreateOrganizationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItems

`func (o *IamCreateOrganizationInput) GetAccountItems() []IamAccountItem`

GetAccountItems returns the AccountItems field if non-nil, zero value otherwise.

### GetAccountItemsOk

`func (o *IamCreateOrganizationInput) GetAccountItemsOk() (*[]IamAccountItem, bool)`

GetAccountItemsOk returns a tuple with the AccountItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItems

`func (o *IamCreateOrganizationInput) SetAccountItems(v []IamAccountItem)`

SetAccountItems sets AccountItems field to given value.

### HasAccountItems

`func (o *IamCreateOrganizationInput) HasAccountItems() bool`

HasAccountItems returns a boolean if a field has been set.

### GetAccountMenu

`func (o *IamCreateOrganizationInput) GetAccountMenu() string`

GetAccountMenu returns the AccountMenu field if non-nil, zero value otherwise.

### GetAccountMenuOk

`func (o *IamCreateOrganizationInput) GetAccountMenuOk() (*string, bool)`

GetAccountMenuOk returns a tuple with the AccountMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMenu

`func (o *IamCreateOrganizationInput) SetAccountMenu(v string)`

SetAccountMenu sets AccountMenu field to given value.

### HasAccountMenu

`func (o *IamCreateOrganizationInput) HasAccountMenu() bool`

HasAccountMenu returns a boolean if a field has been set.

### GetBalanceCredit

`func (o *IamCreateOrganizationInput) GetBalanceCredit() float32`

GetBalanceCredit returns the BalanceCredit field if non-nil, zero value otherwise.

### GetBalanceCreditOk

`func (o *IamCreateOrganizationInput) GetBalanceCreditOk() (*float32, bool)`

GetBalanceCreditOk returns a tuple with the BalanceCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCredit

`func (o *IamCreateOrganizationInput) SetBalanceCredit(v float32)`

SetBalanceCredit sets BalanceCredit field to given value.

### HasBalanceCredit

`func (o *IamCreateOrganizationInput) HasBalanceCredit() bool`

HasBalanceCredit returns a boolean if a field has been set.

### GetBalanceCurrency

`func (o *IamCreateOrganizationInput) GetBalanceCurrency() string`

GetBalanceCurrency returns the BalanceCurrency field if non-nil, zero value otherwise.

### GetBalanceCurrencyOk

`func (o *IamCreateOrganizationInput) GetBalanceCurrencyOk() (*string, bool)`

GetBalanceCurrencyOk returns a tuple with the BalanceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCurrency

`func (o *IamCreateOrganizationInput) SetBalanceCurrency(v string)`

SetBalanceCurrency sets BalanceCurrency field to given value.

### HasBalanceCurrency

`func (o *IamCreateOrganizationInput) HasBalanceCurrency() bool`

HasBalanceCurrency returns a boolean if a field has been set.

### GetCountryCodes

`func (o *IamCreateOrganizationInput) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *IamCreateOrganizationInput) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *IamCreateOrganizationInput) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *IamCreateOrganizationInput) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IamCreateOrganizationInput) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamCreateOrganizationInput) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamCreateOrganizationInput) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamCreateOrganizationInput) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamCreateOrganizationInput) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamCreateOrganizationInput) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamCreateOrganizationInput) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamCreateOrganizationInput) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDcrPolicy

`func (o *IamCreateOrganizationInput) GetDcrPolicy() string`

GetDcrPolicy returns the DcrPolicy field if non-nil, zero value otherwise.

### GetDcrPolicyOk

`func (o *IamCreateOrganizationInput) GetDcrPolicyOk() (*string, bool)`

GetDcrPolicyOk returns a tuple with the DcrPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcrPolicy

`func (o *IamCreateOrganizationInput) SetDcrPolicy(v string)`

SetDcrPolicy sets DcrPolicy field to given value.

### HasDcrPolicy

`func (o *IamCreateOrganizationInput) HasDcrPolicy() bool`

HasDcrPolicy returns a boolean if a field has been set.

### GetDefaultApplication

`func (o *IamCreateOrganizationInput) GetDefaultApplication() string`

GetDefaultApplication returns the DefaultApplication field if non-nil, zero value otherwise.

### GetDefaultApplicationOk

`func (o *IamCreateOrganizationInput) GetDefaultApplicationOk() (*string, bool)`

GetDefaultApplicationOk returns a tuple with the DefaultApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApplication

`func (o *IamCreateOrganizationInput) SetDefaultApplication(v string)`

SetDefaultApplication sets DefaultApplication field to given value.

### HasDefaultApplication

`func (o *IamCreateOrganizationInput) HasDefaultApplication() bool`

HasDefaultApplication returns a boolean if a field has been set.

### GetDefaultAvatar

`func (o *IamCreateOrganizationInput) GetDefaultAvatar() string`

GetDefaultAvatar returns the DefaultAvatar field if non-nil, zero value otherwise.

### GetDefaultAvatarOk

`func (o *IamCreateOrganizationInput) GetDefaultAvatarOk() (*string, bool)`

GetDefaultAvatarOk returns a tuple with the DefaultAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAvatar

`func (o *IamCreateOrganizationInput) SetDefaultAvatar(v string)`

SetDefaultAvatar sets DefaultAvatar field to given value.

### HasDefaultAvatar

`func (o *IamCreateOrganizationInput) HasDefaultAvatar() bool`

HasDefaultAvatar returns a boolean if a field has been set.

### GetDefaultPassword

`func (o *IamCreateOrganizationInput) GetDefaultPassword() string`

GetDefaultPassword returns the DefaultPassword field if non-nil, zero value otherwise.

### GetDefaultPasswordOk

`func (o *IamCreateOrganizationInput) GetDefaultPasswordOk() (*string, bool)`

GetDefaultPasswordOk returns a tuple with the DefaultPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPassword

`func (o *IamCreateOrganizationInput) SetDefaultPassword(v string)`

SetDefaultPassword sets DefaultPassword field to given value.

### HasDefaultPassword

`func (o *IamCreateOrganizationInput) HasDefaultPassword() bool`

HasDefaultPassword returns a boolean if a field has been set.

### GetDeleted

`func (o *IamCreateOrganizationInput) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamCreateOrganizationInput) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamCreateOrganizationInput) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamCreateOrganizationInput) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDisableSignin

`func (o *IamCreateOrganizationInput) GetDisableSignin() bool`

GetDisableSignin returns the DisableSignin field if non-nil, zero value otherwise.

### GetDisableSigninOk

`func (o *IamCreateOrganizationInput) GetDisableSigninOk() (*bool, bool)`

GetDisableSigninOk returns a tuple with the DisableSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSignin

`func (o *IamCreateOrganizationInput) SetDisableSignin(v bool)`

SetDisableSignin sets DisableSignin field to given value.

### HasDisableSignin

`func (o *IamCreateOrganizationInput) HasDisableSignin() bool`

HasDisableSignin returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamCreateOrganizationInput) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamCreateOrganizationInput) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamCreateOrganizationInput) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamCreateOrganizationInput) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableSoftDeletion

`func (o *IamCreateOrganizationInput) GetEnableSoftDeletion() bool`

GetEnableSoftDeletion returns the EnableSoftDeletion field if non-nil, zero value otherwise.

### GetEnableSoftDeletionOk

`func (o *IamCreateOrganizationInput) GetEnableSoftDeletionOk() (*bool, bool)`

GetEnableSoftDeletionOk returns a tuple with the EnableSoftDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSoftDeletion

`func (o *IamCreateOrganizationInput) SetEnableSoftDeletion(v bool)`

SetEnableSoftDeletion sets EnableSoftDeletion field to given value.

### HasEnableSoftDeletion

`func (o *IamCreateOrganizationInput) HasEnableSoftDeletion() bool`

HasEnableSoftDeletion returns a boolean if a field has been set.

### GetEnableTour

`func (o *IamCreateOrganizationInput) GetEnableTour() bool`

GetEnableTour returns the EnableTour field if non-nil, zero value otherwise.

### GetEnableTourOk

`func (o *IamCreateOrganizationInput) GetEnableTourOk() (*bool, bool)`

GetEnableTourOk returns a tuple with the EnableTour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTour

`func (o *IamCreateOrganizationInput) SetEnableTour(v bool)`

SetEnableTour sets EnableTour field to given value.

### HasEnableTour

`func (o *IamCreateOrganizationInput) HasEnableTour() bool`

HasEnableTour returns a boolean if a field has been set.

### GetFailedSigninFrozenTime

`func (o *IamCreateOrganizationInput) GetFailedSigninFrozenTime() int32`

GetFailedSigninFrozenTime returns the FailedSigninFrozenTime field if non-nil, zero value otherwise.

### GetFailedSigninFrozenTimeOk

`func (o *IamCreateOrganizationInput) GetFailedSigninFrozenTimeOk() (*int32, bool)`

GetFailedSigninFrozenTimeOk returns a tuple with the FailedSigninFrozenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedSigninFrozenTime

`func (o *IamCreateOrganizationInput) SetFailedSigninFrozenTime(v int32)`

SetFailedSigninFrozenTime sets FailedSigninFrozenTime field to given value.

### HasFailedSigninFrozenTime

`func (o *IamCreateOrganizationInput) HasFailedSigninFrozenTime() bool`

HasFailedSigninFrozenTime returns a boolean if a field has been set.

### GetFailedSigninLimit

`func (o *IamCreateOrganizationInput) GetFailedSigninLimit() int32`

GetFailedSigninLimit returns the FailedSigninLimit field if non-nil, zero value otherwise.

### GetFailedSigninLimitOk

`func (o *IamCreateOrganizationInput) GetFailedSigninLimitOk() (*int32, bool)`

GetFailedSigninLimitOk returns a tuple with the FailedSigninLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedSigninLimit

`func (o *IamCreateOrganizationInput) SetFailedSigninLimit(v int32)`

SetFailedSigninLimit sets FailedSigninLimit field to given value.

### HasFailedSigninLimit

`func (o *IamCreateOrganizationInput) HasFailedSigninLimit() bool`

HasFailedSigninLimit returns a boolean if a field has been set.

### GetFavicon

`func (o *IamCreateOrganizationInput) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *IamCreateOrganizationInput) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *IamCreateOrganizationInput) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *IamCreateOrganizationInput) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### GetFounder

`func (o *IamCreateOrganizationInput) GetFounder() string`

GetFounder returns the Founder field if non-nil, zero value otherwise.

### GetFounderOk

`func (o *IamCreateOrganizationInput) GetFounderOk() (*string, bool)`

GetFounderOk returns a tuple with the Founder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFounder

`func (o *IamCreateOrganizationInput) SetFounder(v string)`

SetFounder sets Founder field to given value.

### HasFounder

`func (o *IamCreateOrganizationInput) HasFounder() bool`

HasFounder returns a boolean if a field has been set.

### GetHasPrivilegeConsent

`func (o *IamCreateOrganizationInput) GetHasPrivilegeConsent() bool`

GetHasPrivilegeConsent returns the HasPrivilegeConsent field if non-nil, zero value otherwise.

### GetHasPrivilegeConsentOk

`func (o *IamCreateOrganizationInput) GetHasPrivilegeConsentOk() (*bool, bool)`

GetHasPrivilegeConsentOk returns a tuple with the HasPrivilegeConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivilegeConsent

`func (o *IamCreateOrganizationInput) SetHasPrivilegeConsent(v bool)`

SetHasPrivilegeConsent sets HasPrivilegeConsent field to given value.

### HasHasPrivilegeConsent

`func (o *IamCreateOrganizationInput) HasHasPrivilegeConsent() bool`

HasHasPrivilegeConsent returns a boolean if a field has been set.

### GetId

`func (o *IamCreateOrganizationInput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamCreateOrganizationInput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamCreateOrganizationInput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamCreateOrganizationInput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInitScore

`func (o *IamCreateOrganizationInput) GetInitScore() int32`

GetInitScore returns the InitScore field if non-nil, zero value otherwise.

### GetInitScoreOk

`func (o *IamCreateOrganizationInput) GetInitScoreOk() (*int32, bool)`

GetInitScoreOk returns a tuple with the InitScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitScore

`func (o *IamCreateOrganizationInput) SetInitScore(v int32)`

SetInitScore sets InitScore field to given value.

### HasInitScore

`func (o *IamCreateOrganizationInput) HasInitScore() bool`

HasInitScore returns a boolean if a field has been set.

### GetIpRestriction

`func (o *IamCreateOrganizationInput) GetIpRestriction() string`

GetIpRestriction returns the IpRestriction field if non-nil, zero value otherwise.

### GetIpRestrictionOk

`func (o *IamCreateOrganizationInput) GetIpRestrictionOk() (*string, bool)`

GetIpRestrictionOk returns a tuple with the IpRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRestriction

`func (o *IamCreateOrganizationInput) SetIpRestriction(v string)`

SetIpRestriction sets IpRestriction field to given value.

### HasIpRestriction

`func (o *IamCreateOrganizationInput) HasIpRestriction() bool`

HasIpRestriction returns a boolean if a field has been set.

### GetIpWhitelist

`func (o *IamCreateOrganizationInput) GetIpWhitelist() string`

GetIpWhitelist returns the IpWhitelist field if non-nil, zero value otherwise.

### GetIpWhitelistOk

`func (o *IamCreateOrganizationInput) GetIpWhitelistOk() (*string, bool)`

GetIpWhitelistOk returns a tuple with the IpWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelist

`func (o *IamCreateOrganizationInput) SetIpWhitelist(v string)`

SetIpWhitelist sets IpWhitelist field to given value.

### HasIpWhitelist

`func (o *IamCreateOrganizationInput) HasIpWhitelist() bool`

HasIpWhitelist returns a boolean if a field has been set.

### GetIsPersonal

`func (o *IamCreateOrganizationInput) GetIsPersonal() bool`

GetIsPersonal returns the IsPersonal field if non-nil, zero value otherwise.

### GetIsPersonalOk

`func (o *IamCreateOrganizationInput) GetIsPersonalOk() (*bool, bool)`

GetIsPersonalOk returns a tuple with the IsPersonal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersonal

`func (o *IamCreateOrganizationInput) SetIsPersonal(v bool)`

SetIsPersonal sets IsPersonal field to given value.

### HasIsPersonal

`func (o *IamCreateOrganizationInput) HasIsPersonal() bool`

HasIsPersonal returns a boolean if a field has been set.

### GetIsProfilePublic

`func (o *IamCreateOrganizationInput) GetIsProfilePublic() bool`

GetIsProfilePublic returns the IsProfilePublic field if non-nil, zero value otherwise.

### GetIsProfilePublicOk

`func (o *IamCreateOrganizationInput) GetIsProfilePublicOk() (*bool, bool)`

GetIsProfilePublicOk returns a tuple with the IsProfilePublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProfilePublic

`func (o *IamCreateOrganizationInput) SetIsProfilePublic(v bool)`

SetIsProfilePublic sets IsProfilePublic field to given value.

### HasIsProfilePublic

`func (o *IamCreateOrganizationInput) HasIsProfilePublic() bool`

HasIsProfilePublic returns a boolean if a field has been set.

### GetKerberosKdcHost

`func (o *IamCreateOrganizationInput) GetKerberosKdcHost() string`

GetKerberosKdcHost returns the KerberosKdcHost field if non-nil, zero value otherwise.

### GetKerberosKdcHostOk

`func (o *IamCreateOrganizationInput) GetKerberosKdcHostOk() (*string, bool)`

GetKerberosKdcHostOk returns a tuple with the KerberosKdcHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKdcHost

`func (o *IamCreateOrganizationInput) SetKerberosKdcHost(v string)`

SetKerberosKdcHost sets KerberosKdcHost field to given value.

### HasKerberosKdcHost

`func (o *IamCreateOrganizationInput) HasKerberosKdcHost() bool`

HasKerberosKdcHost returns a boolean if a field has been set.

### GetKerberosKeytab

`func (o *IamCreateOrganizationInput) GetKerberosKeytab() string`

GetKerberosKeytab returns the KerberosKeytab field if non-nil, zero value otherwise.

### GetKerberosKeytabOk

`func (o *IamCreateOrganizationInput) GetKerberosKeytabOk() (*string, bool)`

GetKerberosKeytabOk returns a tuple with the KerberosKeytab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKeytab

`func (o *IamCreateOrganizationInput) SetKerberosKeytab(v string)`

SetKerberosKeytab sets KerberosKeytab field to given value.

### HasKerberosKeytab

`func (o *IamCreateOrganizationInput) HasKerberosKeytab() bool`

HasKerberosKeytab returns a boolean if a field has been set.

### GetKerberosRealm

`func (o *IamCreateOrganizationInput) GetKerberosRealm() string`

GetKerberosRealm returns the KerberosRealm field if non-nil, zero value otherwise.

### GetKerberosRealmOk

`func (o *IamCreateOrganizationInput) GetKerberosRealmOk() (*string, bool)`

GetKerberosRealmOk returns a tuple with the KerberosRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosRealm

`func (o *IamCreateOrganizationInput) SetKerberosRealm(v string)`

SetKerberosRealm sets KerberosRealm field to given value.

### HasKerberosRealm

`func (o *IamCreateOrganizationInput) HasKerberosRealm() bool`

HasKerberosRealm returns a boolean if a field has been set.

### GetKerberosServiceName

`func (o *IamCreateOrganizationInput) GetKerberosServiceName() string`

GetKerberosServiceName returns the KerberosServiceName field if non-nil, zero value otherwise.

### GetKerberosServiceNameOk

`func (o *IamCreateOrganizationInput) GetKerberosServiceNameOk() (*string, bool)`

GetKerberosServiceNameOk returns a tuple with the KerberosServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosServiceName

`func (o *IamCreateOrganizationInput) SetKerberosServiceName(v string)`

SetKerberosServiceName sets KerberosServiceName field to given value.

### HasKerberosServiceName

`func (o *IamCreateOrganizationInput) HasKerberosServiceName() bool`

HasKerberosServiceName returns a boolean if a field has been set.

### GetLanguages

`func (o *IamCreateOrganizationInput) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *IamCreateOrganizationInput) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *IamCreateOrganizationInput) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *IamCreateOrganizationInput) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLdapAttributes

`func (o *IamCreateOrganizationInput) GetLdapAttributes() []string`

GetLdapAttributes returns the LdapAttributes field if non-nil, zero value otherwise.

### GetLdapAttributesOk

`func (o *IamCreateOrganizationInput) GetLdapAttributesOk() (*[]string, bool)`

GetLdapAttributesOk returns a tuple with the LdapAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAttributes

`func (o *IamCreateOrganizationInput) SetLdapAttributes(v []string)`

SetLdapAttributes sets LdapAttributes field to given value.

### HasLdapAttributes

`func (o *IamCreateOrganizationInput) HasLdapAttributes() bool`

HasLdapAttributes returns a boolean if a field has been set.

### GetLogo

`func (o *IamCreateOrganizationInput) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *IamCreateOrganizationInput) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *IamCreateOrganizationInput) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *IamCreateOrganizationInput) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetLogoDark

`func (o *IamCreateOrganizationInput) GetLogoDark() string`

GetLogoDark returns the LogoDark field if non-nil, zero value otherwise.

### GetLogoDarkOk

`func (o *IamCreateOrganizationInput) GetLogoDarkOk() (*string, bool)`

GetLogoDarkOk returns a tuple with the LogoDark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoDark

`func (o *IamCreateOrganizationInput) SetLogoDark(v string)`

SetLogoDark sets LogoDark field to given value.

### HasLogoDark

`func (o *IamCreateOrganizationInput) HasLogoDark() bool`

HasLogoDark returns a boolean if a field has been set.

### GetMasterPassword

`func (o *IamCreateOrganizationInput) GetMasterPassword() string`

GetMasterPassword returns the MasterPassword field if non-nil, zero value otherwise.

### GetMasterPasswordOk

`func (o *IamCreateOrganizationInput) GetMasterPasswordOk() (*string, bool)`

GetMasterPasswordOk returns a tuple with the MasterPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterPassword

`func (o *IamCreateOrganizationInput) SetMasterPassword(v string)`

SetMasterPassword sets MasterPassword field to given value.

### HasMasterPassword

`func (o *IamCreateOrganizationInput) HasMasterPassword() bool`

HasMasterPassword returns a boolean if a field has been set.

### GetMasterVerificationCode

`func (o *IamCreateOrganizationInput) GetMasterVerificationCode() string`

GetMasterVerificationCode returns the MasterVerificationCode field if non-nil, zero value otherwise.

### GetMasterVerificationCodeOk

`func (o *IamCreateOrganizationInput) GetMasterVerificationCodeOk() (*string, bool)`

GetMasterVerificationCodeOk returns a tuple with the MasterVerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterVerificationCode

`func (o *IamCreateOrganizationInput) SetMasterVerificationCode(v string)`

SetMasterVerificationCode sets MasterVerificationCode field to given value.

### HasMasterVerificationCode

`func (o *IamCreateOrganizationInput) HasMasterVerificationCode() bool`

HasMasterVerificationCode returns a boolean if a field has been set.

### GetMfaItems

`func (o *IamCreateOrganizationInput) GetMfaItems() []IamMfaItem`

GetMfaItems returns the MfaItems field if non-nil, zero value otherwise.

### GetMfaItemsOk

`func (o *IamCreateOrganizationInput) GetMfaItemsOk() (*[]IamMfaItem, bool)`

GetMfaItemsOk returns a tuple with the MfaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaItems

`func (o *IamCreateOrganizationInput) SetMfaItems(v []IamMfaItem)`

SetMfaItems sets MfaItems field to given value.

### HasMfaItems

`func (o *IamCreateOrganizationInput) HasMfaItems() bool`

HasMfaItems returns a boolean if a field has been set.

### GetMfaRememberInHours

`func (o *IamCreateOrganizationInput) GetMfaRememberInHours() int32`

GetMfaRememberInHours returns the MfaRememberInHours field if non-nil, zero value otherwise.

### GetMfaRememberInHoursOk

`func (o *IamCreateOrganizationInput) GetMfaRememberInHoursOk() (*int32, bool)`

GetMfaRememberInHoursOk returns a tuple with the MfaRememberInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRememberInHours

`func (o *IamCreateOrganizationInput) SetMfaRememberInHours(v int32)`

SetMfaRememberInHours sets MfaRememberInHours field to given value.

### HasMfaRememberInHours

`func (o *IamCreateOrganizationInput) HasMfaRememberInHours() bool`

HasMfaRememberInHours returns a boolean if a field has been set.

### GetName

`func (o *IamCreateOrganizationInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamCreateOrganizationInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamCreateOrganizationInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamCreateOrganizationInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNavItems

`func (o *IamCreateOrganizationInput) GetNavItems() []string`

GetNavItems returns the NavItems field if non-nil, zero value otherwise.

### GetNavItemsOk

`func (o *IamCreateOrganizationInput) GetNavItemsOk() (*[]string, bool)`

GetNavItemsOk returns a tuple with the NavItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavItems

`func (o *IamCreateOrganizationInput) SetNavItems(v []string)`

SetNavItems sets NavItems field to given value.

### HasNavItems

`func (o *IamCreateOrganizationInput) HasNavItems() bool`

HasNavItems returns a boolean if a field has been set.

### GetOrgBalance

`func (o *IamCreateOrganizationInput) GetOrgBalance() float32`

GetOrgBalance returns the OrgBalance field if non-nil, zero value otherwise.

### GetOrgBalanceOk

`func (o *IamCreateOrganizationInput) GetOrgBalanceOk() (*float32, bool)`

GetOrgBalanceOk returns a tuple with the OrgBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgBalance

`func (o *IamCreateOrganizationInput) SetOrgBalance(v float32)`

SetOrgBalance sets OrgBalance field to given value.

### HasOrgBalance

`func (o *IamCreateOrganizationInput) HasOrgBalance() bool`

HasOrgBalance returns a boolean if a field has been set.

### GetOwner

`func (o *IamCreateOrganizationInput) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamCreateOrganizationInput) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamCreateOrganizationInput) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamCreateOrganizationInput) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPasswordExpireDays

`func (o *IamCreateOrganizationInput) GetPasswordExpireDays() int32`

GetPasswordExpireDays returns the PasswordExpireDays field if non-nil, zero value otherwise.

### GetPasswordExpireDaysOk

`func (o *IamCreateOrganizationInput) GetPasswordExpireDaysOk() (*int32, bool)`

GetPasswordExpireDaysOk returns a tuple with the PasswordExpireDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpireDays

`func (o *IamCreateOrganizationInput) SetPasswordExpireDays(v int32)`

SetPasswordExpireDays sets PasswordExpireDays field to given value.

### HasPasswordExpireDays

`func (o *IamCreateOrganizationInput) HasPasswordExpireDays() bool`

HasPasswordExpireDays returns a boolean if a field has been set.

### GetPasswordObfuscatorKey

`func (o *IamCreateOrganizationInput) GetPasswordObfuscatorKey() string`

GetPasswordObfuscatorKey returns the PasswordObfuscatorKey field if non-nil, zero value otherwise.

### GetPasswordObfuscatorKeyOk

`func (o *IamCreateOrganizationInput) GetPasswordObfuscatorKeyOk() (*string, bool)`

GetPasswordObfuscatorKeyOk returns a tuple with the PasswordObfuscatorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordObfuscatorKey

`func (o *IamCreateOrganizationInput) SetPasswordObfuscatorKey(v string)`

SetPasswordObfuscatorKey sets PasswordObfuscatorKey field to given value.

### HasPasswordObfuscatorKey

`func (o *IamCreateOrganizationInput) HasPasswordObfuscatorKey() bool`

HasPasswordObfuscatorKey returns a boolean if a field has been set.

### GetPasswordObfuscatorType

`func (o *IamCreateOrganizationInput) GetPasswordObfuscatorType() string`

GetPasswordObfuscatorType returns the PasswordObfuscatorType field if non-nil, zero value otherwise.

### GetPasswordObfuscatorTypeOk

`func (o *IamCreateOrganizationInput) GetPasswordObfuscatorTypeOk() (*string, bool)`

GetPasswordObfuscatorTypeOk returns a tuple with the PasswordObfuscatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordObfuscatorType

`func (o *IamCreateOrganizationInput) SetPasswordObfuscatorType(v string)`

SetPasswordObfuscatorType sets PasswordObfuscatorType field to given value.

### HasPasswordObfuscatorType

`func (o *IamCreateOrganizationInput) HasPasswordObfuscatorType() bool`

HasPasswordObfuscatorType returns a boolean if a field has been set.

### GetPasswordOptions

`func (o *IamCreateOrganizationInput) GetPasswordOptions() []string`

GetPasswordOptions returns the PasswordOptions field if non-nil, zero value otherwise.

### GetPasswordOptionsOk

`func (o *IamCreateOrganizationInput) GetPasswordOptionsOk() (*[]string, bool)`

GetPasswordOptionsOk returns a tuple with the PasswordOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordOptions

`func (o *IamCreateOrganizationInput) SetPasswordOptions(v []string)`

SetPasswordOptions sets PasswordOptions field to given value.

### HasPasswordOptions

`func (o *IamCreateOrganizationInput) HasPasswordOptions() bool`

HasPasswordOptions returns a boolean if a field has been set.

### GetPasswordSalt

`func (o *IamCreateOrganizationInput) GetPasswordSalt() string`

GetPasswordSalt returns the PasswordSalt field if non-nil, zero value otherwise.

### GetPasswordSaltOk

`func (o *IamCreateOrganizationInput) GetPasswordSaltOk() (*string, bool)`

GetPasswordSaltOk returns a tuple with the PasswordSalt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSalt

`func (o *IamCreateOrganizationInput) SetPasswordSalt(v string)`

SetPasswordSalt sets PasswordSalt field to given value.

### HasPasswordSalt

`func (o *IamCreateOrganizationInput) HasPasswordSalt() bool`

HasPasswordSalt returns a boolean if a field has been set.

### GetPasswordType

`func (o *IamCreateOrganizationInput) GetPasswordType() string`

GetPasswordType returns the PasswordType field if non-nil, zero value otherwise.

### GetPasswordTypeOk

`func (o *IamCreateOrganizationInput) GetPasswordTypeOk() (*string, bool)`

GetPasswordTypeOk returns a tuple with the PasswordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordType

`func (o *IamCreateOrganizationInput) SetPasswordType(v string)`

SetPasswordType sets PasswordType field to given value.

### HasPasswordType

`func (o *IamCreateOrganizationInput) HasPasswordType() bool`

HasPasswordType returns a boolean if a field has been set.

### GetTags

`func (o *IamCreateOrganizationInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IamCreateOrganizationInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IamCreateOrganizationInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IamCreateOrganizationInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThemeData

`func (o *IamCreateOrganizationInput) GetThemeData() IamThemeData`

GetThemeData returns the ThemeData field if non-nil, zero value otherwise.

### GetThemeDataOk

`func (o *IamCreateOrganizationInput) GetThemeDataOk() (*IamThemeData, bool)`

GetThemeDataOk returns a tuple with the ThemeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeData

`func (o *IamCreateOrganizationInput) SetThemeData(v IamThemeData)`

SetThemeData sets ThemeData field to given value.

### HasThemeData

`func (o *IamCreateOrganizationInput) HasThemeData() bool`

HasThemeData returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamCreateOrganizationInput) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamCreateOrganizationInput) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamCreateOrganizationInput) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamCreateOrganizationInput) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUseEmailAsUsername

`func (o *IamCreateOrganizationInput) GetUseEmailAsUsername() bool`

GetUseEmailAsUsername returns the UseEmailAsUsername field if non-nil, zero value otherwise.

### GetUseEmailAsUsernameOk

`func (o *IamCreateOrganizationInput) GetUseEmailAsUsernameOk() (*bool, bool)`

GetUseEmailAsUsernameOk returns a tuple with the UseEmailAsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmailAsUsername

`func (o *IamCreateOrganizationInput) SetUseEmailAsUsername(v bool)`

SetUseEmailAsUsername sets UseEmailAsUsername field to given value.

### HasUseEmailAsUsername

`func (o *IamCreateOrganizationInput) HasUseEmailAsUsername() bool`

HasUseEmailAsUsername returns a boolean if a field has been set.

### GetUsePermanentAvatar

`func (o *IamCreateOrganizationInput) GetUsePermanentAvatar() bool`

GetUsePermanentAvatar returns the UsePermanentAvatar field if non-nil, zero value otherwise.

### GetUsePermanentAvatarOk

`func (o *IamCreateOrganizationInput) GetUsePermanentAvatarOk() (*bool, bool)`

GetUsePermanentAvatarOk returns a tuple with the UsePermanentAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePermanentAvatar

`func (o *IamCreateOrganizationInput) SetUsePermanentAvatar(v bool)`

SetUsePermanentAvatar sets UsePermanentAvatar field to given value.

### HasUsePermanentAvatar

`func (o *IamCreateOrganizationInput) HasUsePermanentAvatar() bool`

HasUsePermanentAvatar returns a boolean if a field has been set.

### GetUserBalance

`func (o *IamCreateOrganizationInput) GetUserBalance() float32`

GetUserBalance returns the UserBalance field if non-nil, zero value otherwise.

### GetUserBalanceOk

`func (o *IamCreateOrganizationInput) GetUserBalanceOk() (*float32, bool)`

GetUserBalanceOk returns a tuple with the UserBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBalance

`func (o *IamCreateOrganizationInput) SetUserBalance(v float32)`

SetUserBalance sets UserBalance field to given value.

### HasUserBalance

`func (o *IamCreateOrganizationInput) HasUserBalance() bool`

HasUserBalance returns a boolean if a field has been set.

### GetUserNavItems

`func (o *IamCreateOrganizationInput) GetUserNavItems() []string`

GetUserNavItems returns the UserNavItems field if non-nil, zero value otherwise.

### GetUserNavItemsOk

`func (o *IamCreateOrganizationInput) GetUserNavItemsOk() (*[]string, bool)`

GetUserNavItemsOk returns a tuple with the UserNavItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNavItems

`func (o *IamCreateOrganizationInput) SetUserNavItems(v []string)`

SetUserNavItems sets UserNavItems field to given value.

### HasUserNavItems

`func (o *IamCreateOrganizationInput) HasUserNavItems() bool`

HasUserNavItems returns a boolean if a field has been set.

### GetUserTypes

`func (o *IamCreateOrganizationInput) GetUserTypes() []string`

GetUserTypes returns the UserTypes field if non-nil, zero value otherwise.

### GetUserTypesOk

`func (o *IamCreateOrganizationInput) GetUserTypesOk() (*[]string, bool)`

GetUserTypesOk returns a tuple with the UserTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTypes

`func (o *IamCreateOrganizationInput) SetUserTypes(v []string)`

SetUserTypes sets UserTypes field to given value.

### HasUserTypes

`func (o *IamCreateOrganizationInput) HasUserTypes() bool`

HasUserTypes returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *IamCreateOrganizationInput) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *IamCreateOrganizationInput) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *IamCreateOrganizationInput) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *IamCreateOrganizationInput) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetWidgetItems

`func (o *IamCreateOrganizationInput) GetWidgetItems() []string`

GetWidgetItems returns the WidgetItems field if non-nil, zero value otherwise.

### GetWidgetItemsOk

`func (o *IamCreateOrganizationInput) GetWidgetItemsOk() (*[]string, bool)`

GetWidgetItemsOk returns a tuple with the WidgetItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetItems

`func (o *IamCreateOrganizationInput) SetWidgetItems(v []string)`

SetWidgetItems sets WidgetItems field to given value.

### HasWidgetItems

`func (o *IamCreateOrganizationInput) HasWidgetItems() bool`

HasWidgetItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



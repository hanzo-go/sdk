# IamObjectOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItems** | Pointer to [**[]IamObjectAccountItem**](IamObjectAccountItem.md) |  | [optional] 
**AccountMenu** | Pointer to **string** |  | [optional] 
**BalanceCredit** | Pointer to **float64** |  | [optional] 
**BalanceCurrency** | Pointer to **string** |  | [optional] 
**CountryCodes** | Pointer to **[]string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**DefaultApplication** | Pointer to **string** |  | [optional] 
**DefaultAvatar** | Pointer to **string** |  | [optional] 
**DefaultPassword** | Pointer to **string** |  | [optional] 
**DisableSignin** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EnableSoftDeletion** | Pointer to **bool** |  | [optional] 
**EnableTour** | Pointer to **bool** |  | [optional] 
**Favicon** | Pointer to **string** |  | [optional] 
**HasPrivilegeConsent** | Pointer to **bool** |  | [optional] 
**InitScore** | Pointer to **int64** |  | [optional] 
**IpRestriction** | Pointer to **string** |  | [optional] 
**IpWhitelist** | Pointer to **string** |  | [optional] 
**IsProfilePublic** | Pointer to **bool** |  | [optional] 
**Languages** | Pointer to **[]string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**LogoDark** | Pointer to **string** |  | [optional] 
**MasterPassword** | Pointer to **string** |  | [optional] 
**MasterVerificationCode** | Pointer to **string** |  | [optional] 
**MfaItems** | Pointer to [**[]IamObjectMfaItem**](IamObjectMfaItem.md) |  | [optional] 
**MfaRememberInHours** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NavItems** | Pointer to **[]string** |  | [optional] 
**OrgBalance** | Pointer to **float64** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**PasswordExpireDays** | Pointer to **int64** |  | [optional] 
**PasswordObfuscatorKey** | Pointer to **string** |  | [optional] 
**PasswordObfuscatorType** | Pointer to **string** |  | [optional] 
**PasswordOptions** | Pointer to **[]string** |  | [optional] 
**PasswordSalt** | Pointer to **string** |  | [optional] 
**PasswordType** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ThemeData** | Pointer to [**IamObjectThemeData**](IamObjectThemeData.md) |  | [optional] 
**UseEmailAsUsername** | Pointer to **bool** |  | [optional] 
**UserBalance** | Pointer to **float64** |  | [optional] 
**UserNavItems** | Pointer to **[]string** |  | [optional] 
**UserTypes** | Pointer to **[]string** |  | [optional] 
**WebsiteUrl** | Pointer to **string** |  | [optional] 
**WidgetItems** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIamObjectOrganization

`func NewIamObjectOrganization() *IamObjectOrganization`

NewIamObjectOrganization instantiates a new IamObjectOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectOrganizationWithDefaults

`func NewIamObjectOrganizationWithDefaults() *IamObjectOrganization`

NewIamObjectOrganizationWithDefaults instantiates a new IamObjectOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItems

`func (o *IamObjectOrganization) GetAccountItems() []IamObjectAccountItem`

GetAccountItems returns the AccountItems field if non-nil, zero value otherwise.

### GetAccountItemsOk

`func (o *IamObjectOrganization) GetAccountItemsOk() (*[]IamObjectAccountItem, bool)`

GetAccountItemsOk returns a tuple with the AccountItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItems

`func (o *IamObjectOrganization) SetAccountItems(v []IamObjectAccountItem)`

SetAccountItems sets AccountItems field to given value.

### HasAccountItems

`func (o *IamObjectOrganization) HasAccountItems() bool`

HasAccountItems returns a boolean if a field has been set.

### GetAccountMenu

`func (o *IamObjectOrganization) GetAccountMenu() string`

GetAccountMenu returns the AccountMenu field if non-nil, zero value otherwise.

### GetAccountMenuOk

`func (o *IamObjectOrganization) GetAccountMenuOk() (*string, bool)`

GetAccountMenuOk returns a tuple with the AccountMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMenu

`func (o *IamObjectOrganization) SetAccountMenu(v string)`

SetAccountMenu sets AccountMenu field to given value.

### HasAccountMenu

`func (o *IamObjectOrganization) HasAccountMenu() bool`

HasAccountMenu returns a boolean if a field has been set.

### GetBalanceCredit

`func (o *IamObjectOrganization) GetBalanceCredit() float64`

GetBalanceCredit returns the BalanceCredit field if non-nil, zero value otherwise.

### GetBalanceCreditOk

`func (o *IamObjectOrganization) GetBalanceCreditOk() (*float64, bool)`

GetBalanceCreditOk returns a tuple with the BalanceCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCredit

`func (o *IamObjectOrganization) SetBalanceCredit(v float64)`

SetBalanceCredit sets BalanceCredit field to given value.

### HasBalanceCredit

`func (o *IamObjectOrganization) HasBalanceCredit() bool`

HasBalanceCredit returns a boolean if a field has been set.

### GetBalanceCurrency

`func (o *IamObjectOrganization) GetBalanceCurrency() string`

GetBalanceCurrency returns the BalanceCurrency field if non-nil, zero value otherwise.

### GetBalanceCurrencyOk

`func (o *IamObjectOrganization) GetBalanceCurrencyOk() (*string, bool)`

GetBalanceCurrencyOk returns a tuple with the BalanceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCurrency

`func (o *IamObjectOrganization) SetBalanceCurrency(v string)`

SetBalanceCurrency sets BalanceCurrency field to given value.

### HasBalanceCurrency

`func (o *IamObjectOrganization) HasBalanceCurrency() bool`

HasBalanceCurrency returns a boolean if a field has been set.

### GetCountryCodes

`func (o *IamObjectOrganization) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *IamObjectOrganization) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *IamObjectOrganization) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *IamObjectOrganization) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamObjectOrganization) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectOrganization) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectOrganization) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectOrganization) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDefaultApplication

`func (o *IamObjectOrganization) GetDefaultApplication() string`

GetDefaultApplication returns the DefaultApplication field if non-nil, zero value otherwise.

### GetDefaultApplicationOk

`func (o *IamObjectOrganization) GetDefaultApplicationOk() (*string, bool)`

GetDefaultApplicationOk returns a tuple with the DefaultApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApplication

`func (o *IamObjectOrganization) SetDefaultApplication(v string)`

SetDefaultApplication sets DefaultApplication field to given value.

### HasDefaultApplication

`func (o *IamObjectOrganization) HasDefaultApplication() bool`

HasDefaultApplication returns a boolean if a field has been set.

### GetDefaultAvatar

`func (o *IamObjectOrganization) GetDefaultAvatar() string`

GetDefaultAvatar returns the DefaultAvatar field if non-nil, zero value otherwise.

### GetDefaultAvatarOk

`func (o *IamObjectOrganization) GetDefaultAvatarOk() (*string, bool)`

GetDefaultAvatarOk returns a tuple with the DefaultAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAvatar

`func (o *IamObjectOrganization) SetDefaultAvatar(v string)`

SetDefaultAvatar sets DefaultAvatar field to given value.

### HasDefaultAvatar

`func (o *IamObjectOrganization) HasDefaultAvatar() bool`

HasDefaultAvatar returns a boolean if a field has been set.

### GetDefaultPassword

`func (o *IamObjectOrganization) GetDefaultPassword() string`

GetDefaultPassword returns the DefaultPassword field if non-nil, zero value otherwise.

### GetDefaultPasswordOk

`func (o *IamObjectOrganization) GetDefaultPasswordOk() (*string, bool)`

GetDefaultPasswordOk returns a tuple with the DefaultPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPassword

`func (o *IamObjectOrganization) SetDefaultPassword(v string)`

SetDefaultPassword sets DefaultPassword field to given value.

### HasDefaultPassword

`func (o *IamObjectOrganization) HasDefaultPassword() bool`

HasDefaultPassword returns a boolean if a field has been set.

### GetDisableSignin

`func (o *IamObjectOrganization) GetDisableSignin() bool`

GetDisableSignin returns the DisableSignin field if non-nil, zero value otherwise.

### GetDisableSigninOk

`func (o *IamObjectOrganization) GetDisableSigninOk() (*bool, bool)`

GetDisableSigninOk returns a tuple with the DisableSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSignin

`func (o *IamObjectOrganization) SetDisableSignin(v bool)`

SetDisableSignin sets DisableSignin field to given value.

### HasDisableSignin

`func (o *IamObjectOrganization) HasDisableSignin() bool`

HasDisableSignin returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamObjectOrganization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamObjectOrganization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamObjectOrganization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamObjectOrganization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableSoftDeletion

`func (o *IamObjectOrganization) GetEnableSoftDeletion() bool`

GetEnableSoftDeletion returns the EnableSoftDeletion field if non-nil, zero value otherwise.

### GetEnableSoftDeletionOk

`func (o *IamObjectOrganization) GetEnableSoftDeletionOk() (*bool, bool)`

GetEnableSoftDeletionOk returns a tuple with the EnableSoftDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSoftDeletion

`func (o *IamObjectOrganization) SetEnableSoftDeletion(v bool)`

SetEnableSoftDeletion sets EnableSoftDeletion field to given value.

### HasEnableSoftDeletion

`func (o *IamObjectOrganization) HasEnableSoftDeletion() bool`

HasEnableSoftDeletion returns a boolean if a field has been set.

### GetEnableTour

`func (o *IamObjectOrganization) GetEnableTour() bool`

GetEnableTour returns the EnableTour field if non-nil, zero value otherwise.

### GetEnableTourOk

`func (o *IamObjectOrganization) GetEnableTourOk() (*bool, bool)`

GetEnableTourOk returns a tuple with the EnableTour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTour

`func (o *IamObjectOrganization) SetEnableTour(v bool)`

SetEnableTour sets EnableTour field to given value.

### HasEnableTour

`func (o *IamObjectOrganization) HasEnableTour() bool`

HasEnableTour returns a boolean if a field has been set.

### GetFavicon

`func (o *IamObjectOrganization) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *IamObjectOrganization) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *IamObjectOrganization) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *IamObjectOrganization) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### GetHasPrivilegeConsent

`func (o *IamObjectOrganization) GetHasPrivilegeConsent() bool`

GetHasPrivilegeConsent returns the HasPrivilegeConsent field if non-nil, zero value otherwise.

### GetHasPrivilegeConsentOk

`func (o *IamObjectOrganization) GetHasPrivilegeConsentOk() (*bool, bool)`

GetHasPrivilegeConsentOk returns a tuple with the HasPrivilegeConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivilegeConsent

`func (o *IamObjectOrganization) SetHasPrivilegeConsent(v bool)`

SetHasPrivilegeConsent sets HasPrivilegeConsent field to given value.

### HasHasPrivilegeConsent

`func (o *IamObjectOrganization) HasHasPrivilegeConsent() bool`

HasHasPrivilegeConsent returns a boolean if a field has been set.

### GetInitScore

`func (o *IamObjectOrganization) GetInitScore() int64`

GetInitScore returns the InitScore field if non-nil, zero value otherwise.

### GetInitScoreOk

`func (o *IamObjectOrganization) GetInitScoreOk() (*int64, bool)`

GetInitScoreOk returns a tuple with the InitScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitScore

`func (o *IamObjectOrganization) SetInitScore(v int64)`

SetInitScore sets InitScore field to given value.

### HasInitScore

`func (o *IamObjectOrganization) HasInitScore() bool`

HasInitScore returns a boolean if a field has been set.

### GetIpRestriction

`func (o *IamObjectOrganization) GetIpRestriction() string`

GetIpRestriction returns the IpRestriction field if non-nil, zero value otherwise.

### GetIpRestrictionOk

`func (o *IamObjectOrganization) GetIpRestrictionOk() (*string, bool)`

GetIpRestrictionOk returns a tuple with the IpRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRestriction

`func (o *IamObjectOrganization) SetIpRestriction(v string)`

SetIpRestriction sets IpRestriction field to given value.

### HasIpRestriction

`func (o *IamObjectOrganization) HasIpRestriction() bool`

HasIpRestriction returns a boolean if a field has been set.

### GetIpWhitelist

`func (o *IamObjectOrganization) GetIpWhitelist() string`

GetIpWhitelist returns the IpWhitelist field if non-nil, zero value otherwise.

### GetIpWhitelistOk

`func (o *IamObjectOrganization) GetIpWhitelistOk() (*string, bool)`

GetIpWhitelistOk returns a tuple with the IpWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelist

`func (o *IamObjectOrganization) SetIpWhitelist(v string)`

SetIpWhitelist sets IpWhitelist field to given value.

### HasIpWhitelist

`func (o *IamObjectOrganization) HasIpWhitelist() bool`

HasIpWhitelist returns a boolean if a field has been set.

### GetIsProfilePublic

`func (o *IamObjectOrganization) GetIsProfilePublic() bool`

GetIsProfilePublic returns the IsProfilePublic field if non-nil, zero value otherwise.

### GetIsProfilePublicOk

`func (o *IamObjectOrganization) GetIsProfilePublicOk() (*bool, bool)`

GetIsProfilePublicOk returns a tuple with the IsProfilePublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProfilePublic

`func (o *IamObjectOrganization) SetIsProfilePublic(v bool)`

SetIsProfilePublic sets IsProfilePublic field to given value.

### HasIsProfilePublic

`func (o *IamObjectOrganization) HasIsProfilePublic() bool`

HasIsProfilePublic returns a boolean if a field has been set.

### GetLanguages

`func (o *IamObjectOrganization) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *IamObjectOrganization) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *IamObjectOrganization) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *IamObjectOrganization) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLogo

`func (o *IamObjectOrganization) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *IamObjectOrganization) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *IamObjectOrganization) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *IamObjectOrganization) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetLogoDark

`func (o *IamObjectOrganization) GetLogoDark() string`

GetLogoDark returns the LogoDark field if non-nil, zero value otherwise.

### GetLogoDarkOk

`func (o *IamObjectOrganization) GetLogoDarkOk() (*string, bool)`

GetLogoDarkOk returns a tuple with the LogoDark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoDark

`func (o *IamObjectOrganization) SetLogoDark(v string)`

SetLogoDark sets LogoDark field to given value.

### HasLogoDark

`func (o *IamObjectOrganization) HasLogoDark() bool`

HasLogoDark returns a boolean if a field has been set.

### GetMasterPassword

`func (o *IamObjectOrganization) GetMasterPassword() string`

GetMasterPassword returns the MasterPassword field if non-nil, zero value otherwise.

### GetMasterPasswordOk

`func (o *IamObjectOrganization) GetMasterPasswordOk() (*string, bool)`

GetMasterPasswordOk returns a tuple with the MasterPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterPassword

`func (o *IamObjectOrganization) SetMasterPassword(v string)`

SetMasterPassword sets MasterPassword field to given value.

### HasMasterPassword

`func (o *IamObjectOrganization) HasMasterPassword() bool`

HasMasterPassword returns a boolean if a field has been set.

### GetMasterVerificationCode

`func (o *IamObjectOrganization) GetMasterVerificationCode() string`

GetMasterVerificationCode returns the MasterVerificationCode field if non-nil, zero value otherwise.

### GetMasterVerificationCodeOk

`func (o *IamObjectOrganization) GetMasterVerificationCodeOk() (*string, bool)`

GetMasterVerificationCodeOk returns a tuple with the MasterVerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterVerificationCode

`func (o *IamObjectOrganization) SetMasterVerificationCode(v string)`

SetMasterVerificationCode sets MasterVerificationCode field to given value.

### HasMasterVerificationCode

`func (o *IamObjectOrganization) HasMasterVerificationCode() bool`

HasMasterVerificationCode returns a boolean if a field has been set.

### GetMfaItems

`func (o *IamObjectOrganization) GetMfaItems() []IamObjectMfaItem`

GetMfaItems returns the MfaItems field if non-nil, zero value otherwise.

### GetMfaItemsOk

`func (o *IamObjectOrganization) GetMfaItemsOk() (*[]IamObjectMfaItem, bool)`

GetMfaItemsOk returns a tuple with the MfaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaItems

`func (o *IamObjectOrganization) SetMfaItems(v []IamObjectMfaItem)`

SetMfaItems sets MfaItems field to given value.

### HasMfaItems

`func (o *IamObjectOrganization) HasMfaItems() bool`

HasMfaItems returns a boolean if a field has been set.

### GetMfaRememberInHours

`func (o *IamObjectOrganization) GetMfaRememberInHours() int64`

GetMfaRememberInHours returns the MfaRememberInHours field if non-nil, zero value otherwise.

### GetMfaRememberInHoursOk

`func (o *IamObjectOrganization) GetMfaRememberInHoursOk() (*int64, bool)`

GetMfaRememberInHoursOk returns a tuple with the MfaRememberInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRememberInHours

`func (o *IamObjectOrganization) SetMfaRememberInHours(v int64)`

SetMfaRememberInHours sets MfaRememberInHours field to given value.

### HasMfaRememberInHours

`func (o *IamObjectOrganization) HasMfaRememberInHours() bool`

HasMfaRememberInHours returns a boolean if a field has been set.

### GetName

`func (o *IamObjectOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectOrganization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectOrganization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNavItems

`func (o *IamObjectOrganization) GetNavItems() []string`

GetNavItems returns the NavItems field if non-nil, zero value otherwise.

### GetNavItemsOk

`func (o *IamObjectOrganization) GetNavItemsOk() (*[]string, bool)`

GetNavItemsOk returns a tuple with the NavItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavItems

`func (o *IamObjectOrganization) SetNavItems(v []string)`

SetNavItems sets NavItems field to given value.

### HasNavItems

`func (o *IamObjectOrganization) HasNavItems() bool`

HasNavItems returns a boolean if a field has been set.

### GetOrgBalance

`func (o *IamObjectOrganization) GetOrgBalance() float64`

GetOrgBalance returns the OrgBalance field if non-nil, zero value otherwise.

### GetOrgBalanceOk

`func (o *IamObjectOrganization) GetOrgBalanceOk() (*float64, bool)`

GetOrgBalanceOk returns a tuple with the OrgBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgBalance

`func (o *IamObjectOrganization) SetOrgBalance(v float64)`

SetOrgBalance sets OrgBalance field to given value.

### HasOrgBalance

`func (o *IamObjectOrganization) HasOrgBalance() bool`

HasOrgBalance returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectOrganization) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectOrganization) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectOrganization) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectOrganization) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPasswordExpireDays

`func (o *IamObjectOrganization) GetPasswordExpireDays() int64`

GetPasswordExpireDays returns the PasswordExpireDays field if non-nil, zero value otherwise.

### GetPasswordExpireDaysOk

`func (o *IamObjectOrganization) GetPasswordExpireDaysOk() (*int64, bool)`

GetPasswordExpireDaysOk returns a tuple with the PasswordExpireDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpireDays

`func (o *IamObjectOrganization) SetPasswordExpireDays(v int64)`

SetPasswordExpireDays sets PasswordExpireDays field to given value.

### HasPasswordExpireDays

`func (o *IamObjectOrganization) HasPasswordExpireDays() bool`

HasPasswordExpireDays returns a boolean if a field has been set.

### GetPasswordObfuscatorKey

`func (o *IamObjectOrganization) GetPasswordObfuscatorKey() string`

GetPasswordObfuscatorKey returns the PasswordObfuscatorKey field if non-nil, zero value otherwise.

### GetPasswordObfuscatorKeyOk

`func (o *IamObjectOrganization) GetPasswordObfuscatorKeyOk() (*string, bool)`

GetPasswordObfuscatorKeyOk returns a tuple with the PasswordObfuscatorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordObfuscatorKey

`func (o *IamObjectOrganization) SetPasswordObfuscatorKey(v string)`

SetPasswordObfuscatorKey sets PasswordObfuscatorKey field to given value.

### HasPasswordObfuscatorKey

`func (o *IamObjectOrganization) HasPasswordObfuscatorKey() bool`

HasPasswordObfuscatorKey returns a boolean if a field has been set.

### GetPasswordObfuscatorType

`func (o *IamObjectOrganization) GetPasswordObfuscatorType() string`

GetPasswordObfuscatorType returns the PasswordObfuscatorType field if non-nil, zero value otherwise.

### GetPasswordObfuscatorTypeOk

`func (o *IamObjectOrganization) GetPasswordObfuscatorTypeOk() (*string, bool)`

GetPasswordObfuscatorTypeOk returns a tuple with the PasswordObfuscatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordObfuscatorType

`func (o *IamObjectOrganization) SetPasswordObfuscatorType(v string)`

SetPasswordObfuscatorType sets PasswordObfuscatorType field to given value.

### HasPasswordObfuscatorType

`func (o *IamObjectOrganization) HasPasswordObfuscatorType() bool`

HasPasswordObfuscatorType returns a boolean if a field has been set.

### GetPasswordOptions

`func (o *IamObjectOrganization) GetPasswordOptions() []string`

GetPasswordOptions returns the PasswordOptions field if non-nil, zero value otherwise.

### GetPasswordOptionsOk

`func (o *IamObjectOrganization) GetPasswordOptionsOk() (*[]string, bool)`

GetPasswordOptionsOk returns a tuple with the PasswordOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordOptions

`func (o *IamObjectOrganization) SetPasswordOptions(v []string)`

SetPasswordOptions sets PasswordOptions field to given value.

### HasPasswordOptions

`func (o *IamObjectOrganization) HasPasswordOptions() bool`

HasPasswordOptions returns a boolean if a field has been set.

### GetPasswordSalt

`func (o *IamObjectOrganization) GetPasswordSalt() string`

GetPasswordSalt returns the PasswordSalt field if non-nil, zero value otherwise.

### GetPasswordSaltOk

`func (o *IamObjectOrganization) GetPasswordSaltOk() (*string, bool)`

GetPasswordSaltOk returns a tuple with the PasswordSalt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSalt

`func (o *IamObjectOrganization) SetPasswordSalt(v string)`

SetPasswordSalt sets PasswordSalt field to given value.

### HasPasswordSalt

`func (o *IamObjectOrganization) HasPasswordSalt() bool`

HasPasswordSalt returns a boolean if a field has been set.

### GetPasswordType

`func (o *IamObjectOrganization) GetPasswordType() string`

GetPasswordType returns the PasswordType field if non-nil, zero value otherwise.

### GetPasswordTypeOk

`func (o *IamObjectOrganization) GetPasswordTypeOk() (*string, bool)`

GetPasswordTypeOk returns a tuple with the PasswordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordType

`func (o *IamObjectOrganization) SetPasswordType(v string)`

SetPasswordType sets PasswordType field to given value.

### HasPasswordType

`func (o *IamObjectOrganization) HasPasswordType() bool`

HasPasswordType returns a boolean if a field has been set.

### GetTags

`func (o *IamObjectOrganization) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IamObjectOrganization) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IamObjectOrganization) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IamObjectOrganization) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThemeData

`func (o *IamObjectOrganization) GetThemeData() IamObjectThemeData`

GetThemeData returns the ThemeData field if non-nil, zero value otherwise.

### GetThemeDataOk

`func (o *IamObjectOrganization) GetThemeDataOk() (*IamObjectThemeData, bool)`

GetThemeDataOk returns a tuple with the ThemeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeData

`func (o *IamObjectOrganization) SetThemeData(v IamObjectThemeData)`

SetThemeData sets ThemeData field to given value.

### HasThemeData

`func (o *IamObjectOrganization) HasThemeData() bool`

HasThemeData returns a boolean if a field has been set.

### GetUseEmailAsUsername

`func (o *IamObjectOrganization) GetUseEmailAsUsername() bool`

GetUseEmailAsUsername returns the UseEmailAsUsername field if non-nil, zero value otherwise.

### GetUseEmailAsUsernameOk

`func (o *IamObjectOrganization) GetUseEmailAsUsernameOk() (*bool, bool)`

GetUseEmailAsUsernameOk returns a tuple with the UseEmailAsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmailAsUsername

`func (o *IamObjectOrganization) SetUseEmailAsUsername(v bool)`

SetUseEmailAsUsername sets UseEmailAsUsername field to given value.

### HasUseEmailAsUsername

`func (o *IamObjectOrganization) HasUseEmailAsUsername() bool`

HasUseEmailAsUsername returns a boolean if a field has been set.

### GetUserBalance

`func (o *IamObjectOrganization) GetUserBalance() float64`

GetUserBalance returns the UserBalance field if non-nil, zero value otherwise.

### GetUserBalanceOk

`func (o *IamObjectOrganization) GetUserBalanceOk() (*float64, bool)`

GetUserBalanceOk returns a tuple with the UserBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBalance

`func (o *IamObjectOrganization) SetUserBalance(v float64)`

SetUserBalance sets UserBalance field to given value.

### HasUserBalance

`func (o *IamObjectOrganization) HasUserBalance() bool`

HasUserBalance returns a boolean if a field has been set.

### GetUserNavItems

`func (o *IamObjectOrganization) GetUserNavItems() []string`

GetUserNavItems returns the UserNavItems field if non-nil, zero value otherwise.

### GetUserNavItemsOk

`func (o *IamObjectOrganization) GetUserNavItemsOk() (*[]string, bool)`

GetUserNavItemsOk returns a tuple with the UserNavItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNavItems

`func (o *IamObjectOrganization) SetUserNavItems(v []string)`

SetUserNavItems sets UserNavItems field to given value.

### HasUserNavItems

`func (o *IamObjectOrganization) HasUserNavItems() bool`

HasUserNavItems returns a boolean if a field has been set.

### GetUserTypes

`func (o *IamObjectOrganization) GetUserTypes() []string`

GetUserTypes returns the UserTypes field if non-nil, zero value otherwise.

### GetUserTypesOk

`func (o *IamObjectOrganization) GetUserTypesOk() (*[]string, bool)`

GetUserTypesOk returns a tuple with the UserTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTypes

`func (o *IamObjectOrganization) SetUserTypes(v []string)`

SetUserTypes sets UserTypes field to given value.

### HasUserTypes

`func (o *IamObjectOrganization) HasUserTypes() bool`

HasUserTypes returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *IamObjectOrganization) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *IamObjectOrganization) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *IamObjectOrganization) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *IamObjectOrganization) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetWidgetItems

`func (o *IamObjectOrganization) GetWidgetItems() []string`

GetWidgetItems returns the WidgetItems field if non-nil, zero value otherwise.

### GetWidgetItemsOk

`func (o *IamObjectOrganization) GetWidgetItemsOk() (*[]string, bool)`

GetWidgetItemsOk returns a tuple with the WidgetItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetItems

`func (o *IamObjectOrganization) SetWidgetItems(v []string)`

SetWidgetItems sets WidgetItems field to given value.

### HasWidgetItems

`func (o *IamObjectOrganization) HasWidgetItems() bool`

HasWidgetItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ChatStartupConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppTitle** | Pointer to **string** |  | [optional] 
**SocialLogins** | Pointer to **[]string** |  | [optional] 
**EmailLoginEnabled** | Pointer to **bool** |  | [optional] 
**RegistrationEnabled** | Pointer to **bool** |  | [optional] 
**SocialLoginEnabled** | Pointer to **bool** |  | [optional] 
**PasswordResetEnabled** | Pointer to **bool** |  | [optional] 
**SharedLinksEnabled** | Pointer to **bool** |  | [optional] 
**PublicSharedLinksEnabled** | Pointer to **bool** |  | [optional] 
**ServerDomain** | Pointer to **string** |  | [optional] 
**HelpAndFaqURL** | Pointer to **string** |  | [optional] 
**GoogleLoginEnabled** | Pointer to **bool** |  | [optional] 
**GithubLoginEnabled** | Pointer to **bool** |  | [optional] 
**DiscordLoginEnabled** | Pointer to **bool** |  | [optional] 
**FacebookLoginEnabled** | Pointer to **bool** |  | [optional] 
**AppleLoginEnabled** | Pointer to **bool** |  | [optional] 
**OpenidLoginEnabled** | Pointer to **bool** |  | [optional] 
**OpenidLabel** | Pointer to **string** |  | [optional] 
**SamlLoginEnabled** | Pointer to **bool** |  | [optional] 
**Interface** | Pointer to **map[string]interface{}** |  | [optional] 
**Balance** | Pointer to **map[string]interface{}** |  | [optional] 
**ModelSpecs** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewChatStartupConfig

`func NewChatStartupConfig() *ChatStartupConfig`

NewChatStartupConfig instantiates a new ChatStartupConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatStartupConfigWithDefaults

`func NewChatStartupConfigWithDefaults() *ChatStartupConfig`

NewChatStartupConfigWithDefaults instantiates a new ChatStartupConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppTitle

`func (o *ChatStartupConfig) GetAppTitle() string`

GetAppTitle returns the AppTitle field if non-nil, zero value otherwise.

### GetAppTitleOk

`func (o *ChatStartupConfig) GetAppTitleOk() (*string, bool)`

GetAppTitleOk returns a tuple with the AppTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTitle

`func (o *ChatStartupConfig) SetAppTitle(v string)`

SetAppTitle sets AppTitle field to given value.

### HasAppTitle

`func (o *ChatStartupConfig) HasAppTitle() bool`

HasAppTitle returns a boolean if a field has been set.

### GetSocialLogins

`func (o *ChatStartupConfig) GetSocialLogins() []string`

GetSocialLogins returns the SocialLogins field if non-nil, zero value otherwise.

### GetSocialLoginsOk

`func (o *ChatStartupConfig) GetSocialLoginsOk() (*[]string, bool)`

GetSocialLoginsOk returns a tuple with the SocialLogins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialLogins

`func (o *ChatStartupConfig) SetSocialLogins(v []string)`

SetSocialLogins sets SocialLogins field to given value.

### HasSocialLogins

`func (o *ChatStartupConfig) HasSocialLogins() bool`

HasSocialLogins returns a boolean if a field has been set.

### GetEmailLoginEnabled

`func (o *ChatStartupConfig) GetEmailLoginEnabled() bool`

GetEmailLoginEnabled returns the EmailLoginEnabled field if non-nil, zero value otherwise.

### GetEmailLoginEnabledOk

`func (o *ChatStartupConfig) GetEmailLoginEnabledOk() (*bool, bool)`

GetEmailLoginEnabledOk returns a tuple with the EmailLoginEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailLoginEnabled

`func (o *ChatStartupConfig) SetEmailLoginEnabled(v bool)`

SetEmailLoginEnabled sets EmailLoginEnabled field to given value.

### HasEmailLoginEnabled

`func (o *ChatStartupConfig) HasEmailLoginEnabled() bool`

HasEmailLoginEnabled returns a boolean if a field has been set.

### GetRegistrationEnabled

`func (o *ChatStartupConfig) GetRegistrationEnabled() bool`

GetRegistrationEnabled returns the RegistrationEnabled field if non-nil, zero value otherwise.

### GetRegistrationEnabledOk

`func (o *ChatStartupConfig) GetRegistrationEnabledOk() (*bool, bool)`

GetRegistrationEnabledOk returns a tuple with the RegistrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationEnabled

`func (o *ChatStartupConfig) SetRegistrationEnabled(v bool)`

SetRegistrationEnabled sets RegistrationEnabled field to given value.

### HasRegistrationEnabled

`func (o *ChatStartupConfig) HasRegistrationEnabled() bool`

HasRegistrationEnabled returns a boolean if a field has been set.

### GetSocialLoginEnabled

`func (o *ChatStartupConfig) GetSocialLoginEnabled() bool`

GetSocialLoginEnabled returns the SocialLoginEnabled field if non-nil, zero value otherwise.

### GetSocialLoginEnabledOk

`func (o *ChatStartupConfig) GetSocialLoginEnabledOk() (*bool, bool)`

GetSocialLoginEnabledOk returns a tuple with the SocialLoginEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialLoginEnabled

`func (o *ChatStartupConfig) SetSocialLoginEnabled(v bool)`

SetSocialLoginEnabled sets SocialLoginEnabled field to given value.

### HasSocialLoginEnabled

`func (o *ChatStartupConfig) HasSocialLoginEnabled() bool`

HasSocialLoginEnabled returns a boolean if a field has been set.

### GetPasswordResetEnabled

`func (o *ChatStartupConfig) GetPasswordResetEnabled() bool`

GetPasswordResetEnabled returns the PasswordResetEnabled field if non-nil, zero value otherwise.

### GetPasswordResetEnabledOk

`func (o *ChatStartupConfig) GetPasswordResetEnabledOk() (*bool, bool)`

GetPasswordResetEnabledOk returns a tuple with the PasswordResetEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetEnabled

`func (o *ChatStartupConfig) SetPasswordResetEnabled(v bool)`

SetPasswordResetEnabled sets PasswordResetEnabled field to given value.

### HasPasswordResetEnabled

`func (o *ChatStartupConfig) HasPasswordResetEnabled() bool`

HasPasswordResetEnabled returns a boolean if a field has been set.

### GetSharedLinksEnabled

`func (o *ChatStartupConfig) GetSharedLinksEnabled() bool`

GetSharedLinksEnabled returns the SharedLinksEnabled field if non-nil, zero value otherwise.

### GetSharedLinksEnabledOk

`func (o *ChatStartupConfig) GetSharedLinksEnabledOk() (*bool, bool)`

GetSharedLinksEnabledOk returns a tuple with the SharedLinksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedLinksEnabled

`func (o *ChatStartupConfig) SetSharedLinksEnabled(v bool)`

SetSharedLinksEnabled sets SharedLinksEnabled field to given value.

### HasSharedLinksEnabled

`func (o *ChatStartupConfig) HasSharedLinksEnabled() bool`

HasSharedLinksEnabled returns a boolean if a field has been set.

### GetPublicSharedLinksEnabled

`func (o *ChatStartupConfig) GetPublicSharedLinksEnabled() bool`

GetPublicSharedLinksEnabled returns the PublicSharedLinksEnabled field if non-nil, zero value otherwise.

### GetPublicSharedLinksEnabledOk

`func (o *ChatStartupConfig) GetPublicSharedLinksEnabledOk() (*bool, bool)`

GetPublicSharedLinksEnabledOk returns a tuple with the PublicSharedLinksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicSharedLinksEnabled

`func (o *ChatStartupConfig) SetPublicSharedLinksEnabled(v bool)`

SetPublicSharedLinksEnabled sets PublicSharedLinksEnabled field to given value.

### HasPublicSharedLinksEnabled

`func (o *ChatStartupConfig) HasPublicSharedLinksEnabled() bool`

HasPublicSharedLinksEnabled returns a boolean if a field has been set.

### GetServerDomain

`func (o *ChatStartupConfig) GetServerDomain() string`

GetServerDomain returns the ServerDomain field if non-nil, zero value otherwise.

### GetServerDomainOk

`func (o *ChatStartupConfig) GetServerDomainOk() (*string, bool)`

GetServerDomainOk returns a tuple with the ServerDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDomain

`func (o *ChatStartupConfig) SetServerDomain(v string)`

SetServerDomain sets ServerDomain field to given value.

### HasServerDomain

`func (o *ChatStartupConfig) HasServerDomain() bool`

HasServerDomain returns a boolean if a field has been set.

### GetHelpAndFaqURL

`func (o *ChatStartupConfig) GetHelpAndFaqURL() string`

GetHelpAndFaqURL returns the HelpAndFaqURL field if non-nil, zero value otherwise.

### GetHelpAndFaqURLOk

`func (o *ChatStartupConfig) GetHelpAndFaqURLOk() (*string, bool)`

GetHelpAndFaqURLOk returns a tuple with the HelpAndFaqURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpAndFaqURL

`func (o *ChatStartupConfig) SetHelpAndFaqURL(v string)`

SetHelpAndFaqURL sets HelpAndFaqURL field to given value.

### HasHelpAndFaqURL

`func (o *ChatStartupConfig) HasHelpAndFaqURL() bool`

HasHelpAndFaqURL returns a boolean if a field has been set.

### GetGoogleLoginEnabled

`func (o *ChatStartupConfig) GetGoogleLoginEnabled() bool`

GetGoogleLoginEnabled returns the GoogleLoginEnabled field if non-nil, zero value otherwise.

### GetGoogleLoginEnabledOk

`func (o *ChatStartupConfig) GetGoogleLoginEnabledOk() (*bool, bool)`

GetGoogleLoginEnabledOk returns a tuple with the GoogleLoginEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleLoginEnabled

`func (o *ChatStartupConfig) SetGoogleLoginEnabled(v bool)`

SetGoogleLoginEnabled sets GoogleLoginEnabled field to given value.

### HasGoogleLoginEnabled

`func (o *ChatStartupConfig) HasGoogleLoginEnabled() bool`

HasGoogleLoginEnabled returns a boolean if a field has been set.

### GetGithubLoginEnabled

`func (o *ChatStartupConfig) GetGithubLoginEnabled() bool`

GetGithubLoginEnabled returns the GithubLoginEnabled field if non-nil, zero value otherwise.

### GetGithubLoginEnabledOk

`func (o *ChatStartupConfig) GetGithubLoginEnabledOk() (*bool, bool)`

GetGithubLoginEnabledOk returns a tuple with the GithubLoginEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubLoginEnabled

`func (o *ChatStartupConfig) SetGithubLoginEnabled(v bool)`

SetGithubLoginEnabled sets GithubLoginEnabled field to given value.

### HasGithubLoginEnabled

`func (o *ChatStartupConfig) HasGithubLoginEnabled() bool`

HasGithubLoginEnabled returns a boolean if a field has been set.

### GetDiscordLoginEnabled

`func (o *ChatStartupConfig) GetDiscordLoginEnabled() bool`

GetDiscordLoginEnabled returns the DiscordLoginEnabled field if non-nil, zero value otherwise.

### GetDiscordLoginEnabledOk

`func (o *ChatStartupConfig) GetDiscordLoginEnabledOk() (*bool, bool)`

GetDiscordLoginEnabledOk returns a tuple with the DiscordLoginEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordLoginEnabled

`func (o *ChatStartupConfig) SetDiscordLoginEnabled(v bool)`

SetDiscordLoginEnabled sets DiscordLoginEnabled field to given value.

### HasDiscordLoginEnabled

`func (o *ChatStartupConfig) HasDiscordLoginEnabled() bool`

HasDiscordLoginEnabled returns a boolean if a field has been set.

### GetFacebookLoginEnabled

`func (o *ChatStartupConfig) GetFacebookLoginEnabled() bool`

GetFacebookLoginEnabled returns the FacebookLoginEnabled field if non-nil, zero value otherwise.

### GetFacebookLoginEnabledOk

`func (o *ChatStartupConfig) GetFacebookLoginEnabledOk() (*bool, bool)`

GetFacebookLoginEnabledOk returns a tuple with the FacebookLoginEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookLoginEnabled

`func (o *ChatStartupConfig) SetFacebookLoginEnabled(v bool)`

SetFacebookLoginEnabled sets FacebookLoginEnabled field to given value.

### HasFacebookLoginEnabled

`func (o *ChatStartupConfig) HasFacebookLoginEnabled() bool`

HasFacebookLoginEnabled returns a boolean if a field has been set.

### GetAppleLoginEnabled

`func (o *ChatStartupConfig) GetAppleLoginEnabled() bool`

GetAppleLoginEnabled returns the AppleLoginEnabled field if non-nil, zero value otherwise.

### GetAppleLoginEnabledOk

`func (o *ChatStartupConfig) GetAppleLoginEnabledOk() (*bool, bool)`

GetAppleLoginEnabledOk returns a tuple with the AppleLoginEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleLoginEnabled

`func (o *ChatStartupConfig) SetAppleLoginEnabled(v bool)`

SetAppleLoginEnabled sets AppleLoginEnabled field to given value.

### HasAppleLoginEnabled

`func (o *ChatStartupConfig) HasAppleLoginEnabled() bool`

HasAppleLoginEnabled returns a boolean if a field has been set.

### GetOpenidLoginEnabled

`func (o *ChatStartupConfig) GetOpenidLoginEnabled() bool`

GetOpenidLoginEnabled returns the OpenidLoginEnabled field if non-nil, zero value otherwise.

### GetOpenidLoginEnabledOk

`func (o *ChatStartupConfig) GetOpenidLoginEnabledOk() (*bool, bool)`

GetOpenidLoginEnabledOk returns a tuple with the OpenidLoginEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenidLoginEnabled

`func (o *ChatStartupConfig) SetOpenidLoginEnabled(v bool)`

SetOpenidLoginEnabled sets OpenidLoginEnabled field to given value.

### HasOpenidLoginEnabled

`func (o *ChatStartupConfig) HasOpenidLoginEnabled() bool`

HasOpenidLoginEnabled returns a boolean if a field has been set.

### GetOpenidLabel

`func (o *ChatStartupConfig) GetOpenidLabel() string`

GetOpenidLabel returns the OpenidLabel field if non-nil, zero value otherwise.

### GetOpenidLabelOk

`func (o *ChatStartupConfig) GetOpenidLabelOk() (*string, bool)`

GetOpenidLabelOk returns a tuple with the OpenidLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenidLabel

`func (o *ChatStartupConfig) SetOpenidLabel(v string)`

SetOpenidLabel sets OpenidLabel field to given value.

### HasOpenidLabel

`func (o *ChatStartupConfig) HasOpenidLabel() bool`

HasOpenidLabel returns a boolean if a field has been set.

### GetSamlLoginEnabled

`func (o *ChatStartupConfig) GetSamlLoginEnabled() bool`

GetSamlLoginEnabled returns the SamlLoginEnabled field if non-nil, zero value otherwise.

### GetSamlLoginEnabledOk

`func (o *ChatStartupConfig) GetSamlLoginEnabledOk() (*bool, bool)`

GetSamlLoginEnabledOk returns a tuple with the SamlLoginEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlLoginEnabled

`func (o *ChatStartupConfig) SetSamlLoginEnabled(v bool)`

SetSamlLoginEnabled sets SamlLoginEnabled field to given value.

### HasSamlLoginEnabled

`func (o *ChatStartupConfig) HasSamlLoginEnabled() bool`

HasSamlLoginEnabled returns a boolean if a field has been set.

### GetInterface

`func (o *ChatStartupConfig) GetInterface() map[string]interface{}`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *ChatStartupConfig) GetInterfaceOk() (*map[string]interface{}, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *ChatStartupConfig) SetInterface(v map[string]interface{})`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *ChatStartupConfig) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetBalance

`func (o *ChatStartupConfig) GetBalance() map[string]interface{}`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *ChatStartupConfig) GetBalanceOk() (*map[string]interface{}, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *ChatStartupConfig) SetBalance(v map[string]interface{})`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *ChatStartupConfig) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetModelSpecs

`func (o *ChatStartupConfig) GetModelSpecs() map[string]interface{}`

GetModelSpecs returns the ModelSpecs field if non-nil, zero value otherwise.

### GetModelSpecsOk

`func (o *ChatStartupConfig) GetModelSpecsOk() (*map[string]interface{}, bool)`

GetModelSpecsOk returns a tuple with the ModelSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelSpecs

`func (o *ChatStartupConfig) SetModelSpecs(v map[string]interface{})`

SetModelSpecs sets ModelSpecs field to given value.

### HasModelSpecs

`func (o *ChatStartupConfig) HasModelSpecs() bool`

HasModelSpecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IamRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cert** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EnableCodeSignin** | Pointer to **bool** | EnableCodeSignin offers sign-in by an emailed or texted one-time code beside the password. A POINTER for the same reason as IsShared: a plain bool reads as false on every reconcile that says nothing and would switch the method off for every app whose caller never mentioned it. | [optional] 
**ExpireInHours** | Pointer to **float64** | ExpireInHours and RefreshExpireInHours are the application&#39;s token lifetimes. They are the ONLY declarative way to say that a refresh token must OUTLIVE its access token: with neither stated, oidc.refreshTTL clamps the refresh lifetime to the access lifetime, so the refresh_token grant the registration advertises expires at the same instant as the token it was meant to renew and can never be exercised. &#x60;hanzo-cli&#x60; sat in exactly that state — a browser re-login every hour, and a live refresh returning 401.  POINTERS, for the same reason as IsShared: a plain float would read as 0 on every reconcile that says nothing and reset a deliberate lifetime back to the default. Nil means \&quot;not stated, leave it\&quot;. | [optional] 
**GrantTypes** | Pointer to **[]string** |  | [optional] 
**IsShared** | Pointer to **bool** | IsShared declares that this application serves EVERY organization, not only the one named in Organization. It is the honest description of a brand app — hanzo-id, hanzo-chat, a brand console — whose customers each live in their own tenant: self-service onboarding moves a founder OUT of the brand org, so &#x60;user.Owner !&#x3D; app.Organization&#x60; is the steady state and the app really does serve every org. Application.ServesOrg reads it as one of the three ways to say yes.  A POINTER because omission must PRESERVE. This upsert is the operator&#39;s steady-state reconcile and most callers say nothing about sharing; a plain bool would read as false on every one of them and silently un-share an app — the same shape of accident that de-secreted apps through update-application. Nil means \&quot;not stated, leave it\&quot;; only an explicit true or false moves it. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Public** | Pointer to **bool** | Public declares a client that CANNOT hold a credential — a browser SPA, a CLI, a desktop app. It proves itself with PKCE instead, and the token endpoint treats \&quot;no stored secret\&quot; as exactly that (token.go: a secret is verified only when one is stored). Without this flag every upsert minted a secret, so a public client could never be registered at all and its browser code-&gt;token exchange 401&#39;d &#x60;invalid_client&#x60; forever. | [optional] 
**RedirectUris** | Pointer to **[]string** |  | [optional] 
**RefreshExpireInHours** | Pointer to **float64** |  | [optional] 

## Methods

### NewIamRegistration

`func NewIamRegistration() *IamRegistration`

NewIamRegistration instantiates a new IamRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamRegistrationWithDefaults

`func NewIamRegistrationWithDefaults() *IamRegistration`

NewIamRegistrationWithDefaults instantiates a new IamRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCert

`func (o *IamRegistration) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *IamRegistration) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *IamRegistration) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *IamRegistration) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetClientId

`func (o *IamRegistration) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IamRegistration) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IamRegistration) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IamRegistration) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *IamRegistration) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IamRegistration) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IamRegistration) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *IamRegistration) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamRegistration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamRegistration) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamRegistration) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamRegistration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableCodeSignin

`func (o *IamRegistration) GetEnableCodeSignin() bool`

GetEnableCodeSignin returns the EnableCodeSignin field if non-nil, zero value otherwise.

### GetEnableCodeSigninOk

`func (o *IamRegistration) GetEnableCodeSigninOk() (*bool, bool)`

GetEnableCodeSigninOk returns a tuple with the EnableCodeSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCodeSignin

`func (o *IamRegistration) SetEnableCodeSignin(v bool)`

SetEnableCodeSignin sets EnableCodeSignin field to given value.

### HasEnableCodeSignin

`func (o *IamRegistration) HasEnableCodeSignin() bool`

HasEnableCodeSignin returns a boolean if a field has been set.

### GetExpireInHours

`func (o *IamRegistration) GetExpireInHours() float64`

GetExpireInHours returns the ExpireInHours field if non-nil, zero value otherwise.

### GetExpireInHoursOk

`func (o *IamRegistration) GetExpireInHoursOk() (*float64, bool)`

GetExpireInHoursOk returns a tuple with the ExpireInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireInHours

`func (o *IamRegistration) SetExpireInHours(v float64)`

SetExpireInHours sets ExpireInHours field to given value.

### HasExpireInHours

`func (o *IamRegistration) HasExpireInHours() bool`

HasExpireInHours returns a boolean if a field has been set.

### GetGrantTypes

`func (o *IamRegistration) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *IamRegistration) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *IamRegistration) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *IamRegistration) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetIsShared

`func (o *IamRegistration) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *IamRegistration) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *IamRegistration) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *IamRegistration) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetName

`func (o *IamRegistration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamRegistration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamRegistration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamRegistration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *IamRegistration) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamRegistration) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamRegistration) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamRegistration) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPublic

`func (o *IamRegistration) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *IamRegistration) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *IamRegistration) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *IamRegistration) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRedirectUris

`func (o *IamRegistration) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *IamRegistration) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *IamRegistration) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *IamRegistration) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetRefreshExpireInHours

`func (o *IamRegistration) GetRefreshExpireInHours() float64`

GetRefreshExpireInHours returns the RefreshExpireInHours field if non-nil, zero value otherwise.

### GetRefreshExpireInHoursOk

`func (o *IamRegistration) GetRefreshExpireInHoursOk() (*float64, bool)`

GetRefreshExpireInHoursOk returns a tuple with the RefreshExpireInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshExpireInHours

`func (o *IamRegistration) SetRefreshExpireInHours(v float64)`

SetRefreshExpireInHours sets RefreshExpireInHours field to given value.

### HasRefreshExpireInHours

`func (o *IamRegistration) HasRefreshExpireInHours() bool`

HasRefreshExpireInHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



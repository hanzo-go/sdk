# IamUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | API credentials. AccessSecret / AccessSecretHash / the OAuth tokens are bearer material. AccessSecretHash MUST persist (orm stores via JSON; a json:\&quot;-\&quot; field is never saved), so it carries a real json tag and the handler&#39;s redact() strips it (and AccessSecret + the token fields) before responding. | [optional] 
**AccessSecret** | Pointer to **string** |  | [optional] 
**AccessSecretHash** | Pointer to **string** |  | [optional] 
**AccessToken** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **[]string** |  | [optional] 
**Addresses** | Pointer to [**[]IamAddress**](IamAddress.md) |  | [optional] 
**Adfs** | Pointer to **string** |  | [optional] 
**Affiliation** | Pointer to **string** |  | [optional] 
**Alipay** | Pointer to **string** |  | [optional] 
**Amazon** | Pointer to **string** |  | [optional] 
**Apple** | Pointer to **string** |  | [optional] 
**ApplicationScopes** | Pointer to [**[]IamConsentRecord**](IamConsentRecord.md) |  | [optional] 
**Auth0** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**AvatarType** | Pointer to **string** |  | [optional] 
**Azuread** | Pointer to **string** |  | [optional] 
**Azureadb2c** | Pointer to **string** |  | [optional] 
**Baidu** | Pointer to **string** |  | [optional] 
**Balance** | Pointer to **float32** | Balance mirrors v1 for lossless migration but is authoritative in Commerce (billing.hanzo.ai), not here — do not write it from IAM. | [optional] 
**BalanceCredit** | Pointer to **float32** |  | [optional] 
**BalanceCurrency** | Pointer to **string** |  | [optional] 
**Battlenet** | Pointer to **string** |  | [optional] 
**Bilibili** | Pointer to **string** |  | [optional] 
**Bio** | Pointer to **string** |  | [optional] 
**Birthday** | Pointer to **string** |  | [optional] 
**Bitbucket** | Pointer to **string** |  | [optional] 
**Box** | Pointer to **string** |  | [optional] 
**Cart** | Pointer to [**[]IamCartItem**](IamCartItem.md) |  | [optional] 
**Cloudfoundry** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedIp** | Pointer to **string** | Sign-in provenance. | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Custom** | Pointer to **string** |  | [optional] 
**Custom2** | Pointer to **string** |  | [optional] 
**Custom3** | Pointer to **string** |  | [optional] 
**Custom4** | Pointer to **string** |  | [optional] 
**Custom5** | Pointer to **string** |  | [optional] 
**Custom6** | Pointer to **string** |  | [optional] 
**Custom7** | Pointer to **string** |  | [optional] 
**Custom8** | Pointer to **string** |  | [optional] 
**Custom9** | Pointer to **string** |  | [optional] 
**Custom10** | Pointer to **string** |  | [optional] 
**Dailymotion** | Pointer to **string** |  | [optional] 
**Deezer** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedTime** | Pointer to **string** |  | [optional] 
**Digitalocean** | Pointer to **string** |  | [optional] 
**Dingtalk** | Pointer to **string** |  | [optional] 
**Discord** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** | Profile. | [optional] 
**Douyin** | Pointer to **string** |  | [optional] 
**Dropbox** | Pointer to **string** |  | [optional] 
**Education** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EmailVerified** | Pointer to **bool** |  | [optional] 
**Eveonline** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**FaceIds** | Pointer to [**[]IamFaceId**](IamFaceId.md) |  | [optional] 
**Facebook** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**Fitbit** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**Gitea** | Pointer to **string** |  | [optional] 
**Gitee** | Pointer to **string** |  | [optional] 
**Github** | Pointer to **string** | Linked federated-identity subjects, one column per connector (v1 parity). | [optional] 
**Gitlab** | Pointer to **string** |  | [optional] 
**Google** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Heroku** | Pointer to **string** |  | [optional] 
**Homepage** | Pointer to **string** |  | [optional] 
**Iam** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Id is the user&#39;s STABLE OPAQUE identifier — the value the OIDC &#x60;sub&#x60; claim carries. It is the v1 the legacy surface per-row UUID (e.g. \&quot;e7d7fda0-4c53-4508-9d35-7ec892b7e5d7\&quot;), migrated verbatim so a user&#39;s &#x60;sub&#x60; is byte-identical across the cutover: every live session, external reference, and the downstream money-path principal keyed on &#x60;sub&#x60; survive unchanged. A user minted natively in v2 is assigned a fresh UUID here on create, so the &#x60;sub&#x60; is ALWAYS a stable opaque id going forward — never the (Owner, Name) pair, which is mutable (a rename would otherwise silently reissue identity).  It is distinct from the embedded orm.Model STORAGE KEY — the value the datastore locks and looks a row up by — which is NOT (Owner, Name) for every row: a MIGRATED legacy row is stamped \&quot;owner/name\&quot; (SetId in the migrator), but a v2-native users.Create&#39;d row is NOT — Create allocates rather than pinning a key, so its storage key is a store-assigned surrogate id (a decimal string like \&quot;17847909129933610000001\&quot;). (Owner, Name) is therefore the natural/QUERY key (unique, indexed), not necessarily the storage key: resolve a row for a locked write by its REAL key (store.GetUserByName(...).Key().Encode(), which stamps both shapes — see internal/oidc updateUser), never by assuming \&quot;owner/name\&quot;. This Id is a first-class, indexed DOMAIN field; its json tag \&quot;id\&quot; dominates the promoted orm.Model &#x60;Id_&#x60; (also \&quot;id\&quot;) by shallower depth, so the persisted record&#39;s \&quot;id\&quot; is this UUID — exactly the v1 shape. A row that carries no Id (a not-yet-assigned pre-cutover user) falls back to the (Owner, Name) subject at mint; every other path resolves &#x60;sub&#x60;→user by Id. | [optional] 
**IdCard** | Pointer to **string** |  | [optional] 
**IdCardType** | Pointer to **string** |  | [optional] 
**Influxcloud** | Pointer to **string** |  | [optional] 
**Infoflow** | Pointer to **string** |  | [optional] 
**Instagram** | Pointer to **string** |  | [optional] 
**Intercom** | Pointer to **string** |  | [optional] 
**Invitation** | Pointer to **string** |  | [optional] 
**InvitationCode** | Pointer to **string** |  | [optional] 
**IpWhitelist** | Pointer to **string** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**IsDefaultAvatar** | Pointer to **bool** | State flags. | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**IsForbidden** | Pointer to **bool** |  | [optional] 
**IsOnline** | Pointer to **bool** |  | [optional] 
**IsVerified** | Pointer to **bool** |  | [optional] 
**Kakao** | Pointer to **string** |  | [optional] 
**Karma** | Pointer to **int32** |  | [optional] 
**Kwai** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Lark** | Pointer to **string** |  | [optional] 
**LastChangePasswordTime** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**LastSigninIp** | Pointer to **string** |  | [optional] 
**LastSigninTime** | Pointer to **string** |  | [optional] 
**LastSigninWrongTime** | Pointer to **string** |  | [optional] 
**Lastfm** | Pointer to **string** |  | [optional] 
**Ldap** | Pointer to **string** |  | [optional] 
**Line** | Pointer to **string** |  | [optional] 
**Linkedin** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Mailru** | Pointer to **string** |  | [optional] 
**ManagedAccounts** | Pointer to [**[]IamManagedAccount**](IamManagedAccount.md) |  | [optional] 
**Meetup** | Pointer to **string** |  | [optional] 
**MfaAccounts** | Pointer to [**[]IamMfaAccount**](IamMfaAccount.md) |  | [optional] 
**MfaEmailEnabled** | Pointer to **bool** |  | [optional] 
**MfaItems** | Pointer to [**[]IamMfaItem**](IamMfaItem.md) |  | [optional] 
**MfaPhoneEnabled** | Pointer to **bool** |  | [optional] 
**MfaPushEnabled** | Pointer to **bool** |  | [optional] 
**MfaPushProvider** | Pointer to **string** |  | [optional] 
**MfaPushReceiver** | Pointer to **string** |  | [optional] 
**MfaRadiusEnabled** | Pointer to **bool** |  | [optional] 
**MfaRadiusProvider** | Pointer to **string** |  | [optional] 
**MfaRadiusUsername** | Pointer to **string** |  | [optional] 
**MfaRememberDeadline** | Pointer to **string** |  | [optional] 
**MfaRememberDigest** | Pointer to **string** | MfaRememberDigest is the digest of the token held by the ONE browser the deadline above applies to. Without it the deadline is account-wide and \&quot;don&#39;t ask again on this browser\&quot; switches the second factor off everywhere. It is a digest, never the token, so a database dump yields nothing presentable — and it carries a REAL json tag because orm persists via json.Marshal, so &#x60;json:\&quot;-\&quot;&#x60; would never be stored (the trap PasswordHash documents above); Mask() strips it from every response instead. | [optional] 
**Microsoftonline** | Pointer to **string** |  | [optional] 
**MultiFactorAuths** | Pointer to [**[]IamMfaProps**](IamMfaProps.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Naver** | Pointer to **string** |  | [optional] 
**NeedUpdatePassword** | Pointer to **bool** |  | [optional] 
**Nextcloud** | Pointer to **string** |  | [optional] 
**Okta** | Pointer to **string** |  | [optional] 
**Onedrive** | Pointer to **string** |  | [optional] 
**OriginalRefreshToken** | Pointer to **string** |  | [optional] 
**OriginalToken** | Pointer to **string** |  | [optional] 
**Oura** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** | Identity / tenancy. (Owner, Name) is the natural key. | [optional] 
**PasswordHash** | Pointer to **string** | Credential material. PasswordHash is a one-way bcrypt digest and is verify-only. It MUST be persisted (orm serializes the entity to its JSON data column, so a json:\&quot;-\&quot; field would never be stored — that silently broke login), so it carries a real json tag; the users API redact() strips it (and every other secret) from every response. PasswordType and PasswordSalt describe the digest scheme so rows hashed under the legacy argon2id scheme can still be verified and lazily re-hashed to bcrypt. | [optional] 
**PasswordSalt** | Pointer to **string** |  | [optional] 
**PasswordType** | Pointer to **string** |  | [optional] 
**Patreon** | Pointer to **string** |  | [optional] 
**Paypal** | Pointer to **string** |  | [optional] 
**PermanentAvatar** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**[]IamPermission**](IamPermission.md) |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**PreHash** | Pointer to **string** |  | [optional] 
**PreferredMfaType** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**Qq** | Pointer to **string** |  | [optional] 
**Ranking** | Pointer to **int32** |  | [optional] 
**RealName** | Pointer to **string** |  | [optional] 
**RecoveryCodes** | Pointer to **[]string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**RegisterSource** | Pointer to **string** |  | [optional] 
**RegisterType** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to [**[]IamRole**](IamRole.md) | Authorization attachments. Roles and Permissions are computed on read from the authz store and carried here for API parity with v1. | [optional] 
**Salesforce** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**Shopify** | Pointer to **string** |  | [optional] 
**SigninWrongTimes** | Pointer to **int32** |  | [optional] 
**SignupApplication** | Pointer to **string** |  | [optional] 
**Slack** | Pointer to **string** |  | [optional] 
**Soundcloud** | Pointer to **string** |  | [optional] 
**Spotify** | Pointer to **string** |  | [optional] 
**Steam** | Pointer to **string** |  | [optional] 
**Strava** | Pointer to **string** |  | [optional] 
**Stripe** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Telegram** | Pointer to **string** |  | [optional] 
**Tiktok** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TotpSecret** | Pointer to **string** |  | [optional] 
**Tumblr** | Pointer to **string** |  | [optional] 
**Twitch** | Pointer to **string** |  | [optional] 
**Twitter** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Typetalk** | Pointer to **string** |  | [optional] 
**Uber** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 
**VerificationCode** | Pointer to **string** |  | [optional] 
**Vk** | Pointer to **string** |  | [optional] 
**WebauthnCredentials** | Pointer to **[]interface{}** | Multi-factor authentication. TotpSecret and RecoveryCodes are secret verify-only material — the handler strips them from every response. WebauthnCredentials is carried as raw JSON here for lossless migration; the typed passkey model is the sibling WebauthnCredential entity. | [optional] 
**Wechat** | Pointer to **string** |  | [optional] 
**Wecom** | Pointer to **string** |  | [optional] 
**Weibo** | Pointer to **string** |  | [optional] 
**Wepay** | Pointer to **string** |  | [optional] 
**Xero** | Pointer to **string** |  | [optional] 
**Yahoo** | Pointer to **string** |  | [optional] 
**Yammer** | Pointer to **string** |  | [optional] 
**Yandex** | Pointer to **string** |  | [optional] 
**Zoom** | Pointer to **string** |  | [optional] 

## Methods

### NewIamUser

`func NewIamUser() *IamUser`

NewIamUser instantiates a new IamUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUserWithDefaults

`func NewIamUserWithDefaults() *IamUser`

NewIamUserWithDefaults instantiates a new IamUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *IamUser) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *IamUser) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *IamUser) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *IamUser) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessSecret

`func (o *IamUser) GetAccessSecret() string`

GetAccessSecret returns the AccessSecret field if non-nil, zero value otherwise.

### GetAccessSecretOk

`func (o *IamUser) GetAccessSecretOk() (*string, bool)`

GetAccessSecretOk returns a tuple with the AccessSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSecret

`func (o *IamUser) SetAccessSecret(v string)`

SetAccessSecret sets AccessSecret field to given value.

### HasAccessSecret

`func (o *IamUser) HasAccessSecret() bool`

HasAccessSecret returns a boolean if a field has been set.

### GetAccessSecretHash

`func (o *IamUser) GetAccessSecretHash() string`

GetAccessSecretHash returns the AccessSecretHash field if non-nil, zero value otherwise.

### GetAccessSecretHashOk

`func (o *IamUser) GetAccessSecretHashOk() (*string, bool)`

GetAccessSecretHashOk returns a tuple with the AccessSecretHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSecretHash

`func (o *IamUser) SetAccessSecretHash(v string)`

SetAccessSecretHash sets AccessSecretHash field to given value.

### HasAccessSecretHash

`func (o *IamUser) HasAccessSecretHash() bool`

HasAccessSecretHash returns a boolean if a field has been set.

### GetAccessToken

`func (o *IamUser) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *IamUser) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *IamUser) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *IamUser) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAddress

`func (o *IamUser) GetAddress() []string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IamUser) GetAddressOk() (*[]string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IamUser) SetAddress(v []string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IamUser) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddresses

`func (o *IamUser) GetAddresses() []IamAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *IamUser) GetAddressesOk() (*[]IamAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *IamUser) SetAddresses(v []IamAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *IamUser) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAdfs

`func (o *IamUser) GetAdfs() string`

GetAdfs returns the Adfs field if non-nil, zero value otherwise.

### GetAdfsOk

`func (o *IamUser) GetAdfsOk() (*string, bool)`

GetAdfsOk returns a tuple with the Adfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdfs

`func (o *IamUser) SetAdfs(v string)`

SetAdfs sets Adfs field to given value.

### HasAdfs

`func (o *IamUser) HasAdfs() bool`

HasAdfs returns a boolean if a field has been set.

### GetAffiliation

`func (o *IamUser) GetAffiliation() string`

GetAffiliation returns the Affiliation field if non-nil, zero value otherwise.

### GetAffiliationOk

`func (o *IamUser) GetAffiliationOk() (*string, bool)`

GetAffiliationOk returns a tuple with the Affiliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliation

`func (o *IamUser) SetAffiliation(v string)`

SetAffiliation sets Affiliation field to given value.

### HasAffiliation

`func (o *IamUser) HasAffiliation() bool`

HasAffiliation returns a boolean if a field has been set.

### GetAlipay

`func (o *IamUser) GetAlipay() string`

GetAlipay returns the Alipay field if non-nil, zero value otherwise.

### GetAlipayOk

`func (o *IamUser) GetAlipayOk() (*string, bool)`

GetAlipayOk returns a tuple with the Alipay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlipay

`func (o *IamUser) SetAlipay(v string)`

SetAlipay sets Alipay field to given value.

### HasAlipay

`func (o *IamUser) HasAlipay() bool`

HasAlipay returns a boolean if a field has been set.

### GetAmazon

`func (o *IamUser) GetAmazon() string`

GetAmazon returns the Amazon field if non-nil, zero value otherwise.

### GetAmazonOk

`func (o *IamUser) GetAmazonOk() (*string, bool)`

GetAmazonOk returns a tuple with the Amazon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazon

`func (o *IamUser) SetAmazon(v string)`

SetAmazon sets Amazon field to given value.

### HasAmazon

`func (o *IamUser) HasAmazon() bool`

HasAmazon returns a boolean if a field has been set.

### GetApple

`func (o *IamUser) GetApple() string`

GetApple returns the Apple field if non-nil, zero value otherwise.

### GetAppleOk

`func (o *IamUser) GetAppleOk() (*string, bool)`

GetAppleOk returns a tuple with the Apple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApple

`func (o *IamUser) SetApple(v string)`

SetApple sets Apple field to given value.

### HasApple

`func (o *IamUser) HasApple() bool`

HasApple returns a boolean if a field has been set.

### GetApplicationScopes

`func (o *IamUser) GetApplicationScopes() []IamConsentRecord`

GetApplicationScopes returns the ApplicationScopes field if non-nil, zero value otherwise.

### GetApplicationScopesOk

`func (o *IamUser) GetApplicationScopesOk() (*[]IamConsentRecord, bool)`

GetApplicationScopesOk returns a tuple with the ApplicationScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationScopes

`func (o *IamUser) SetApplicationScopes(v []IamConsentRecord)`

SetApplicationScopes sets ApplicationScopes field to given value.

### HasApplicationScopes

`func (o *IamUser) HasApplicationScopes() bool`

HasApplicationScopes returns a boolean if a field has been set.

### GetAuth0

`func (o *IamUser) GetAuth0() string`

GetAuth0 returns the Auth0 field if non-nil, zero value otherwise.

### GetAuth0Ok

`func (o *IamUser) GetAuth0Ok() (*string, bool)`

GetAuth0Ok returns a tuple with the Auth0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth0

`func (o *IamUser) SetAuth0(v string)`

SetAuth0 sets Auth0 field to given value.

### HasAuth0

`func (o *IamUser) HasAuth0() bool`

HasAuth0 returns a boolean if a field has been set.

### GetAvatar

`func (o *IamUser) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *IamUser) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *IamUser) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *IamUser) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetAvatarType

`func (o *IamUser) GetAvatarType() string`

GetAvatarType returns the AvatarType field if non-nil, zero value otherwise.

### GetAvatarTypeOk

`func (o *IamUser) GetAvatarTypeOk() (*string, bool)`

GetAvatarTypeOk returns a tuple with the AvatarType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarType

`func (o *IamUser) SetAvatarType(v string)`

SetAvatarType sets AvatarType field to given value.

### HasAvatarType

`func (o *IamUser) HasAvatarType() bool`

HasAvatarType returns a boolean if a field has been set.

### GetAzuread

`func (o *IamUser) GetAzuread() string`

GetAzuread returns the Azuread field if non-nil, zero value otherwise.

### GetAzureadOk

`func (o *IamUser) GetAzureadOk() (*string, bool)`

GetAzureadOk returns a tuple with the Azuread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzuread

`func (o *IamUser) SetAzuread(v string)`

SetAzuread sets Azuread field to given value.

### HasAzuread

`func (o *IamUser) HasAzuread() bool`

HasAzuread returns a boolean if a field has been set.

### GetAzureadb2c

`func (o *IamUser) GetAzureadb2c() string`

GetAzureadb2c returns the Azureadb2c field if non-nil, zero value otherwise.

### GetAzureadb2cOk

`func (o *IamUser) GetAzureadb2cOk() (*string, bool)`

GetAzureadb2cOk returns a tuple with the Azureadb2c field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureadb2c

`func (o *IamUser) SetAzureadb2c(v string)`

SetAzureadb2c sets Azureadb2c field to given value.

### HasAzureadb2c

`func (o *IamUser) HasAzureadb2c() bool`

HasAzureadb2c returns a boolean if a field has been set.

### GetBaidu

`func (o *IamUser) GetBaidu() string`

GetBaidu returns the Baidu field if non-nil, zero value otherwise.

### GetBaiduOk

`func (o *IamUser) GetBaiduOk() (*string, bool)`

GetBaiduOk returns a tuple with the Baidu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaidu

`func (o *IamUser) SetBaidu(v string)`

SetBaidu sets Baidu field to given value.

### HasBaidu

`func (o *IamUser) HasBaidu() bool`

HasBaidu returns a boolean if a field has been set.

### GetBalance

`func (o *IamUser) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *IamUser) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *IamUser) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *IamUser) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetBalanceCredit

`func (o *IamUser) GetBalanceCredit() float32`

GetBalanceCredit returns the BalanceCredit field if non-nil, zero value otherwise.

### GetBalanceCreditOk

`func (o *IamUser) GetBalanceCreditOk() (*float32, bool)`

GetBalanceCreditOk returns a tuple with the BalanceCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCredit

`func (o *IamUser) SetBalanceCredit(v float32)`

SetBalanceCredit sets BalanceCredit field to given value.

### HasBalanceCredit

`func (o *IamUser) HasBalanceCredit() bool`

HasBalanceCredit returns a boolean if a field has been set.

### GetBalanceCurrency

`func (o *IamUser) GetBalanceCurrency() string`

GetBalanceCurrency returns the BalanceCurrency field if non-nil, zero value otherwise.

### GetBalanceCurrencyOk

`func (o *IamUser) GetBalanceCurrencyOk() (*string, bool)`

GetBalanceCurrencyOk returns a tuple with the BalanceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCurrency

`func (o *IamUser) SetBalanceCurrency(v string)`

SetBalanceCurrency sets BalanceCurrency field to given value.

### HasBalanceCurrency

`func (o *IamUser) HasBalanceCurrency() bool`

HasBalanceCurrency returns a boolean if a field has been set.

### GetBattlenet

`func (o *IamUser) GetBattlenet() string`

GetBattlenet returns the Battlenet field if non-nil, zero value otherwise.

### GetBattlenetOk

`func (o *IamUser) GetBattlenetOk() (*string, bool)`

GetBattlenetOk returns a tuple with the Battlenet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBattlenet

`func (o *IamUser) SetBattlenet(v string)`

SetBattlenet sets Battlenet field to given value.

### HasBattlenet

`func (o *IamUser) HasBattlenet() bool`

HasBattlenet returns a boolean if a field has been set.

### GetBilibili

`func (o *IamUser) GetBilibili() string`

GetBilibili returns the Bilibili field if non-nil, zero value otherwise.

### GetBilibiliOk

`func (o *IamUser) GetBilibiliOk() (*string, bool)`

GetBilibiliOk returns a tuple with the Bilibili field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilibili

`func (o *IamUser) SetBilibili(v string)`

SetBilibili sets Bilibili field to given value.

### HasBilibili

`func (o *IamUser) HasBilibili() bool`

HasBilibili returns a boolean if a field has been set.

### GetBio

`func (o *IamUser) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *IamUser) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *IamUser) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *IamUser) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetBirthday

`func (o *IamUser) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *IamUser) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *IamUser) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *IamUser) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetBitbucket

`func (o *IamUser) GetBitbucket() string`

GetBitbucket returns the Bitbucket field if non-nil, zero value otherwise.

### GetBitbucketOk

`func (o *IamUser) GetBitbucketOk() (*string, bool)`

GetBitbucketOk returns a tuple with the Bitbucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitbucket

`func (o *IamUser) SetBitbucket(v string)`

SetBitbucket sets Bitbucket field to given value.

### HasBitbucket

`func (o *IamUser) HasBitbucket() bool`

HasBitbucket returns a boolean if a field has been set.

### GetBox

`func (o *IamUser) GetBox() string`

GetBox returns the Box field if non-nil, zero value otherwise.

### GetBoxOk

`func (o *IamUser) GetBoxOk() (*string, bool)`

GetBoxOk returns a tuple with the Box field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBox

`func (o *IamUser) SetBox(v string)`

SetBox sets Box field to given value.

### HasBox

`func (o *IamUser) HasBox() bool`

HasBox returns a boolean if a field has been set.

### GetCart

`func (o *IamUser) GetCart() []IamCartItem`

GetCart returns the Cart field if non-nil, zero value otherwise.

### GetCartOk

`func (o *IamUser) GetCartOk() (*[]IamCartItem, bool)`

GetCartOk returns a tuple with the Cart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCart

`func (o *IamUser) SetCart(v []IamCartItem)`

SetCart sets Cart field to given value.

### HasCart

`func (o *IamUser) HasCart() bool`

HasCart returns a boolean if a field has been set.

### GetCloudfoundry

`func (o *IamUser) GetCloudfoundry() string`

GetCloudfoundry returns the Cloudfoundry field if non-nil, zero value otherwise.

### GetCloudfoundryOk

`func (o *IamUser) GetCloudfoundryOk() (*string, bool)`

GetCloudfoundryOk returns a tuple with the Cloudfoundry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudfoundry

`func (o *IamUser) SetCloudfoundry(v string)`

SetCloudfoundry sets Cloudfoundry field to given value.

### HasCloudfoundry

`func (o *IamUser) HasCloudfoundry() bool`

HasCloudfoundry returns a boolean if a field has been set.

### GetCountryCode

`func (o *IamUser) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *IamUser) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *IamUser) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *IamUser) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IamUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedIp

`func (o *IamUser) GetCreatedIp() string`

GetCreatedIp returns the CreatedIp field if non-nil, zero value otherwise.

### GetCreatedIpOk

`func (o *IamUser) GetCreatedIpOk() (*string, bool)`

GetCreatedIpOk returns a tuple with the CreatedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedIp

`func (o *IamUser) SetCreatedIp(v string)`

SetCreatedIp sets CreatedIp field to given value.

### HasCreatedIp

`func (o *IamUser) HasCreatedIp() bool`

HasCreatedIp returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamUser) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamUser) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamUser) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamUser) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrency

`func (o *IamUser) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *IamUser) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *IamUser) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *IamUser) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustom

`func (o *IamUser) GetCustom() string`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *IamUser) GetCustomOk() (*string, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *IamUser) SetCustom(v string)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *IamUser) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetCustom2

`func (o *IamUser) GetCustom2() string`

GetCustom2 returns the Custom2 field if non-nil, zero value otherwise.

### GetCustom2Ok

`func (o *IamUser) GetCustom2Ok() (*string, bool)`

GetCustom2Ok returns a tuple with the Custom2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom2

`func (o *IamUser) SetCustom2(v string)`

SetCustom2 sets Custom2 field to given value.

### HasCustom2

`func (o *IamUser) HasCustom2() bool`

HasCustom2 returns a boolean if a field has been set.

### GetCustom3

`func (o *IamUser) GetCustom3() string`

GetCustom3 returns the Custom3 field if non-nil, zero value otherwise.

### GetCustom3Ok

`func (o *IamUser) GetCustom3Ok() (*string, bool)`

GetCustom3Ok returns a tuple with the Custom3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom3

`func (o *IamUser) SetCustom3(v string)`

SetCustom3 sets Custom3 field to given value.

### HasCustom3

`func (o *IamUser) HasCustom3() bool`

HasCustom3 returns a boolean if a field has been set.

### GetCustom4

`func (o *IamUser) GetCustom4() string`

GetCustom4 returns the Custom4 field if non-nil, zero value otherwise.

### GetCustom4Ok

`func (o *IamUser) GetCustom4Ok() (*string, bool)`

GetCustom4Ok returns a tuple with the Custom4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom4

`func (o *IamUser) SetCustom4(v string)`

SetCustom4 sets Custom4 field to given value.

### HasCustom4

`func (o *IamUser) HasCustom4() bool`

HasCustom4 returns a boolean if a field has been set.

### GetCustom5

`func (o *IamUser) GetCustom5() string`

GetCustom5 returns the Custom5 field if non-nil, zero value otherwise.

### GetCustom5Ok

`func (o *IamUser) GetCustom5Ok() (*string, bool)`

GetCustom5Ok returns a tuple with the Custom5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom5

`func (o *IamUser) SetCustom5(v string)`

SetCustom5 sets Custom5 field to given value.

### HasCustom5

`func (o *IamUser) HasCustom5() bool`

HasCustom5 returns a boolean if a field has been set.

### GetCustom6

`func (o *IamUser) GetCustom6() string`

GetCustom6 returns the Custom6 field if non-nil, zero value otherwise.

### GetCustom6Ok

`func (o *IamUser) GetCustom6Ok() (*string, bool)`

GetCustom6Ok returns a tuple with the Custom6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom6

`func (o *IamUser) SetCustom6(v string)`

SetCustom6 sets Custom6 field to given value.

### HasCustom6

`func (o *IamUser) HasCustom6() bool`

HasCustom6 returns a boolean if a field has been set.

### GetCustom7

`func (o *IamUser) GetCustom7() string`

GetCustom7 returns the Custom7 field if non-nil, zero value otherwise.

### GetCustom7Ok

`func (o *IamUser) GetCustom7Ok() (*string, bool)`

GetCustom7Ok returns a tuple with the Custom7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom7

`func (o *IamUser) SetCustom7(v string)`

SetCustom7 sets Custom7 field to given value.

### HasCustom7

`func (o *IamUser) HasCustom7() bool`

HasCustom7 returns a boolean if a field has been set.

### GetCustom8

`func (o *IamUser) GetCustom8() string`

GetCustom8 returns the Custom8 field if non-nil, zero value otherwise.

### GetCustom8Ok

`func (o *IamUser) GetCustom8Ok() (*string, bool)`

GetCustom8Ok returns a tuple with the Custom8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom8

`func (o *IamUser) SetCustom8(v string)`

SetCustom8 sets Custom8 field to given value.

### HasCustom8

`func (o *IamUser) HasCustom8() bool`

HasCustom8 returns a boolean if a field has been set.

### GetCustom9

`func (o *IamUser) GetCustom9() string`

GetCustom9 returns the Custom9 field if non-nil, zero value otherwise.

### GetCustom9Ok

`func (o *IamUser) GetCustom9Ok() (*string, bool)`

GetCustom9Ok returns a tuple with the Custom9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom9

`func (o *IamUser) SetCustom9(v string)`

SetCustom9 sets Custom9 field to given value.

### HasCustom9

`func (o *IamUser) HasCustom9() bool`

HasCustom9 returns a boolean if a field has been set.

### GetCustom10

`func (o *IamUser) GetCustom10() string`

GetCustom10 returns the Custom10 field if non-nil, zero value otherwise.

### GetCustom10Ok

`func (o *IamUser) GetCustom10Ok() (*string, bool)`

GetCustom10Ok returns a tuple with the Custom10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom10

`func (o *IamUser) SetCustom10(v string)`

SetCustom10 sets Custom10 field to given value.

### HasCustom10

`func (o *IamUser) HasCustom10() bool`

HasCustom10 returns a boolean if a field has been set.

### GetDailymotion

`func (o *IamUser) GetDailymotion() string`

GetDailymotion returns the Dailymotion field if non-nil, zero value otherwise.

### GetDailymotionOk

`func (o *IamUser) GetDailymotionOk() (*string, bool)`

GetDailymotionOk returns a tuple with the Dailymotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailymotion

`func (o *IamUser) SetDailymotion(v string)`

SetDailymotion sets Dailymotion field to given value.

### HasDailymotion

`func (o *IamUser) HasDailymotion() bool`

HasDailymotion returns a boolean if a field has been set.

### GetDeezer

`func (o *IamUser) GetDeezer() string`

GetDeezer returns the Deezer field if non-nil, zero value otherwise.

### GetDeezerOk

`func (o *IamUser) GetDeezerOk() (*string, bool)`

GetDeezerOk returns a tuple with the Deezer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeezer

`func (o *IamUser) SetDeezer(v string)`

SetDeezer sets Deezer field to given value.

### HasDeezer

`func (o *IamUser) HasDeezer() bool`

HasDeezer returns a boolean if a field has been set.

### GetDeleted

`func (o *IamUser) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamUser) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamUser) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamUser) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedTime

`func (o *IamUser) GetDeletedTime() string`

GetDeletedTime returns the DeletedTime field if non-nil, zero value otherwise.

### GetDeletedTimeOk

`func (o *IamUser) GetDeletedTimeOk() (*string, bool)`

GetDeletedTimeOk returns a tuple with the DeletedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedTime

`func (o *IamUser) SetDeletedTime(v string)`

SetDeletedTime sets DeletedTime field to given value.

### HasDeletedTime

`func (o *IamUser) HasDeletedTime() bool`

HasDeletedTime returns a boolean if a field has been set.

### GetDigitalocean

`func (o *IamUser) GetDigitalocean() string`

GetDigitalocean returns the Digitalocean field if non-nil, zero value otherwise.

### GetDigitaloceanOk

`func (o *IamUser) GetDigitaloceanOk() (*string, bool)`

GetDigitaloceanOk returns a tuple with the Digitalocean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalocean

`func (o *IamUser) SetDigitalocean(v string)`

SetDigitalocean sets Digitalocean field to given value.

### HasDigitalocean

`func (o *IamUser) HasDigitalocean() bool`

HasDigitalocean returns a boolean if a field has been set.

### GetDingtalk

`func (o *IamUser) GetDingtalk() string`

GetDingtalk returns the Dingtalk field if non-nil, zero value otherwise.

### GetDingtalkOk

`func (o *IamUser) GetDingtalkOk() (*string, bool)`

GetDingtalkOk returns a tuple with the Dingtalk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDingtalk

`func (o *IamUser) SetDingtalk(v string)`

SetDingtalk sets Dingtalk field to given value.

### HasDingtalk

`func (o *IamUser) HasDingtalk() bool`

HasDingtalk returns a boolean if a field has been set.

### GetDiscord

`func (o *IamUser) GetDiscord() string`

GetDiscord returns the Discord field if non-nil, zero value otherwise.

### GetDiscordOk

`func (o *IamUser) GetDiscordOk() (*string, bool)`

GetDiscordOk returns a tuple with the Discord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscord

`func (o *IamUser) SetDiscord(v string)`

SetDiscord sets Discord field to given value.

### HasDiscord

`func (o *IamUser) HasDiscord() bool`

HasDiscord returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDouyin

`func (o *IamUser) GetDouyin() string`

GetDouyin returns the Douyin field if non-nil, zero value otherwise.

### GetDouyinOk

`func (o *IamUser) GetDouyinOk() (*string, bool)`

GetDouyinOk returns a tuple with the Douyin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDouyin

`func (o *IamUser) SetDouyin(v string)`

SetDouyin sets Douyin field to given value.

### HasDouyin

`func (o *IamUser) HasDouyin() bool`

HasDouyin returns a boolean if a field has been set.

### GetDropbox

`func (o *IamUser) GetDropbox() string`

GetDropbox returns the Dropbox field if non-nil, zero value otherwise.

### GetDropboxOk

`func (o *IamUser) GetDropboxOk() (*string, bool)`

GetDropboxOk returns a tuple with the Dropbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropbox

`func (o *IamUser) SetDropbox(v string)`

SetDropbox sets Dropbox field to given value.

### HasDropbox

`func (o *IamUser) HasDropbox() bool`

HasDropbox returns a boolean if a field has been set.

### GetEducation

`func (o *IamUser) GetEducation() string`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *IamUser) GetEducationOk() (*string, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *IamUser) SetEducation(v string)`

SetEducation sets Education field to given value.

### HasEducation

`func (o *IamUser) HasEducation() bool`

HasEducation returns a boolean if a field has been set.

### GetEmail

`func (o *IamUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IamUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IamUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IamUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailVerified

`func (o *IamUser) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *IamUser) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *IamUser) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *IamUser) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetEveonline

`func (o *IamUser) GetEveonline() string`

GetEveonline returns the Eveonline field if non-nil, zero value otherwise.

### GetEveonlineOk

`func (o *IamUser) GetEveonlineOk() (*string, bool)`

GetEveonlineOk returns a tuple with the Eveonline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEveonline

`func (o *IamUser) SetEveonline(v string)`

SetEveonline sets Eveonline field to given value.

### HasEveonline

`func (o *IamUser) HasEveonline() bool`

HasEveonline returns a boolean if a field has been set.

### GetExternalId

`func (o *IamUser) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *IamUser) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *IamUser) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *IamUser) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFaceIds

`func (o *IamUser) GetFaceIds() []IamFaceId`

GetFaceIds returns the FaceIds field if non-nil, zero value otherwise.

### GetFaceIdsOk

`func (o *IamUser) GetFaceIdsOk() (*[]IamFaceId, bool)`

GetFaceIdsOk returns a tuple with the FaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceIds

`func (o *IamUser) SetFaceIds(v []IamFaceId)`

SetFaceIds sets FaceIds field to given value.

### HasFaceIds

`func (o *IamUser) HasFaceIds() bool`

HasFaceIds returns a boolean if a field has been set.

### GetFacebook

`func (o *IamUser) GetFacebook() string`

GetFacebook returns the Facebook field if non-nil, zero value otherwise.

### GetFacebookOk

`func (o *IamUser) GetFacebookOk() (*string, bool)`

GetFacebookOk returns a tuple with the Facebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebook

`func (o *IamUser) SetFacebook(v string)`

SetFacebook sets Facebook field to given value.

### HasFacebook

`func (o *IamUser) HasFacebook() bool`

HasFacebook returns a boolean if a field has been set.

### GetFirstName

`func (o *IamUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *IamUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *IamUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *IamUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetFitbit

`func (o *IamUser) GetFitbit() string`

GetFitbit returns the Fitbit field if non-nil, zero value otherwise.

### GetFitbitOk

`func (o *IamUser) GetFitbitOk() (*string, bool)`

GetFitbitOk returns a tuple with the Fitbit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFitbit

`func (o *IamUser) SetFitbit(v string)`

SetFitbit sets Fitbit field to given value.

### HasFitbit

`func (o *IamUser) HasFitbit() bool`

HasFitbit returns a boolean if a field has been set.

### GetGender

`func (o *IamUser) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *IamUser) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *IamUser) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *IamUser) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetGitea

`func (o *IamUser) GetGitea() string`

GetGitea returns the Gitea field if non-nil, zero value otherwise.

### GetGiteaOk

`func (o *IamUser) GetGiteaOk() (*string, bool)`

GetGiteaOk returns a tuple with the Gitea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitea

`func (o *IamUser) SetGitea(v string)`

SetGitea sets Gitea field to given value.

### HasGitea

`func (o *IamUser) HasGitea() bool`

HasGitea returns a boolean if a field has been set.

### GetGitee

`func (o *IamUser) GetGitee() string`

GetGitee returns the Gitee field if non-nil, zero value otherwise.

### GetGiteeOk

`func (o *IamUser) GetGiteeOk() (*string, bool)`

GetGiteeOk returns a tuple with the Gitee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitee

`func (o *IamUser) SetGitee(v string)`

SetGitee sets Gitee field to given value.

### HasGitee

`func (o *IamUser) HasGitee() bool`

HasGitee returns a boolean if a field has been set.

### GetGithub

`func (o *IamUser) GetGithub() string`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *IamUser) GetGithubOk() (*string, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *IamUser) SetGithub(v string)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *IamUser) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetGitlab

`func (o *IamUser) GetGitlab() string`

GetGitlab returns the Gitlab field if non-nil, zero value otherwise.

### GetGitlabOk

`func (o *IamUser) GetGitlabOk() (*string, bool)`

GetGitlabOk returns a tuple with the Gitlab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlab

`func (o *IamUser) SetGitlab(v string)`

SetGitlab sets Gitlab field to given value.

### HasGitlab

`func (o *IamUser) HasGitlab() bool`

HasGitlab returns a boolean if a field has been set.

### GetGoogle

`func (o *IamUser) GetGoogle() string`

GetGoogle returns the Google field if non-nil, zero value otherwise.

### GetGoogleOk

`func (o *IamUser) GetGoogleOk() (*string, bool)`

GetGoogleOk returns a tuple with the Google field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogle

`func (o *IamUser) SetGoogle(v string)`

SetGoogle sets Google field to given value.

### HasGoogle

`func (o *IamUser) HasGoogle() bool`

HasGoogle returns a boolean if a field has been set.

### GetGroups

`func (o *IamUser) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IamUser) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IamUser) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *IamUser) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHash

`func (o *IamUser) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *IamUser) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *IamUser) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *IamUser) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetHeroku

`func (o *IamUser) GetHeroku() string`

GetHeroku returns the Heroku field if non-nil, zero value otherwise.

### GetHerokuOk

`func (o *IamUser) GetHerokuOk() (*string, bool)`

GetHerokuOk returns a tuple with the Heroku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroku

`func (o *IamUser) SetHeroku(v string)`

SetHeroku sets Heroku field to given value.

### HasHeroku

`func (o *IamUser) HasHeroku() bool`

HasHeroku returns a boolean if a field has been set.

### GetHomepage

`func (o *IamUser) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *IamUser) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *IamUser) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *IamUser) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetIam

`func (o *IamUser) GetIam() string`

GetIam returns the Iam field if non-nil, zero value otherwise.

### GetIamOk

`func (o *IamUser) GetIamOk() (*string, bool)`

GetIamOk returns a tuple with the Iam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIam

`func (o *IamUser) SetIam(v string)`

SetIam sets Iam field to given value.

### HasIam

`func (o *IamUser) HasIam() bool`

HasIam returns a boolean if a field has been set.

### GetId

`func (o *IamUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdCard

`func (o *IamUser) GetIdCard() string`

GetIdCard returns the IdCard field if non-nil, zero value otherwise.

### GetIdCardOk

`func (o *IamUser) GetIdCardOk() (*string, bool)`

GetIdCardOk returns a tuple with the IdCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCard

`func (o *IamUser) SetIdCard(v string)`

SetIdCard sets IdCard field to given value.

### HasIdCard

`func (o *IamUser) HasIdCard() bool`

HasIdCard returns a boolean if a field has been set.

### GetIdCardType

`func (o *IamUser) GetIdCardType() string`

GetIdCardType returns the IdCardType field if non-nil, zero value otherwise.

### GetIdCardTypeOk

`func (o *IamUser) GetIdCardTypeOk() (*string, bool)`

GetIdCardTypeOk returns a tuple with the IdCardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCardType

`func (o *IamUser) SetIdCardType(v string)`

SetIdCardType sets IdCardType field to given value.

### HasIdCardType

`func (o *IamUser) HasIdCardType() bool`

HasIdCardType returns a boolean if a field has been set.

### GetInfluxcloud

`func (o *IamUser) GetInfluxcloud() string`

GetInfluxcloud returns the Influxcloud field if non-nil, zero value otherwise.

### GetInfluxcloudOk

`func (o *IamUser) GetInfluxcloudOk() (*string, bool)`

GetInfluxcloudOk returns a tuple with the Influxcloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfluxcloud

`func (o *IamUser) SetInfluxcloud(v string)`

SetInfluxcloud sets Influxcloud field to given value.

### HasInfluxcloud

`func (o *IamUser) HasInfluxcloud() bool`

HasInfluxcloud returns a boolean if a field has been set.

### GetInfoflow

`func (o *IamUser) GetInfoflow() string`

GetInfoflow returns the Infoflow field if non-nil, zero value otherwise.

### GetInfoflowOk

`func (o *IamUser) GetInfoflowOk() (*string, bool)`

GetInfoflowOk returns a tuple with the Infoflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoflow

`func (o *IamUser) SetInfoflow(v string)`

SetInfoflow sets Infoflow field to given value.

### HasInfoflow

`func (o *IamUser) HasInfoflow() bool`

HasInfoflow returns a boolean if a field has been set.

### GetInstagram

`func (o *IamUser) GetInstagram() string`

GetInstagram returns the Instagram field if non-nil, zero value otherwise.

### GetInstagramOk

`func (o *IamUser) GetInstagramOk() (*string, bool)`

GetInstagramOk returns a tuple with the Instagram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagram

`func (o *IamUser) SetInstagram(v string)`

SetInstagram sets Instagram field to given value.

### HasInstagram

`func (o *IamUser) HasInstagram() bool`

HasInstagram returns a boolean if a field has been set.

### GetIntercom

`func (o *IamUser) GetIntercom() string`

GetIntercom returns the Intercom field if non-nil, zero value otherwise.

### GetIntercomOk

`func (o *IamUser) GetIntercomOk() (*string, bool)`

GetIntercomOk returns a tuple with the Intercom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntercom

`func (o *IamUser) SetIntercom(v string)`

SetIntercom sets Intercom field to given value.

### HasIntercom

`func (o *IamUser) HasIntercom() bool`

HasIntercom returns a boolean if a field has been set.

### GetInvitation

`func (o *IamUser) GetInvitation() string`

GetInvitation returns the Invitation field if non-nil, zero value otherwise.

### GetInvitationOk

`func (o *IamUser) GetInvitationOk() (*string, bool)`

GetInvitationOk returns a tuple with the Invitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitation

`func (o *IamUser) SetInvitation(v string)`

SetInvitation sets Invitation field to given value.

### HasInvitation

`func (o *IamUser) HasInvitation() bool`

HasInvitation returns a boolean if a field has been set.

### GetInvitationCode

`func (o *IamUser) GetInvitationCode() string`

GetInvitationCode returns the InvitationCode field if non-nil, zero value otherwise.

### GetInvitationCodeOk

`func (o *IamUser) GetInvitationCodeOk() (*string, bool)`

GetInvitationCodeOk returns a tuple with the InvitationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationCode

`func (o *IamUser) SetInvitationCode(v string)`

SetInvitationCode sets InvitationCode field to given value.

### HasInvitationCode

`func (o *IamUser) HasInvitationCode() bool`

HasInvitationCode returns a boolean if a field has been set.

### GetIpWhitelist

`func (o *IamUser) GetIpWhitelist() string`

GetIpWhitelist returns the IpWhitelist field if non-nil, zero value otherwise.

### GetIpWhitelistOk

`func (o *IamUser) GetIpWhitelistOk() (*string, bool)`

GetIpWhitelistOk returns a tuple with the IpWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelist

`func (o *IamUser) SetIpWhitelist(v string)`

SetIpWhitelist sets IpWhitelist field to given value.

### HasIpWhitelist

`func (o *IamUser) HasIpWhitelist() bool`

HasIpWhitelist returns a boolean if a field has been set.

### GetIsAdmin

`func (o *IamUser) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *IamUser) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *IamUser) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *IamUser) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetIsDefaultAvatar

`func (o *IamUser) GetIsDefaultAvatar() bool`

GetIsDefaultAvatar returns the IsDefaultAvatar field if non-nil, zero value otherwise.

### GetIsDefaultAvatarOk

`func (o *IamUser) GetIsDefaultAvatarOk() (*bool, bool)`

GetIsDefaultAvatarOk returns a tuple with the IsDefaultAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAvatar

`func (o *IamUser) SetIsDefaultAvatar(v bool)`

SetIsDefaultAvatar sets IsDefaultAvatar field to given value.

### HasIsDefaultAvatar

`func (o *IamUser) HasIsDefaultAvatar() bool`

HasIsDefaultAvatar returns a boolean if a field has been set.

### GetIsDeleted

`func (o *IamUser) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *IamUser) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *IamUser) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *IamUser) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsForbidden

`func (o *IamUser) GetIsForbidden() bool`

GetIsForbidden returns the IsForbidden field if non-nil, zero value otherwise.

### GetIsForbiddenOk

`func (o *IamUser) GetIsForbiddenOk() (*bool, bool)`

GetIsForbiddenOk returns a tuple with the IsForbidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForbidden

`func (o *IamUser) SetIsForbidden(v bool)`

SetIsForbidden sets IsForbidden field to given value.

### HasIsForbidden

`func (o *IamUser) HasIsForbidden() bool`

HasIsForbidden returns a boolean if a field has been set.

### GetIsOnline

`func (o *IamUser) GetIsOnline() bool`

GetIsOnline returns the IsOnline field if non-nil, zero value otherwise.

### GetIsOnlineOk

`func (o *IamUser) GetIsOnlineOk() (*bool, bool)`

GetIsOnlineOk returns a tuple with the IsOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnline

`func (o *IamUser) SetIsOnline(v bool)`

SetIsOnline sets IsOnline field to given value.

### HasIsOnline

`func (o *IamUser) HasIsOnline() bool`

HasIsOnline returns a boolean if a field has been set.

### GetIsVerified

`func (o *IamUser) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *IamUser) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *IamUser) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *IamUser) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetKakao

`func (o *IamUser) GetKakao() string`

GetKakao returns the Kakao field if non-nil, zero value otherwise.

### GetKakaoOk

`func (o *IamUser) GetKakaoOk() (*string, bool)`

GetKakaoOk returns a tuple with the Kakao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKakao

`func (o *IamUser) SetKakao(v string)`

SetKakao sets Kakao field to given value.

### HasKakao

`func (o *IamUser) HasKakao() bool`

HasKakao returns a boolean if a field has been set.

### GetKarma

`func (o *IamUser) GetKarma() int32`

GetKarma returns the Karma field if non-nil, zero value otherwise.

### GetKarmaOk

`func (o *IamUser) GetKarmaOk() (*int32, bool)`

GetKarmaOk returns a tuple with the Karma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKarma

`func (o *IamUser) SetKarma(v int32)`

SetKarma sets Karma field to given value.

### HasKarma

`func (o *IamUser) HasKarma() bool`

HasKarma returns a boolean if a field has been set.

### GetKwai

`func (o *IamUser) GetKwai() string`

GetKwai returns the Kwai field if non-nil, zero value otherwise.

### GetKwaiOk

`func (o *IamUser) GetKwaiOk() (*string, bool)`

GetKwaiOk returns a tuple with the Kwai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKwai

`func (o *IamUser) SetKwai(v string)`

SetKwai sets Kwai field to given value.

### HasKwai

`func (o *IamUser) HasKwai() bool`

HasKwai returns a boolean if a field has been set.

### GetLanguage

`func (o *IamUser) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *IamUser) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *IamUser) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *IamUser) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLark

`func (o *IamUser) GetLark() string`

GetLark returns the Lark field if non-nil, zero value otherwise.

### GetLarkOk

`func (o *IamUser) GetLarkOk() (*string, bool)`

GetLarkOk returns a tuple with the Lark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLark

`func (o *IamUser) SetLark(v string)`

SetLark sets Lark field to given value.

### HasLark

`func (o *IamUser) HasLark() bool`

HasLark returns a boolean if a field has been set.

### GetLastChangePasswordTime

`func (o *IamUser) GetLastChangePasswordTime() string`

GetLastChangePasswordTime returns the LastChangePasswordTime field if non-nil, zero value otherwise.

### GetLastChangePasswordTimeOk

`func (o *IamUser) GetLastChangePasswordTimeOk() (*string, bool)`

GetLastChangePasswordTimeOk returns a tuple with the LastChangePasswordTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangePasswordTime

`func (o *IamUser) SetLastChangePasswordTime(v string)`

SetLastChangePasswordTime sets LastChangePasswordTime field to given value.

### HasLastChangePasswordTime

`func (o *IamUser) HasLastChangePasswordTime() bool`

HasLastChangePasswordTime returns a boolean if a field has been set.

### GetLastName

`func (o *IamUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *IamUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *IamUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *IamUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLastSigninIp

`func (o *IamUser) GetLastSigninIp() string`

GetLastSigninIp returns the LastSigninIp field if non-nil, zero value otherwise.

### GetLastSigninIpOk

`func (o *IamUser) GetLastSigninIpOk() (*string, bool)`

GetLastSigninIpOk returns a tuple with the LastSigninIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSigninIp

`func (o *IamUser) SetLastSigninIp(v string)`

SetLastSigninIp sets LastSigninIp field to given value.

### HasLastSigninIp

`func (o *IamUser) HasLastSigninIp() bool`

HasLastSigninIp returns a boolean if a field has been set.

### GetLastSigninTime

`func (o *IamUser) GetLastSigninTime() string`

GetLastSigninTime returns the LastSigninTime field if non-nil, zero value otherwise.

### GetLastSigninTimeOk

`func (o *IamUser) GetLastSigninTimeOk() (*string, bool)`

GetLastSigninTimeOk returns a tuple with the LastSigninTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSigninTime

`func (o *IamUser) SetLastSigninTime(v string)`

SetLastSigninTime sets LastSigninTime field to given value.

### HasLastSigninTime

`func (o *IamUser) HasLastSigninTime() bool`

HasLastSigninTime returns a boolean if a field has been set.

### GetLastSigninWrongTime

`func (o *IamUser) GetLastSigninWrongTime() string`

GetLastSigninWrongTime returns the LastSigninWrongTime field if non-nil, zero value otherwise.

### GetLastSigninWrongTimeOk

`func (o *IamUser) GetLastSigninWrongTimeOk() (*string, bool)`

GetLastSigninWrongTimeOk returns a tuple with the LastSigninWrongTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSigninWrongTime

`func (o *IamUser) SetLastSigninWrongTime(v string)`

SetLastSigninWrongTime sets LastSigninWrongTime field to given value.

### HasLastSigninWrongTime

`func (o *IamUser) HasLastSigninWrongTime() bool`

HasLastSigninWrongTime returns a boolean if a field has been set.

### GetLastfm

`func (o *IamUser) GetLastfm() string`

GetLastfm returns the Lastfm field if non-nil, zero value otherwise.

### GetLastfmOk

`func (o *IamUser) GetLastfmOk() (*string, bool)`

GetLastfmOk returns a tuple with the Lastfm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastfm

`func (o *IamUser) SetLastfm(v string)`

SetLastfm sets Lastfm field to given value.

### HasLastfm

`func (o *IamUser) HasLastfm() bool`

HasLastfm returns a boolean if a field has been set.

### GetLdap

`func (o *IamUser) GetLdap() string`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *IamUser) GetLdapOk() (*string, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *IamUser) SetLdap(v string)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *IamUser) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetLine

`func (o *IamUser) GetLine() string`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *IamUser) GetLineOk() (*string, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *IamUser) SetLine(v string)`

SetLine sets Line field to given value.

### HasLine

`func (o *IamUser) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetLinkedin

`func (o *IamUser) GetLinkedin() string`

GetLinkedin returns the Linkedin field if non-nil, zero value otherwise.

### GetLinkedinOk

`func (o *IamUser) GetLinkedinOk() (*string, bool)`

GetLinkedinOk returns a tuple with the Linkedin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedin

`func (o *IamUser) SetLinkedin(v string)`

SetLinkedin sets Linkedin field to given value.

### HasLinkedin

`func (o *IamUser) HasLinkedin() bool`

HasLinkedin returns a boolean if a field has been set.

### GetLocation

`func (o *IamUser) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *IamUser) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *IamUser) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *IamUser) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMailru

`func (o *IamUser) GetMailru() string`

GetMailru returns the Mailru field if non-nil, zero value otherwise.

### GetMailruOk

`func (o *IamUser) GetMailruOk() (*string, bool)`

GetMailruOk returns a tuple with the Mailru field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailru

`func (o *IamUser) SetMailru(v string)`

SetMailru sets Mailru field to given value.

### HasMailru

`func (o *IamUser) HasMailru() bool`

HasMailru returns a boolean if a field has been set.

### GetManagedAccounts

`func (o *IamUser) GetManagedAccounts() []IamManagedAccount`

GetManagedAccounts returns the ManagedAccounts field if non-nil, zero value otherwise.

### GetManagedAccountsOk

`func (o *IamUser) GetManagedAccountsOk() (*[]IamManagedAccount, bool)`

GetManagedAccountsOk returns a tuple with the ManagedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAccounts

`func (o *IamUser) SetManagedAccounts(v []IamManagedAccount)`

SetManagedAccounts sets ManagedAccounts field to given value.

### HasManagedAccounts

`func (o *IamUser) HasManagedAccounts() bool`

HasManagedAccounts returns a boolean if a field has been set.

### GetMeetup

`func (o *IamUser) GetMeetup() string`

GetMeetup returns the Meetup field if non-nil, zero value otherwise.

### GetMeetupOk

`func (o *IamUser) GetMeetupOk() (*string, bool)`

GetMeetupOk returns a tuple with the Meetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetup

`func (o *IamUser) SetMeetup(v string)`

SetMeetup sets Meetup field to given value.

### HasMeetup

`func (o *IamUser) HasMeetup() bool`

HasMeetup returns a boolean if a field has been set.

### GetMfaAccounts

`func (o *IamUser) GetMfaAccounts() []IamMfaAccount`

GetMfaAccounts returns the MfaAccounts field if non-nil, zero value otherwise.

### GetMfaAccountsOk

`func (o *IamUser) GetMfaAccountsOk() (*[]IamMfaAccount, bool)`

GetMfaAccountsOk returns a tuple with the MfaAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaAccounts

`func (o *IamUser) SetMfaAccounts(v []IamMfaAccount)`

SetMfaAccounts sets MfaAccounts field to given value.

### HasMfaAccounts

`func (o *IamUser) HasMfaAccounts() bool`

HasMfaAccounts returns a boolean if a field has been set.

### GetMfaEmailEnabled

`func (o *IamUser) GetMfaEmailEnabled() bool`

GetMfaEmailEnabled returns the MfaEmailEnabled field if non-nil, zero value otherwise.

### GetMfaEmailEnabledOk

`func (o *IamUser) GetMfaEmailEnabledOk() (*bool, bool)`

GetMfaEmailEnabledOk returns a tuple with the MfaEmailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaEmailEnabled

`func (o *IamUser) SetMfaEmailEnabled(v bool)`

SetMfaEmailEnabled sets MfaEmailEnabled field to given value.

### HasMfaEmailEnabled

`func (o *IamUser) HasMfaEmailEnabled() bool`

HasMfaEmailEnabled returns a boolean if a field has been set.

### GetMfaItems

`func (o *IamUser) GetMfaItems() []IamMfaItem`

GetMfaItems returns the MfaItems field if non-nil, zero value otherwise.

### GetMfaItemsOk

`func (o *IamUser) GetMfaItemsOk() (*[]IamMfaItem, bool)`

GetMfaItemsOk returns a tuple with the MfaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaItems

`func (o *IamUser) SetMfaItems(v []IamMfaItem)`

SetMfaItems sets MfaItems field to given value.

### HasMfaItems

`func (o *IamUser) HasMfaItems() bool`

HasMfaItems returns a boolean if a field has been set.

### GetMfaPhoneEnabled

`func (o *IamUser) GetMfaPhoneEnabled() bool`

GetMfaPhoneEnabled returns the MfaPhoneEnabled field if non-nil, zero value otherwise.

### GetMfaPhoneEnabledOk

`func (o *IamUser) GetMfaPhoneEnabledOk() (*bool, bool)`

GetMfaPhoneEnabledOk returns a tuple with the MfaPhoneEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaPhoneEnabled

`func (o *IamUser) SetMfaPhoneEnabled(v bool)`

SetMfaPhoneEnabled sets MfaPhoneEnabled field to given value.

### HasMfaPhoneEnabled

`func (o *IamUser) HasMfaPhoneEnabled() bool`

HasMfaPhoneEnabled returns a boolean if a field has been set.

### GetMfaPushEnabled

`func (o *IamUser) GetMfaPushEnabled() bool`

GetMfaPushEnabled returns the MfaPushEnabled field if non-nil, zero value otherwise.

### GetMfaPushEnabledOk

`func (o *IamUser) GetMfaPushEnabledOk() (*bool, bool)`

GetMfaPushEnabledOk returns a tuple with the MfaPushEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaPushEnabled

`func (o *IamUser) SetMfaPushEnabled(v bool)`

SetMfaPushEnabled sets MfaPushEnabled field to given value.

### HasMfaPushEnabled

`func (o *IamUser) HasMfaPushEnabled() bool`

HasMfaPushEnabled returns a boolean if a field has been set.

### GetMfaPushProvider

`func (o *IamUser) GetMfaPushProvider() string`

GetMfaPushProvider returns the MfaPushProvider field if non-nil, zero value otherwise.

### GetMfaPushProviderOk

`func (o *IamUser) GetMfaPushProviderOk() (*string, bool)`

GetMfaPushProviderOk returns a tuple with the MfaPushProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaPushProvider

`func (o *IamUser) SetMfaPushProvider(v string)`

SetMfaPushProvider sets MfaPushProvider field to given value.

### HasMfaPushProvider

`func (o *IamUser) HasMfaPushProvider() bool`

HasMfaPushProvider returns a boolean if a field has been set.

### GetMfaPushReceiver

`func (o *IamUser) GetMfaPushReceiver() string`

GetMfaPushReceiver returns the MfaPushReceiver field if non-nil, zero value otherwise.

### GetMfaPushReceiverOk

`func (o *IamUser) GetMfaPushReceiverOk() (*string, bool)`

GetMfaPushReceiverOk returns a tuple with the MfaPushReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaPushReceiver

`func (o *IamUser) SetMfaPushReceiver(v string)`

SetMfaPushReceiver sets MfaPushReceiver field to given value.

### HasMfaPushReceiver

`func (o *IamUser) HasMfaPushReceiver() bool`

HasMfaPushReceiver returns a boolean if a field has been set.

### GetMfaRadiusEnabled

`func (o *IamUser) GetMfaRadiusEnabled() bool`

GetMfaRadiusEnabled returns the MfaRadiusEnabled field if non-nil, zero value otherwise.

### GetMfaRadiusEnabledOk

`func (o *IamUser) GetMfaRadiusEnabledOk() (*bool, bool)`

GetMfaRadiusEnabledOk returns a tuple with the MfaRadiusEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRadiusEnabled

`func (o *IamUser) SetMfaRadiusEnabled(v bool)`

SetMfaRadiusEnabled sets MfaRadiusEnabled field to given value.

### HasMfaRadiusEnabled

`func (o *IamUser) HasMfaRadiusEnabled() bool`

HasMfaRadiusEnabled returns a boolean if a field has been set.

### GetMfaRadiusProvider

`func (o *IamUser) GetMfaRadiusProvider() string`

GetMfaRadiusProvider returns the MfaRadiusProvider field if non-nil, zero value otherwise.

### GetMfaRadiusProviderOk

`func (o *IamUser) GetMfaRadiusProviderOk() (*string, bool)`

GetMfaRadiusProviderOk returns a tuple with the MfaRadiusProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRadiusProvider

`func (o *IamUser) SetMfaRadiusProvider(v string)`

SetMfaRadiusProvider sets MfaRadiusProvider field to given value.

### HasMfaRadiusProvider

`func (o *IamUser) HasMfaRadiusProvider() bool`

HasMfaRadiusProvider returns a boolean if a field has been set.

### GetMfaRadiusUsername

`func (o *IamUser) GetMfaRadiusUsername() string`

GetMfaRadiusUsername returns the MfaRadiusUsername field if non-nil, zero value otherwise.

### GetMfaRadiusUsernameOk

`func (o *IamUser) GetMfaRadiusUsernameOk() (*string, bool)`

GetMfaRadiusUsernameOk returns a tuple with the MfaRadiusUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRadiusUsername

`func (o *IamUser) SetMfaRadiusUsername(v string)`

SetMfaRadiusUsername sets MfaRadiusUsername field to given value.

### HasMfaRadiusUsername

`func (o *IamUser) HasMfaRadiusUsername() bool`

HasMfaRadiusUsername returns a boolean if a field has been set.

### GetMfaRememberDeadline

`func (o *IamUser) GetMfaRememberDeadline() string`

GetMfaRememberDeadline returns the MfaRememberDeadline field if non-nil, zero value otherwise.

### GetMfaRememberDeadlineOk

`func (o *IamUser) GetMfaRememberDeadlineOk() (*string, bool)`

GetMfaRememberDeadlineOk returns a tuple with the MfaRememberDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRememberDeadline

`func (o *IamUser) SetMfaRememberDeadline(v string)`

SetMfaRememberDeadline sets MfaRememberDeadline field to given value.

### HasMfaRememberDeadline

`func (o *IamUser) HasMfaRememberDeadline() bool`

HasMfaRememberDeadline returns a boolean if a field has been set.

### GetMfaRememberDigest

`func (o *IamUser) GetMfaRememberDigest() string`

GetMfaRememberDigest returns the MfaRememberDigest field if non-nil, zero value otherwise.

### GetMfaRememberDigestOk

`func (o *IamUser) GetMfaRememberDigestOk() (*string, bool)`

GetMfaRememberDigestOk returns a tuple with the MfaRememberDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRememberDigest

`func (o *IamUser) SetMfaRememberDigest(v string)`

SetMfaRememberDigest sets MfaRememberDigest field to given value.

### HasMfaRememberDigest

`func (o *IamUser) HasMfaRememberDigest() bool`

HasMfaRememberDigest returns a boolean if a field has been set.

### GetMicrosoftonline

`func (o *IamUser) GetMicrosoftonline() string`

GetMicrosoftonline returns the Microsoftonline field if non-nil, zero value otherwise.

### GetMicrosoftonlineOk

`func (o *IamUser) GetMicrosoftonlineOk() (*string, bool)`

GetMicrosoftonlineOk returns a tuple with the Microsoftonline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftonline

`func (o *IamUser) SetMicrosoftonline(v string)`

SetMicrosoftonline sets Microsoftonline field to given value.

### HasMicrosoftonline

`func (o *IamUser) HasMicrosoftonline() bool`

HasMicrosoftonline returns a boolean if a field has been set.

### GetMultiFactorAuths

`func (o *IamUser) GetMultiFactorAuths() []IamMfaProps`

GetMultiFactorAuths returns the MultiFactorAuths field if non-nil, zero value otherwise.

### GetMultiFactorAuthsOk

`func (o *IamUser) GetMultiFactorAuthsOk() (*[]IamMfaProps, bool)`

GetMultiFactorAuthsOk returns a tuple with the MultiFactorAuths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiFactorAuths

`func (o *IamUser) SetMultiFactorAuths(v []IamMfaProps)`

SetMultiFactorAuths sets MultiFactorAuths field to given value.

### HasMultiFactorAuths

`func (o *IamUser) HasMultiFactorAuths() bool`

HasMultiFactorAuths returns a boolean if a field has been set.

### GetName

`func (o *IamUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNaver

`func (o *IamUser) GetNaver() string`

GetNaver returns the Naver field if non-nil, zero value otherwise.

### GetNaverOk

`func (o *IamUser) GetNaverOk() (*string, bool)`

GetNaverOk returns a tuple with the Naver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaver

`func (o *IamUser) SetNaver(v string)`

SetNaver sets Naver field to given value.

### HasNaver

`func (o *IamUser) HasNaver() bool`

HasNaver returns a boolean if a field has been set.

### GetNeedUpdatePassword

`func (o *IamUser) GetNeedUpdatePassword() bool`

GetNeedUpdatePassword returns the NeedUpdatePassword field if non-nil, zero value otherwise.

### GetNeedUpdatePasswordOk

`func (o *IamUser) GetNeedUpdatePasswordOk() (*bool, bool)`

GetNeedUpdatePasswordOk returns a tuple with the NeedUpdatePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedUpdatePassword

`func (o *IamUser) SetNeedUpdatePassword(v bool)`

SetNeedUpdatePassword sets NeedUpdatePassword field to given value.

### HasNeedUpdatePassword

`func (o *IamUser) HasNeedUpdatePassword() bool`

HasNeedUpdatePassword returns a boolean if a field has been set.

### GetNextcloud

`func (o *IamUser) GetNextcloud() string`

GetNextcloud returns the Nextcloud field if non-nil, zero value otherwise.

### GetNextcloudOk

`func (o *IamUser) GetNextcloudOk() (*string, bool)`

GetNextcloudOk returns a tuple with the Nextcloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextcloud

`func (o *IamUser) SetNextcloud(v string)`

SetNextcloud sets Nextcloud field to given value.

### HasNextcloud

`func (o *IamUser) HasNextcloud() bool`

HasNextcloud returns a boolean if a field has been set.

### GetOkta

`func (o *IamUser) GetOkta() string`

GetOkta returns the Okta field if non-nil, zero value otherwise.

### GetOktaOk

`func (o *IamUser) GetOktaOk() (*string, bool)`

GetOktaOk returns a tuple with the Okta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOkta

`func (o *IamUser) SetOkta(v string)`

SetOkta sets Okta field to given value.

### HasOkta

`func (o *IamUser) HasOkta() bool`

HasOkta returns a boolean if a field has been set.

### GetOnedrive

`func (o *IamUser) GetOnedrive() string`

GetOnedrive returns the Onedrive field if non-nil, zero value otherwise.

### GetOnedriveOk

`func (o *IamUser) GetOnedriveOk() (*string, bool)`

GetOnedriveOk returns a tuple with the Onedrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnedrive

`func (o *IamUser) SetOnedrive(v string)`

SetOnedrive sets Onedrive field to given value.

### HasOnedrive

`func (o *IamUser) HasOnedrive() bool`

HasOnedrive returns a boolean if a field has been set.

### GetOriginalRefreshToken

`func (o *IamUser) GetOriginalRefreshToken() string`

GetOriginalRefreshToken returns the OriginalRefreshToken field if non-nil, zero value otherwise.

### GetOriginalRefreshTokenOk

`func (o *IamUser) GetOriginalRefreshTokenOk() (*string, bool)`

GetOriginalRefreshTokenOk returns a tuple with the OriginalRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalRefreshToken

`func (o *IamUser) SetOriginalRefreshToken(v string)`

SetOriginalRefreshToken sets OriginalRefreshToken field to given value.

### HasOriginalRefreshToken

`func (o *IamUser) HasOriginalRefreshToken() bool`

HasOriginalRefreshToken returns a boolean if a field has been set.

### GetOriginalToken

`func (o *IamUser) GetOriginalToken() string`

GetOriginalToken returns the OriginalToken field if non-nil, zero value otherwise.

### GetOriginalTokenOk

`func (o *IamUser) GetOriginalTokenOk() (*string, bool)`

GetOriginalTokenOk returns a tuple with the OriginalToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalToken

`func (o *IamUser) SetOriginalToken(v string)`

SetOriginalToken sets OriginalToken field to given value.

### HasOriginalToken

`func (o *IamUser) HasOriginalToken() bool`

HasOriginalToken returns a boolean if a field has been set.

### GetOura

`func (o *IamUser) GetOura() string`

GetOura returns the Oura field if non-nil, zero value otherwise.

### GetOuraOk

`func (o *IamUser) GetOuraOk() (*string, bool)`

GetOuraOk returns a tuple with the Oura field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOura

`func (o *IamUser) SetOura(v string)`

SetOura sets Oura field to given value.

### HasOura

`func (o *IamUser) HasOura() bool`

HasOura returns a boolean if a field has been set.

### GetOwner

`func (o *IamUser) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamUser) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamUser) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamUser) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPasswordHash

`func (o *IamUser) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *IamUser) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *IamUser) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *IamUser) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetPasswordSalt

`func (o *IamUser) GetPasswordSalt() string`

GetPasswordSalt returns the PasswordSalt field if non-nil, zero value otherwise.

### GetPasswordSaltOk

`func (o *IamUser) GetPasswordSaltOk() (*string, bool)`

GetPasswordSaltOk returns a tuple with the PasswordSalt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSalt

`func (o *IamUser) SetPasswordSalt(v string)`

SetPasswordSalt sets PasswordSalt field to given value.

### HasPasswordSalt

`func (o *IamUser) HasPasswordSalt() bool`

HasPasswordSalt returns a boolean if a field has been set.

### GetPasswordType

`func (o *IamUser) GetPasswordType() string`

GetPasswordType returns the PasswordType field if non-nil, zero value otherwise.

### GetPasswordTypeOk

`func (o *IamUser) GetPasswordTypeOk() (*string, bool)`

GetPasswordTypeOk returns a tuple with the PasswordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordType

`func (o *IamUser) SetPasswordType(v string)`

SetPasswordType sets PasswordType field to given value.

### HasPasswordType

`func (o *IamUser) HasPasswordType() bool`

HasPasswordType returns a boolean if a field has been set.

### GetPatreon

`func (o *IamUser) GetPatreon() string`

GetPatreon returns the Patreon field if non-nil, zero value otherwise.

### GetPatreonOk

`func (o *IamUser) GetPatreonOk() (*string, bool)`

GetPatreonOk returns a tuple with the Patreon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatreon

`func (o *IamUser) SetPatreon(v string)`

SetPatreon sets Patreon field to given value.

### HasPatreon

`func (o *IamUser) HasPatreon() bool`

HasPatreon returns a boolean if a field has been set.

### GetPaypal

`func (o *IamUser) GetPaypal() string`

GetPaypal returns the Paypal field if non-nil, zero value otherwise.

### GetPaypalOk

`func (o *IamUser) GetPaypalOk() (*string, bool)`

GetPaypalOk returns a tuple with the Paypal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypal

`func (o *IamUser) SetPaypal(v string)`

SetPaypal sets Paypal field to given value.

### HasPaypal

`func (o *IamUser) HasPaypal() bool`

HasPaypal returns a boolean if a field has been set.

### GetPermanentAvatar

`func (o *IamUser) GetPermanentAvatar() string`

GetPermanentAvatar returns the PermanentAvatar field if non-nil, zero value otherwise.

### GetPermanentAvatarOk

`func (o *IamUser) GetPermanentAvatarOk() (*string, bool)`

GetPermanentAvatarOk returns a tuple with the PermanentAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentAvatar

`func (o *IamUser) SetPermanentAvatar(v string)`

SetPermanentAvatar sets PermanentAvatar field to given value.

### HasPermanentAvatar

`func (o *IamUser) HasPermanentAvatar() bool`

HasPermanentAvatar returns a boolean if a field has been set.

### GetPermissions

`func (o *IamUser) GetPermissions() []IamPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *IamUser) GetPermissionsOk() (*[]IamPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *IamUser) SetPermissions(v []IamPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *IamUser) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPhone

`func (o *IamUser) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *IamUser) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *IamUser) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *IamUser) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPreHash

`func (o *IamUser) GetPreHash() string`

GetPreHash returns the PreHash field if non-nil, zero value otherwise.

### GetPreHashOk

`func (o *IamUser) GetPreHashOk() (*string, bool)`

GetPreHashOk returns a tuple with the PreHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreHash

`func (o *IamUser) SetPreHash(v string)`

SetPreHash sets PreHash field to given value.

### HasPreHash

`func (o *IamUser) HasPreHash() bool`

HasPreHash returns a boolean if a field has been set.

### GetPreferredMfaType

`func (o *IamUser) GetPreferredMfaType() string`

GetPreferredMfaType returns the PreferredMfaType field if non-nil, zero value otherwise.

### GetPreferredMfaTypeOk

`func (o *IamUser) GetPreferredMfaTypeOk() (*string, bool)`

GetPreferredMfaTypeOk returns a tuple with the PreferredMfaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMfaType

`func (o *IamUser) SetPreferredMfaType(v string)`

SetPreferredMfaType sets PreferredMfaType field to given value.

### HasPreferredMfaType

`func (o *IamUser) HasPreferredMfaType() bool`

HasPreferredMfaType returns a boolean if a field has been set.

### GetProperties

`func (o *IamUser) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IamUser) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IamUser) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *IamUser) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetQq

`func (o *IamUser) GetQq() string`

GetQq returns the Qq field if non-nil, zero value otherwise.

### GetQqOk

`func (o *IamUser) GetQqOk() (*string, bool)`

GetQqOk returns a tuple with the Qq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQq

`func (o *IamUser) SetQq(v string)`

SetQq sets Qq field to given value.

### HasQq

`func (o *IamUser) HasQq() bool`

HasQq returns a boolean if a field has been set.

### GetRanking

`func (o *IamUser) GetRanking() int32`

GetRanking returns the Ranking field if non-nil, zero value otherwise.

### GetRankingOk

`func (o *IamUser) GetRankingOk() (*int32, bool)`

GetRankingOk returns a tuple with the Ranking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanking

`func (o *IamUser) SetRanking(v int32)`

SetRanking sets Ranking field to given value.

### HasRanking

`func (o *IamUser) HasRanking() bool`

HasRanking returns a boolean if a field has been set.

### GetRealName

`func (o *IamUser) GetRealName() string`

GetRealName returns the RealName field if non-nil, zero value otherwise.

### GetRealNameOk

`func (o *IamUser) GetRealNameOk() (*string, bool)`

GetRealNameOk returns a tuple with the RealName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealName

`func (o *IamUser) SetRealName(v string)`

SetRealName sets RealName field to given value.

### HasRealName

`func (o *IamUser) HasRealName() bool`

HasRealName returns a boolean if a field has been set.

### GetRecoveryCodes

`func (o *IamUser) GetRecoveryCodes() []string`

GetRecoveryCodes returns the RecoveryCodes field if non-nil, zero value otherwise.

### GetRecoveryCodesOk

`func (o *IamUser) GetRecoveryCodesOk() (*[]string, bool)`

GetRecoveryCodesOk returns a tuple with the RecoveryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryCodes

`func (o *IamUser) SetRecoveryCodes(v []string)`

SetRecoveryCodes sets RecoveryCodes field to given value.

### HasRecoveryCodes

`func (o *IamUser) HasRecoveryCodes() bool`

HasRecoveryCodes returns a boolean if a field has been set.

### GetRegion

`func (o *IamUser) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *IamUser) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *IamUser) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *IamUser) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRegisterSource

`func (o *IamUser) GetRegisterSource() string`

GetRegisterSource returns the RegisterSource field if non-nil, zero value otherwise.

### GetRegisterSourceOk

`func (o *IamUser) GetRegisterSourceOk() (*string, bool)`

GetRegisterSourceOk returns a tuple with the RegisterSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterSource

`func (o *IamUser) SetRegisterSource(v string)`

SetRegisterSource sets RegisterSource field to given value.

### HasRegisterSource

`func (o *IamUser) HasRegisterSource() bool`

HasRegisterSource returns a boolean if a field has been set.

### GetRegisterType

`func (o *IamUser) GetRegisterType() string`

GetRegisterType returns the RegisterType field if non-nil, zero value otherwise.

### GetRegisterTypeOk

`func (o *IamUser) GetRegisterTypeOk() (*string, bool)`

GetRegisterTypeOk returns a tuple with the RegisterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterType

`func (o *IamUser) SetRegisterType(v string)`

SetRegisterType sets RegisterType field to given value.

### HasRegisterType

`func (o *IamUser) HasRegisterType() bool`

HasRegisterType returns a boolean if a field has been set.

### GetRoles

`func (o *IamUser) GetRoles() []IamRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamUser) GetRolesOk() (*[]IamRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamUser) SetRoles(v []IamRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSalesforce

`func (o *IamUser) GetSalesforce() string`

GetSalesforce returns the Salesforce field if non-nil, zero value otherwise.

### GetSalesforceOk

`func (o *IamUser) GetSalesforceOk() (*string, bool)`

GetSalesforceOk returns a tuple with the Salesforce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesforce

`func (o *IamUser) SetSalesforce(v string)`

SetSalesforce sets Salesforce field to given value.

### HasSalesforce

`func (o *IamUser) HasSalesforce() bool`

HasSalesforce returns a boolean if a field has been set.

### GetScore

`func (o *IamUser) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *IamUser) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *IamUser) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *IamUser) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetShopify

`func (o *IamUser) GetShopify() string`

GetShopify returns the Shopify field if non-nil, zero value otherwise.

### GetShopifyOk

`func (o *IamUser) GetShopifyOk() (*string, bool)`

GetShopifyOk returns a tuple with the Shopify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopify

`func (o *IamUser) SetShopify(v string)`

SetShopify sets Shopify field to given value.

### HasShopify

`func (o *IamUser) HasShopify() bool`

HasShopify returns a boolean if a field has been set.

### GetSigninWrongTimes

`func (o *IamUser) GetSigninWrongTimes() int32`

GetSigninWrongTimes returns the SigninWrongTimes field if non-nil, zero value otherwise.

### GetSigninWrongTimesOk

`func (o *IamUser) GetSigninWrongTimesOk() (*int32, bool)`

GetSigninWrongTimesOk returns a tuple with the SigninWrongTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigninWrongTimes

`func (o *IamUser) SetSigninWrongTimes(v int32)`

SetSigninWrongTimes sets SigninWrongTimes field to given value.

### HasSigninWrongTimes

`func (o *IamUser) HasSigninWrongTimes() bool`

HasSigninWrongTimes returns a boolean if a field has been set.

### GetSignupApplication

`func (o *IamUser) GetSignupApplication() string`

GetSignupApplication returns the SignupApplication field if non-nil, zero value otherwise.

### GetSignupApplicationOk

`func (o *IamUser) GetSignupApplicationOk() (*string, bool)`

GetSignupApplicationOk returns a tuple with the SignupApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupApplication

`func (o *IamUser) SetSignupApplication(v string)`

SetSignupApplication sets SignupApplication field to given value.

### HasSignupApplication

`func (o *IamUser) HasSignupApplication() bool`

HasSignupApplication returns a boolean if a field has been set.

### GetSlack

`func (o *IamUser) GetSlack() string`

GetSlack returns the Slack field if non-nil, zero value otherwise.

### GetSlackOk

`func (o *IamUser) GetSlackOk() (*string, bool)`

GetSlackOk returns a tuple with the Slack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlack

`func (o *IamUser) SetSlack(v string)`

SetSlack sets Slack field to given value.

### HasSlack

`func (o *IamUser) HasSlack() bool`

HasSlack returns a boolean if a field has been set.

### GetSoundcloud

`func (o *IamUser) GetSoundcloud() string`

GetSoundcloud returns the Soundcloud field if non-nil, zero value otherwise.

### GetSoundcloudOk

`func (o *IamUser) GetSoundcloudOk() (*string, bool)`

GetSoundcloudOk returns a tuple with the Soundcloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoundcloud

`func (o *IamUser) SetSoundcloud(v string)`

SetSoundcloud sets Soundcloud field to given value.

### HasSoundcloud

`func (o *IamUser) HasSoundcloud() bool`

HasSoundcloud returns a boolean if a field has been set.

### GetSpotify

`func (o *IamUser) GetSpotify() string`

GetSpotify returns the Spotify field if non-nil, zero value otherwise.

### GetSpotifyOk

`func (o *IamUser) GetSpotifyOk() (*string, bool)`

GetSpotifyOk returns a tuple with the Spotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotify

`func (o *IamUser) SetSpotify(v string)`

SetSpotify sets Spotify field to given value.

### HasSpotify

`func (o *IamUser) HasSpotify() bool`

HasSpotify returns a boolean if a field has been set.

### GetSteam

`func (o *IamUser) GetSteam() string`

GetSteam returns the Steam field if non-nil, zero value otherwise.

### GetSteamOk

`func (o *IamUser) GetSteamOk() (*string, bool)`

GetSteamOk returns a tuple with the Steam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteam

`func (o *IamUser) SetSteam(v string)`

SetSteam sets Steam field to given value.

### HasSteam

`func (o *IamUser) HasSteam() bool`

HasSteam returns a boolean if a field has been set.

### GetStrava

`func (o *IamUser) GetStrava() string`

GetStrava returns the Strava field if non-nil, zero value otherwise.

### GetStravaOk

`func (o *IamUser) GetStravaOk() (*string, bool)`

GetStravaOk returns a tuple with the Strava field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrava

`func (o *IamUser) SetStrava(v string)`

SetStrava sets Strava field to given value.

### HasStrava

`func (o *IamUser) HasStrava() bool`

HasStrava returns a boolean if a field has been set.

### GetStripe

`func (o *IamUser) GetStripe() string`

GetStripe returns the Stripe field if non-nil, zero value otherwise.

### GetStripeOk

`func (o *IamUser) GetStripeOk() (*string, bool)`

GetStripeOk returns a tuple with the Stripe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripe

`func (o *IamUser) SetStripe(v string)`

SetStripe sets Stripe field to given value.

### HasStripe

`func (o *IamUser) HasStripe() bool`

HasStripe returns a boolean if a field has been set.

### GetTag

`func (o *IamUser) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *IamUser) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *IamUser) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *IamUser) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTelegram

`func (o *IamUser) GetTelegram() string`

GetTelegram returns the Telegram field if non-nil, zero value otherwise.

### GetTelegramOk

`func (o *IamUser) GetTelegramOk() (*string, bool)`

GetTelegramOk returns a tuple with the Telegram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegram

`func (o *IamUser) SetTelegram(v string)`

SetTelegram sets Telegram field to given value.

### HasTelegram

`func (o *IamUser) HasTelegram() bool`

HasTelegram returns a boolean if a field has been set.

### GetTiktok

`func (o *IamUser) GetTiktok() string`

GetTiktok returns the Tiktok field if non-nil, zero value otherwise.

### GetTiktokOk

`func (o *IamUser) GetTiktokOk() (*string, bool)`

GetTiktokOk returns a tuple with the Tiktok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiktok

`func (o *IamUser) SetTiktok(v string)`

SetTiktok sets Tiktok field to given value.

### HasTiktok

`func (o *IamUser) HasTiktok() bool`

HasTiktok returns a boolean if a field has been set.

### GetTitle

`func (o *IamUser) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IamUser) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IamUser) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IamUser) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTotpSecret

`func (o *IamUser) GetTotpSecret() string`

GetTotpSecret returns the TotpSecret field if non-nil, zero value otherwise.

### GetTotpSecretOk

`func (o *IamUser) GetTotpSecretOk() (*string, bool)`

GetTotpSecretOk returns a tuple with the TotpSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpSecret

`func (o *IamUser) SetTotpSecret(v string)`

SetTotpSecret sets TotpSecret field to given value.

### HasTotpSecret

`func (o *IamUser) HasTotpSecret() bool`

HasTotpSecret returns a boolean if a field has been set.

### GetTumblr

`func (o *IamUser) GetTumblr() string`

GetTumblr returns the Tumblr field if non-nil, zero value otherwise.

### GetTumblrOk

`func (o *IamUser) GetTumblrOk() (*string, bool)`

GetTumblrOk returns a tuple with the Tumblr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTumblr

`func (o *IamUser) SetTumblr(v string)`

SetTumblr sets Tumblr field to given value.

### HasTumblr

`func (o *IamUser) HasTumblr() bool`

HasTumblr returns a boolean if a field has been set.

### GetTwitch

`func (o *IamUser) GetTwitch() string`

GetTwitch returns the Twitch field if non-nil, zero value otherwise.

### GetTwitchOk

`func (o *IamUser) GetTwitchOk() (*string, bool)`

GetTwitchOk returns a tuple with the Twitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitch

`func (o *IamUser) SetTwitch(v string)`

SetTwitch sets Twitch field to given value.

### HasTwitch

`func (o *IamUser) HasTwitch() bool`

HasTwitch returns a boolean if a field has been set.

### GetTwitter

`func (o *IamUser) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *IamUser) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *IamUser) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *IamUser) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.

### GetType

`func (o *IamUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IamUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IamUser) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IamUser) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypetalk

`func (o *IamUser) GetTypetalk() string`

GetTypetalk returns the Typetalk field if non-nil, zero value otherwise.

### GetTypetalkOk

`func (o *IamUser) GetTypetalkOk() (*string, bool)`

GetTypetalkOk returns a tuple with the Typetalk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypetalk

`func (o *IamUser) SetTypetalk(v string)`

SetTypetalk sets Typetalk field to given value.

### HasTypetalk

`func (o *IamUser) HasTypetalk() bool`

HasTypetalk returns a boolean if a field has been set.

### GetUber

`func (o *IamUser) GetUber() string`

GetUber returns the Uber field if non-nil, zero value otherwise.

### GetUberOk

`func (o *IamUser) GetUberOk() (*string, bool)`

GetUberOk returns a tuple with the Uber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUber

`func (o *IamUser) SetUber(v string)`

SetUber sets Uber field to given value.

### HasUber

`func (o *IamUser) HasUber() bool`

HasUber returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *IamUser) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *IamUser) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *IamUser) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *IamUser) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetVerificationCode

`func (o *IamUser) GetVerificationCode() string`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *IamUser) GetVerificationCodeOk() (*string, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *IamUser) SetVerificationCode(v string)`

SetVerificationCode sets VerificationCode field to given value.

### HasVerificationCode

`func (o *IamUser) HasVerificationCode() bool`

HasVerificationCode returns a boolean if a field has been set.

### GetVk

`func (o *IamUser) GetVk() string`

GetVk returns the Vk field if non-nil, zero value otherwise.

### GetVkOk

`func (o *IamUser) GetVkOk() (*string, bool)`

GetVkOk returns a tuple with the Vk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVk

`func (o *IamUser) SetVk(v string)`

SetVk sets Vk field to given value.

### HasVk

`func (o *IamUser) HasVk() bool`

HasVk returns a boolean if a field has been set.

### GetWebauthnCredentials

`func (o *IamUser) GetWebauthnCredentials() []interface{}`

GetWebauthnCredentials returns the WebauthnCredentials field if non-nil, zero value otherwise.

### GetWebauthnCredentialsOk

`func (o *IamUser) GetWebauthnCredentialsOk() (*[]interface{}, bool)`

GetWebauthnCredentialsOk returns a tuple with the WebauthnCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebauthnCredentials

`func (o *IamUser) SetWebauthnCredentials(v []interface{})`

SetWebauthnCredentials sets WebauthnCredentials field to given value.

### HasWebauthnCredentials

`func (o *IamUser) HasWebauthnCredentials() bool`

HasWebauthnCredentials returns a boolean if a field has been set.

### GetWechat

`func (o *IamUser) GetWechat() string`

GetWechat returns the Wechat field if non-nil, zero value otherwise.

### GetWechatOk

`func (o *IamUser) GetWechatOk() (*string, bool)`

GetWechatOk returns a tuple with the Wechat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWechat

`func (o *IamUser) SetWechat(v string)`

SetWechat sets Wechat field to given value.

### HasWechat

`func (o *IamUser) HasWechat() bool`

HasWechat returns a boolean if a field has been set.

### GetWecom

`func (o *IamUser) GetWecom() string`

GetWecom returns the Wecom field if non-nil, zero value otherwise.

### GetWecomOk

`func (o *IamUser) GetWecomOk() (*string, bool)`

GetWecomOk returns a tuple with the Wecom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWecom

`func (o *IamUser) SetWecom(v string)`

SetWecom sets Wecom field to given value.

### HasWecom

`func (o *IamUser) HasWecom() bool`

HasWecom returns a boolean if a field has been set.

### GetWeibo

`func (o *IamUser) GetWeibo() string`

GetWeibo returns the Weibo field if non-nil, zero value otherwise.

### GetWeiboOk

`func (o *IamUser) GetWeiboOk() (*string, bool)`

GetWeiboOk returns a tuple with the Weibo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeibo

`func (o *IamUser) SetWeibo(v string)`

SetWeibo sets Weibo field to given value.

### HasWeibo

`func (o *IamUser) HasWeibo() bool`

HasWeibo returns a boolean if a field has been set.

### GetWepay

`func (o *IamUser) GetWepay() string`

GetWepay returns the Wepay field if non-nil, zero value otherwise.

### GetWepayOk

`func (o *IamUser) GetWepayOk() (*string, bool)`

GetWepayOk returns a tuple with the Wepay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWepay

`func (o *IamUser) SetWepay(v string)`

SetWepay sets Wepay field to given value.

### HasWepay

`func (o *IamUser) HasWepay() bool`

HasWepay returns a boolean if a field has been set.

### GetXero

`func (o *IamUser) GetXero() string`

GetXero returns the Xero field if non-nil, zero value otherwise.

### GetXeroOk

`func (o *IamUser) GetXeroOk() (*string, bool)`

GetXeroOk returns a tuple with the Xero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXero

`func (o *IamUser) SetXero(v string)`

SetXero sets Xero field to given value.

### HasXero

`func (o *IamUser) HasXero() bool`

HasXero returns a boolean if a field has been set.

### GetYahoo

`func (o *IamUser) GetYahoo() string`

GetYahoo returns the Yahoo field if non-nil, zero value otherwise.

### GetYahooOk

`func (o *IamUser) GetYahooOk() (*string, bool)`

GetYahooOk returns a tuple with the Yahoo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYahoo

`func (o *IamUser) SetYahoo(v string)`

SetYahoo sets Yahoo field to given value.

### HasYahoo

`func (o *IamUser) HasYahoo() bool`

HasYahoo returns a boolean if a field has been set.

### GetYammer

`func (o *IamUser) GetYammer() string`

GetYammer returns the Yammer field if non-nil, zero value otherwise.

### GetYammerOk

`func (o *IamUser) GetYammerOk() (*string, bool)`

GetYammerOk returns a tuple with the Yammer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYammer

`func (o *IamUser) SetYammer(v string)`

SetYammer sets Yammer field to given value.

### HasYammer

`func (o *IamUser) HasYammer() bool`

HasYammer returns a boolean if a field has been set.

### GetYandex

`func (o *IamUser) GetYandex() string`

GetYandex returns the Yandex field if non-nil, zero value otherwise.

### GetYandexOk

`func (o *IamUser) GetYandexOk() (*string, bool)`

GetYandexOk returns a tuple with the Yandex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYandex

`func (o *IamUser) SetYandex(v string)`

SetYandex sets Yandex field to given value.

### HasYandex

`func (o *IamUser) HasYandex() bool`

HasYandex returns a boolean if a field has been set.

### GetZoom

`func (o *IamUser) GetZoom() string`

GetZoom returns the Zoom field if non-nil, zero value otherwise.

### GetZoomOk

`func (o *IamUser) GetZoomOk() (*string, bool)`

GetZoomOk returns a tuple with the Zoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoom

`func (o *IamUser) SetZoom(v string)`

SetZoom sets Zoom field to given value.

### HasZoom

`func (o *IamUser) HasZoom() bool`

HasZoom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



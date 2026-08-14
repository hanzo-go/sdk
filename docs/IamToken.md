# IamToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**AccessTokenHash** | Pointer to **string** |  | [optional] 
**Application** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**CodeChallenge** | Pointer to **string** |  | [optional] 
**CodeChallengeMethod** | Pointer to **string** |  | [optional] 
**CodeExpireIn** | Pointer to **int32** |  | [optional] 
**CodeIsUsed** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Nonce** | Pointer to **string** | Nonce is the OIDC authorize nonce, stored on the code and echoed into the id_token minted at the exchange (OIDC Core §3.1.3.6) so a relying party binds the id_token to its own request and detects replay. | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**PublicGrant** | Pointer to **bool** | PublicGrant records that this grant was established WITHOUT client authentication — a PKCE code exchange from a client that presented no secret. Whether a client is confidential is a property of the GRANT, not only of the registration: &#x60;hanzo-cli&#x60; and every @hanzo/iam SPA keep a registered secret for a BACKEND path while the surface that actually signs in is a public PKCE client that cannot hold one. authorizationCodeGrant already makes exactly that bounded relaxation; this is the same fact, recorded so refreshTokenGrant can honour it instead of demanding a secret the client never had (which 401s invalid_client and kills the session at the access token&#39;s expiry). Carried across rotation, so the second refresh behaves like the first. | [optional] 
**RedirectUri** | Pointer to **string** | RedirectUri binds the authorization code to the exact redirect URI of the authorize request (RFC 6749 §4.1.3): the token endpoint refuses a code redeemed with a different redirect_uri, closing code-injection across a client&#39;s registered URIs. | [optional] 
**RefreshConsumed** | Pointer to **bool** |  | [optional] 
**RefreshExpireIn** | Pointer to **int32** |  | [optional] 
**RefreshFamily** | Pointer to **string** | Refresh-token rotation state (v2). Each refresh belongs to a family (the grant); rotation mints a new row in the same family and marks the prior one consumed. Presenting a consumed refresh is reuse — the whole family is revoked (RFC 9700 §4.14.2). RefreshExpireIn is the refresh token&#39;s own absolute expiry (unix), independent of the access token&#39;s shorter life. | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**RefreshTokenHash** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** | RFC 8707 resource indicator | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**UserCode** | Pointer to **string** |  | [optional] 

## Methods

### NewIamToken

`func NewIamToken() *IamToken`

NewIamToken instantiates a new IamToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamTokenWithDefaults

`func NewIamTokenWithDefaults() *IamToken`

NewIamTokenWithDefaults instantiates a new IamToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *IamToken) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *IamToken) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *IamToken) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *IamToken) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAccessTokenHash

`func (o *IamToken) GetAccessTokenHash() string`

GetAccessTokenHash returns the AccessTokenHash field if non-nil, zero value otherwise.

### GetAccessTokenHashOk

`func (o *IamToken) GetAccessTokenHashOk() (*string, bool)`

GetAccessTokenHashOk returns a tuple with the AccessTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenHash

`func (o *IamToken) SetAccessTokenHash(v string)`

SetAccessTokenHash sets AccessTokenHash field to given value.

### HasAccessTokenHash

`func (o *IamToken) HasAccessTokenHash() bool`

HasAccessTokenHash returns a boolean if a field has been set.

### GetApplication

`func (o *IamToken) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *IamToken) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *IamToken) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *IamToken) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCode

`func (o *IamToken) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IamToken) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IamToken) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *IamToken) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCodeChallenge

`func (o *IamToken) GetCodeChallenge() string`

GetCodeChallenge returns the CodeChallenge field if non-nil, zero value otherwise.

### GetCodeChallengeOk

`func (o *IamToken) GetCodeChallengeOk() (*string, bool)`

GetCodeChallengeOk returns a tuple with the CodeChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeChallenge

`func (o *IamToken) SetCodeChallenge(v string)`

SetCodeChallenge sets CodeChallenge field to given value.

### HasCodeChallenge

`func (o *IamToken) HasCodeChallenge() bool`

HasCodeChallenge returns a boolean if a field has been set.

### GetCodeChallengeMethod

`func (o *IamToken) GetCodeChallengeMethod() string`

GetCodeChallengeMethod returns the CodeChallengeMethod field if non-nil, zero value otherwise.

### GetCodeChallengeMethodOk

`func (o *IamToken) GetCodeChallengeMethodOk() (*string, bool)`

GetCodeChallengeMethodOk returns a tuple with the CodeChallengeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeChallengeMethod

`func (o *IamToken) SetCodeChallengeMethod(v string)`

SetCodeChallengeMethod sets CodeChallengeMethod field to given value.

### HasCodeChallengeMethod

`func (o *IamToken) HasCodeChallengeMethod() bool`

HasCodeChallengeMethod returns a boolean if a field has been set.

### GetCodeExpireIn

`func (o *IamToken) GetCodeExpireIn() int32`

GetCodeExpireIn returns the CodeExpireIn field if non-nil, zero value otherwise.

### GetCodeExpireInOk

`func (o *IamToken) GetCodeExpireInOk() (*int32, bool)`

GetCodeExpireInOk returns a tuple with the CodeExpireIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeExpireIn

`func (o *IamToken) SetCodeExpireIn(v int32)`

SetCodeExpireIn sets CodeExpireIn field to given value.

### HasCodeExpireIn

`func (o *IamToken) HasCodeExpireIn() bool`

HasCodeExpireIn returns a boolean if a field has been set.

### GetCodeIsUsed

`func (o *IamToken) GetCodeIsUsed() bool`

GetCodeIsUsed returns the CodeIsUsed field if non-nil, zero value otherwise.

### GetCodeIsUsedOk

`func (o *IamToken) GetCodeIsUsedOk() (*bool, bool)`

GetCodeIsUsedOk returns a tuple with the CodeIsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeIsUsed

`func (o *IamToken) SetCodeIsUsed(v bool)`

SetCodeIsUsed sets CodeIsUsed field to given value.

### HasCodeIsUsed

`func (o *IamToken) HasCodeIsUsed() bool`

HasCodeIsUsed returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IamToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamToken) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamToken) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamToken) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamToken) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeleted

`func (o *IamToken) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamToken) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamToken) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamToken) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetExpiresIn

`func (o *IamToken) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *IamToken) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *IamToken) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *IamToken) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetId

`func (o *IamToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IamToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNonce

`func (o *IamToken) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *IamToken) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *IamToken) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *IamToken) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetOrganization

`func (o *IamToken) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamToken) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamToken) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamToken) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOwner

`func (o *IamToken) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamToken) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamToken) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamToken) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPublicGrant

`func (o *IamToken) GetPublicGrant() bool`

GetPublicGrant returns the PublicGrant field if non-nil, zero value otherwise.

### GetPublicGrantOk

`func (o *IamToken) GetPublicGrantOk() (*bool, bool)`

GetPublicGrantOk returns a tuple with the PublicGrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicGrant

`func (o *IamToken) SetPublicGrant(v bool)`

SetPublicGrant sets PublicGrant field to given value.

### HasPublicGrant

`func (o *IamToken) HasPublicGrant() bool`

HasPublicGrant returns a boolean if a field has been set.

### GetRedirectUri

`func (o *IamToken) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *IamToken) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *IamToken) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *IamToken) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetRefreshConsumed

`func (o *IamToken) GetRefreshConsumed() bool`

GetRefreshConsumed returns the RefreshConsumed field if non-nil, zero value otherwise.

### GetRefreshConsumedOk

`func (o *IamToken) GetRefreshConsumedOk() (*bool, bool)`

GetRefreshConsumedOk returns a tuple with the RefreshConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshConsumed

`func (o *IamToken) SetRefreshConsumed(v bool)`

SetRefreshConsumed sets RefreshConsumed field to given value.

### HasRefreshConsumed

`func (o *IamToken) HasRefreshConsumed() bool`

HasRefreshConsumed returns a boolean if a field has been set.

### GetRefreshExpireIn

`func (o *IamToken) GetRefreshExpireIn() int32`

GetRefreshExpireIn returns the RefreshExpireIn field if non-nil, zero value otherwise.

### GetRefreshExpireInOk

`func (o *IamToken) GetRefreshExpireInOk() (*int32, bool)`

GetRefreshExpireInOk returns a tuple with the RefreshExpireIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshExpireIn

`func (o *IamToken) SetRefreshExpireIn(v int32)`

SetRefreshExpireIn sets RefreshExpireIn field to given value.

### HasRefreshExpireIn

`func (o *IamToken) HasRefreshExpireIn() bool`

HasRefreshExpireIn returns a boolean if a field has been set.

### GetRefreshFamily

`func (o *IamToken) GetRefreshFamily() string`

GetRefreshFamily returns the RefreshFamily field if non-nil, zero value otherwise.

### GetRefreshFamilyOk

`func (o *IamToken) GetRefreshFamilyOk() (*string, bool)`

GetRefreshFamilyOk returns a tuple with the RefreshFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshFamily

`func (o *IamToken) SetRefreshFamily(v string)`

SetRefreshFamily sets RefreshFamily field to given value.

### HasRefreshFamily

`func (o *IamToken) HasRefreshFamily() bool`

HasRefreshFamily returns a boolean if a field has been set.

### GetRefreshToken

`func (o *IamToken) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *IamToken) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *IamToken) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *IamToken) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetRefreshTokenHash

`func (o *IamToken) GetRefreshTokenHash() string`

GetRefreshTokenHash returns the RefreshTokenHash field if non-nil, zero value otherwise.

### GetRefreshTokenHashOk

`func (o *IamToken) GetRefreshTokenHashOk() (*string, bool)`

GetRefreshTokenHashOk returns a tuple with the RefreshTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenHash

`func (o *IamToken) SetRefreshTokenHash(v string)`

SetRefreshTokenHash sets RefreshTokenHash field to given value.

### HasRefreshTokenHash

`func (o *IamToken) HasRefreshTokenHash() bool`

HasRefreshTokenHash returns a boolean if a field has been set.

### GetResource

`func (o *IamToken) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IamToken) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IamToken) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *IamToken) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetScope

`func (o *IamToken) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IamToken) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IamToken) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *IamToken) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTokenType

`func (o *IamToken) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *IamToken) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *IamToken) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *IamToken) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamToken) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamToken) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamToken) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamToken) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *IamToken) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamToken) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamToken) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *IamToken) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserCode

`func (o *IamToken) GetUserCode() string`

GetUserCode returns the UserCode field if non-nil, zero value otherwise.

### GetUserCodeOk

`func (o *IamToken) GetUserCodeOk() (*string, bool)`

GetUserCodeOk returns a tuple with the UserCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCode

`func (o *IamToken) SetUserCode(v string)`

SetUserCode sets UserCode field to given value.

### HasUserCode

`func (o *IamToken) HasUserCode() bool`

HasUserCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



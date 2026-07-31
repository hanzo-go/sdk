# IamControllersTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assertion** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**ClientAssertion** | Pointer to **string** |  | [optional] 
**ClientAssertionType** | Pointer to **string** |  | [optional] 
**GrantType** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**CodeVerifier** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Nonce** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**SubjectToken** | Pointer to **string** |  | [optional] 
**SubjectTokenType** | Pointer to **string** |  | [optional] 
**Audience** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** | RFC 8707 resource indicator | [optional] 
**AccessKey** | Pointer to **string** |  | [optional] 
**AccessSecret** | Pointer to **string** |  | [optional] 

## Methods

### NewIamControllersTokenRequest

`func NewIamControllersTokenRequest() *IamControllersTokenRequest`

NewIamControllersTokenRequest instantiates a new IamControllersTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamControllersTokenRequestWithDefaults

`func NewIamControllersTokenRequestWithDefaults() *IamControllersTokenRequest`

NewIamControllersTokenRequestWithDefaults instantiates a new IamControllersTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertion

`func (o *IamControllersTokenRequest) GetAssertion() string`

GetAssertion returns the Assertion field if non-nil, zero value otherwise.

### GetAssertionOk

`func (o *IamControllersTokenRequest) GetAssertionOk() (*string, bool)`

GetAssertionOk returns a tuple with the Assertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertion

`func (o *IamControllersTokenRequest) SetAssertion(v string)`

SetAssertion sets Assertion field to given value.

### HasAssertion

`func (o *IamControllersTokenRequest) HasAssertion() bool`

HasAssertion returns a boolean if a field has been set.

### GetClientId

`func (o *IamControllersTokenRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IamControllersTokenRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IamControllersTokenRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IamControllersTokenRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *IamControllersTokenRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IamControllersTokenRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IamControllersTokenRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *IamControllersTokenRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientAssertion

`func (o *IamControllersTokenRequest) GetClientAssertion() string`

GetClientAssertion returns the ClientAssertion field if non-nil, zero value otherwise.

### GetClientAssertionOk

`func (o *IamControllersTokenRequest) GetClientAssertionOk() (*string, bool)`

GetClientAssertionOk returns a tuple with the ClientAssertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAssertion

`func (o *IamControllersTokenRequest) SetClientAssertion(v string)`

SetClientAssertion sets ClientAssertion field to given value.

### HasClientAssertion

`func (o *IamControllersTokenRequest) HasClientAssertion() bool`

HasClientAssertion returns a boolean if a field has been set.

### GetClientAssertionType

`func (o *IamControllersTokenRequest) GetClientAssertionType() string`

GetClientAssertionType returns the ClientAssertionType field if non-nil, zero value otherwise.

### GetClientAssertionTypeOk

`func (o *IamControllersTokenRequest) GetClientAssertionTypeOk() (*string, bool)`

GetClientAssertionTypeOk returns a tuple with the ClientAssertionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAssertionType

`func (o *IamControllersTokenRequest) SetClientAssertionType(v string)`

SetClientAssertionType sets ClientAssertionType field to given value.

### HasClientAssertionType

`func (o *IamControllersTokenRequest) HasClientAssertionType() bool`

HasClientAssertionType returns a boolean if a field has been set.

### GetGrantType

`func (o *IamControllersTokenRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *IamControllersTokenRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *IamControllersTokenRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *IamControllersTokenRequest) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetCode

`func (o *IamControllersTokenRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IamControllersTokenRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IamControllersTokenRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *IamControllersTokenRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCodeVerifier

`func (o *IamControllersTokenRequest) GetCodeVerifier() string`

GetCodeVerifier returns the CodeVerifier field if non-nil, zero value otherwise.

### GetCodeVerifierOk

`func (o *IamControllersTokenRequest) GetCodeVerifierOk() (*string, bool)`

GetCodeVerifierOk returns a tuple with the CodeVerifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeVerifier

`func (o *IamControllersTokenRequest) SetCodeVerifier(v string)`

SetCodeVerifier sets CodeVerifier field to given value.

### HasCodeVerifier

`func (o *IamControllersTokenRequest) HasCodeVerifier() bool`

HasCodeVerifier returns a boolean if a field has been set.

### GetScope

`func (o *IamControllersTokenRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IamControllersTokenRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IamControllersTokenRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *IamControllersTokenRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetNonce

`func (o *IamControllersTokenRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *IamControllersTokenRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *IamControllersTokenRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *IamControllersTokenRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetUsername

`func (o *IamControllersTokenRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IamControllersTokenRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IamControllersTokenRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IamControllersTokenRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *IamControllersTokenRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IamControllersTokenRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IamControllersTokenRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IamControllersTokenRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTag

`func (o *IamControllersTokenRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *IamControllersTokenRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *IamControllersTokenRequest) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *IamControllersTokenRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetAvatar

`func (o *IamControllersTokenRequest) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *IamControllersTokenRequest) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *IamControllersTokenRequest) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *IamControllersTokenRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetRefreshToken

`func (o *IamControllersTokenRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *IamControllersTokenRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *IamControllersTokenRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *IamControllersTokenRequest) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetSubjectToken

`func (o *IamControllersTokenRequest) GetSubjectToken() string`

GetSubjectToken returns the SubjectToken field if non-nil, zero value otherwise.

### GetSubjectTokenOk

`func (o *IamControllersTokenRequest) GetSubjectTokenOk() (*string, bool)`

GetSubjectTokenOk returns a tuple with the SubjectToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectToken

`func (o *IamControllersTokenRequest) SetSubjectToken(v string)`

SetSubjectToken sets SubjectToken field to given value.

### HasSubjectToken

`func (o *IamControllersTokenRequest) HasSubjectToken() bool`

HasSubjectToken returns a boolean if a field has been set.

### GetSubjectTokenType

`func (o *IamControllersTokenRequest) GetSubjectTokenType() string`

GetSubjectTokenType returns the SubjectTokenType field if non-nil, zero value otherwise.

### GetSubjectTokenTypeOk

`func (o *IamControllersTokenRequest) GetSubjectTokenTypeOk() (*string, bool)`

GetSubjectTokenTypeOk returns a tuple with the SubjectTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectTokenType

`func (o *IamControllersTokenRequest) SetSubjectTokenType(v string)`

SetSubjectTokenType sets SubjectTokenType field to given value.

### HasSubjectTokenType

`func (o *IamControllersTokenRequest) HasSubjectTokenType() bool`

HasSubjectTokenType returns a boolean if a field has been set.

### GetAudience

`func (o *IamControllersTokenRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *IamControllersTokenRequest) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *IamControllersTokenRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *IamControllersTokenRequest) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetResource

`func (o *IamControllersTokenRequest) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IamControllersTokenRequest) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IamControllersTokenRequest) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *IamControllersTokenRequest) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetAccessKey

`func (o *IamControllersTokenRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *IamControllersTokenRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *IamControllersTokenRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *IamControllersTokenRequest) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessSecret

`func (o *IamControllersTokenRequest) GetAccessSecret() string`

GetAccessSecret returns the AccessSecret field if non-nil, zero value otherwise.

### GetAccessSecretOk

`func (o *IamControllersTokenRequest) GetAccessSecretOk() (*string, bool)`

GetAccessSecretOk returns a tuple with the AccessSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSecret

`func (o *IamControllersTokenRequest) SetAccessSecret(v string)`

SetAccessSecret sets AccessSecret field to given value.

### HasAccessSecret

`func (o *IamControllersTokenRequest) HasAccessSecret() bool`

HasAccessSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



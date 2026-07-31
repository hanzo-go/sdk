# IamObjectToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**AccessTokenHash** | Pointer to **string** |  | [optional] 
**Application** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**CodeChallenge** | Pointer to **string** |  | [optional] 
**CodeExpireIn** | Pointer to **int64** |  | [optional] 
**CodeIsUsed** | Pointer to **bool** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**ExpiresIn** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**RefreshTokenHash** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectToken

`func NewIamObjectToken() *IamObjectToken`

NewIamObjectToken instantiates a new IamObjectToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectTokenWithDefaults

`func NewIamObjectTokenWithDefaults() *IamObjectToken`

NewIamObjectTokenWithDefaults instantiates a new IamObjectToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *IamObjectToken) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *IamObjectToken) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *IamObjectToken) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *IamObjectToken) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAccessTokenHash

`func (o *IamObjectToken) GetAccessTokenHash() string`

GetAccessTokenHash returns the AccessTokenHash field if non-nil, zero value otherwise.

### GetAccessTokenHashOk

`func (o *IamObjectToken) GetAccessTokenHashOk() (*string, bool)`

GetAccessTokenHashOk returns a tuple with the AccessTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenHash

`func (o *IamObjectToken) SetAccessTokenHash(v string)`

SetAccessTokenHash sets AccessTokenHash field to given value.

### HasAccessTokenHash

`func (o *IamObjectToken) HasAccessTokenHash() bool`

HasAccessTokenHash returns a boolean if a field has been set.

### GetApplication

`func (o *IamObjectToken) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *IamObjectToken) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *IamObjectToken) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *IamObjectToken) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCode

`func (o *IamObjectToken) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IamObjectToken) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IamObjectToken) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *IamObjectToken) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCodeChallenge

`func (o *IamObjectToken) GetCodeChallenge() string`

GetCodeChallenge returns the CodeChallenge field if non-nil, zero value otherwise.

### GetCodeChallengeOk

`func (o *IamObjectToken) GetCodeChallengeOk() (*string, bool)`

GetCodeChallengeOk returns a tuple with the CodeChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeChallenge

`func (o *IamObjectToken) SetCodeChallenge(v string)`

SetCodeChallenge sets CodeChallenge field to given value.

### HasCodeChallenge

`func (o *IamObjectToken) HasCodeChallenge() bool`

HasCodeChallenge returns a boolean if a field has been set.

### GetCodeExpireIn

`func (o *IamObjectToken) GetCodeExpireIn() int64`

GetCodeExpireIn returns the CodeExpireIn field if non-nil, zero value otherwise.

### GetCodeExpireInOk

`func (o *IamObjectToken) GetCodeExpireInOk() (*int64, bool)`

GetCodeExpireInOk returns a tuple with the CodeExpireIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeExpireIn

`func (o *IamObjectToken) SetCodeExpireIn(v int64)`

SetCodeExpireIn sets CodeExpireIn field to given value.

### HasCodeExpireIn

`func (o *IamObjectToken) HasCodeExpireIn() bool`

HasCodeExpireIn returns a boolean if a field has been set.

### GetCodeIsUsed

`func (o *IamObjectToken) GetCodeIsUsed() bool`

GetCodeIsUsed returns the CodeIsUsed field if non-nil, zero value otherwise.

### GetCodeIsUsedOk

`func (o *IamObjectToken) GetCodeIsUsedOk() (*bool, bool)`

GetCodeIsUsedOk returns a tuple with the CodeIsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeIsUsed

`func (o *IamObjectToken) SetCodeIsUsed(v bool)`

SetCodeIsUsed sets CodeIsUsed field to given value.

### HasCodeIsUsed

`func (o *IamObjectToken) HasCodeIsUsed() bool`

HasCodeIsUsed returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamObjectToken) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectToken) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectToken) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectToken) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetExpiresIn

`func (o *IamObjectToken) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *IamObjectToken) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *IamObjectToken) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *IamObjectToken) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetName

`func (o *IamObjectToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *IamObjectToken) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamObjectToken) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamObjectToken) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamObjectToken) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectToken) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectToken) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectToken) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectToken) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRefreshToken

`func (o *IamObjectToken) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *IamObjectToken) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *IamObjectToken) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *IamObjectToken) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetRefreshTokenHash

`func (o *IamObjectToken) GetRefreshTokenHash() string`

GetRefreshTokenHash returns the RefreshTokenHash field if non-nil, zero value otherwise.

### GetRefreshTokenHashOk

`func (o *IamObjectToken) GetRefreshTokenHashOk() (*string, bool)`

GetRefreshTokenHashOk returns a tuple with the RefreshTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenHash

`func (o *IamObjectToken) SetRefreshTokenHash(v string)`

SetRefreshTokenHash sets RefreshTokenHash field to given value.

### HasRefreshTokenHash

`func (o *IamObjectToken) HasRefreshTokenHash() bool`

HasRefreshTokenHash returns a boolean if a field has been set.

### GetScope

`func (o *IamObjectToken) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IamObjectToken) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IamObjectToken) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *IamObjectToken) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTokenType

`func (o *IamObjectToken) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *IamObjectToken) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *IamObjectToken) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *IamObjectToken) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetUser

`func (o *IamObjectToken) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamObjectToken) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamObjectToken) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *IamObjectToken) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



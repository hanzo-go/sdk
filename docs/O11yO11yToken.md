# O11yO11yToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | AccessToken authenticates requests until it expires. | [optional] 
**ExpiresIn** | Pointer to **int32** | ExpiresIn is the access token&#39;s lifetime in seconds. | [optional] 
**RefreshToken** | Pointer to **string** | RefreshToken buys the next pair via rotateSession. | [optional] 
**TokenType** | Pointer to **string** | TokenType is how to present the access token, e.g. bearer. | [optional] 

## Methods

### NewO11yO11yToken

`func NewO11yO11yToken() *O11yO11yToken`

NewO11yO11yToken instantiates a new O11yO11yToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yTokenWithDefaults

`func NewO11yO11yTokenWithDefaults() *O11yO11yToken`

NewO11yO11yTokenWithDefaults instantiates a new O11yO11yToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *O11yO11yToken) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *O11yO11yToken) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *O11yO11yToken) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *O11yO11yToken) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *O11yO11yToken) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *O11yO11yToken) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *O11yO11yToken) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *O11yO11yToken) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetRefreshToken

`func (o *O11yO11yToken) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *O11yO11yToken) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *O11yO11yToken) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *O11yO11yToken) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetTokenType

`func (o *O11yO11yToken) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *O11yO11yToken) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *O11yO11yToken) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *O11yO11yToken) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



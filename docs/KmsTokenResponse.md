# KmsTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**AccessTokenMaxTTL** | Pointer to **int32** |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] [default to "Bearer"]

## Methods

### NewKmsTokenResponse

`func NewKmsTokenResponse() *KmsTokenResponse`

NewKmsTokenResponse instantiates a new KmsTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsTokenResponseWithDefaults

`func NewKmsTokenResponseWithDefaults() *KmsTokenResponse`

NewKmsTokenResponseWithDefaults instantiates a new KmsTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *KmsTokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *KmsTokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *KmsTokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *KmsTokenResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *KmsTokenResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *KmsTokenResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *KmsTokenResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *KmsTokenResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetAccessTokenMaxTTL

`func (o *KmsTokenResponse) GetAccessTokenMaxTTL() int32`

GetAccessTokenMaxTTL returns the AccessTokenMaxTTL field if non-nil, zero value otherwise.

### GetAccessTokenMaxTTLOk

`func (o *KmsTokenResponse) GetAccessTokenMaxTTLOk() (*int32, bool)`

GetAccessTokenMaxTTLOk returns a tuple with the AccessTokenMaxTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenMaxTTL

`func (o *KmsTokenResponse) SetAccessTokenMaxTTL(v int32)`

SetAccessTokenMaxTTL sets AccessTokenMaxTTL field to given value.

### HasAccessTokenMaxTTL

`func (o *KmsTokenResponse) HasAccessTokenMaxTTL() bool`

HasAccessTokenMaxTTL returns a boolean if a field has been set.

### GetTokenType

`func (o *KmsTokenResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *KmsTokenResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *KmsTokenResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *KmsTokenResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



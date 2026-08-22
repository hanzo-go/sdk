# KmsToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | AccessToken is IAM&#39;s own JWT, verbatim. Its &#x60;owner&#x60; claim scopes it to exactly one org, which is what every secret operation then reads. | [optional] 
**ExpiresIn** | Pointer to **int32** | ExpiresIn is the token&#39;s lifetime in seconds, as IAM reported it. | [optional] 
**TokenType** | Pointer to **string** | TokenType is &#x60;Bearer&#x60;. | [optional] 

## Methods

### NewKmsToken

`func NewKmsToken() *KmsToken`

NewKmsToken instantiates a new KmsToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsTokenWithDefaults

`func NewKmsTokenWithDefaults() *KmsToken`

NewKmsTokenWithDefaults instantiates a new KmsToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *KmsToken) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *KmsToken) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *KmsToken) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *KmsToken) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *KmsToken) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *KmsToken) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *KmsToken) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *KmsToken) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetTokenType

`func (o *KmsToken) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *KmsToken) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *KmsToken) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *KmsToken) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



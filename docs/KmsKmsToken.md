# KmsKmsToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | AccessToken is IAM&#39;s own JWT, verbatim. Its &#x60;owner&#x60; claim scopes it to exactly one org, which is what every secret operation then reads. | [optional] 
**ExpiresIn** | Pointer to **int64** | ExpiresIn is the token&#39;s lifetime in seconds, as IAM reported it. | [optional] 
**TokenType** | Pointer to **string** | TokenType is &#x60;Bearer&#x60;. | [optional] 

## Methods

### NewKmsKmsToken

`func NewKmsKmsToken() *KmsKmsToken`

NewKmsKmsToken instantiates a new KmsKmsToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsKmsTokenWithDefaults

`func NewKmsKmsTokenWithDefaults() *KmsKmsToken`

NewKmsKmsTokenWithDefaults instantiates a new KmsKmsToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *KmsKmsToken) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *KmsKmsToken) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *KmsKmsToken) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *KmsKmsToken) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *KmsKmsToken) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *KmsKmsToken) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *KmsKmsToken) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *KmsKmsToken) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetTokenType

`func (o *KmsKmsToken) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *KmsKmsToken) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *KmsKmsToken) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *KmsKmsToken) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



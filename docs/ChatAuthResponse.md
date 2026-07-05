# ChatAuthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**ChatUser**](ChatUser.md) |  | [optional] 

## Methods

### NewChatAuthResponse

`func NewChatAuthResponse() *ChatAuthResponse`

NewChatAuthResponse instantiates a new ChatAuthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatAuthResponseWithDefaults

`func NewChatAuthResponseWithDefaults() *ChatAuthResponse`

NewChatAuthResponseWithDefaults instantiates a new ChatAuthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ChatAuthResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ChatAuthResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ChatAuthResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ChatAuthResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *ChatAuthResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *ChatAuthResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *ChatAuthResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *ChatAuthResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetUser

`func (o *ChatAuthResponse) GetUser() ChatUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChatAuthResponse) GetUserOk() (*ChatUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChatAuthResponse) SetUser(v ChatUser)`

SetUser sets User field to given value.

### HasUser

`func (o *ChatAuthResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



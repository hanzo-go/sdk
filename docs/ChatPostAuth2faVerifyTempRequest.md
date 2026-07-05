# ChatPostAuth2faVerifyTempRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TempToken** | **string** |  | 
**Token** | **string** |  | 
**BackupCode** | Pointer to **string** |  | [optional] 

## Methods

### NewChatPostAuth2faVerifyTempRequest

`func NewChatPostAuth2faVerifyTempRequest(tempToken string, token string, ) *ChatPostAuth2faVerifyTempRequest`

NewChatPostAuth2faVerifyTempRequest instantiates a new ChatPostAuth2faVerifyTempRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatPostAuth2faVerifyTempRequestWithDefaults

`func NewChatPostAuth2faVerifyTempRequestWithDefaults() *ChatPostAuth2faVerifyTempRequest`

NewChatPostAuth2faVerifyTempRequestWithDefaults instantiates a new ChatPostAuth2faVerifyTempRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTempToken

`func (o *ChatPostAuth2faVerifyTempRequest) GetTempToken() string`

GetTempToken returns the TempToken field if non-nil, zero value otherwise.

### GetTempTokenOk

`func (o *ChatPostAuth2faVerifyTempRequest) GetTempTokenOk() (*string, bool)`

GetTempTokenOk returns a tuple with the TempToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempToken

`func (o *ChatPostAuth2faVerifyTempRequest) SetTempToken(v string)`

SetTempToken sets TempToken field to given value.


### GetToken

`func (o *ChatPostAuth2faVerifyTempRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ChatPostAuth2faVerifyTempRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ChatPostAuth2faVerifyTempRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetBackupCode

`func (o *ChatPostAuth2faVerifyTempRequest) GetBackupCode() string`

GetBackupCode returns the BackupCode field if non-nil, zero value otherwise.

### GetBackupCodeOk

`func (o *ChatPostAuth2faVerifyTempRequest) GetBackupCodeOk() (*string, bool)`

GetBackupCodeOk returns a tuple with the BackupCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCode

`func (o *ChatPostAuth2faVerifyTempRequest) SetBackupCode(v string)`

SetBackupCode sets BackupCode field to given value.

### HasBackupCode

`func (o *ChatPostAuth2faVerifyTempRequest) HasBackupCode() bool`

HasBackupCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



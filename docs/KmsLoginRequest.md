# KmsLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | **string** |  | 
**ProviderAuthToken** | Pointer to **string** |  | [optional] 

## Methods

### NewKmsLoginRequest

`func NewKmsLoginRequest(email string, password string, ) *KmsLoginRequest`

NewKmsLoginRequest instantiates a new KmsLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsLoginRequestWithDefaults

`func NewKmsLoginRequestWithDefaults() *KmsLoginRequest`

NewKmsLoginRequestWithDefaults instantiates a new KmsLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *KmsLoginRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *KmsLoginRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *KmsLoginRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *KmsLoginRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *KmsLoginRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *KmsLoginRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetProviderAuthToken

`func (o *KmsLoginRequest) GetProviderAuthToken() string`

GetProviderAuthToken returns the ProviderAuthToken field if non-nil, zero value otherwise.

### GetProviderAuthTokenOk

`func (o *KmsLoginRequest) GetProviderAuthTokenOk() (*string, bool)`

GetProviderAuthTokenOk returns a tuple with the ProviderAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAuthToken

`func (o *KmsLoginRequest) SetProviderAuthToken(v string)`

SetProviderAuthToken sets ProviderAuthToken field to given value.

### HasProviderAuthToken

`func (o *KmsLoginRequest) HasProviderAuthToken() bool`

HasProviderAuthToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



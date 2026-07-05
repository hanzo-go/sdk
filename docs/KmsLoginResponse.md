# KmsLoginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**IsMfaEnabled** | Pointer to **bool** |  | [optional] 
**MfaMethod** | Pointer to **string** |  | [optional] 

## Methods

### NewKmsLoginResponse

`func NewKmsLoginResponse() *KmsLoginResponse`

NewKmsLoginResponse instantiates a new KmsLoginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsLoginResponseWithDefaults

`func NewKmsLoginResponseWithDefaults() *KmsLoginResponse`

NewKmsLoginResponseWithDefaults instantiates a new KmsLoginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *KmsLoginResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KmsLoginResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KmsLoginResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *KmsLoginResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetIsMfaEnabled

`func (o *KmsLoginResponse) GetIsMfaEnabled() bool`

GetIsMfaEnabled returns the IsMfaEnabled field if non-nil, zero value otherwise.

### GetIsMfaEnabledOk

`func (o *KmsLoginResponse) GetIsMfaEnabledOk() (*bool, bool)`

GetIsMfaEnabledOk returns a tuple with the IsMfaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMfaEnabled

`func (o *KmsLoginResponse) SetIsMfaEnabled(v bool)`

SetIsMfaEnabled sets IsMfaEnabled field to given value.

### HasIsMfaEnabled

`func (o *KmsLoginResponse) HasIsMfaEnabled() bool`

HasIsMfaEnabled returns a boolean if a field has been set.

### GetMfaMethod

`func (o *KmsLoginResponse) GetMfaMethod() string`

GetMfaMethod returns the MfaMethod field if non-nil, zero value otherwise.

### GetMfaMethodOk

`func (o *KmsLoginResponse) GetMfaMethodOk() (*string, bool)`

GetMfaMethodOk returns a tuple with the MfaMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaMethod

`func (o *KmsLoginResponse) SetMfaMethod(v string)`

SetMfaMethod sets MfaMethod field to given value.

### HasMfaMethod

`func (o *KmsLoginResponse) HasMfaMethod() bool`

HasMfaMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



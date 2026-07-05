# KmsCreateSharedSecretRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretValue** | **string** |  | 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAfterViews** | Pointer to **int32** |  | [optional] 
**AccessType** | Pointer to **string** |  | [optional] [default to "anyone"]

## Methods

### NewKmsCreateSharedSecretRequest

`func NewKmsCreateSharedSecretRequest(secretValue string, ) *KmsCreateSharedSecretRequest`

NewKmsCreateSharedSecretRequest instantiates a new KmsCreateSharedSecretRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCreateSharedSecretRequestWithDefaults

`func NewKmsCreateSharedSecretRequestWithDefaults() *KmsCreateSharedSecretRequest`

NewKmsCreateSharedSecretRequestWithDefaults instantiates a new KmsCreateSharedSecretRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretValue

`func (o *KmsCreateSharedSecretRequest) GetSecretValue() string`

GetSecretValue returns the SecretValue field if non-nil, zero value otherwise.

### GetSecretValueOk

`func (o *KmsCreateSharedSecretRequest) GetSecretValueOk() (*string, bool)`

GetSecretValueOk returns a tuple with the SecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretValue

`func (o *KmsCreateSharedSecretRequest) SetSecretValue(v string)`

SetSecretValue sets SecretValue field to given value.


### GetExpiresAt

`func (o *KmsCreateSharedSecretRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *KmsCreateSharedSecretRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *KmsCreateSharedSecretRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *KmsCreateSharedSecretRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetExpiresAfterViews

`func (o *KmsCreateSharedSecretRequest) GetExpiresAfterViews() int32`

GetExpiresAfterViews returns the ExpiresAfterViews field if non-nil, zero value otherwise.

### GetExpiresAfterViewsOk

`func (o *KmsCreateSharedSecretRequest) GetExpiresAfterViewsOk() (*int32, bool)`

GetExpiresAfterViewsOk returns a tuple with the ExpiresAfterViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAfterViews

`func (o *KmsCreateSharedSecretRequest) SetExpiresAfterViews(v int32)`

SetExpiresAfterViews sets ExpiresAfterViews field to given value.

### HasExpiresAfterViews

`func (o *KmsCreateSharedSecretRequest) HasExpiresAfterViews() bool`

HasExpiresAfterViews returns a boolean if a field has been set.

### GetAccessType

`func (o *KmsCreateSharedSecretRequest) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *KmsCreateSharedSecretRequest) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *KmsCreateSharedSecretRequest) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *KmsCreateSharedSecretRequest) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



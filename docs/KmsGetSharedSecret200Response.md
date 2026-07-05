# KmsGetSharedSecret200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretValue** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewKmsGetSharedSecret200Response

`func NewKmsGetSharedSecret200Response() *KmsGetSharedSecret200Response`

NewKmsGetSharedSecret200Response instantiates a new KmsGetSharedSecret200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsGetSharedSecret200ResponseWithDefaults

`func NewKmsGetSharedSecret200ResponseWithDefaults() *KmsGetSharedSecret200Response`

NewKmsGetSharedSecret200ResponseWithDefaults instantiates a new KmsGetSharedSecret200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretValue

`func (o *KmsGetSharedSecret200Response) GetSecretValue() string`

GetSecretValue returns the SecretValue field if non-nil, zero value otherwise.

### GetSecretValueOk

`func (o *KmsGetSharedSecret200Response) GetSecretValueOk() (*string, bool)`

GetSecretValueOk returns a tuple with the SecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretValue

`func (o *KmsGetSharedSecret200Response) SetSecretValue(v string)`

SetSecretValue sets SecretValue field to given value.

### HasSecretValue

`func (o *KmsGetSharedSecret200Response) HasSecretValue() bool`

HasSecretValue returns a boolean if a field has been set.

### GetExpiresAt

`func (o *KmsGetSharedSecret200Response) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *KmsGetSharedSecret200Response) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *KmsGetSharedSecret200Response) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *KmsGetSharedSecret200Response) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



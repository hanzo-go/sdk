# KmsGetSecret200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to [**KmsSecret**](KmsSecret.md) |  | [optional] 

## Methods

### NewKmsGetSecret200Response

`func NewKmsGetSecret200Response() *KmsGetSecret200Response`

NewKmsGetSecret200Response instantiates a new KmsGetSecret200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsGetSecret200ResponseWithDefaults

`func NewKmsGetSecret200ResponseWithDefaults() *KmsGetSecret200Response`

NewKmsGetSecret200ResponseWithDefaults instantiates a new KmsGetSecret200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *KmsGetSecret200Response) GetSecret() KmsSecret`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *KmsGetSecret200Response) GetSecretOk() (*KmsSecret, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *KmsGetSecret200Response) SetSecret(v KmsSecret)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *KmsGetSecret200Response) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



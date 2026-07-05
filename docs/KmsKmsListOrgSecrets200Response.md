# KmsKmsListOrgSecrets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secrets** | Pointer to [**[]KmsKmsListOrgSecrets200ResponseSecretsInner**](KmsKmsListOrgSecrets200ResponseSecretsInner.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Truncated** | Pointer to **bool** |  | [optional] 

## Methods

### NewKmsKmsListOrgSecrets200Response

`func NewKmsKmsListOrgSecrets200Response() *KmsKmsListOrgSecrets200Response`

NewKmsKmsListOrgSecrets200Response instantiates a new KmsKmsListOrgSecrets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsKmsListOrgSecrets200ResponseWithDefaults

`func NewKmsKmsListOrgSecrets200ResponseWithDefaults() *KmsKmsListOrgSecrets200Response`

NewKmsKmsListOrgSecrets200ResponseWithDefaults instantiates a new KmsKmsListOrgSecrets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecrets

`func (o *KmsKmsListOrgSecrets200Response) GetSecrets() []KmsKmsListOrgSecrets200ResponseSecretsInner`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *KmsKmsListOrgSecrets200Response) GetSecretsOk() (*[]KmsKmsListOrgSecrets200ResponseSecretsInner, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *KmsKmsListOrgSecrets200Response) SetSecrets(v []KmsKmsListOrgSecrets200ResponseSecretsInner)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *KmsKmsListOrgSecrets200Response) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetCount

`func (o *KmsKmsListOrgSecrets200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *KmsKmsListOrgSecrets200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *KmsKmsListOrgSecrets200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *KmsKmsListOrgSecrets200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTruncated

`func (o *KmsKmsListOrgSecrets200Response) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *KmsKmsListOrgSecrets200Response) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *KmsKmsListOrgSecrets200Response) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *KmsKmsListOrgSecrets200Response) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



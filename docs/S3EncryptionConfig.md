# S3EncryptionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SseAlgorithm** | Pointer to **string** |  | [optional] 
**KmsMasterKeyId** | Pointer to **string** | KMS key ID for SSE-KMS | [optional] 

## Methods

### NewS3EncryptionConfig

`func NewS3EncryptionConfig() *S3EncryptionConfig`

NewS3EncryptionConfig instantiates a new S3EncryptionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3EncryptionConfigWithDefaults

`func NewS3EncryptionConfigWithDefaults() *S3EncryptionConfig`

NewS3EncryptionConfigWithDefaults instantiates a new S3EncryptionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSseAlgorithm

`func (o *S3EncryptionConfig) GetSseAlgorithm() string`

GetSseAlgorithm returns the SseAlgorithm field if non-nil, zero value otherwise.

### GetSseAlgorithmOk

`func (o *S3EncryptionConfig) GetSseAlgorithmOk() (*string, bool)`

GetSseAlgorithmOk returns a tuple with the SseAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSseAlgorithm

`func (o *S3EncryptionConfig) SetSseAlgorithm(v string)`

SetSseAlgorithm sets SseAlgorithm field to given value.

### HasSseAlgorithm

`func (o *S3EncryptionConfig) HasSseAlgorithm() bool`

HasSseAlgorithm returns a boolean if a field has been set.

### GetKmsMasterKeyId

`func (o *S3EncryptionConfig) GetKmsMasterKeyId() string`

GetKmsMasterKeyId returns the KmsMasterKeyId field if non-nil, zero value otherwise.

### GetKmsMasterKeyIdOk

`func (o *S3EncryptionConfig) GetKmsMasterKeyIdOk() (*string, bool)`

GetKmsMasterKeyIdOk returns a tuple with the KmsMasterKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsMasterKeyId

`func (o *S3EncryptionConfig) SetKmsMasterKeyId(v string)`

SetKmsMasterKeyId sets KmsMasterKeyId field to given value.

### HasKmsMasterKeyId

`func (o *S3EncryptionConfig) HasKmsMasterKeyId() bool`

HasKmsMasterKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



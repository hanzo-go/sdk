# S3CreateServiceAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to [**S3BucketPolicy**](S3BucketPolicy.md) |  | [optional] 

## Methods

### NewS3CreateServiceAccountRequest

`func NewS3CreateServiceAccountRequest() *S3CreateServiceAccountRequest`

NewS3CreateServiceAccountRequest instantiates a new S3CreateServiceAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3CreateServiceAccountRequestWithDefaults

`func NewS3CreateServiceAccountRequestWithDefaults() *S3CreateServiceAccountRequest`

NewS3CreateServiceAccountRequestWithDefaults instantiates a new S3CreateServiceAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *S3CreateServiceAccountRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *S3CreateServiceAccountRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *S3CreateServiceAccountRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *S3CreateServiceAccountRequest) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *S3CreateServiceAccountRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *S3CreateServiceAccountRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *S3CreateServiceAccountRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *S3CreateServiceAccountRequest) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetPolicy

`func (o *S3CreateServiceAccountRequest) GetPolicy() S3BucketPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *S3CreateServiceAccountRequest) GetPolicyOk() (*S3BucketPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *S3CreateServiceAccountRequest) SetPolicy(v S3BucketPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *S3CreateServiceAccountRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



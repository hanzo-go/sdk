# S3PresignResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresIn** | Pointer to **int64** | seconds until the URL expires | [optional] 
**Key** | Pointer to **string** | Key is the object key the URL was signed for, relative to the bucket root and path-cleaned — so it is what the store will actually read or write, which is not always the string the caller sent. The signature covers this one bucket and this one key: a URL minted here reaches nothing else. | [optional] 
**Method** | Pointer to **string** | \&quot;PUT\&quot; (upload) or \&quot;GET\&quot; (download) | [optional] 
**Url** | Pointer to **string** | presigned URL the browser follows directly | [optional] 

## Methods

### NewS3PresignResponse

`func NewS3PresignResponse() *S3PresignResponse`

NewS3PresignResponse instantiates a new S3PresignResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3PresignResponseWithDefaults

`func NewS3PresignResponseWithDefaults() *S3PresignResponse`

NewS3PresignResponseWithDefaults instantiates a new S3PresignResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresIn

`func (o *S3PresignResponse) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *S3PresignResponse) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *S3PresignResponse) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *S3PresignResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetKey

`func (o *S3PresignResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *S3PresignResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *S3PresignResponse) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *S3PresignResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMethod

`func (o *S3PresignResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *S3PresignResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *S3PresignResponse) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *S3PresignResponse) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUrl

`func (o *S3PresignResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *S3PresignResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *S3PresignResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *S3PresignResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



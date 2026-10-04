# S3UploadGone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Upload** | Pointer to **string** | Upload is the id of the upload abandoned; its stored parts are deleted. | [optional] 

## Methods

### NewS3UploadGone

`func NewS3UploadGone() *S3UploadGone`

NewS3UploadGone instantiates a new S3UploadGone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3UploadGoneWithDefaults

`func NewS3UploadGoneWithDefaults() *S3UploadGone`

NewS3UploadGoneWithDefaults instantiates a new S3UploadGone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpload

`func (o *S3UploadGone) GetUpload() string`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *S3UploadGone) GetUploadOk() (*string, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *S3UploadGone) SetUpload(v string)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *S3UploadGone) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



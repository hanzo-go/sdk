# S3UploadStarted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key is the object key the upload is bound to. | [optional] 
**PartSize** | Pointer to **int64** | PartSize is the size every part but the last must be, in bytes. The number of parts is the file&#39;s size divided by it, rounded up. | [optional] 
**Upload** | Pointer to **string** | Upload is the upload&#39;s id, which every later call names. | [optional] 

## Methods

### NewS3UploadStarted

`func NewS3UploadStarted() *S3UploadStarted`

NewS3UploadStarted instantiates a new S3UploadStarted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3UploadStartedWithDefaults

`func NewS3UploadStartedWithDefaults() *S3UploadStarted`

NewS3UploadStartedWithDefaults instantiates a new S3UploadStarted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *S3UploadStarted) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *S3UploadStarted) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *S3UploadStarted) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *S3UploadStarted) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPartSize

`func (o *S3UploadStarted) GetPartSize() int64`

GetPartSize returns the PartSize field if non-nil, zero value otherwise.

### GetPartSizeOk

`func (o *S3UploadStarted) GetPartSizeOk() (*int64, bool)`

GetPartSizeOk returns a tuple with the PartSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartSize

`func (o *S3UploadStarted) SetPartSize(v int64)`

SetPartSize sets PartSize field to given value.

### HasPartSize

`func (o *S3UploadStarted) HasPartSize() bool`

HasPartSize returns a boolean if a field has been set.

### GetUpload

`func (o *S3UploadStarted) GetUpload() string`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *S3UploadStarted) GetUploadOk() (*string, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *S3UploadStarted) SetUpload(v string)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *S3UploadStarted) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



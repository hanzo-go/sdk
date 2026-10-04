# S3StoredParts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key is the object key it is bound to. | [optional] 
**Parts** | Pointer to [**[]S3StoredPart**](S3StoredPart.md) | Parts are the parts stored, in order. A client resuming sends the others. | [optional] 
**Upload** | Pointer to **string** | Upload is the upload&#39;s id. | [optional] 

## Methods

### NewS3StoredParts

`func NewS3StoredParts() *S3StoredParts`

NewS3StoredParts instantiates a new S3StoredParts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3StoredPartsWithDefaults

`func NewS3StoredPartsWithDefaults() *S3StoredParts`

NewS3StoredPartsWithDefaults instantiates a new S3StoredParts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *S3StoredParts) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *S3StoredParts) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *S3StoredParts) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *S3StoredParts) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetParts

`func (o *S3StoredParts) GetParts() []S3StoredPart`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *S3StoredParts) GetPartsOk() (*[]S3StoredPart, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *S3StoredParts) SetParts(v []S3StoredPart)`

SetParts sets Parts field to given value.

### HasParts

`func (o *S3StoredParts) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetUpload

`func (o *S3StoredParts) GetUpload() string`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *S3StoredParts) GetUploadOk() (*string, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *S3StoredParts) SetUpload(v string)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *S3StoredParts) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



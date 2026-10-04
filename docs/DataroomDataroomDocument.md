# DataroomDataroomDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | ContentType is the mime type recorded at upload, null when none was sent. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the document was uploaded, in unix milliseconds. | [optional] 
**FileKey** | Pointer to **string** | FileKey is the opaque object-storage key the bytes are stored under. It is scoped to the tenant&#39;s own key prefix and is not a URL. | [optional] 
**FileSize** | Pointer to **int64** | FileSize is the stored byte count, null when it was not recorded. | [optional] 
**Id** | Pointer to **string** | ID is the document id. | [optional] 
**Name** | Pointer to **string** | Name is the document&#39;s display name. | [optional] 
**NumPages** | Pointer to **int64** | NumPages is the page count, null when it was not supplied at upload. | [optional] 
**Type** | Pointer to **string** | Type is the document&#39;s kind, null when it was not recorded. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when the document row last changed, in unix milliseconds. | [optional] 

## Methods

### NewDataroomDataroomDocument

`func NewDataroomDataroomDocument() *DataroomDataroomDocument`

NewDataroomDataroomDocument instantiates a new DataroomDataroomDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomDataroomDocumentWithDefaults

`func NewDataroomDataroomDocumentWithDefaults() *DataroomDataroomDocument`

NewDataroomDataroomDocumentWithDefaults instantiates a new DataroomDataroomDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *DataroomDataroomDocument) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *DataroomDataroomDocument) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *DataroomDataroomDocument) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *DataroomDataroomDocument) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DataroomDataroomDocument) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataroomDataroomDocument) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataroomDataroomDocument) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DataroomDataroomDocument) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFileKey

`func (o *DataroomDataroomDocument) GetFileKey() string`

GetFileKey returns the FileKey field if non-nil, zero value otherwise.

### GetFileKeyOk

`func (o *DataroomDataroomDocument) GetFileKeyOk() (*string, bool)`

GetFileKeyOk returns a tuple with the FileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKey

`func (o *DataroomDataroomDocument) SetFileKey(v string)`

SetFileKey sets FileKey field to given value.

### HasFileKey

`func (o *DataroomDataroomDocument) HasFileKey() bool`

HasFileKey returns a boolean if a field has been set.

### GetFileSize

`func (o *DataroomDataroomDocument) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *DataroomDataroomDocument) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *DataroomDataroomDocument) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *DataroomDataroomDocument) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetId

`func (o *DataroomDataroomDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataroomDataroomDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataroomDataroomDocument) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataroomDataroomDocument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DataroomDataroomDocument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataroomDataroomDocument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataroomDataroomDocument) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataroomDataroomDocument) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumPages

`func (o *DataroomDataroomDocument) GetNumPages() int64`

GetNumPages returns the NumPages field if non-nil, zero value otherwise.

### GetNumPagesOk

`func (o *DataroomDataroomDocument) GetNumPagesOk() (*int64, bool)`

GetNumPagesOk returns a tuple with the NumPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPages

`func (o *DataroomDataroomDocument) SetNumPages(v int64)`

SetNumPages sets NumPages field to given value.

### HasNumPages

`func (o *DataroomDataroomDocument) HasNumPages() bool`

HasNumPages returns a boolean if a field has been set.

### GetType

`func (o *DataroomDataroomDocument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataroomDataroomDocument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataroomDataroomDocument) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DataroomDataroomDocument) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DataroomDataroomDocument) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DataroomDataroomDocument) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DataroomDataroomDocument) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DataroomDataroomDocument) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



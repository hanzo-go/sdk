# DataroomDataroomMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | ContentType is the mime type recorded at upload, null when none was sent. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the document was uploaded, in unix milliseconds. | [optional] 
**DataroomDocumentId** | Pointer to **string** | DataroomDocumentId is the membership id — this document&#39;s place in THIS room, distinct from the document id. | [optional] 
**FileKey** | Pointer to **string** | FileKey is the opaque object-storage key the bytes are stored under. | [optional] 
**FileSize** | Pointer to **int64** | FileSize is the stored byte count, null when it was not recorded. | [optional] 
**Id** | Pointer to **string** | ID is the document id. | [optional] 
**Name** | Pointer to **string** | Name is the document&#39;s display name. | [optional] 
**NumPages** | Pointer to **int64** | NumPages is the page count, null when it was not supplied at upload. | [optional] 
**OrderIndex** | Pointer to **int64** | OrderIndex is the document&#39;s place in the viewer&#39;s list, null when it was added without one. Unordered documents sort last. | [optional] 
**Type** | Pointer to **string** | Type is the document&#39;s kind, null when it was not recorded. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when the document row last changed, in unix milliseconds. | [optional] 

## Methods

### NewDataroomDataroomMember

`func NewDataroomDataroomMember() *DataroomDataroomMember`

NewDataroomDataroomMember instantiates a new DataroomDataroomMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomDataroomMemberWithDefaults

`func NewDataroomDataroomMemberWithDefaults() *DataroomDataroomMember`

NewDataroomDataroomMemberWithDefaults instantiates a new DataroomDataroomMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *DataroomDataroomMember) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *DataroomDataroomMember) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *DataroomDataroomMember) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *DataroomDataroomMember) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DataroomDataroomMember) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataroomDataroomMember) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataroomDataroomMember) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DataroomDataroomMember) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDataroomDocumentId

`func (o *DataroomDataroomMember) GetDataroomDocumentId() string`

GetDataroomDocumentId returns the DataroomDocumentId field if non-nil, zero value otherwise.

### GetDataroomDocumentIdOk

`func (o *DataroomDataroomMember) GetDataroomDocumentIdOk() (*string, bool)`

GetDataroomDocumentIdOk returns a tuple with the DataroomDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataroomDocumentId

`func (o *DataroomDataroomMember) SetDataroomDocumentId(v string)`

SetDataroomDocumentId sets DataroomDocumentId field to given value.

### HasDataroomDocumentId

`func (o *DataroomDataroomMember) HasDataroomDocumentId() bool`

HasDataroomDocumentId returns a boolean if a field has been set.

### GetFileKey

`func (o *DataroomDataroomMember) GetFileKey() string`

GetFileKey returns the FileKey field if non-nil, zero value otherwise.

### GetFileKeyOk

`func (o *DataroomDataroomMember) GetFileKeyOk() (*string, bool)`

GetFileKeyOk returns a tuple with the FileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKey

`func (o *DataroomDataroomMember) SetFileKey(v string)`

SetFileKey sets FileKey field to given value.

### HasFileKey

`func (o *DataroomDataroomMember) HasFileKey() bool`

HasFileKey returns a boolean if a field has been set.

### GetFileSize

`func (o *DataroomDataroomMember) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *DataroomDataroomMember) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *DataroomDataroomMember) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *DataroomDataroomMember) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetId

`func (o *DataroomDataroomMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataroomDataroomMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataroomDataroomMember) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataroomDataroomMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DataroomDataroomMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataroomDataroomMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataroomDataroomMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataroomDataroomMember) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumPages

`func (o *DataroomDataroomMember) GetNumPages() int64`

GetNumPages returns the NumPages field if non-nil, zero value otherwise.

### GetNumPagesOk

`func (o *DataroomDataroomMember) GetNumPagesOk() (*int64, bool)`

GetNumPagesOk returns a tuple with the NumPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPages

`func (o *DataroomDataroomMember) SetNumPages(v int64)`

SetNumPages sets NumPages field to given value.

### HasNumPages

`func (o *DataroomDataroomMember) HasNumPages() bool`

HasNumPages returns a boolean if a field has been set.

### GetOrderIndex

`func (o *DataroomDataroomMember) GetOrderIndex() int64`

GetOrderIndex returns the OrderIndex field if non-nil, zero value otherwise.

### GetOrderIndexOk

`func (o *DataroomDataroomMember) GetOrderIndexOk() (*int64, bool)`

GetOrderIndexOk returns a tuple with the OrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderIndex

`func (o *DataroomDataroomMember) SetOrderIndex(v int64)`

SetOrderIndex sets OrderIndex field to given value.

### HasOrderIndex

`func (o *DataroomDataroomMember) HasOrderIndex() bool`

HasOrderIndex returns a boolean if a field has been set.

### GetType

`func (o *DataroomDataroomMember) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataroomDataroomMember) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataroomDataroomMember) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DataroomDataroomMember) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DataroomDataroomMember) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DataroomDataroomMember) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DataroomDataroomMember) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DataroomDataroomMember) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



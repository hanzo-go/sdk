# DataroomMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | ContentType is the mime type recorded at upload, null when none was sent. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the document was uploaded, in unix milliseconds. | [optional] 
**DataroomDocumentId** | Pointer to **string** | DataroomDocumentId is the membership id — this document&#39;s place in THIS room, distinct from the document id. | [optional] 
**FileKey** | Pointer to **string** | FileKey is the opaque object-storage key the bytes are stored under. | [optional] 
**FileSize** | Pointer to **int32** | FileSize is the stored byte count, null when it was not recorded. | [optional] 
**Id** | Pointer to **string** | ID is the document id. | [optional] 
**Name** | Pointer to **string** | Name is the document&#39;s display name. | [optional] 
**NumPages** | Pointer to **int32** | NumPages is the page count, null when it was not supplied at upload. | [optional] 
**OrderIndex** | Pointer to **int32** | OrderIndex is the document&#39;s place in the viewer&#39;s list, null when it was added without one. Unordered documents sort last. | [optional] 
**Type** | Pointer to **string** | Type is the document&#39;s kind, null when it was not recorded. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is when the document row last changed, in unix milliseconds. | [optional] 

## Methods

### NewDataroomMember

`func NewDataroomMember() *DataroomMember`

NewDataroomMember instantiates a new DataroomMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomMemberWithDefaults

`func NewDataroomMemberWithDefaults() *DataroomMember`

NewDataroomMemberWithDefaults instantiates a new DataroomMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *DataroomMember) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *DataroomMember) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *DataroomMember) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *DataroomMember) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DataroomMember) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataroomMember) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataroomMember) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DataroomMember) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDataroomDocumentId

`func (o *DataroomMember) GetDataroomDocumentId() string`

GetDataroomDocumentId returns the DataroomDocumentId field if non-nil, zero value otherwise.

### GetDataroomDocumentIdOk

`func (o *DataroomMember) GetDataroomDocumentIdOk() (*string, bool)`

GetDataroomDocumentIdOk returns a tuple with the DataroomDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataroomDocumentId

`func (o *DataroomMember) SetDataroomDocumentId(v string)`

SetDataroomDocumentId sets DataroomDocumentId field to given value.

### HasDataroomDocumentId

`func (o *DataroomMember) HasDataroomDocumentId() bool`

HasDataroomDocumentId returns a boolean if a field has been set.

### GetFileKey

`func (o *DataroomMember) GetFileKey() string`

GetFileKey returns the FileKey field if non-nil, zero value otherwise.

### GetFileKeyOk

`func (o *DataroomMember) GetFileKeyOk() (*string, bool)`

GetFileKeyOk returns a tuple with the FileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKey

`func (o *DataroomMember) SetFileKey(v string)`

SetFileKey sets FileKey field to given value.

### HasFileKey

`func (o *DataroomMember) HasFileKey() bool`

HasFileKey returns a boolean if a field has been set.

### GetFileSize

`func (o *DataroomMember) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *DataroomMember) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *DataroomMember) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *DataroomMember) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetId

`func (o *DataroomMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataroomMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataroomMember) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataroomMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DataroomMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataroomMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataroomMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataroomMember) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumPages

`func (o *DataroomMember) GetNumPages() int32`

GetNumPages returns the NumPages field if non-nil, zero value otherwise.

### GetNumPagesOk

`func (o *DataroomMember) GetNumPagesOk() (*int32, bool)`

GetNumPagesOk returns a tuple with the NumPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPages

`func (o *DataroomMember) SetNumPages(v int32)`

SetNumPages sets NumPages field to given value.

### HasNumPages

`func (o *DataroomMember) HasNumPages() bool`

HasNumPages returns a boolean if a field has been set.

### GetOrderIndex

`func (o *DataroomMember) GetOrderIndex() int32`

GetOrderIndex returns the OrderIndex field if non-nil, zero value otherwise.

### GetOrderIndexOk

`func (o *DataroomMember) GetOrderIndexOk() (*int32, bool)`

GetOrderIndexOk returns a tuple with the OrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderIndex

`func (o *DataroomMember) SetOrderIndex(v int32)`

SetOrderIndex sets OrderIndex field to given value.

### HasOrderIndex

`func (o *DataroomMember) HasOrderIndex() bool`

HasOrderIndex returns a boolean if a field has been set.

### GetType

`func (o *DataroomMember) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataroomMember) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataroomMember) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DataroomMember) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DataroomMember) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DataroomMember) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DataroomMember) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DataroomMember) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



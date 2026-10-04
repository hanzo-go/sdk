# DataroomDataroomLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDownload** | Pointer to **bool** | AllowDownload is whether a visitor may download, rather than only view. | [optional] 
**AllowList** | Pointer to **[]string** | AllowList narrows which addresses pass the email gate. An entry may be a full address, an \&quot;@domain.com\&quot; suffix, or a bare \&quot;domain.com\&quot;. An EMPTY list admits everyone. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the link was minted, in unix milliseconds. | [optional] 
**DataroomId** | Pointer to **string** | DataroomId is the room the link opens, null for a single-document link. | [optional] 
**DenyList** | Pointer to **[]string** | DenyList rejects addresses, in the same three forms as the allow list, and is checked BEFORE it — so deny always wins. | [optional] 
**DocumentId** | Pointer to **string** | DocumentId is the document the link opens, null for a room link. | [optional] 
**EmailProtected** | Pointer to **bool** | EmailProtected is whether a visitor must state an address to enter. | [optional] 
**ExpiresAt** | Pointer to **int64** | ExpiresAt is when the link closes, in unix milliseconds; null never expires. | [optional] 
**HasPassword** | Pointer to **bool** | HasPassword reports THAT a password is set. The stored form is a bcrypt hash and no route returns it. | [optional] 
**Id** | Pointer to **string** | ID is the link id — the public token a visitor opens the room with. | [optional] 
**IsArchived** | Pointer to **bool** | IsArchived is whether the link has been retired. | [optional] 
**LinkType** | Pointer to **string** | LinkType is DATAROOM_LINK or DOCUMENT_LINK. | [optional] 
**Name** | Pointer to **string** | Name is the link&#39;s label, null when none was given. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when the link last changed, in unix milliseconds. | [optional] 

## Methods

### NewDataroomDataroomLink

`func NewDataroomDataroomLink() *DataroomDataroomLink`

NewDataroomDataroomLink instantiates a new DataroomDataroomLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomDataroomLinkWithDefaults

`func NewDataroomDataroomLinkWithDefaults() *DataroomDataroomLink`

NewDataroomDataroomLinkWithDefaults instantiates a new DataroomDataroomLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowDownload

`func (o *DataroomDataroomLink) GetAllowDownload() bool`

GetAllowDownload returns the AllowDownload field if non-nil, zero value otherwise.

### GetAllowDownloadOk

`func (o *DataroomDataroomLink) GetAllowDownloadOk() (*bool, bool)`

GetAllowDownloadOk returns a tuple with the AllowDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDownload

`func (o *DataroomDataroomLink) SetAllowDownload(v bool)`

SetAllowDownload sets AllowDownload field to given value.

### HasAllowDownload

`func (o *DataroomDataroomLink) HasAllowDownload() bool`

HasAllowDownload returns a boolean if a field has been set.

### GetAllowList

`func (o *DataroomDataroomLink) GetAllowList() []string`

GetAllowList returns the AllowList field if non-nil, zero value otherwise.

### GetAllowListOk

`func (o *DataroomDataroomLink) GetAllowListOk() (*[]string, bool)`

GetAllowListOk returns a tuple with the AllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowList

`func (o *DataroomDataroomLink) SetAllowList(v []string)`

SetAllowList sets AllowList field to given value.

### HasAllowList

`func (o *DataroomDataroomLink) HasAllowList() bool`

HasAllowList returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DataroomDataroomLink) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataroomDataroomLink) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataroomDataroomLink) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DataroomDataroomLink) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDataroomId

`func (o *DataroomDataroomLink) GetDataroomId() string`

GetDataroomId returns the DataroomId field if non-nil, zero value otherwise.

### GetDataroomIdOk

`func (o *DataroomDataroomLink) GetDataroomIdOk() (*string, bool)`

GetDataroomIdOk returns a tuple with the DataroomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataroomId

`func (o *DataroomDataroomLink) SetDataroomId(v string)`

SetDataroomId sets DataroomId field to given value.

### HasDataroomId

`func (o *DataroomDataroomLink) HasDataroomId() bool`

HasDataroomId returns a boolean if a field has been set.

### GetDenyList

`func (o *DataroomDataroomLink) GetDenyList() []string`

GetDenyList returns the DenyList field if non-nil, zero value otherwise.

### GetDenyListOk

`func (o *DataroomDataroomLink) GetDenyListOk() (*[]string, bool)`

GetDenyListOk returns a tuple with the DenyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyList

`func (o *DataroomDataroomLink) SetDenyList(v []string)`

SetDenyList sets DenyList field to given value.

### HasDenyList

`func (o *DataroomDataroomLink) HasDenyList() bool`

HasDenyList returns a boolean if a field has been set.

### GetDocumentId

`func (o *DataroomDataroomLink) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *DataroomDataroomLink) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *DataroomDataroomLink) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *DataroomDataroomLink) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetEmailProtected

`func (o *DataroomDataroomLink) GetEmailProtected() bool`

GetEmailProtected returns the EmailProtected field if non-nil, zero value otherwise.

### GetEmailProtectedOk

`func (o *DataroomDataroomLink) GetEmailProtectedOk() (*bool, bool)`

GetEmailProtectedOk returns a tuple with the EmailProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailProtected

`func (o *DataroomDataroomLink) SetEmailProtected(v bool)`

SetEmailProtected sets EmailProtected field to given value.

### HasEmailProtected

`func (o *DataroomDataroomLink) HasEmailProtected() bool`

HasEmailProtected returns a boolean if a field has been set.

### GetExpiresAt

`func (o *DataroomDataroomLink) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DataroomDataroomLink) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DataroomDataroomLink) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DataroomDataroomLink) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetHasPassword

`func (o *DataroomDataroomLink) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *DataroomDataroomLink) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *DataroomDataroomLink) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *DataroomDataroomLink) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetId

`func (o *DataroomDataroomLink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataroomDataroomLink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataroomDataroomLink) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataroomDataroomLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsArchived

`func (o *DataroomDataroomLink) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *DataroomDataroomLink) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *DataroomDataroomLink) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *DataroomDataroomLink) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetLinkType

`func (o *DataroomDataroomLink) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *DataroomDataroomLink) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *DataroomDataroomLink) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.

### HasLinkType

`func (o *DataroomDataroomLink) HasLinkType() bool`

HasLinkType returns a boolean if a field has been set.

### GetName

`func (o *DataroomDataroomLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataroomDataroomLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataroomDataroomLink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataroomDataroomLink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DataroomDataroomLink) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DataroomDataroomLink) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DataroomDataroomLink) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DataroomDataroomLink) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



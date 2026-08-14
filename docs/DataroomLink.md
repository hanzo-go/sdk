# DataroomLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDownload** | Pointer to **bool** | AllowDownload is whether a visitor may download, rather than only view. | [optional] 
**AllowList** | Pointer to **[]string** | AllowList narrows which addresses pass the email gate. An entry may be a full address, an \&quot;@domain.com\&quot; suffix, or a bare \&quot;domain.com\&quot;. An EMPTY list admits everyone. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the link was minted, in unix milliseconds. | [optional] 
**DataroomId** | Pointer to **string** | DataroomId is the room the link opens, null for a single-document link. | [optional] 
**DenyList** | Pointer to **[]string** | DenyList rejects addresses, in the same three forms as the allow list, and is checked BEFORE it — so deny always wins. | [optional] 
**DocumentId** | Pointer to **string** | DocumentId is the document the link opens, null for a room link. | [optional] 
**EmailProtected** | Pointer to **bool** | EmailProtected is whether a visitor must state an address to enter. | [optional] 
**ExpiresAt** | Pointer to **int32** | ExpiresAt is when the link closes, in unix milliseconds; null never expires. | [optional] 
**HasPassword** | Pointer to **bool** | HasPassword reports THAT a password is set. The stored form is a bcrypt hash and no route returns it. | [optional] 
**Id** | Pointer to **string** | ID is the link id — the public token a visitor opens the room with. | [optional] 
**IsArchived** | Pointer to **bool** | IsArchived is whether the link has been retired. | [optional] 
**LinkType** | Pointer to **string** | LinkType is DATAROOM_LINK or DOCUMENT_LINK. | [optional] 
**Name** | Pointer to **string** | Name is the link&#39;s label, null when none was given. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is when the link last changed, in unix milliseconds. | [optional] 

## Methods

### NewDataroomLink

`func NewDataroomLink() *DataroomLink`

NewDataroomLink instantiates a new DataroomLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomLinkWithDefaults

`func NewDataroomLinkWithDefaults() *DataroomLink`

NewDataroomLinkWithDefaults instantiates a new DataroomLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowDownload

`func (o *DataroomLink) GetAllowDownload() bool`

GetAllowDownload returns the AllowDownload field if non-nil, zero value otherwise.

### GetAllowDownloadOk

`func (o *DataroomLink) GetAllowDownloadOk() (*bool, bool)`

GetAllowDownloadOk returns a tuple with the AllowDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDownload

`func (o *DataroomLink) SetAllowDownload(v bool)`

SetAllowDownload sets AllowDownload field to given value.

### HasAllowDownload

`func (o *DataroomLink) HasAllowDownload() bool`

HasAllowDownload returns a boolean if a field has been set.

### GetAllowList

`func (o *DataroomLink) GetAllowList() []string`

GetAllowList returns the AllowList field if non-nil, zero value otherwise.

### GetAllowListOk

`func (o *DataroomLink) GetAllowListOk() (*[]string, bool)`

GetAllowListOk returns a tuple with the AllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowList

`func (o *DataroomLink) SetAllowList(v []string)`

SetAllowList sets AllowList field to given value.

### HasAllowList

`func (o *DataroomLink) HasAllowList() bool`

HasAllowList returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DataroomLink) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataroomLink) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataroomLink) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DataroomLink) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDataroomId

`func (o *DataroomLink) GetDataroomId() string`

GetDataroomId returns the DataroomId field if non-nil, zero value otherwise.

### GetDataroomIdOk

`func (o *DataroomLink) GetDataroomIdOk() (*string, bool)`

GetDataroomIdOk returns a tuple with the DataroomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataroomId

`func (o *DataroomLink) SetDataroomId(v string)`

SetDataroomId sets DataroomId field to given value.

### HasDataroomId

`func (o *DataroomLink) HasDataroomId() bool`

HasDataroomId returns a boolean if a field has been set.

### GetDenyList

`func (o *DataroomLink) GetDenyList() []string`

GetDenyList returns the DenyList field if non-nil, zero value otherwise.

### GetDenyListOk

`func (o *DataroomLink) GetDenyListOk() (*[]string, bool)`

GetDenyListOk returns a tuple with the DenyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyList

`func (o *DataroomLink) SetDenyList(v []string)`

SetDenyList sets DenyList field to given value.

### HasDenyList

`func (o *DataroomLink) HasDenyList() bool`

HasDenyList returns a boolean if a field has been set.

### GetDocumentId

`func (o *DataroomLink) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *DataroomLink) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *DataroomLink) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *DataroomLink) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetEmailProtected

`func (o *DataroomLink) GetEmailProtected() bool`

GetEmailProtected returns the EmailProtected field if non-nil, zero value otherwise.

### GetEmailProtectedOk

`func (o *DataroomLink) GetEmailProtectedOk() (*bool, bool)`

GetEmailProtectedOk returns a tuple with the EmailProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailProtected

`func (o *DataroomLink) SetEmailProtected(v bool)`

SetEmailProtected sets EmailProtected field to given value.

### HasEmailProtected

`func (o *DataroomLink) HasEmailProtected() bool`

HasEmailProtected returns a boolean if a field has been set.

### GetExpiresAt

`func (o *DataroomLink) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DataroomLink) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DataroomLink) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DataroomLink) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetHasPassword

`func (o *DataroomLink) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *DataroomLink) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *DataroomLink) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *DataroomLink) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetId

`func (o *DataroomLink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataroomLink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataroomLink) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataroomLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsArchived

`func (o *DataroomLink) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *DataroomLink) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *DataroomLink) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *DataroomLink) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetLinkType

`func (o *DataroomLink) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *DataroomLink) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *DataroomLink) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.

### HasLinkType

`func (o *DataroomLink) HasLinkType() bool`

HasLinkType returns a boolean if a field has been set.

### GetName

`func (o *DataroomLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataroomLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataroomLink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataroomLink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DataroomLink) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DataroomLink) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DataroomLink) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DataroomLink) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



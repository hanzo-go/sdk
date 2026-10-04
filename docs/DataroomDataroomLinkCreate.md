# DataroomDataroomLinkCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDownload** | Pointer to **interface{}** |  | [optional] 
**AllowList** | Pointer to **[]interface{}** | AllowList narrows which addresses pass the email gate. Optional; an entry may be a full address (\&quot;ada@example.com\&quot;), an \&quot;@domain.com\&quot; suffix, or a bare \&quot;domain.com\&quot;. An omitted or EMPTY list admits everyone, so a link with no list enforces the email gate alone. | [optional] 
**DataroomId** | Pointer to **interface{}** |  | [optional] 
**DenyList** | Pointer to **[]interface{}** | DenyList rejects addresses, in the same three forms as the allow list. Optional. It is checked BEFORE the allow list, so deny always wins. | [optional] 
**DocumentId** | Pointer to **interface{}** |  | [optional] 
**EmailProtected** | Pointer to **interface{}** |  | [optional] 
**ExpiresAt** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **interface{}** |  | [optional] 
**Password** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewDataroomDataroomLinkCreate

`func NewDataroomDataroomLinkCreate() *DataroomDataroomLinkCreate`

NewDataroomDataroomLinkCreate instantiates a new DataroomDataroomLinkCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomDataroomLinkCreateWithDefaults

`func NewDataroomDataroomLinkCreateWithDefaults() *DataroomDataroomLinkCreate`

NewDataroomDataroomLinkCreateWithDefaults instantiates a new DataroomDataroomLinkCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowDownload

`func (o *DataroomDataroomLinkCreate) GetAllowDownload() interface{}`

GetAllowDownload returns the AllowDownload field if non-nil, zero value otherwise.

### GetAllowDownloadOk

`func (o *DataroomDataroomLinkCreate) GetAllowDownloadOk() (*interface{}, bool)`

GetAllowDownloadOk returns a tuple with the AllowDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDownload

`func (o *DataroomDataroomLinkCreate) SetAllowDownload(v interface{})`

SetAllowDownload sets AllowDownload field to given value.

### HasAllowDownload

`func (o *DataroomDataroomLinkCreate) HasAllowDownload() bool`

HasAllowDownload returns a boolean if a field has been set.

### SetAllowDownloadNil

`func (o *DataroomDataroomLinkCreate) SetAllowDownloadNil(b bool)`

 SetAllowDownloadNil sets the value for AllowDownload to be an explicit nil

### UnsetAllowDownload
`func (o *DataroomDataroomLinkCreate) UnsetAllowDownload()`

UnsetAllowDownload ensures that no value is present for AllowDownload, not even an explicit nil
### GetAllowList

`func (o *DataroomDataroomLinkCreate) GetAllowList() []interface{}`

GetAllowList returns the AllowList field if non-nil, zero value otherwise.

### GetAllowListOk

`func (o *DataroomDataroomLinkCreate) GetAllowListOk() (*[]interface{}, bool)`

GetAllowListOk returns a tuple with the AllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowList

`func (o *DataroomDataroomLinkCreate) SetAllowList(v []interface{})`

SetAllowList sets AllowList field to given value.

### HasAllowList

`func (o *DataroomDataroomLinkCreate) HasAllowList() bool`

HasAllowList returns a boolean if a field has been set.

### GetDataroomId

`func (o *DataroomDataroomLinkCreate) GetDataroomId() interface{}`

GetDataroomId returns the DataroomId field if non-nil, zero value otherwise.

### GetDataroomIdOk

`func (o *DataroomDataroomLinkCreate) GetDataroomIdOk() (*interface{}, bool)`

GetDataroomIdOk returns a tuple with the DataroomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataroomId

`func (o *DataroomDataroomLinkCreate) SetDataroomId(v interface{})`

SetDataroomId sets DataroomId field to given value.

### HasDataroomId

`func (o *DataroomDataroomLinkCreate) HasDataroomId() bool`

HasDataroomId returns a boolean if a field has been set.

### SetDataroomIdNil

`func (o *DataroomDataroomLinkCreate) SetDataroomIdNil(b bool)`

 SetDataroomIdNil sets the value for DataroomId to be an explicit nil

### UnsetDataroomId
`func (o *DataroomDataroomLinkCreate) UnsetDataroomId()`

UnsetDataroomId ensures that no value is present for DataroomId, not even an explicit nil
### GetDenyList

`func (o *DataroomDataroomLinkCreate) GetDenyList() []interface{}`

GetDenyList returns the DenyList field if non-nil, zero value otherwise.

### GetDenyListOk

`func (o *DataroomDataroomLinkCreate) GetDenyListOk() (*[]interface{}, bool)`

GetDenyListOk returns a tuple with the DenyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyList

`func (o *DataroomDataroomLinkCreate) SetDenyList(v []interface{})`

SetDenyList sets DenyList field to given value.

### HasDenyList

`func (o *DataroomDataroomLinkCreate) HasDenyList() bool`

HasDenyList returns a boolean if a field has been set.

### GetDocumentId

`func (o *DataroomDataroomLinkCreate) GetDocumentId() interface{}`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *DataroomDataroomLinkCreate) GetDocumentIdOk() (*interface{}, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *DataroomDataroomLinkCreate) SetDocumentId(v interface{})`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *DataroomDataroomLinkCreate) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### SetDocumentIdNil

`func (o *DataroomDataroomLinkCreate) SetDocumentIdNil(b bool)`

 SetDocumentIdNil sets the value for DocumentId to be an explicit nil

### UnsetDocumentId
`func (o *DataroomDataroomLinkCreate) UnsetDocumentId()`

UnsetDocumentId ensures that no value is present for DocumentId, not even an explicit nil
### GetEmailProtected

`func (o *DataroomDataroomLinkCreate) GetEmailProtected() interface{}`

GetEmailProtected returns the EmailProtected field if non-nil, zero value otherwise.

### GetEmailProtectedOk

`func (o *DataroomDataroomLinkCreate) GetEmailProtectedOk() (*interface{}, bool)`

GetEmailProtectedOk returns a tuple with the EmailProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailProtected

`func (o *DataroomDataroomLinkCreate) SetEmailProtected(v interface{})`

SetEmailProtected sets EmailProtected field to given value.

### HasEmailProtected

`func (o *DataroomDataroomLinkCreate) HasEmailProtected() bool`

HasEmailProtected returns a boolean if a field has been set.

### SetEmailProtectedNil

`func (o *DataroomDataroomLinkCreate) SetEmailProtectedNil(b bool)`

 SetEmailProtectedNil sets the value for EmailProtected to be an explicit nil

### UnsetEmailProtected
`func (o *DataroomDataroomLinkCreate) UnsetEmailProtected()`

UnsetEmailProtected ensures that no value is present for EmailProtected, not even an explicit nil
### GetExpiresAt

`func (o *DataroomDataroomLinkCreate) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DataroomDataroomLinkCreate) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DataroomDataroomLinkCreate) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DataroomDataroomLinkCreate) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *DataroomDataroomLinkCreate) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *DataroomDataroomLinkCreate) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetName

`func (o *DataroomDataroomLinkCreate) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataroomDataroomLinkCreate) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataroomDataroomLinkCreate) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *DataroomDataroomLinkCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DataroomDataroomLinkCreate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DataroomDataroomLinkCreate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPassword

`func (o *DataroomDataroomLinkCreate) GetPassword() interface{}`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DataroomDataroomLinkCreate) GetPasswordOk() (*interface{}, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DataroomDataroomLinkCreate) SetPassword(v interface{})`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DataroomDataroomLinkCreate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *DataroomDataroomLinkCreate) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *DataroomDataroomLinkCreate) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



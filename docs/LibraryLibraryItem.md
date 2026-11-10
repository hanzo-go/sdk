# LibraryLibraryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | Pointer to **string** | By is the owner as a reader names them: the name, else the address, the identity boundary carried when the item was made. | [optional] 
**Bytes** | Pointer to **int64** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Download** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Mime** | Pointer to **string** |  | [optional] 
**Mine** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Pages** | Pointer to **int64** |  | [optional] 
**Preview** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to **string** |  | [optional] 
**Seconds** | Pointer to **float64** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] 
**Upload** | Pointer to [**LibraryLibraryUpload**](LibraryLibraryUpload.md) | Upload is where a reserved item&#39;s bytes go, on the reservation response before it is committed. | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Width** | Pointer to **int64** |  | [optional] 

## Methods

### NewLibraryLibraryItem

`func NewLibraryLibraryItem() *LibraryLibraryItem`

NewLibraryLibraryItem instantiates a new LibraryLibraryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryLibraryItemWithDefaults

`func NewLibraryLibraryItemWithDefaults() *LibraryLibraryItem`

NewLibraryLibraryItemWithDefaults instantiates a new LibraryLibraryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *LibraryLibraryItem) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *LibraryLibraryItem) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *LibraryLibraryItem) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *LibraryLibraryItem) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetBytes

`func (o *LibraryLibraryItem) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *LibraryLibraryItem) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *LibraryLibraryItem) SetBytes(v int64)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *LibraryLibraryItem) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetCreated

`func (o *LibraryLibraryItem) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LibraryLibraryItem) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LibraryLibraryItem) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *LibraryLibraryItem) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDownload

`func (o *LibraryLibraryItem) GetDownload() string`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *LibraryLibraryItem) GetDownloadOk() (*string, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *LibraryLibraryItem) SetDownload(v string)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *LibraryLibraryItem) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetExpires

`func (o *LibraryLibraryItem) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *LibraryLibraryItem) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *LibraryLibraryItem) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *LibraryLibraryItem) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetHeight

`func (o *LibraryLibraryItem) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *LibraryLibraryItem) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *LibraryLibraryItem) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *LibraryLibraryItem) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *LibraryLibraryItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LibraryLibraryItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LibraryLibraryItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LibraryLibraryItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *LibraryLibraryItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LibraryLibraryItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LibraryLibraryItem) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *LibraryLibraryItem) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMime

`func (o *LibraryLibraryItem) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *LibraryLibraryItem) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *LibraryLibraryItem) SetMime(v string)`

SetMime sets Mime field to given value.

### HasMime

`func (o *LibraryLibraryItem) HasMime() bool`

HasMime returns a boolean if a field has been set.

### GetMine

`func (o *LibraryLibraryItem) GetMine() bool`

GetMine returns the Mine field if non-nil, zero value otherwise.

### GetMineOk

`func (o *LibraryLibraryItem) GetMineOk() (*bool, bool)`

GetMineOk returns a tuple with the Mine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMine

`func (o *LibraryLibraryItem) SetMine(v bool)`

SetMine sets Mine field to given value.

### HasMine

`func (o *LibraryLibraryItem) HasMine() bool`

HasMine returns a boolean if a field has been set.

### GetName

`func (o *LibraryLibraryItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LibraryLibraryItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LibraryLibraryItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LibraryLibraryItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrigin

`func (o *LibraryLibraryItem) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *LibraryLibraryItem) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *LibraryLibraryItem) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *LibraryLibraryItem) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetOwner

`func (o *LibraryLibraryItem) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *LibraryLibraryItem) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *LibraryLibraryItem) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *LibraryLibraryItem) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPages

`func (o *LibraryLibraryItem) GetPages() int64`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *LibraryLibraryItem) GetPagesOk() (*int64, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *LibraryLibraryItem) SetPages(v int64)`

SetPages sets Pages field to given value.

### HasPages

`func (o *LibraryLibraryItem) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetPreview

`func (o *LibraryLibraryItem) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *LibraryLibraryItem) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *LibraryLibraryItem) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *LibraryLibraryItem) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetProject

`func (o *LibraryLibraryItem) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *LibraryLibraryItem) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *LibraryLibraryItem) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *LibraryLibraryItem) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRef

`func (o *LibraryLibraryItem) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *LibraryLibraryItem) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *LibraryLibraryItem) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *LibraryLibraryItem) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSeconds

`func (o *LibraryLibraryItem) GetSeconds() float64`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *LibraryLibraryItem) GetSecondsOk() (*float64, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *LibraryLibraryItem) SetSeconds(v float64)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *LibraryLibraryItem) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.

### GetShared

`func (o *LibraryLibraryItem) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *LibraryLibraryItem) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *LibraryLibraryItem) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *LibraryLibraryItem) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSource

`func (o *LibraryLibraryItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LibraryLibraryItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LibraryLibraryItem) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LibraryLibraryItem) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUpdated

`func (o *LibraryLibraryItem) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *LibraryLibraryItem) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *LibraryLibraryItem) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *LibraryLibraryItem) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUpload

`func (o *LibraryLibraryItem) GetUpload() LibraryLibraryUpload`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *LibraryLibraryItem) GetUploadOk() (*LibraryLibraryUpload, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *LibraryLibraryItem) SetUpload(v LibraryLibraryUpload)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *LibraryLibraryItem) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetUrl

`func (o *LibraryLibraryItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LibraryLibraryItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LibraryLibraryItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LibraryLibraryItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWidth

`func (o *LibraryLibraryItem) GetWidth() int64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *LibraryLibraryItem) GetWidthOk() (*int64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *LibraryLibraryItem) SetWidth(v int64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *LibraryLibraryItem) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



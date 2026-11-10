# StudioItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Ref** | Pointer to **string** |  | [optional] 
**Seconds** | Pointer to **float64** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | URL opens the file inline; Download saves it; Preview is its still. | [optional] 
**Width** | Pointer to **int64** |  | [optional] 

## Methods

### NewStudioItem

`func NewStudioItem() *StudioItem`

NewStudioItem instantiates a new StudioItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioItemWithDefaults

`func NewStudioItemWithDefaults() *StudioItem`

NewStudioItemWithDefaults instantiates a new StudioItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *StudioItem) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *StudioItem) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *StudioItem) SetBytes(v int64)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *StudioItem) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetCreated

`func (o *StudioItem) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StudioItem) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StudioItem) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StudioItem) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDownload

`func (o *StudioItem) GetDownload() string`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *StudioItem) GetDownloadOk() (*string, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *StudioItem) SetDownload(v string)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *StudioItem) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetExpires

`func (o *StudioItem) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *StudioItem) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *StudioItem) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *StudioItem) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetHeight

`func (o *StudioItem) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *StudioItem) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *StudioItem) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *StudioItem) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *StudioItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StudioItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StudioItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StudioItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *StudioItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StudioItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StudioItem) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *StudioItem) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMime

`func (o *StudioItem) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *StudioItem) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *StudioItem) SetMime(v string)`

SetMime sets Mime field to given value.

### HasMime

`func (o *StudioItem) HasMime() bool`

HasMime returns a boolean if a field has been set.

### GetMine

`func (o *StudioItem) GetMine() bool`

GetMine returns the Mine field if non-nil, zero value otherwise.

### GetMineOk

`func (o *StudioItem) GetMineOk() (*bool, bool)`

GetMineOk returns a tuple with the Mine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMine

`func (o *StudioItem) SetMine(v bool)`

SetMine sets Mine field to given value.

### HasMine

`func (o *StudioItem) HasMine() bool`

HasMine returns a boolean if a field has been set.

### GetName

`func (o *StudioItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StudioItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StudioItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StudioItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrigin

`func (o *StudioItem) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *StudioItem) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *StudioItem) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *StudioItem) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetOwner

`func (o *StudioItem) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *StudioItem) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *StudioItem) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *StudioItem) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPages

`func (o *StudioItem) GetPages() int64`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *StudioItem) GetPagesOk() (*int64, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *StudioItem) SetPages(v int64)`

SetPages sets Pages field to given value.

### HasPages

`func (o *StudioItem) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetPreview

`func (o *StudioItem) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *StudioItem) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *StudioItem) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *StudioItem) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetRef

`func (o *StudioItem) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *StudioItem) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *StudioItem) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *StudioItem) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSeconds

`func (o *StudioItem) GetSeconds() float64`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *StudioItem) GetSecondsOk() (*float64, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *StudioItem) SetSeconds(v float64)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *StudioItem) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.

### GetShared

`func (o *StudioItem) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *StudioItem) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *StudioItem) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *StudioItem) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSource

`func (o *StudioItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StudioItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StudioItem) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StudioItem) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUrl

`func (o *StudioItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StudioItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StudioItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StudioItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWidth

`func (o *StudioItem) GetWidth() int64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *StudioItem) GetWidthOk() (*int64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *StudioItem) SetWidth(v int64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *StudioItem) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



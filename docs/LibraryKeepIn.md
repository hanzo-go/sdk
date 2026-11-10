# LibraryKeepIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **int64** | Bytes reserves an upload of this size, at most 2 GiB, when data is empty, and returns a signed PUT in upload. | [optional] 
**Data** | Pointer to **string** | Data is the bytes, base64, at most 10 MiB; omit when reserving with bytes. | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Key** | Pointer to **string** | Key makes the save idempotent: the same source and key answer the item already kept. | [optional] 
**Kind** | Pointer to **string** | Kind is image, video, audio, document, slides, type, model, code or file; absent derives it from mime. | [optional] 
**Mime** | Pointer to **string** | Mime is its media type. Required. | [optional] 
**Name** | Pointer to **string** | Name is what it is called, a file name. Required. | [optional] 
**Origin** | Pointer to **string** | Origin is the path on this site back to what made it, as /chat?thread&#x3D;…. | [optional] 
**Pages** | Pointer to **int64** |  | [optional] 
**Preview** | Pointer to **string** | Preview is a still that shows it, base64, at most 2 MiB. | [optional] 
**Project** | Pointer to **string** | Project is one of the organization&#39;s projects, by slug. | [optional] 
**Ref** | Pointer to **string** | Ref is what made it: the conversation, the run. | [optional] 
**Seconds** | Pointer to **float64** |  | [optional] 
**Source** | Pointer to **string** | Source is the product it came from: chat, dev, studio or workflow. Required. | [optional] 
**Width** | Pointer to **int64** | Width, Height, Seconds and Pages describe it when known. | [optional] 

## Methods

### NewLibraryKeepIn

`func NewLibraryKeepIn() *LibraryKeepIn`

NewLibraryKeepIn instantiates a new LibraryKeepIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryKeepInWithDefaults

`func NewLibraryKeepInWithDefaults() *LibraryKeepIn`

NewLibraryKeepInWithDefaults instantiates a new LibraryKeepIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *LibraryKeepIn) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *LibraryKeepIn) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *LibraryKeepIn) SetBytes(v int64)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *LibraryKeepIn) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetData

`func (o *LibraryKeepIn) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LibraryKeepIn) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LibraryKeepIn) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *LibraryKeepIn) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeight

`func (o *LibraryKeepIn) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *LibraryKeepIn) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *LibraryKeepIn) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *LibraryKeepIn) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetKey

`func (o *LibraryKeepIn) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LibraryKeepIn) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LibraryKeepIn) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *LibraryKeepIn) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKind

`func (o *LibraryKeepIn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LibraryKeepIn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LibraryKeepIn) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *LibraryKeepIn) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMime

`func (o *LibraryKeepIn) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *LibraryKeepIn) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *LibraryKeepIn) SetMime(v string)`

SetMime sets Mime field to given value.

### HasMime

`func (o *LibraryKeepIn) HasMime() bool`

HasMime returns a boolean if a field has been set.

### GetName

`func (o *LibraryKeepIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LibraryKeepIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LibraryKeepIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LibraryKeepIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrigin

`func (o *LibraryKeepIn) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *LibraryKeepIn) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *LibraryKeepIn) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *LibraryKeepIn) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetPages

`func (o *LibraryKeepIn) GetPages() int64`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *LibraryKeepIn) GetPagesOk() (*int64, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *LibraryKeepIn) SetPages(v int64)`

SetPages sets Pages field to given value.

### HasPages

`func (o *LibraryKeepIn) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetPreview

`func (o *LibraryKeepIn) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *LibraryKeepIn) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *LibraryKeepIn) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *LibraryKeepIn) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetProject

`func (o *LibraryKeepIn) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *LibraryKeepIn) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *LibraryKeepIn) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *LibraryKeepIn) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRef

`func (o *LibraryKeepIn) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *LibraryKeepIn) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *LibraryKeepIn) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *LibraryKeepIn) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSeconds

`func (o *LibraryKeepIn) GetSeconds() float64`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *LibraryKeepIn) GetSecondsOk() (*float64, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *LibraryKeepIn) SetSeconds(v float64)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *LibraryKeepIn) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.

### GetSource

`func (o *LibraryKeepIn) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LibraryKeepIn) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LibraryKeepIn) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LibraryKeepIn) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetWidth

`func (o *LibraryKeepIn) GetWidth() int64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *LibraryKeepIn) GetWidthOk() (*int64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *LibraryKeepIn) SetWidth(v int64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *LibraryKeepIn) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



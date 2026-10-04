# ProviderSlackFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the file&#39;s Slack ID and the handle the file read takes. | [optional] 
**Mimetype** | Pointer to **string** | Mimetype is the type Slack recorded at upload. It is the uploader&#39;s word, so it describes the file rather than deciding how the bytes are served. | [optional] 
**Name** | Pointer to **string** | Name is the filename as uploaded. | [optional] 
**Permalink** | Pointer to **string** | Permalink opens the file in Slack, for a person. It needs a Slack session and is never the bytes. | [optional] 
**Size** | Pointer to **int64** | Size is the file&#39;s length in bytes. | [optional] 
**Title** | Pointer to **string** | Title is the title Slack displays, which may differ from the filename. | [optional] 

## Methods

### NewProviderSlackFile

`func NewProviderSlackFile() *ProviderSlackFile`

NewProviderSlackFile instantiates a new ProviderSlackFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSlackFileWithDefaults

`func NewProviderSlackFileWithDefaults() *ProviderSlackFile`

NewProviderSlackFileWithDefaults instantiates a new ProviderSlackFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProviderSlackFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderSlackFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderSlackFile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProviderSlackFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMimetype

`func (o *ProviderSlackFile) GetMimetype() string`

GetMimetype returns the Mimetype field if non-nil, zero value otherwise.

### GetMimetypeOk

`func (o *ProviderSlackFile) GetMimetypeOk() (*string, bool)`

GetMimetypeOk returns a tuple with the Mimetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimetype

`func (o *ProviderSlackFile) SetMimetype(v string)`

SetMimetype sets Mimetype field to given value.

### HasMimetype

`func (o *ProviderSlackFile) HasMimetype() bool`

HasMimetype returns a boolean if a field has been set.

### GetName

`func (o *ProviderSlackFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderSlackFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderSlackFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderSlackFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermalink

`func (o *ProviderSlackFile) GetPermalink() string`

GetPermalink returns the Permalink field if non-nil, zero value otherwise.

### GetPermalinkOk

`func (o *ProviderSlackFile) GetPermalinkOk() (*string, bool)`

GetPermalinkOk returns a tuple with the Permalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermalink

`func (o *ProviderSlackFile) SetPermalink(v string)`

SetPermalink sets Permalink field to given value.

### HasPermalink

`func (o *ProviderSlackFile) HasPermalink() bool`

HasPermalink returns a boolean if a field has been set.

### GetSize

`func (o *ProviderSlackFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ProviderSlackFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ProviderSlackFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ProviderSlackFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTitle

`func (o *ProviderSlackFile) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProviderSlackFile) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProviderSlackFile) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProviderSlackFile) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



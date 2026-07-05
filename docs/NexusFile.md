# NexusFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]NexusFile**](NexusFile.md) |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**IsLeaf** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewNexusFile

`func NewNexusFile() *NexusFile`

NewNexusFile instantiates a new NexusFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexusFileWithDefaults

`func NewNexusFileWithDefaults() *NexusFile`

NewNexusFileWithDefaults instantiates a new NexusFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *NexusFile) GetChildren() []NexusFile`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *NexusFile) GetChildrenOk() (*[]NexusFile, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *NexusFile) SetChildren(v []NexusFile)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *NexusFile) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCreatedTime

`func (o *NexusFile) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *NexusFile) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *NexusFile) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *NexusFile) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetIsLeaf

`func (o *NexusFile) GetIsLeaf() bool`

GetIsLeaf returns the IsLeaf field if non-nil, zero value otherwise.

### GetIsLeafOk

`func (o *NexusFile) GetIsLeafOk() (*bool, bool)`

GetIsLeafOk returns a tuple with the IsLeaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLeaf

`func (o *NexusFile) SetIsLeaf(v bool)`

SetIsLeaf sets IsLeaf field to given value.

### HasIsLeaf

`func (o *NexusFile) HasIsLeaf() bool`

HasIsLeaf returns a boolean if a field has been set.

### GetKey

`func (o *NexusFile) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NexusFile) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *NexusFile) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *NexusFile) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSize

`func (o *NexusFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *NexusFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *NexusFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *NexusFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTitle

`func (o *NexusFile) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NexusFile) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NexusFile) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NexusFile) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *NexusFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NexusFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NexusFile) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NexusFile) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



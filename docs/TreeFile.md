# TreeFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]TreeFile**](TreeFile.md) |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**IsLeaf** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewTreeFile

`func NewTreeFile() *TreeFile`

NewTreeFile instantiates a new TreeFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTreeFileWithDefaults

`func NewTreeFileWithDefaults() *TreeFile`

NewTreeFileWithDefaults instantiates a new TreeFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *TreeFile) GetChildren() []TreeFile`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *TreeFile) GetChildrenOk() (*[]TreeFile, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *TreeFile) SetChildren(v []TreeFile)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *TreeFile) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCreatedTime

`func (o *TreeFile) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *TreeFile) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *TreeFile) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *TreeFile) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetIsLeaf

`func (o *TreeFile) GetIsLeaf() bool`

GetIsLeaf returns the IsLeaf field if non-nil, zero value otherwise.

### GetIsLeafOk

`func (o *TreeFile) GetIsLeafOk() (*bool, bool)`

GetIsLeafOk returns a tuple with the IsLeaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLeaf

`func (o *TreeFile) SetIsLeaf(v bool)`

SetIsLeaf sets IsLeaf field to given value.

### HasIsLeaf

`func (o *TreeFile) HasIsLeaf() bool`

HasIsLeaf returns a boolean if a field has been set.

### GetKey

`func (o *TreeFile) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TreeFile) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TreeFile) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TreeFile) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSize

`func (o *TreeFile) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TreeFile) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TreeFile) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *TreeFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTitle

`func (o *TreeFile) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TreeFile) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TreeFile) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TreeFile) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *TreeFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TreeFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TreeFile) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TreeFile) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



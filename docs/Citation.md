# Citation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndLine** | Pointer to **int64** | EndLine is its last line, inclusive. | [optional] 
**File** | Pointer to **string** | File is the path inside the repo, relative to its root. | [optional] 
**Line** | Pointer to **int64** | Line is the first line of the cited region, 1-based. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository the cited code lives in (\&quot;owner/name\&quot;), absent when the ask was already scoped to one. | [optional] 
**Symbol** | Pointer to **string** | Symbol is the declaration the region belongs to, when it belongs to one. | [optional] 

## Methods

### NewCitation

`func NewCitation() *Citation`

NewCitation instantiates a new Citation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCitationWithDefaults

`func NewCitationWithDefaults() *Citation`

NewCitationWithDefaults instantiates a new Citation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndLine

`func (o *Citation) GetEndLine() int64`

GetEndLine returns the EndLine field if non-nil, zero value otherwise.

### GetEndLineOk

`func (o *Citation) GetEndLineOk() (*int64, bool)`

GetEndLineOk returns a tuple with the EndLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLine

`func (o *Citation) SetEndLine(v int64)`

SetEndLine sets EndLine field to given value.

### HasEndLine

`func (o *Citation) HasEndLine() bool`

HasEndLine returns a boolean if a field has been set.

### GetFile

`func (o *Citation) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *Citation) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *Citation) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *Citation) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetLine

`func (o *Citation) GetLine() int64`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *Citation) GetLineOk() (*int64, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *Citation) SetLine(v int64)`

SetLine sets Line field to given value.

### HasLine

`func (o *Citation) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetRepo

`func (o *Citation) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *Citation) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *Citation) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *Citation) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSymbol

`func (o *Citation) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Citation) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Citation) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Citation) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



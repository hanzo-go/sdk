# TreeEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lang** | Pointer to **string** | Lang is the language the indexer parsed the file as (\&quot;go\&quot;, \&quot;python\&quot;, …), or empty when it recognised none — in which case Symbols is 0 because nothing was extracted, not because the file declares nothing. | [optional] 
**Path** | Pointer to **string** | Path is the file, relative to the repo root. The list is ordered by it, so a reader can see module layout without sorting. | [optional] 
**Symbols** | Pointer to **int32** | Symbols is how many top-level declarations the file defines. A file with none is still listed: the file set is the authority here and the counts decorate it. | [optional] 

## Methods

### NewTreeEntry

`func NewTreeEntry() *TreeEntry`

NewTreeEntry instantiates a new TreeEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTreeEntryWithDefaults

`func NewTreeEntryWithDefaults() *TreeEntry`

NewTreeEntryWithDefaults instantiates a new TreeEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLang

`func (o *TreeEntry) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *TreeEntry) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *TreeEntry) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *TreeEntry) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetPath

`func (o *TreeEntry) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TreeEntry) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TreeEntry) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *TreeEntry) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSymbols

`func (o *TreeEntry) GetSymbols() int32`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *TreeEntry) GetSymbolsOk() (*int32, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *TreeEntry) SetSymbols(v int32)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *TreeEntry) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



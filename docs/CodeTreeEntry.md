# CodeTreeEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lang** | Pointer to **string** | Lang is the language the indexer parsed the file as (\&quot;go\&quot;, \&quot;python\&quot;, …), or empty when it recognised none — in which case Symbols is 0 because nothing was extracted, not because the file declares nothing. | [optional] 
**Path** | Pointer to **string** | Path is the file, relative to the repo root. The list is ordered by it, so a reader can see module layout without sorting. | [optional] 
**Symbols** | Pointer to **int64** | Symbols is how many top-level declarations the file defines. A file with none is still listed: the file set is the authority here and the counts decorate it. | [optional] 

## Methods

### NewCodeTreeEntry

`func NewCodeTreeEntry() *CodeTreeEntry`

NewCodeTreeEntry instantiates a new CodeTreeEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeTreeEntryWithDefaults

`func NewCodeTreeEntryWithDefaults() *CodeTreeEntry`

NewCodeTreeEntryWithDefaults instantiates a new CodeTreeEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLang

`func (o *CodeTreeEntry) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *CodeTreeEntry) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *CodeTreeEntry) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *CodeTreeEntry) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetPath

`func (o *CodeTreeEntry) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CodeTreeEntry) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CodeTreeEntry) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CodeTreeEntry) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSymbols

`func (o *CodeTreeEntry) GetSymbols() int64`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *CodeTreeEntry) GetSymbolsOk() (*int64, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *CodeTreeEntry) SetSymbols(v int64)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *CodeTreeEntry) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



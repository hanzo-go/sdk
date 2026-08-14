# Answer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cold** | Pointer to **bool** | Cold reports that this request paid to PREPARE the revision — the tree write, the dependency fetch and the language server&#39;s first index. It is the billed event, surfaced so a caller can see what it was charged for. | [optional] 
**Completions** | Pointer to [**[]Completion**](Completion.md) |  | [optional] 
**Diagnostics** | Pointer to [**[]Diagnostic**](Diagnostic.md) |  | [optional] 
**Hover** | Pointer to **string** |  | [optional] 
**Lang** | Pointer to **string** |  | [optional] 
**Locations** | Pointer to [**[]Location**](Location.md) |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Rev** | Pointer to **string** |  | [optional] 
**Symbols** | Pointer to [**[]Symbol**](Symbol.md) |  | [optional] 

## Methods

### NewAnswer

`func NewAnswer() *Answer`

NewAnswer instantiates a new Answer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerWithDefaults

`func NewAnswerWithDefaults() *Answer`

NewAnswerWithDefaults instantiates a new Answer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCold

`func (o *Answer) GetCold() bool`

GetCold returns the Cold field if non-nil, zero value otherwise.

### GetColdOk

`func (o *Answer) GetColdOk() (*bool, bool)`

GetColdOk returns a tuple with the Cold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCold

`func (o *Answer) SetCold(v bool)`

SetCold sets Cold field to given value.

### HasCold

`func (o *Answer) HasCold() bool`

HasCold returns a boolean if a field has been set.

### GetCompletions

`func (o *Answer) GetCompletions() []Completion`

GetCompletions returns the Completions field if non-nil, zero value otherwise.

### GetCompletionsOk

`func (o *Answer) GetCompletionsOk() (*[]Completion, bool)`

GetCompletionsOk returns a tuple with the Completions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletions

`func (o *Answer) SetCompletions(v []Completion)`

SetCompletions sets Completions field to given value.

### HasCompletions

`func (o *Answer) HasCompletions() bool`

HasCompletions returns a boolean if a field has been set.

### GetDiagnostics

`func (o *Answer) GetDiagnostics() []Diagnostic`

GetDiagnostics returns the Diagnostics field if non-nil, zero value otherwise.

### GetDiagnosticsOk

`func (o *Answer) GetDiagnosticsOk() (*[]Diagnostic, bool)`

GetDiagnosticsOk returns a tuple with the Diagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnostics

`func (o *Answer) SetDiagnostics(v []Diagnostic)`

SetDiagnostics sets Diagnostics field to given value.

### HasDiagnostics

`func (o *Answer) HasDiagnostics() bool`

HasDiagnostics returns a boolean if a field has been set.

### GetHover

`func (o *Answer) GetHover() string`

GetHover returns the Hover field if non-nil, zero value otherwise.

### GetHoverOk

`func (o *Answer) GetHoverOk() (*string, bool)`

GetHoverOk returns a tuple with the Hover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHover

`func (o *Answer) SetHover(v string)`

SetHover sets Hover field to given value.

### HasHover

`func (o *Answer) HasHover() bool`

HasHover returns a boolean if a field has been set.

### GetLang

`func (o *Answer) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *Answer) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *Answer) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *Answer) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetLocations

`func (o *Answer) GetLocations() []Location`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *Answer) GetLocationsOk() (*[]Location, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *Answer) SetLocations(v []Location)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *Answer) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetOp

`func (o *Answer) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *Answer) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *Answer) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *Answer) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetPath

`func (o *Answer) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Answer) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Answer) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Answer) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRepo

`func (o *Answer) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *Answer) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *Answer) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *Answer) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRev

`func (o *Answer) GetRev() string`

GetRev returns the Rev field if non-nil, zero value otherwise.

### GetRevOk

`func (o *Answer) GetRevOk() (*string, bool)`

GetRevOk returns a tuple with the Rev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRev

`func (o *Answer) SetRev(v string)`

SetRev sets Rev field to given value.

### HasRev

`func (o *Answer) HasRev() bool`

HasRev returns a boolean if a field has been set.

### GetSymbols

`func (o *Answer) GetSymbols() []Symbol`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *Answer) GetSymbolsOk() (*[]Symbol, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *Answer) SetSymbols(v []Symbol)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *Answer) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



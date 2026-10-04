# LspAnswer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cold** | Pointer to **bool** | Cold reports that this request paid to PREPARE the revision — the tree write, the dependency fetch and the language server&#39;s first index. It is the billed event, surfaced so a caller can see what it was charged for. | [optional] 
**Completions** | Pointer to [**[]LspCompletion**](LspCompletion.md) | Completions is complete&#39;s answer: the candidates at the position, typed and resolved through the repository&#39;s dependencies rather than guessed from text. | [optional] 
**Diagnostics** | Pointer to [**[]LspDiagnostic**](LspDiagnostic.md) | Diagnostics is diagnostics&#39; answer: every problem the server finds in the whole file, position ignored. Empty means it found none. | [optional] 
**Hover** | Pointer to **string** | Hover is hover&#39;s answer: the type and documentation as the language server itself renders them, so it is prose meant to be shown, not parsed. | [optional] 
**Lang** | Pointer to **string** | Lang is the language the server that answered speaks (\&quot;go\&quot;), as the daemon reports it. Empty when the daemon named none. | [optional] 
**Locations** | Pointer to [**[]LspLocation**](LspLocation.md) | Locations is locate&#39;s answer: where the symbol is defined, referenced, typed or implemented, per the relation asked for. Empty means the server resolved nothing there, which is an answer. | [optional] 
**Op** | Pointer to **string** | Op is the question that was asked: hover, locate, symbols, diagnostics or complete. It names which result field below is the populated one. | [optional] 
**Path** | Pointer to **string** | Path is the repo-relative file the question was about, echoed back. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository the question was about, echoed back. | [optional] 
**Rev** | Pointer to **string** | Rev is the RESOLVED commit sha, never the branch or tag that was asked for. It is what makes an answer re-askable: a branch moves, this does not. | [optional] 
**Symbols** | Pointer to [**[]LspSymbol**](LspSymbol.md) | Symbols is symbols&#39; answer: the file&#39;s whole outline, position ignored. | [optional] 

## Methods

### NewLspAnswer

`func NewLspAnswer() *LspAnswer`

NewLspAnswer instantiates a new LspAnswer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLspAnswerWithDefaults

`func NewLspAnswerWithDefaults() *LspAnswer`

NewLspAnswerWithDefaults instantiates a new LspAnswer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCold

`func (o *LspAnswer) GetCold() bool`

GetCold returns the Cold field if non-nil, zero value otherwise.

### GetColdOk

`func (o *LspAnswer) GetColdOk() (*bool, bool)`

GetColdOk returns a tuple with the Cold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCold

`func (o *LspAnswer) SetCold(v bool)`

SetCold sets Cold field to given value.

### HasCold

`func (o *LspAnswer) HasCold() bool`

HasCold returns a boolean if a field has been set.

### GetCompletions

`func (o *LspAnswer) GetCompletions() []LspCompletion`

GetCompletions returns the Completions field if non-nil, zero value otherwise.

### GetCompletionsOk

`func (o *LspAnswer) GetCompletionsOk() (*[]LspCompletion, bool)`

GetCompletionsOk returns a tuple with the Completions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletions

`func (o *LspAnswer) SetCompletions(v []LspCompletion)`

SetCompletions sets Completions field to given value.

### HasCompletions

`func (o *LspAnswer) HasCompletions() bool`

HasCompletions returns a boolean if a field has been set.

### GetDiagnostics

`func (o *LspAnswer) GetDiagnostics() []LspDiagnostic`

GetDiagnostics returns the Diagnostics field if non-nil, zero value otherwise.

### GetDiagnosticsOk

`func (o *LspAnswer) GetDiagnosticsOk() (*[]LspDiagnostic, bool)`

GetDiagnosticsOk returns a tuple with the Diagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnostics

`func (o *LspAnswer) SetDiagnostics(v []LspDiagnostic)`

SetDiagnostics sets Diagnostics field to given value.

### HasDiagnostics

`func (o *LspAnswer) HasDiagnostics() bool`

HasDiagnostics returns a boolean if a field has been set.

### GetHover

`func (o *LspAnswer) GetHover() string`

GetHover returns the Hover field if non-nil, zero value otherwise.

### GetHoverOk

`func (o *LspAnswer) GetHoverOk() (*string, bool)`

GetHoverOk returns a tuple with the Hover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHover

`func (o *LspAnswer) SetHover(v string)`

SetHover sets Hover field to given value.

### HasHover

`func (o *LspAnswer) HasHover() bool`

HasHover returns a boolean if a field has been set.

### GetLang

`func (o *LspAnswer) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *LspAnswer) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *LspAnswer) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *LspAnswer) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetLocations

`func (o *LspAnswer) GetLocations() []LspLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *LspAnswer) GetLocationsOk() (*[]LspLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *LspAnswer) SetLocations(v []LspLocation)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *LspAnswer) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetOp

`func (o *LspAnswer) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *LspAnswer) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *LspAnswer) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *LspAnswer) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetPath

`func (o *LspAnswer) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LspAnswer) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LspAnswer) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LspAnswer) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRepo

`func (o *LspAnswer) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *LspAnswer) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *LspAnswer) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *LspAnswer) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRev

`func (o *LspAnswer) GetRev() string`

GetRev returns the Rev field if non-nil, zero value otherwise.

### GetRevOk

`func (o *LspAnswer) GetRevOk() (*string, bool)`

GetRevOk returns a tuple with the Rev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRev

`func (o *LspAnswer) SetRev(v string)`

SetRev sets Rev field to given value.

### HasRev

`func (o *LspAnswer) HasRev() bool`

HasRev returns a boolean if a field has been set.

### GetSymbols

`func (o *LspAnswer) GetSymbols() []LspSymbol`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *LspAnswer) GetSymbolsOk() (*[]LspSymbol, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *LspAnswer) SetSymbols(v []LspSymbol)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *LspAnswer) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



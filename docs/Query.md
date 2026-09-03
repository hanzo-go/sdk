# Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Character** | Pointer to **int64** | Character is a 0-based UTF-16 code-unit offset within Line, per the LSP specification — not a byte offset and not a rune index. | [optional] 
**Line** | Pointer to **int64** | Line is 0-based, per the LSP specification. | [optional] 
**Path** | Pointer to **string** | Path is the repo-relative file, e.g. \&quot;apps/lsp/lsp.go\&quot;. | [optional] 
**Relation** | Pointer to **string** | Relation refines locate: definition, reference, type or implementation. Empty means definition. Every other op ignores it. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository NAME within the caller&#39;s own org, e.g. \&quot;cloud\&quot;. Not a URL and not an owner/name pair: the owner is the validated principal&#39;s org, so this names a repository the caller already owns. | [optional] 
**Rev** | Pointer to **string** | Rev is a branch, tag or commit sha. Empty means the default branch. It is resolved to a commit before anything else happens, so an answer is always about one immutable tree. | [optional] 

## Methods

### NewQuery

`func NewQuery() *Query`

NewQuery instantiates a new Query object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryWithDefaults

`func NewQueryWithDefaults() *Query`

NewQueryWithDefaults instantiates a new Query object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharacter

`func (o *Query) GetCharacter() int64`

GetCharacter returns the Character field if non-nil, zero value otherwise.

### GetCharacterOk

`func (o *Query) GetCharacterOk() (*int64, bool)`

GetCharacterOk returns a tuple with the Character field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacter

`func (o *Query) SetCharacter(v int64)`

SetCharacter sets Character field to given value.

### HasCharacter

`func (o *Query) HasCharacter() bool`

HasCharacter returns a boolean if a field has been set.

### GetLine

`func (o *Query) GetLine() int64`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *Query) GetLineOk() (*int64, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *Query) SetLine(v int64)`

SetLine sets Line field to given value.

### HasLine

`func (o *Query) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetPath

`func (o *Query) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Query) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Query) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Query) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRelation

`func (o *Query) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *Query) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *Query) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *Query) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetRepo

`func (o *Query) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *Query) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *Query) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *Query) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRev

`func (o *Query) GetRev() string`

GetRev returns the Rev field if non-nil, zero value otherwise.

### GetRevOk

`func (o *Query) GetRevOk() (*string, bool)`

GetRevOk returns a tuple with the Rev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRev

`func (o *Query) SetRev(v string)`

SetRev sets Rev field to given value.

### HasRev

`func (o *Query) HasRev() bool`

HasRev returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



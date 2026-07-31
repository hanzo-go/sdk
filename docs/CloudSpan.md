# CloudSpan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndLine** | Pointer to **int32** |  | [optional] 
**File** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Line** | Pointer to **int32** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** | context: match | definition | caller | [optional] 
**Score** | Pointer to **float32** |  | [optional] 
**Snippet** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Tier** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudSpan

`func NewCloudSpan() *CloudSpan`

NewCloudSpan instantiates a new CloudSpan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSpanWithDefaults

`func NewCloudSpanWithDefaults() *CloudSpan`

NewCloudSpanWithDefaults instantiates a new CloudSpan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndLine

`func (o *CloudSpan) GetEndLine() int32`

GetEndLine returns the EndLine field if non-nil, zero value otherwise.

### GetEndLineOk

`func (o *CloudSpan) GetEndLineOk() (*int32, bool)`

GetEndLineOk returns a tuple with the EndLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLine

`func (o *CloudSpan) SetEndLine(v int32)`

SetEndLine sets EndLine field to given value.

### HasEndLine

`func (o *CloudSpan) HasEndLine() bool`

HasEndLine returns a boolean if a field has been set.

### GetFile

`func (o *CloudSpan) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *CloudSpan) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *CloudSpan) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *CloudSpan) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetKind

`func (o *CloudSpan) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudSpan) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudSpan) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudSpan) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLine

`func (o *CloudSpan) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *CloudSpan) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *CloudSpan) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *CloudSpan) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetRepo

`func (o *CloudSpan) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudSpan) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudSpan) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudSpan) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRole

`func (o *CloudSpan) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CloudSpan) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CloudSpan) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CloudSpan) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetScore

`func (o *CloudSpan) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CloudSpan) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CloudSpan) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *CloudSpan) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSnippet

`func (o *CloudSpan) GetSnippet() string`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *CloudSpan) GetSnippetOk() (*string, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *CloudSpan) SetSnippet(v string)`

SetSnippet sets Snippet field to given value.

### HasSnippet

`func (o *CloudSpan) HasSnippet() bool`

HasSnippet returns a boolean if a field has been set.

### GetSymbol

`func (o *CloudSpan) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CloudSpan) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CloudSpan) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CloudSpan) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTier

`func (o *CloudSpan) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *CloudSpan) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *CloudSpan) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *CloudSpan) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



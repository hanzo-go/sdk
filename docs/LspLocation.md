# LspLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**External** | Pointer to **bool** | External is true when the answer left the repository — the case a static index cannot answer, and the reason this service resolves through dependencies. | [optional] 
**Path** | Pointer to **string** | Path is repo-relative while External is false, and the module coordinate (\&quot;golang.org/x/mod@v0.14.0/semver/semver.go\&quot;) once it is true. | [optional] 
**Range** | Pointer to [**LspRange**](LspRange.md) | Range is the span inside that file, in LSP positions. | [optional] 

## Methods

### NewLspLocation

`func NewLspLocation() *LspLocation`

NewLspLocation instantiates a new LspLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLspLocationWithDefaults

`func NewLspLocationWithDefaults() *LspLocation`

NewLspLocationWithDefaults instantiates a new LspLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternal

`func (o *LspLocation) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *LspLocation) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *LspLocation) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *LspLocation) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetPath

`func (o *LspLocation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LspLocation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LspLocation) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LspLocation) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRange

`func (o *LspLocation) GetRange() LspRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *LspLocation) GetRangeOk() (*LspRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *LspLocation) SetRange(v LspRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *LspLocation) HasRange() bool`

HasRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# LspSymbol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** | Detail is the server&#39;s short elaboration, typically the signature. Absent when it offered none. | [optional] 
**Kind** | Pointer to **int64** | Kind is the LSP SymbolKind number (5 class, 6 method, 12 function, 23 struct, …), passed through rather than translated to a word — these callers already speak LSP, and inventing a second vocabulary is how the two drift. | [optional] 
**Name** | Pointer to **string** | Name is the declared identifier. | [optional] 
**Range** | Pointer to [**LspRange**](LspRange.md) | Range is the declaration&#39;s span in the file. | [optional] 

## Methods

### NewLspSymbol

`func NewLspSymbol() *LspSymbol`

NewLspSymbol instantiates a new LspSymbol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLspSymbolWithDefaults

`func NewLspSymbolWithDefaults() *LspSymbol`

NewLspSymbolWithDefaults instantiates a new LspSymbol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *LspSymbol) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *LspSymbol) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *LspSymbol) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *LspSymbol) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetKind

`func (o *LspSymbol) GetKind() int64`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LspSymbol) GetKindOk() (*int64, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LspSymbol) SetKind(v int64)`

SetKind sets Kind field to given value.

### HasKind

`func (o *LspSymbol) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *LspSymbol) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LspSymbol) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LspSymbol) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LspSymbol) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRange

`func (o *LspSymbol) GetRange() LspRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *LspSymbol) GetRangeOk() (*LspRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *LspSymbol) SetRange(v LspRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *LspSymbol) HasRange() bool`

HasRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



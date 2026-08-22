# Symbol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** | Detail is the server&#39;s short elaboration, typically the signature. Absent when it offered none. | [optional] 
**Kind** | Pointer to **int32** | Kind is the LSP SymbolKind number (5 class, 6 method, 12 function, 23 struct, …), passed through rather than translated to a word — these callers already speak LSP, and inventing a second vocabulary is how the two drift. | [optional] 
**Name** | Pointer to **string** | Name is the declared identifier. | [optional] 
**Range** | Pointer to [**Range**](Range.md) | Range is the declaration&#39;s span in the file. | [optional] 

## Methods

### NewSymbol

`func NewSymbol() *Symbol`

NewSymbol instantiates a new Symbol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolWithDefaults

`func NewSymbolWithDefaults() *Symbol`

NewSymbolWithDefaults instantiates a new Symbol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *Symbol) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Symbol) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Symbol) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *Symbol) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetKind

`func (o *Symbol) GetKind() int32`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Symbol) GetKindOk() (*int32, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Symbol) SetKind(v int32)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Symbol) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *Symbol) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Symbol) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Symbol) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Symbol) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRange

`func (o *Symbol) GetRange() Range`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *Symbol) GetRangeOk() (*Range, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *Symbol) SetRange(v Range)`

SetRange sets Range field to given value.

### HasRange

`func (o *Symbol) HasRange() bool`

HasRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Position

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Character** | Pointer to **int32** | Character is a 0-based UTF-16 code-unit offset within Line, per the LSP specification: not a byte offset and not a rune index. An emoji before the cursor counts as one here and as two in Go&#39;s arithmetic. | [optional] 
**Line** | Pointer to **int32** | Line is 0-BASED, per the LSP specification — one less than the line an editor shows a human. | [optional] 

## Methods

### NewPosition

`func NewPosition() *Position`

NewPosition instantiates a new Position object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionWithDefaults

`func NewPositionWithDefaults() *Position`

NewPositionWithDefaults instantiates a new Position object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharacter

`func (o *Position) GetCharacter() int32`

GetCharacter returns the Character field if non-nil, zero value otherwise.

### GetCharacterOk

`func (o *Position) GetCharacterOk() (*int32, bool)`

GetCharacterOk returns a tuple with the Character field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacter

`func (o *Position) SetCharacter(v int32)`

SetCharacter sets Character field to given value.

### HasCharacter

`func (o *Position) HasCharacter() bool`

HasCharacter returns a boolean if a field has been set.

### GetLine

`func (o *Position) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *Position) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *Position) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *Position) HasLine() bool`

HasLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



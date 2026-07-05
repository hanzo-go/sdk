# AutomationsCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PieceCount** | Pointer to **int32** |  | [optional] 
**Pieces** | Pointer to [**[]AutomationsPieceMetadata**](AutomationsPieceMetadata.md) |  | [optional] 

## Methods

### NewAutomationsCatalog

`func NewAutomationsCatalog() *AutomationsCatalog`

NewAutomationsCatalog instantiates a new AutomationsCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationsCatalogWithDefaults

`func NewAutomationsCatalogWithDefaults() *AutomationsCatalog`

NewAutomationsCatalogWithDefaults instantiates a new AutomationsCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPieceCount

`func (o *AutomationsCatalog) GetPieceCount() int32`

GetPieceCount returns the PieceCount field if non-nil, zero value otherwise.

### GetPieceCountOk

`func (o *AutomationsCatalog) GetPieceCountOk() (*int32, bool)`

GetPieceCountOk returns a tuple with the PieceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPieceCount

`func (o *AutomationsCatalog) SetPieceCount(v int32)`

SetPieceCount sets PieceCount field to given value.

### HasPieceCount

`func (o *AutomationsCatalog) HasPieceCount() bool`

HasPieceCount returns a boolean if a field has been set.

### GetPieces

`func (o *AutomationsCatalog) GetPieces() []AutomationsPieceMetadata`

GetPieces returns the Pieces field if non-nil, zero value otherwise.

### GetPiecesOk

`func (o *AutomationsCatalog) GetPiecesOk() (*[]AutomationsPieceMetadata, bool)`

GetPiecesOk returns a tuple with the Pieces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPieces

`func (o *AutomationsCatalog) SetPieces(v []AutomationsPieceMetadata)`

SetPieces sets Pieces field to given value.

### HasPieces

`func (o *AutomationsCatalog) HasPieces() bool`

HasPieces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



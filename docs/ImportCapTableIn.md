# ImportCapTableIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | Pointer to **string** | Range is an optional A1 range within the sheet; empty reads the default range. | [optional] 
**SpreadsheetId** | Pointer to **string** | SpreadsheetID is a Google Sheets id. Required. | [optional] 

## Methods

### NewImportCapTableIn

`func NewImportCapTableIn() *ImportCapTableIn`

NewImportCapTableIn instantiates a new ImportCapTableIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportCapTableInWithDefaults

`func NewImportCapTableInWithDefaults() *ImportCapTableIn`

NewImportCapTableInWithDefaults instantiates a new ImportCapTableIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *ImportCapTableIn) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *ImportCapTableIn) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *ImportCapTableIn) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *ImportCapTableIn) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSpreadsheetId

`func (o *ImportCapTableIn) GetSpreadsheetId() string`

GetSpreadsheetId returns the SpreadsheetId field if non-nil, zero value otherwise.

### GetSpreadsheetIdOk

`func (o *ImportCapTableIn) GetSpreadsheetIdOk() (*string, bool)`

GetSpreadsheetIdOk returns a tuple with the SpreadsheetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpreadsheetId

`func (o *ImportCapTableIn) SetSpreadsheetId(v string)`

SetSpreadsheetId sets SpreadsheetId field to given value.

### HasSpreadsheetId

`func (o *ImportCapTableIn) HasSpreadsheetId() bool`

HasSpreadsheetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



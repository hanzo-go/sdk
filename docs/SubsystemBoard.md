# SubsystemBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** |  | [optional] 
**Range** | Pointer to **string** |  | [optional] 
**Rows** | Pointer to [**[]SubsystemRow**](SubsystemRow.md) |  | [optional] 
**Sources** | Pointer to [**[]SourceStatus**](SourceStatus.md) |  | [optional] 
**Start** | Pointer to **string** |  | [optional] 
**Totals** | Pointer to [**SubsystemTotals**](SubsystemTotals.md) |  | [optional] 

## Methods

### NewSubsystemBoard

`func NewSubsystemBoard() *SubsystemBoard`

NewSubsystemBoard instantiates a new SubsystemBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubsystemBoardWithDefaults

`func NewSubsystemBoardWithDefaults() *SubsystemBoard`

NewSubsystemBoardWithDefaults instantiates a new SubsystemBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *SubsystemBoard) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SubsystemBoard) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SubsystemBoard) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SubsystemBoard) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetRange

`func (o *SubsystemBoard) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *SubsystemBoard) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *SubsystemBoard) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *SubsystemBoard) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRows

`func (o *SubsystemBoard) GetRows() []SubsystemRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *SubsystemBoard) GetRowsOk() (*[]SubsystemRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *SubsystemBoard) SetRows(v []SubsystemRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *SubsystemBoard) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetSources

`func (o *SubsystemBoard) GetSources() []SourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *SubsystemBoard) GetSourcesOk() (*[]SourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *SubsystemBoard) SetSources(v []SourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *SubsystemBoard) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetStart

`func (o *SubsystemBoard) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SubsystemBoard) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SubsystemBoard) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *SubsystemBoard) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTotals

`func (o *SubsystemBoard) GetTotals() SubsystemTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *SubsystemBoard) GetTotalsOk() (*SubsystemTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *SubsystemBoard) SetTotals(v SubsystemTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *SubsystemBoard) HasTotals() bool`

HasTotals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



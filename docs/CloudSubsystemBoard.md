# CloudSubsystemBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** |  | [optional] 
**Range** | Pointer to **string** |  | [optional] 
**Rows** | Pointer to [**[]CloudSubsystemRow**](CloudSubsystemRow.md) |  | [optional] 
**Sources** | Pointer to [**[]CloudSourceStatus**](CloudSourceStatus.md) |  | [optional] 
**Start** | Pointer to **string** |  | [optional] 
**Totals** | Pointer to [**CloudSubsystemTotals**](CloudSubsystemTotals.md) |  | [optional] 

## Methods

### NewCloudSubsystemBoard

`func NewCloudSubsystemBoard() *CloudSubsystemBoard`

NewCloudSubsystemBoard instantiates a new CloudSubsystemBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSubsystemBoardWithDefaults

`func NewCloudSubsystemBoardWithDefaults() *CloudSubsystemBoard`

NewCloudSubsystemBoardWithDefaults instantiates a new CloudSubsystemBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *CloudSubsystemBoard) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CloudSubsystemBoard) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CloudSubsystemBoard) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CloudSubsystemBoard) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetRange

`func (o *CloudSubsystemBoard) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CloudSubsystemBoard) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CloudSubsystemBoard) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *CloudSubsystemBoard) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRows

`func (o *CloudSubsystemBoard) GetRows() []CloudSubsystemRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *CloudSubsystemBoard) GetRowsOk() (*[]CloudSubsystemRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *CloudSubsystemBoard) SetRows(v []CloudSubsystemRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *CloudSubsystemBoard) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetSources

`func (o *CloudSubsystemBoard) GetSources() []CloudSourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CloudSubsystemBoard) GetSourcesOk() (*[]CloudSourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CloudSubsystemBoard) SetSources(v []CloudSourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *CloudSubsystemBoard) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetStart

`func (o *CloudSubsystemBoard) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CloudSubsystemBoard) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CloudSubsystemBoard) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *CloudSubsystemBoard) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTotals

`func (o *CloudSubsystemBoard) GetTotals() CloudSubsystemTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *CloudSubsystemBoard) GetTotalsOk() (*CloudSubsystemTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *CloudSubsystemBoard) SetTotals(v CloudSubsystemTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *CloudSubsystemBoard) HasTotals() bool`

HasTotals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



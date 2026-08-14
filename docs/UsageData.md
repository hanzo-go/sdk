# UsageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByModel** | Pointer to [**[]UsageByModel**](UsageByModel.md) |  | [optional] 
**Series** | Pointer to [**[]UsagePoint**](UsagePoint.md) |  | [optional] 
**Totals** | Pointer to [**UsageTotals**](UsageTotals.md) |  | [optional] 

## Methods

### NewUsageData

`func NewUsageData() *UsageData`

NewUsageData instantiates a new UsageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageDataWithDefaults

`func NewUsageDataWithDefaults() *UsageData`

NewUsageDataWithDefaults instantiates a new UsageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByModel

`func (o *UsageData) GetByModel() []UsageByModel`

GetByModel returns the ByModel field if non-nil, zero value otherwise.

### GetByModelOk

`func (o *UsageData) GetByModelOk() (*[]UsageByModel, bool)`

GetByModelOk returns a tuple with the ByModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByModel

`func (o *UsageData) SetByModel(v []UsageByModel)`

SetByModel sets ByModel field to given value.

### HasByModel

`func (o *UsageData) HasByModel() bool`

HasByModel returns a boolean if a field has been set.

### GetSeries

`func (o *UsageData) GetSeries() []UsagePoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *UsageData) GetSeriesOk() (*[]UsagePoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *UsageData) SetSeries(v []UsagePoint)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *UsageData) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetTotals

`func (o *UsageData) GetTotals() UsageTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *UsageData) GetTotalsOk() (*UsageTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *UsageData) SetTotals(v UsageTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *UsageData) HasTotals() bool`

HasTotals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



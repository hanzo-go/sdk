# AdminUsageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Totals** | Pointer to [**AdminUsageTotals**](AdminUsageTotals.md) |  | [optional] 
**Series** | Pointer to [**[]AdminUsagePoint**](AdminUsagePoint.md) |  | [optional] 
**ByProduct** | Pointer to [**[]AdminUsageByProduct**](AdminUsageByProduct.md) |  | [optional] 

## Methods

### NewAdminUsageData

`func NewAdminUsageData() *AdminUsageData`

NewAdminUsageData instantiates a new AdminUsageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUsageDataWithDefaults

`func NewAdminUsageDataWithDefaults() *AdminUsageData`

NewAdminUsageDataWithDefaults instantiates a new AdminUsageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotals

`func (o *AdminUsageData) GetTotals() AdminUsageTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *AdminUsageData) GetTotalsOk() (*AdminUsageTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *AdminUsageData) SetTotals(v AdminUsageTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *AdminUsageData) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetSeries

`func (o *AdminUsageData) GetSeries() []AdminUsagePoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *AdminUsageData) GetSeriesOk() (*[]AdminUsagePoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *AdminUsageData) SetSeries(v []AdminUsagePoint)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *AdminUsageData) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetByProduct

`func (o *AdminUsageData) GetByProduct() []AdminUsageByProduct`

GetByProduct returns the ByProduct field if non-nil, zero value otherwise.

### GetByProductOk

`func (o *AdminUsageData) GetByProductOk() (*[]AdminUsageByProduct, bool)`

GetByProductOk returns a tuple with the ByProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByProduct

`func (o *AdminUsageData) SetByProduct(v []AdminUsageByProduct)`

SetByProduct sets ByProduct field to given value.

### HasByProduct

`func (o *AdminUsageData) HasByProduct() bool`

HasByProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



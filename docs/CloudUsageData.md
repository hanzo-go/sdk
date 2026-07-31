# CloudUsageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByProduct** | Pointer to [**[]CloudUsageByProduct**](CloudUsageByProduct.md) |  | [optional] 
**Series** | Pointer to [**[]CloudUsagePoint**](CloudUsagePoint.md) |  | [optional] 
**Totals** | Pointer to [**CloudUsageTotals**](CloudUsageTotals.md) |  | [optional] 

## Methods

### NewCloudUsageData

`func NewCloudUsageData() *CloudUsageData`

NewCloudUsageData instantiates a new CloudUsageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudUsageDataWithDefaults

`func NewCloudUsageDataWithDefaults() *CloudUsageData`

NewCloudUsageDataWithDefaults instantiates a new CloudUsageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByProduct

`func (o *CloudUsageData) GetByProduct() []CloudUsageByProduct`

GetByProduct returns the ByProduct field if non-nil, zero value otherwise.

### GetByProductOk

`func (o *CloudUsageData) GetByProductOk() (*[]CloudUsageByProduct, bool)`

GetByProductOk returns a tuple with the ByProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByProduct

`func (o *CloudUsageData) SetByProduct(v []CloudUsageByProduct)`

SetByProduct sets ByProduct field to given value.

### HasByProduct

`func (o *CloudUsageData) HasByProduct() bool`

HasByProduct returns a boolean if a field has been set.

### GetSeries

`func (o *CloudUsageData) GetSeries() []CloudUsagePoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *CloudUsageData) GetSeriesOk() (*[]CloudUsagePoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *CloudUsageData) SetSeries(v []CloudUsagePoint)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *CloudUsageData) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetTotals

`func (o *CloudUsageData) GetTotals() CloudUsageTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *CloudUsageData) GetTotalsOk() (*CloudUsageTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *CloudUsageData) SetTotals(v CloudUsageTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *CloudUsageData) HasTotals() bool`

HasTotals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



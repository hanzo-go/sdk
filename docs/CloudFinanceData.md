# CloudFinanceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to [**CloudFinanceCost**](CloudFinanceCost.md) |  | [optional] 
**Derived** | Pointer to [**CloudFinanceDerived**](CloudFinanceDerived.md) |  | [optional] 
**GeneratedAt** | Pointer to **string** |  | [optional] 
**Revenue** | Pointer to [**CloudFinanceRevenue**](CloudFinanceRevenue.md) |  | [optional] 
**Sources** | Pointer to [**[]CloudSourceStatus**](CloudSourceStatus.md) |  | [optional] 

## Methods

### NewCloudFinanceData

`func NewCloudFinanceData() *CloudFinanceData`

NewCloudFinanceData instantiates a new CloudFinanceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFinanceDataWithDefaults

`func NewCloudFinanceDataWithDefaults() *CloudFinanceData`

NewCloudFinanceDataWithDefaults instantiates a new CloudFinanceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *CloudFinanceData) GetCost() CloudFinanceCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CloudFinanceData) GetCostOk() (*CloudFinanceCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CloudFinanceData) SetCost(v CloudFinanceCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CloudFinanceData) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDerived

`func (o *CloudFinanceData) GetDerived() CloudFinanceDerived`

GetDerived returns the Derived field if non-nil, zero value otherwise.

### GetDerivedOk

`func (o *CloudFinanceData) GetDerivedOk() (*CloudFinanceDerived, bool)`

GetDerivedOk returns a tuple with the Derived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerived

`func (o *CloudFinanceData) SetDerived(v CloudFinanceDerived)`

SetDerived sets Derived field to given value.

### HasDerived

`func (o *CloudFinanceData) HasDerived() bool`

HasDerived returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *CloudFinanceData) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *CloudFinanceData) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *CloudFinanceData) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *CloudFinanceData) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetRevenue

`func (o *CloudFinanceData) GetRevenue() CloudFinanceRevenue`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CloudFinanceData) GetRevenueOk() (*CloudFinanceRevenue, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CloudFinanceData) SetRevenue(v CloudFinanceRevenue)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *CloudFinanceData) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetSources

`func (o *CloudFinanceData) GetSources() []CloudSourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CloudFinanceData) GetSourcesOk() (*[]CloudSourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CloudFinanceData) SetSources(v []CloudSourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *CloudFinanceData) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



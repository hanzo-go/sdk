# FinanceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to [**FinanceCost**](FinanceCost.md) |  | [optional] 
**Derived** | Pointer to [**FinanceDerived**](FinanceDerived.md) |  | [optional] 
**GeneratedAt** | Pointer to **string** |  | [optional] 
**Revenue** | Pointer to [**FinanceRevenue**](FinanceRevenue.md) |  | [optional] 
**Sources** | Pointer to [**[]SourceStatus**](SourceStatus.md) |  | [optional] 

## Methods

### NewFinanceData

`func NewFinanceData() *FinanceData`

NewFinanceData instantiates a new FinanceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinanceDataWithDefaults

`func NewFinanceDataWithDefaults() *FinanceData`

NewFinanceDataWithDefaults instantiates a new FinanceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *FinanceData) GetCost() FinanceCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *FinanceData) GetCostOk() (*FinanceCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *FinanceData) SetCost(v FinanceCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *FinanceData) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDerived

`func (o *FinanceData) GetDerived() FinanceDerived`

GetDerived returns the Derived field if non-nil, zero value otherwise.

### GetDerivedOk

`func (o *FinanceData) GetDerivedOk() (*FinanceDerived, bool)`

GetDerivedOk returns a tuple with the Derived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerived

`func (o *FinanceData) SetDerived(v FinanceDerived)`

SetDerived sets Derived field to given value.

### HasDerived

`func (o *FinanceData) HasDerived() bool`

HasDerived returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *FinanceData) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *FinanceData) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *FinanceData) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *FinanceData) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetRevenue

`func (o *FinanceData) GetRevenue() FinanceRevenue`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *FinanceData) GetRevenueOk() (*FinanceRevenue, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *FinanceData) SetRevenue(v FinanceRevenue)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *FinanceData) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetSources

`func (o *FinanceData) GetSources() []SourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *FinanceData) GetSourcesOk() (*[]SourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *FinanceData) SetSources(v []SourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *FinanceData) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



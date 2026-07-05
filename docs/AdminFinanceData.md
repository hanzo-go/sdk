# AdminFinanceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to [**AdminFinanceCost**](AdminFinanceCost.md) |  | [optional] 
**Revenue** | Pointer to [**AdminFinanceRevenue**](AdminFinanceRevenue.md) |  | [optional] 
**Derived** | Pointer to [**AdminFinanceDerived**](AdminFinanceDerived.md) |  | [optional] 
**GeneratedAt** | Pointer to **string** |  | [optional] 
**Sources** | Pointer to [**[]AdminSourceStatus**](AdminSourceStatus.md) |  | [optional] 

## Methods

### NewAdminFinanceData

`func NewAdminFinanceData() *AdminFinanceData`

NewAdminFinanceData instantiates a new AdminFinanceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminFinanceDataWithDefaults

`func NewAdminFinanceDataWithDefaults() *AdminFinanceData`

NewAdminFinanceDataWithDefaults instantiates a new AdminFinanceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *AdminFinanceData) GetCost() AdminFinanceCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *AdminFinanceData) GetCostOk() (*AdminFinanceCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *AdminFinanceData) SetCost(v AdminFinanceCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *AdminFinanceData) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetRevenue

`func (o *AdminFinanceData) GetRevenue() AdminFinanceRevenue`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *AdminFinanceData) GetRevenueOk() (*AdminFinanceRevenue, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *AdminFinanceData) SetRevenue(v AdminFinanceRevenue)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *AdminFinanceData) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetDerived

`func (o *AdminFinanceData) GetDerived() AdminFinanceDerived`

GetDerived returns the Derived field if non-nil, zero value otherwise.

### GetDerivedOk

`func (o *AdminFinanceData) GetDerivedOk() (*AdminFinanceDerived, bool)`

GetDerivedOk returns a tuple with the Derived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerived

`func (o *AdminFinanceData) SetDerived(v AdminFinanceDerived)`

SetDerived sets Derived field to given value.

### HasDerived

`func (o *AdminFinanceData) HasDerived() bool`

HasDerived returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *AdminFinanceData) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *AdminFinanceData) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *AdminFinanceData) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *AdminFinanceData) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetSources

`func (o *AdminFinanceData) GetSources() []AdminSourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *AdminFinanceData) GetSourcesOk() (*[]AdminSourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *AdminFinanceData) SetSources(v []AdminSourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *AdminFinanceData) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



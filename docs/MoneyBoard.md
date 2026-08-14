# MoneyBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByOrg** | Pointer to [**[]MoneyOrgRow**](MoneyOrgRow.md) |  | [optional] 
**Credits** | Pointer to [**MoneyCredits**](MoneyCredits.md) |  | [optional] 
**GeneratedAt** | Pointer to **string** |  | [optional] 
**Infrastructure** | Pointer to [**MoneyInfra**](MoneyInfra.md) |  | [optional] 
**Margin** | Pointer to [**MoneyMargin**](MoneyMargin.md) |  | [optional] 
**Revenue** | Pointer to [**MoneyRevenue**](MoneyRevenue.md) |  | [optional] 
**Sources** | Pointer to [**[]SourceStatus**](SourceStatus.md) |  | [optional] 

## Methods

### NewMoneyBoard

`func NewMoneyBoard() *MoneyBoard`

NewMoneyBoard instantiates a new MoneyBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyBoardWithDefaults

`func NewMoneyBoardWithDefaults() *MoneyBoard`

NewMoneyBoardWithDefaults instantiates a new MoneyBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByOrg

`func (o *MoneyBoard) GetByOrg() []MoneyOrgRow`

GetByOrg returns the ByOrg field if non-nil, zero value otherwise.

### GetByOrgOk

`func (o *MoneyBoard) GetByOrgOk() (*[]MoneyOrgRow, bool)`

GetByOrgOk returns a tuple with the ByOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByOrg

`func (o *MoneyBoard) SetByOrg(v []MoneyOrgRow)`

SetByOrg sets ByOrg field to given value.

### HasByOrg

`func (o *MoneyBoard) HasByOrg() bool`

HasByOrg returns a boolean if a field has been set.

### GetCredits

`func (o *MoneyBoard) GetCredits() MoneyCredits`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *MoneyBoard) GetCreditsOk() (*MoneyCredits, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *MoneyBoard) SetCredits(v MoneyCredits)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *MoneyBoard) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *MoneyBoard) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *MoneyBoard) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *MoneyBoard) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *MoneyBoard) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetInfrastructure

`func (o *MoneyBoard) GetInfrastructure() MoneyInfra`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *MoneyBoard) GetInfrastructureOk() (*MoneyInfra, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *MoneyBoard) SetInfrastructure(v MoneyInfra)`

SetInfrastructure sets Infrastructure field to given value.

### HasInfrastructure

`func (o *MoneyBoard) HasInfrastructure() bool`

HasInfrastructure returns a boolean if a field has been set.

### GetMargin

`func (o *MoneyBoard) GetMargin() MoneyMargin`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *MoneyBoard) GetMarginOk() (*MoneyMargin, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *MoneyBoard) SetMargin(v MoneyMargin)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *MoneyBoard) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetRevenue

`func (o *MoneyBoard) GetRevenue() MoneyRevenue`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *MoneyBoard) GetRevenueOk() (*MoneyRevenue, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *MoneyBoard) SetRevenue(v MoneyRevenue)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *MoneyBoard) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetSources

`func (o *MoneyBoard) GetSources() []SourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *MoneyBoard) GetSourcesOk() (*[]SourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *MoneyBoard) SetSources(v []SourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *MoneyBoard) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



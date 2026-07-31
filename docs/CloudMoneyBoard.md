# CloudMoneyBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByOrg** | Pointer to [**[]CloudMoneyOrgRow**](CloudMoneyOrgRow.md) |  | [optional] 
**Credits** | Pointer to [**CloudMoneyCredits**](CloudMoneyCredits.md) |  | [optional] 
**GeneratedAt** | Pointer to **string** |  | [optional] 
**Infrastructure** | Pointer to [**CloudMoneyInfra**](CloudMoneyInfra.md) |  | [optional] 
**Margin** | Pointer to [**CloudMoneyMargin**](CloudMoneyMargin.md) |  | [optional] 
**Revenue** | Pointer to [**CloudMoneyRevenue**](CloudMoneyRevenue.md) |  | [optional] 
**Sources** | Pointer to [**[]CloudSourceStatus**](CloudSourceStatus.md) |  | [optional] 

## Methods

### NewCloudMoneyBoard

`func NewCloudMoneyBoard() *CloudMoneyBoard`

NewCloudMoneyBoard instantiates a new CloudMoneyBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMoneyBoardWithDefaults

`func NewCloudMoneyBoardWithDefaults() *CloudMoneyBoard`

NewCloudMoneyBoardWithDefaults instantiates a new CloudMoneyBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByOrg

`func (o *CloudMoneyBoard) GetByOrg() []CloudMoneyOrgRow`

GetByOrg returns the ByOrg field if non-nil, zero value otherwise.

### GetByOrgOk

`func (o *CloudMoneyBoard) GetByOrgOk() (*[]CloudMoneyOrgRow, bool)`

GetByOrgOk returns a tuple with the ByOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByOrg

`func (o *CloudMoneyBoard) SetByOrg(v []CloudMoneyOrgRow)`

SetByOrg sets ByOrg field to given value.

### HasByOrg

`func (o *CloudMoneyBoard) HasByOrg() bool`

HasByOrg returns a boolean if a field has been set.

### GetCredits

`func (o *CloudMoneyBoard) GetCredits() CloudMoneyCredits`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *CloudMoneyBoard) GetCreditsOk() (*CloudMoneyCredits, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *CloudMoneyBoard) SetCredits(v CloudMoneyCredits)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *CloudMoneyBoard) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *CloudMoneyBoard) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *CloudMoneyBoard) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *CloudMoneyBoard) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *CloudMoneyBoard) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetInfrastructure

`func (o *CloudMoneyBoard) GetInfrastructure() CloudMoneyInfra`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *CloudMoneyBoard) GetInfrastructureOk() (*CloudMoneyInfra, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *CloudMoneyBoard) SetInfrastructure(v CloudMoneyInfra)`

SetInfrastructure sets Infrastructure field to given value.

### HasInfrastructure

`func (o *CloudMoneyBoard) HasInfrastructure() bool`

HasInfrastructure returns a boolean if a field has been set.

### GetMargin

`func (o *CloudMoneyBoard) GetMargin() CloudMoneyMargin`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *CloudMoneyBoard) GetMarginOk() (*CloudMoneyMargin, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *CloudMoneyBoard) SetMargin(v CloudMoneyMargin)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *CloudMoneyBoard) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetRevenue

`func (o *CloudMoneyBoard) GetRevenue() CloudMoneyRevenue`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CloudMoneyBoard) GetRevenueOk() (*CloudMoneyRevenue, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CloudMoneyBoard) SetRevenue(v CloudMoneyRevenue)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *CloudMoneyBoard) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetSources

`func (o *CloudMoneyBoard) GetSources() []CloudSourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CloudMoneyBoard) GetSourcesOk() (*[]CloudSourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CloudMoneyBoard) SetSources(v []CloudSourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *CloudMoneyBoard) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# OverviewData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveProducts** | Pointer to **int32** |  | [optional] 
**CreditsCents** | Pointer to **int32** |  | [optional] 
**Drift** | Pointer to **int32** |  | [optional] 
**LastSync** | Pointer to **string** |  | [optional] 
**Orgs** | Pointer to **int32** |  | [optional] 
**Products** | Pointer to **int32** |  | [optional] 
**Sources** | Pointer to [**[]SourceStatus**](SourceStatus.md) |  | [optional] 
**SpendCents30d** | Pointer to **int32** |  | [optional] 
**Tokens30d** | Pointer to **int32** |  | [optional] 
**Users** | Pointer to **int32** |  | [optional] 

## Methods

### NewOverviewData

`func NewOverviewData() *OverviewData`

NewOverviewData instantiates a new OverviewData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewDataWithDefaults

`func NewOverviewDataWithDefaults() *OverviewData`

NewOverviewDataWithDefaults instantiates a new OverviewData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveProducts

`func (o *OverviewData) GetActiveProducts() int32`

GetActiveProducts returns the ActiveProducts field if non-nil, zero value otherwise.

### GetActiveProductsOk

`func (o *OverviewData) GetActiveProductsOk() (*int32, bool)`

GetActiveProductsOk returns a tuple with the ActiveProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProducts

`func (o *OverviewData) SetActiveProducts(v int32)`

SetActiveProducts sets ActiveProducts field to given value.

### HasActiveProducts

`func (o *OverviewData) HasActiveProducts() bool`

HasActiveProducts returns a boolean if a field has been set.

### GetCreditsCents

`func (o *OverviewData) GetCreditsCents() int32`

GetCreditsCents returns the CreditsCents field if non-nil, zero value otherwise.

### GetCreditsCentsOk

`func (o *OverviewData) GetCreditsCentsOk() (*int32, bool)`

GetCreditsCentsOk returns a tuple with the CreditsCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsCents

`func (o *OverviewData) SetCreditsCents(v int32)`

SetCreditsCents sets CreditsCents field to given value.

### HasCreditsCents

`func (o *OverviewData) HasCreditsCents() bool`

HasCreditsCents returns a boolean if a field has been set.

### GetDrift

`func (o *OverviewData) GetDrift() int32`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *OverviewData) GetDriftOk() (*int32, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *OverviewData) SetDrift(v int32)`

SetDrift sets Drift field to given value.

### HasDrift

`func (o *OverviewData) HasDrift() bool`

HasDrift returns a boolean if a field has been set.

### GetLastSync

`func (o *OverviewData) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *OverviewData) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *OverviewData) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *OverviewData) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetOrgs

`func (o *OverviewData) GetOrgs() int32`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *OverviewData) GetOrgsOk() (*int32, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *OverviewData) SetOrgs(v int32)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *OverviewData) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetProducts

`func (o *OverviewData) GetProducts() int32`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *OverviewData) GetProductsOk() (*int32, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *OverviewData) SetProducts(v int32)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *OverviewData) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetSources

`func (o *OverviewData) GetSources() []SourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *OverviewData) GetSourcesOk() (*[]SourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *OverviewData) SetSources(v []SourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *OverviewData) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetSpendCents30d

`func (o *OverviewData) GetSpendCents30d() int32`

GetSpendCents30d returns the SpendCents30d field if non-nil, zero value otherwise.

### GetSpendCents30dOk

`func (o *OverviewData) GetSpendCents30dOk() (*int32, bool)`

GetSpendCents30dOk returns a tuple with the SpendCents30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents30d

`func (o *OverviewData) SetSpendCents30d(v int32)`

SetSpendCents30d sets SpendCents30d field to given value.

### HasSpendCents30d

`func (o *OverviewData) HasSpendCents30d() bool`

HasSpendCents30d returns a boolean if a field has been set.

### GetTokens30d

`func (o *OverviewData) GetTokens30d() int32`

GetTokens30d returns the Tokens30d field if non-nil, zero value otherwise.

### GetTokens30dOk

`func (o *OverviewData) GetTokens30dOk() (*int32, bool)`

GetTokens30dOk returns a tuple with the Tokens30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens30d

`func (o *OverviewData) SetTokens30d(v int32)`

SetTokens30d sets Tokens30d field to given value.

### HasTokens30d

`func (o *OverviewData) HasTokens30d() bool`

HasTokens30d returns a boolean if a field has been set.

### GetUsers

`func (o *OverviewData) GetUsers() int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *OverviewData) GetUsersOk() (*int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *OverviewData) SetUsers(v int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *OverviewData) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



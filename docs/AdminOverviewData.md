# AdminOverviewData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orgs** | Pointer to **int32** |  | [optional] 
**Users** | Pointer to **int32** |  | [optional] 
**Products** | Pointer to **int32** |  | [optional] 
**ActiveProducts** | Pointer to **int32** |  | [optional] 
**Drift** | Pointer to **int32** |  | [optional] 
**SpendCents30d** | Pointer to **int64** |  | [optional] 
**Tokens30d** | Pointer to **int64** |  | [optional] 
**CreditsCents** | Pointer to **int64** |  | [optional] 
**LastSync** | Pointer to **string** |  | [optional] 
**Sources** | Pointer to [**[]AdminSourceStatus**](AdminSourceStatus.md) |  | [optional] 

## Methods

### NewAdminOverviewData

`func NewAdminOverviewData() *AdminOverviewData`

NewAdminOverviewData instantiates a new AdminOverviewData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminOverviewDataWithDefaults

`func NewAdminOverviewDataWithDefaults() *AdminOverviewData`

NewAdminOverviewDataWithDefaults instantiates a new AdminOverviewData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgs

`func (o *AdminOverviewData) GetOrgs() int32`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *AdminOverviewData) GetOrgsOk() (*int32, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *AdminOverviewData) SetOrgs(v int32)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *AdminOverviewData) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetUsers

`func (o *AdminOverviewData) GetUsers() int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AdminOverviewData) GetUsersOk() (*int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AdminOverviewData) SetUsers(v int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AdminOverviewData) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetProducts

`func (o *AdminOverviewData) GetProducts() int32`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *AdminOverviewData) GetProductsOk() (*int32, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *AdminOverviewData) SetProducts(v int32)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *AdminOverviewData) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetActiveProducts

`func (o *AdminOverviewData) GetActiveProducts() int32`

GetActiveProducts returns the ActiveProducts field if non-nil, zero value otherwise.

### GetActiveProductsOk

`func (o *AdminOverviewData) GetActiveProductsOk() (*int32, bool)`

GetActiveProductsOk returns a tuple with the ActiveProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProducts

`func (o *AdminOverviewData) SetActiveProducts(v int32)`

SetActiveProducts sets ActiveProducts field to given value.

### HasActiveProducts

`func (o *AdminOverviewData) HasActiveProducts() bool`

HasActiveProducts returns a boolean if a field has been set.

### GetDrift

`func (o *AdminOverviewData) GetDrift() int32`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *AdminOverviewData) GetDriftOk() (*int32, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *AdminOverviewData) SetDrift(v int32)`

SetDrift sets Drift field to given value.

### HasDrift

`func (o *AdminOverviewData) HasDrift() bool`

HasDrift returns a boolean if a field has been set.

### GetSpendCents30d

`func (o *AdminOverviewData) GetSpendCents30d() int64`

GetSpendCents30d returns the SpendCents30d field if non-nil, zero value otherwise.

### GetSpendCents30dOk

`func (o *AdminOverviewData) GetSpendCents30dOk() (*int64, bool)`

GetSpendCents30dOk returns a tuple with the SpendCents30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents30d

`func (o *AdminOverviewData) SetSpendCents30d(v int64)`

SetSpendCents30d sets SpendCents30d field to given value.

### HasSpendCents30d

`func (o *AdminOverviewData) HasSpendCents30d() bool`

HasSpendCents30d returns a boolean if a field has been set.

### GetTokens30d

`func (o *AdminOverviewData) GetTokens30d() int64`

GetTokens30d returns the Tokens30d field if non-nil, zero value otherwise.

### GetTokens30dOk

`func (o *AdminOverviewData) GetTokens30dOk() (*int64, bool)`

GetTokens30dOk returns a tuple with the Tokens30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens30d

`func (o *AdminOverviewData) SetTokens30d(v int64)`

SetTokens30d sets Tokens30d field to given value.

### HasTokens30d

`func (o *AdminOverviewData) HasTokens30d() bool`

HasTokens30d returns a boolean if a field has been set.

### GetCreditsCents

`func (o *AdminOverviewData) GetCreditsCents() int64`

GetCreditsCents returns the CreditsCents field if non-nil, zero value otherwise.

### GetCreditsCentsOk

`func (o *AdminOverviewData) GetCreditsCentsOk() (*int64, bool)`

GetCreditsCentsOk returns a tuple with the CreditsCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsCents

`func (o *AdminOverviewData) SetCreditsCents(v int64)`

SetCreditsCents sets CreditsCents field to given value.

### HasCreditsCents

`func (o *AdminOverviewData) HasCreditsCents() bool`

HasCreditsCents returns a boolean if a field has been set.

### GetLastSync

`func (o *AdminOverviewData) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *AdminOverviewData) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *AdminOverviewData) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *AdminOverviewData) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetSources

`func (o *AdminOverviewData) GetSources() []AdminSourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *AdminOverviewData) GetSourcesOk() (*[]AdminSourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *AdminOverviewData) SetSources(v []AdminSourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *AdminOverviewData) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



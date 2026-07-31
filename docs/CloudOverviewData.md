# CloudOverviewData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveProducts** | Pointer to **int32** |  | [optional] 
**CreditsCents** | Pointer to **int32** |  | [optional] 
**Drift** | Pointer to **int32** |  | [optional] 
**LastSync** | Pointer to **string** |  | [optional] 
**Orgs** | Pointer to **int32** |  | [optional] 
**Products** | Pointer to **int32** |  | [optional] 
**Sources** | Pointer to [**[]CloudSourceStatus**](CloudSourceStatus.md) |  | [optional] 
**SpendCents30d** | Pointer to **int32** |  | [optional] 
**Tokens30d** | Pointer to **int32** |  | [optional] 
**Users** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudOverviewData

`func NewCloudOverviewData() *CloudOverviewData`

NewCloudOverviewData instantiates a new CloudOverviewData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudOverviewDataWithDefaults

`func NewCloudOverviewDataWithDefaults() *CloudOverviewData`

NewCloudOverviewDataWithDefaults instantiates a new CloudOverviewData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveProducts

`func (o *CloudOverviewData) GetActiveProducts() int32`

GetActiveProducts returns the ActiveProducts field if non-nil, zero value otherwise.

### GetActiveProductsOk

`func (o *CloudOverviewData) GetActiveProductsOk() (*int32, bool)`

GetActiveProductsOk returns a tuple with the ActiveProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProducts

`func (o *CloudOverviewData) SetActiveProducts(v int32)`

SetActiveProducts sets ActiveProducts field to given value.

### HasActiveProducts

`func (o *CloudOverviewData) HasActiveProducts() bool`

HasActiveProducts returns a boolean if a field has been set.

### GetCreditsCents

`func (o *CloudOverviewData) GetCreditsCents() int32`

GetCreditsCents returns the CreditsCents field if non-nil, zero value otherwise.

### GetCreditsCentsOk

`func (o *CloudOverviewData) GetCreditsCentsOk() (*int32, bool)`

GetCreditsCentsOk returns a tuple with the CreditsCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsCents

`func (o *CloudOverviewData) SetCreditsCents(v int32)`

SetCreditsCents sets CreditsCents field to given value.

### HasCreditsCents

`func (o *CloudOverviewData) HasCreditsCents() bool`

HasCreditsCents returns a boolean if a field has been set.

### GetDrift

`func (o *CloudOverviewData) GetDrift() int32`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *CloudOverviewData) GetDriftOk() (*int32, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *CloudOverviewData) SetDrift(v int32)`

SetDrift sets Drift field to given value.

### HasDrift

`func (o *CloudOverviewData) HasDrift() bool`

HasDrift returns a boolean if a field has been set.

### GetLastSync

`func (o *CloudOverviewData) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *CloudOverviewData) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *CloudOverviewData) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *CloudOverviewData) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetOrgs

`func (o *CloudOverviewData) GetOrgs() int32`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *CloudOverviewData) GetOrgsOk() (*int32, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *CloudOverviewData) SetOrgs(v int32)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *CloudOverviewData) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetProducts

`func (o *CloudOverviewData) GetProducts() int32`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CloudOverviewData) GetProductsOk() (*int32, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CloudOverviewData) SetProducts(v int32)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *CloudOverviewData) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetSources

`func (o *CloudOverviewData) GetSources() []CloudSourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CloudOverviewData) GetSourcesOk() (*[]CloudSourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CloudOverviewData) SetSources(v []CloudSourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *CloudOverviewData) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetSpendCents30d

`func (o *CloudOverviewData) GetSpendCents30d() int32`

GetSpendCents30d returns the SpendCents30d field if non-nil, zero value otherwise.

### GetSpendCents30dOk

`func (o *CloudOverviewData) GetSpendCents30dOk() (*int32, bool)`

GetSpendCents30dOk returns a tuple with the SpendCents30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents30d

`func (o *CloudOverviewData) SetSpendCents30d(v int32)`

SetSpendCents30d sets SpendCents30d field to given value.

### HasSpendCents30d

`func (o *CloudOverviewData) HasSpendCents30d() bool`

HasSpendCents30d returns a boolean if a field has been set.

### GetTokens30d

`func (o *CloudOverviewData) GetTokens30d() int32`

GetTokens30d returns the Tokens30d field if non-nil, zero value otherwise.

### GetTokens30dOk

`func (o *CloudOverviewData) GetTokens30dOk() (*int32, bool)`

GetTokens30dOk returns a tuple with the Tokens30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens30d

`func (o *CloudOverviewData) SetTokens30d(v int32)`

SetTokens30d sets Tokens30d field to given value.

### HasTokens30d

`func (o *CloudOverviewData) HasTokens30d() bool`

HasTokens30d returns a boolean if a field has been set.

### GetUsers

`func (o *CloudOverviewData) GetUsers() int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *CloudOverviewData) GetUsersOk() (*int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *CloudOverviewData) SetUsers(v int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *CloudOverviewData) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



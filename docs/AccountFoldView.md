# AccountFoldView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**CloudAccountView**](CloudAccountView.md) | Account is the account as it is now recorded. | [optional] 
**Clusters** | Pointer to [**[]ClusterResult**](ClusterResult.md) | Clusters is one entry per cluster discovered in the account. It is empty when discovery itself failed, which leaves the previously folded set untouched rather than mass-detaching it. | [optional] 

## Methods

### NewAccountFoldView

`func NewAccountFoldView() *AccountFoldView`

NewAccountFoldView instantiates a new AccountFoldView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountFoldViewWithDefaults

`func NewAccountFoldViewWithDefaults() *AccountFoldView`

NewAccountFoldViewWithDefaults instantiates a new AccountFoldView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AccountFoldView) GetAccount() CloudAccountView`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AccountFoldView) GetAccountOk() (*CloudAccountView, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AccountFoldView) SetAccount(v CloudAccountView)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AccountFoldView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetClusters

`func (o *AccountFoldView) GetClusters() []ClusterResult`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *AccountFoldView) GetClustersOk() (*[]ClusterResult, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *AccountFoldView) SetClusters(v []ClusterResult)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *AccountFoldView) HasClusters() bool`

HasClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



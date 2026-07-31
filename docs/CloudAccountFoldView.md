# CloudAccountFoldView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**CloudCloudAccountView**](CloudCloudAccountView.md) | Account is the account as it is now recorded. | [optional] 
**Clusters** | Pointer to [**[]CloudClusterResult**](CloudClusterResult.md) | Clusters is one entry per cluster discovered in the account. It is empty when discovery itself failed, which leaves the previously folded set untouched rather than mass-detaching it. | [optional] 

## Methods

### NewCloudAccountFoldView

`func NewCloudAccountFoldView() *CloudAccountFoldView`

NewCloudAccountFoldView instantiates a new CloudAccountFoldView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAccountFoldViewWithDefaults

`func NewCloudAccountFoldViewWithDefaults() *CloudAccountFoldView`

NewCloudAccountFoldViewWithDefaults instantiates a new CloudAccountFoldView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudAccountFoldView) GetAccount() CloudCloudAccountView`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudAccountFoldView) GetAccountOk() (*CloudCloudAccountView, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudAccountFoldView) SetAccount(v CloudCloudAccountView)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudAccountFoldView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetClusters

`func (o *CloudAccountFoldView) GetClusters() []CloudClusterResult`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *CloudAccountFoldView) GetClustersOk() (*[]CloudClusterResult, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *CloudAccountFoldView) SetClusters(v []CloudClusterResult)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *CloudAccountFoldView) HasClusters() bool`

HasClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



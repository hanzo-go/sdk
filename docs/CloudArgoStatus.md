# CloudArgoStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to [**CloudArgoHealth**](CloudArgoHealth.md) |  | [optional] 
**ReconciledAt** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to [**[]CloudArgoResourceStatus**](CloudArgoResourceStatus.md) |  | [optional] 
**Summary** | Pointer to [**CloudArgoSummary**](CloudArgoSummary.md) |  | [optional] 
**Sync** | Pointer to [**CloudArgoSyncStatus**](CloudArgoSyncStatus.md) |  | [optional] 

## Methods

### NewCloudArgoStatus

`func NewCloudArgoStatus() *CloudArgoStatus`

NewCloudArgoStatus instantiates a new CloudArgoStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudArgoStatusWithDefaults

`func NewCloudArgoStatusWithDefaults() *CloudArgoStatus`

NewCloudArgoStatusWithDefaults instantiates a new CloudArgoStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *CloudArgoStatus) GetHealth() CloudArgoHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *CloudArgoStatus) GetHealthOk() (*CloudArgoHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *CloudArgoStatus) SetHealth(v CloudArgoHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *CloudArgoStatus) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetReconciledAt

`func (o *CloudArgoStatus) GetReconciledAt() string`

GetReconciledAt returns the ReconciledAt field if non-nil, zero value otherwise.

### GetReconciledAtOk

`func (o *CloudArgoStatus) GetReconciledAtOk() (*string, bool)`

GetReconciledAtOk returns a tuple with the ReconciledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciledAt

`func (o *CloudArgoStatus) SetReconciledAt(v string)`

SetReconciledAt sets ReconciledAt field to given value.

### HasReconciledAt

`func (o *CloudArgoStatus) HasReconciledAt() bool`

HasReconciledAt returns a boolean if a field has been set.

### GetResources

`func (o *CloudArgoStatus) GetResources() []CloudArgoResourceStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CloudArgoStatus) GetResourcesOk() (*[]CloudArgoResourceStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CloudArgoStatus) SetResources(v []CloudArgoResourceStatus)`

SetResources sets Resources field to given value.

### HasResources

`func (o *CloudArgoStatus) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSummary

`func (o *CloudArgoStatus) GetSummary() CloudArgoSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CloudArgoStatus) GetSummaryOk() (*CloudArgoSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CloudArgoStatus) SetSummary(v CloudArgoSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CloudArgoStatus) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetSync

`func (o *CloudArgoStatus) GetSync() CloudArgoSyncStatus`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *CloudArgoStatus) GetSyncOk() (*CloudArgoSyncStatus, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *CloudArgoStatus) SetSync(v CloudArgoSyncStatus)`

SetSync sets Sync field to given value.

### HasSync

`func (o *CloudArgoStatus) HasSync() bool`

HasSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



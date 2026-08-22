# ArgoStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to [**ArgoHealth**](ArgoHealth.md) | Health is the application&#39;s reconciled health. | [optional] 
**ReconciledAt** | Pointer to **string** | ReconciledAt is when the desired state was last compared against the cluster, RFC 3339. Empty for an App CR — the projection derives its verdict at read time and nothing records a comparison — and CD&#39;s own status.reconciledAt for a CD row. | [optional] 
**Resources** | Pointer to [**[]ArgoResourceStatus**](ArgoResourceStatus.md) | Resources are the objects the application owns. EMPTY on the list — filling it would walk the cluster once per row — and populated only by the read of ONE application, which is what makes that the detail view. | [optional] 
**Summary** | Pointer to [**ArgoSummary**](ArgoSummary.md) | Summary is the small aggregate the list column renders: the images. | [optional] 
**Sync** | Pointer to [**ArgoSyncStatus**](ArgoSyncStatus.md) | Sync is the declared-versus-running verdict and what it was reached against. | [optional] 

## Methods

### NewArgoStatus

`func NewArgoStatus() *ArgoStatus`

NewArgoStatus instantiates a new ArgoStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoStatusWithDefaults

`func NewArgoStatusWithDefaults() *ArgoStatus`

NewArgoStatusWithDefaults instantiates a new ArgoStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *ArgoStatus) GetHealth() ArgoHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ArgoStatus) GetHealthOk() (*ArgoHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ArgoStatus) SetHealth(v ArgoHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ArgoStatus) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetReconciledAt

`func (o *ArgoStatus) GetReconciledAt() string`

GetReconciledAt returns the ReconciledAt field if non-nil, zero value otherwise.

### GetReconciledAtOk

`func (o *ArgoStatus) GetReconciledAtOk() (*string, bool)`

GetReconciledAtOk returns a tuple with the ReconciledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciledAt

`func (o *ArgoStatus) SetReconciledAt(v string)`

SetReconciledAt sets ReconciledAt field to given value.

### HasReconciledAt

`func (o *ArgoStatus) HasReconciledAt() bool`

HasReconciledAt returns a boolean if a field has been set.

### GetResources

`func (o *ArgoStatus) GetResources() []ArgoResourceStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ArgoStatus) GetResourcesOk() (*[]ArgoResourceStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ArgoStatus) SetResources(v []ArgoResourceStatus)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ArgoStatus) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSummary

`func (o *ArgoStatus) GetSummary() ArgoSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ArgoStatus) GetSummaryOk() (*ArgoSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ArgoStatus) SetSummary(v ArgoSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ArgoStatus) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetSync

`func (o *ArgoStatus) GetSync() ArgoSyncStatus`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *ArgoStatus) GetSyncOk() (*ArgoSyncStatus, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *ArgoStatus) SetSync(v ArgoSyncStatus)`

SetSync sets Sync field to given value.

### HasSync

`func (o *ArgoStatus) HasSync() bool`

HasSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



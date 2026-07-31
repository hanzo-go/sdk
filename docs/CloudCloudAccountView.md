# CloudCloudAccountView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the provider&#39;s human label for it, e.g. the DigitalOcean team email. | [optional] 
**Clusters** | Pointer to **[]string** | Clusters is the fleet names this account currently owns. Unlinking detaches exactly these and nothing else. | [optional] 
**ExternalId** | Pointer to **string** | ExternalID is the provider&#39;s own identifier for the account — a DigitalOcean account uuid, an AWS account id, a GCP project. | [optional] 
**Label** | Pointer to **string** | Label is the org-chosen name for this account within the provider, which is how a second account at the same provider is addressed. Defaults to \&quot;default\&quot;. | [optional] 
**LinkedAt** | Pointer to **string** | LinkedAt is when the account was first linked, RFC3339 UTC. Re-linking the same label keeps the original. | [optional] 
**Project** | Pointer to **string** | Project is the fleet shard the account&#39;s clusters were folded into, recorded at link time so a later sync or unlink acts on the same shard. | [optional] 
**Provider** | Pointer to **string** | Provider is the cloud the account belongs to: digitalocean, aws, gcp or azure. | [optional] 
**SyncedAt** | Pointer to **string** | SyncedAt is when it was last discovered, RFC3339 UTC. | [optional] 

## Methods

### NewCloudCloudAccountView

`func NewCloudCloudAccountView() *CloudCloudAccountView`

NewCloudCloudAccountView instantiates a new CloudCloudAccountView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCloudAccountViewWithDefaults

`func NewCloudCloudAccountViewWithDefaults() *CloudCloudAccountView`

NewCloudCloudAccountViewWithDefaults instantiates a new CloudCloudAccountView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudCloudAccountView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudCloudAccountView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudCloudAccountView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudCloudAccountView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetClusters

`func (o *CloudCloudAccountView) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *CloudCloudAccountView) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *CloudCloudAccountView) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *CloudCloudAccountView) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetExternalId

`func (o *CloudCloudAccountView) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CloudCloudAccountView) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CloudCloudAccountView) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CloudCloudAccountView) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLabel

`func (o *CloudCloudAccountView) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CloudCloudAccountView) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CloudCloudAccountView) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CloudCloudAccountView) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLinkedAt

`func (o *CloudCloudAccountView) GetLinkedAt() string`

GetLinkedAt returns the LinkedAt field if non-nil, zero value otherwise.

### GetLinkedAtOk

`func (o *CloudCloudAccountView) GetLinkedAtOk() (*string, bool)`

GetLinkedAtOk returns a tuple with the LinkedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAt

`func (o *CloudCloudAccountView) SetLinkedAt(v string)`

SetLinkedAt sets LinkedAt field to given value.

### HasLinkedAt

`func (o *CloudCloudAccountView) HasLinkedAt() bool`

HasLinkedAt returns a boolean if a field has been set.

### GetProject

`func (o *CloudCloudAccountView) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudCloudAccountView) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudCloudAccountView) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudCloudAccountView) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProvider

`func (o *CloudCloudAccountView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudCloudAccountView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudCloudAccountView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudCloudAccountView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSyncedAt

`func (o *CloudCloudAccountView) GetSyncedAt() string`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *CloudCloudAccountView) GetSyncedAtOk() (*string, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *CloudCloudAccountView) SetSyncedAt(v string)`

SetSyncedAt sets SyncedAt field to given value.

### HasSyncedAt

`func (o *CloudCloudAccountView) HasSyncedAt() bool`

HasSyncedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



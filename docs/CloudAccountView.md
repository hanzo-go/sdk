# CloudAccountView

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

### NewCloudAccountView

`func NewCloudAccountView() *CloudAccountView`

NewCloudAccountView instantiates a new CloudAccountView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAccountViewWithDefaults

`func NewCloudAccountViewWithDefaults() *CloudAccountView`

NewCloudAccountViewWithDefaults instantiates a new CloudAccountView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudAccountView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudAccountView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudAccountView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudAccountView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetClusters

`func (o *CloudAccountView) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *CloudAccountView) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *CloudAccountView) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *CloudAccountView) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetExternalId

`func (o *CloudAccountView) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CloudAccountView) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CloudAccountView) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CloudAccountView) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLabel

`func (o *CloudAccountView) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CloudAccountView) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CloudAccountView) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CloudAccountView) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLinkedAt

`func (o *CloudAccountView) GetLinkedAt() string`

GetLinkedAt returns the LinkedAt field if non-nil, zero value otherwise.

### GetLinkedAtOk

`func (o *CloudAccountView) GetLinkedAtOk() (*string, bool)`

GetLinkedAtOk returns a tuple with the LinkedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAt

`func (o *CloudAccountView) SetLinkedAt(v string)`

SetLinkedAt sets LinkedAt field to given value.

### HasLinkedAt

`func (o *CloudAccountView) HasLinkedAt() bool`

HasLinkedAt returns a boolean if a field has been set.

### GetProject

`func (o *CloudAccountView) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudAccountView) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudAccountView) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudAccountView) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProvider

`func (o *CloudAccountView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudAccountView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudAccountView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudAccountView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSyncedAt

`func (o *CloudAccountView) GetSyncedAt() string`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *CloudAccountView) GetSyncedAtOk() (*string, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *CloudAccountView) SetSyncedAt(v string)`

SetSyncedAt sets SyncedAt field to given value.

### HasSyncedAt

`func (o *CloudAccountView) HasSyncedAt() bool`

HasSyncedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DeployArgoProjectSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterResourceWhitelist** | Pointer to [**[]DeployArgoGroupKind**](DeployArgoGroupKind.md) | ClusterResourceWhitelist are the cluster-scoped kinds it may create — [{group:\&quot;*\&quot;, kind:\&quot;*\&quot;}] on a synthesized project. | [optional] 
**Description** | Pointer to **string** | Description is the project&#39;s human label: the IAM project&#39;s display name, or its description when it has no display name. Absent when IAM carries neither. | [optional] 
**Destinations** | Pointer to [**[]DeployArgoDestination**](DeployArgoDestination.md) | Destinations are the cluster/namespace pairs it may write to — a single {server:\&quot;*\&quot;, namespace:\&quot;*\&quot;} on a synthesized project, for the same reason. | [optional] 
**SourceRepos** | Pointer to **[]string** | SourceRepos are the git repos applications in this project may pull from. [\&quot;*\&quot;] for every project this plane synthesizes or reflects from IAM: the boundary that actually holds on this platform is the IAM org, resolved before a row is ever projected, so the projected fence is deliberately permissive and is NOT an authorization statement. | [optional] 

## Methods

### NewDeployArgoProjectSpec

`func NewDeployArgoProjectSpec() *DeployArgoProjectSpec`

NewDeployArgoProjectSpec instantiates a new DeployArgoProjectSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployArgoProjectSpecWithDefaults

`func NewDeployArgoProjectSpecWithDefaults() *DeployArgoProjectSpec`

NewDeployArgoProjectSpecWithDefaults instantiates a new DeployArgoProjectSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterResourceWhitelist

`func (o *DeployArgoProjectSpec) GetClusterResourceWhitelist() []DeployArgoGroupKind`

GetClusterResourceWhitelist returns the ClusterResourceWhitelist field if non-nil, zero value otherwise.

### GetClusterResourceWhitelistOk

`func (o *DeployArgoProjectSpec) GetClusterResourceWhitelistOk() (*[]DeployArgoGroupKind, bool)`

GetClusterResourceWhitelistOk returns a tuple with the ClusterResourceWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterResourceWhitelist

`func (o *DeployArgoProjectSpec) SetClusterResourceWhitelist(v []DeployArgoGroupKind)`

SetClusterResourceWhitelist sets ClusterResourceWhitelist field to given value.

### HasClusterResourceWhitelist

`func (o *DeployArgoProjectSpec) HasClusterResourceWhitelist() bool`

HasClusterResourceWhitelist returns a boolean if a field has been set.

### GetDescription

`func (o *DeployArgoProjectSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeployArgoProjectSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeployArgoProjectSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeployArgoProjectSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestinations

`func (o *DeployArgoProjectSpec) GetDestinations() []DeployArgoDestination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *DeployArgoProjectSpec) GetDestinationsOk() (*[]DeployArgoDestination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *DeployArgoProjectSpec) SetDestinations(v []DeployArgoDestination)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *DeployArgoProjectSpec) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetSourceRepos

`func (o *DeployArgoProjectSpec) GetSourceRepos() []string`

GetSourceRepos returns the SourceRepos field if non-nil, zero value otherwise.

### GetSourceReposOk

`func (o *DeployArgoProjectSpec) GetSourceReposOk() (*[]string, bool)`

GetSourceReposOk returns a tuple with the SourceRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRepos

`func (o *DeployArgoProjectSpec) SetSourceRepos(v []string)`

SetSourceRepos sets SourceRepos field to given value.

### HasSourceRepos

`func (o *DeployArgoProjectSpec) HasSourceRepos() bool`

HasSourceRepos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



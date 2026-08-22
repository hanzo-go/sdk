# ArgoProjectSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterResourceWhitelist** | Pointer to [**[]ArgoGroupKind**](ArgoGroupKind.md) | ClusterResourceWhitelist are the cluster-scoped kinds it may create — [{group:\&quot;*\&quot;, kind:\&quot;*\&quot;}] on a synthesized project. | [optional] 
**Description** | Pointer to **string** | Description is the project&#39;s human label: the IAM project&#39;s display name, or its description when it has no display name. Absent when IAM carries neither. | [optional] 
**Destinations** | Pointer to [**[]ArgoDestination**](ArgoDestination.md) | Destinations are the cluster/namespace pairs it may write to — a single {server:\&quot;*\&quot;, namespace:\&quot;*\&quot;} on a synthesized project, for the same reason. | [optional] 
**SourceRepos** | Pointer to **[]string** | SourceRepos are the git repos applications in this project may pull from. [\&quot;*\&quot;] for every project this plane synthesizes or reflects from IAM: the boundary that actually holds on this platform is the IAM org, resolved before a row is ever projected, so the projected fence is deliberately permissive and is NOT an authorization statement. | [optional] 

## Methods

### NewArgoProjectSpec

`func NewArgoProjectSpec() *ArgoProjectSpec`

NewArgoProjectSpec instantiates a new ArgoProjectSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoProjectSpecWithDefaults

`func NewArgoProjectSpecWithDefaults() *ArgoProjectSpec`

NewArgoProjectSpecWithDefaults instantiates a new ArgoProjectSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterResourceWhitelist

`func (o *ArgoProjectSpec) GetClusterResourceWhitelist() []ArgoGroupKind`

GetClusterResourceWhitelist returns the ClusterResourceWhitelist field if non-nil, zero value otherwise.

### GetClusterResourceWhitelistOk

`func (o *ArgoProjectSpec) GetClusterResourceWhitelistOk() (*[]ArgoGroupKind, bool)`

GetClusterResourceWhitelistOk returns a tuple with the ClusterResourceWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterResourceWhitelist

`func (o *ArgoProjectSpec) SetClusterResourceWhitelist(v []ArgoGroupKind)`

SetClusterResourceWhitelist sets ClusterResourceWhitelist field to given value.

### HasClusterResourceWhitelist

`func (o *ArgoProjectSpec) HasClusterResourceWhitelist() bool`

HasClusterResourceWhitelist returns a boolean if a field has been set.

### GetDescription

`func (o *ArgoProjectSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArgoProjectSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArgoProjectSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArgoProjectSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestinations

`func (o *ArgoProjectSpec) GetDestinations() []ArgoDestination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *ArgoProjectSpec) GetDestinationsOk() (*[]ArgoDestination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *ArgoProjectSpec) SetDestinations(v []ArgoDestination)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *ArgoProjectSpec) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetSourceRepos

`func (o *ArgoProjectSpec) GetSourceRepos() []string`

GetSourceRepos returns the SourceRepos field if non-nil, zero value otherwise.

### GetSourceReposOk

`func (o *ArgoProjectSpec) GetSourceReposOk() (*[]string, bool)`

GetSourceReposOk returns a tuple with the SourceRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRepos

`func (o *ArgoProjectSpec) SetSourceRepos(v []string)`

SetSourceRepos sets SourceRepos field to given value.

### HasSourceRepos

`func (o *ArgoProjectSpec) HasSourceRepos() bool`

HasSourceRepos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



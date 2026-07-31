# CloudArgoProjectSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterResourceWhitelist** | Pointer to [**[]CloudArgoGroupKind**](CloudArgoGroupKind.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Destinations** | Pointer to [**[]CloudArgoDestination**](CloudArgoDestination.md) |  | [optional] 
**SourceRepos** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCloudArgoProjectSpec

`func NewCloudArgoProjectSpec() *CloudArgoProjectSpec`

NewCloudArgoProjectSpec instantiates a new CloudArgoProjectSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudArgoProjectSpecWithDefaults

`func NewCloudArgoProjectSpecWithDefaults() *CloudArgoProjectSpec`

NewCloudArgoProjectSpecWithDefaults instantiates a new CloudArgoProjectSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterResourceWhitelist

`func (o *CloudArgoProjectSpec) GetClusterResourceWhitelist() []CloudArgoGroupKind`

GetClusterResourceWhitelist returns the ClusterResourceWhitelist field if non-nil, zero value otherwise.

### GetClusterResourceWhitelistOk

`func (o *CloudArgoProjectSpec) GetClusterResourceWhitelistOk() (*[]CloudArgoGroupKind, bool)`

GetClusterResourceWhitelistOk returns a tuple with the ClusterResourceWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterResourceWhitelist

`func (o *CloudArgoProjectSpec) SetClusterResourceWhitelist(v []CloudArgoGroupKind)`

SetClusterResourceWhitelist sets ClusterResourceWhitelist field to given value.

### HasClusterResourceWhitelist

`func (o *CloudArgoProjectSpec) HasClusterResourceWhitelist() bool`

HasClusterResourceWhitelist returns a boolean if a field has been set.

### GetDescription

`func (o *CloudArgoProjectSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudArgoProjectSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudArgoProjectSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudArgoProjectSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestinations

`func (o *CloudArgoProjectSpec) GetDestinations() []CloudArgoDestination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *CloudArgoProjectSpec) GetDestinationsOk() (*[]CloudArgoDestination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *CloudArgoProjectSpec) SetDestinations(v []CloudArgoDestination)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *CloudArgoProjectSpec) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetSourceRepos

`func (o *CloudArgoProjectSpec) GetSourceRepos() []string`

GetSourceRepos returns the SourceRepos field if non-nil, zero value otherwise.

### GetSourceReposOk

`func (o *CloudArgoProjectSpec) GetSourceReposOk() (*[]string, bool)`

GetSourceReposOk returns a tuple with the SourceRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRepos

`func (o *CloudArgoProjectSpec) SetSourceRepos(v []string)`

SetSourceRepos sets SourceRepos field to given value.

### HasSourceRepos

`func (o *CloudArgoProjectSpec) HasSourceRepos() bool`

HasSourceRepos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



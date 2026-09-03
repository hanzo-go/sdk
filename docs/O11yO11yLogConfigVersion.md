# O11yO11yLogConfigVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **string** | Config is the rendered collector config the version deployed. | [optional] 
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when the version was created. | [optional] 
**CreatedBy** | Pointer to **string** | CreatedBy is the id of who created the version. | [optional] 
**CreatedByName** | Pointer to **string** | CreatedByName is the display name of who created the version. | [optional] 
**DeployResult** | Pointer to **string** | DeployResult is the deployment&#39;s outcome message. | [optional] 
**DeploySequence** | Pointer to **int64** | DeploySequence orders this deployment among the version&#39;s deployments. | [optional] 
**DeployStatus** | Pointer to **string** | DeployStatus is where the deployment stands, e.g. dirty, deploying, deployed, in_progress, failed, unknown. | [optional] 
**ElementType** | Pointer to **string** | ElementType is the config element the version carries — log_pipelines. | [optional] 
**Id** | Pointer to **string** | ID is the version record&#39;s id. | [optional] 
**LastHash** | Pointer to **string** | LastHash is the deployed config&#39;s hash. | [optional] 
**OrgId** | Pointer to **string** | OrgID is the org the version belongs to. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when the version last changed. | [optional] 
**UpdatedBy** | Pointer to **string** | UpdatedBy is the id of who last changed it. | [optional] 
**Version** | Pointer to **int64** | Version is the config version number. | [optional] 

## Methods

### NewO11yO11yLogConfigVersion

`func NewO11yO11yLogConfigVersion() *O11yO11yLogConfigVersion`

NewO11yO11yLogConfigVersion instantiates a new O11yO11yLogConfigVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLogConfigVersionWithDefaults

`func NewO11yO11yLogConfigVersionWithDefaults() *O11yO11yLogConfigVersion`

NewO11yO11yLogConfigVersionWithDefaults instantiates a new O11yO11yLogConfigVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *O11yO11yLogConfigVersion) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *O11yO11yLogConfigVersion) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *O11yO11yLogConfigVersion) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *O11yO11yLogConfigVersion) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11yO11yLogConfigVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yLogConfigVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yLogConfigVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yLogConfigVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *O11yO11yLogConfigVersion) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *O11yO11yLogConfigVersion) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *O11yO11yLogConfigVersion) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *O11yO11yLogConfigVersion) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByName

`func (o *O11yO11yLogConfigVersion) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *O11yO11yLogConfigVersion) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *O11yO11yLogConfigVersion) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *O11yO11yLogConfigVersion) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetDeployResult

`func (o *O11yO11yLogConfigVersion) GetDeployResult() string`

GetDeployResult returns the DeployResult field if non-nil, zero value otherwise.

### GetDeployResultOk

`func (o *O11yO11yLogConfigVersion) GetDeployResultOk() (*string, bool)`

GetDeployResultOk returns a tuple with the DeployResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployResult

`func (o *O11yO11yLogConfigVersion) SetDeployResult(v string)`

SetDeployResult sets DeployResult field to given value.

### HasDeployResult

`func (o *O11yO11yLogConfigVersion) HasDeployResult() bool`

HasDeployResult returns a boolean if a field has been set.

### GetDeploySequence

`func (o *O11yO11yLogConfigVersion) GetDeploySequence() int64`

GetDeploySequence returns the DeploySequence field if non-nil, zero value otherwise.

### GetDeploySequenceOk

`func (o *O11yO11yLogConfigVersion) GetDeploySequenceOk() (*int64, bool)`

GetDeploySequenceOk returns a tuple with the DeploySequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploySequence

`func (o *O11yO11yLogConfigVersion) SetDeploySequence(v int64)`

SetDeploySequence sets DeploySequence field to given value.

### HasDeploySequence

`func (o *O11yO11yLogConfigVersion) HasDeploySequence() bool`

HasDeploySequence returns a boolean if a field has been set.

### GetDeployStatus

`func (o *O11yO11yLogConfigVersion) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *O11yO11yLogConfigVersion) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *O11yO11yLogConfigVersion) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.

### HasDeployStatus

`func (o *O11yO11yLogConfigVersion) HasDeployStatus() bool`

HasDeployStatus returns a boolean if a field has been set.

### GetElementType

`func (o *O11yO11yLogConfigVersion) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *O11yO11yLogConfigVersion) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *O11yO11yLogConfigVersion) SetElementType(v string)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *O11yO11yLogConfigVersion) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yLogConfigVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yLogConfigVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yLogConfigVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yLogConfigVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastHash

`func (o *O11yO11yLogConfigVersion) GetLastHash() string`

GetLastHash returns the LastHash field if non-nil, zero value otherwise.

### GetLastHashOk

`func (o *O11yO11yLogConfigVersion) GetLastHashOk() (*string, bool)`

GetLastHashOk returns a tuple with the LastHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHash

`func (o *O11yO11yLogConfigVersion) SetLastHash(v string)`

SetLastHash sets LastHash field to given value.

### HasLastHash

`func (o *O11yO11yLogConfigVersion) HasLastHash() bool`

HasLastHash returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yO11yLogConfigVersion) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yO11yLogConfigVersion) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yO11yLogConfigVersion) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yO11yLogConfigVersion) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yLogConfigVersion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yLogConfigVersion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yLogConfigVersion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yLogConfigVersion) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *O11yO11yLogConfigVersion) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *O11yO11yLogConfigVersion) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *O11yO11yLogConfigVersion) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *O11yO11yLogConfigVersion) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetVersion

`func (o *O11yO11yLogConfigVersion) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *O11yO11yLogConfigVersion) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *O11yO11yLogConfigVersion) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *O11yO11yLogConfigVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



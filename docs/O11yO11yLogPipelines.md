# O11yO11yLogPipelines

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
**ElementType** | Pointer to **string** | ElementType is the config element this version carries — log_pipelines. | [optional] 
**History** | Pointer to [**[]O11yO11yLogConfigVersion**](O11yO11yLogConfigVersion.md) | History is the recent version history, newest first. | [optional] 
**Id** | Pointer to **string** | ID is the version record&#39;s id. | [optional] 
**LastHash** | Pointer to **string** | LastHash is the deployed config&#39;s hash. | [optional] 
**OrgId** | Pointer to **string** | OrgID is the org the version belongs to. | [optional] 
**Pipelines** | Pointer to [**[]O11yO11yLogPipeline**](O11yO11yLogPipeline.md) | Pipelines are the version&#39;s pipelines, in order. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when the version last changed. | [optional] 
**UpdatedBy** | Pointer to **string** | UpdatedBy is the id of who last changed it. | [optional] 
**Version** | Pointer to **int64** | Version is the config version number. | [optional] 

## Methods

### NewO11yO11yLogPipelines

`func NewO11yO11yLogPipelines() *O11yO11yLogPipelines`

NewO11yO11yLogPipelines instantiates a new O11yO11yLogPipelines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLogPipelinesWithDefaults

`func NewO11yO11yLogPipelinesWithDefaults() *O11yO11yLogPipelines`

NewO11yO11yLogPipelinesWithDefaults instantiates a new O11yO11yLogPipelines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *O11yO11yLogPipelines) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *O11yO11yLogPipelines) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *O11yO11yLogPipelines) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *O11yO11yLogPipelines) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11yO11yLogPipelines) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yLogPipelines) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yLogPipelines) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yLogPipelines) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *O11yO11yLogPipelines) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *O11yO11yLogPipelines) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *O11yO11yLogPipelines) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *O11yO11yLogPipelines) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByName

`func (o *O11yO11yLogPipelines) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *O11yO11yLogPipelines) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *O11yO11yLogPipelines) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *O11yO11yLogPipelines) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetDeployResult

`func (o *O11yO11yLogPipelines) GetDeployResult() string`

GetDeployResult returns the DeployResult field if non-nil, zero value otherwise.

### GetDeployResultOk

`func (o *O11yO11yLogPipelines) GetDeployResultOk() (*string, bool)`

GetDeployResultOk returns a tuple with the DeployResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployResult

`func (o *O11yO11yLogPipelines) SetDeployResult(v string)`

SetDeployResult sets DeployResult field to given value.

### HasDeployResult

`func (o *O11yO11yLogPipelines) HasDeployResult() bool`

HasDeployResult returns a boolean if a field has been set.

### GetDeploySequence

`func (o *O11yO11yLogPipelines) GetDeploySequence() int64`

GetDeploySequence returns the DeploySequence field if non-nil, zero value otherwise.

### GetDeploySequenceOk

`func (o *O11yO11yLogPipelines) GetDeploySequenceOk() (*int64, bool)`

GetDeploySequenceOk returns a tuple with the DeploySequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploySequence

`func (o *O11yO11yLogPipelines) SetDeploySequence(v int64)`

SetDeploySequence sets DeploySequence field to given value.

### HasDeploySequence

`func (o *O11yO11yLogPipelines) HasDeploySequence() bool`

HasDeploySequence returns a boolean if a field has been set.

### GetDeployStatus

`func (o *O11yO11yLogPipelines) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *O11yO11yLogPipelines) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *O11yO11yLogPipelines) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.

### HasDeployStatus

`func (o *O11yO11yLogPipelines) HasDeployStatus() bool`

HasDeployStatus returns a boolean if a field has been set.

### GetElementType

`func (o *O11yO11yLogPipelines) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *O11yO11yLogPipelines) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *O11yO11yLogPipelines) SetElementType(v string)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *O11yO11yLogPipelines) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetHistory

`func (o *O11yO11yLogPipelines) GetHistory() []O11yO11yLogConfigVersion`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *O11yO11yLogPipelines) GetHistoryOk() (*[]O11yO11yLogConfigVersion, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *O11yO11yLogPipelines) SetHistory(v []O11yO11yLogConfigVersion)`

SetHistory sets History field to given value.

### HasHistory

`func (o *O11yO11yLogPipelines) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yLogPipelines) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yLogPipelines) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yLogPipelines) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yLogPipelines) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastHash

`func (o *O11yO11yLogPipelines) GetLastHash() string`

GetLastHash returns the LastHash field if non-nil, zero value otherwise.

### GetLastHashOk

`func (o *O11yO11yLogPipelines) GetLastHashOk() (*string, bool)`

GetLastHashOk returns a tuple with the LastHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHash

`func (o *O11yO11yLogPipelines) SetLastHash(v string)`

SetLastHash sets LastHash field to given value.

### HasLastHash

`func (o *O11yO11yLogPipelines) HasLastHash() bool`

HasLastHash returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yO11yLogPipelines) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yO11yLogPipelines) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yO11yLogPipelines) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yO11yLogPipelines) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPipelines

`func (o *O11yO11yLogPipelines) GetPipelines() []O11yO11yLogPipeline`

GetPipelines returns the Pipelines field if non-nil, zero value otherwise.

### GetPipelinesOk

`func (o *O11yO11yLogPipelines) GetPipelinesOk() (*[]O11yO11yLogPipeline, bool)`

GetPipelinesOk returns a tuple with the Pipelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelines

`func (o *O11yO11yLogPipelines) SetPipelines(v []O11yO11yLogPipeline)`

SetPipelines sets Pipelines field to given value.

### HasPipelines

`func (o *O11yO11yLogPipelines) HasPipelines() bool`

HasPipelines returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yLogPipelines) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yLogPipelines) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yLogPipelines) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yLogPipelines) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *O11yO11yLogPipelines) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *O11yO11yLogPipelines) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *O11yO11yLogPipelines) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *O11yO11yLogPipelines) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetVersion

`func (o *O11yO11yLogPipelines) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *O11yO11yLogPipelines) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *O11yO11yLogPipelines) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *O11yO11yLogPipelines) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



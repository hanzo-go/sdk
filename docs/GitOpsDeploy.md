# GitOpsDeploy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Automated** | Pointer to **bool** | Automated is whether CD started this deploy itself, from its own polling of the tracked git ref (initiatedBy.automated), rather than someone asking for it. | [optional] 
**DeployedAt** | Pointer to **string** | DeployedAt is when the apply finished, RFC 3339. Absent when CD recorded none. | [optional] 
**Id** | Pointer to **int64** | ID is CD&#39;s own sequence number for this deploy (status.history[].id). It increases with every applied revision, so the largest id in &#x60;history&#x60; is the most recent deploy — which is the first entry, since the list is reversed. | [optional] 
**Revision** | Pointer to **string** | Revision is the git commit this deploy applied, as CD recorded it. | [optional] 
**StartedAt** | Pointer to **string** | StartedAt is when CD began applying the revision (deployStartedAt), RFC 3339. Absent when CD recorded none. | [optional] 

## Methods

### NewGitOpsDeploy

`func NewGitOpsDeploy() *GitOpsDeploy`

NewGitOpsDeploy instantiates a new GitOpsDeploy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitOpsDeployWithDefaults

`func NewGitOpsDeployWithDefaults() *GitOpsDeploy`

NewGitOpsDeployWithDefaults instantiates a new GitOpsDeploy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomated

`func (o *GitOpsDeploy) GetAutomated() bool`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *GitOpsDeploy) GetAutomatedOk() (*bool, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *GitOpsDeploy) SetAutomated(v bool)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *GitOpsDeploy) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetDeployedAt

`func (o *GitOpsDeploy) GetDeployedAt() string`

GetDeployedAt returns the DeployedAt field if non-nil, zero value otherwise.

### GetDeployedAtOk

`func (o *GitOpsDeploy) GetDeployedAtOk() (*string, bool)`

GetDeployedAtOk returns a tuple with the DeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAt

`func (o *GitOpsDeploy) SetDeployedAt(v string)`

SetDeployedAt sets DeployedAt field to given value.

### HasDeployedAt

`func (o *GitOpsDeploy) HasDeployedAt() bool`

HasDeployedAt returns a boolean if a field has been set.

### GetId

`func (o *GitOpsDeploy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitOpsDeploy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitOpsDeploy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GitOpsDeploy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRevision

`func (o *GitOpsDeploy) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *GitOpsDeploy) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *GitOpsDeploy) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *GitOpsDeploy) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetStartedAt

`func (o *GitOpsDeploy) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GitOpsDeploy) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GitOpsDeploy) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GitOpsDeploy) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



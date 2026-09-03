# DeploymentView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** | ApplicationID is the app this deployed — the app&#39;s &#x60;id&#x60;, not its slug. | [optional] 
**BuildId** | Pointer to **string** | BuildID is the build record behind a git deploy, whose logs and status live at /v1/platform/builds. Empty for an image deploy. | [optional] 
**Commit** | Pointer to **string** | Commit is the git ref this built — the commit a deploy or a push named, else the app&#39;s branch. Empty for an image deploy, which builds nothing. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the attempt was recorded, unix seconds. | [optional] 
**Id** | Pointer to **string** | ID is the deployment&#39;s id (&#x60;dep_…&#x60;), minted when the attempt is recorded. The app&#39;s currentDeploymentId points at one of these. | [optional] 
**Image** | Pointer to **string** | Image is the full &#x60;repo:tag&#x60; this deployment put in the CR. For a git deploy it is the ref the in-cluster build pushes to, known before the build runs. | [optional] 
**Message** | Pointer to **string** | Message is why this attempt is not live: the failure, or the note that a newer deployment went live before this build finished. Empty while it is fine. | [optional] 
**Org** | Pointer to **string** | Org is the tenant the deployment belongs to, from the validated identity. | [optional] 
**Source** | Pointer to **string** | Source is which lane produced it: &#x60;git&#x60; (built from the repo) or &#x60;image&#x60; (an already-built ref deployed as-is, including promote and rollback). | [optional] 
**Status** | Pointer to **string** | Status is where the attempt got to: &#x60;building&#x60; while its image is being built, &#x60;deploying&#x60; once its CR reached the cluster — which is the terminal success state, the app&#39;s own status is what turns &#x60;live&#x60; — &#x60;error&#x60; with the reason in Message, or &#x60;superseded&#x60; when a newer version went live first. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is its last transition, unix seconds — so for a terminal deployment it is when it reached that state. | [optional] 
**Version** | Pointer to **int64** | Version counts this app&#39;s deployments, from 1 and monotonically. It is what ORDERS them: a deploy only goes live if no higher version already is, so a build that finishes late is superseded instead of overwriting a newer one. | [optional] 

## Methods

### NewDeploymentView

`func NewDeploymentView() *DeploymentView`

NewDeploymentView instantiates a new DeploymentView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentViewWithDefaults

`func NewDeploymentViewWithDefaults() *DeploymentView`

NewDeploymentViewWithDefaults instantiates a new DeploymentView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *DeploymentView) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *DeploymentView) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *DeploymentView) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *DeploymentView) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetBuildId

`func (o *DeploymentView) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *DeploymentView) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *DeploymentView) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *DeploymentView) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### GetCommit

`func (o *DeploymentView) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *DeploymentView) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *DeploymentView) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *DeploymentView) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DeploymentView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DeploymentView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *DeploymentView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *DeploymentView) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *DeploymentView) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *DeploymentView) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *DeploymentView) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMessage

`func (o *DeploymentView) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeploymentView) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeploymentView) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeploymentView) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrg

`func (o *DeploymentView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *DeploymentView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *DeploymentView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *DeploymentView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSource

`func (o *DeploymentView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DeploymentView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DeploymentView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DeploymentView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DeploymentView) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentView) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentView) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeploymentView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *DeploymentView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeploymentView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeploymentView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeploymentView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



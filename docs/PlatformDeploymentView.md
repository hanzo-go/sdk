# PlatformDeploymentView

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

### NewPlatformDeploymentView

`func NewPlatformDeploymentView() *PlatformDeploymentView`

NewPlatformDeploymentView instantiates a new PlatformDeploymentView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformDeploymentViewWithDefaults

`func NewPlatformDeploymentViewWithDefaults() *PlatformDeploymentView`

NewPlatformDeploymentViewWithDefaults instantiates a new PlatformDeploymentView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *PlatformDeploymentView) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PlatformDeploymentView) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PlatformDeploymentView) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *PlatformDeploymentView) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetBuildId

`func (o *PlatformDeploymentView) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *PlatformDeploymentView) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *PlatformDeploymentView) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *PlatformDeploymentView) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### GetCommit

`func (o *PlatformDeploymentView) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *PlatformDeploymentView) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *PlatformDeploymentView) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *PlatformDeploymentView) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PlatformDeploymentView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlatformDeploymentView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlatformDeploymentView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PlatformDeploymentView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *PlatformDeploymentView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlatformDeploymentView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlatformDeploymentView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlatformDeploymentView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *PlatformDeploymentView) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PlatformDeploymentView) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PlatformDeploymentView) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PlatformDeploymentView) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMessage

`func (o *PlatformDeploymentView) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PlatformDeploymentView) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PlatformDeploymentView) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PlatformDeploymentView) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrg

`func (o *PlatformDeploymentView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PlatformDeploymentView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PlatformDeploymentView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PlatformDeploymentView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSource

`func (o *PlatformDeploymentView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PlatformDeploymentView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PlatformDeploymentView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PlatformDeploymentView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *PlatformDeploymentView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlatformDeploymentView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlatformDeploymentView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlatformDeploymentView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PlatformDeploymentView) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlatformDeploymentView) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlatformDeploymentView) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PlatformDeploymentView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *PlatformDeploymentView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PlatformDeploymentView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PlatformDeploymentView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PlatformDeploymentView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



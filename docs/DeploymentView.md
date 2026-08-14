# DeploymentView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** |  | [optional] 
**BuildId** | Pointer to **string** |  | [optional] 
**Commit** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

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

`func (o *DeploymentView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentView) SetCreatedAt(v int32)`

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

`func (o *DeploymentView) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentView) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentView) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeploymentView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *DeploymentView) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeploymentView) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeploymentView) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeploymentView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



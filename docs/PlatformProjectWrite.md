# PlatformProjectWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | Pointer to **string** | Commit is the sha the change created, \&quot;\&quot; when nothing had to move. | [optional] 
**Live** | Pointer to **bool** | Live says whether the generator will see it: true only on main. | [optional] 
**Mode** | Pointer to **string** | Mode is branch (a review; nothing deploys) or commit (main). | [optional] 
**Moved** | Pointer to [**[]PlatformRelabel**](PlatformRelabel.md) | Moved are the declarations whose partOf moved, and from what. | [optional] 
**Project** | Pointer to **string** | Project is the project the change leaves the moved apps in. | [optional] 
**Ref** | Pointer to **string** | Ref is the branch the change was pushed to, main for a commit. | [optional] 
**Review** | Pointer to **string** | Review is where to open the pull request for a branch. | [optional] 

## Methods

### NewPlatformProjectWrite

`func NewPlatformProjectWrite() *PlatformProjectWrite`

NewPlatformProjectWrite instantiates a new PlatformProjectWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformProjectWriteWithDefaults

`func NewPlatformProjectWriteWithDefaults() *PlatformProjectWrite`

NewPlatformProjectWriteWithDefaults instantiates a new PlatformProjectWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *PlatformProjectWrite) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *PlatformProjectWrite) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *PlatformProjectWrite) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *PlatformProjectWrite) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetLive

`func (o *PlatformProjectWrite) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *PlatformProjectWrite) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *PlatformProjectWrite) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *PlatformProjectWrite) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetMode

`func (o *PlatformProjectWrite) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PlatformProjectWrite) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PlatformProjectWrite) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PlatformProjectWrite) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMoved

`func (o *PlatformProjectWrite) GetMoved() []PlatformRelabel`

GetMoved returns the Moved field if non-nil, zero value otherwise.

### GetMovedOk

`func (o *PlatformProjectWrite) GetMovedOk() (*[]PlatformRelabel, bool)`

GetMovedOk returns a tuple with the Moved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoved

`func (o *PlatformProjectWrite) SetMoved(v []PlatformRelabel)`

SetMoved sets Moved field to given value.

### HasMoved

`func (o *PlatformProjectWrite) HasMoved() bool`

HasMoved returns a boolean if a field has been set.

### GetProject

`func (o *PlatformProjectWrite) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PlatformProjectWrite) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PlatformProjectWrite) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PlatformProjectWrite) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRef

`func (o *PlatformProjectWrite) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *PlatformProjectWrite) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *PlatformProjectWrite) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *PlatformProjectWrite) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetReview

`func (o *PlatformProjectWrite) GetReview() string`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *PlatformProjectWrite) GetReviewOk() (*string, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *PlatformProjectWrite) SetReview(v string)`

SetReview sets Review field to given value.

### HasReview

`func (o *PlatformProjectWrite) HasReview() bool`

HasReview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



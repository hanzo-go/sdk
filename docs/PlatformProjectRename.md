# PlatformProjectRename

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Mode is &#x60;branch&#x60; (the default) or &#x60;commit&#x60;. | [optional] 
**Name** | Pointer to **string** | Name is its new name, which no app may already name. | [optional] 
**Org** | Pointer to **string** | Org names the project&#39;s owner, defaulting to the caller&#39;s own scope. | [optional] 
**Project** | Pointer to **string** | Project is the project to rename, from the path. | [optional] 

## Methods

### NewPlatformProjectRename

`func NewPlatformProjectRename() *PlatformProjectRename`

NewPlatformProjectRename instantiates a new PlatformProjectRename object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformProjectRenameWithDefaults

`func NewPlatformProjectRenameWithDefaults() *PlatformProjectRename`

NewPlatformProjectRenameWithDefaults instantiates a new PlatformProjectRename object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *PlatformProjectRename) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PlatformProjectRename) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PlatformProjectRename) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PlatformProjectRename) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *PlatformProjectRename) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformProjectRename) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformProjectRename) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformProjectRename) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *PlatformProjectRename) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PlatformProjectRename) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PlatformProjectRename) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PlatformProjectRename) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *PlatformProjectRename) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PlatformProjectRename) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PlatformProjectRename) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PlatformProjectRename) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



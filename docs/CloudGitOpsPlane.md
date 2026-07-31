# CloudGitOpsPlane

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]CloudGitOpsApp**](CloudGitOpsApp.md) |  | [optional] 
**Installed** | Pointer to **bool** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudGitOpsPlane

`func NewCloudGitOpsPlane() *CloudGitOpsPlane`

NewCloudGitOpsPlane instantiates a new CloudGitOpsPlane object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGitOpsPlaneWithDefaults

`func NewCloudGitOpsPlaneWithDefaults() *CloudGitOpsPlane`

NewCloudGitOpsPlaneWithDefaults instantiates a new CloudGitOpsPlane object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *CloudGitOpsPlane) GetApplications() []CloudGitOpsApp`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *CloudGitOpsPlane) GetApplicationsOk() (*[]CloudGitOpsApp, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *CloudGitOpsPlane) SetApplications(v []CloudGitOpsApp)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *CloudGitOpsPlane) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetInstalled

`func (o *CloudGitOpsPlane) GetInstalled() bool`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *CloudGitOpsPlane) GetInstalledOk() (*bool, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *CloudGitOpsPlane) SetInstalled(v bool)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *CloudGitOpsPlane) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.

### GetReason

`func (o *CloudGitOpsPlane) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CloudGitOpsPlane) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CloudGitOpsPlane) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CloudGitOpsPlane) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



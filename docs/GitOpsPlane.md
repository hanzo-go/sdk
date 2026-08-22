# GitOpsPlane

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]GitOpsApp**](GitOpsApp.md) | Applications is every CD Application in the cluster, ordered by namespace then name. Empty (never null) when the plane is not installed, and equally empty when it is installed and tracks nothing — Installed is what separates those two. | [optional] 
**Installed** | Pointer to **bool** | Installed is whether this cluster serves the CD Application CRD at all. False is a fact about the cluster, not a failure of the request: the caller says \&quot;no CD plane here\&quot; rather than rendering an error it cannot act on. | [optional] 
**Reason** | Pointer to **string** | Reason says why the plane is absent, in words a caller can show. Empty when Installed. | [optional] 

## Methods

### NewGitOpsPlane

`func NewGitOpsPlane() *GitOpsPlane`

NewGitOpsPlane instantiates a new GitOpsPlane object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitOpsPlaneWithDefaults

`func NewGitOpsPlaneWithDefaults() *GitOpsPlane`

NewGitOpsPlaneWithDefaults instantiates a new GitOpsPlane object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *GitOpsPlane) GetApplications() []GitOpsApp`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *GitOpsPlane) GetApplicationsOk() (*[]GitOpsApp, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *GitOpsPlane) SetApplications(v []GitOpsApp)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *GitOpsPlane) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetInstalled

`func (o *GitOpsPlane) GetInstalled() bool`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *GitOpsPlane) GetInstalledOk() (*bool, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *GitOpsPlane) SetInstalled(v bool)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *GitOpsPlane) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.

### GetReason

`func (o *GitOpsPlane) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GitOpsPlane) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GitOpsPlane) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GitOpsPlane) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



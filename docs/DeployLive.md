# DeployLive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Installed** | Pointer to **bool** | Installed is whether this cluster serves the CD Application CRD at all. False is a fact about the cluster, not a failure of the request. | [optional] 
**Reason** | Pointer to **string** | Reason says why the plane is absent. Empty when Installed. | [optional] 
**Services** | Pointer to [**[]DeployRunning**](DeployRunning.md) | Services is every visible Application, ordered by namespace then name. Empty, never null. | [optional] 

## Methods

### NewDeployLive

`func NewDeployLive() *DeployLive`

NewDeployLive instantiates a new DeployLive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployLiveWithDefaults

`func NewDeployLiveWithDefaults() *DeployLive`

NewDeployLiveWithDefaults instantiates a new DeployLive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstalled

`func (o *DeployLive) GetInstalled() bool`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *DeployLive) GetInstalledOk() (*bool, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *DeployLive) SetInstalled(v bool)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *DeployLive) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.

### GetReason

`func (o *DeployLive) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DeployLive) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DeployLive) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DeployLive) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetServices

`func (o *DeployLive) GetServices() []DeployRunning`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *DeployLive) GetServicesOk() (*[]DeployRunning, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *DeployLive) SetServices(v []DeployRunning)`

SetServices sets Services field to given value.

### HasServices

`func (o *DeployLive) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



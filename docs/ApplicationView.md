# ApplicationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to [**[]EnvVariable**](EnvVariable.md) |  | [optional] 
**Deployments** | Pointer to [**[]DeploymentDetail**](DeploymentDetail.md) |  | [optional] 
**Events** | Pointer to [**[]ApplicationEvent**](ApplicationEvent.md) |  | [optional] 
**Metrics** | Pointer to [**ResourceMetrics**](ResourceMetrics.md) |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Services** | Pointer to [**[]ServiceDetail**](ServiceDetail.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewApplicationView

`func NewApplicationView() *ApplicationView`

NewApplicationView instantiates a new ApplicationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationViewWithDefaults

`func NewApplicationViewWithDefaults() *ApplicationView`

NewApplicationViewWithDefaults instantiates a new ApplicationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *ApplicationView) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ApplicationView) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ApplicationView) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ApplicationView) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCredentials

`func (o *ApplicationView) GetCredentials() []EnvVariable`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ApplicationView) GetCredentialsOk() (*[]EnvVariable, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ApplicationView) SetCredentials(v []EnvVariable)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ApplicationView) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetDeployments

`func (o *ApplicationView) GetDeployments() []DeploymentDetail`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *ApplicationView) GetDeploymentsOk() (*[]DeploymentDetail, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *ApplicationView) SetDeployments(v []DeploymentDetail)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *ApplicationView) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetEvents

`func (o *ApplicationView) GetEvents() []ApplicationEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ApplicationView) GetEventsOk() (*[]ApplicationEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ApplicationView) SetEvents(v []ApplicationEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ApplicationView) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetMetrics

`func (o *ApplicationView) GetMetrics() ResourceMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ApplicationView) GetMetricsOk() (*ResourceMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ApplicationView) SetMetrics(v ResourceMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *ApplicationView) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetNamespace

`func (o *ApplicationView) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ApplicationView) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ApplicationView) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ApplicationView) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetServices

`func (o *ApplicationView) GetServices() []ServiceDetail`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ApplicationView) GetServicesOk() (*[]ServiceDetail, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ApplicationView) SetServices(v []ServiceDetail)`

SetServices sets Services field to given value.

### HasServices

`func (o *ApplicationView) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetStatus

`func (o *ApplicationView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplicationView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplicationView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplicationView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ServiceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIP** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**ExternalHost** | Pointer to **string** |  | [optional] 
**ExternalIP** | Pointer to **string** |  | [optional] 
**InternalHost** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to [**[]ServicePort**](ServicePort.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceDetail

`func NewServiceDetail() *ServiceDetail`

NewServiceDetail instantiates a new ServiceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDetailWithDefaults

`func NewServiceDetailWithDefaults() *ServiceDetail`

NewServiceDetailWithDefaults instantiates a new ServiceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIP

`func (o *ServiceDetail) GetClusterIP() string`

GetClusterIP returns the ClusterIP field if non-nil, zero value otherwise.

### GetClusterIPOk

`func (o *ServiceDetail) GetClusterIPOk() (*string, bool)`

GetClusterIPOk returns a tuple with the ClusterIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIP

`func (o *ServiceDetail) SetClusterIP(v string)`

SetClusterIP sets ClusterIP field to given value.

### HasClusterIP

`func (o *ServiceDetail) HasClusterIP() bool`

HasClusterIP returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ServiceDetail) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ServiceDetail) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ServiceDetail) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ServiceDetail) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetExternalHost

`func (o *ServiceDetail) GetExternalHost() string`

GetExternalHost returns the ExternalHost field if non-nil, zero value otherwise.

### GetExternalHostOk

`func (o *ServiceDetail) GetExternalHostOk() (*string, bool)`

GetExternalHostOk returns a tuple with the ExternalHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHost

`func (o *ServiceDetail) SetExternalHost(v string)`

SetExternalHost sets ExternalHost field to given value.

### HasExternalHost

`func (o *ServiceDetail) HasExternalHost() bool`

HasExternalHost returns a boolean if a field has been set.

### GetExternalIP

`func (o *ServiceDetail) GetExternalIP() string`

GetExternalIP returns the ExternalIP field if non-nil, zero value otherwise.

### GetExternalIPOk

`func (o *ServiceDetail) GetExternalIPOk() (*string, bool)`

GetExternalIPOk returns a tuple with the ExternalIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIP

`func (o *ServiceDetail) SetExternalIP(v string)`

SetExternalIP sets ExternalIP field to given value.

### HasExternalIP

`func (o *ServiceDetail) HasExternalIP() bool`

HasExternalIP returns a boolean if a field has been set.

### GetInternalHost

`func (o *ServiceDetail) GetInternalHost() string`

GetInternalHost returns the InternalHost field if non-nil, zero value otherwise.

### GetInternalHostOk

`func (o *ServiceDetail) GetInternalHostOk() (*string, bool)`

GetInternalHostOk returns a tuple with the InternalHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHost

`func (o *ServiceDetail) SetInternalHost(v string)`

SetInternalHost sets InternalHost field to given value.

### HasInternalHost

`func (o *ServiceDetail) HasInternalHost() bool`

HasInternalHost returns a boolean if a field has been set.

### GetName

`func (o *ServiceDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPorts

`func (o *ServiceDetail) GetPorts() []ServicePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ServiceDetail) GetPortsOk() (*[]ServicePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ServiceDetail) SetPorts(v []ServicePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ServiceDetail) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetType

`func (o *ServiceDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceDetail) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



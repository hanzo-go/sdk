# ServiceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to [**[]ServiceView**](ServiceView.md) | Services is every registered service with its live waitlist mode. | [optional] 

## Methods

### NewServiceList

`func NewServiceList() *ServiceList`

NewServiceList instantiates a new ServiceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceListWithDefaults

`func NewServiceListWithDefaults() *ServiceList`

NewServiceListWithDefaults instantiates a new ServiceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *ServiceList) GetServices() []ServiceView`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ServiceList) GetServicesOk() (*[]ServiceView, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ServiceList) SetServices(v []ServiceView)`

SetServices sets Services field to given value.

### HasServices

`func (o *ServiceList) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ServiceOne

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to [**ServiceView**](ServiceView.md) | Service is the row as it stands after the write, live mode included. | [optional] 

## Methods

### NewServiceOne

`func NewServiceOne() *ServiceOne`

NewServiceOne instantiates a new ServiceOne object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceOneWithDefaults

`func NewServiceOneWithDefaults() *ServiceOne`

NewServiceOneWithDefaults instantiates a new ServiceOne object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ServiceOne) GetService() ServiceView`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceOne) GetServiceOk() (*ServiceView, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceOne) SetService(v ServiceView)`

SetService sets Service field to given value.

### HasService

`func (o *ServiceOne) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



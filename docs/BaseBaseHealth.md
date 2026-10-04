# BaseBaseHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | Service is \&quot;base\&quot; — which subsystem answered. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; when the subsystem is serving. | [optional] 

## Methods

### NewBaseBaseHealth

`func NewBaseBaseHealth() *BaseBaseHealth`

NewBaseBaseHealth instantiates a new BaseBaseHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseBaseHealthWithDefaults

`func NewBaseBaseHealthWithDefaults() *BaseBaseHealth`

NewBaseBaseHealthWithDefaults instantiates a new BaseBaseHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *BaseBaseHealth) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BaseBaseHealth) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BaseBaseHealth) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *BaseBaseHealth) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *BaseBaseHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseBaseHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseBaseHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseBaseHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



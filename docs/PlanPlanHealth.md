# PlanPlanHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | Service names the subsystem that answered. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; whenever this subsystem is mounted. | [optional] 

## Methods

### NewPlanPlanHealth

`func NewPlanPlanHealth() *PlanPlanHealth`

NewPlanPlanHealth instantiates a new PlanPlanHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanPlanHealthWithDefaults

`func NewPlanPlanHealthWithDefaults() *PlanPlanHealth`

NewPlanPlanHealthWithDefaults instantiates a new PlanPlanHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *PlanPlanHealth) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *PlanPlanHealth) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *PlanPlanHealth) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *PlanPlanHealth) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *PlanPlanHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlanPlanHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlanPlanHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlanPlanHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



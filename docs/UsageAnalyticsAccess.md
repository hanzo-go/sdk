# UsageAnalyticsAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**UsageAnalyticsGrant**](UsageAnalyticsGrant.md) | Access is what that plan grants. | [optional] 
**Plan** | Pointer to **string** | Plan echoes the plan id that was resolved, exactly as it was asked for. | [optional] 

## Methods

### NewUsageAnalyticsAccess

`func NewUsageAnalyticsAccess() *UsageAnalyticsAccess`

NewUsageAnalyticsAccess instantiates a new UsageAnalyticsAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageAnalyticsAccessWithDefaults

`func NewUsageAnalyticsAccessWithDefaults() *UsageAnalyticsAccess`

NewUsageAnalyticsAccessWithDefaults instantiates a new UsageAnalyticsAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *UsageAnalyticsAccess) GetAccess() UsageAnalyticsGrant`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UsageAnalyticsAccess) GetAccessOk() (*UsageAnalyticsGrant, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UsageAnalyticsAccess) SetAccess(v UsageAnalyticsGrant)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *UsageAnalyticsAccess) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetPlan

`func (o *UsageAnalyticsAccess) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UsageAnalyticsAccess) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UsageAnalyticsAccess) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UsageAnalyticsAccess) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



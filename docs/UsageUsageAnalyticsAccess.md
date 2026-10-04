# UsageUsageAnalyticsAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**UsageUsageAnalyticsGrant**](UsageUsageAnalyticsGrant.md) | Access is what that plan grants. | [optional] 
**Plan** | Pointer to **string** | Plan echoes the plan id that was resolved, exactly as it was asked for. | [optional] 

## Methods

### NewUsageUsageAnalyticsAccess

`func NewUsageUsageAnalyticsAccess() *UsageUsageAnalyticsAccess`

NewUsageUsageAnalyticsAccess instantiates a new UsageUsageAnalyticsAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageUsageAnalyticsAccessWithDefaults

`func NewUsageUsageAnalyticsAccessWithDefaults() *UsageUsageAnalyticsAccess`

NewUsageUsageAnalyticsAccessWithDefaults instantiates a new UsageUsageAnalyticsAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *UsageUsageAnalyticsAccess) GetAccess() UsageUsageAnalyticsGrant`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UsageUsageAnalyticsAccess) GetAccessOk() (*UsageUsageAnalyticsGrant, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UsageUsageAnalyticsAccess) SetAccess(v UsageUsageAnalyticsGrant)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *UsageUsageAnalyticsAccess) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetPlan

`func (o *UsageUsageAnalyticsAccess) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UsageUsageAnalyticsAccess) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UsageUsageAnalyticsAccess) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UsageUsageAnalyticsAccess) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



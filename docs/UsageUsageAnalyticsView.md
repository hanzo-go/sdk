# UsageUsageAnalyticsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** | End is the window&#39;s exclusive end, RFC3339 UTC. | [optional] 
**Export** | Pointer to **bool** | Export is whether the resolved plan allows exporting these rows. | [optional] 
**Plan** | Pointer to **string** | Plan echoes the plan id the entitlement was resolved from. | [optional] 
**Providers** | Pointer to [**UsageProviderBreakdown**](UsageProviderBreakdown.md) | Providers is the per-provider roll-up over the window. | [optional] 
**Range** | Pointer to **string** | Range is the label that was ASKED for. A plan whose retention is shorter than that window is served the retention instead, so read start and end for the window the rows actually cover and retentionDays for the reason — on a clamped read the label is longer than what was served. | [optional] 
**RetentionDays** | Pointer to **int64** | RetentionDays is how far back the resolved plan allows reading. | [optional] 
**Scope** | Pointer to [**UsageUsageScope**](UsageUsageScope.md) | Scope is the tenant the rows were read under — the validated principal&#39;s org. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s inclusive start, RFC3339 UTC, AFTER the retention clamp — so it may be later than the start that was asked for. | [optional] 

## Methods

### NewUsageUsageAnalyticsView

`func NewUsageUsageAnalyticsView() *UsageUsageAnalyticsView`

NewUsageUsageAnalyticsView instantiates a new UsageUsageAnalyticsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageUsageAnalyticsViewWithDefaults

`func NewUsageUsageAnalyticsViewWithDefaults() *UsageUsageAnalyticsView`

NewUsageUsageAnalyticsViewWithDefaults instantiates a new UsageUsageAnalyticsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *UsageUsageAnalyticsView) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *UsageUsageAnalyticsView) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *UsageUsageAnalyticsView) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *UsageUsageAnalyticsView) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetExport

`func (o *UsageUsageAnalyticsView) GetExport() bool`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *UsageUsageAnalyticsView) GetExportOk() (*bool, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *UsageUsageAnalyticsView) SetExport(v bool)`

SetExport sets Export field to given value.

### HasExport

`func (o *UsageUsageAnalyticsView) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetPlan

`func (o *UsageUsageAnalyticsView) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UsageUsageAnalyticsView) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UsageUsageAnalyticsView) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UsageUsageAnalyticsView) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProviders

`func (o *UsageUsageAnalyticsView) GetProviders() UsageProviderBreakdown`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *UsageUsageAnalyticsView) GetProvidersOk() (*UsageProviderBreakdown, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *UsageUsageAnalyticsView) SetProviders(v UsageProviderBreakdown)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *UsageUsageAnalyticsView) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetRange

`func (o *UsageUsageAnalyticsView) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *UsageUsageAnalyticsView) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *UsageUsageAnalyticsView) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *UsageUsageAnalyticsView) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRetentionDays

`func (o *UsageUsageAnalyticsView) GetRetentionDays() int64`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *UsageUsageAnalyticsView) GetRetentionDaysOk() (*int64, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *UsageUsageAnalyticsView) SetRetentionDays(v int64)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *UsageUsageAnalyticsView) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.

### GetScope

`func (o *UsageUsageAnalyticsView) GetScope() UsageUsageScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UsageUsageAnalyticsView) GetScopeOk() (*UsageUsageScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UsageUsageAnalyticsView) SetScope(v UsageUsageScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UsageUsageAnalyticsView) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStart

`func (o *UsageUsageAnalyticsView) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *UsageUsageAnalyticsView) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *UsageUsageAnalyticsView) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *UsageUsageAnalyticsView) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



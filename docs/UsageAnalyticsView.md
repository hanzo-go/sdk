# UsageAnalyticsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** | End is the window&#39;s exclusive end, RFC3339 UTC. | [optional] 
**Export** | Pointer to **bool** | Export is whether the resolved plan allows exporting these rows. | [optional] 
**Plan** | Pointer to **string** | Plan echoes the plan id the entitlement was resolved from. | [optional] 
**Providers** | Pointer to [**ProviderBreakdown**](ProviderBreakdown.md) | Providers is the per-provider roll-up over the window. | [optional] 
**Range** | Pointer to **string** | Range is the label that was ASKED for. A plan whose retention is shorter than that window is served the retention instead, so read start and end for the window the rows actually cover and retentionDays for the reason — on a clamped read the label is longer than what was served. | [optional] 
**RetentionDays** | Pointer to **int32** | RetentionDays is how far back the resolved plan allows reading. | [optional] 
**Scope** | Pointer to [**UsageScope**](UsageScope.md) | Scope is the tenant the rows were read under — the validated principal&#39;s org. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s inclusive start, RFC3339 UTC, AFTER the retention clamp — so it may be later than the start that was asked for. | [optional] 

## Methods

### NewUsageAnalyticsView

`func NewUsageAnalyticsView() *UsageAnalyticsView`

NewUsageAnalyticsView instantiates a new UsageAnalyticsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageAnalyticsViewWithDefaults

`func NewUsageAnalyticsViewWithDefaults() *UsageAnalyticsView`

NewUsageAnalyticsViewWithDefaults instantiates a new UsageAnalyticsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *UsageAnalyticsView) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *UsageAnalyticsView) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *UsageAnalyticsView) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *UsageAnalyticsView) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetExport

`func (o *UsageAnalyticsView) GetExport() bool`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *UsageAnalyticsView) GetExportOk() (*bool, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *UsageAnalyticsView) SetExport(v bool)`

SetExport sets Export field to given value.

### HasExport

`func (o *UsageAnalyticsView) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetPlan

`func (o *UsageAnalyticsView) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UsageAnalyticsView) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UsageAnalyticsView) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UsageAnalyticsView) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProviders

`func (o *UsageAnalyticsView) GetProviders() ProviderBreakdown`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *UsageAnalyticsView) GetProvidersOk() (*ProviderBreakdown, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *UsageAnalyticsView) SetProviders(v ProviderBreakdown)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *UsageAnalyticsView) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetRange

`func (o *UsageAnalyticsView) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *UsageAnalyticsView) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *UsageAnalyticsView) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *UsageAnalyticsView) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRetentionDays

`func (o *UsageAnalyticsView) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *UsageAnalyticsView) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *UsageAnalyticsView) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *UsageAnalyticsView) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.

### GetScope

`func (o *UsageAnalyticsView) GetScope() UsageScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UsageAnalyticsView) GetScopeOk() (*UsageScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UsageAnalyticsView) SetScope(v UsageScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UsageAnalyticsView) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStart

`func (o *UsageAnalyticsView) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *UsageAnalyticsView) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *UsageAnalyticsView) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *UsageAnalyticsView) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



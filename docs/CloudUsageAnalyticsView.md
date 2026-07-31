# CloudUsageAnalyticsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** | End is the window&#39;s exclusive end, RFC3339 UTC. | [optional] 
**Export** | Pointer to **bool** | Export is whether the resolved plan allows exporting these rows. | [optional] 
**Plan** | Pointer to **string** | Plan echoes the plan id the entitlement was resolved from. | [optional] 
**Providers** | Pointer to [**CloudProviderBreakdown**](CloudProviderBreakdown.md) | Providers is the per-provider roll-up over the window. | [optional] 
**Range** | Pointer to **string** | Range is the window label that was served, which is what was asked for. | [optional] 
**RetentionDays** | Pointer to **int32** | RetentionDays is how far back the resolved plan allows reading. | [optional] 
**Scope** | Pointer to [**CloudUsageScope**](CloudUsageScope.md) | Scope is the tenant the rows were read under — the validated principal&#39;s org. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s inclusive start, RFC3339 UTC, AFTER the retention clamp — so it may be later than the start that was asked for. | [optional] 

## Methods

### NewCloudUsageAnalyticsView

`func NewCloudUsageAnalyticsView() *CloudUsageAnalyticsView`

NewCloudUsageAnalyticsView instantiates a new CloudUsageAnalyticsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudUsageAnalyticsViewWithDefaults

`func NewCloudUsageAnalyticsViewWithDefaults() *CloudUsageAnalyticsView`

NewCloudUsageAnalyticsViewWithDefaults instantiates a new CloudUsageAnalyticsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *CloudUsageAnalyticsView) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CloudUsageAnalyticsView) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CloudUsageAnalyticsView) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CloudUsageAnalyticsView) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetExport

`func (o *CloudUsageAnalyticsView) GetExport() bool`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *CloudUsageAnalyticsView) GetExportOk() (*bool, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *CloudUsageAnalyticsView) SetExport(v bool)`

SetExport sets Export field to given value.

### HasExport

`func (o *CloudUsageAnalyticsView) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetPlan

`func (o *CloudUsageAnalyticsView) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CloudUsageAnalyticsView) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CloudUsageAnalyticsView) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CloudUsageAnalyticsView) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProviders

`func (o *CloudUsageAnalyticsView) GetProviders() CloudProviderBreakdown`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *CloudUsageAnalyticsView) GetProvidersOk() (*CloudProviderBreakdown, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *CloudUsageAnalyticsView) SetProviders(v CloudProviderBreakdown)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *CloudUsageAnalyticsView) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetRange

`func (o *CloudUsageAnalyticsView) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CloudUsageAnalyticsView) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CloudUsageAnalyticsView) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *CloudUsageAnalyticsView) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRetentionDays

`func (o *CloudUsageAnalyticsView) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *CloudUsageAnalyticsView) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *CloudUsageAnalyticsView) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *CloudUsageAnalyticsView) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.

### GetScope

`func (o *CloudUsageAnalyticsView) GetScope() CloudUsageScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CloudUsageAnalyticsView) GetScopeOk() (*CloudUsageScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CloudUsageAnalyticsView) SetScope(v CloudUsageScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CloudUsageAnalyticsView) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStart

`func (o *CloudUsageAnalyticsView) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CloudUsageAnalyticsView) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CloudUsageAnalyticsView) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *CloudUsageAnalyticsView) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



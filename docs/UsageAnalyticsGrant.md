# UsageAnalyticsGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | Pointer to **bool** | Datastore is whether the plan may read GET /v1/usage/analytics at all. The free floor is false, and that is what a catalog outage resolves to. | [optional] 
**Export** | Pointer to **bool** | Export is whether the plan may export the analytics it can read. | [optional] 
**RetentionDays** | Pointer to **int64** | RetentionDays is how far back the plan may read. GET /v1/usage/analytics clamps a custom window&#39;s start to this, so an older &#x60;start&#x60; returns the clamped window rather than an error. | [optional] 

## Methods

### NewUsageAnalyticsGrant

`func NewUsageAnalyticsGrant() *UsageAnalyticsGrant`

NewUsageAnalyticsGrant instantiates a new UsageAnalyticsGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageAnalyticsGrantWithDefaults

`func NewUsageAnalyticsGrantWithDefaults() *UsageAnalyticsGrant`

NewUsageAnalyticsGrantWithDefaults instantiates a new UsageAnalyticsGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *UsageAnalyticsGrant) GetDatastore() bool`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *UsageAnalyticsGrant) GetDatastoreOk() (*bool, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *UsageAnalyticsGrant) SetDatastore(v bool)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *UsageAnalyticsGrant) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetExport

`func (o *UsageAnalyticsGrant) GetExport() bool`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *UsageAnalyticsGrant) GetExportOk() (*bool, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *UsageAnalyticsGrant) SetExport(v bool)`

SetExport sets Export field to given value.

### HasExport

`func (o *UsageAnalyticsGrant) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetRetentionDays

`func (o *UsageAnalyticsGrant) GetRetentionDays() int64`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *UsageAnalyticsGrant) GetRetentionDaysOk() (*int64, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *UsageAnalyticsGrant) SetRetentionDays(v int64)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *UsageAnalyticsGrant) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



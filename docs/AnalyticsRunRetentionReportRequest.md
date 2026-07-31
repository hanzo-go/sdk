# AnalyticsRunRetentionReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebsiteId** | **string** |  | 
**DateRange** | [**AnalyticsDateRange**](AnalyticsDateRange.md) |  | 
**Timezone** | **string** |  | 

## Methods

### NewAnalyticsRunRetentionReportRequest

`func NewAnalyticsRunRetentionReportRequest(websiteId string, dateRange AnalyticsDateRange, timezone string, ) *AnalyticsRunRetentionReportRequest`

NewAnalyticsRunRetentionReportRequest instantiates a new AnalyticsRunRetentionReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsRunRetentionReportRequestWithDefaults

`func NewAnalyticsRunRetentionReportRequestWithDefaults() *AnalyticsRunRetentionReportRequest`

NewAnalyticsRunRetentionReportRequestWithDefaults instantiates a new AnalyticsRunRetentionReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebsiteId

`func (o *AnalyticsRunRetentionReportRequest) GetWebsiteId() string`

GetWebsiteId returns the WebsiteId field if non-nil, zero value otherwise.

### GetWebsiteIdOk

`func (o *AnalyticsRunRetentionReportRequest) GetWebsiteIdOk() (*string, bool)`

GetWebsiteIdOk returns a tuple with the WebsiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteId

`func (o *AnalyticsRunRetentionReportRequest) SetWebsiteId(v string)`

SetWebsiteId sets WebsiteId field to given value.


### GetDateRange

`func (o *AnalyticsRunRetentionReportRequest) GetDateRange() AnalyticsDateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *AnalyticsRunRetentionReportRequest) GetDateRangeOk() (*AnalyticsDateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *AnalyticsRunRetentionReportRequest) SetDateRange(v AnalyticsDateRange)`

SetDateRange sets DateRange field to given value.


### GetTimezone

`func (o *AnalyticsRunRetentionReportRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AnalyticsRunRetentionReportRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AnalyticsRunRetentionReportRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



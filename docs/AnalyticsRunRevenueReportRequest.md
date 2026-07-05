# AnalyticsRunRevenueReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebsiteId** | **string** |  | 
**DateRange** | [**AnalyticsDateRange**](AnalyticsDateRange.md) |  | 
**Currency** | **string** |  | 
**Timezone** | **string** |  | 

## Methods

### NewAnalyticsRunRevenueReportRequest

`func NewAnalyticsRunRevenueReportRequest(websiteId string, dateRange AnalyticsDateRange, currency string, timezone string, ) *AnalyticsRunRevenueReportRequest`

NewAnalyticsRunRevenueReportRequest instantiates a new AnalyticsRunRevenueReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsRunRevenueReportRequestWithDefaults

`func NewAnalyticsRunRevenueReportRequestWithDefaults() *AnalyticsRunRevenueReportRequest`

NewAnalyticsRunRevenueReportRequestWithDefaults instantiates a new AnalyticsRunRevenueReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebsiteId

`func (o *AnalyticsRunRevenueReportRequest) GetWebsiteId() string`

GetWebsiteId returns the WebsiteId field if non-nil, zero value otherwise.

### GetWebsiteIdOk

`func (o *AnalyticsRunRevenueReportRequest) GetWebsiteIdOk() (*string, bool)`

GetWebsiteIdOk returns a tuple with the WebsiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteId

`func (o *AnalyticsRunRevenueReportRequest) SetWebsiteId(v string)`

SetWebsiteId sets WebsiteId field to given value.


### GetDateRange

`func (o *AnalyticsRunRevenueReportRequest) GetDateRange() AnalyticsDateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *AnalyticsRunRevenueReportRequest) GetDateRangeOk() (*AnalyticsDateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *AnalyticsRunRevenueReportRequest) SetDateRange(v AnalyticsDateRange)`

SetDateRange sets DateRange field to given value.


### GetCurrency

`func (o *AnalyticsRunRevenueReportRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AnalyticsRunRevenueReportRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AnalyticsRunRevenueReportRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTimezone

`func (o *AnalyticsRunRevenueReportRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AnalyticsRunRevenueReportRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AnalyticsRunRevenueReportRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AnalyticsRunInsightsReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebsiteId** | **string** |  | 
**DateRange** | [**AnalyticsDateRange**](AnalyticsDateRange.md) |  | 
**Fields** | [**[]AnalyticsRunInsightsReportRequestFieldsInner**](AnalyticsRunInsightsReportRequestFieldsInner.md) |  | 
**Filters** | [**[]AnalyticsRunInsightsReportRequestFiltersInner**](AnalyticsRunInsightsReportRequestFiltersInner.md) |  | 

## Methods

### NewAnalyticsRunInsightsReportRequest

`func NewAnalyticsRunInsightsReportRequest(websiteId string, dateRange AnalyticsDateRange, fields []AnalyticsRunInsightsReportRequestFieldsInner, filters []AnalyticsRunInsightsReportRequestFiltersInner, ) *AnalyticsRunInsightsReportRequest`

NewAnalyticsRunInsightsReportRequest instantiates a new AnalyticsRunInsightsReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsRunInsightsReportRequestWithDefaults

`func NewAnalyticsRunInsightsReportRequestWithDefaults() *AnalyticsRunInsightsReportRequest`

NewAnalyticsRunInsightsReportRequestWithDefaults instantiates a new AnalyticsRunInsightsReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebsiteId

`func (o *AnalyticsRunInsightsReportRequest) GetWebsiteId() string`

GetWebsiteId returns the WebsiteId field if non-nil, zero value otherwise.

### GetWebsiteIdOk

`func (o *AnalyticsRunInsightsReportRequest) GetWebsiteIdOk() (*string, bool)`

GetWebsiteIdOk returns a tuple with the WebsiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteId

`func (o *AnalyticsRunInsightsReportRequest) SetWebsiteId(v string)`

SetWebsiteId sets WebsiteId field to given value.


### GetDateRange

`func (o *AnalyticsRunInsightsReportRequest) GetDateRange() AnalyticsDateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *AnalyticsRunInsightsReportRequest) GetDateRangeOk() (*AnalyticsDateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *AnalyticsRunInsightsReportRequest) SetDateRange(v AnalyticsDateRange)`

SetDateRange sets DateRange field to given value.


### GetFields

`func (o *AnalyticsRunInsightsReportRequest) GetFields() []AnalyticsRunInsightsReportRequestFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AnalyticsRunInsightsReportRequest) GetFieldsOk() (*[]AnalyticsRunInsightsReportRequestFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AnalyticsRunInsightsReportRequest) SetFields(v []AnalyticsRunInsightsReportRequestFieldsInner)`

SetFields sets Fields field to given value.


### GetFilters

`func (o *AnalyticsRunInsightsReportRequest) GetFilters() []AnalyticsRunInsightsReportRequestFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AnalyticsRunInsightsReportRequest) GetFiltersOk() (*[]AnalyticsRunInsightsReportRequestFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AnalyticsRunInsightsReportRequest) SetFilters(v []AnalyticsRunInsightsReportRequestFiltersInner)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



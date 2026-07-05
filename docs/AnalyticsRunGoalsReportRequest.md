# AnalyticsRunGoalsReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebsiteId** | **string** |  | 
**DateRange** | [**AnalyticsDateRange**](AnalyticsDateRange.md) |  | 
**Goals** | [**[]AnalyticsRunGoalsReportRequestGoalsInner**](AnalyticsRunGoalsReportRequestGoalsInner.md) |  | 

## Methods

### NewAnalyticsRunGoalsReportRequest

`func NewAnalyticsRunGoalsReportRequest(websiteId string, dateRange AnalyticsDateRange, goals []AnalyticsRunGoalsReportRequestGoalsInner, ) *AnalyticsRunGoalsReportRequest`

NewAnalyticsRunGoalsReportRequest instantiates a new AnalyticsRunGoalsReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsRunGoalsReportRequestWithDefaults

`func NewAnalyticsRunGoalsReportRequestWithDefaults() *AnalyticsRunGoalsReportRequest`

NewAnalyticsRunGoalsReportRequestWithDefaults instantiates a new AnalyticsRunGoalsReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebsiteId

`func (o *AnalyticsRunGoalsReportRequest) GetWebsiteId() string`

GetWebsiteId returns the WebsiteId field if non-nil, zero value otherwise.

### GetWebsiteIdOk

`func (o *AnalyticsRunGoalsReportRequest) GetWebsiteIdOk() (*string, bool)`

GetWebsiteIdOk returns a tuple with the WebsiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteId

`func (o *AnalyticsRunGoalsReportRequest) SetWebsiteId(v string)`

SetWebsiteId sets WebsiteId field to given value.


### GetDateRange

`func (o *AnalyticsRunGoalsReportRequest) GetDateRange() AnalyticsDateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *AnalyticsRunGoalsReportRequest) GetDateRangeOk() (*AnalyticsDateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *AnalyticsRunGoalsReportRequest) SetDateRange(v AnalyticsDateRange)`

SetDateRange sets DateRange field to given value.


### GetGoals

`func (o *AnalyticsRunGoalsReportRequest) GetGoals() []AnalyticsRunGoalsReportRequestGoalsInner`

GetGoals returns the Goals field if non-nil, zero value otherwise.

### GetGoalsOk

`func (o *AnalyticsRunGoalsReportRequest) GetGoalsOk() (*[]AnalyticsRunGoalsReportRequestGoalsInner, bool)`

GetGoalsOk returns a tuple with the Goals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoals

`func (o *AnalyticsRunGoalsReportRequest) SetGoals(v []AnalyticsRunGoalsReportRequestGoalsInner)`

SetGoals sets Goals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AnalyticsRunFunnelReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebsiteId** | **string** |  | 
**DateRange** | [**AnalyticsDateRange**](AnalyticsDateRange.md) |  | 
**Window** | **float32** | Funnel window in minutes | 
**Steps** | [**[]AnalyticsRunAttributionReportRequestStepsInner**](AnalyticsRunAttributionReportRequestStepsInner.md) |  | 

## Methods

### NewAnalyticsRunFunnelReportRequest

`func NewAnalyticsRunFunnelReportRequest(websiteId string, dateRange AnalyticsDateRange, window float32, steps []AnalyticsRunAttributionReportRequestStepsInner, ) *AnalyticsRunFunnelReportRequest`

NewAnalyticsRunFunnelReportRequest instantiates a new AnalyticsRunFunnelReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsRunFunnelReportRequestWithDefaults

`func NewAnalyticsRunFunnelReportRequestWithDefaults() *AnalyticsRunFunnelReportRequest`

NewAnalyticsRunFunnelReportRequestWithDefaults instantiates a new AnalyticsRunFunnelReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebsiteId

`func (o *AnalyticsRunFunnelReportRequest) GetWebsiteId() string`

GetWebsiteId returns the WebsiteId field if non-nil, zero value otherwise.

### GetWebsiteIdOk

`func (o *AnalyticsRunFunnelReportRequest) GetWebsiteIdOk() (*string, bool)`

GetWebsiteIdOk returns a tuple with the WebsiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteId

`func (o *AnalyticsRunFunnelReportRequest) SetWebsiteId(v string)`

SetWebsiteId sets WebsiteId field to given value.


### GetDateRange

`func (o *AnalyticsRunFunnelReportRequest) GetDateRange() AnalyticsDateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *AnalyticsRunFunnelReportRequest) GetDateRangeOk() (*AnalyticsDateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *AnalyticsRunFunnelReportRequest) SetDateRange(v AnalyticsDateRange)`

SetDateRange sets DateRange field to given value.


### GetWindow

`func (o *AnalyticsRunFunnelReportRequest) GetWindow() float32`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *AnalyticsRunFunnelReportRequest) GetWindowOk() (*float32, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *AnalyticsRunFunnelReportRequest) SetWindow(v float32)`

SetWindow sets Window field to given value.


### GetSteps

`func (o *AnalyticsRunFunnelReportRequest) GetSteps() []AnalyticsRunAttributionReportRequestStepsInner`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *AnalyticsRunFunnelReportRequest) GetStepsOk() (*[]AnalyticsRunAttributionReportRequestStepsInner, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *AnalyticsRunFunnelReportRequest) SetSteps(v []AnalyticsRunAttributionReportRequestStepsInner)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AnalyticsRunJourneyReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebsiteId** | **string** |  | 
**DateRange** | [**AnalyticsDateRange**](AnalyticsDateRange.md) |  | 
**Steps** | **int32** |  | 
**StartStep** | Pointer to **string** |  | [optional] 
**EndStep** | Pointer to **string** |  | [optional] 

## Methods

### NewAnalyticsRunJourneyReportRequest

`func NewAnalyticsRunJourneyReportRequest(websiteId string, dateRange AnalyticsDateRange, steps int32, ) *AnalyticsRunJourneyReportRequest`

NewAnalyticsRunJourneyReportRequest instantiates a new AnalyticsRunJourneyReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsRunJourneyReportRequestWithDefaults

`func NewAnalyticsRunJourneyReportRequestWithDefaults() *AnalyticsRunJourneyReportRequest`

NewAnalyticsRunJourneyReportRequestWithDefaults instantiates a new AnalyticsRunJourneyReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebsiteId

`func (o *AnalyticsRunJourneyReportRequest) GetWebsiteId() string`

GetWebsiteId returns the WebsiteId field if non-nil, zero value otherwise.

### GetWebsiteIdOk

`func (o *AnalyticsRunJourneyReportRequest) GetWebsiteIdOk() (*string, bool)`

GetWebsiteIdOk returns a tuple with the WebsiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteId

`func (o *AnalyticsRunJourneyReportRequest) SetWebsiteId(v string)`

SetWebsiteId sets WebsiteId field to given value.


### GetDateRange

`func (o *AnalyticsRunJourneyReportRequest) GetDateRange() AnalyticsDateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *AnalyticsRunJourneyReportRequest) GetDateRangeOk() (*AnalyticsDateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *AnalyticsRunJourneyReportRequest) SetDateRange(v AnalyticsDateRange)`

SetDateRange sets DateRange field to given value.


### GetSteps

`func (o *AnalyticsRunJourneyReportRequest) GetSteps() int32`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *AnalyticsRunJourneyReportRequest) GetStepsOk() (*int32, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *AnalyticsRunJourneyReportRequest) SetSteps(v int32)`

SetSteps sets Steps field to given value.


### GetStartStep

`func (o *AnalyticsRunJourneyReportRequest) GetStartStep() string`

GetStartStep returns the StartStep field if non-nil, zero value otherwise.

### GetStartStepOk

`func (o *AnalyticsRunJourneyReportRequest) GetStartStepOk() (*string, bool)`

GetStartStepOk returns a tuple with the StartStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartStep

`func (o *AnalyticsRunJourneyReportRequest) SetStartStep(v string)`

SetStartStep sets StartStep field to given value.

### HasStartStep

`func (o *AnalyticsRunJourneyReportRequest) HasStartStep() bool`

HasStartStep returns a boolean if a field has been set.

### GetEndStep

`func (o *AnalyticsRunJourneyReportRequest) GetEndStep() string`

GetEndStep returns the EndStep field if non-nil, zero value otherwise.

### GetEndStepOk

`func (o *AnalyticsRunJourneyReportRequest) GetEndStepOk() (*string, bool)`

GetEndStepOk returns a tuple with the EndStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndStep

`func (o *AnalyticsRunJourneyReportRequest) SetEndStep(v string)`

SetEndStep sets EndStep field to given value.

### HasEndStep

`func (o *AnalyticsRunJourneyReportRequest) HasEndStep() bool`

HasEndStep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AnalyticsRunAttributionReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebsiteId** | **string** |  | 
**DateRange** | [**AnalyticsDateRange**](AnalyticsDateRange.md) |  | 
**Model** | **string** |  | 
**Steps** | [**[]AnalyticsRunAttributionReportRequestStepsInner**](AnalyticsRunAttributionReportRequestStepsInner.md) |  | 
**Currency** | Pointer to **string** |  | [optional] 

## Methods

### NewAnalyticsRunAttributionReportRequest

`func NewAnalyticsRunAttributionReportRequest(websiteId string, dateRange AnalyticsDateRange, model string, steps []AnalyticsRunAttributionReportRequestStepsInner, ) *AnalyticsRunAttributionReportRequest`

NewAnalyticsRunAttributionReportRequest instantiates a new AnalyticsRunAttributionReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsRunAttributionReportRequestWithDefaults

`func NewAnalyticsRunAttributionReportRequestWithDefaults() *AnalyticsRunAttributionReportRequest`

NewAnalyticsRunAttributionReportRequestWithDefaults instantiates a new AnalyticsRunAttributionReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebsiteId

`func (o *AnalyticsRunAttributionReportRequest) GetWebsiteId() string`

GetWebsiteId returns the WebsiteId field if non-nil, zero value otherwise.

### GetWebsiteIdOk

`func (o *AnalyticsRunAttributionReportRequest) GetWebsiteIdOk() (*string, bool)`

GetWebsiteIdOk returns a tuple with the WebsiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteId

`func (o *AnalyticsRunAttributionReportRequest) SetWebsiteId(v string)`

SetWebsiteId sets WebsiteId field to given value.


### GetDateRange

`func (o *AnalyticsRunAttributionReportRequest) GetDateRange() AnalyticsDateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *AnalyticsRunAttributionReportRequest) GetDateRangeOk() (*AnalyticsDateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *AnalyticsRunAttributionReportRequest) SetDateRange(v AnalyticsDateRange)`

SetDateRange sets DateRange field to given value.


### GetModel

`func (o *AnalyticsRunAttributionReportRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AnalyticsRunAttributionReportRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AnalyticsRunAttributionReportRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetSteps

`func (o *AnalyticsRunAttributionReportRequest) GetSteps() []AnalyticsRunAttributionReportRequestStepsInner`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *AnalyticsRunAttributionReportRequest) GetStepsOk() (*[]AnalyticsRunAttributionReportRequestStepsInner, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *AnalyticsRunAttributionReportRequest) SetSteps(v []AnalyticsRunAttributionReportRequestStepsInner)`

SetSteps sets Steps field to given value.


### GetCurrency

`func (o *AnalyticsRunAttributionReportRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AnalyticsRunAttributionReportRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AnalyticsRunAttributionReportRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AnalyticsRunAttributionReportRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



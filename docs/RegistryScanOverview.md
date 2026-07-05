# RegistryScanOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportId** | Pointer to **string** |  | [optional] 
**ScanStatus** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**CompletePercent** | Pointer to **int32** |  | [optional] 
**Summary** | Pointer to [**RegistryScanOverviewSummary**](RegistryScanOverviewSummary.md) |  | [optional] 

## Methods

### NewRegistryScanOverview

`func NewRegistryScanOverview() *RegistryScanOverview`

NewRegistryScanOverview instantiates a new RegistryScanOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryScanOverviewWithDefaults

`func NewRegistryScanOverviewWithDefaults() *RegistryScanOverview`

NewRegistryScanOverviewWithDefaults instantiates a new RegistryScanOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportId

`func (o *RegistryScanOverview) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *RegistryScanOverview) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *RegistryScanOverview) SetReportId(v string)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *RegistryScanOverview) HasReportId() bool`

HasReportId returns a boolean if a field has been set.

### GetScanStatus

`func (o *RegistryScanOverview) GetScanStatus() string`

GetScanStatus returns the ScanStatus field if non-nil, zero value otherwise.

### GetScanStatusOk

`func (o *RegistryScanOverview) GetScanStatusOk() (*string, bool)`

GetScanStatusOk returns a tuple with the ScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatus

`func (o *RegistryScanOverview) SetScanStatus(v string)`

SetScanStatus sets ScanStatus field to given value.

### HasScanStatus

`func (o *RegistryScanOverview) HasScanStatus() bool`

HasScanStatus returns a boolean if a field has been set.

### GetSeverity

`func (o *RegistryScanOverview) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RegistryScanOverview) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RegistryScanOverview) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *RegistryScanOverview) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStartTime

`func (o *RegistryScanOverview) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RegistryScanOverview) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RegistryScanOverview) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RegistryScanOverview) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *RegistryScanOverview) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RegistryScanOverview) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *RegistryScanOverview) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *RegistryScanOverview) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetCompletePercent

`func (o *RegistryScanOverview) GetCompletePercent() int32`

GetCompletePercent returns the CompletePercent field if non-nil, zero value otherwise.

### GetCompletePercentOk

`func (o *RegistryScanOverview) GetCompletePercentOk() (*int32, bool)`

GetCompletePercentOk returns a tuple with the CompletePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletePercent

`func (o *RegistryScanOverview) SetCompletePercent(v int32)`

SetCompletePercent sets CompletePercent field to given value.

### HasCompletePercent

`func (o *RegistryScanOverview) HasCompletePercent() bool`

HasCompletePercent returns a boolean if a field has been set.

### GetSummary

`func (o *RegistryScanOverview) GetSummary() RegistryScanOverviewSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RegistryScanOverview) GetSummaryOk() (*RegistryScanOverviewSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RegistryScanOverview) SetSummary(v RegistryScanOverviewSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *RegistryScanOverview) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



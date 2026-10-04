# DataroomDataroomPageStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgDuration** | Pointer to **int64** | AvgDuration is totalDuration divided by views, rounded; 0 when unviewed. | [optional] 
**PageNumber** | Pointer to **int64** | PageNumber is the page these counts are for. | [optional] 
**TotalDuration** | Pointer to **int64** | TotalDuration is the summed dwell measure reported for the page. | [optional] 
**Views** | Pointer to **int64** | Views is how many times the page was viewed. | [optional] 

## Methods

### NewDataroomDataroomPageStat

`func NewDataroomDataroomPageStat() *DataroomDataroomPageStat`

NewDataroomDataroomPageStat instantiates a new DataroomDataroomPageStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomDataroomPageStatWithDefaults

`func NewDataroomDataroomPageStatWithDefaults() *DataroomDataroomPageStat`

NewDataroomDataroomPageStatWithDefaults instantiates a new DataroomDataroomPageStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgDuration

`func (o *DataroomDataroomPageStat) GetAvgDuration() int64`

GetAvgDuration returns the AvgDuration field if non-nil, zero value otherwise.

### GetAvgDurationOk

`func (o *DataroomDataroomPageStat) GetAvgDurationOk() (*int64, bool)`

GetAvgDurationOk returns a tuple with the AvgDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDuration

`func (o *DataroomDataroomPageStat) SetAvgDuration(v int64)`

SetAvgDuration sets AvgDuration field to given value.

### HasAvgDuration

`func (o *DataroomDataroomPageStat) HasAvgDuration() bool`

HasAvgDuration returns a boolean if a field has been set.

### GetPageNumber

`func (o *DataroomDataroomPageStat) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DataroomDataroomPageStat) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DataroomDataroomPageStat) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DataroomDataroomPageStat) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetTotalDuration

`func (o *DataroomDataroomPageStat) GetTotalDuration() int64`

GetTotalDuration returns the TotalDuration field if non-nil, zero value otherwise.

### GetTotalDurationOk

`func (o *DataroomDataroomPageStat) GetTotalDurationOk() (*int64, bool)`

GetTotalDurationOk returns a tuple with the TotalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDuration

`func (o *DataroomDataroomPageStat) SetTotalDuration(v int64)`

SetTotalDuration sets TotalDuration field to given value.

### HasTotalDuration

`func (o *DataroomDataroomPageStat) HasTotalDuration() bool`

HasTotalDuration returns a boolean if a field has been set.

### GetViews

`func (o *DataroomDataroomPageStat) GetViews() int64`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *DataroomDataroomPageStat) GetViewsOk() (*int64, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *DataroomDataroomPageStat) SetViews(v int64)`

SetViews sets Views field to given value.

### HasViews

`func (o *DataroomDataroomPageStat) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



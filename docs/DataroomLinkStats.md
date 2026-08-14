# DataroomLinkStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkId** | Pointer to **string** | LinkId is the link these counts are for. | [optional] 
**Pages** | Pointer to [**[]DataroomPageStat**](DataroomPageStat.md) | Pages is the per-page breakdown, in page order. | [optional] 
**TotalPageViews** | Pointer to **int32** | TotalPageViews is how many page views the link received. | [optional] 
**TotalViews** | Pointer to **int32** | TotalViews is how many viewing sessions the link opened. | [optional] 

## Methods

### NewDataroomLinkStats

`func NewDataroomLinkStats() *DataroomLinkStats`

NewDataroomLinkStats instantiates a new DataroomLinkStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomLinkStatsWithDefaults

`func NewDataroomLinkStatsWithDefaults() *DataroomLinkStats`

NewDataroomLinkStatsWithDefaults instantiates a new DataroomLinkStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkId

`func (o *DataroomLinkStats) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *DataroomLinkStats) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *DataroomLinkStats) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *DataroomLinkStats) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetPages

`func (o *DataroomLinkStats) GetPages() []DataroomPageStat`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *DataroomLinkStats) GetPagesOk() (*[]DataroomPageStat, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *DataroomLinkStats) SetPages(v []DataroomPageStat)`

SetPages sets Pages field to given value.

### HasPages

`func (o *DataroomLinkStats) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetTotalPageViews

`func (o *DataroomLinkStats) GetTotalPageViews() int32`

GetTotalPageViews returns the TotalPageViews field if non-nil, zero value otherwise.

### GetTotalPageViewsOk

`func (o *DataroomLinkStats) GetTotalPageViewsOk() (*int32, bool)`

GetTotalPageViewsOk returns a tuple with the TotalPageViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPageViews

`func (o *DataroomLinkStats) SetTotalPageViews(v int32)`

SetTotalPageViews sets TotalPageViews field to given value.

### HasTotalPageViews

`func (o *DataroomLinkStats) HasTotalPageViews() bool`

HasTotalPageViews returns a boolean if a field has been set.

### GetTotalViews

`func (o *DataroomLinkStats) GetTotalViews() int32`

GetTotalViews returns the TotalViews field if non-nil, zero value otherwise.

### GetTotalViewsOk

`func (o *DataroomLinkStats) GetTotalViewsOk() (*int32, bool)`

GetTotalViewsOk returns a tuple with the TotalViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalViews

`func (o *DataroomLinkStats) SetTotalViews(v int32)`

SetTotalViews sets TotalViews field to given value.

### HasTotalViews

`func (o *DataroomLinkStats) HasTotalViews() bool`

HasTotalViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



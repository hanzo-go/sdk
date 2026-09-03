# DataroomStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataroomId** | Pointer to **string** | DataroomId is the room these counts are for. | [optional] 
**Links** | Pointer to [**[]DataroomLinkStats**](DataroomLinkStats.md) | Links is the same per-page breakdown for each link into the room. | [optional] 
**TotalPageViews** | Pointer to **int64** | TotalPageViews is the room&#39;s page views across every link. | [optional] 
**TotalViews** | Pointer to **int64** | TotalViews is the room&#39;s viewing sessions across every link. | [optional] 

## Methods

### NewDataroomStats

`func NewDataroomStats() *DataroomStats`

NewDataroomStats instantiates a new DataroomStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomStatsWithDefaults

`func NewDataroomStatsWithDefaults() *DataroomStats`

NewDataroomStatsWithDefaults instantiates a new DataroomStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataroomId

`func (o *DataroomStats) GetDataroomId() string`

GetDataroomId returns the DataroomId field if non-nil, zero value otherwise.

### GetDataroomIdOk

`func (o *DataroomStats) GetDataroomIdOk() (*string, bool)`

GetDataroomIdOk returns a tuple with the DataroomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataroomId

`func (o *DataroomStats) SetDataroomId(v string)`

SetDataroomId sets DataroomId field to given value.

### HasDataroomId

`func (o *DataroomStats) HasDataroomId() bool`

HasDataroomId returns a boolean if a field has been set.

### GetLinks

`func (o *DataroomStats) GetLinks() []DataroomLinkStats`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DataroomStats) GetLinksOk() (*[]DataroomLinkStats, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DataroomStats) SetLinks(v []DataroomLinkStats)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DataroomStats) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTotalPageViews

`func (o *DataroomStats) GetTotalPageViews() int64`

GetTotalPageViews returns the TotalPageViews field if non-nil, zero value otherwise.

### GetTotalPageViewsOk

`func (o *DataroomStats) GetTotalPageViewsOk() (*int64, bool)`

GetTotalPageViewsOk returns a tuple with the TotalPageViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPageViews

`func (o *DataroomStats) SetTotalPageViews(v int64)`

SetTotalPageViews sets TotalPageViews field to given value.

### HasTotalPageViews

`func (o *DataroomStats) HasTotalPageViews() bool`

HasTotalPageViews returns a boolean if a field has been set.

### GetTotalViews

`func (o *DataroomStats) GetTotalViews() int64`

GetTotalViews returns the TotalViews field if non-nil, zero value otherwise.

### GetTotalViewsOk

`func (o *DataroomStats) GetTotalViewsOk() (*int64, bool)`

GetTotalViewsOk returns a tuple with the TotalViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalViews

`func (o *DataroomStats) SetTotalViews(v int64)`

SetTotalViews sets TotalViews field to given value.

### HasTotalViews

`func (o *DataroomStats) HasTotalViews() bool`

HasTotalViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



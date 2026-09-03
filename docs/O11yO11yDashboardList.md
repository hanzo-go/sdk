# O11yO11yDashboardList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dashboards** | Pointer to [**[]O11yO11yDashboardListItem**](O11yO11yDashboardListItem.md) | Dashboards are the rows for this page. | [optional] 
**Tags** | Pointer to [**[]O11yO11yDashboardTag**](O11yO11yDashboardTag.md) | Tags are all tags in use across the org&#39;s dashboards. | [optional] 
**Total** | Pointer to **int64** | Total is the count across all pages. | [optional] 

## Methods

### NewO11yO11yDashboardList

`func NewO11yO11yDashboardList() *O11yO11yDashboardList`

NewO11yO11yDashboardList instantiates a new O11yO11yDashboardList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDashboardListWithDefaults

`func NewO11yO11yDashboardListWithDefaults() *O11yO11yDashboardList`

NewO11yO11yDashboardListWithDefaults instantiates a new O11yO11yDashboardList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboards

`func (o *O11yO11yDashboardList) GetDashboards() []O11yO11yDashboardListItem`

GetDashboards returns the Dashboards field if non-nil, zero value otherwise.

### GetDashboardsOk

`func (o *O11yO11yDashboardList) GetDashboardsOk() (*[]O11yO11yDashboardListItem, bool)`

GetDashboardsOk returns a tuple with the Dashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboards

`func (o *O11yO11yDashboardList) SetDashboards(v []O11yO11yDashboardListItem)`

SetDashboards sets Dashboards field to given value.

### HasDashboards

`func (o *O11yO11yDashboardList) HasDashboards() bool`

HasDashboards returns a boolean if a field has been set.

### GetTags

`func (o *O11yO11yDashboardList) GetTags() []O11yO11yDashboardTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yO11yDashboardList) GetTagsOk() (*[]O11yO11yDashboardTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yO11yDashboardList) SetTags(v []O11yO11yDashboardTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yO11yDashboardList) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTotal

`func (o *O11yO11yDashboardList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yO11yDashboardList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yO11yDashboardList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yO11yDashboardList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



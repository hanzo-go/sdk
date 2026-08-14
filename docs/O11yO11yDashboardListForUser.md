# O11yO11yDashboardListForUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dashboards** | Pointer to [**[]O11yO11yDashboardListItemForUser**](O11yO11yDashboardListItemForUser.md) | Dashboards are the rows for this page, each with the caller&#39;s pin state. | [optional] 
**Tags** | Pointer to [**[]O11yO11yDashboardTag**](O11yO11yDashboardTag.md) | Tags are all tags in use across the org&#39;s dashboards. | [optional] 
**Total** | Pointer to **int32** | Total is the count across all pages. | [optional] 

## Methods

### NewO11yO11yDashboardListForUser

`func NewO11yO11yDashboardListForUser() *O11yO11yDashboardListForUser`

NewO11yO11yDashboardListForUser instantiates a new O11yO11yDashboardListForUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDashboardListForUserWithDefaults

`func NewO11yO11yDashboardListForUserWithDefaults() *O11yO11yDashboardListForUser`

NewO11yO11yDashboardListForUserWithDefaults instantiates a new O11yO11yDashboardListForUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboards

`func (o *O11yO11yDashboardListForUser) GetDashboards() []O11yO11yDashboardListItemForUser`

GetDashboards returns the Dashboards field if non-nil, zero value otherwise.

### GetDashboardsOk

`func (o *O11yO11yDashboardListForUser) GetDashboardsOk() (*[]O11yO11yDashboardListItemForUser, bool)`

GetDashboardsOk returns a tuple with the Dashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboards

`func (o *O11yO11yDashboardListForUser) SetDashboards(v []O11yO11yDashboardListItemForUser)`

SetDashboards sets Dashboards field to given value.

### HasDashboards

`func (o *O11yO11yDashboardListForUser) HasDashboards() bool`

HasDashboards returns a boolean if a field has been set.

### GetTags

`func (o *O11yO11yDashboardListForUser) GetTags() []O11yO11yDashboardTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yO11yDashboardListForUser) GetTagsOk() (*[]O11yO11yDashboardTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yO11yDashboardListForUser) SetTags(v []O11yO11yDashboardTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yO11yDashboardListForUser) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTotal

`func (o *O11yO11yDashboardListForUser) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yO11yDashboardListForUser) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yO11yDashboardListForUser) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yO11yDashboardListForUser) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



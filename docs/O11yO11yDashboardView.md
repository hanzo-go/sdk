# O11yO11yDashboardView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when the view was created. | [optional] 
**Data** | Pointer to [**O11yO11yDashboardViewData**](O11yO11yDashboardViewData.md) | Data is the listing state the view captures. | [optional] 
**Id** | Pointer to **string** | ID is the saved view&#39;s id. | [optional] 
**Name** | Pointer to **string** | Name is the saved view&#39;s name. | [optional] 
**OrgId** | Pointer to **string** | OrgID is the org the view belongs to. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when it last changed. | [optional] 

## Methods

### NewO11yO11yDashboardView

`func NewO11yO11yDashboardView() *O11yO11yDashboardView`

NewO11yO11yDashboardView instantiates a new O11yO11yDashboardView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDashboardViewWithDefaults

`func NewO11yO11yDashboardViewWithDefaults() *O11yO11yDashboardView`

NewO11yO11yDashboardViewWithDefaults instantiates a new O11yO11yDashboardView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yO11yDashboardView) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yDashboardView) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yDashboardView) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yDashboardView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetData

`func (o *O11yO11yDashboardView) GetData() O11yO11yDashboardViewData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yDashboardView) GetDataOk() (*O11yO11yDashboardViewData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yDashboardView) SetData(v O11yO11yDashboardViewData)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yDashboardView) HasData() bool`

HasData returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yDashboardView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yDashboardView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yDashboardView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yDashboardView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yDashboardView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yDashboardView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yDashboardView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yDashboardView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yO11yDashboardView) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yO11yDashboardView) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yO11yDashboardView) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yO11yDashboardView) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yDashboardView) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yDashboardView) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yDashboardView) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yDashboardView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# O11yO11yDashboardViewPostable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yO11yDashboardViewData**](O11yO11yDashboardViewData.md) | Data is the listing state the view captures. | [optional] 
**Name** | Pointer to **string** | Name is the saved view&#39;s name; at most 32 characters, no surrounding space. | [optional] 

## Methods

### NewO11yO11yDashboardViewPostable

`func NewO11yO11yDashboardViewPostable() *O11yO11yDashboardViewPostable`

NewO11yO11yDashboardViewPostable instantiates a new O11yO11yDashboardViewPostable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDashboardViewPostableWithDefaults

`func NewO11yO11yDashboardViewPostableWithDefaults() *O11yO11yDashboardViewPostable`

NewO11yO11yDashboardViewPostableWithDefaults instantiates a new O11yO11yDashboardViewPostable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yDashboardViewPostable) GetData() O11yO11yDashboardViewData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yDashboardViewPostable) GetDataOk() (*O11yO11yDashboardViewData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yDashboardViewPostable) SetData(v O11yO11yDashboardViewData)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yDashboardViewPostable) HasData() bool`

HasData returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yDashboardViewPostable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yDashboardViewPostable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yDashboardViewPostable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yDashboardViewPostable) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



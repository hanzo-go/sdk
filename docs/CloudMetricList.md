# CloudMetricList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudMetricRow**](CloudMetricRow.md) | Data is one row per prompt the org owns. | [optional] 

## Methods

### NewCloudMetricList

`func NewCloudMetricList() *CloudMetricList`

NewCloudMetricList instantiates a new CloudMetricList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMetricListWithDefaults

`func NewCloudMetricListWithDefaults() *CloudMetricList`

NewCloudMetricListWithDefaults instantiates a new CloudMetricList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudMetricList) GetData() []CloudMetricRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudMetricList) GetDataOk() (*[]CloudMetricRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudMetricList) SetData(v []CloudMetricRow)`

SetData sets Data field to given value.

### HasData

`func (o *CloudMetricList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



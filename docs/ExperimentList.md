# ExperimentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Trial**](Trial.md) | Data is the org&#39;s experiments, ordered by project then id. | [optional] 
**Total** | Pointer to **int64** | Total is how many rows Data holds. | [optional] 

## Methods

### NewExperimentList

`func NewExperimentList() *ExperimentList`

NewExperimentList instantiates a new ExperimentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentListWithDefaults

`func NewExperimentListWithDefaults() *ExperimentList`

NewExperimentListWithDefaults instantiates a new ExperimentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ExperimentList) GetData() []Trial`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExperimentList) GetDataOk() (*[]Trial, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExperimentList) SetData(v []Trial)`

SetData sets Data field to given value.

### HasData

`func (o *ExperimentList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *ExperimentList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ExperimentList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ExperimentList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ExperimentList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DatasetList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]DatasetView**](DatasetView.md) | Data is the caller org&#39;s datasets, newest first, bounded by limit. | [optional] 

## Methods

### NewDatasetList

`func NewDatasetList() *DatasetList`

NewDatasetList instantiates a new DatasetList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetListWithDefaults

`func NewDatasetListWithDefaults() *DatasetList`

NewDatasetListWithDefaults instantiates a new DatasetList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DatasetList) GetData() []DatasetView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DatasetList) GetDataOk() (*[]DatasetView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DatasetList) SetData(v []DatasetView)`

SetData sets Data field to given value.

### HasData

`func (o *DatasetList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



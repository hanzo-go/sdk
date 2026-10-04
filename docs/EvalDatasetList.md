# EvalDatasetList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]EvalDatasetView**](EvalDatasetView.md) | Data is the caller org&#39;s datasets, newest first, bounded by limit. | [optional] 

## Methods

### NewEvalDatasetList

`func NewEvalDatasetList() *EvalDatasetList`

NewEvalDatasetList instantiates a new EvalDatasetList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalDatasetListWithDefaults

`func NewEvalDatasetListWithDefaults() *EvalDatasetList`

NewEvalDatasetListWithDefaults instantiates a new EvalDatasetList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EvalDatasetList) GetData() []EvalDatasetView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EvalDatasetList) GetDataOk() (*[]EvalDatasetView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EvalDatasetList) SetData(v []EvalDatasetView)`

SetData sets Data field to given value.

### HasData

`func (o *EvalDatasetList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



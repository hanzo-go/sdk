# EvalScoreList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]EvalScoreView**](EvalScoreView.md) | Data is the caller org&#39;s score events matching the filters, bounded by limit. | [optional] 

## Methods

### NewEvalScoreList

`func NewEvalScoreList() *EvalScoreList`

NewEvalScoreList instantiates a new EvalScoreList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalScoreListWithDefaults

`func NewEvalScoreListWithDefaults() *EvalScoreList`

NewEvalScoreListWithDefaults instantiates a new EvalScoreList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EvalScoreList) GetData() []EvalScoreView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EvalScoreList) GetDataOk() (*[]EvalScoreView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EvalScoreList) SetData(v []EvalScoreView)`

SetData sets Data field to given value.

### HasData

`func (o *EvalScoreList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



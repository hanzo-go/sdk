# EvalScoreConfigList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]EvalScoreConfigView**](EvalScoreConfigView.md) | Data is the caller org&#39;s rubrics, bounded by limit. | [optional] 

## Methods

### NewEvalScoreConfigList

`func NewEvalScoreConfigList() *EvalScoreConfigList`

NewEvalScoreConfigList instantiates a new EvalScoreConfigList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalScoreConfigListWithDefaults

`func NewEvalScoreConfigListWithDefaults() *EvalScoreConfigList`

NewEvalScoreConfigListWithDefaults instantiates a new EvalScoreConfigList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EvalScoreConfigList) GetData() []EvalScoreConfigView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EvalScoreConfigList) GetDataOk() (*[]EvalScoreConfigView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EvalScoreConfigList) SetData(v []EvalScoreConfigView)`

SetData sets Data field to given value.

### HasData

`func (o *EvalScoreConfigList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



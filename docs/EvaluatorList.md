# EvaluatorList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]EvaluatorView**](EvaluatorView.md) | Data is the caller org&#39;s judges, bounded by limit. | [optional] 

## Methods

### NewEvaluatorList

`func NewEvaluatorList() *EvaluatorList`

NewEvaluatorList instantiates a new EvaluatorList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluatorListWithDefaults

`func NewEvaluatorListWithDefaults() *EvaluatorList`

NewEvaluatorListWithDefaults instantiates a new EvaluatorList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EvaluatorList) GetData() []EvaluatorView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EvaluatorList) GetDataOk() (*[]EvaluatorView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EvaluatorList) SetData(v []EvaluatorView)`

SetData sets Data field to given value.

### HasData

`func (o *EvaluatorList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



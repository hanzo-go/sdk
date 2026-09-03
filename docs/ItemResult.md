# ItemResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error is why this example produced no score — the model, the judge, or the run&#39;s deadline. A result carrying one is not counted in Scored. | [optional] 
**ItemId** | Pointer to **string** | ItemID is the example that was scored. | [optional] 
**Output** | Pointer to **string** | Output is what the model under test answered, truncated at 2000 characters. | [optional] 
**Score** | Pointer to **float64** | Score is the judge&#39;s grade. | [optional] 
**TraceId** | Pointer to **string** | TraceID is the model call this result came from. | [optional] 

## Methods

### NewItemResult

`func NewItemResult() *ItemResult`

NewItemResult instantiates a new ItemResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemResultWithDefaults

`func NewItemResultWithDefaults() *ItemResult`

NewItemResultWithDefaults instantiates a new ItemResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ItemResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ItemResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ItemResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ItemResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetItemId

`func (o *ItemResult) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemResult) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemResult) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ItemResult) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetOutput

`func (o *ItemResult) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ItemResult) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ItemResult) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ItemResult) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetScore

`func (o *ItemResult) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ItemResult) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ItemResult) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *ItemResult) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetTraceId

`func (o *ItemResult) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *ItemResult) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *ItemResult) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *ItemResult) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



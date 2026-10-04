# EvalScoreView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment is the grader&#39;s reasoning, truncated at 2000 characters. | [optional] 
**DataType** | Pointer to **string** | DataType is NUMERIC, CATEGORICAL or BOOLEAN. | [optional] 
**Id** | Pointer to **string** | ID is the score event&#39;s handle. | [optional] 
**Name** | Pointer to **string** | Name is the score name, which a rubric of the same name governs. | [optional] 
**RunName** | Pointer to **string** | RunName is the run this score was recorded under, when it came from one. | [optional] 
**StringValue** | Pointer to **string** | StringValue is the label of a CATEGORICAL score. | [optional] 
**Timestamp** | Pointer to **string** | Timestamp is when the score was recorded. | [optional] 
**TraceId** | Pointer to **string** | TraceID is the model call this score grades, when it grades one. | [optional] 
**Value** | Pointer to **float64** | Value is the numeric score; for BOOLEAN it is 0 or 1. | [optional] 

## Methods

### NewEvalScoreView

`func NewEvalScoreView() *EvalScoreView`

NewEvalScoreView instantiates a new EvalScoreView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalScoreViewWithDefaults

`func NewEvalScoreViewWithDefaults() *EvalScoreView`

NewEvalScoreViewWithDefaults instantiates a new EvalScoreView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *EvalScoreView) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *EvalScoreView) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *EvalScoreView) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *EvalScoreView) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDataType

`func (o *EvalScoreView) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *EvalScoreView) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *EvalScoreView) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *EvalScoreView) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetId

`func (o *EvalScoreView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EvalScoreView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EvalScoreView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EvalScoreView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EvalScoreView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvalScoreView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvalScoreView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EvalScoreView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRunName

`func (o *EvalScoreView) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *EvalScoreView) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *EvalScoreView) SetRunName(v string)`

SetRunName sets RunName field to given value.

### HasRunName

`func (o *EvalScoreView) HasRunName() bool`

HasRunName returns a boolean if a field has been set.

### GetStringValue

`func (o *EvalScoreView) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *EvalScoreView) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *EvalScoreView) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *EvalScoreView) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### GetTimestamp

`func (o *EvalScoreView) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EvalScoreView) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EvalScoreView) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EvalScoreView) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTraceId

`func (o *EvalScoreView) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *EvalScoreView) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *EvalScoreView) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *EvalScoreView) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetValue

`func (o *EvalScoreView) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EvalScoreView) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EvalScoreView) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *EvalScoreView) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



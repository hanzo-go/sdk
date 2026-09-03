# ScoreView

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

### NewScoreView

`func NewScoreView() *ScoreView`

NewScoreView instantiates a new ScoreView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoreViewWithDefaults

`func NewScoreViewWithDefaults() *ScoreView`

NewScoreViewWithDefaults instantiates a new ScoreView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ScoreView) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ScoreView) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ScoreView) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ScoreView) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDataType

`func (o *ScoreView) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ScoreView) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ScoreView) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *ScoreView) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetId

`func (o *ScoreView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScoreView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScoreView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScoreView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ScoreView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScoreView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScoreView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScoreView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRunName

`func (o *ScoreView) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *ScoreView) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *ScoreView) SetRunName(v string)`

SetRunName sets RunName field to given value.

### HasRunName

`func (o *ScoreView) HasRunName() bool`

HasRunName returns a boolean if a field has been set.

### GetStringValue

`func (o *ScoreView) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *ScoreView) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *ScoreView) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *ScoreView) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### GetTimestamp

`func (o *ScoreView) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ScoreView) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ScoreView) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ScoreView) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTraceId

`func (o *ScoreView) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *ScoreView) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *ScoreView) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *ScoreView) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetValue

`func (o *ScoreView) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ScoreView) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ScoreView) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ScoreView) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



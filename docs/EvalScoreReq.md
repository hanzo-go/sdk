# EvalScoreReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment is the grader&#39;s reasoning, truncated at 2000 characters. | [optional] 
**DataType** | Pointer to **string** | DataType is NUMERIC, CATEGORICAL or BOOLEAN. A declared rubric overrides it — a caller cannot claim a type the org&#39;s rubric contradicts. | [optional] 
**DatasetItemId** | Pointer to **string** | ItemID attaches the score to one graded example. | [optional] 
**DatasetName** | Pointer to **string** | Dataset attaches the score to one dataset. | [optional] 
**Name** | **string** | Name is the score name. A rubric of the same name, if the org has declared one, decides this score&#39;s type and the values it may take. | 
**RunName** | Pointer to **string** | RunName attaches the score to one run. | [optional] 
**StringValue** | Pointer to **string** | StringValue is the label of a CATEGORICAL score, which must be one the rubric allows. | [optional] 
**TraceId** | Pointer to **string** | TraceID attaches the score to one model call. | [optional] 
**Value** | Pointer to **float64** | Value is the numeric score, which must be finite: NaN and Inf are refused. A BOOLEAN score takes 0 or 1. | [optional] 

## Methods

### NewEvalScoreReq

`func NewEvalScoreReq(name string, ) *EvalScoreReq`

NewEvalScoreReq instantiates a new EvalScoreReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalScoreReqWithDefaults

`func NewEvalScoreReqWithDefaults() *EvalScoreReq`

NewEvalScoreReqWithDefaults instantiates a new EvalScoreReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *EvalScoreReq) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *EvalScoreReq) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *EvalScoreReq) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *EvalScoreReq) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDataType

`func (o *EvalScoreReq) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *EvalScoreReq) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *EvalScoreReq) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *EvalScoreReq) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDatasetItemId

`func (o *EvalScoreReq) GetDatasetItemId() string`

GetDatasetItemId returns the DatasetItemId field if non-nil, zero value otherwise.

### GetDatasetItemIdOk

`func (o *EvalScoreReq) GetDatasetItemIdOk() (*string, bool)`

GetDatasetItemIdOk returns a tuple with the DatasetItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetItemId

`func (o *EvalScoreReq) SetDatasetItemId(v string)`

SetDatasetItemId sets DatasetItemId field to given value.

### HasDatasetItemId

`func (o *EvalScoreReq) HasDatasetItemId() bool`

HasDatasetItemId returns a boolean if a field has been set.

### GetDatasetName

`func (o *EvalScoreReq) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *EvalScoreReq) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *EvalScoreReq) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.

### HasDatasetName

`func (o *EvalScoreReq) HasDatasetName() bool`

HasDatasetName returns a boolean if a field has been set.

### GetName

`func (o *EvalScoreReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvalScoreReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvalScoreReq) SetName(v string)`

SetName sets Name field to given value.


### GetRunName

`func (o *EvalScoreReq) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *EvalScoreReq) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *EvalScoreReq) SetRunName(v string)`

SetRunName sets RunName field to given value.

### HasRunName

`func (o *EvalScoreReq) HasRunName() bool`

HasRunName returns a boolean if a field has been set.

### GetStringValue

`func (o *EvalScoreReq) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *EvalScoreReq) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *EvalScoreReq) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *EvalScoreReq) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### GetTraceId

`func (o *EvalScoreReq) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *EvalScoreReq) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *EvalScoreReq) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *EvalScoreReq) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetValue

`func (o *EvalScoreReq) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EvalScoreReq) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EvalScoreReq) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *EvalScoreReq) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



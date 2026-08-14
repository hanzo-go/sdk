# O11yO11yLLMIngestScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment is a free-text note. | [optional] 
**DataType** | Pointer to **string** | DataType is the score&#39;s kind — NUMERIC or CATEGORICAL. Defaults from the value when empty. | [optional] 
**Name** | Pointer to **string** | Name is the score&#39;s name, e.g. helpfulness. Required. | [optional] 
**ObservationId** | Pointer to **string** | ObservationID is the single observation the score attaches to, when narrowed to one. | [optional] 
**Source** | Pointer to **string** | Source is where the score came from, e.g. API, EVAL. Defaults to API. | [optional] 
**StringValue** | Pointer to **string** | StringValue is the categorical score, when the score is categorical. | [optional] 
**TraceId** | Pointer to **string** | TraceID is the trace the score attaches to. Required. | [optional] 
**Value** | Pointer to **float32** | Value is the numeric score. | [optional] 

## Methods

### NewO11yO11yLLMIngestScore

`func NewO11yO11yLLMIngestScore() *O11yO11yLLMIngestScore`

NewO11yO11yLLMIngestScore instantiates a new O11yO11yLLMIngestScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMIngestScoreWithDefaults

`func NewO11yO11yLLMIngestScoreWithDefaults() *O11yO11yLLMIngestScore`

NewO11yO11yLLMIngestScoreWithDefaults instantiates a new O11yO11yLLMIngestScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *O11yO11yLLMIngestScore) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *O11yO11yLLMIngestScore) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *O11yO11yLLMIngestScore) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *O11yO11yLLMIngestScore) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDataType

`func (o *O11yO11yLLMIngestScore) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *O11yO11yLLMIngestScore) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *O11yO11yLLMIngestScore) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *O11yO11yLLMIngestScore) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yLLMIngestScore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yLLMIngestScore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yLLMIngestScore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yLLMIngestScore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObservationId

`func (o *O11yO11yLLMIngestScore) GetObservationId() string`

GetObservationId returns the ObservationId field if non-nil, zero value otherwise.

### GetObservationIdOk

`func (o *O11yO11yLLMIngestScore) GetObservationIdOk() (*string, bool)`

GetObservationIdOk returns a tuple with the ObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationId

`func (o *O11yO11yLLMIngestScore) SetObservationId(v string)`

SetObservationId sets ObservationId field to given value.

### HasObservationId

`func (o *O11yO11yLLMIngestScore) HasObservationId() bool`

HasObservationId returns a boolean if a field has been set.

### GetSource

`func (o *O11yO11yLLMIngestScore) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *O11yO11yLLMIngestScore) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *O11yO11yLLMIngestScore) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *O11yO11yLLMIngestScore) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStringValue

`func (o *O11yO11yLLMIngestScore) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *O11yO11yLLMIngestScore) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *O11yO11yLLMIngestScore) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *O11yO11yLLMIngestScore) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### GetTraceId

`func (o *O11yO11yLLMIngestScore) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *O11yO11yLLMIngestScore) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *O11yO11yLLMIngestScore) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *O11yO11yLLMIngestScore) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetValue

`func (o *O11yO11yLLMIngestScore) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yO11yLLMIngestScore) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yO11yLLMIngestScore) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yO11yLLMIngestScore) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



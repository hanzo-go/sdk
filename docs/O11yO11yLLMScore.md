# O11yO11yLLMScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment is a free-text note. | [optional] 
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when the score was stored. | [optional] 
**CreatedBy** | Pointer to **string** | CreatedBy is who created the score. | [optional] 
**DataType** | Pointer to **string** | DataType is the score&#39;s kind — NUMERIC or CATEGORICAL. | [optional] 
**Id** | Pointer to **string** | ID is the score&#39;s id. | [optional] 
**Name** | Pointer to **string** | Name is the score&#39;s name. | [optional] 
**ObservationId** | Pointer to **string** | ObservationID is the observation the score attaches to, when narrowed. | [optional] 
**Source** | Pointer to **string** | Source is where the score came from, e.g. API, EVAL. | [optional] 
**StringValue** | Pointer to **string** | StringValue is the categorical score, when the score is categorical. | [optional] 
**Timestamp** | Pointer to **time.Time** | Timestamp is the score&#39;s own event time. | [optional] 
**TraceId** | Pointer to **string** | TraceID is the trace the score attaches to. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when the score last changed. | [optional] 
**Value** | Pointer to **float32** | Value is the numeric score. | [optional] 

## Methods

### NewO11yO11yLLMScore

`func NewO11yO11yLLMScore() *O11yO11yLLMScore`

NewO11yO11yLLMScore instantiates a new O11yO11yLLMScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMScoreWithDefaults

`func NewO11yO11yLLMScoreWithDefaults() *O11yO11yLLMScore`

NewO11yO11yLLMScoreWithDefaults instantiates a new O11yO11yLLMScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *O11yO11yLLMScore) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *O11yO11yLLMScore) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *O11yO11yLLMScore) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *O11yO11yLLMScore) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11yO11yLLMScore) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yLLMScore) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yLLMScore) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yLLMScore) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *O11yO11yLLMScore) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *O11yO11yLLMScore) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *O11yO11yLLMScore) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *O11yO11yLLMScore) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDataType

`func (o *O11yO11yLLMScore) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *O11yO11yLLMScore) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *O11yO11yLLMScore) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *O11yO11yLLMScore) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yLLMScore) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yLLMScore) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yLLMScore) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yLLMScore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yLLMScore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yLLMScore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yLLMScore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yLLMScore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObservationId

`func (o *O11yO11yLLMScore) GetObservationId() string`

GetObservationId returns the ObservationId field if non-nil, zero value otherwise.

### GetObservationIdOk

`func (o *O11yO11yLLMScore) GetObservationIdOk() (*string, bool)`

GetObservationIdOk returns a tuple with the ObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationId

`func (o *O11yO11yLLMScore) SetObservationId(v string)`

SetObservationId sets ObservationId field to given value.

### HasObservationId

`func (o *O11yO11yLLMScore) HasObservationId() bool`

HasObservationId returns a boolean if a field has been set.

### GetSource

`func (o *O11yO11yLLMScore) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *O11yO11yLLMScore) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *O11yO11yLLMScore) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *O11yO11yLLMScore) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStringValue

`func (o *O11yO11yLLMScore) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *O11yO11yLLMScore) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *O11yO11yLLMScore) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *O11yO11yLLMScore) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### GetTimestamp

`func (o *O11yO11yLLMScore) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *O11yO11yLLMScore) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *O11yO11yLLMScore) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *O11yO11yLLMScore) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTraceId

`func (o *O11yO11yLLMScore) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *O11yO11yLLMScore) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *O11yO11yLLMScore) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *O11yO11yLLMScore) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yLLMScore) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yLLMScore) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yLLMScore) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yLLMScore) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValue

`func (o *O11yO11yLLMScore) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yO11yLLMScore) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yO11yLLMScore) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yO11yLLMScore) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



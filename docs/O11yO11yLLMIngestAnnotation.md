# O11yO11yLLMIngestAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content is the note itself. Required. | [optional] 
**ObservationId** | Pointer to **string** | ObservationID is the single observation the annotation attaches to, when narrowed to one. | [optional] 
**Queue** | Pointer to **string** | Queue is the review queue to file the annotation in. | [optional] 
**Status** | Pointer to **string** | Status is the annotation&#39;s initial review status. Defaults to PENDING. | [optional] 
**TraceId** | Pointer to **string** | TraceID is the trace the annotation attaches to. Required. | [optional] 

## Methods

### NewO11yO11yLLMIngestAnnotation

`func NewO11yO11yLLMIngestAnnotation() *O11yO11yLLMIngestAnnotation`

NewO11yO11yLLMIngestAnnotation instantiates a new O11yO11yLLMIngestAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMIngestAnnotationWithDefaults

`func NewO11yO11yLLMIngestAnnotationWithDefaults() *O11yO11yLLMIngestAnnotation`

NewO11yO11yLLMIngestAnnotationWithDefaults instantiates a new O11yO11yLLMIngestAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *O11yO11yLLMIngestAnnotation) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *O11yO11yLLMIngestAnnotation) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *O11yO11yLLMIngestAnnotation) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *O11yO11yLLMIngestAnnotation) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetObservationId

`func (o *O11yO11yLLMIngestAnnotation) GetObservationId() string`

GetObservationId returns the ObservationId field if non-nil, zero value otherwise.

### GetObservationIdOk

`func (o *O11yO11yLLMIngestAnnotation) GetObservationIdOk() (*string, bool)`

GetObservationIdOk returns a tuple with the ObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationId

`func (o *O11yO11yLLMIngestAnnotation) SetObservationId(v string)`

SetObservationId sets ObservationId field to given value.

### HasObservationId

`func (o *O11yO11yLLMIngestAnnotation) HasObservationId() bool`

HasObservationId returns a boolean if a field has been set.

### GetQueue

`func (o *O11yO11yLLMIngestAnnotation) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *O11yO11yLLMIngestAnnotation) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *O11yO11yLLMIngestAnnotation) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *O11yO11yLLMIngestAnnotation) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yLLMIngestAnnotation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yLLMIngestAnnotation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yLLMIngestAnnotation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yLLMIngestAnnotation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTraceId

`func (o *O11yO11yLLMIngestAnnotation) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *O11yO11yLLMIngestAnnotation) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *O11yO11yLLMIngestAnnotation) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *O11yO11yLLMIngestAnnotation) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



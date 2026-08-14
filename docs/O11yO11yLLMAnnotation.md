# O11yO11yLLMAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** | Author is who wrote it. | [optional] 
**Content** | Pointer to **string** | Content is the note itself. | [optional] 
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when the annotation was stored. | [optional] 
**Id** | Pointer to **string** | ID is the annotation&#39;s id. | [optional] 
**ObservationId** | Pointer to **string** | ObservationID is the observation the annotation attaches to, when narrowed. | [optional] 
**Queue** | Pointer to **string** | Queue is the review queue the annotation sits in, when queued. | [optional] 
**Status** | Pointer to **string** | Status is the annotation&#39;s review status, e.g. PENDING. | [optional] 
**TraceId** | Pointer to **string** | TraceID is the trace the annotation attaches to. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when the annotation last changed. | [optional] 

## Methods

### NewO11yO11yLLMAnnotation

`func NewO11yO11yLLMAnnotation() *O11yO11yLLMAnnotation`

NewO11yO11yLLMAnnotation instantiates a new O11yO11yLLMAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMAnnotationWithDefaults

`func NewO11yO11yLLMAnnotationWithDefaults() *O11yO11yLLMAnnotation`

NewO11yO11yLLMAnnotationWithDefaults instantiates a new O11yO11yLLMAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *O11yO11yLLMAnnotation) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *O11yO11yLLMAnnotation) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *O11yO11yLLMAnnotation) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *O11yO11yLLMAnnotation) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetContent

`func (o *O11yO11yLLMAnnotation) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *O11yO11yLLMAnnotation) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *O11yO11yLLMAnnotation) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *O11yO11yLLMAnnotation) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11yO11yLLMAnnotation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yLLMAnnotation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yLLMAnnotation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yLLMAnnotation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yLLMAnnotation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yLLMAnnotation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yLLMAnnotation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yLLMAnnotation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObservationId

`func (o *O11yO11yLLMAnnotation) GetObservationId() string`

GetObservationId returns the ObservationId field if non-nil, zero value otherwise.

### GetObservationIdOk

`func (o *O11yO11yLLMAnnotation) GetObservationIdOk() (*string, bool)`

GetObservationIdOk returns a tuple with the ObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationId

`func (o *O11yO11yLLMAnnotation) SetObservationId(v string)`

SetObservationId sets ObservationId field to given value.

### HasObservationId

`func (o *O11yO11yLLMAnnotation) HasObservationId() bool`

HasObservationId returns a boolean if a field has been set.

### GetQueue

`func (o *O11yO11yLLMAnnotation) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *O11yO11yLLMAnnotation) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *O11yO11yLLMAnnotation) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *O11yO11yLLMAnnotation) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yLLMAnnotation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yLLMAnnotation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yLLMAnnotation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yLLMAnnotation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTraceId

`func (o *O11yO11yLLMAnnotation) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *O11yO11yLLMAnnotation) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *O11yO11yLLMAnnotation) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *O11yO11yLLMAnnotation) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yLLMAnnotation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yLLMAnnotation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yLLMAnnotation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yLLMAnnotation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



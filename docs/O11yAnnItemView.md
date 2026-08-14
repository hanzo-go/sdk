# O11yAnnItemView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **string** | Assignee is the reviewer it is for, omitted when unassigned. | [optional] 
**CompletedAt** | Pointer to **string** | CompletedAt is when it was reviewed, omitted while pending. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when it was enqueued, RFC3339 in UTC. | [optional] 
**Id** | Pointer to **string** | ID is the item&#39;s id. | [optional] 
**ObjectId** | Pointer to **string** | ObjectID is the referenced object&#39;s id. | [optional] 
**ObjectType** | Pointer to **string** | ObjectType is what it references: TRACE, OBSERVATION or SESSION. | [optional] 
**ObservationId** | Pointer to **string** | ObservationID echoes objectId when objectType is OBSERVATION. | [optional] 
**QueueId** | Pointer to **string** | QueueID is the queue it belongs to. | [optional] 
**SessionId** | Pointer to **string** | SessionID echoes objectId when objectType is SESSION. | [optional] 
**Status** | Pointer to **string** | Status is PENDING or COMPLETED. | [optional] 
**TraceId** | Pointer to **string** | TraceID echoes objectId when objectType is TRACE. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when it last changed, RFC3339 in UTC. | [optional] 

## Methods

### NewO11yAnnItemView

`func NewO11yAnnItemView() *O11yAnnItemView`

NewO11yAnnItemView instantiates a new O11yAnnItemView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAnnItemViewWithDefaults

`func NewO11yAnnItemViewWithDefaults() *O11yAnnItemView`

NewO11yAnnItemViewWithDefaults instantiates a new O11yAnnItemView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *O11yAnnItemView) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *O11yAnnItemView) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *O11yAnnItemView) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *O11yAnnItemView) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetCompletedAt

`func (o *O11yAnnItemView) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *O11yAnnItemView) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *O11yAnnItemView) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *O11yAnnItemView) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11yAnnItemView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yAnnItemView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yAnnItemView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yAnnItemView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *O11yAnnItemView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yAnnItemView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yAnnItemView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yAnnItemView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectId

`func (o *O11yAnnItemView) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *O11yAnnItemView) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *O11yAnnItemView) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *O11yAnnItemView) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectType

`func (o *O11yAnnItemView) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *O11yAnnItemView) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *O11yAnnItemView) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *O11yAnnItemView) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObservationId

`func (o *O11yAnnItemView) GetObservationId() string`

GetObservationId returns the ObservationId field if non-nil, zero value otherwise.

### GetObservationIdOk

`func (o *O11yAnnItemView) GetObservationIdOk() (*string, bool)`

GetObservationIdOk returns a tuple with the ObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationId

`func (o *O11yAnnItemView) SetObservationId(v string)`

SetObservationId sets ObservationId field to given value.

### HasObservationId

`func (o *O11yAnnItemView) HasObservationId() bool`

HasObservationId returns a boolean if a field has been set.

### GetQueueId

`func (o *O11yAnnItemView) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *O11yAnnItemView) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *O11yAnnItemView) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *O11yAnnItemView) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetSessionId

`func (o *O11yAnnItemView) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *O11yAnnItemView) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *O11yAnnItemView) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *O11yAnnItemView) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetStatus

`func (o *O11yAnnItemView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yAnnItemView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yAnnItemView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yAnnItemView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTraceId

`func (o *O11yAnnItemView) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *O11yAnnItemView) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *O11yAnnItemView) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *O11yAnnItemView) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yAnnItemView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yAnnItemView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yAnnItemView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yAnnItemView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



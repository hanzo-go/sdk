# CloudAnnItemView

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

### NewCloudAnnItemView

`func NewCloudAnnItemView() *CloudAnnItemView`

NewCloudAnnItemView instantiates a new CloudAnnItemView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAnnItemViewWithDefaults

`func NewCloudAnnItemViewWithDefaults() *CloudAnnItemView`

NewCloudAnnItemViewWithDefaults instantiates a new CloudAnnItemView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *CloudAnnItemView) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *CloudAnnItemView) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *CloudAnnItemView) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *CloudAnnItemView) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetCompletedAt

`func (o *CloudAnnItemView) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *CloudAnnItemView) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *CloudAnnItemView) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *CloudAnnItemView) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudAnnItemView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudAnnItemView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudAnnItemView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudAnnItemView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CloudAnnItemView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAnnItemView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAnnItemView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAnnItemView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectId

`func (o *CloudAnnItemView) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *CloudAnnItemView) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *CloudAnnItemView) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *CloudAnnItemView) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectType

`func (o *CloudAnnItemView) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudAnnItemView) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudAnnItemView) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *CloudAnnItemView) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObservationId

`func (o *CloudAnnItemView) GetObservationId() string`

GetObservationId returns the ObservationId field if non-nil, zero value otherwise.

### GetObservationIdOk

`func (o *CloudAnnItemView) GetObservationIdOk() (*string, bool)`

GetObservationIdOk returns a tuple with the ObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationId

`func (o *CloudAnnItemView) SetObservationId(v string)`

SetObservationId sets ObservationId field to given value.

### HasObservationId

`func (o *CloudAnnItemView) HasObservationId() bool`

HasObservationId returns a boolean if a field has been set.

### GetQueueId

`func (o *CloudAnnItemView) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *CloudAnnItemView) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *CloudAnnItemView) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *CloudAnnItemView) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetSessionId

`func (o *CloudAnnItemView) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CloudAnnItemView) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CloudAnnItemView) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CloudAnnItemView) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAnnItemView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAnnItemView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAnnItemView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAnnItemView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTraceId

`func (o *CloudAnnItemView) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *CloudAnnItemView) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *CloudAnnItemView) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *CloudAnnItemView) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudAnnItemView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudAnnItemView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudAnnItemView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudAnnItemView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



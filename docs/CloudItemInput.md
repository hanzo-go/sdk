# CloudItemInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **string** | Assignee is the reviewer this item is for, up to 512 characters. | [optional] 
**ObjectId** | Pointer to **string** | ObjectID is the referenced object&#39;s id, paired with objectType. | [optional] 
**ObjectType** | Pointer to **string** | ObjectType is TRACE, OBSERVATION or SESSION — the generic form, paired with objectId. | [optional] 
**ObservationId** | Pointer to **string** | ObservationID references an observation — the console-friendly form of objectType&#x3D;OBSERVATION. | [optional] 
**SessionId** | Pointer to **string** | SessionID references a session — the console-friendly form of objectType&#x3D;SESSION. | [optional] 
**TraceId** | Pointer to **string** | TraceID references a trace — the console-friendly form of objectType&#x3D;TRACE. | [optional] 

## Methods

### NewCloudItemInput

`func NewCloudItemInput() *CloudItemInput`

NewCloudItemInput instantiates a new CloudItemInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudItemInputWithDefaults

`func NewCloudItemInputWithDefaults() *CloudItemInput`

NewCloudItemInputWithDefaults instantiates a new CloudItemInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *CloudItemInput) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *CloudItemInput) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *CloudItemInput) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *CloudItemInput) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetObjectId

`func (o *CloudItemInput) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *CloudItemInput) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *CloudItemInput) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *CloudItemInput) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectType

`func (o *CloudItemInput) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudItemInput) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudItemInput) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *CloudItemInput) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObservationId

`func (o *CloudItemInput) GetObservationId() string`

GetObservationId returns the ObservationId field if non-nil, zero value otherwise.

### GetObservationIdOk

`func (o *CloudItemInput) GetObservationIdOk() (*string, bool)`

GetObservationIdOk returns a tuple with the ObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationId

`func (o *CloudItemInput) SetObservationId(v string)`

SetObservationId sets ObservationId field to given value.

### HasObservationId

`func (o *CloudItemInput) HasObservationId() bool`

HasObservationId returns a boolean if a field has been set.

### GetSessionId

`func (o *CloudItemInput) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CloudItemInput) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CloudItemInput) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CloudItemInput) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetTraceId

`func (o *CloudItemInput) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *CloudItemInput) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *CloudItemInput) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *CloudItemInput) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



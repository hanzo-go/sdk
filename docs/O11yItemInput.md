# O11yItemInput

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

### NewO11yItemInput

`func NewO11yItemInput() *O11yItemInput`

NewO11yItemInput instantiates a new O11yItemInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yItemInputWithDefaults

`func NewO11yItemInputWithDefaults() *O11yItemInput`

NewO11yItemInputWithDefaults instantiates a new O11yItemInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *O11yItemInput) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *O11yItemInput) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *O11yItemInput) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *O11yItemInput) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetObjectId

`func (o *O11yItemInput) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *O11yItemInput) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *O11yItemInput) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *O11yItemInput) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectType

`func (o *O11yItemInput) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *O11yItemInput) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *O11yItemInput) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *O11yItemInput) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObservationId

`func (o *O11yItemInput) GetObservationId() string`

GetObservationId returns the ObservationId field if non-nil, zero value otherwise.

### GetObservationIdOk

`func (o *O11yItemInput) GetObservationIdOk() (*string, bool)`

GetObservationIdOk returns a tuple with the ObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationId

`func (o *O11yItemInput) SetObservationId(v string)`

SetObservationId sets ObservationId field to given value.

### HasObservationId

`func (o *O11yItemInput) HasObservationId() bool`

HasObservationId returns a boolean if a field has been set.

### GetSessionId

`func (o *O11yItemInput) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *O11yItemInput) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *O11yItemInput) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *O11yItemInput) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetTraceId

`func (o *O11yItemInput) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *O11yItemInput) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *O11yItemInput) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *O11yItemInput) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



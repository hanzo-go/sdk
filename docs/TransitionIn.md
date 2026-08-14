# TransitionIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doctype** | Pointer to **string** | DocType is the content type to act on, from the path. | [optional] 
**Name** | Pointer to **string** | Name is the document to act on, from the path. | [optional] 
**ScheduleAt** | Pointer to **string** | ScheduleAt is an ISO-8601 go-live time handed to the channel&#39;s own scheduler; \&quot;\&quot; distributes now. | [optional] 
**To** | Pointer to **string** | To is the lifecycle state to move to. Required, and the move must be a legal edge from the item&#39;s current state. | [optional] 

## Methods

### NewTransitionIn

`func NewTransitionIn() *TransitionIn`

NewTransitionIn instantiates a new TransitionIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitionInWithDefaults

`func NewTransitionInWithDefaults() *TransitionIn`

NewTransitionInWithDefaults instantiates a new TransitionIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoctype

`func (o *TransitionIn) GetDoctype() string`

GetDoctype returns the Doctype field if non-nil, zero value otherwise.

### GetDoctypeOk

`func (o *TransitionIn) GetDoctypeOk() (*string, bool)`

GetDoctypeOk returns a tuple with the Doctype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctype

`func (o *TransitionIn) SetDoctype(v string)`

SetDoctype sets Doctype field to given value.

### HasDoctype

`func (o *TransitionIn) HasDoctype() bool`

HasDoctype returns a boolean if a field has been set.

### GetName

`func (o *TransitionIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransitionIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransitionIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TransitionIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduleAt

`func (o *TransitionIn) GetScheduleAt() string`

GetScheduleAt returns the ScheduleAt field if non-nil, zero value otherwise.

### GetScheduleAtOk

`func (o *TransitionIn) GetScheduleAtOk() (*string, bool)`

GetScheduleAtOk returns a tuple with the ScheduleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleAt

`func (o *TransitionIn) SetScheduleAt(v string)`

SetScheduleAt sets ScheduleAt field to given value.

### HasScheduleAt

`func (o *TransitionIn) HasScheduleAt() bool`

HasScheduleAt returns a boolean if a field has been set.

### GetTo

`func (o *TransitionIn) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TransitionIn) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TransitionIn) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *TransitionIn) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



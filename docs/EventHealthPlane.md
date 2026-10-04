# EventHealthPlane

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bus** | Pointer to **string** | Bus is the address this process reaches the plane at. | [optional] 
**Ready** | Pointer to **bool** | Ready reports whether an ingest would succeed right now. False is a 503. | [optional] 
**Reason** | Pointer to **string** | Reason is the plane&#39;s own failure text, present only when Ready is false. | [optional] 
**Stream** | Pointer to **string** | Stream is the JetStream stream every signal lands on. | [optional] 

## Methods

### NewEventHealthPlane

`func NewEventHealthPlane() *EventHealthPlane`

NewEventHealthPlane instantiates a new EventHealthPlane object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventHealthPlaneWithDefaults

`func NewEventHealthPlaneWithDefaults() *EventHealthPlane`

NewEventHealthPlaneWithDefaults instantiates a new EventHealthPlane object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBus

`func (o *EventHealthPlane) GetBus() string`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *EventHealthPlane) GetBusOk() (*string, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *EventHealthPlane) SetBus(v string)`

SetBus sets Bus field to given value.

### HasBus

`func (o *EventHealthPlane) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetReady

`func (o *EventHealthPlane) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *EventHealthPlane) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *EventHealthPlane) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *EventHealthPlane) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetReason

`func (o *EventHealthPlane) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EventHealthPlane) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EventHealthPlane) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EventHealthPlane) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStream

`func (o *EventHealthPlane) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *EventHealthPlane) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *EventHealthPlane) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *EventHealthPlane) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



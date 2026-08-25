# RecordIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Room** | **string** | Room is the LiveKit room, named the way the office client names one (&#x60;&lt;workspace&gt;_&lt;name&gt;_&lt;id&gt;&#x60;). Its leading segment is what binds the room to a tenant, and it is the segment the caller&#39;s membership is checked against. | 

## Methods

### NewRecordIn

`func NewRecordIn(room string, ) *RecordIn`

NewRecordIn instantiates a new RecordIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordInWithDefaults

`func NewRecordInWithDefaults() *RecordIn`

NewRecordInWithDefaults instantiates a new RecordIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoom

`func (o *RecordIn) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *RecordIn) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *RecordIn) SetRoom(v string)`

SetRoom sets Room field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



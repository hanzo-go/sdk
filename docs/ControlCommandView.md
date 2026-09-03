# ControlCommandView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** | Command is what was asked, from a closed four: pause, resume, stop, message. It is an INTENT — the poller decides what to do about it, and the session&#39;s status changes only when the poller reports back that it did. | [optional] 
**Message** | Pointer to **string** | Message is the text that came with the command: what to say into the run for &#x60;message&#x60;, and the cancellation reason for &#x60;stop&#x60;. Up to 16 KiB. Empty on a bare pause or resume. | [optional] 
**Payload** | Pointer to **interface{}** |  | [optional] 
**Seq** | Pointer to **int64** | Seq is this command&#39;s position in the session&#39;s log — the same monotonic number every other turn is ordered by, so a command sits in the transcript where it was issued. Send the highest one you applied back as &#x60;after&#x60; and it is never redelivered. | [optional] 

## Methods

### NewControlCommandView

`func NewControlCommandView() *ControlCommandView`

NewControlCommandView instantiates a new ControlCommandView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlCommandViewWithDefaults

`func NewControlCommandViewWithDefaults() *ControlCommandView`

NewControlCommandViewWithDefaults instantiates a new ControlCommandView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *ControlCommandView) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ControlCommandView) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ControlCommandView) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *ControlCommandView) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetMessage

`func (o *ControlCommandView) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ControlCommandView) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ControlCommandView) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ControlCommandView) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPayload

`func (o *ControlCommandView) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ControlCommandView) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ControlCommandView) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ControlCommandView) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *ControlCommandView) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *ControlCommandView) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetSeq

`func (o *ControlCommandView) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *ControlCommandView) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *ControlCommandView) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *ControlCommandView) HasSeq() bool`

HasSeq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



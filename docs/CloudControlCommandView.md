# CloudControlCommandView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **interface{}** |  | [optional] 
**Seq** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudControlCommandView

`func NewCloudControlCommandView() *CloudControlCommandView`

NewCloudControlCommandView instantiates a new CloudControlCommandView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudControlCommandViewWithDefaults

`func NewCloudControlCommandViewWithDefaults() *CloudControlCommandView`

NewCloudControlCommandViewWithDefaults instantiates a new CloudControlCommandView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *CloudControlCommandView) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CloudControlCommandView) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CloudControlCommandView) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CloudControlCommandView) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetMessage

`func (o *CloudControlCommandView) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudControlCommandView) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudControlCommandView) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CloudControlCommandView) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPayload

`func (o *CloudControlCommandView) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CloudControlCommandView) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CloudControlCommandView) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CloudControlCommandView) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *CloudControlCommandView) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *CloudControlCommandView) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetSeq

`func (o *CloudControlCommandView) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *CloudControlCommandView) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *CloudControlCommandView) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *CloudControlCommandView) HasSeq() bool`

HasSeq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DestinationTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error is the platform&#39;s rejection, present only on a failed send. | [optional] 
**Message** | Pointer to **string** | Message is the platform&#39;s own note about the send, present only on success. | [optional] 
**Ok** | Pointer to **bool** | OK is true when the platform accepted the synthetic event. | [optional] 
**Sent** | Pointer to **int32** | Sent is how many events the platform accepted, present only on success. | [optional] 

## Methods

### NewDestinationTest

`func NewDestinationTest() *DestinationTest`

NewDestinationTest instantiates a new DestinationTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationTestWithDefaults

`func NewDestinationTestWithDefaults() *DestinationTest`

NewDestinationTestWithDefaults instantiates a new DestinationTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *DestinationTest) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DestinationTest) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DestinationTest) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DestinationTest) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *DestinationTest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DestinationTest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DestinationTest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DestinationTest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOk

`func (o *DestinationTest) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *DestinationTest) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *DestinationTest) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *DestinationTest) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetSent

`func (o *DestinationTest) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *DestinationTest) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *DestinationTest) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *DestinationTest) HasSent() bool`

HasSent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



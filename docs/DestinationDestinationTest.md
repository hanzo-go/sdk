# DestinationDestinationTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error is the platform&#39;s rejection, present only on a failed send. | [optional] 
**Message** | Pointer to **string** | Message is the platform&#39;s own note about the send, present only on success. | [optional] 
**Ok** | Pointer to **bool** | OK is true when the platform accepted the synthetic event. | [optional] 
**Sent** | Pointer to **int64** | Sent is how many events the platform accepted, present only on success. | [optional] 

## Methods

### NewDestinationDestinationTest

`func NewDestinationDestinationTest() *DestinationDestinationTest`

NewDestinationDestinationTest instantiates a new DestinationDestinationTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationDestinationTestWithDefaults

`func NewDestinationDestinationTestWithDefaults() *DestinationDestinationTest`

NewDestinationDestinationTestWithDefaults instantiates a new DestinationDestinationTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *DestinationDestinationTest) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DestinationDestinationTest) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DestinationDestinationTest) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DestinationDestinationTest) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *DestinationDestinationTest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DestinationDestinationTest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DestinationDestinationTest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DestinationDestinationTest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOk

`func (o *DestinationDestinationTest) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *DestinationDestinationTest) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *DestinationDestinationTest) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *DestinationDestinationTest) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetSent

`func (o *DestinationDestinationTest) GetSent() int64`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *DestinationDestinationTest) GetSentOk() (*int64, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *DestinationDestinationTest) SetSent(v int64)`

SetSent sets Sent field to given value.

### HasSent

`func (o *DestinationDestinationTest) HasSent() bool`

HasSent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



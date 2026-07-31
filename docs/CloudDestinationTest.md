# CloudDestinationTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error is the platform&#39;s rejection, present only on a failed send. | [optional] 
**Message** | Pointer to **string** | Message is the platform&#39;s own note about the send, present only on success. | [optional] 
**Ok** | Pointer to **bool** | OK is true when the platform accepted the synthetic event. | [optional] 
**Sent** | Pointer to **int32** | Sent is how many events the platform accepted, present only on success. | [optional] 

## Methods

### NewCloudDestinationTest

`func NewCloudDestinationTest() *CloudDestinationTest`

NewCloudDestinationTest instantiates a new CloudDestinationTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDestinationTestWithDefaults

`func NewCloudDestinationTestWithDefaults() *CloudDestinationTest`

NewCloudDestinationTestWithDefaults instantiates a new CloudDestinationTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *CloudDestinationTest) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CloudDestinationTest) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CloudDestinationTest) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CloudDestinationTest) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *CloudDestinationTest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudDestinationTest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudDestinationTest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CloudDestinationTest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOk

`func (o *CloudDestinationTest) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *CloudDestinationTest) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *CloudDestinationTest) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *CloudDestinationTest) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetSent

`func (o *CloudDestinationTest) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *CloudDestinationTest) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *CloudDestinationTest) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *CloudDestinationTest) HasSent() bool`

HasSent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



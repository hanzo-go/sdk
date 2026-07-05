# ConsoleError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewConsoleError

`func NewConsoleError() *ConsoleError`

NewConsoleError instantiates a new ConsoleError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleErrorWithDefaults

`func NewConsoleErrorWithDefaults() *ConsoleError`

NewConsoleErrorWithDefaults instantiates a new ConsoleError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ConsoleError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConsoleError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConsoleError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConsoleError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetError

`func (o *ConsoleError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ConsoleError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ConsoleError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ConsoleError) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



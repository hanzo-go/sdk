# OpenaiToolCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Function** | Pointer to [**OpenaiFunctionCall**](OpenaiFunctionCall.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewOpenaiToolCall

`func NewOpenaiToolCall() *OpenaiToolCall`

NewOpenaiToolCall instantiates a new OpenaiToolCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenaiToolCallWithDefaults

`func NewOpenaiToolCallWithDefaults() *OpenaiToolCall`

NewOpenaiToolCallWithDefaults instantiates a new OpenaiToolCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunction

`func (o *OpenaiToolCall) GetFunction() OpenaiFunctionCall`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *OpenaiToolCall) GetFunctionOk() (*OpenaiFunctionCall, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *OpenaiToolCall) SetFunction(v OpenaiFunctionCall)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *OpenaiToolCall) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetId

`func (o *OpenaiToolCall) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenaiToolCall) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenaiToolCall) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OpenaiToolCall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *OpenaiToolCall) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *OpenaiToolCall) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *OpenaiToolCall) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *OpenaiToolCall) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetType

`func (o *OpenaiToolCall) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpenaiToolCall) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpenaiToolCall) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OpenaiToolCall) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



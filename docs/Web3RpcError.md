# Web3RpcError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int64** | Code is the JSON-RPC error code the chain reported, passed through as it came. -32603 (internal error) is the one value this deployment mints itself, for an upstream that could not be reached at all. | [optional] 
**Message** | Pointer to **string** | Message is the chain&#39;s own explanation, e.g. \&quot;execution reverted\&quot;. It is \&quot;upstream unavailable\&quot; when the deployment minted the error rather than the chain — that is the one message this side writes. | [optional] 

## Methods

### NewWeb3RpcError

`func NewWeb3RpcError() *Web3RpcError`

NewWeb3RpcError instantiates a new Web3RpcError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeb3RpcErrorWithDefaults

`func NewWeb3RpcErrorWithDefaults() *Web3RpcError`

NewWeb3RpcErrorWithDefaults instantiates a new Web3RpcError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Web3RpcError) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Web3RpcError) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Web3RpcError) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *Web3RpcError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *Web3RpcError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Web3RpcError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Web3RpcError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Web3RpcError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



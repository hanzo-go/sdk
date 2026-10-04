# Web3RpcOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**Web3RpcError**](Web3RpcError.md) | Error is the JSON-RPC error object, present instead of Result. Its presence is the ONLY way a failure shows up here: the HTTP status stays 200, because that is what a standard JSON-RPC client parses. | [optional] 
**Id** | Pointer to **interface{}** |  | [optional] 
**Jsonrpc** | Pointer to **string** | JSONRPC is always \&quot;2.0\&quot;. An upstream that omits it has it filled in, so a client never has to cope with a response that is missing the one field telling it which protocol it is reading. | [optional] 
**Result** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewWeb3RpcOut

`func NewWeb3RpcOut() *Web3RpcOut`

NewWeb3RpcOut instantiates a new Web3RpcOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeb3RpcOutWithDefaults

`func NewWeb3RpcOutWithDefaults() *Web3RpcOut`

NewWeb3RpcOutWithDefaults instantiates a new Web3RpcOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *Web3RpcOut) GetError() Web3RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Web3RpcOut) GetErrorOk() (*Web3RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Web3RpcOut) SetError(v Web3RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *Web3RpcOut) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *Web3RpcOut) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Web3RpcOut) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Web3RpcOut) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *Web3RpcOut) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Web3RpcOut) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Web3RpcOut) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetJsonrpc

`func (o *Web3RpcOut) GetJsonrpc() string`

GetJsonrpc returns the Jsonrpc field if non-nil, zero value otherwise.

### GetJsonrpcOk

`func (o *Web3RpcOut) GetJsonrpcOk() (*string, bool)`

GetJsonrpcOk returns a tuple with the Jsonrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonrpc

`func (o *Web3RpcOut) SetJsonrpc(v string)`

SetJsonrpc sets Jsonrpc field to given value.

### HasJsonrpc

`func (o *Web3RpcOut) HasJsonrpc() bool`

HasJsonrpc returns a boolean if a field has been set.

### GetResult

`func (o *Web3RpcOut) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Web3RpcOut) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Web3RpcOut) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *Web3RpcOut) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *Web3RpcOut) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *Web3RpcOut) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



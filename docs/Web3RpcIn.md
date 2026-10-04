# Web3RpcIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | Pointer to **string** | Chain is the registry id, from the URL. | [optional] 
**Id** | Pointer to **interface{}** |  | [optional] 
**Jsonrpc** | Pointer to **string** | JSONRPC must be \&quot;2.0\&quot; when present. | [optional] 
**Method** | Pointer to **string** | Method is the RPC method, e.g. eth_getBalance. | [optional] 
**Params** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewWeb3RpcIn

`func NewWeb3RpcIn() *Web3RpcIn`

NewWeb3RpcIn instantiates a new Web3RpcIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeb3RpcInWithDefaults

`func NewWeb3RpcInWithDefaults() *Web3RpcIn`

NewWeb3RpcInWithDefaults instantiates a new Web3RpcIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *Web3RpcIn) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *Web3RpcIn) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *Web3RpcIn) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *Web3RpcIn) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetId

`func (o *Web3RpcIn) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Web3RpcIn) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Web3RpcIn) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *Web3RpcIn) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Web3RpcIn) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Web3RpcIn) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetJsonrpc

`func (o *Web3RpcIn) GetJsonrpc() string`

GetJsonrpc returns the Jsonrpc field if non-nil, zero value otherwise.

### GetJsonrpcOk

`func (o *Web3RpcIn) GetJsonrpcOk() (*string, bool)`

GetJsonrpcOk returns a tuple with the Jsonrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonrpc

`func (o *Web3RpcIn) SetJsonrpc(v string)`

SetJsonrpc sets Jsonrpc field to given value.

### HasJsonrpc

`func (o *Web3RpcIn) HasJsonrpc() bool`

HasJsonrpc returns a boolean if a field has been set.

### GetMethod

`func (o *Web3RpcIn) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Web3RpcIn) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Web3RpcIn) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Web3RpcIn) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetParams

`func (o *Web3RpcIn) GetParams() interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Web3RpcIn) GetParamsOk() (*interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Web3RpcIn) SetParams(v interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *Web3RpcIn) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *Web3RpcIn) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *Web3RpcIn) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



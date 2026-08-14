# RpcOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 
**Id** | Pointer to **interface{}** |  | [optional] 
**Jsonrpc** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewRpcOut

`func NewRpcOut() *RpcOut`

NewRpcOut instantiates a new RpcOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcOutWithDefaults

`func NewRpcOutWithDefaults() *RpcOut`

NewRpcOutWithDefaults instantiates a new RpcOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *RpcOut) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RpcOut) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RpcOut) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *RpcOut) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *RpcOut) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RpcOut) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RpcOut) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *RpcOut) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RpcOut) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RpcOut) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetJsonrpc

`func (o *RpcOut) GetJsonrpc() string`

GetJsonrpc returns the Jsonrpc field if non-nil, zero value otherwise.

### GetJsonrpcOk

`func (o *RpcOut) GetJsonrpcOk() (*string, bool)`

GetJsonrpcOk returns a tuple with the Jsonrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonrpc

`func (o *RpcOut) SetJsonrpc(v string)`

SetJsonrpc sets Jsonrpc field to given value.

### HasJsonrpc

`func (o *RpcOut) HasJsonrpc() bool`

HasJsonrpc returns a boolean if a field has been set.

### GetResult

`func (o *RpcOut) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RpcOut) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RpcOut) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *RpcOut) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *RpcOut) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *RpcOut) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



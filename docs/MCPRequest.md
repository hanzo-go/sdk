# MCPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **interface{}** |  | [optional] 
**Jsonrpc** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewMCPRequest

`func NewMCPRequest() *MCPRequest`

NewMCPRequest instantiates a new MCPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPRequestWithDefaults

`func NewMCPRequestWithDefaults() *MCPRequest`

NewMCPRequestWithDefaults instantiates a new MCPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MCPRequest) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MCPRequest) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MCPRequest) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *MCPRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MCPRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MCPRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetJsonrpc

`func (o *MCPRequest) GetJsonrpc() string`

GetJsonrpc returns the Jsonrpc field if non-nil, zero value otherwise.

### GetJsonrpcOk

`func (o *MCPRequest) GetJsonrpcOk() (*string, bool)`

GetJsonrpcOk returns a tuple with the Jsonrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonrpc

`func (o *MCPRequest) SetJsonrpc(v string)`

SetJsonrpc sets Jsonrpc field to given value.

### HasJsonrpc

`func (o *MCPRequest) HasJsonrpc() bool`

HasJsonrpc returns a boolean if a field has been set.

### GetMethod

`func (o *MCPRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *MCPRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *MCPRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *MCPRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetParams

`func (o *MCPRequest) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *MCPRequest) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *MCPRequest) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *MCPRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



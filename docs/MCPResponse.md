# MCPResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**MCPError**](MCPError.md) |  | [optional] 
**Id** | Pointer to **interface{}** |  | [optional] 
**Jsonrpc** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewMCPResponse

`func NewMCPResponse() *MCPResponse`

NewMCPResponse instantiates a new MCPResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPResponseWithDefaults

`func NewMCPResponseWithDefaults() *MCPResponse`

NewMCPResponseWithDefaults instantiates a new MCPResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *MCPResponse) GetError() MCPError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MCPResponse) GetErrorOk() (*MCPError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MCPResponse) SetError(v MCPError)`

SetError sets Error field to given value.

### HasError

`func (o *MCPResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *MCPResponse) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MCPResponse) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MCPResponse) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *MCPResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MCPResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MCPResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetJsonrpc

`func (o *MCPResponse) GetJsonrpc() string`

GetJsonrpc returns the Jsonrpc field if non-nil, zero value otherwise.

### GetJsonrpcOk

`func (o *MCPResponse) GetJsonrpcOk() (*string, bool)`

GetJsonrpcOk returns a tuple with the Jsonrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonrpc

`func (o *MCPResponse) SetJsonrpc(v string)`

SetJsonrpc sets Jsonrpc field to given value.

### HasJsonrpc

`func (o *MCPResponse) HasJsonrpc() bool`

HasJsonrpc returns a boolean if a field has been set.

### GetResult

`func (o *MCPResponse) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MCPResponse) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *MCPResponse) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *MCPResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



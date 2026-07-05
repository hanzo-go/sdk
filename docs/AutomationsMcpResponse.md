# AutomationsMcpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jsonrpc** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **interface{}** |  | [optional] 
**Result** | Pointer to **interface{}** |  | [optional] 
**Error** | Pointer to [**AutomationsMcpResponseError**](AutomationsMcpResponseError.md) |  | [optional] 

## Methods

### NewAutomationsMcpResponse

`func NewAutomationsMcpResponse() *AutomationsMcpResponse`

NewAutomationsMcpResponse instantiates a new AutomationsMcpResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationsMcpResponseWithDefaults

`func NewAutomationsMcpResponseWithDefaults() *AutomationsMcpResponse`

NewAutomationsMcpResponseWithDefaults instantiates a new AutomationsMcpResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonrpc

`func (o *AutomationsMcpResponse) GetJsonrpc() string`

GetJsonrpc returns the Jsonrpc field if non-nil, zero value otherwise.

### GetJsonrpcOk

`func (o *AutomationsMcpResponse) GetJsonrpcOk() (*string, bool)`

GetJsonrpcOk returns a tuple with the Jsonrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonrpc

`func (o *AutomationsMcpResponse) SetJsonrpc(v string)`

SetJsonrpc sets Jsonrpc field to given value.

### HasJsonrpc

`func (o *AutomationsMcpResponse) HasJsonrpc() bool`

HasJsonrpc returns a boolean if a field has been set.

### GetId

`func (o *AutomationsMcpResponse) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationsMcpResponse) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationsMcpResponse) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *AutomationsMcpResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AutomationsMcpResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AutomationsMcpResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetResult

`func (o *AutomationsMcpResponse) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AutomationsMcpResponse) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AutomationsMcpResponse) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *AutomationsMcpResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *AutomationsMcpResponse) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *AutomationsMcpResponse) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetError

`func (o *AutomationsMcpResponse) GetError() AutomationsMcpResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AutomationsMcpResponse) GetErrorOk() (*AutomationsMcpResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AutomationsMcpResponse) SetError(v AutomationsMcpResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *AutomationsMcpResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



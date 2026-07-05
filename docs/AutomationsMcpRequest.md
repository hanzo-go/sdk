# AutomationsMcpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jsonrpc** | **string** |  | 
**Id** | Pointer to **interface{}** |  | [optional] 
**Method** | **string** |  | 
**Params** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewAutomationsMcpRequest

`func NewAutomationsMcpRequest(jsonrpc string, method string, ) *AutomationsMcpRequest`

NewAutomationsMcpRequest instantiates a new AutomationsMcpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationsMcpRequestWithDefaults

`func NewAutomationsMcpRequestWithDefaults() *AutomationsMcpRequest`

NewAutomationsMcpRequestWithDefaults instantiates a new AutomationsMcpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonrpc

`func (o *AutomationsMcpRequest) GetJsonrpc() string`

GetJsonrpc returns the Jsonrpc field if non-nil, zero value otherwise.

### GetJsonrpcOk

`func (o *AutomationsMcpRequest) GetJsonrpcOk() (*string, bool)`

GetJsonrpcOk returns a tuple with the Jsonrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonrpc

`func (o *AutomationsMcpRequest) SetJsonrpc(v string)`

SetJsonrpc sets Jsonrpc field to given value.


### GetId

`func (o *AutomationsMcpRequest) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationsMcpRequest) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationsMcpRequest) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *AutomationsMcpRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AutomationsMcpRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AutomationsMcpRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMethod

`func (o *AutomationsMcpRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AutomationsMcpRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AutomationsMcpRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetParams

`func (o *AutomationsMcpRequest) GetParams() interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *AutomationsMcpRequest) GetParamsOk() (*interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *AutomationsMcpRequest) SetParams(v interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *AutomationsMcpRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *AutomationsMcpRequest) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *AutomationsMcpRequest) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



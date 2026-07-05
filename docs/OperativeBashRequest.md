# OperativeBashRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** | Bash command to execute | [optional] 
**Restart** | Pointer to **bool** | Restart the bash session | [optional] [default to false]

## Methods

### NewOperativeBashRequest

`func NewOperativeBashRequest() *OperativeBashRequest`

NewOperativeBashRequest instantiates a new OperativeBashRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperativeBashRequestWithDefaults

`func NewOperativeBashRequestWithDefaults() *OperativeBashRequest`

NewOperativeBashRequestWithDefaults instantiates a new OperativeBashRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *OperativeBashRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *OperativeBashRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *OperativeBashRequest) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *OperativeBashRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetRestart

`func (o *OperativeBashRequest) GetRestart() bool`

GetRestart returns the Restart field if non-nil, zero value otherwise.

### GetRestartOk

`func (o *OperativeBashRequest) GetRestartOk() (*bool, bool)`

GetRestartOk returns a tuple with the Restart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestart

`func (o *OperativeBashRequest) SetRestart(v bool)`

SetRestart sets Restart field to given value.

### HasRestart

`func (o *OperativeBashRequest) HasRestart() bool`

HasRestart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



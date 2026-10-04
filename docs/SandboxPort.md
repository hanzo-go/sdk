# SandboxPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host is the preview host that serves it, the same for as long as the sandbox lives, parked or running. | [optional] 
**Port** | Pointer to **int64** | Port is the TCP port. | [optional] 

## Methods

### NewSandboxPort

`func NewSandboxPort() *SandboxPort`

NewSandboxPort instantiates a new SandboxPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxPortWithDefaults

`func NewSandboxPortWithDefaults() *SandboxPort`

NewSandboxPortWithDefaults instantiates a new SandboxPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SandboxPort) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SandboxPort) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SandboxPort) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SandboxPort) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *SandboxPort) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SandboxPort) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SandboxPort) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SandboxPort) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



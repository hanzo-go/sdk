# SandboxPreviewIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the sandbox, from the path. | [optional] 
**Port** | Pointer to **int64** | Port is the TCP port inside the sandbox to open, 1 to 65535. Whatever listens on it on localhost is what the address serves. | [optional] 

## Methods

### NewSandboxPreviewIn

`func NewSandboxPreviewIn() *SandboxPreviewIn`

NewSandboxPreviewIn instantiates a new SandboxPreviewIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxPreviewInWithDefaults

`func NewSandboxPreviewInWithDefaults() *SandboxPreviewIn`

NewSandboxPreviewInWithDefaults instantiates a new SandboxPreviewIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SandboxPreviewIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SandboxPreviewIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SandboxPreviewIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SandboxPreviewIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPort

`func (o *SandboxPreviewIn) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SandboxPreviewIn) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SandboxPreviewIn) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SandboxPreviewIn) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



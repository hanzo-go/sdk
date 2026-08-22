# SandboxList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sandboxes** | Pointer to [**[]Sandbox**](Sandbox.md) | Sandboxes are the caller org&#39;s sandboxes matching the filter. Never null. | [optional] 

## Methods

### NewSandboxList

`func NewSandboxList() *SandboxList`

NewSandboxList instantiates a new SandboxList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxListWithDefaults

`func NewSandboxListWithDefaults() *SandboxList`

NewSandboxListWithDefaults instantiates a new SandboxList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSandboxes

`func (o *SandboxList) GetSandboxes() []Sandbox`

GetSandboxes returns the Sandboxes field if non-nil, zero value otherwise.

### GetSandboxesOk

`func (o *SandboxList) GetSandboxesOk() (*[]Sandbox, bool)`

GetSandboxesOk returns a tuple with the Sandboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxes

`func (o *SandboxList) SetSandboxes(v []Sandbox)`

SetSandboxes sets Sandboxes field to given value.

### HasSandboxes

`func (o *SandboxList) HasSandboxes() bool`

HasSandboxes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



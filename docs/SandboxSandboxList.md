# SandboxSandboxList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sandboxes** | Pointer to [**[]SandboxSandbox**](SandboxSandbox.md) | Sandboxes are the sandboxes the caller holds that match the filter: the ones they leased, or every one in the org for an admin of it. Never null. | [optional] 

## Methods

### NewSandboxSandboxList

`func NewSandboxSandboxList() *SandboxSandboxList`

NewSandboxSandboxList instantiates a new SandboxSandboxList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxSandboxListWithDefaults

`func NewSandboxSandboxListWithDefaults() *SandboxSandboxList`

NewSandboxSandboxListWithDefaults instantiates a new SandboxSandboxList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSandboxes

`func (o *SandboxSandboxList) GetSandboxes() []SandboxSandbox`

GetSandboxes returns the Sandboxes field if non-nil, zero value otherwise.

### GetSandboxesOk

`func (o *SandboxSandboxList) GetSandboxesOk() (*[]SandboxSandbox, bool)`

GetSandboxesOk returns a tuple with the Sandboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxes

`func (o *SandboxSandboxList) SetSandboxes(v []SandboxSandbox)`

SetSandboxes sets Sandboxes field to given value.

### HasSandboxes

`func (o *SandboxSandboxList) HasSandboxes() bool`

HasSandboxes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AgentsRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to **string** | Caller input appended to the agent&#39;s instructions (capped at 128 KiB). | [optional] 

## Methods

### NewAgentsRunRequest

`func NewAgentsRunRequest() *AgentsRunRequest`

NewAgentsRunRequest instantiates a new AgentsRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsRunRequestWithDefaults

`func NewAgentsRunRequestWithDefaults() *AgentsRunRequest`

NewAgentsRunRequestWithDefaults instantiates a new AgentsRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *AgentsRunRequest) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *AgentsRunRequest) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *AgentsRunRequest) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *AgentsRunRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



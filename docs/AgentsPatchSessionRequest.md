# AgentsPatchSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewAgentsPatchSessionRequest

`func NewAgentsPatchSessionRequest() *AgentsPatchSessionRequest`

NewAgentsPatchSessionRequest instantiates a new AgentsPatchSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsPatchSessionRequestWithDefaults

`func NewAgentsPatchSessionRequestWithDefaults() *AgentsPatchSessionRequest`

NewAgentsPatchSessionRequestWithDefaults instantiates a new AgentsPatchSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AgentsPatchSessionRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentsPatchSessionRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentsPatchSessionRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentsPatchSessionRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *AgentsPatchSessionRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AgentsPatchSessionRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AgentsPatchSessionRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AgentsPatchSessionRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



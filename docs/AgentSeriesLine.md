# AgentSeriesLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | agent name | [optional] 
**Points** | Pointer to [**[]AgentSeriesPoint**](AgentSeriesPoint.md) | Points is one bucket per interval across the whole window, in time order and never sparse: a bucket with no runs is present with v 0, so two lines drawn from two agents share an x-axis without the client aligning anything. The window decides the count — 24 hourly for 24H, 7 daily, 30 daily. | [optional] 

## Methods

### NewAgentSeriesLine

`func NewAgentSeriesLine() *AgentSeriesLine`

NewAgentSeriesLine instantiates a new AgentSeriesLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentSeriesLineWithDefaults

`func NewAgentSeriesLineWithDefaults() *AgentSeriesLine`

NewAgentSeriesLineWithDefaults instantiates a new AgentSeriesLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AgentSeriesLine) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AgentSeriesLine) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AgentSeriesLine) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AgentSeriesLine) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPoints

`func (o *AgentSeriesLine) GetPoints() []AgentSeriesPoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *AgentSeriesLine) GetPointsOk() (*[]AgentSeriesPoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *AgentSeriesLine) SetPoints(v []AgentSeriesPoint)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *AgentSeriesLine) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



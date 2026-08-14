# SourceToolList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | Source is the source these tools came from. | [optional] 
**Tools** | Pointer to [**[]Tool**](Tool.md) | Tools is the caller&#39;s tools from that source. Never null. | [optional] 

## Methods

### NewSourceToolList

`func NewSourceToolList() *SourceToolList`

NewSourceToolList instantiates a new SourceToolList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceToolListWithDefaults

`func NewSourceToolListWithDefaults() *SourceToolList`

NewSourceToolListWithDefaults instantiates a new SourceToolList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *SourceToolList) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SourceToolList) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SourceToolList) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *SourceToolList) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTools

`func (o *SourceToolList) GetTools() []Tool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *SourceToolList) GetToolsOk() (*[]Tool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *SourceToolList) SetTools(v []Tool)`

SetTools sets Tools field to given value.

### HasTools

`func (o *SourceToolList) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



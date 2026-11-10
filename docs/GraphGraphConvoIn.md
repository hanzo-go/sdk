# GraphGraphConvoIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | **string** | At is the default valid instant for any turn that omits its own At, RFC 3339. Required. | 
**DryRun** | Pointer to **bool** | DryRun when true returns the extracted assertions without writing them to the store. Absent or false files them through the assertion admission path. | [optional] 
**Source** | **string** | Source names the conversation these turns came from, such as chat/thread-42. Required. | 
**Turns** | [**[]GraphGraphTurn**](GraphGraphTurn.md) | Turns is the ordered dialogue turns to read. Required. | 

## Methods

### NewGraphGraphConvoIn

`func NewGraphGraphConvoIn(at string, source string, turns []GraphGraphTurn, ) *GraphGraphConvoIn`

NewGraphGraphConvoIn instantiates a new GraphGraphConvoIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphConvoInWithDefaults

`func NewGraphGraphConvoInWithDefaults() *GraphGraphConvoIn`

NewGraphGraphConvoInWithDefaults instantiates a new GraphGraphConvoIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *GraphGraphConvoIn) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *GraphGraphConvoIn) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *GraphGraphConvoIn) SetAt(v string)`

SetAt sets At field to given value.


### GetDryRun

`func (o *GraphGraphConvoIn) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *GraphGraphConvoIn) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *GraphGraphConvoIn) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *GraphGraphConvoIn) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetSource

`func (o *GraphGraphConvoIn) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GraphGraphConvoIn) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GraphGraphConvoIn) SetSource(v string)`

SetSource sets Source field to given value.


### GetTurns

`func (o *GraphGraphConvoIn) GetTurns() []GraphGraphTurn`

GetTurns returns the Turns field if non-nil, zero value otherwise.

### GetTurnsOk

`func (o *GraphGraphConvoIn) GetTurnsOk() (*[]GraphGraphTurn, bool)`

GetTurnsOk returns a tuple with the Turns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurns

`func (o *GraphGraphConvoIn) SetTurns(v []GraphGraphTurn)`

SetTurns sets Turns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



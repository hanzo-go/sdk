# O11yO11yQueryRangePreviewIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompositeQuery** | Pointer to [**O11yQuerybuildertypesv5CompositeQuery**](O11yQuerybuildertypesv5CompositeQuery.md) |  | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**FormatOptions** | Pointer to [**O11yFormatOptions**](O11yFormatOptions.md) |  | [optional] 
**NoCache** | Pointer to **bool** |  | [optional] 
**RequestType** | Pointer to **interface{}** |  | [optional] 
**SchemaVersion** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Variables** | Pointer to [**map[string]O11yVariableItem**](O11yVariableItem.md) |  | [optional] 
**Verbose** | Pointer to **string** | Verbose selects the answer&#39;s depth. Empty or \&quot;true\&quot; renders the underlying Datastore SQL with EXPLAIN and granule analysis; \&quot;false\&quot; returns only the per-query valid/error/warnings verdict with no Datastore round trips. | [optional] 

## Methods

### NewO11yO11yQueryRangePreviewIn

`func NewO11yO11yQueryRangePreviewIn() *O11yO11yQueryRangePreviewIn`

NewO11yO11yQueryRangePreviewIn instantiates a new O11yO11yQueryRangePreviewIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yQueryRangePreviewInWithDefaults

`func NewO11yO11yQueryRangePreviewInWithDefaults() *O11yO11yQueryRangePreviewIn`

NewO11yO11yQueryRangePreviewInWithDefaults instantiates a new O11yO11yQueryRangePreviewIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompositeQuery

`func (o *O11yO11yQueryRangePreviewIn) GetCompositeQuery() O11yQuerybuildertypesv5CompositeQuery`

GetCompositeQuery returns the CompositeQuery field if non-nil, zero value otherwise.

### GetCompositeQueryOk

`func (o *O11yO11yQueryRangePreviewIn) GetCompositeQueryOk() (*O11yQuerybuildertypesv5CompositeQuery, bool)`

GetCompositeQueryOk returns a tuple with the CompositeQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeQuery

`func (o *O11yO11yQueryRangePreviewIn) SetCompositeQuery(v O11yQuerybuildertypesv5CompositeQuery)`

SetCompositeQuery sets CompositeQuery field to given value.

### HasCompositeQuery

`func (o *O11yO11yQueryRangePreviewIn) HasCompositeQuery() bool`

HasCompositeQuery returns a boolean if a field has been set.

### GetEnd

`func (o *O11yO11yQueryRangePreviewIn) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yO11yQueryRangePreviewIn) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yO11yQueryRangePreviewIn) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yO11yQueryRangePreviewIn) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFormatOptions

`func (o *O11yO11yQueryRangePreviewIn) GetFormatOptions() O11yFormatOptions`

GetFormatOptions returns the FormatOptions field if non-nil, zero value otherwise.

### GetFormatOptionsOk

`func (o *O11yO11yQueryRangePreviewIn) GetFormatOptionsOk() (*O11yFormatOptions, bool)`

GetFormatOptionsOk returns a tuple with the FormatOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatOptions

`func (o *O11yO11yQueryRangePreviewIn) SetFormatOptions(v O11yFormatOptions)`

SetFormatOptions sets FormatOptions field to given value.

### HasFormatOptions

`func (o *O11yO11yQueryRangePreviewIn) HasFormatOptions() bool`

HasFormatOptions returns a boolean if a field has been set.

### GetNoCache

`func (o *O11yO11yQueryRangePreviewIn) GetNoCache() bool`

GetNoCache returns the NoCache field if non-nil, zero value otherwise.

### GetNoCacheOk

`func (o *O11yO11yQueryRangePreviewIn) GetNoCacheOk() (*bool, bool)`

GetNoCacheOk returns a tuple with the NoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCache

`func (o *O11yO11yQueryRangePreviewIn) SetNoCache(v bool)`

SetNoCache sets NoCache field to given value.

### HasNoCache

`func (o *O11yO11yQueryRangePreviewIn) HasNoCache() bool`

HasNoCache returns a boolean if a field has been set.

### GetRequestType

`func (o *O11yO11yQueryRangePreviewIn) GetRequestType() interface{}`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *O11yO11yQueryRangePreviewIn) GetRequestTypeOk() (*interface{}, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *O11yO11yQueryRangePreviewIn) SetRequestType(v interface{})`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *O11yO11yQueryRangePreviewIn) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### SetRequestTypeNil

`func (o *O11yO11yQueryRangePreviewIn) SetRequestTypeNil(b bool)`

 SetRequestTypeNil sets the value for RequestType to be an explicit nil

### UnsetRequestType
`func (o *O11yO11yQueryRangePreviewIn) UnsetRequestType()`

UnsetRequestType ensures that no value is present for RequestType, not even an explicit nil
### GetSchemaVersion

`func (o *O11yO11yQueryRangePreviewIn) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *O11yO11yQueryRangePreviewIn) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *O11yO11yQueryRangePreviewIn) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *O11yO11yQueryRangePreviewIn) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetStart

`func (o *O11yO11yQueryRangePreviewIn) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yO11yQueryRangePreviewIn) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yO11yQueryRangePreviewIn) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yO11yQueryRangePreviewIn) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetVariables

`func (o *O11yO11yQueryRangePreviewIn) GetVariables() map[string]O11yVariableItem`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *O11yO11yQueryRangePreviewIn) GetVariablesOk() (*map[string]O11yVariableItem, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *O11yO11yQueryRangePreviewIn) SetVariables(v map[string]O11yVariableItem)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *O11yO11yQueryRangePreviewIn) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetVerbose

`func (o *O11yO11yQueryRangePreviewIn) GetVerbose() string`

GetVerbose returns the Verbose field if non-nil, zero value otherwise.

### GetVerboseOk

`func (o *O11yO11yQueryRangePreviewIn) GetVerboseOk() (*string, bool)`

GetVerboseOk returns a tuple with the Verbose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbose

`func (o *O11yO11yQueryRangePreviewIn) SetVerbose(v string)`

SetVerbose sets Verbose field to given value.

### HasVerbose

`func (o *O11yO11yQueryRangePreviewIn) HasVerbose() bool`

HasVerbose returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



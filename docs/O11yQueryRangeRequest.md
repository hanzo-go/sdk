# O11yQueryRangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompositeQuery** | Pointer to [**O11yQuerybuildertypesv5CompositeQuery**](O11yQuerybuildertypesv5CompositeQuery.md) | CompositeQuery is the composite query to use for the request. | [optional] 
**End** | Pointer to **int32** | End is the end time of the query in epoch milliseconds. | [optional] 
**FormatOptions** | Pointer to [**O11yFormatOptions**](O11yFormatOptions.md) |  | [optional] 
**NoCache** | Pointer to **bool** | NoCache is a flag to disable caching for the request. | [optional] 
**RequestType** | Pointer to **interface{}** |  | [optional] 
**SchemaVersion** | Pointer to **string** | SchemaVersion is the version of the schema to use for the request payload. | [optional] 
**Start** | Pointer to **int32** | Start is the start time of the query in epoch milliseconds. | [optional] 
**Variables** | Pointer to [**map[string]O11yVariableItem**](O11yVariableItem.md) | Variables is the variables to use for the request. | [optional] 

## Methods

### NewO11yQueryRangeRequest

`func NewO11yQueryRangeRequest() *O11yQueryRangeRequest`

NewO11yQueryRangeRequest instantiates a new O11yQueryRangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yQueryRangeRequestWithDefaults

`func NewO11yQueryRangeRequestWithDefaults() *O11yQueryRangeRequest`

NewO11yQueryRangeRequestWithDefaults instantiates a new O11yQueryRangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompositeQuery

`func (o *O11yQueryRangeRequest) GetCompositeQuery() O11yQuerybuildertypesv5CompositeQuery`

GetCompositeQuery returns the CompositeQuery field if non-nil, zero value otherwise.

### GetCompositeQueryOk

`func (o *O11yQueryRangeRequest) GetCompositeQueryOk() (*O11yQuerybuildertypesv5CompositeQuery, bool)`

GetCompositeQueryOk returns a tuple with the CompositeQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeQuery

`func (o *O11yQueryRangeRequest) SetCompositeQuery(v O11yQuerybuildertypesv5CompositeQuery)`

SetCompositeQuery sets CompositeQuery field to given value.

### HasCompositeQuery

`func (o *O11yQueryRangeRequest) HasCompositeQuery() bool`

HasCompositeQuery returns a boolean if a field has been set.

### GetEnd

`func (o *O11yQueryRangeRequest) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yQueryRangeRequest) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yQueryRangeRequest) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yQueryRangeRequest) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFormatOptions

`func (o *O11yQueryRangeRequest) GetFormatOptions() O11yFormatOptions`

GetFormatOptions returns the FormatOptions field if non-nil, zero value otherwise.

### GetFormatOptionsOk

`func (o *O11yQueryRangeRequest) GetFormatOptionsOk() (*O11yFormatOptions, bool)`

GetFormatOptionsOk returns a tuple with the FormatOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatOptions

`func (o *O11yQueryRangeRequest) SetFormatOptions(v O11yFormatOptions)`

SetFormatOptions sets FormatOptions field to given value.

### HasFormatOptions

`func (o *O11yQueryRangeRequest) HasFormatOptions() bool`

HasFormatOptions returns a boolean if a field has been set.

### GetNoCache

`func (o *O11yQueryRangeRequest) GetNoCache() bool`

GetNoCache returns the NoCache field if non-nil, zero value otherwise.

### GetNoCacheOk

`func (o *O11yQueryRangeRequest) GetNoCacheOk() (*bool, bool)`

GetNoCacheOk returns a tuple with the NoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCache

`func (o *O11yQueryRangeRequest) SetNoCache(v bool)`

SetNoCache sets NoCache field to given value.

### HasNoCache

`func (o *O11yQueryRangeRequest) HasNoCache() bool`

HasNoCache returns a boolean if a field has been set.

### GetRequestType

`func (o *O11yQueryRangeRequest) GetRequestType() interface{}`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *O11yQueryRangeRequest) GetRequestTypeOk() (*interface{}, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *O11yQueryRangeRequest) SetRequestType(v interface{})`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *O11yQueryRangeRequest) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### SetRequestTypeNil

`func (o *O11yQueryRangeRequest) SetRequestTypeNil(b bool)`

 SetRequestTypeNil sets the value for RequestType to be an explicit nil

### UnsetRequestType
`func (o *O11yQueryRangeRequest) UnsetRequestType()`

UnsetRequestType ensures that no value is present for RequestType, not even an explicit nil
### GetSchemaVersion

`func (o *O11yQueryRangeRequest) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *O11yQueryRangeRequest) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *O11yQueryRangeRequest) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *O11yQueryRangeRequest) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetStart

`func (o *O11yQueryRangeRequest) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yQueryRangeRequest) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yQueryRangeRequest) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yQueryRangeRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetVariables

`func (o *O11yQueryRangeRequest) GetVariables() map[string]O11yVariableItem`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *O11yQueryRangeRequest) GetVariablesOk() (*map[string]O11yVariableItem, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *O11yQueryRangeRequest) SetVariables(v map[string]O11yVariableItem)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *O11yQueryRangeRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



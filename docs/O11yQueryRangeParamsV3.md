# O11yQueryRangeParamsV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompositeQuery** | Pointer to [**O11yCompositeQuery**](O11yCompositeQuery.md) |  | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**FormatForWeb** | Pointer to **bool** |  | [optional] 
**NoCache** | Pointer to **bool** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Step** | Pointer to **int32** | step is in seconds; used for prometheus queries | [optional] 
**Variables** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewO11yQueryRangeParamsV3

`func NewO11yQueryRangeParamsV3() *O11yQueryRangeParamsV3`

NewO11yQueryRangeParamsV3 instantiates a new O11yQueryRangeParamsV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yQueryRangeParamsV3WithDefaults

`func NewO11yQueryRangeParamsV3WithDefaults() *O11yQueryRangeParamsV3`

NewO11yQueryRangeParamsV3WithDefaults instantiates a new O11yQueryRangeParamsV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompositeQuery

`func (o *O11yQueryRangeParamsV3) GetCompositeQuery() O11yCompositeQuery`

GetCompositeQuery returns the CompositeQuery field if non-nil, zero value otherwise.

### GetCompositeQueryOk

`func (o *O11yQueryRangeParamsV3) GetCompositeQueryOk() (*O11yCompositeQuery, bool)`

GetCompositeQueryOk returns a tuple with the CompositeQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeQuery

`func (o *O11yQueryRangeParamsV3) SetCompositeQuery(v O11yCompositeQuery)`

SetCompositeQuery sets CompositeQuery field to given value.

### HasCompositeQuery

`func (o *O11yQueryRangeParamsV3) HasCompositeQuery() bool`

HasCompositeQuery returns a boolean if a field has been set.

### GetEnd

`func (o *O11yQueryRangeParamsV3) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yQueryRangeParamsV3) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yQueryRangeParamsV3) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yQueryRangeParamsV3) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFormatForWeb

`func (o *O11yQueryRangeParamsV3) GetFormatForWeb() bool`

GetFormatForWeb returns the FormatForWeb field if non-nil, zero value otherwise.

### GetFormatForWebOk

`func (o *O11yQueryRangeParamsV3) GetFormatForWebOk() (*bool, bool)`

GetFormatForWebOk returns a tuple with the FormatForWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatForWeb

`func (o *O11yQueryRangeParamsV3) SetFormatForWeb(v bool)`

SetFormatForWeb sets FormatForWeb field to given value.

### HasFormatForWeb

`func (o *O11yQueryRangeParamsV3) HasFormatForWeb() bool`

HasFormatForWeb returns a boolean if a field has been set.

### GetNoCache

`func (o *O11yQueryRangeParamsV3) GetNoCache() bool`

GetNoCache returns the NoCache field if non-nil, zero value otherwise.

### GetNoCacheOk

`func (o *O11yQueryRangeParamsV3) GetNoCacheOk() (*bool, bool)`

GetNoCacheOk returns a tuple with the NoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCache

`func (o *O11yQueryRangeParamsV3) SetNoCache(v bool)`

SetNoCache sets NoCache field to given value.

### HasNoCache

`func (o *O11yQueryRangeParamsV3) HasNoCache() bool`

HasNoCache returns a boolean if a field has been set.

### GetStart

`func (o *O11yQueryRangeParamsV3) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yQueryRangeParamsV3) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yQueryRangeParamsV3) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yQueryRangeParamsV3) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStep

`func (o *O11yQueryRangeParamsV3) GetStep() int32`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *O11yQueryRangeParamsV3) GetStepOk() (*int32, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *O11yQueryRangeParamsV3) SetStep(v int32)`

SetStep sets Step field to given value.

### HasStep

`func (o *O11yQueryRangeParamsV3) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetVariables

`func (o *O11yQueryRangeParamsV3) GetVariables() map[string]map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *O11yQueryRangeParamsV3) GetVariablesOk() (*map[string]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *O11yQueryRangeParamsV3) SetVariables(v map[string]map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *O11yQueryRangeParamsV3) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



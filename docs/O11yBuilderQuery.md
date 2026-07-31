# O11yBuilderQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **int64** | Range start (Unix ms). | [optional] 
**End** | Pointer to **int64** | Range end (Unix ms). | [optional] 
**Step** | Pointer to **int32** | Step resolution in seconds. | [optional] 
**CompositeQuery** | Pointer to **map[string]interface{}** | Composite query definition (queryType, panelType, builderQueries). | [optional] 

## Methods

### NewO11yBuilderQuery

`func NewO11yBuilderQuery() *O11yBuilderQuery`

NewO11yBuilderQuery instantiates a new O11yBuilderQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yBuilderQueryWithDefaults

`func NewO11yBuilderQueryWithDefaults() *O11yBuilderQuery`

NewO11yBuilderQueryWithDefaults instantiates a new O11yBuilderQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *O11yBuilderQuery) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yBuilderQuery) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yBuilderQuery) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yBuilderQuery) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *O11yBuilderQuery) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yBuilderQuery) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yBuilderQuery) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yBuilderQuery) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStep

`func (o *O11yBuilderQuery) GetStep() int32`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *O11yBuilderQuery) GetStepOk() (*int32, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *O11yBuilderQuery) SetStep(v int32)`

SetStep sets Step field to given value.

### HasStep

`func (o *O11yBuilderQuery) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetCompositeQuery

`func (o *O11yBuilderQuery) GetCompositeQuery() map[string]interface{}`

GetCompositeQuery returns the CompositeQuery field if non-nil, zero value otherwise.

### GetCompositeQueryOk

`func (o *O11yBuilderQuery) GetCompositeQueryOk() (*map[string]interface{}, bool)`

GetCompositeQueryOk returns a tuple with the CompositeQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeQuery

`func (o *O11yBuilderQuery) SetCompositeQuery(v map[string]interface{})`

SetCompositeQuery sets CompositeQuery field to given value.

### HasCompositeQuery

`func (o *O11yBuilderQuery) HasCompositeQuery() bool`

HasCompositeQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IndexIndexQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **interface{}** |  | [optional] 
**Limit** | Pointer to **int64** | Limit is how many hits to return. Absent means 20; the ceiling is 1000. | [optional] 
**Offset** | Pointer to **int64** | Offset is where to start. Absent means 0. | [optional] 
**Q** | Pointer to **string** | Q is the search text. Typos are forgiven. An empty Q matches everything, which is how a client lists an index by relevance rather than by insertion order. | [optional] 

## Methods

### NewIndexIndexQuery

`func NewIndexIndexQuery() *IndexIndexQuery`

NewIndexIndexQuery instantiates a new IndexIndexQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexIndexQueryWithDefaults

`func NewIndexIndexQueryWithDefaults() *IndexIndexQuery`

NewIndexIndexQueryWithDefaults instantiates a new IndexIndexQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *IndexIndexQuery) GetFilter() interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *IndexIndexQuery) GetFilterOk() (*interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *IndexIndexQuery) SetFilter(v interface{})`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *IndexIndexQuery) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *IndexIndexQuery) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *IndexIndexQuery) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetLimit

`func (o *IndexIndexQuery) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *IndexIndexQuery) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *IndexIndexQuery) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *IndexIndexQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *IndexIndexQuery) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *IndexIndexQuery) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *IndexIndexQuery) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *IndexIndexQuery) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetQ

`func (o *IndexIndexQuery) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *IndexIndexQuery) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *IndexIndexQuery) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *IndexIndexQuery) HasQ() bool`

HasQ returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



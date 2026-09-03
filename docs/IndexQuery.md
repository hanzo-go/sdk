# IndexQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **interface{}** |  | [optional] 
**Limit** | Pointer to **int64** | Limit is how many hits to return. Absent means 20; the ceiling is 1000. | [optional] 
**Offset** | Pointer to **int64** | Offset is where to start. Absent means 0. | [optional] 
**Q** | Pointer to **string** | Q is the search text. Typos are forgiven. An empty Q matches everything, which is how a client lists an index by relevance rather than by insertion order. | [optional] 

## Methods

### NewIndexQuery

`func NewIndexQuery() *IndexQuery`

NewIndexQuery instantiates a new IndexQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexQueryWithDefaults

`func NewIndexQueryWithDefaults() *IndexQuery`

NewIndexQueryWithDefaults instantiates a new IndexQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *IndexQuery) GetFilter() interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *IndexQuery) GetFilterOk() (*interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *IndexQuery) SetFilter(v interface{})`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *IndexQuery) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *IndexQuery) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *IndexQuery) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetLimit

`func (o *IndexQuery) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *IndexQuery) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *IndexQuery) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *IndexQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *IndexQuery) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *IndexQuery) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *IndexQuery) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *IndexQuery) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetQ

`func (o *IndexQuery) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *IndexQuery) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *IndexQuery) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *IndexQuery) HasQ() bool`

HasQ returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



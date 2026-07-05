# SearchSettingsFaceting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxValuesPerFacet** | Pointer to **int32** |  | [optional] 
**SortFacetValuesBy** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSearchSettingsFaceting

`func NewSearchSettingsFaceting() *SearchSettingsFaceting`

NewSearchSettingsFaceting instantiates a new SearchSettingsFaceting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSettingsFacetingWithDefaults

`func NewSearchSettingsFacetingWithDefaults() *SearchSettingsFaceting`

NewSearchSettingsFacetingWithDefaults instantiates a new SearchSettingsFaceting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxValuesPerFacet

`func (o *SearchSettingsFaceting) GetMaxValuesPerFacet() int32`

GetMaxValuesPerFacet returns the MaxValuesPerFacet field if non-nil, zero value otherwise.

### GetMaxValuesPerFacetOk

`func (o *SearchSettingsFaceting) GetMaxValuesPerFacetOk() (*int32, bool)`

GetMaxValuesPerFacetOk returns a tuple with the MaxValuesPerFacet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValuesPerFacet

`func (o *SearchSettingsFaceting) SetMaxValuesPerFacet(v int32)`

SetMaxValuesPerFacet sets MaxValuesPerFacet field to given value.

### HasMaxValuesPerFacet

`func (o *SearchSettingsFaceting) HasMaxValuesPerFacet() bool`

HasMaxValuesPerFacet returns a boolean if a field has been set.

### GetSortFacetValuesBy

`func (o *SearchSettingsFaceting) GetSortFacetValuesBy() map[string]string`

GetSortFacetValuesBy returns the SortFacetValuesBy field if non-nil, zero value otherwise.

### GetSortFacetValuesByOk

`func (o *SearchSettingsFaceting) GetSortFacetValuesByOk() (*map[string]string, bool)`

GetSortFacetValuesByOk returns a tuple with the SortFacetValuesBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortFacetValuesBy

`func (o *SearchSettingsFaceting) SetSortFacetValuesBy(v map[string]string)`

SetSortFacetValuesBy sets SortFacetValuesBy field to given value.

### HasSortFacetValuesBy

`func (o *SearchSettingsFaceting) HasSortFacetValuesBy() bool`

HasSortFacetValuesBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



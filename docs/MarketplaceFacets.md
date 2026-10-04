# MarketplaceFacets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **map[string]int64** | Category counts listings in each category; an uncategorized one is not counted. | [optional] 
**Kind** | Pointer to **map[string]int64** | Kind counts listings of each kind. | [optional] 
**Price** | Pointer to **map[string]int64** | Price counts free and priced listings. | [optional] 
**Rating** | Pointer to **map[string]int64** | Rating counts listings by their rounded rating, \&quot;1\&quot; to \&quot;5\&quot;, and those with no rating yet as \&quot;unrated\&quot;. | [optional] 

## Methods

### NewMarketplaceFacets

`func NewMarketplaceFacets() *MarketplaceFacets`

NewMarketplaceFacets instantiates a new MarketplaceFacets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceFacetsWithDefaults

`func NewMarketplaceFacetsWithDefaults() *MarketplaceFacets`

NewMarketplaceFacetsWithDefaults instantiates a new MarketplaceFacets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *MarketplaceFacets) GetCategory() map[string]int64`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MarketplaceFacets) GetCategoryOk() (*map[string]int64, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MarketplaceFacets) SetCategory(v map[string]int64)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MarketplaceFacets) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetKind

`func (o *MarketplaceFacets) GetKind() map[string]int64`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MarketplaceFacets) GetKindOk() (*map[string]int64, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MarketplaceFacets) SetKind(v map[string]int64)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MarketplaceFacets) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPrice

`func (o *MarketplaceFacets) GetPrice() map[string]int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarketplaceFacets) GetPriceOk() (*map[string]int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarketplaceFacets) SetPrice(v map[string]int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MarketplaceFacets) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetRating

`func (o *MarketplaceFacets) GetRating() map[string]int64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *MarketplaceFacets) GetRatingOk() (*map[string]int64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *MarketplaceFacets) SetRating(v map[string]int64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *MarketplaceFacets) HasRating() bool`

HasRating returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Top

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** | End is the window&#39;s exclusive upper bound, RFC3339 UTC. | [optional] 
**Models** | Pointer to [**TopModels**](TopModels.md) | Models ranks the window&#39;s LLM models by spend — real per-org data. | [optional] 
**Products** | Pointer to [**TopProducts**](TopProducts.md) | Products ranks the window&#39;s products by revenue. | [optional] 
**Range** | Pointer to **string** | Range is the window that was actually applied: 24h, 7d, 30d or custom. | [optional] 
**Scope** | Pointer to [**Scope**](Scope.md) | Scope names the tenant these rankings belong to. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s inclusive lower bound, RFC3339 UTC. | [optional] 
**TopPages** | Pointer to [**Breakdown**](Breakdown.md) | Pages ranks the paths visitors requested, by pageviews. | [optional] 
**TopReferrers** | Pointer to [**Breakdown**](Breakdown.md) | Referrers ranks the external domains visitors arrived from, by pageviews. | [optional] 
**TopSources** | Pointer to [**Breakdown**](Breakdown.md) | Sources ranks the utm_source campaigns visitors arrived on, by pageviews. | [optional] 

## Methods

### NewTop

`func NewTop() *Top`

NewTop instantiates a new Top object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopWithDefaults

`func NewTopWithDefaults() *Top`

NewTopWithDefaults instantiates a new Top object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *Top) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Top) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Top) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Top) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetModels

`func (o *Top) GetModels() TopModels`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *Top) GetModelsOk() (*TopModels, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *Top) SetModels(v TopModels)`

SetModels sets Models field to given value.

### HasModels

`func (o *Top) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetProducts

`func (o *Top) GetProducts() TopProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *Top) GetProductsOk() (*TopProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *Top) SetProducts(v TopProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *Top) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetRange

`func (o *Top) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *Top) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *Top) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *Top) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScope

`func (o *Top) GetScope() Scope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Top) GetScopeOk() (*Scope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Top) SetScope(v Scope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Top) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStart

`func (o *Top) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Top) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Top) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *Top) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTopPages

`func (o *Top) GetTopPages() Breakdown`

GetTopPages returns the TopPages field if non-nil, zero value otherwise.

### GetTopPagesOk

`func (o *Top) GetTopPagesOk() (*Breakdown, bool)`

GetTopPagesOk returns a tuple with the TopPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopPages

`func (o *Top) SetTopPages(v Breakdown)`

SetTopPages sets TopPages field to given value.

### HasTopPages

`func (o *Top) HasTopPages() bool`

HasTopPages returns a boolean if a field has been set.

### GetTopReferrers

`func (o *Top) GetTopReferrers() Breakdown`

GetTopReferrers returns the TopReferrers field if non-nil, zero value otherwise.

### GetTopReferrersOk

`func (o *Top) GetTopReferrersOk() (*Breakdown, bool)`

GetTopReferrersOk returns a tuple with the TopReferrers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopReferrers

`func (o *Top) SetTopReferrers(v Breakdown)`

SetTopReferrers sets TopReferrers field to given value.

### HasTopReferrers

`func (o *Top) HasTopReferrers() bool`

HasTopReferrers returns a boolean if a field has been set.

### GetTopSources

`func (o *Top) GetTopSources() Breakdown`

GetTopSources returns the TopSources field if non-nil, zero value otherwise.

### GetTopSourcesOk

`func (o *Top) GetTopSourcesOk() (*Breakdown, bool)`

GetTopSourcesOk returns a tuple with the TopSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSources

`func (o *Top) SetTopSources(v Breakdown)`

SetTopSources sets TopSources field to given value.

### HasTopSources

`func (o *Top) HasTopSources() bool`

HasTopSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



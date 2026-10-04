# EventTop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** | End is the window&#39;s exclusive upper bound, RFC3339 UTC. | [optional] 
**Models** | Pointer to [**EventTopModels**](EventTopModels.md) | Models ranks the window&#39;s LLM models by spend — real per-org data. | [optional] 
**Products** | Pointer to [**EventTopProducts**](EventTopProducts.md) | Products ranks the window&#39;s products by revenue. | [optional] 
**Range** | Pointer to **string** | Range is the window that was actually applied: 24h, 7d, 30d or custom. | [optional] 
**Scope** | Pointer to [**EventScope**](EventScope.md) | Scope names the tenant these rankings belong to. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s inclusive lower bound, RFC3339 UTC. | [optional] 
**TopPages** | Pointer to [**EventBreakdown**](EventBreakdown.md) | Pages ranks the paths visitors requested, by pageviews. | [optional] 
**TopReferrers** | Pointer to [**EventBreakdown**](EventBreakdown.md) | Referrers ranks the external domains visitors arrived from, by pageviews. | [optional] 
**TopSources** | Pointer to [**EventBreakdown**](EventBreakdown.md) | Sources ranks the utm_source campaigns visitors arrived on, by pageviews. | [optional] 

## Methods

### NewEventTop

`func NewEventTop() *EventTop`

NewEventTop instantiates a new EventTop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTopWithDefaults

`func NewEventTopWithDefaults() *EventTop`

NewEventTopWithDefaults instantiates a new EventTop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *EventTop) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *EventTop) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *EventTop) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *EventTop) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetModels

`func (o *EventTop) GetModels() EventTopModels`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *EventTop) GetModelsOk() (*EventTopModels, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *EventTop) SetModels(v EventTopModels)`

SetModels sets Models field to given value.

### HasModels

`func (o *EventTop) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetProducts

`func (o *EventTop) GetProducts() EventTopProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *EventTop) GetProductsOk() (*EventTopProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *EventTop) SetProducts(v EventTopProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *EventTop) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetRange

`func (o *EventTop) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *EventTop) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *EventTop) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *EventTop) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScope

`func (o *EventTop) GetScope() EventScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EventTop) GetScopeOk() (*EventScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EventTop) SetScope(v EventScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *EventTop) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStart

`func (o *EventTop) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *EventTop) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *EventTop) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *EventTop) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTopPages

`func (o *EventTop) GetTopPages() EventBreakdown`

GetTopPages returns the TopPages field if non-nil, zero value otherwise.

### GetTopPagesOk

`func (o *EventTop) GetTopPagesOk() (*EventBreakdown, bool)`

GetTopPagesOk returns a tuple with the TopPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopPages

`func (o *EventTop) SetTopPages(v EventBreakdown)`

SetTopPages sets TopPages field to given value.

### HasTopPages

`func (o *EventTop) HasTopPages() bool`

HasTopPages returns a boolean if a field has been set.

### GetTopReferrers

`func (o *EventTop) GetTopReferrers() EventBreakdown`

GetTopReferrers returns the TopReferrers field if non-nil, zero value otherwise.

### GetTopReferrersOk

`func (o *EventTop) GetTopReferrersOk() (*EventBreakdown, bool)`

GetTopReferrersOk returns a tuple with the TopReferrers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopReferrers

`func (o *EventTop) SetTopReferrers(v EventBreakdown)`

SetTopReferrers sets TopReferrers field to given value.

### HasTopReferrers

`func (o *EventTop) HasTopReferrers() bool`

HasTopReferrers returns a boolean if a field has been set.

### GetTopSources

`func (o *EventTop) GetTopSources() EventBreakdown`

GetTopSources returns the TopSources field if non-nil, zero value otherwise.

### GetTopSourcesOk

`func (o *EventTop) GetTopSourcesOk() (*EventBreakdown, bool)`

GetTopSourcesOk returns a tuple with the TopSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSources

`func (o *EventTop) SetTopSources(v EventBreakdown)`

SetTopSources sets TopSources field to given value.

### HasTopSources

`func (o *EventTop) HasTopSources() bool`

HasTopSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



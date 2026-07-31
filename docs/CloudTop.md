# CloudTop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** | End is the window&#39;s exclusive upper bound, RFC3339 UTC. | [optional] 
**Models** | Pointer to [**CloudTopModels**](CloudTopModels.md) | Models ranks the window&#39;s LLM models by spend — real per-org data. | [optional] 
**Products** | Pointer to [**CloudTopProducts**](CloudTopProducts.md) | Products ranks the window&#39;s products by revenue. | [optional] 
**Range** | Pointer to **string** | Range is the window that was actually applied: 24h, 7d, 30d or custom. | [optional] 
**Scope** | Pointer to [**CloudScope**](CloudScope.md) | Scope names the tenant these rankings belong to. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s inclusive lower bound, RFC3339 UTC. | [optional] 
**TopPages** | Pointer to [**CloudBreakdown**](CloudBreakdown.md) | Pages ranks the paths visitors requested, by pageviews. | [optional] 
**TopReferrers** | Pointer to [**CloudBreakdown**](CloudBreakdown.md) | Referrers ranks the external domains visitors arrived from, by pageviews. | [optional] 
**TopSources** | Pointer to [**CloudBreakdown**](CloudBreakdown.md) | Sources ranks the utm_source campaigns visitors arrived on, by pageviews. | [optional] 

## Methods

### NewCloudTop

`func NewCloudTop() *CloudTop`

NewCloudTop instantiates a new CloudTop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTopWithDefaults

`func NewCloudTopWithDefaults() *CloudTop`

NewCloudTopWithDefaults instantiates a new CloudTop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *CloudTop) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CloudTop) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CloudTop) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CloudTop) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetModels

`func (o *CloudTop) GetModels() CloudTopModels`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *CloudTop) GetModelsOk() (*CloudTopModels, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *CloudTop) SetModels(v CloudTopModels)`

SetModels sets Models field to given value.

### HasModels

`func (o *CloudTop) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetProducts

`func (o *CloudTop) GetProducts() CloudTopProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CloudTop) GetProductsOk() (*CloudTopProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CloudTop) SetProducts(v CloudTopProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *CloudTop) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetRange

`func (o *CloudTop) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CloudTop) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CloudTop) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *CloudTop) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScope

`func (o *CloudTop) GetScope() CloudScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CloudTop) GetScopeOk() (*CloudScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CloudTop) SetScope(v CloudScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CloudTop) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStart

`func (o *CloudTop) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CloudTop) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CloudTop) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *CloudTop) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTopPages

`func (o *CloudTop) GetTopPages() CloudBreakdown`

GetTopPages returns the TopPages field if non-nil, zero value otherwise.

### GetTopPagesOk

`func (o *CloudTop) GetTopPagesOk() (*CloudBreakdown, bool)`

GetTopPagesOk returns a tuple with the TopPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopPages

`func (o *CloudTop) SetTopPages(v CloudBreakdown)`

SetTopPages sets TopPages field to given value.

### HasTopPages

`func (o *CloudTop) HasTopPages() bool`

HasTopPages returns a boolean if a field has been set.

### GetTopReferrers

`func (o *CloudTop) GetTopReferrers() CloudBreakdown`

GetTopReferrers returns the TopReferrers field if non-nil, zero value otherwise.

### GetTopReferrersOk

`func (o *CloudTop) GetTopReferrersOk() (*CloudBreakdown, bool)`

GetTopReferrersOk returns a tuple with the TopReferrers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopReferrers

`func (o *CloudTop) SetTopReferrers(v CloudBreakdown)`

SetTopReferrers sets TopReferrers field to given value.

### HasTopReferrers

`func (o *CloudTop) HasTopReferrers() bool`

HasTopReferrers returns a boolean if a field has been set.

### GetTopSources

`func (o *CloudTop) GetTopSources() CloudBreakdown`

GetTopSources returns the TopSources field if non-nil, zero value otherwise.

### GetTopSourcesOk

`func (o *CloudTop) GetTopSourcesOk() (*CloudBreakdown, bool)`

GetTopSourcesOk returns a tuple with the TopSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSources

`func (o *CloudTop) SetTopSources(v CloudBreakdown)`

SetTopSources sets TopSources field to given value.

### HasTopSources

`func (o *CloudTop) HasTopSources() bool`

HasTopSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



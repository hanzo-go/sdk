# MarketingPromoList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MarketingPromoStatus**](MarketingPromoStatus.md) | Data is every promo in the deployment, oldest first, each with its live counters. The list is fleet-wide rather than per-org. It is normally EMPTY: nothing seeds a promo, and the migration purges the one that once shipped by accident. | [optional] 

## Methods

### NewMarketingPromoList

`func NewMarketingPromoList() *MarketingPromoList`

NewMarketingPromoList instantiates a new MarketingPromoList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingPromoListWithDefaults

`func NewMarketingPromoListWithDefaults() *MarketingPromoList`

NewMarketingPromoListWithDefaults instantiates a new MarketingPromoList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MarketingPromoList) GetData() []MarketingPromoStatus`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MarketingPromoList) GetDataOk() (*[]MarketingPromoStatus, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MarketingPromoList) SetData(v []MarketingPromoStatus)`

SetData sets Data field to given value.

### HasData

`func (o *MarketingPromoList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



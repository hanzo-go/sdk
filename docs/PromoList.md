# PromoList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PromoStatus**](PromoStatus.md) | Data is every promo in the deployment, oldest first, each with its live counters. The list is fleet-wide rather than per-org. It is normally EMPTY: nothing seeds a promo, and the migration purges the one that once shipped by accident. | [optional] 

## Methods

### NewPromoList

`func NewPromoList() *PromoList`

NewPromoList instantiates a new PromoList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromoListWithDefaults

`func NewPromoListWithDefaults() *PromoList`

NewPromoListWithDefaults instantiates a new PromoList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PromoList) GetData() []PromoStatus`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PromoList) GetDataOk() (*[]PromoStatus, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PromoList) SetData(v []PromoStatus)`

SetData sets Data field to given value.

### HasData

`func (o *PromoList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Settlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affiliate** | Pointer to [**AdminAffiliateView**](AdminAffiliateView.md) | Affiliate is the row re-read AFTER the payout, so its paidCents and pendingCents already account for the row beside it. | [optional] 
**Payout** | Pointer to [**Remittance**](Remittance.md) | Payout is the payout row just recorded. | [optional] 

## Methods

### NewSettlement

`func NewSettlement() *Settlement`

NewSettlement instantiates a new Settlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementWithDefaults

`func NewSettlementWithDefaults() *Settlement`

NewSettlementWithDefaults instantiates a new Settlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffiliate

`func (o *Settlement) GetAffiliate() AdminAffiliateView`

GetAffiliate returns the Affiliate field if non-nil, zero value otherwise.

### GetAffiliateOk

`func (o *Settlement) GetAffiliateOk() (*AdminAffiliateView, bool)`

GetAffiliateOk returns a tuple with the Affiliate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliate

`func (o *Settlement) SetAffiliate(v AdminAffiliateView)`

SetAffiliate sets Affiliate field to given value.

### HasAffiliate

`func (o *Settlement) HasAffiliate() bool`

HasAffiliate returns a boolean if a field has been set.

### GetPayout

`func (o *Settlement) GetPayout() Remittance`

GetPayout returns the Payout field if non-nil, zero value otherwise.

### GetPayoutOk

`func (o *Settlement) GetPayoutOk() (*Remittance, bool)`

GetPayoutOk returns a tuple with the Payout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayout

`func (o *Settlement) SetPayout(v Remittance)`

SetPayout sets Payout field to given value.

### HasPayout

`func (o *Settlement) HasPayout() bool`

HasPayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



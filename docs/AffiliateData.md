# AffiliateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affiliate** | Pointer to [**AdminAffiliateView**](AdminAffiliateView.md) | Affiliate is the row as it stands AFTER the action that returned it. Its referredCount is 0 here: these single-affiliate answers do not run the count. | [optional] 

## Methods

### NewAffiliateData

`func NewAffiliateData() *AffiliateData`

NewAffiliateData instantiates a new AffiliateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliateDataWithDefaults

`func NewAffiliateDataWithDefaults() *AffiliateData`

NewAffiliateDataWithDefaults instantiates a new AffiliateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffiliate

`func (o *AffiliateData) GetAffiliate() AdminAffiliateView`

GetAffiliate returns the Affiliate field if non-nil, zero value otherwise.

### GetAffiliateOk

`func (o *AffiliateData) GetAffiliateOk() (*AdminAffiliateView, bool)`

GetAffiliateOk returns a tuple with the Affiliate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliate

`func (o *AffiliateData) SetAffiliate(v AdminAffiliateView)`

SetAffiliate sets Affiliate field to given value.

### HasAffiliate

`func (o *AffiliateData) HasAffiliate() bool`

HasAffiliate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



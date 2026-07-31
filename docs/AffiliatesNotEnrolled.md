# AffiliatesNotEnrolled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAffiliate** | Pointer to **bool** |  | [optional] 
**DefaultRateBps** | Pointer to **int64** | The default commission rate a new affiliate gets, in basis points (2000 &#x3D; 20%). | [optional] 

## Methods

### NewAffiliatesNotEnrolled

`func NewAffiliatesNotEnrolled() *AffiliatesNotEnrolled`

NewAffiliatesNotEnrolled instantiates a new AffiliatesNotEnrolled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliatesNotEnrolledWithDefaults

`func NewAffiliatesNotEnrolledWithDefaults() *AffiliatesNotEnrolled`

NewAffiliatesNotEnrolledWithDefaults instantiates a new AffiliatesNotEnrolled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAffiliate

`func (o *AffiliatesNotEnrolled) GetIsAffiliate() bool`

GetIsAffiliate returns the IsAffiliate field if non-nil, zero value otherwise.

### GetIsAffiliateOk

`func (o *AffiliatesNotEnrolled) GetIsAffiliateOk() (*bool, bool)`

GetIsAffiliateOk returns a tuple with the IsAffiliate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAffiliate

`func (o *AffiliatesNotEnrolled) SetIsAffiliate(v bool)`

SetIsAffiliate sets IsAffiliate field to given value.

### HasIsAffiliate

`func (o *AffiliatesNotEnrolled) HasIsAffiliate() bool`

HasIsAffiliate returns a boolean if a field has been set.

### GetDefaultRateBps

`func (o *AffiliatesNotEnrolled) GetDefaultRateBps() int64`

GetDefaultRateBps returns the DefaultRateBps field if non-nil, zero value otherwise.

### GetDefaultRateBpsOk

`func (o *AffiliatesNotEnrolled) GetDefaultRateBpsOk() (*int64, bool)`

GetDefaultRateBpsOk returns a tuple with the DefaultRateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRateBps

`func (o *AffiliatesNotEnrolled) SetDefaultRateBps(v int64)`

SetDefaultRateBps sets DefaultRateBps field to given value.

### HasDefaultRateBps

`func (o *AffiliatesNotEnrolled) HasDefaultRateBps() bool`

HasDefaultRateBps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



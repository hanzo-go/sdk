# AffiliateLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAffiliate** | Pointer to **bool** | IsAffiliate says whether the caller org has an affiliate record. On false only maxLinks comes back — there are no links, and there is no link to mint until the org applies and is approved. | [optional] 
**Links** | Pointer to [**[]CodeView**](CodeView.md) | Links is the caller&#39;s share links, each with its URL and funnel. | [optional] 
**MaxLinks** | Pointer to **int32** | MaxLinks is how many share links one affiliate may hold. | [optional] 
**Status** | Pointer to **string** | Status is the caller&#39;s affiliate status: \&quot;applied\&quot;, \&quot;approved\&quot; or \&quot;suspended\&quot;; absent for a non-affiliate. Minting a link requires \&quot;approved\&quot;, because a link that cannot accrue quietly loses the referral. | [optional] 

## Methods

### NewAffiliateLinks

`func NewAffiliateLinks() *AffiliateLinks`

NewAffiliateLinks instantiates a new AffiliateLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliateLinksWithDefaults

`func NewAffiliateLinksWithDefaults() *AffiliateLinks`

NewAffiliateLinksWithDefaults instantiates a new AffiliateLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAffiliate

`func (o *AffiliateLinks) GetIsAffiliate() bool`

GetIsAffiliate returns the IsAffiliate field if non-nil, zero value otherwise.

### GetIsAffiliateOk

`func (o *AffiliateLinks) GetIsAffiliateOk() (*bool, bool)`

GetIsAffiliateOk returns a tuple with the IsAffiliate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAffiliate

`func (o *AffiliateLinks) SetIsAffiliate(v bool)`

SetIsAffiliate sets IsAffiliate field to given value.

### HasIsAffiliate

`func (o *AffiliateLinks) HasIsAffiliate() bool`

HasIsAffiliate returns a boolean if a field has been set.

### GetLinks

`func (o *AffiliateLinks) GetLinks() []CodeView`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AffiliateLinks) GetLinksOk() (*[]CodeView, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AffiliateLinks) SetLinks(v []CodeView)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AffiliateLinks) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMaxLinks

`func (o *AffiliateLinks) GetMaxLinks() int32`

GetMaxLinks returns the MaxLinks field if non-nil, zero value otherwise.

### GetMaxLinksOk

`func (o *AffiliateLinks) GetMaxLinksOk() (*int32, bool)`

GetMaxLinksOk returns a tuple with the MaxLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLinks

`func (o *AffiliateLinks) SetMaxLinks(v int32)`

SetMaxLinks sets MaxLinks field to given value.

### HasMaxLinks

`func (o *AffiliateLinks) HasMaxLinks() bool`

HasMaxLinks returns a boolean if a field has been set.

### GetStatus

`func (o *AffiliateLinks) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AffiliateLinks) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AffiliateLinks) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AffiliateLinks) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



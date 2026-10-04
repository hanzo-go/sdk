# MarketplaceDispute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int64** | At is when, unix seconds. | [optional] 
**By** | Pointer to **string** | By is the org that disputed. | [optional] 
**Contested** | Pointer to **bool** | Contested is the clock&#39;s own dispute: the payment stopped clearing when the review window closed. A party&#39;s dispute never carries it, whatever the party is named. | [optional] 
**Reason** | Pointer to **string** | Reason is the disputing party&#39;s words. | [optional] 

## Methods

### NewMarketplaceDispute

`func NewMarketplaceDispute() *MarketplaceDispute`

NewMarketplaceDispute instantiates a new MarketplaceDispute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceDisputeWithDefaults

`func NewMarketplaceDisputeWithDefaults() *MarketplaceDispute`

NewMarketplaceDisputeWithDefaults instantiates a new MarketplaceDispute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *MarketplaceDispute) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *MarketplaceDispute) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *MarketplaceDispute) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *MarketplaceDispute) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBy

`func (o *MarketplaceDispute) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *MarketplaceDispute) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *MarketplaceDispute) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *MarketplaceDispute) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetContested

`func (o *MarketplaceDispute) GetContested() bool`

GetContested returns the Contested field if non-nil, zero value otherwise.

### GetContestedOk

`func (o *MarketplaceDispute) GetContestedOk() (*bool, bool)`

GetContestedOk returns a tuple with the Contested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContested

`func (o *MarketplaceDispute) SetContested(v bool)`

SetContested sets Contested field to given value.

### HasContested

`func (o *MarketplaceDispute) HasContested() bool`

HasContested returns a boolean if a field has been set.

### GetReason

`func (o *MarketplaceDispute) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MarketplaceDispute) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MarketplaceDispute) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MarketplaceDispute) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MarketplaceDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int64** | At is when, unix seconds. | [optional] 
**Hash** | Pointer to **string** | Hash commits to what was delivered: the seller&#39;s own 32-byte hash, or the SHA-256 of the note and URL. It is the escrow&#39;s delivery hash on chain. | [optional] 
**Note** | Pointer to **string** | Note says what was delivered. | [optional] 
**Url** | Pointer to **string** | URL is where it is, https. | [optional] 

## Methods

### NewMarketplaceDelivery

`func NewMarketplaceDelivery() *MarketplaceDelivery`

NewMarketplaceDelivery instantiates a new MarketplaceDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceDeliveryWithDefaults

`func NewMarketplaceDeliveryWithDefaults() *MarketplaceDelivery`

NewMarketplaceDeliveryWithDefaults instantiates a new MarketplaceDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *MarketplaceDelivery) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *MarketplaceDelivery) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *MarketplaceDelivery) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *MarketplaceDelivery) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetHash

`func (o *MarketplaceDelivery) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *MarketplaceDelivery) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *MarketplaceDelivery) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *MarketplaceDelivery) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetNote

`func (o *MarketplaceDelivery) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *MarketplaceDelivery) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *MarketplaceDelivery) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *MarketplaceDelivery) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetUrl

`func (o *MarketplaceDelivery) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MarketplaceDelivery) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MarketplaceDelivery) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MarketplaceDelivery) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MarketplaceSeller

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documented** | Pointer to **bool** | Documented is true while the seller&#39;s tax form is certified and valid — what its agents&#39; TaxPrincipalCredentials rest on. False as well when that could not be asked just now. | [optional] 
**Org** | Pointer to **string** | Org is the seller org. | [optional] 
**Reputation** | Pointer to [**MarketplaceReputation**](MarketplaceReputation.md) | Reputation is what buyers made of the seller, across its listings. | [optional] 

## Methods

### NewMarketplaceSeller

`func NewMarketplaceSeller() *MarketplaceSeller`

NewMarketplaceSeller instantiates a new MarketplaceSeller object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceSellerWithDefaults

`func NewMarketplaceSellerWithDefaults() *MarketplaceSeller`

NewMarketplaceSellerWithDefaults instantiates a new MarketplaceSeller object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumented

`func (o *MarketplaceSeller) GetDocumented() bool`

GetDocumented returns the Documented field if non-nil, zero value otherwise.

### GetDocumentedOk

`func (o *MarketplaceSeller) GetDocumentedOk() (*bool, bool)`

GetDocumentedOk returns a tuple with the Documented field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumented

`func (o *MarketplaceSeller) SetDocumented(v bool)`

SetDocumented sets Documented field to given value.

### HasDocumented

`func (o *MarketplaceSeller) HasDocumented() bool`

HasDocumented returns a boolean if a field has been set.

### GetOrg

`func (o *MarketplaceSeller) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *MarketplaceSeller) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *MarketplaceSeller) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *MarketplaceSeller) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetReputation

`func (o *MarketplaceSeller) GetReputation() MarketplaceReputation`

GetReputation returns the Reputation field if non-nil, zero value otherwise.

### GetReputationOk

`func (o *MarketplaceSeller) GetReputationOk() (*MarketplaceReputation, bool)`

GetReputationOk returns a tuple with the Reputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputation

`func (o *MarketplaceSeller) SetReputation(v MarketplaceReputation)`

SetReputation sets Reputation field to given value.

### HasReputation

`func (o *MarketplaceSeller) HasReputation() bool`

HasReputation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



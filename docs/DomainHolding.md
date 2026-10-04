# DomainHolding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostCents** | Pointer to **int64** | wholesale cost | [optional] 
**Domain** | Pointer to **string** | the name owned | [optional] 
**ExpiresAt** | Pointer to **string** | when the registration lapses, RFC3339 | [optional] 
**Nameservers** | Pointer to **[]string** | the authoritative nameservers the name points at | [optional] 
**Order** | Pointer to **int64** | registrar order id | [optional] 
**Org** | Pointer to **string** | the org that owns the domain | [optional] 
**PriceCents** | Pointer to **int64** | what the customer paid (sell) | [optional] 
**RegisteredAt** | Pointer to **int64** | unix seconds | [optional] 

## Methods

### NewDomainHolding

`func NewDomainHolding() *DomainHolding`

NewDomainHolding instantiates a new DomainHolding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainHoldingWithDefaults

`func NewDomainHoldingWithDefaults() *DomainHolding`

NewDomainHoldingWithDefaults instantiates a new DomainHolding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostCents

`func (o *DomainHolding) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *DomainHolding) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *DomainHolding) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *DomainHolding) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetDomain

`func (o *DomainHolding) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainHolding) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainHolding) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DomainHolding) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetExpiresAt

`func (o *DomainHolding) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DomainHolding) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DomainHolding) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DomainHolding) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetNameservers

`func (o *DomainHolding) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *DomainHolding) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *DomainHolding) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *DomainHolding) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetOrder

`func (o *DomainHolding) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *DomainHolding) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *DomainHolding) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *DomainHolding) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrg

`func (o *DomainHolding) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *DomainHolding) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *DomainHolding) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *DomainHolding) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPriceCents

`func (o *DomainHolding) GetPriceCents() int64`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *DomainHolding) GetPriceCentsOk() (*int64, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *DomainHolding) SetPriceCents(v int64)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *DomainHolding) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.

### GetRegisteredAt

`func (o *DomainHolding) GetRegisteredAt() int64`

GetRegisteredAt returns the RegisteredAt field if non-nil, zero value otherwise.

### GetRegisteredAtOk

`func (o *DomainHolding) GetRegisteredAtOk() (*int64, bool)`

GetRegisteredAtOk returns a tuple with the RegisteredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAt

`func (o *DomainHolding) SetRegisteredAt(v int64)`

SetRegisteredAt sets RegisteredAt field to given value.

### HasRegisteredAt

`func (o *DomainHolding) HasRegisteredAt() bool`

HasRegisteredAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



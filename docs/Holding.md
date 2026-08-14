# Holding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostCents** | Pointer to **int32** | wholesale cost | [optional] 
**Domain** | Pointer to **string** | the name owned | [optional] 
**ExpiresAt** | Pointer to **string** | when the registration lapses, RFC3339 | [optional] 
**Nameservers** | Pointer to **[]string** | the authoritative nameservers the name points at | [optional] 
**Order** | Pointer to **int32** | registrar order id | [optional] 
**Org** | Pointer to **string** | the org that owns the domain | [optional] 
**PriceCents** | Pointer to **int32** | what the customer paid (sell) | [optional] 
**RegisteredAt** | Pointer to **int32** | unix seconds | [optional] 

## Methods

### NewHolding

`func NewHolding() *Holding`

NewHolding instantiates a new Holding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldingWithDefaults

`func NewHoldingWithDefaults() *Holding`

NewHoldingWithDefaults instantiates a new Holding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostCents

`func (o *Holding) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *Holding) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *Holding) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *Holding) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetDomain

`func (o *Holding) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Holding) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Holding) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Holding) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Holding) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Holding) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Holding) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Holding) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetNameservers

`func (o *Holding) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *Holding) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *Holding) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *Holding) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetOrder

`func (o *Holding) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Holding) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Holding) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Holding) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrg

`func (o *Holding) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Holding) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Holding) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Holding) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPriceCents

`func (o *Holding) GetPriceCents() int32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *Holding) GetPriceCentsOk() (*int32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *Holding) SetPriceCents(v int32)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *Holding) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.

### GetRegisteredAt

`func (o *Holding) GetRegisteredAt() int32`

GetRegisteredAt returns the RegisteredAt field if non-nil, zero value otherwise.

### GetRegisteredAtOk

`func (o *Holding) GetRegisteredAtOk() (*int32, bool)`

GetRegisteredAtOk returns a tuple with the RegisteredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAt

`func (o *Holding) SetRegisteredAt(v int32)`

SetRegisteredAt sets RegisteredAt field to given value.

### HasRegisteredAt

`func (o *Holding) HasRegisteredAt() bool`

HasRegisteredAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# EventEconomicEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** | Agent is the principal that acted for the buyer, as IAM names it. Only the buyer&#39;s copy carries it: it may name a person, and who inside the buyer acted is none of the seller&#39;s records. | [optional] 
**At** | Pointer to **int64** | At is when the payment was made, unix seconds, as the rail recorded it. | [optional] 
**Buyer** | Pointer to **string** | Buyer is the org that paid. | [optional] 
**Category** | Pointer to **string** | Category is what the payment was for — service, goods, transfer, refund or royalty — as the rail recorded the parties agreeing it. Empty on an event stated before version 2, and on a rail that records no purpose. | [optional] 
**Chain** | Pointer to **string** | Chain is the CAIP-2 network an on-chain payment moved on. | [optional] 
**Corrects** | Pointer to **string** | Corrects is the id of the event this one restates, when it is a correction. | [optional] 
**Currency** | Pointer to **string** | Currency is what the amounts are in: an ISO 4217 code or an asset symbol. | [optional] 
**Fee** | Pointer to **string** | Fee is what the rail kept of Gross, an exact decimal in Currency. | [optional] 
**Fmv** | Pointer to **string** | FMV is Gross&#39;s fair market value in U.S. dollars when it was paid, an exact decimal. | [optional] 
**Gross** | Pointer to **string** | Gross is what the buyer paid, an exact decimal in Currency. | [optional] 
**Id** | Pointer to **string** | ID is the event&#39;s identity, derived from its rail and proof. | [optional] 
**Jurisdictions** | Pointer to **[]string** | Jurisdictions are the ISO 3166 codes the payment is subject to, when the rail knows them. | [optional] 
**Parent** | Pointer to **string** | Parent is the agent that spawned Agent, when one did. Only the buyer&#39;s copy carries it, for Agent&#39;s reason. | [optional] 
**Proof** | Pointer to **string** | Proof is the rail&#39;s own reference for the payment: a settlement id, a chain transaction hash or a processor reference. | [optional] 
**Rail** | Pointer to **string** | Rail is how the money moved: x402, ledger, chain, ach, wire, card, network or other. | [optional] 
**Seller** | Pointer to **string** | Seller is the org that was paid. | [optional] 
**Service** | Pointer to **string** | Service is what was bought, in the rail&#39;s own words. | [optional] 
**Tax** | Pointer to **string** | Tax is the tax collected within Gross, an exact decimal in Currency. | [optional] 
**Tx** | Pointer to **string** | Tx is the on-chain transaction hash of an on-chain payment. | [optional] 
**V** | Pointer to **int64** | V is the schema version the event was stated in. | [optional] 

## Methods

### NewEventEconomicEvent

`func NewEventEconomicEvent() *EventEconomicEvent`

NewEventEconomicEvent instantiates a new EventEconomicEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEconomicEventWithDefaults

`func NewEventEconomicEventWithDefaults() *EventEconomicEvent`

NewEventEconomicEventWithDefaults instantiates a new EventEconomicEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *EventEconomicEvent) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *EventEconomicEvent) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *EventEconomicEvent) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *EventEconomicEvent) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetAt

`func (o *EventEconomicEvent) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *EventEconomicEvent) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *EventEconomicEvent) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *EventEconomicEvent) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBuyer

`func (o *EventEconomicEvent) GetBuyer() string`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *EventEconomicEvent) GetBuyerOk() (*string, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *EventEconomicEvent) SetBuyer(v string)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *EventEconomicEvent) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetCategory

`func (o *EventEconomicEvent) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EventEconomicEvent) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EventEconomicEvent) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EventEconomicEvent) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetChain

`func (o *EventEconomicEvent) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *EventEconomicEvent) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *EventEconomicEvent) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *EventEconomicEvent) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetCorrects

`func (o *EventEconomicEvent) GetCorrects() string`

GetCorrects returns the Corrects field if non-nil, zero value otherwise.

### GetCorrectsOk

`func (o *EventEconomicEvent) GetCorrectsOk() (*string, bool)`

GetCorrectsOk returns a tuple with the Corrects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrects

`func (o *EventEconomicEvent) SetCorrects(v string)`

SetCorrects sets Corrects field to given value.

### HasCorrects

`func (o *EventEconomicEvent) HasCorrects() bool`

HasCorrects returns a boolean if a field has been set.

### GetCurrency

`func (o *EventEconomicEvent) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *EventEconomicEvent) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *EventEconomicEvent) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *EventEconomicEvent) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetFee

`func (o *EventEconomicEvent) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *EventEconomicEvent) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *EventEconomicEvent) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *EventEconomicEvent) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFmv

`func (o *EventEconomicEvent) GetFmv() string`

GetFmv returns the Fmv field if non-nil, zero value otherwise.

### GetFmvOk

`func (o *EventEconomicEvent) GetFmvOk() (*string, bool)`

GetFmvOk returns a tuple with the Fmv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFmv

`func (o *EventEconomicEvent) SetFmv(v string)`

SetFmv sets Fmv field to given value.

### HasFmv

`func (o *EventEconomicEvent) HasFmv() bool`

HasFmv returns a boolean if a field has been set.

### GetGross

`func (o *EventEconomicEvent) GetGross() string`

GetGross returns the Gross field if non-nil, zero value otherwise.

### GetGrossOk

`func (o *EventEconomicEvent) GetGrossOk() (*string, bool)`

GetGrossOk returns a tuple with the Gross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGross

`func (o *EventEconomicEvent) SetGross(v string)`

SetGross sets Gross field to given value.

### HasGross

`func (o *EventEconomicEvent) HasGross() bool`

HasGross returns a boolean if a field has been set.

### GetId

`func (o *EventEconomicEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventEconomicEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventEconomicEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventEconomicEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJurisdictions

`func (o *EventEconomicEvent) GetJurisdictions() []string`

GetJurisdictions returns the Jurisdictions field if non-nil, zero value otherwise.

### GetJurisdictionsOk

`func (o *EventEconomicEvent) GetJurisdictionsOk() (*[]string, bool)`

GetJurisdictionsOk returns a tuple with the Jurisdictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdictions

`func (o *EventEconomicEvent) SetJurisdictions(v []string)`

SetJurisdictions sets Jurisdictions field to given value.

### HasJurisdictions

`func (o *EventEconomicEvent) HasJurisdictions() bool`

HasJurisdictions returns a boolean if a field has been set.

### GetParent

`func (o *EventEconomicEvent) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *EventEconomicEvent) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *EventEconomicEvent) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *EventEconomicEvent) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProof

`func (o *EventEconomicEvent) GetProof() string`

GetProof returns the Proof field if non-nil, zero value otherwise.

### GetProofOk

`func (o *EventEconomicEvent) GetProofOk() (*string, bool)`

GetProofOk returns a tuple with the Proof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProof

`func (o *EventEconomicEvent) SetProof(v string)`

SetProof sets Proof field to given value.

### HasProof

`func (o *EventEconomicEvent) HasProof() bool`

HasProof returns a boolean if a field has been set.

### GetRail

`func (o *EventEconomicEvent) GetRail() string`

GetRail returns the Rail field if non-nil, zero value otherwise.

### GetRailOk

`func (o *EventEconomicEvent) GetRailOk() (*string, bool)`

GetRailOk returns a tuple with the Rail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRail

`func (o *EventEconomicEvent) SetRail(v string)`

SetRail sets Rail field to given value.

### HasRail

`func (o *EventEconomicEvent) HasRail() bool`

HasRail returns a boolean if a field has been set.

### GetSeller

`func (o *EventEconomicEvent) GetSeller() string`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *EventEconomicEvent) GetSellerOk() (*string, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *EventEconomicEvent) SetSeller(v string)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *EventEconomicEvent) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetService

`func (o *EventEconomicEvent) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *EventEconomicEvent) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *EventEconomicEvent) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *EventEconomicEvent) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTax

`func (o *EventEconomicEvent) GetTax() string`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *EventEconomicEvent) GetTaxOk() (*string, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *EventEconomicEvent) SetTax(v string)`

SetTax sets Tax field to given value.

### HasTax

`func (o *EventEconomicEvent) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTx

`func (o *EventEconomicEvent) GetTx() string`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *EventEconomicEvent) GetTxOk() (*string, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *EventEconomicEvent) SetTx(v string)`

SetTx sets Tx field to given value.

### HasTx

`func (o *EventEconomicEvent) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetV

`func (o *EventEconomicEvent) GetV() int64`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *EventEconomicEvent) GetVOk() (*int64, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *EventEconomicEvent) SetV(v int64)`

SetV sets V field to given value.

### HasV

`func (o *EventEconomicEvent) HasV() bool`

HasV returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



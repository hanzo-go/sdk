# PricingComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basis** | Pointer to **string** | Basis is what it is metered by, in the fewest words that are true. | [optional] 
**Name** | Pointer to **string** | Name is the component&#39;s name on the spend breakdown. | [optional] 
**Note** | Pointer to **string** | Note is what the component means, in one sentence. | [optional] 
**Quoted** | Pointer to **bool** | Quoted says whether a call is priced BEFORE it runs. A quoted component can refuse a call that would breach a budget; a measured one is booked from what it used and is bounded by the headroom under the per-task cap instead. | [optional] 
**Tiers** | Pointer to **[]string** | Tiers are the disjoint tiers this component meters in, where it has them. | [optional] 
**Title** | Pointer to **string** | Title is how it reads on a price list. | [optional] 

## Methods

### NewPricingComponent

`func NewPricingComponent() *PricingComponent`

NewPricingComponent instantiates a new PricingComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingComponentWithDefaults

`func NewPricingComponentWithDefaults() *PricingComponent`

NewPricingComponentWithDefaults instantiates a new PricingComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasis

`func (o *PricingComponent) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *PricingComponent) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *PricingComponent) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *PricingComponent) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### GetName

`func (o *PricingComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PricingComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PricingComponent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PricingComponent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNote

`func (o *PricingComponent) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *PricingComponent) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *PricingComponent) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *PricingComponent) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetQuoted

`func (o *PricingComponent) GetQuoted() bool`

GetQuoted returns the Quoted field if non-nil, zero value otherwise.

### GetQuotedOk

`func (o *PricingComponent) GetQuotedOk() (*bool, bool)`

GetQuotedOk returns a tuple with the Quoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoted

`func (o *PricingComponent) SetQuoted(v bool)`

SetQuoted sets Quoted field to given value.

### HasQuoted

`func (o *PricingComponent) HasQuoted() bool`

HasQuoted returns a boolean if a field has been set.

### GetTiers

`func (o *PricingComponent) GetTiers() []string`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *PricingComponent) GetTiersOk() (*[]string, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *PricingComponent) SetTiers(v []string)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *PricingComponent) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetTitle

`func (o *PricingComponent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PricingComponent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PricingComponent) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PricingComponent) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



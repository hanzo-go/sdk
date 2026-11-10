# PricingUserEnablementItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effective** | Pointer to **bool** | Effective is whether the caller may use the item right now, which is the field to branch on: true for any ga item, for a beta item while the platform&#39;s &#x60;beta&#x60; switch is on for the caller, and never for an off one. | [optional] 
**Id** | Pointer to **string** | ID is the item within that namespace — a model id, a provider name, or a feature key. | [optional] 
**Kind** | Pointer to **string** | Kind is the namespace the id lives in: \&quot;model\&quot;, \&quot;provider\&quot; or \&quot;feature\&quot;. | [optional] 
**State** | Pointer to **string** | State is the item&#39;s GLOBAL availability — \&quot;off\&quot;, \&quot;beta\&quot; or \&quot;ga\&quot; — which is the operator&#39;s setting and not this caller&#39;s answer. Effective is that. | [optional] 

## Methods

### NewPricingUserEnablementItem

`func NewPricingUserEnablementItem() *PricingUserEnablementItem`

NewPricingUserEnablementItem instantiates a new PricingUserEnablementItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingUserEnablementItemWithDefaults

`func NewPricingUserEnablementItemWithDefaults() *PricingUserEnablementItem`

NewPricingUserEnablementItemWithDefaults instantiates a new PricingUserEnablementItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffective

`func (o *PricingUserEnablementItem) GetEffective() bool`

GetEffective returns the Effective field if non-nil, zero value otherwise.

### GetEffectiveOk

`func (o *PricingUserEnablementItem) GetEffectiveOk() (*bool, bool)`

GetEffectiveOk returns a tuple with the Effective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffective

`func (o *PricingUserEnablementItem) SetEffective(v bool)`

SetEffective sets Effective field to given value.

### HasEffective

`func (o *PricingUserEnablementItem) HasEffective() bool`

HasEffective returns a boolean if a field has been set.

### GetId

`func (o *PricingUserEnablementItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PricingUserEnablementItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PricingUserEnablementItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PricingUserEnablementItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *PricingUserEnablementItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PricingUserEnablementItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PricingUserEnablementItem) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PricingUserEnablementItem) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetState

`func (o *PricingUserEnablementItem) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PricingUserEnablementItem) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PricingUserEnablementItem) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PricingUserEnablementItem) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



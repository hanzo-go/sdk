# Promo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Active is false for a promo that is no longer offered; an inactive promo quotes as ineligible and refuses to redeem. | [optional] 
**Code** | Pointer to **string** | Code is the promo id, e.g. \&quot;first1000\&quot;. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is unix seconds. | [optional] 
**Description** | Pointer to **string** | Description is the human-readable offer. | [optional] 
**MaxRedemptions** | Pointer to **int64** | MaxRedemptions is the hard fleet-wide cap; the redemption past it is declined. | [optional] 
**PercentOff** | Pointer to **int64** | PercentOff is the discount applied to ONE month&#39;s list price. | [optional] 
**Plans** | Pointer to **string** | Plans is the csv of eligible plan ids (\&quot;pro,max,team\&quot;). | [optional] 
**TeamSeatCap** | Pointer to **int64** | TeamSeatCap is how many Team seats bill at the promo rate; seats beyond it bill at list. | [optional] 

## Methods

### NewPromo

`func NewPromo() *Promo`

NewPromo instantiates a new Promo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromoWithDefaults

`func NewPromoWithDefaults() *Promo`

NewPromoWithDefaults instantiates a new Promo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *Promo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Promo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Promo) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Promo) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCode

`func (o *Promo) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Promo) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Promo) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Promo) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Promo) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Promo) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Promo) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Promo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Promo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Promo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Promo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Promo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxRedemptions

`func (o *Promo) GetMaxRedemptions() int64`

GetMaxRedemptions returns the MaxRedemptions field if non-nil, zero value otherwise.

### GetMaxRedemptionsOk

`func (o *Promo) GetMaxRedemptionsOk() (*int64, bool)`

GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptions

`func (o *Promo) SetMaxRedemptions(v int64)`

SetMaxRedemptions sets MaxRedemptions field to given value.

### HasMaxRedemptions

`func (o *Promo) HasMaxRedemptions() bool`

HasMaxRedemptions returns a boolean if a field has been set.

### GetPercentOff

`func (o *Promo) GetPercentOff() int64`

GetPercentOff returns the PercentOff field if non-nil, zero value otherwise.

### GetPercentOffOk

`func (o *Promo) GetPercentOffOk() (*int64, bool)`

GetPercentOffOk returns a tuple with the PercentOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOff

`func (o *Promo) SetPercentOff(v int64)`

SetPercentOff sets PercentOff field to given value.

### HasPercentOff

`func (o *Promo) HasPercentOff() bool`

HasPercentOff returns a boolean if a field has been set.

### GetPlans

`func (o *Promo) GetPlans() string`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *Promo) GetPlansOk() (*string, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *Promo) SetPlans(v string)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *Promo) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetTeamSeatCap

`func (o *Promo) GetTeamSeatCap() int64`

GetTeamSeatCap returns the TeamSeatCap field if non-nil, zero value otherwise.

### GetTeamSeatCapOk

`func (o *Promo) GetTeamSeatCapOk() (*int64, bool)`

GetTeamSeatCapOk returns a tuple with the TeamSeatCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamSeatCap

`func (o *Promo) SetTeamSeatCap(v int64)`

SetTeamSeatCap sets TeamSeatCap field to given value.

### HasTeamSeatCap

`func (o *Promo) HasTeamSeatCap() bool`

HasTeamSeatCap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CloudPromo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Active is false for a promo that is no longer offered; an inactive promo quotes as ineligible and refuses to redeem. | [optional] 
**Code** | Pointer to **string** | Code is the promo id, e.g. \&quot;first1000\&quot;. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is unix seconds. | [optional] 
**Description** | Pointer to **string** | Description is the human-readable offer. | [optional] 
**MaxRedemptions** | Pointer to **int32** | MaxRedemptions is the hard fleet-wide cap; the redemption past it is declined. | [optional] 
**PercentOff** | Pointer to **int32** | PercentOff is the discount applied to ONE month&#39;s list price. | [optional] 
**Plans** | Pointer to **string** | Plans is the csv of eligible plan ids (\&quot;pro,max,team\&quot;). | [optional] 
**TeamSeatCap** | Pointer to **int32** | TeamSeatCap is how many Team seats bill at the promo rate; seats beyond it bill at list. | [optional] 

## Methods

### NewCloudPromo

`func NewCloudPromo() *CloudPromo`

NewCloudPromo instantiates a new CloudPromo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPromoWithDefaults

`func NewCloudPromoWithDefaults() *CloudPromo`

NewCloudPromoWithDefaults instantiates a new CloudPromo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CloudPromo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CloudPromo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CloudPromo) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CloudPromo) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCode

`func (o *CloudPromo) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CloudPromo) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CloudPromo) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CloudPromo) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudPromo) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudPromo) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudPromo) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudPromo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *CloudPromo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudPromo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudPromo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudPromo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxRedemptions

`func (o *CloudPromo) GetMaxRedemptions() int32`

GetMaxRedemptions returns the MaxRedemptions field if non-nil, zero value otherwise.

### GetMaxRedemptionsOk

`func (o *CloudPromo) GetMaxRedemptionsOk() (*int32, bool)`

GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptions

`func (o *CloudPromo) SetMaxRedemptions(v int32)`

SetMaxRedemptions sets MaxRedemptions field to given value.

### HasMaxRedemptions

`func (o *CloudPromo) HasMaxRedemptions() bool`

HasMaxRedemptions returns a boolean if a field has been set.

### GetPercentOff

`func (o *CloudPromo) GetPercentOff() int32`

GetPercentOff returns the PercentOff field if non-nil, zero value otherwise.

### GetPercentOffOk

`func (o *CloudPromo) GetPercentOffOk() (*int32, bool)`

GetPercentOffOk returns a tuple with the PercentOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOff

`func (o *CloudPromo) SetPercentOff(v int32)`

SetPercentOff sets PercentOff field to given value.

### HasPercentOff

`func (o *CloudPromo) HasPercentOff() bool`

HasPercentOff returns a boolean if a field has been set.

### GetPlans

`func (o *CloudPromo) GetPlans() string`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *CloudPromo) GetPlansOk() (*string, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *CloudPromo) SetPlans(v string)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *CloudPromo) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetTeamSeatCap

`func (o *CloudPromo) GetTeamSeatCap() int32`

GetTeamSeatCap returns the TeamSeatCap field if non-nil, zero value otherwise.

### GetTeamSeatCapOk

`func (o *CloudPromo) GetTeamSeatCapOk() (*int32, bool)`

GetTeamSeatCapOk returns a tuple with the TeamSeatCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamSeatCap

`func (o *CloudPromo) SetTeamSeatCap(v int32)`

SetTeamSeatCap sets TeamSeatCap field to given value.

### HasTeamSeatCap

`func (o *CloudPromo) HasTeamSeatCap() bool`

HasTeamSeatCap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



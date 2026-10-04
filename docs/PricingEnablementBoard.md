# PricingEnablementBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Betas** | Pointer to [**[]PricingUserEnablementItem**](PricingUserEnablementItem.md) | Betas are the subset of Items the caller&#39;s org may still opt into. | [optional] 
**Items** | Pointer to [**[]PricingUserEnablementItem**](PricingUserEnablementItem.md) | Items is every managed item, each resolved for the caller&#39;s org. | [optional] 
**Org** | Pointer to **string** | Org is the org this view was resolved for; empty for a caller with no validated principal, who sees only the generally-available items. | [optional] 

## Methods

### NewPricingEnablementBoard

`func NewPricingEnablementBoard() *PricingEnablementBoard`

NewPricingEnablementBoard instantiates a new PricingEnablementBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingEnablementBoardWithDefaults

`func NewPricingEnablementBoardWithDefaults() *PricingEnablementBoard`

NewPricingEnablementBoardWithDefaults instantiates a new PricingEnablementBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBetas

`func (o *PricingEnablementBoard) GetBetas() []PricingUserEnablementItem`

GetBetas returns the Betas field if non-nil, zero value otherwise.

### GetBetasOk

`func (o *PricingEnablementBoard) GetBetasOk() (*[]PricingUserEnablementItem, bool)`

GetBetasOk returns a tuple with the Betas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetas

`func (o *PricingEnablementBoard) SetBetas(v []PricingUserEnablementItem)`

SetBetas sets Betas field to given value.

### HasBetas

`func (o *PricingEnablementBoard) HasBetas() bool`

HasBetas returns a boolean if a field has been set.

### GetItems

`func (o *PricingEnablementBoard) GetItems() []PricingUserEnablementItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PricingEnablementBoard) GetItemsOk() (*[]PricingUserEnablementItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PricingEnablementBoard) SetItems(v []PricingUserEnablementItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *PricingEnablementBoard) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOrg

`func (o *PricingEnablementBoard) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PricingEnablementBoard) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PricingEnablementBoard) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PricingEnablementBoard) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



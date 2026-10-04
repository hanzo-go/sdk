# BillingCreditGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**AmountCents** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**EffectiveAt** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int64** |  | [optional] 
**RemainingCents** | Pointer to **int64** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Voided** | Pointer to **bool** |  | [optional] 

## Methods

### NewBillingCreditGrant

`func NewBillingCreditGrant() *BillingCreditGrant`

NewBillingCreditGrant instantiates a new BillingCreditGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCreditGrantWithDefaults

`func NewBillingCreditGrantWithDefaults() *BillingCreditGrant`

NewBillingCreditGrantWithDefaults instantiates a new BillingCreditGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *BillingCreditGrant) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BillingCreditGrant) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BillingCreditGrant) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *BillingCreditGrant) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAmountCents

`func (o *BillingCreditGrant) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *BillingCreditGrant) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *BillingCreditGrant) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *BillingCreditGrant) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BillingCreditGrant) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingCreditGrant) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingCreditGrant) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BillingCreditGrant) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingCreditGrant) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingCreditGrant) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingCreditGrant) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingCreditGrant) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEffectiveAt

`func (o *BillingCreditGrant) GetEffectiveAt() string`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *BillingCreditGrant) GetEffectiveAtOk() (*string, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *BillingCreditGrant) SetEffectiveAt(v string)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *BillingCreditGrant) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *BillingCreditGrant) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *BillingCreditGrant) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *BillingCreditGrant) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *BillingCreditGrant) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *BillingCreditGrant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingCreditGrant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingCreditGrant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingCreditGrant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BillingCreditGrant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingCreditGrant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingCreditGrant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingCreditGrant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *BillingCreditGrant) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BillingCreditGrant) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BillingCreditGrant) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *BillingCreditGrant) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRemainingCents

`func (o *BillingCreditGrant) GetRemainingCents() int64`

GetRemainingCents returns the RemainingCents field if non-nil, zero value otherwise.

### GetRemainingCentsOk

`func (o *BillingCreditGrant) GetRemainingCentsOk() (*int64, bool)`

GetRemainingCentsOk returns a tuple with the RemainingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCents

`func (o *BillingCreditGrant) SetRemainingCents(v int64)`

SetRemainingCents sets RemainingCents field to given value.

### HasRemainingCents

`func (o *BillingCreditGrant) HasRemainingCents() bool`

HasRemainingCents returns a boolean if a field has been set.

### GetTags

`func (o *BillingCreditGrant) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BillingCreditGrant) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BillingCreditGrant) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BillingCreditGrant) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUserId

`func (o *BillingCreditGrant) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BillingCreditGrant) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BillingCreditGrant) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BillingCreditGrant) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVoided

`func (o *BillingCreditGrant) GetVoided() bool`

GetVoided returns the Voided field if non-nil, zero value otherwise.

### GetVoidedOk

`func (o *BillingCreditGrant) GetVoidedOk() (*bool, bool)`

GetVoidedOk returns a tuple with the Voided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoided

`func (o *BillingCreditGrant) SetVoided(v bool)`

SetVoided sets Voided field to given value.

### HasVoided

`func (o *BillingCreditGrant) HasVoided() bool`

HasVoided returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


